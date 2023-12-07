package processor

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bitspawngg/bitspawn-api/config"
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services/queue"
	"github.com/bitspawngg/bitspawn-api/services/tournament"
	"github.com/bitspawngg/bitspawn-api/utils/tasks"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type TxProcessor struct {
	DB   *models.DB
	log  *logrus.Entry
	conf config.Config

	SQSClient *queue.SqsClient
	tsvc      *tournament.TournamentService

	q chan int
}

func NewTxProcessor(db *models.DB, log *logrus.Logger, conf config.AwsConfig, tsvc *tournament.TournamentService, sqsSvc *queue.SQSService) *TxProcessor {

	tp := &TxProcessor{
		DB:        db,
		log:       log.WithField("service", "processor"),
		conf:      config.Config{},
		SQSClient: sqsSvc.Client(conf.SQSNameTx, queue.WithDelay(0)),
		tsvc:      tsvc,
		q:         make(chan int),
	}
	tp.Start()
	return tp
}

func (tp *TxProcessor) Start() {
	ticker := time.NewTicker(20 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				// handle pull
				tp.handlePull()
			case <-tp.q:
				ticker.Stop()
				return
			}
		}
	}()
}

func (tp *TxProcessor) Close() error {
	tp.q <- 0
	return nil
}

func (tp *TxProcessor) handlePull() {
	err := tp.SQSClient.LongPull(tp.AddJob)
	if err != nil {
		tp.log.Warn("no data pull from queue, ", err.Error())
		//todo: add mechanism to increase the ticker skip
	}
}

func (tp *TxProcessor) AddJob(message *sqs.Message) {
	//////To check message type
	meta := &tournament.Meta{}
	err := json.Unmarshal([]byte(*message.Body), meta)
	if err != nil {
		tp.log.Error("error in unmarshal message body: " + err.Error())
		return
	}
	id, err := uuid.FromString(meta.Author)
	if err != nil {
		tp.log.Error("error in parse uuid : " + err.Error())
		return
	}
	_ = tasks.GlobalTaskPool.Job(id, tp.Handle(meta, message))
}

func (tp *TxProcessor) Handle(meta *tournament.Meta, msg *sqs.Message) func() {
	return func() {
		err := tp.tsvc.ExecuteStateMachine(meta.TournamentID, meta.Author, tournament.Action(meta.Action), meta)
		if err != nil {
			tp.log.Error("error in ExecuteStateMachine: ", err)
			return
		}
		err = tp.SQSClient.DeleteMsg(msg)
		if err != nil {
			tp.log.Error("error in DeleteMsg: ", err)
			return
		}
		tournamentInfo, err := tp.DB.GetTournamentData(meta.TournamentID)
		if err != nil {
			tp.log.Error("error in GetTournamentData for tournament ", meta.TournamentID, ": ", err)
			return
		}
		note := models.Notification{
			Icon:     tournamentInfo.ThumbnailUrl,
			Keyword:  tournamentInfo.TournamentName,
			Link:     "/tournament/fetch/" + meta.TournamentID,
			Message:  "You have successfully " + meta.Action + "ed Tournament " + tournamentInfo.TournamentName,
			Type:     "Tournament",
			Username: meta.Author,
		}
		_ = tp.DB.CreateNotification(&note)
	}

}

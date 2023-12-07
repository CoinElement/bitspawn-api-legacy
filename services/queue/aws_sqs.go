package queue

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bitspawngg/bitspawn-api/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type SQSService struct {
	log  *logrus.Entry
	sess *session.Session
	conf config.AwsConfig
}

type SqsClient struct {
	log       *logrus.Entry
	sqsSvc    *sqs.SQS
	queueName string
	delay     *int64
}

func NewSQSService(log *logrus.Logger, conf config.AwsConfig) *SQSService {
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))

	svr := &SQSService{
		sess: sess,
		log:  log.WithField("services", "bitspawn_poa"),
		conf: conf,
	}
	return svr
}

func (s *SQSService) Client(queueName string, opts ...Option) *SqsClient {
	fields := s.log.Data
	fields["queue"] = queueName
	defaultDelay := int64(0)
	cli := &SqsClient{
		log:       s.log.WithFields(fields),
		sqsSvc:    sqs.New(s.sess),
		queueName: queueName,
		delay:     &defaultDelay,
	}
	for _, opt := range opts {
		opt(cli)
	}
	return cli
}

func (c *SqsClient) SendMsg(msg Message) error {
	urlResp, err := c.sqsSvc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &c.queueName,
	})
	if err != nil {
		c.log.Error(err)
		return errors.Wrap(err, "failed to get queue url: ")
	}
	queueURL := urlResp.QueueUrl

	result, err := c.sqsSvc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: c.delay,
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(msg.Title),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(msg.Author),
			},
			////// to add Attribute type
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String(msg.Body),
		QueueUrl:    queueURL,
	})

	if err != nil {
		c.log.Error(err)
		return err
	}

	c.log.Info("Success", *result.MessageId)

	return nil
}

func (c *SqsClient) GetUrl() *string {
	urlResp, err := c.sqsSvc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &c.queueName,
	})
	if err != nil {
		c.log.Error(err)
	}
	return urlResp.QueueUrl
}

func (c *SqsClient) ShortPollSqs(chn chan<- *sqs.Message) error {
	title := "Title"
	author := "Author"
	output, err := c.sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:              c.GetUrl(),
		MessageAttributeNames: []*string{&title, &author},
	})
	if err != nil {
		return fmt.Errorf("failed to fetch sqs message: %v", err)
	}

	c.log.Infof("pulled %d messages\n", len(output.Messages))
	for _, message := range output.Messages {
		chn <- message
	}

	return nil
}

func (c *SqsClient) LongPull(fn func(message *sqs.Message)) error {
	title := "Title"
	author := "Author"
	waitTime := int64(20)
	go func() {
		output, err := c.sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:              c.GetUrl(),
			MessageAttributeNames: []*string{&title, &author},
			WaitTimeSeconds:       &waitTime,
		})
		if err != nil {
			c.log.Warn("no msg pulled, ", err.Error())
		}
		c.log.Infof("pulled %d messages\n", len(output.Messages))
		for _, message := range output.Messages {
			fn(message)
		}
	}()
	return nil
}

func (c *SqsClient) DeleteMsg(msg *sqs.Message) error {
	_, err := c.sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      c.GetUrl(),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		return fmt.Errorf("failed to delete sqs message: %v", err)
	}

	return nil
}

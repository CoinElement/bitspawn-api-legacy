/*

 */

package challenge

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bitspawngg/bitspawn-api/services/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services"
	"github.com/bitspawngg/bitspawn-api/services/hdkey"
	"github.com/bitspawngg/bitspawn-api/services/poa"
)

type ChallengeService struct {
	DB  *models.DB
	log *logrus.Entry
	bsc *poa.BitspawnPoaClient

	conf *config.Service
}

type Meta struct {
	ChallengeID string `json:"challengeId" binding:"required"`
	ClubName    string `json:"clubName"`
	Members     []int  `json:"members"`
	TipAmount   int    `json:"tipAmount"`
	Platform    string `json:"platform"`
	PlayerID    string `json:"playerId"`
}

func NewChallengeService(db *models.DB, log *logrus.Logger, bsc *poa.BitspawnPoaClient, service *config.Service) *ChallengeService {
	return &ChallengeService{
		DB:  db,
		log: log.WithField("service", "challenge"),
		bsc: bsc,

		conf: service,
	}
}

func (tsvc *ChallengeService) GetAdminUser() string {
	return tsvc.conf.GetConfig().AdminUserID
}

func (tsvc *ChallengeService) CreateStateMachine(challenge *models.Challenge, user *models.UserAccount) error {
	err := tsvc.checkEth(user)
	if err != nil {
		tsvc.log.Error("error in checkEth: ", err)
		return fmt.Errorf("error in checkEth: %v", err)
	}
	userPrivateKey, err := getPrivateKey(user)
	if err != nil {
		tsvc.log.Error("error in getPrivateKey: ", err)
		return fmt.Errorf("error in getPrivateKey: %v", err)
	}

	ef := services.ConvertEthToWei(challenge.EntryFee)
	address, _, _, err := tsvc.bsc.DeployTournament(userPrivateKey, challenge.FeeType, 0, 0, ef, nil)

	if err != nil {
		tsvc.log.Error("error in deploying challenge ", challenge.ChallengeID, err)
		return fmt.Errorf("error in deploying challenge: %v", err)
	}

	challenge.ContractAddress = address.String()
	err = tsvc.DB.CreateChallengeData(challenge)
	if err != nil {
		tsvc.log.Error("error saving new challenge data to db: ", err)
		return fmt.Errorf("error saving new challenge data to db: %v", err)
	}

	return nil
}

func (tsvc *ChallengeService) ExecuteStateMachine(challengeId, username, action string, meta *Meta) error {
	challenge, err := tsvc.DB.GetChallengeData(challengeId)
	if err != nil {
		tsvc.log.Error("error in GetChallengeData for challenge ", challengeId, ": ", err)
		return fmt.Errorf("error in GetChallengeData for challenge %s: %v", challengeId, err)
	}
	err = tsvc.UpdateStatusFromBlockchain(challenge)
	if err != nil {
		tsvc.log.Error("error in Update Status from blockchain for challenge ", challengeId, ": ", err)
	}
	user, err := tsvc.DB.GetUserProfile(username)
	if err != nil {
		tsvc.log.Error("error in GetUser ", username, ": ", err)
		return fmt.Errorf("error in GetUser %s: %v", username, err)
	}
	err = tsvc.checkEth(user)
	if err != nil {
		tsvc.log.Error("error in checkEth: ", err)
		return fmt.Errorf("error in checkEth: %v", err)
	}
	state := challenge.Status
	switch action {
	case "Cancel":
		err = tsvc.Cancel(challenge, user, meta)
	case "Complete":
		err = tsvc.Complete(challenge, user, meta)
	case "Fund":
		err = tsvc.Fund(challenge, user, meta)
	case "Register":
		err = tsvc.Register(challenge, user, meta)
	default:
		err = fmt.Errorf("Action %s is not allowed in the current state %s", action, state)
	}
	if err != nil {
		tsvc.log.Error("error in ExecuteStateMachine for challenge ", challengeId, ": ", err.Error())
		return fmt.Errorf("error in ExecuteStateMachine for challenge %s: %v", challengeId, err)
	}

	return nil
}

func (tsvc *ChallengeService) UpdateStatusFromBlockchain(challenge *models.Challenge) error {
	status, err := tsvc.bsc.GetTournamentStatus(challenge.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting Challenge Status: %v", err)
	}
	if status == "Ongoing" { // do not change, use existing value in DB
		status = challenge.Status
	}

	err = tsvc.DB.UpdateChallengeState(challenge.ChallengeID, status)
	if err != nil {
		return fmt.Errorf("error updating challenge state to db")
	}
	return nil
}

func (tsvc *ChallengeService) Cancel(challenge *models.Challenge, user *models.UserAccount, meta *Meta) error {
	statusOnChain, err := tsvc.bsc.GetTournamentStatus(challenge.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting challenge Status: %v", err)
	}
	if statusOnChain == "Completed" {
		// challenge is already completed, bypass Cancel action
		return nil
	}

	if user.Username != challenge.OrganizerID {
		return fmt.Errorf("challenge not accessible by user")
	}
	userPrivateKey, err := getPrivateKey(user)
	if err != nil {
		tsvc.log.Error("error in getPrivateKey: ", err)
		return fmt.Errorf("error in getPrivateKey: %v", err)
	}

	contractAddress := common.HexToAddress(challenge.ContractAddress)
	_, err = tsvc.bsc.CancelTournament(userPrivateKey, contractAddress)
	if err != nil {
		return fmt.Errorf("error in CancelChallenge: %v", err)
	}

	_ = tsvc.DB.UpdateChallengeState(challenge.ChallengeID, "Cancelled")

	return nil
}

func (tsvc *ChallengeService) Complete(challenge *models.Challenge, user *models.UserAccount, meta *Meta) error {
	if strings.ToUpper(challenge.Status) != "READY" {
		return fmt.Errorf("challenge not in Ready status")
	}
	if time.Now().Before(challenge.CutoffDate.Add(time.Duration(challenge.MinuteTimeWindow) * time.Minute)) {
		return fmt.Errorf("challenge not yet past deadline")
	}
	if user.Username != challenge.OrganizerID {
		return fmt.Errorf("challenge not accessible by user")
	}
	userPrivateKey, err := getPrivateKey(user)
	if err != nil {
		tsvc.log.Error("error in getPrivateKey: ", err)
		return fmt.Errorf("error in getPrivateKey: %v", err)
	}

	challengeRecords, err := tsvc.DB.ReportChallengeWinners(challenge)
	if err != nil {
		return fmt.Errorf("error in ReportChallengeWinners: %v", err)
	}
	finalPlacementsAddress := make([]common.Address, 0, challenge.NumberOfWinners)
	for i := 0; i < int(challenge.NumberOfWinners); i++ {
		finalPlacementsAddress = append(finalPlacementsAddress, common.HexToAddress(challengeRecords[i].PublicAddress))
	}

	prizePercentage := strings.Split(challenge.PrizeAllocation, ",")
	prizePerTenThousand := []*big.Int{}
	for _, pp := range prizePercentage {
		ppInt, _ := strconv.ParseInt(pp, 10, 64)
		prizePerTenThousand = append(prizePerTenThousand, big.NewInt(ppInt*100)) // scale by 100
	}

	contractAddress := common.HexToAddress(challenge.ContractAddress)
	_, err = tsvc.bsc.CompleteTournament(userPrivateKey, contractAddress, finalPlacementsAddress, prizePerTenThousand)
	if err != nil {
		return fmt.Errorf("error in CompleteChallenge: %v", err)
	}

	_ = tsvc.DB.UpdateChallengeState(challenge.ChallengeID, "Completed")
	return nil
}

func (tsvc *ChallengeService) Fund(challenge *models.Challenge, user *models.UserAccount, meta *Meta) error {
	tipAmount := meta.TipAmount
	statusOnChain, err := tsvc.bsc.GetTournamentStatus(challenge.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting Challenge Status: %v", err)
	}
	if statusOnChain == "Completed" {
		// tournament is already completed, bypass fund action
		return nil
	}

	userPrivateKey, err := getPrivateKey(user)
	if err != nil {
		tsvc.log.Error("error in getPrivateKey: ", err)
		return fmt.Errorf("error in getPrivateKey: %v", err)
	}

	contractAddress := common.HexToAddress(challenge.ContractAddress)
	fundAmount := services.ConvertEthToWei(strconv.Itoa(tipAmount))
	_, err = tsvc.bsc.FundTournament(userPrivateKey, challenge.FeeType, contractAddress, fundAmount)
	if err != nil {
		return fmt.Errorf("error in Fund Challenge: %v", err)
	}

	_ = tsvc.DB.UpdateChallengeOrganizerContribute(challenge.ChallengeID, int64(tipAmount))

	return nil
}

func (tsvc *ChallengeService) Register(challenge *models.Challenge, user *models.UserAccount, meta *Meta) error {
	tipAmount := meta.TipAmount
	if challenge.Status != "Registration" && challenge.Status != "Ready" {
		return fmt.Errorf("challenge not in Registration or Ready status")
	}
	if challenge.EntryOnce {
		records, err := tsvc.DB.GetChallengeRecordsByUser(challenge.ChallengeID.String(), user.Sub)
		if err != nil {
			return fmt.Errorf("cannot decide whether user has registered or not: %v", err)
		} else if len(records) > 0 {
			return fmt.Errorf("user cannot register in this challenge more than once")
		}
	}
	userPrivateKey, err := getPrivateKey(user)
	if err != nil {
		tsvc.log.Error("error in getPrivateKey: ", err)
		return fmt.Errorf("error in getPrivateKey: %v", err)
	}

	challengeRecord := models.ChallengeRecord{
		UserId:            user.Sub,
		ChallengeId:       challenge.ChallengeID.String(),
		Platform:          meta.Platform,
		PlayerId:          meta.PlayerID,
		PublicAddress:     user.PublicAddress,
		RegisterDate:      challenge.StartDate,
		ChallengeDeadline: challenge.CutoffDate,
		ChallengeType:     challenge.ChallengeType,
		Scoring:           challenge.Scoring,
	}
	err = tsvc.DB.InsertChallengeRecord(&challengeRecord)
	if err != nil {
		return fmt.Errorf("error inserting challenge record: %v", err)
	}

	if challenge.EntryFee != "0" {
		contractAddress := common.HexToAddress(challenge.ContractAddress)
		minEntryFee, _ := strconv.Atoi(challenge.EntryFee)
		entryFeeWithTip := services.ConvertEthToWei(strconv.Itoa(minEntryFee + tipAmount))
		_, err = tsvc.bsc.RegisterTournament(userPrivateKey, challenge.FeeType, contractAddress, entryFeeWithTip)
		if err != nil {
			revertErr := tsvc.DB.DeleteChallengeRecord(&challengeRecord)
			if revertErr != nil {
				revertErr := tsvc.DB.DeleteChallengeRecord(&challengeRecord)
				if revertErr != nil {
					revertErr := tsvc.DB.DeleteChallengeRecord(&challengeRecord)
					if revertErr != nil {
						tsvc.log.Error("error in DeleteChallengeRecord: ", revertErr)
					}
				}
			}
			return fmt.Errorf("error inserting challenge record: %v", err)
		}
	}

	if challenge.ParticipantCount >= int(challenge.MinParticipants) {
		_ = tsvc.DB.UpdateChallengeState(challenge.ChallengeID, "Ready")
	}
	return nil
}

func (tsvc *ChallengeService) ReportScore(challengeRecord models.ChallengeRecord) error {
	var scoreMap map[string]*string
	score, scoreMap, teamName, err := tsvc.QueryWeightedScore(challengeRecord)
	if err != nil {
		return fmt.Errorf("Error in QueryChallengeScore: %v", err)
	}
	challengeRecord.ScoreMap = scoreMap
	challengeRecord.Score = score
	if challengeRecord.ChallengeDeadline.Before(time.Now()) {
		challengeRecord.ScoreReported = true
	}
	challengeRecord.TeamName = teamName
	err = tsvc.DB.UpdateChallengeRecord(&challengeRecord)
	if err != nil {
		return fmt.Errorf("Error in updating challenge record %d: %v", challengeRecord.ID, err)
	}
	return nil
}

func (tsvc *ChallengeService) checkEth(user *models.UserAccount) error {
	ethBalance, err := tsvc.bsc.GetEthBalance(user.PublicAddress)
	if err != nil {
		return fmt.Errorf("error in GetEthBalance: %v", err)
	}
	ethBalanceFloat, _ := ethBalance.Float64()
	if ethBalanceFloat < 0.05 {
		if user.EthGiftedAt.After(time.Now().Add(-1440 * time.Minute)) {
			return errors.New("spam user")
		}
		_, err = tsvc.bsc.GiftEth(user.PublicAddress)
		if err != nil {
			return fmt.Errorf("error in GiftEth: %v", err)
		}
		err = tsvc.DB.UpdateEthGiftedAt(user.Sub)
		if err != nil {
			return fmt.Errorf("error in UpdateEthGiftedAt: %v", err)
		}
	}
	return nil
}

func getPrivateKey(user *models.UserAccount) (string, error) {
	subStr := user.Sub
	privateKey, err := hdkey.GeneratePrivateKeyFromUUID(subStr)
	if err != nil {
		return "", errors.New("fail to generate privateKey from sub")
	}
	return privateKey, nil
}

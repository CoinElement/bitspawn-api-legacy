/*

 */

package tournament

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/bitspawngg/bitspawn-api/config"
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services"
	"github.com/bitspawngg/bitspawn-api/services/hdkey"
	"github.com/bitspawngg/bitspawn-api/services/poa"
	"github.com/bitspawngg/bitspawn-api/services/queue"
)

type TournamentService struct {
	DB  *models.DB
	log *logrus.Entry
	bsc *poa.BitspawnPoaClient
	sqs *queue.SqsClient
}

type Meta struct {
	Action       string `json:"action,omitempty"`
	Author       string `json:"author,omitempty"`
	TournamentID string `json:"tournamentId" binding:"required"`
	ClubName     string `json:"clubName,omitempty"`
	Members      []int  `json:"members,omitempty"`
	TipAmount    int    `json:"tipAmount,omitempty"`
}

func NewTournamentService(db *models.DB, log *logrus.Logger, bsc *poa.BitspawnPoaClient, sqsSvc *queue.SQSService, config *config.Config) *TournamentService {
	return &TournamentService{
		DB:  db,
		log: log.WithField("service", "tournament"),
		bsc: bsc,
		sqs: sqsSvc.Client(config.AwsConfig().SQSNameTx, queue.WithDelay(2)),
	}
}

func (tsvc *TournamentService) ExecuteStateMachine(tournamentId, sub string, action Action, meta *Meta) error {
	tournament, err := tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tsvc.log.Error("error in GetTournamentData for tournament ", tournamentId, ": ", err)
		return fmt.Errorf("error in GetTournamentData for tournament %s: %v", tournamentId, err)
	}
	user, err := tsvc.DB.GetUser(sub)
	if err != nil {
		tsvc.log.Error("error in GetUser ", sub, ": ", err)
		return fmt.Errorf("error in GetUser %s: %v", sub, err)
	}
	err = tsvc.checkEth(user)
	if err != nil {
		tsvc.log.Error("error in checkEth: ", err)
		return fmt.Errorf("error in checkEth: %v", err)
	}
	state := tournament.Status
	switch action {

	case REGISTER:
		err = tsvc.register(tournament, user, meta)
	case UNREGISTER:
		err = tsvc.unregister(tournament, user, meta)
	case CANCEL:
		err = tsvc.cancel(tournament, user, meta)
	case COMPLETE:
		err = tsvc.complete(tournament, user, meta)
	case FUND:
		err = tsvc.fund(tournament, user, meta)
	case DEPLOY:
		err = tsvc.deploy(tournament, user, meta)
	// case "RegisterTeam":
	// 	err = tsvc.RegisterTeam(tournament, user, meta)
	// case "UnregisterTeam":
	// 	err = tsvc.UnregisterTeam(tournament, user, meta)
	default:
		err = fmt.Errorf("Action %s is not allowed in the current state %s", action, state)
	}
	if err == nil {
		// do nothing
	} else if err.Error() == "No Permission" {
		tsvc.log.Error("error in ExecuteStateMachine for tournament ", tournamentId, ": ", err.Error())
		return fmt.Errorf("No Permission")
	} else if err != nil {
		tsvc.log.Error("error in ExecuteStateMachine for tournament ", tournamentId, ": ", err.Error())
		return fmt.Errorf("error in ExecuteStateMachine for tournament %s: %v", tournamentId, err)
	}

	return nil
}

func (tsvc *TournamentService) UpdateStatus(tournament *models.TournamentData) error {
	status, err := tsvc.bsc.GetTournamentStatus(tournament.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting Tournament Status: %v", err)
	}
	if strings.ToUpper(status) == "ONGOING" { // do not change, use existing value in DB
		status = tournament.Status
	}
	participantCount, err := tsvc.bsc.GetParticipantCount(tournament.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting Participant Count: %v", err)
	}
	prizePool, err := tsvc.bsc.GetPrizePool(tournament.ContractAddress, tournament.FeeType)
	if err != nil {
		return fmt.Errorf("error getting Prize Pool: %v", err)
	}
	err = tsvc.DB.UpdateTournamentState(tournament.TournamentID, status, participantCount, prizePool)
	if err != nil {
		return fmt.Errorf("error updating tournament state to db")
	}
	return nil
}

func (tsvc *TournamentService) deploy(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
	ef := services.ConvertEthToWei(tournament.EntryFee)
	userPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(user.Sub)
	if err != nil {
		return err
	}

	address, _, _, err := tsvc.bsc.DeployTournament(userPrivateKey, tournament.FeeType, tournament.FeePercentage, tournament.OrganizerPercentage, ef, nil)

	if err != nil {
		tsvc.log.Error("error in deploying tournament: ", err)
		return err
	}

	tournamnetContractAddress := address.String()

	td := models.TournamentData{
		ContractAddress: tournamnetContractAddress,
		Status:          "REGISTRATION",
		InviteOnly:      tournament.InviteOnly,
	}

	err = tsvc.DB.UpdateTournamentData(tournament.TournamentID, &td)
	if err != nil {
		return err
	}
	return nil
}

func (tsvc *TournamentService) cancel(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
	statusOnChain, err := tsvc.bsc.GetTournamentStatus(tournament.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting Tournament Status: %v", err)
	}
	if statusOnChain == "Completed" {
		// tournament is already completed, bypass Cancel action
		return nil
	}

	if tournament.OrganizerID != user.Username {
		return fmt.Errorf("No Permission")
	}

	userPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(user.Sub)
	if err != nil {
		return fmt.Errorf("fail to generate private key for %s", user.Username)
	}

	contractAddress := common.HexToAddress(tournament.ContractAddress)
	_, err = tsvc.bsc.CancelTournament(userPrivateKey, contractAddress)
	if err != nil {
		return fmt.Errorf("error in CancelTournament: %v", err)
	}

	_ = tsvc.DB.UpdateTournamentStatus(meta.TournamentID, "CANCELLED")

	return nil
}

func (tsvc *TournamentService) complete(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
	statusOnChain, err := tsvc.bsc.GetTournamentStatus(tournament.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting Tournament Status: %v", err)
	}
	if statusOnChain == "Completed" {
		// tournament is already completed, bypass Complete action
		return nil
	}

	if user.Username != tournament.OrganizerID {
		return fmt.Errorf("tournament not accessible by user")
	}
	userPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(user.Sub)
	if err != nil {
		return fmt.Errorf("fail to generate private key for %s", user.Username)
	}

	playRecords, err := tsvc.DB.GetPlayRecordByTournament(tournament.TournamentID)
	if err != nil {
		return fmt.Errorf("error in GetPlayRecordByTournament: %v", err)
	}

	teamSizeOfRank := make(map[int]int64)
	for _, p := range playRecords {
		if p.RankingNumber != 0 {
			teamSizeOfRank[p.RankingNumber] += 1
		}
	}

	prizePercentage := strings.Split(tournament.PrizeAllocation, ",")
	prizePerTenThousand := []*big.Int{}
	finalPlacementsAddress := make([]common.Address, 0)
	for _, p := range playRecords {
		if p.RankingNumber != 0 {
			finalPlacementsAddress = append(finalPlacementsAddress, common.HexToAddress(p.PublicAddress))
			teamPrizes, _ := strconv.ParseInt(prizePercentage[p.RankingNumber-1], 10, 64)
			prizePerPlayer := teamPrizes * 100 / teamSizeOfRank[p.RankingNumber] // scale by 100
			prizePerTenThousand = append(prizePerTenThousand, big.NewInt(prizePerPlayer))
		}
	}

	contractAddress := common.HexToAddress(tournament.ContractAddress)
	_, err = tsvc.bsc.CompleteTournament(userPrivateKey, contractAddress, finalPlacementsAddress, prizePerTenThousand)
	if err != nil {
		return fmt.Errorf("error in CompleteTournament: %v", err)
	}

	_ = tsvc.DB.UpdateTournamentStatus(meta.TournamentID, "COMPLETED")

	return nil
}

func (tsvc *TournamentService) fund(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
	tipAmount := meta.TipAmount
	statusOnChain, err := tsvc.bsc.GetTournamentStatus(tournament.ContractAddress)
	if err != nil {
		return fmt.Errorf("error getting Tournament Status: %v", err)
	}
	if statusOnChain == "Completed" {
		// tournament is already completed, bypass fund action
		return nil
	}

	userPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(user.Sub)
	if err != nil {
		return fmt.Errorf("fail to generate private key for %s", user.Username)
	}

	contractAddress := common.HexToAddress(tournament.ContractAddress)
	fundAmount := services.ConvertEthToWei(strconv.Itoa(tipAmount))
	_, err = tsvc.bsc.FundTournament(userPrivateKey, tournament.FeeType, contractAddress, fundAmount)
	if err != nil {
		return fmt.Errorf("error in FundTournament: %v", err)
	}
	tournament.OrganizerContribute = tournament.OrganizerContribute + int64(tipAmount)
	tournament.TotalPrizePool = tournament.TotalPrizePool + int64(tipAmount)
	_ = tsvc.DB.UpdateTournamentData(meta.TournamentID, tournament)

	return nil
}

func (tsvc *TournamentService) register(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
	entryFee, err := strconv.ParseInt(tournament.EntryFee, 10, 64)
	if err != nil {
		return fmt.Errorf("Entry Fee format wrong in tournament %s", tournament.TournamentID)
	}
	if strings.ToUpper(tournament.Status) != "REGISTRATION" {
		return fmt.Errorf("tournament not in Registration status")
	}
	userPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(user.Sub)
	if err != nil {
		return fmt.Errorf("fail to generate private key for %s", user.Username)
	}

	// isRegistered, err := tsvc.bsc.IsRegistered(user.PublicAddress, tournament.ContractAddress)
	// if err != nil {
	// 	return fmt.Errorf("error in IsRegistered")
	// }
	// if isRegistered == true {
	// 	// user already registered, do nothing
	// } else {
	contractAddress := common.HexToAddress(tournament.ContractAddress)
	entryFeeInWei := services.ConvertEthToWei(tournament.EntryFee)
	_, err = tsvc.bsc.RegisterTournament(userPrivateKey, tournament.FeeType, contractAddress, entryFeeInWei)
	if err != nil {
		return fmt.Errorf("register tournament transaction failed")
	}

	tournament.ParticipantCount = tournament.ParticipantCount + 1
	tournament.TotalPrizePool = tournament.TotalPrizePool + entryFee
	_ = tsvc.DB.UpdateTournamentData(meta.TournamentID, tournament)

	return nil
}

func (tsvc *TournamentService) unregister(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
	entryFee, err := strconv.ParseInt(tournament.EntryFee, 10, 64)
	if err != nil {
		return fmt.Errorf("Entry Fee format wrong in tournament %s", tournament.TournamentID)
	}
	if strings.ToUpper(tournament.Status) != "REGISTRATION" {
		return fmt.Errorf("tournament not in Registration status")
	}
	userPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(user.Sub)
	if err != nil {
		return fmt.Errorf("fail to generate private key for %s", user.Username)
	}

	isRegistered, err := tsvc.bsc.IsRegistered(user.PublicAddress, tournament.ContractAddress)
	if err != nil {
		return fmt.Errorf("error in IsRegistered")
	}
	if !isRegistered {
		// user already unregistered, do nothing
	} else {
		contractAddress := common.HexToAddress(tournament.ContractAddress)
		_, err = tsvc.bsc.UnregisterTournament(userPrivateKey, contractAddress)
		if err != nil {
			return fmt.Errorf("Unregister tournament transaction failed")
		}
	}

	tournament.ParticipantCount = tournament.ParticipantCount - 1
	tournament.TotalPrizePool = tournament.TotalPrizePool - entryFee
	_ = tsvc.DB.UpdateTournamentData(meta.TournamentID, tournament)

	return nil
}

// func (tsvc *TournamentService) RegisterTeam(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
// 	tsvc.log.Info(meta)
// 	clubName := meta.ClubName
// 	members := meta.Members
// 	tipAmount := meta.TipAmount
// 	if tournament.Status != "Registration" {
// 		return fmt.Errorf("tournament not in Registration status")
// 	}
// 	teamSize, err := tsvc.getTeamSize(tournament.GameType, tournament.GameSubtype)
// 	if err != nil {
// 		return fmt.Errorf("error in getTeamSize: %v", err)
// 	}
// 	if len(members) != teamSize {
// 		return fmt.Errorf("wrong team size: expect %d, got %d", teamSize, len(members))
// 	}

// 	role, err := tsvc.DB.GetClubRole(clubName, tournament.GameType, user.Username)
// 	if role != "OWNER" && role != "MANAGER" {
// 		return fmt.Errorf("cannot register team: user not a club manager")
// 	}

// 	club, err := tsvc.DB.GetClubByName(clubName, tournament.GameType)
// 	if err != nil {
// 		return fmt.Errorf("error in GetClubByName: %v", err)
// 	}

// 	// Begin club Account validation
// 	clubPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(club.Sub)
// 	if err != nil {
// 		return fmt.Errorf("fail to generate private key for %s: %v", user.Username, err)
// 	}
// 	err = tsvc.checkClubEth(&club)
// 	if err != nil {
// 		return fmt.Errorf("error in checkClubEth: %v", err)
// 	}
// 	// End of club Account validation

// 	playRecordsToInsert := []models.PlayRecord{}
// 	for _, memberId := range members {
// 		//player, err := tsvc.DB.GetUserProfileByDisplayName(displayName)
// 		player, err := tsvc.DB.GetClubMemberByID(memberId)
// 		if err != nil {
// 			return fmt.Errorf("player is not a club member")
// 		}
// 		_, err = tsvc.DB.GetClubRole(clubName, tournament.GameType, player.Username)
// 		if err != nil {
// 			if err.Error() == "user not found" {
// 				return fmt.Errorf("player is not a member of this club")
// 			} else {
// 				return fmt.Errorf("error in GetClubRole: %v", err)
// 			}
// 		}
// 		pr := models.PlayRecord{
// 			UserId:        player.Username,
// 			TournamentId:  tournament.TournamentID.String(),
// 			PublicAddress: club.PublicAddress,
// 			Club:          club.ClubName,
// 			GameType:      tournament.GameType,
// 		}
// 		playRecordsToInsert = append(playRecordsToInsert, pr)
// 	}
// 	err = tsvc.DB.BatchInsertPlayRecords(playRecordsToInsert)
// 	if err != nil {
// 		return fmt.Errorf("error inserting play records: %v", err)
// 	}

// 	contractAddress := common.HexToAddress(tournament.ContractAddress)
// 	minEntryFee, _ := strconv.Atoi(tournament.EntryFee)
// 	entryFeeWithTip := services.ConvertEthToWei(strconv.Itoa(minEntryFee + tipAmount))
// 	_, err = tsvc.bsc.RegisterTeam(clubPrivateKey, contractAddress, entryFeeWithTip)
// 	if err != nil {
// 		tsvc.log.Error("error in registerTeam: ", err)
// 		revertErr := tsvc.DB.BatchDeletePlayRecords(playRecordsToInsert)
// 		if revertErr != nil {
// 			tsvc.log.Error("error in BatchDeletePlayRecords: ", revertErr)
// 		}
// 		return fmt.Errorf("register team transaction failed")
// 	}
// 	return nil
// }

// func (tsvc *TournamentService) UnregisterTeam(tournament *models.TournamentData, user *models.UserAccount, meta *Meta) error {
// 	clubName := meta.ClubName
// 	if tournament.Status != "Registration" {
// 		return fmt.Errorf("tournament not in Registration status")
// 	}

// 	role, err := tsvc.DB.GetClubRole(clubName, tournament.GameType, user.Username)
// 	if role != "OWNER" && role != "MANAGER" {
// 		return fmt.Errorf("cannot unregister team: user not a club manager")
// 	}

// 	club, err := tsvc.DB.GetClubByName(clubName, tournament.GameType)
// 	if err != nil {
// 		return fmt.Errorf("error in GetClubByName: %v", err)
// 	}

// 	// Begin club Account validation
// 	clubPrivateKey, err := hdkey.GeneratePrivateKeyFromUUID(club.Sub)
// 	if err != nil {
// 		return fmt.Errorf("fail to generate private key for %s: %v", user.Username, err)
// 	}
// 	err = tsvc.checkClubEth(&club)
// 	if err != nil {
// 		return fmt.Errorf("error in checkClubEth: %v", err)
// 	}
// 	// End of club Account validation

// 	playRecords, err := tsvc.DB.GetPlayRecordByTournamentAndPublicAddress(tournament.TournamentID.String(), club.PublicAddress)
// 	if err != nil {
// 		return fmt.Errorf("error in GetPlayRecordByTournamentAndPublicAddress: %v", err)
// 	}
// 	if len(playRecords) < 1 {
// 		return fmt.Errorf("your club is not registered in this tournament")
// 	}

// 	err = tsvc.DB.BatchDeletePlayRecords(playRecords)
// 	if err != nil {
// 		return fmt.Errorf("error deleting play records: %v", err)
// 	}

// 	contractAddress := common.HexToAddress(tournament.ContractAddress)
// 	_, err = tsvc.bsc.UnregisterTeam(clubPrivateKey, contractAddress)
// 	if err != nil {
// 		revertErr := tsvc.DB.BatchInsertPlayRecords(playRecords)
// 		if revertErr != nil {
// 			tsvc.log.Error("error in BatchInsertPlayRecords: ", revertErr)
// 		}
// 		return fmt.Errorf("Unregister team transaction failed")
// 	}
// 	return nil
// }

func (tsvc *TournamentService) RegisterUserToTournament(user *models.UserAccount, tournamentInfo *models.TournamentData) error {
	playRecord := models.PlayRecord{
		UserId:        user.Username,
		TournamentId:  tournamentInfo.TournamentID,
		PublicAddress: user.PublicAddress,
		GameType:      tournamentInfo.GameType,
	}
	err := tsvc.DB.InsertPlayRecord(&playRecord)
	if err != nil {
		tsvc.log.Errorf("error in InsertPlayRecord: %v", err)
		return fmt.Errorf("error in InsertPlayRecord: %v", err)
	}

	if tournamentInfo.EntryFee != "0" { // Only write blockchain if there is entry fee
		meta := Meta{
			Action:       string(REGISTER),
			Author:       user.Sub,
			TournamentID: tournamentInfo.TournamentID,
		}
		metaJSON, err := json.Marshal(meta)
		if err != nil {
			tsvc.log.Errorf("error in marshal meta: %v", err)
			return fmt.Errorf("error in marshal meta: %v", err)
		}
		err = tsvc.sqs.SendMsg(queue.Message{Title: string(REGISTER), Author: user.Sub, Body: string(metaJSON)})
		if err != nil {
			tsvc.log.Errorf("error in HandleSendMessage: %v", err)
			_, _ = tsvc.DB.DeletePlayRecord(user.Username, tournamentInfo.TournamentID)
			return fmt.Errorf("error in HandleSendMessage: %v", err)
		}
	}
	return nil
}

func (tsvc *TournamentService) checkEth(user *models.UserAccount) error {
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

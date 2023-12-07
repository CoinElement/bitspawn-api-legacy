/*

 */

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitspawngg/bitspawn-api/enum"
	"strconv"
	"strings"
	"time"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services/hdkey"
	"github.com/bitspawngg/bitspawn-api/services/queue"
	"github.com/bitspawngg/bitspawn-api/services/tournament"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type FormTeamRegister struct {
	ClubName     string `json:"clubName" binding:"required"`
	Members      []int  `json:"members" binding:"required"`
	TournamentID string `json:"tournamentId" binding:"required"`
	TipAmount    int    `json:"tipAmount"`
}

func (tc *TournamentController) HandleTeamRegister(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		tc.InternalErrorResponse(c, "user not found")
		return
	}

	form := tournament.Meta{}
	if err := c.ShouldBindJSON(&form); err != nil {
		tc.BadRequestResponse(c, err.Error())
		return
	}
	if form.ClubName == "" {
		tc.BadRequestResponse(c, "missing club name")
		return
	}

	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(form.TournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", form.TournamentID, ": ", err)
		tc.DBErrorResponse(c, "cannot get tournament data")
		return
	}

	if tournamentInfo.InviteOnly && user.Username != tournamentInfo.OrganizerID {
		tc.AuthErrorResponse(c, "this tournament is invite only")
		return
		// if strings.ToUpper(tournamentInfo.GameSubtype) == "1V1" || tournamentInfo.GameSubtype == "SOLO" {
		// 	tc.BadRequestResponse(c, "teams cannot apply for this individual tournament!")
		// 	return
		// }
		// err := tc.tsvc.DB.FetchExistingApplication(form.TournamentID, form.ClubName, "Team Tournament Application")
		// if err == nil {
		// 	tc.BadRequestResponse(c, "Application already exists")
		// 	return
		// } else if err.Error() == "record not found" {
		// 	// do nothing
		// } else if err != nil {
		// 	tc.log.Error("error in get application data for tournament ", form.TournamentID, "&", form.ClubName, ": ", err)
		// 	tc.DBErrorResponse(c, "error getting application info")
		// 	return
		// }

		// list, _ := json.Marshal(form.Members)
		// teamList := string(list)
		// application := models.Application{
		// 	Id:            form.TournamentID,
		// 	ApplicantName: form.ClubName,
		// 	Reviewer:      tournamentInfo.OrganizerID,
		// 	Category:      "Team Tournament Application",
		// 	TeamList:      teamList,
		// }

		// err = tc.tsvc.DB.CreateApplication(&application)
		// if err != nil {
		// 	tc.log.Error("error saving application data to db: ", err)
		// 	tc.DBErrorResponse(c, err.Error())
		// 	return
		// }

		// tc.SuccessResponse(c, "you have sent the application for joining tournament", nil)
		// return
	}

	// Begin user Account validation
	_, err = tc.getPrivKey(c)
	if err != nil {
		tc.log.Error(err)
		tc.AuthErrorResponse(c, err.Error())
		return
	}
	err = tc.checkEth(user)
	if err != nil {
		tc.log.Error(err)
		tc.AuthErrorResponse(c, err.Error())
		return
	}
	// End of user Account validation

	err = tc.tsvc.ExecuteStateMachine(form.TournamentID, user.Username, "RegisterTeam", &form)
	if err != nil {
		tc.log.Error("error in RegisterTeam for tournament ", form.TournamentID, ": ", err)
		tc.InternalErrorResponse(c, "error in RegisterTeam for tournament: "+err.Error())
		return
	}

	tc.SuccessResponse(c, "team tournament registered successfully", nil)
}

type FormTournamentRegister struct {
	TournamentID string `json:"tournamentId" binding:"required"`
	TipAmount    int    `json:"tipAmount"`
}

func (tc *TournamentController) HandleTournamentRegister(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := tournament.Meta{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(form.TournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", form.TournamentID, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}
	if strings.ToUpper(tournamentInfo.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "tournament is not in Registration state", "")
		return
	}

	if tournamentInfo.InviteOnly && user.Username != tournamentInfo.OrganizerID {
		oldApplications, err := tc.tsvc.DB.FindTournamentApplication(form.TournamentID, user.Sub)
		if err != nil {
			tc.log.Error("error in FindTournamentApplication for tournament ", form.TournamentID, "&", user.Username, ": ", err)
			InternalErrorResponseV2(c, "error in FindTournamentApplication", err.Error())
			return
		}
		if len(oldApplications) > 0 {
			BadRequestResponseV2(c, "Already requested to join this tournament", "")
			return
		}
		err = tc.tsvc.DB.CreateTournamentApplication(form.TournamentID, user.Sub)
		if err != nil {
			tc.log.Error("error saving tournament application to db: ", err)
			InternalErrorResponseV2(c, "error saving tournament application to db", err.Error())
			return
		}
	} else {
		// Begin user Account validation
		_, err = tc.getPrivKey(c)
		if err != nil {
			tc.log.Error("error in getPrivKey: " + err.Error())
			InternalErrorResponseV2(c, "error in getPrivKey", err.Error())
			return
		}
		err = tc.checkEth(user)
		if err != nil {
			tc.log.Error("error in checkEth: " + err.Error())
			InternalErrorResponseV2(c, "error in checkEth", err.Error())
			return
		}
		// End of user Account validation

		if tournamentInfo.EntryFee != "0" { // check user balance if there is entry fee
			balance, err := tc.bsc.GetSPWNBalance(user.PublicAddress, tournamentInfo.FeeType)
			if err != nil {
				tc.log.Error("cannot get SPWN balance")
				InternalErrorResponseV2(c, "cannot get SPWN balance", err.Error())
				return
			}
			truncatedBalance, _ := balance.Int64()

			entryFeeInt, err := strconv.ParseInt(tournamentInfo.EntryFee, 10, 64)
			if err != nil {
				tc.log.Error("error in converting entryFee to int64")
				InternalErrorResponseV2(c, "error in converting entryFee to int64", "")
				return
			}
			if truncatedBalance < entryFeeInt {
				BadRequestResponseV2(c, "insufficient balance", "")
				return
			}
		}

		err = tc.tsvc.RegisterUserToTournament(user, tournamentInfo)
		if err != nil {
			tc.log.Errorf("error in RegisterUserToTournament: %v", err)
			InternalErrorResponseV2(c, "error in RegisterUserToTournament", err.Error())
			return
		}
	}

	allParticipants, err := tc.tsvc.DB.GetAllParticipantsInfo(form.TournamentID)
	if err != nil {
		tc.log.Error("error getting tournament participants", form.TournamentID, err)
		InternalErrorResponseV2(c, "error getting tournament participants", "")
		return
	}
	assignedParticipants, err := tc.tsvc.DB.GetAssignedParticipants(form.TournamentID)
	if err != nil {
		tc.log.Error("error in GetAssignedParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetAssignedParticipants", err.Error())
		return
	}
	var assignedParticipantsUsernames []string
	for _, p := range assignedParticipants {
		assignedParticipantsUsernames = append(assignedParticipantsUsernames, p.Username)
	}
	unassignedParticipants := []models.UserInfo{}
	for _, p := range allParticipants {
		if !stringInSlice(p.Username, assignedParticipantsUsernames) {
			unassignedParticipants = append(unassignedParticipants, p)
		}
	}

	SuccessResponseV2(c, gin.H{
		"tournamentId":           form.TournamentID,
		"participantsList":       allParticipants,
		"assignedParticipants":   assignedParticipants,
		"unassignedParticipants": unassignedParticipants,
	})
}

type FormTeamUnregister struct {
	ClubName     string `json:"clubName" binding:"required"`
	TournamentID string `json:"tournamentId" binding:"required"`
}

func (tc *TournamentController) HandleTeamUnregister(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		tc.InternalErrorResponse(c, "user not found")
		return
	}

	form := tournament.Meta{}
	if err := c.ShouldBindJSON(&form); err != nil {
		tc.BadRequestResponse(c, err.Error())
		return
	}
	if form.ClubName == "" {
		tc.BadRequestResponse(c, "missing club name")
		return
	}

	err := tc.tsvc.ExecuteStateMachine(form.TournamentID, user.Username, "UnregisterTeam", &form)
	if err != nil {
		tc.log.Error("error in Unregister team in tournament ", form.TournamentID, ": ", err)
		tc.InternalErrorResponse(c, "error in Unregister team in tournament: "+err.Error())
		return
	}

	tc.SuccessResponse(c, "team unregistered successfully", nil)
}

type FormTournamentUnregister struct {
	TournamentID string `json:"tournamentId" binding:"required"`
}

func (tc *TournamentController) HandleTournamentUnregister(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	// Begin user Account validation
	_, err := tc.getPrivKey(c)
	if err != nil {
		tc.log.Error(err)
		tc.AuthErrorResponse(c, err.Error())
		return
	}
	err = tc.checkEth(user)
	if err != nil {
		tc.log.Error(err)
		tc.AuthErrorResponse(c, err.Error())
		return
	}
	// End of user Account validation

	form := tournament.Meta{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(form.TournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", form.TournamentID, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}
	if strings.ToUpper(tournamentInfo.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "tournament not in Registration status", "")
		return
	}

	deletedRecord, err := tc.tsvc.DB.DeletePlayRecord(user.Username, tournamentInfo.TournamentID)
	if err != nil {
		tc.log.Error("error in DeletePlayRecord: " + err.Error())
		InternalErrorResponseV2(c, "error in DeletePlayRecord: ", err.Error())
		return
	}

	if tournamentInfo.EntryFee != "0" { // Only write blockchain if there is entry fee
		form.Action = string(tournament.UNREGISTER)
		form.Author = user.Sub
		form.TournamentID = tournamentInfo.TournamentID
		formJSON, _ := json.Marshal(form)
		err = tc.SQSClient.SendMsg(queue.Message{Title: string(tournament.UNREGISTER), Author: user.Sub, Body: string(formJSON)})
		if err != nil {
			tc.log.Error("error in HandleSendMessage: " + err.Error())
			InternalErrorResponseV2(c, "error in HandleSendMessage: ", err.Error())
			_ = tc.tsvc.DB.InsertPlayRecord(deletedRecord)
			return
		}
	}

	allParticipants, err := tc.DB.GetAllParticipantsInfo(form.TournamentID)
	if err != nil {
		tc.log.Error("error getting tournament participants", form.TournamentID, err)
		InternalErrorResponseV2(c, "error getting tournament participants", "")
		return
	}
	assignedParticipants, err := tc.DB.GetAssignedParticipants(form.TournamentID)
	if err != nil {
		tc.log.Error("error in GetAssignedParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetAssignedParticipants", err.Error())
		return
	}
	var assignedParticipantsUsernames []string
	for _, p := range assignedParticipants {
		assignedParticipantsUsernames = append(assignedParticipantsUsernames, p.Username)
	}
	unassignedParticipants := []models.UserInfo{}
	for _, p := range allParticipants {
		if !stringInSlice(p.Username, assignedParticipantsUsernames) {
			unassignedParticipants = append(unassignedParticipants, p)
		}
	}

	SuccessResponseV2(c, gin.H{
		"tournamentId":           form.TournamentID,
		"participantsList":       allParticipants,
		"assignedParticipants":   assignedParticipants,
		"unassignedParticipants": unassignedParticipants,
	})
}

type FormTournamentFund struct {
	TournamentID string `json:"tournament_id"`
	Funds        string `json:"funds"`
}

func (tc *TournamentController) HandleV2TournamentFund(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	// Begin user Account validation
	err := tc.checkEth(user)
	if err != nil {
		tc.log.Error(err)
		InternalErrorResponseV2(c, "fail to fill gas", err.Error())
		return
	}
	// End of user Account validation

	form := tournament.Meta{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TournamentID == "" || form.TipAmount <= 0 {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	tournamentId := form.TournamentID
	tournamentData, err := tc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error getting tournament: ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	if strings.ToUpper(tournamentData.Status) != "REGISTRATION" && strings.ToUpper(tournamentData.Status) != "STARTED" {
		BadRequestResponseV2(c, "tournament not in Registration or Started status", "")
	}
	tResponse, err := tc.DB.FormatTournamentForResponse(*tournamentData)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}
	tResponse.TotalPrizePool += int64(form.TipAmount)
	tResponse.FundContribute += int64(form.TipAmount)

	balance, err := tc.bsc.GetSPWNBalance(user.PublicAddress, tResponse.FeeType)
	if err != nil {
		tc.log.Error("cannot get SPWN balance")
		InternalErrorResponseV2(c, "cannot get SPWN balance", err.Error())
		return
	}
	truncatedBalance, _ := balance.Int64()
	if truncatedBalance < int64(form.TipAmount) {
		BadRequestResponseV2(c, "insufficient balance in wallet", "")
		return
	}

	form.Action = string(tournament.FUND)
	form.Author = user.Sub
	form.TournamentID = tournamentId
	formJSON, err := json.Marshal(form)
	if err != nil {
		tc.log.Error("error in marshal form meta: " + err.Error())
		InternalErrorResponseV2(c, "error in marshal form meta: ", err.Error())
		return
	}
	err = tc.SQSClient.SendMsg(queue.Message{Title: string(tournament.FUND), Author: user.Sub, Body: string(formJSON)})
	if err != nil {
		tc.log.Error("error in HandleSendMessage: " + err.Error())
		InternalErrorResponseV2(c, "error in HandleSendMessage: ", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

type FormTournamentBannerUpload struct {
	File string `json:"file"`
}

func (tc *TournamentController) HandleTournamentBannerUpload(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tournament, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error getting tournament: ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}

	moderatorUsernames := strings.Split(tournament.Roles, ",")
	if tournament.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	maxSize := int64(4096000)
	err = c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		BadRequestResponseV2(c, "Image too large. Max Size: 4M", "")
		return
	}
	file, fileHeader, err := c.Request.FormFile("banner")
	if err != nil {
		tc.log.Error("Could not get uploaded file", err)
		InternalErrorResponseV2(c, "Could not get uploaded file", err.Error())
		return
	}
	defer file.Close()
	tc.log.Info("fileHeader: ", fileHeader.Filename)
	awsConfig := tc.conf.AwsConfig()
	bucketname := awsConfig.S3BucketName
	tempBannerUrl, _ := tc.S3UploadClient.HandleS3TournamentBannerUpload(bucketname, user.PublicAddress, file)

	tResponse, err := tc.DB.UpdateTournamentBanner(tournamentID, tempBannerUrl)
	if err != nil {
		tc.log.Error("error updating tournament banner", err)
		InternalErrorResponseV2(c, "updating tournament banner fails", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

func (tc *TournamentController) HandleTournamentLogoUpload(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tournament, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error getting tournament: ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}

	moderatorUsernames := strings.Split(tournament.Roles, ",")
	if tournament.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	maxSize := int64(4096000)
	err = c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		BadRequestResponseV2(c, "Image too large. Max Size: 4M", "")
		return
	}
	file, fileHeader, err := c.Request.FormFile("logo")
	if err != nil {
		tc.log.Error("Could not get uploaded file", err)
		InternalErrorResponseV2(c, "Could not get uploaded file", err.Error())
		return
	}
	defer file.Close()
	tc.log.Info("fileHeader: ", fileHeader.Filename)
	awsConfig := tc.conf.AwsConfig()
	bucketname := awsConfig.S3BucketName
	tempLogoUrl, _ := tc.S3UploadClient.HandleS3TournamentLogoUpload(bucketname, user.PublicAddress, file)

	tResponse, err := tc.DB.UpdateTournamentLogo(tournamentID, tempLogoUrl)
	if err != nil {
		tc.log.Error("error updating tournament logo", err)
		InternalErrorResponseV2(c, "updating tournament logo fails", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

func (tc *TournamentController) HandleTournamentThumbnailUpload(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tournament, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error getting tournament: ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}

	moderatorUsernames := strings.Split(tournament.Roles, ",")
	if tournament.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	maxSize := int64(4096000)
	err = c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		BadRequestResponseV2(c, "Image too large. Max Size: 4M", "")
		return
	}
	file, fileHeader, err := c.Request.FormFile("thumbnail")
	if err != nil {
		tc.log.Error("Could not get uploaded file", err)
		InternalErrorResponseV2(c, "Could not get uploaded file", err.Error())
		return
	}
	defer file.Close()
	tc.log.Info("fileHeader: ", fileHeader.Filename)
	awsConfig := tc.conf.AwsConfig()
	bucketname := awsConfig.S3BucketName
	tempThumbnailUrl, _ := tc.S3UploadClient.HandleS3TournamentThumbnailUpload(bucketname, user.PublicAddress, file)

	tResponse, err := tc.DB.UpdateTournamentThumbnail(tournamentID, tempThumbnailUrl)
	if err != nil {
		tc.log.Error("error updating tournament banner", err)
		InternalErrorResponseV2(c, "updating tournament banner fails", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

func (tc *TournamentController) HandleTournamentEnum(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	gameTypeStructs, err := tc.DB.GetGameType()
	if err != nil {
		tc.log.Error("error get game type from db", err)
		InternalErrorResponseV2(c, "error get game type from db", err.Error())
		return
	}
	if len(gameTypeStructs) == 0 {
		tc.log.Error("there are no game types in db")
		InternalErrorResponseV2(c, "there are no game types in db", "")
		return
	}

	gameTypes := []string{}
	for _, gameTypeStruct := range gameTypeStructs {
		gameTypes = append(gameTypes, gameTypeStruct.GameType)
	}

	tournamentFormatStructs, err := tc.DB.ListFormat()
	if err != nil {
		tc.log.Error("error get tournament formats from db", err)
		InternalErrorResponseV2(c, "error get tournament formats from db", err.Error())
		return
	}
	if len(tournamentFormatStructs) == 0 {
		tc.log.Error("there are no tournament formats in db")
		InternalErrorResponseV2(c, "there are no tournament formats in db", "")
		return
	}

	tournamentFormats := []string{}
	for _, tournamentFormat := range tournamentFormatStructs {
		tournamentFormats = append(tournamentFormats, tournamentFormat.Format)
	}

	t := gin.H{
		"gameTypes":         gameTypes,
		"tournamentFormats": tournamentFormats,
	}
	SuccessResponseV2(c, t)
}

type FormBracketUpdate struct {
	TournamentId     string                `json:"tournamentId"`
	TournamentFormat enum.TournamentFormat `json:"tournamentFormat"`
	NumberOfTeams    int                   `json:"numberOfTeams"`
}

func (tc *TournamentController) HandleV2TournamentBracketUpdate(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormBracketUpdate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TournamentId == "" || form.TournamentFormat == "" || form.NumberOfTeams == 0 {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}
	if !form.TournamentFormat.IsValid() {
		BadRequestResponseV2(c, "Only Single elimination format is supported", "")
		return
	}
	if form.NumberOfTeams < 2 {
		BadRequestResponseV2(c, "Number of Teams must be at least 2", "")
		return
	}

	tournamentId := form.TournamentId
	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}

	moderators := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderators) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	t, err := tc.DB.FormatTournamentForResponse(*tournamentInfo)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}
	t.TournamentFormat = form.TournamentFormat
	t.NumberOfTeams = form.NumberOfTeams

	if strings.ToUpper(tournamentInfo.Status) != "DRAFT" && strings.ToUpper(tournamentInfo.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "the tournament is not in Draft or Registration status", "")
		return
	}

	if !tournamentInfo.IsManual {
		BadRequestResponseV2(c, "You can only manage brackets on a manual tournament!", "")
		return
	}

	err = tc.DB.UpdateTournament(tournamentId, t)
	if err != nil {
		tc.log.Error("error in Update Tournament: ", err)
		InternalErrorResponseV2(c, "cannot update tournament", err.Error())
		return
	}

	err = tc.DB.DeleteMatchesByTournamentId(tournamentId)
	if err != nil {
		tc.log.Error("error in DeleteMatchesByTournamentId: ", err)
		InternalErrorResponseV2(c, "error in DeleteMatchesByTournamentId", err.Error())
		return
	}
	err = tc.orgSvc.CreateManualMatchSchedule(tournamentId)
	if err != nil {
		tc.log.Error("error in CreateMatchSchedule for ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "CreateMatchSchedule failed", err.Error())
		return
	}

	updatedMatches, err := tc.tsvc.DB.GetMatchesByTournament(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetMatchesByTournament: ", err)
		InternalErrorResponseV2(c, "error in GetMatchesByTournament", err.Error())
		return
	}

	matchesOutput, _ := tc.DB.FormatMatchesForOutput(updatedMatches)
	SuccessResponseV2(c, gin.H{
		"tournamentDetails": t,
		"matchList":         matchesOutput,
	})
}

type FormBracketTeamsSwap struct {
	TournamentId string `json:"tournamentId"`
	TeamAId      string `json:"teamAId"`
	TeamBId      string `json:"teamBId"`
}

func (tc *TournamentController) HandleV2TournamentBracketTeamsSwap(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormBracketTeamsSwap{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TournamentId == "" || form.TeamAId == "" || form.TeamBId == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	tournamentId := form.TournamentId
	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}

	moderators := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderators) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	if strings.ToUpper(tournamentInfo.Status) != "READY" && strings.ToUpper(tournamentInfo.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "the tournament is not in Ready or Registration status", "")
		return
	}

	if !tournamentInfo.IsManual {
		BadRequestResponseV2(c, "You can only manage brackets on a manual tournament!", "")
		return
	}

	err = tc.DB.SwapTeamsV2(tournamentId, form.TeamAId, form.TeamBId)
	if err != nil {
		tc.log.Errorf("error in SwapTeamsV2(%s, %s, %s):%v", tournamentId, form.TeamAId, form.TeamBId, err)
		InternalErrorResponseV2(c, "error in SwapTeamsV2", err.Error())
		return
	}

	updatedMatches, err := tc.tsvc.DB.GetMatchesByTournament(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetMatchesByTournament: ", err)
		InternalErrorResponseV2(c, "error in GetMatchesByTournament", err.Error())
		return
	}

	matchesOutput, _ := tc.DB.FormatMatchesForOutput(updatedMatches)
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"matchList":    matchesOutput,
	})
}

type FormRoundMatchBestOfN struct {
	TournamentId string `json:"tournamentId"`
	Round        int    `json:"round"`
	BestOfN      int    `json:"bestOfN"`
}

func (tc *TournamentController) HandleV2TournamentRoundMatchBestOfN(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormRoundMatchBestOfN{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentId := form.TournamentId
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}
	if form.Round < 1 || form.BestOfN < 1 {
		BadRequestResponseV2(c, "invalid input parameter", "")
		return
	}

	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}

	moderators := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderators) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	if strings.ToUpper(tournamentInfo.Status) != "DRAFT" &&
		strings.ToUpper(tournamentInfo.Status) != "REGISTRATION" &&
		strings.ToUpper(tournamentInfo.Status) != "STARTED" {
		BadRequestResponseV2(c, "tournament not in Draft, Registration or Started status", "")
		return
	}

	if !tournamentInfo.IsManual {
		BadRequestResponseV2(c, "You can only manage brackets on a manual tournament!", "")
		return
	}

	updatedMatches, err := tc.orgSvc.UpdateRoundBestOfN(tournamentId, form.Round, form.BestOfN)
	if err != nil {
		tc.log.Error("error in UpdateRoundsBestOfN for ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "UpdateRoundsBestOfN failed", err.Error())
		return
	}

	matchesOutput, _ := tc.DB.FormatMatchesForOutput(updatedMatches)

	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"matchList":    matchesOutput,
	})
}

type FormRoundTime struct {
	TournamentId string    `json:"tournamentId"`
	Round        int       `json:"round"`
	Time         time.Time `json:"time"`
}

func (tc *TournamentController) HandleV2TournamentRoundMatchTime(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormRoundTime{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentId := form.TournamentId
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}
	if form.Round < 1 {
		BadRequestResponseV2(c, "invalid input parameter", "")
		return
	}

	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}

	moderators := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderators) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	if strings.ToUpper(tournamentInfo.Status) != "READY" && strings.ToUpper(tournamentInfo.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "the tournament is not in Ready or Registration status", "")
		return
	}

	if !tournamentInfo.IsManual {
		BadRequestResponseV2(c, "You can only manage brackets on a manual tournament!", "")
		return
	}

	updatedMatches, err := tc.orgSvc.UpdateRoundTime(tournamentId, form.Round, form.Time)
	if err != nil {
		tc.log.Error("error in UpdateRoundTime for ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "UpdateRoundTime failed", err.Error())
		return
	}

	matchesOutput, _ := tc.DB.FormatMatchesForOutput(updatedMatches)

	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"matchList":    matchesOutput,
	})
}

func (tc *TournamentController) getPrivKey(c *gin.Context) (string, error) {
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		return "", errors.New("token not found")
	}

	authCofig := *tc.AuthStore
	cognitoRegion := authCofig.CognitoRegion()
	cognitoUserPoolID := authCofig.CognitoUserPoolID()

	jwtToken, err := tc.AuthStore.ParseJWT(token, cognitoRegion, cognitoUserPoolID)
	if err != nil || !jwtToken.Valid {
		// jwt is not valid
		return "", errors.New("token is not valid")
	}

	subStr := jwtToken.Claims.(jwt.MapClaims)["sub"].(string)
	privateKey, err := hdkey.GeneratePrivateKeyFromUUID(subStr)
	if err != nil {
		return "", errors.New("fail to create privateKey")
	}
	return privateKey, nil
}

func (tc *TournamentController) checkEth(user *models.UserAccount) error {
	ethBalance, err := tc.bsc.GetEthBalance(user.PublicAddress)
	if err != nil {
		return fmt.Errorf("error in GetEthBalance: %v", err)
	}
	ethBalanceFloat, _ := ethBalance.Float64()
	if ethBalanceFloat < 0.05 {
		if user.EthGiftedAt.After(time.Now().Add(-1440 * time.Minute)) {
			return errors.New("spam user")
		}
		_, err = tc.bsc.GiftEth(user.PublicAddress)
		if err != nil {
			return fmt.Errorf("error in GiftEth: %v", err)
		}
		err = tc.DB.UpdateEthGiftedAt(user.Sub)
		if err != nil {
			return fmt.Errorf("error in UpdateEthGiftedAt: %v", err)
		}
	}
	return nil
}

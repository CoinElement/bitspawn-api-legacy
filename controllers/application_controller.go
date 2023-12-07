/*

 */

package controllers

import (
	"strconv"
	"strings"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/gin-gonic/gin"
)

type ApplicationController struct {
	*TournamentController
}

func NewApplicationController(tc *TournamentController) *ApplicationController {
	return &ApplicationController{tc}
}

func (ac *ApplicationController) HandleTournamentListApplications(c *gin.Context) {
	user := ac.userFromContext(c)
	if user == nil {
		ac.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentId := c.Param("tournamentId")
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tournamentInfo, err := ac.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ac.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	moderatorUsernames := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Only moderators can perform this action", "")
		return
	}

	outputApplications, err := ac.tsvc.DB.ListApplicationByTournament(tournamentId)
	if err != nil {
		ac.log.Error("error in ListApplicationByTournament", err)
		InternalErrorResponseV2(c, "error in ListApplicationByTournament", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"joinList":     outputApplications,
	})
}

type FormApplicationAcceptOrDecline struct {
	RequestID string `json:"requestId"`
}

func (ac *ApplicationController) HandleTournamentApplicationAccept(c *gin.Context) {
	user := ac.userFromContext(c)
	if user == nil {
		ac.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormApplicationAcceptOrDecline{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.RequestID == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	applicationToDelete, err := ac.tsvc.DB.FetchTournamentApplication(form.RequestID)
	if err != nil {
		ac.log.Error("error in FetchTournamentApplication: ", err)
		InternalErrorResponseV2(c, "error in FetchTournamentApplication", err.Error())
		return
	}
	tournamentId := applicationToDelete.TournamentID

	tournamentInfo, err := ac.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ac.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	moderatorUsernames := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Only moderators can perform this action", "")
		return
	}

	applicant := applicationToDelete.UserAccount

	existingPlayRecords, err := ac.tsvc.DB.GetPlayRecordByTournamentAndPublicAddress(tournamentId, applicant.PublicAddress)
	if err != nil {
		ac.log.Error("error in GetPlayRecordByTournamentAndPublicAddress: ", err)
		InternalErrorResponseV2(c, "error in GetPlayRecordByTournamentAndPublicAddress", err.Error())
		return
	}
	if len(existingPlayRecords) > 0 {
		BadRequestResponseV2(c, "player has already registered in this tournament", "")
		return
	}

	if tournamentInfo.EntryFee != "0" { // check user balance if there is entry fee
		err = ac.checkEth(applicant)
		if err != nil {
			ac.log.Errorf("error in checkEth for user - %s: %s", applicant.Username, err.Error())
			InternalErrorResponseV2(c, "error in checkEth for applicant", err.Error())
			return
		}

		balance, err := ac.bsc.GetSPWNBalance(applicant.PublicAddress, tournamentInfo.FeeType)
		if err != nil {
			ac.log.Error("cannot get SPWN balance")
			InternalErrorResponseV2(c, "cannot get SPWN balance", err.Error())
			return
		}
		truncatedBalance, _ := balance.Int64()

		entryFeeInt, err := strconv.ParseInt(tournamentInfo.EntryFee, 10, 64)
		if err != nil {
			ac.log.Error("error in converting entryFee to int64")
			InternalErrorResponseV2(c, "error in converting entryFee to int64", "")
			return
		}
		if truncatedBalance < entryFeeInt {
			BadRequestResponseV2(c, "insufficient balance", "")
			return
		}
	}

	err = ac.tsvc.DB.DeleteTournamentApplication(form.RequestID)
	if err != nil {
		ac.log.Error("error in DeleteTournamentApplication: ", err)
		InternalErrorResponseV2(c, "error in DeleteTournamentApplication", err.Error())
		return
	}

	err = ac.tsvc.RegisterUserToTournament(applicant, tournamentInfo)
	if err != nil {
		ac.log.Errorf("error in RegisterUserToTournament: %v", err)
		InternalErrorResponseV2(c, "error in RegisterUserToTournament", err.Error())
		return
	}

	outputApplications, err := ac.tsvc.DB.ListApplicationByTournament(tournamentId)
	if err != nil {
		ac.log.Error("error in ListApplicationByTournament", err)
		InternalErrorResponseV2(c, "error in ListApplicationByTournament", err.Error())
		return
	}

	allParticipants, err := ac.tsvc.DB.GetAllParticipantsInfo(tournamentId)
	if err != nil {
		ac.log.Error("error getting tournament participants", tournamentId, err)
		InternalErrorResponseV2(c, "error getting tournament participants", err.Error())
		return
	}
	assignedParticipants, err := ac.tsvc.DB.GetAssignedParticipants(tournamentId)
	if err != nil {
		ac.log.Error("error in GetAssignedParticipants: ", err)
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
		"tournamentId":           tournamentId,
		"joinList":               outputApplications,
		"participantsList":       allParticipants,
		"unassignedParticipants": unassignedParticipants,
	})
}

func (ac *ApplicationController) HandleTournamentApplicationDecline(c *gin.Context) {
	user := ac.userFromContext(c)
	if user == nil {
		ac.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormApplicationAcceptOrDecline{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.RequestID == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	applicationToDelete, err := ac.tsvc.DB.FetchTournamentApplication(form.RequestID)
	if err != nil {
		ac.log.Error("error in FetchTournamentApplication: ", err)
		InternalErrorResponseV2(c, "error in FetchTournamentApplication", err.Error())
		return
	}
	tournamentId := applicationToDelete.TournamentID

	tournamentInfo, err := ac.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ac.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	moderatorUsernames := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Only moderators can perform this action", "")
		return
	}

	err = ac.tsvc.DB.DeleteTournamentApplication(form.RequestID)
	if err != nil {
		ac.log.Error("error in DeleteTournamentApplication: ", err)
		InternalErrorResponseV2(c, "error in DeleteTournamentApplication", err.Error())
		return
	}

	outputApplications, err := ac.tsvc.DB.ListApplicationByTournament(tournamentId)
	if err != nil {
		ac.log.Error("error in ListApplicationByTournament", err)
		InternalErrorResponseV2(c, "error in ListApplicationByTournament", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"joinList":     outputApplications,
	})
}

type FormApplicationCreate struct {
	TournamentID string `json:"tournamentId"`
}

func (ac *ApplicationController) HandleApplicationCreate(c *gin.Context) {
	user := ac.userFromContext(c)
	if user == nil {
		ac.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormApplicationCreate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}
	if form.TournamentID == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}
	tournamentId := form.TournamentID

	tournamentInfo, err := ac.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ac.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	if !tournamentInfo.InviteOnly {
		BadRequestResponseV2(c, "this tournament is not invite only", "")
		return
	}
	if tournamentInfo.Status != "REGISTRATION" {
		BadRequestResponseV2(c, "tournament is not in Registration state", "")
		return
	}

	applicant := user
	existingPlayRecords, err := ac.tsvc.DB.GetPlayRecordByTournamentAndPublicAddress(tournamentId, applicant.PublicAddress)
	if err != nil {
		ac.log.Error("error in GetPlayRecordByTournamentAndPublicAddress: ", err)
		InternalErrorResponseV2(c, "error in GetPlayRecordByTournamentAndPublicAddress", err.Error())
		return
	}
	if len(existingPlayRecords) > 0 {
		BadRequestResponseV2(c, "player has already registered in this tournament", "")
		return
	}

	oldApplications, err := ac.tsvc.DB.FindTournamentApplication(form.TournamentID, user.Sub)
	if err != nil {
		ac.log.Error("error in FindTournamentApplication for tournament ", form.TournamentID, "&", user.Username, ": ", err)
		InternalErrorResponseV2(c, "error in FindTournamentApplication", err.Error())
		return
	}
	if len(oldApplications) > 0 {
		BadRequestResponseV2(c, "Already requested to join this tournament", "")
		return
	}

	oldInvitations, err := ac.tsvc.DB.FindTournamentInvite(tournamentId, applicant.Sub)
	if err != nil {
		ac.log.Error("error in FindTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in FindTournamentInvite", err.Error())
		return
	}
	if len(oldInvitations) > 0 {
		BadRequestResponseV2(c, "Already invited player to join this tournament", "")
		return
	}

	err = ac.tsvc.DB.CreateTournamentApplication(form.TournamentID, user.Sub)
	if err != nil {
		ac.log.Error("error saving tournament application to db: ", err)
		InternalErrorResponseV2(c, "error saving tournament application to db", err.Error())
		return
	}

	outputApplicationList, err := ac.tsvc.DB.ListApplicationByTournament(tournamentId)
	if err != nil {
		ac.log.Error("error in ListApplicationByTournament", err)
		InternalErrorResponseV2(c, "error in ListApplicationByTournament", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"joinList":     outputApplicationList,
	})
}

type FormApplicationCancel struct {
	RequestID string `json:"requestId"`
}

func (ac *ApplicationController) HandleApplicationCancel(c *gin.Context) {
	user := ac.userFromContext(c)
	if user == nil {
		ac.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormApplicationCancel{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.RequestID == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	applicationToCancel, err := ac.tsvc.DB.FetchTournamentApplication(form.RequestID)
	if err != nil {
		ac.log.Error("error in FetchTournamentApplication: ", err)
		InternalErrorResponseV2(c, "error in FetchTournamentApplication", err.Error())
		return
	}
	if user.Sub != applicationToCancel.Sub {
		BadRequestResponseV2(c, "Only player themselves can perform this action", "")
		return
	}

	err = ac.tsvc.DB.DeleteTournamentApplication(form.RequestID)
	if err != nil {
		ac.log.Error("error in DeleteTournamentApplication: ", err)
		InternalErrorResponseV2(c, "error in DeleteTournamentApplication", err.Error())
		return
	}

	outputApplications, err := ac.tsvc.DB.ListActiveApplicationsByUser(user)
	if err != nil {
		ac.log.Error("error in ListActiveApplicationsByUser: ", err)
		InternalErrorResponseV2(c, "error in ListActiveApplicationsByUser", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": applicationToCancel.TournamentID,
		"userJoinList": outputApplications,
	})
}

func (ac *ApplicationController) HandleTournamentListUserRequests(c *gin.Context) {
	user := ac.userFromContext(c)
	if user == nil {
		ac.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}
	outputApplications, err := ac.tsvc.DB.ListActiveApplicationsByUser(user)
	if err != nil {
		ac.log.Error("error in ListActiveApplicationsByUser", err)
		InternalErrorResponseV2(c, "error in ListActiveApplicationsByUser", err.Error())
		return
	}
	outputInvites, err := ac.tsvc.DB.ListActiveInvitesByUser(user)
	if err != nil {
		ac.log.Error("error in ListActiveInvitesByUser", err)
		InternalErrorResponseV2(c, "error in ListActiveInvitesByUser", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"inviteList": outputInvites,
		"joinList":   outputApplications,
	})
}

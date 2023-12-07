/*

 */

package controllers

import (
	"strconv"
	"strings"

	"github.com/bitspawngg/bitspawn-api/services/tournament"
	"github.com/gin-gonic/gin"
)

type InvitationController struct {
	*TournamentController
}

func NewInvitationController(tc *TournamentController) *InvitationController {
	return &InvitationController{tc}
}

type FormInvitationsCreate struct {
	TournamentID string `json:"tournamentId"`
	DisplayName  string `json:"displayName"`
}

func (ic *InvitationController) HandleInvitationCreate(c *gin.Context) {
	user := ic.userFromContext(c)
	if user == nil {
		ic.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormInvitationsCreate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}
	if form.TournamentID == "" || form.DisplayName == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}
	tournamentId := form.TournamentID

	tournamentInfo, err := ic.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ic.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	moderatorUsernames := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Only moderators can perform this action", "")
		return
	}

	invitee, err := ic.tsvc.DB.GetUserProfileByDisplayName(form.DisplayName)
	if err != nil {
		ic.log.Error("error getting user ", form.DisplayName, ": ", err)
		InternalErrorResponseV2(c, "Error getting user "+form.DisplayName, err.Error())
		return
	}

	existingPlayRecords, err := ic.tsvc.DB.GetPlayRecordByTournamentAndPublicAddress(tournamentId, invitee.PublicAddress)
	if err != nil {
		ic.log.Error("error in GetPlayRecordByTournamentAndPublicAddress: ", err)
		InternalErrorResponseV2(c, "error in GetPlayRecordByTournamentAndPublicAddress", err.Error())
		return
	}
	if len(existingPlayRecords) > 0 {
		BadRequestResponseV2(c, "player has already registered in this tournament", "")
		return
	}

	oldInvitations, err := ic.tsvc.DB.FindTournamentInvite(tournamentId, invitee.Sub)
	if err != nil {
		ic.log.Error("error in FindTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in FindTournamentInvite", err.Error())
		return
	}
	if len(oldInvitations) > 0 {
		BadRequestResponseV2(c, "Already invited player to join this tournament", "")
		return
	}
	err = ic.tsvc.DB.CreateTournamentInvite(form.TournamentID, invitee.Sub)
	if err != nil {
		ic.log.Error("error saving tournament invitation to db: ", err)
		InternalErrorResponseV2(c, "error saving tournament invitation to db", err.Error())
		return
	}

	outputInviteList, err := ic.tsvc.DB.ListInviteByTournament(tournamentId)
	if err != nil {
		ic.log.Error("error in ListInviteByTournament", err)
		InternalErrorResponseV2(c, "error in ListInviteByTournament", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"inviteList":   outputInviteList,
	})
}

func (ic *InvitationController) HandleTournamentInvitationFetch(c *gin.Context) {
	user := ic.userFromContext(c)
	if user == nil {
		ic.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentId := c.Param("tournamentId")
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tournamentInfo, err := ic.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ic.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	moderatorUsernames := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Only moderators can perform this action", "")
		return
	}

	outputInvites, err := ic.tsvc.DB.ListInviteByTournament(tournamentId)
	if err != nil {
		ic.log.Error("error in ListInviteByTournament", err)
		InternalErrorResponseV2(c, "error in ListInviteByTournament", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"inviteList":   outputInvites,
	})
}

func (ic *InvitationController) HandleTournamentAllRequestsFetch(c *gin.Context) {
	user := ic.userFromContext(c)
	if user == nil {
		ic.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentId := c.Param("tournamentId")
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tournamentInfo, err := ic.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ic.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	moderatorUsernames := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Only moderators can perform this action", "")
		return
	}

	outputApplications, err := ic.tsvc.DB.ListApplicationByTournament(tournamentId)
	if err != nil {
		ic.log.Error("error in ListApplicationByTournament", err)
		InternalErrorResponseV2(c, "error in ListApplicationByTournament", err.Error())
		return
	}
	outputInvites, err := ic.tsvc.DB.ListInviteByTournament(tournamentId)
	if err != nil {
		ic.log.Error("error in ListInviteByTournament", err)
		InternalErrorResponseV2(c, "error in ListInviteByTournament", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"inviteList":   outputInvites,
		"joinList":     outputApplications,
	})
}

type FormInvitationCancel struct {
	RequestID string `json:"requestId"`
}

func (ic *InvitationController) HandleInvitationCancel(c *gin.Context) {
	user := ic.userFromContext(c)
	if user == nil {
		ic.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormInvitationCancel{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.RequestID == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	InviteToCancel, err := ic.tsvc.DB.FetchTournamentInvite(form.RequestID)
	if err != nil {
		ic.log.Error("error in FetchTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in FetchTournamentInvite", err.Error())
		return
	}
	tournamentId := InviteToCancel.TournamentID

	tournamentInfo, err := ic.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		ic.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	moderatorUsernames := strings.Split(tournamentInfo.Roles, ",")
	if tournamentInfo.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Only moderators can perform this action", "")
		return
	}

	err = ic.tsvc.DB.DeleteTournamentInvite(form.RequestID)
	if err != nil {
		ic.log.Error("error in DeleteTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in DeleteTournamentInvite", err.Error())
		return
	}

	outputInvites, err := ic.tsvc.DB.ListInviteByTournament(tournamentId)
	if err != nil {
		ic.log.Error("error in ListInviteByTournament", err)
		InternalErrorResponseV2(c, "error in ListInviteByTournament", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"inviteList":   outputInvites,
	})
}

type FormInvitationsAcceptOrDecline struct {
	RequestID string `json:"requestId"`
}

func (ic *InvitationController) HandleInvitationDecline(c *gin.Context) {
	user := ic.userFromContext(c)
	if user == nil {
		ic.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormInvitationsAcceptOrDecline{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.RequestID == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	inviteToDelete, err := ic.tsvc.DB.FetchTournamentInvite(form.RequestID)
	if err != nil {
		ic.log.Error("error in FetchTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in FetchTournamentInvite", err.Error())
		return
	}
	if inviteToDelete.Sub != user.Sub {
		BadRequestResponseV2(c, "this invite request is not for you", "")
		return
	}

	err = ic.tsvc.DB.DeleteTournamentInvite(form.RequestID)
	if err != nil {
		ic.log.Error("error in DeleteTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in DeleteTournamentInvite", err.Error())
		return
	}

	outputUserInvites, err := ic.tsvc.DB.ListActiveInvitesByUser(user)
	if err != nil {
		ic.log.Error("error in ListActiveInvitesByUser", err)
		InternalErrorResponseV2(c, "error in ListActiveInvitesByUser", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"userInviteList": outputUserInvites,
		"tournamentId":   inviteToDelete.TournamentID,
	})
}

func (ic *InvitationController) HandleInvitationAccept(c *gin.Context) {
	user := ic.userFromContext(c)
	if user == nil {
		ic.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormInvitationsAcceptOrDecline{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.RequestID == "" {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}

	inviteToDelete, err := ic.tsvc.DB.FetchTournamentInvite(form.RequestID)
	if err != nil {
		ic.log.Error("error in FetchTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in FetchTournamentInvite", err.Error())
		return
	}
	if inviteToDelete.Sub != user.Sub {
		BadRequestResponseV2(c, "this invite request is not for you", "")
		return
	}

	existingPlayRecords, err := ic.tsvc.DB.GetPlayRecordByTournamentAndPublicAddress(inviteToDelete.TournamentID, inviteToDelete.UserAccount.PublicAddress)
	if err != nil {
		ic.log.Error("error in GetPlayRecordByTournamentAndPublicAddress: ", err)
		InternalErrorResponseV2(c, "error in GetPlayRecordByTournamentAndPublicAddress", err.Error())
		return
	}
	if len(existingPlayRecords) > 0 {
		BadRequestResponseV2(c, "player has already registered in this tournament", "")
		return
	}

	tournamentInfo, err := ic.tsvc.DB.GetTournamentData(inviteToDelete.TournamentID)
	if err != nil {
		ic.log.Error("error getting tournament", err)
		InternalErrorResponseV2(c, "error getting tournament", "")
		return
	}

	if tournamentInfo.Status != string(tournament.REGISTRATION) {
		BadRequestResponseV2(c, "tournament is not in Registration status", "")
		return
	}

	if tournamentInfo.EntryFee != "0" { // check user balance if there is entry fee
		err = ic.checkEth(user)
		if err != nil {
			ic.log.Errorf("error in checkEth for user - %s: %s", inviteToDelete.UserAccount.Username, err.Error())
			InternalErrorResponseV2(c, "error in checkEth for applicant", err.Error())
			return
		}

		balance, err := ic.bsc.GetSPWNBalance(user.PublicAddress, tournamentInfo.FeeType)
		if err != nil {
			ic.log.Error("cannot get SPWN balance")
			InternalErrorResponseV2(c, "cannot get SPWN balance", err.Error())
			return
		}
		truncatedBalance, _ := balance.Int64()

		entryFeeInt, err := strconv.ParseInt(tournamentInfo.EntryFee, 10, 64)
		if err != nil {
			ic.log.Error("error in converting entryFee to int64")
			InternalErrorResponseV2(c, "error in converting entryFee to int64", "")
			return
		}
		if truncatedBalance < entryFeeInt {
			BadRequestResponseV2(c, "insufficient balance", "")
			return
		}
	}

	err = ic.tsvc.DB.DeleteTournamentInvite(form.RequestID)
	if err != nil {
		ic.log.Error("error in DeleteTournamentInvite: ", err)
		InternalErrorResponseV2(c, "error in DeleteTournamentInvite", err.Error())
		return
	}

	err = ic.tsvc.RegisterUserToTournament(user, tournamentInfo)
	if err != nil {
		ic.log.Errorf("error in RegisterUserToTournament: %v", err)
		InternalErrorResponseV2(c, "error in RegisterUserToTournament", err.Error())
		return
	}

	outputUserInvites, err := ic.tsvc.DB.ListActiveInvitesByUser(user)
	if err != nil {
		ic.log.Error("error in ListActiveInvitesByUser", err)
		InternalErrorResponseV2(c, "error in ListActiveInvitesByUser", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"userInviteList": outputUserInvites,
		"tournamentId":   inviteToDelete.TournamentID,
	})
}

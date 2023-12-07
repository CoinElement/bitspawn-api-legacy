/*

 */

package controllers

import (
	"github.com/bitspawngg/bitspawn-api/models"
	organizer "github.com/bitspawngg/bitspawn-api/services/match"
	"github.com/gin-gonic/gin"
)

type RolesInfo struct {
	DisplayName string `json:"displayName"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (tc *TournamentController) HandleV2TournamentDetails(c *gin.Context) {
	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tournament, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error getting tournament", err)
		InternalErrorResponseV2(c, "error getting tournament", "")
		return
	}

	tournamentForResponse, err := tc.DB.FormatTournamentForResponse(*tournament)
	if err != nil {
		tc.log.Error("error in FormatTournamentForResponse", err)
		InternalErrorResponseV2(c, "error formatting tournament", "")
		return
	}

	matches, err := tc.tsvc.DB.GetMatchesByTournament(tournamentID)
	if err != nil {
		tc.log.Error("error in GetMatchesByTournament: ", err)
		InternalErrorResponseV2(c, "error in GetMatchesByTournament", err.Error())
		return
	}

	scoreboard := organizer.GetScoreboard(matches)
	responseData := gin.H{
		"tournamentDetails": tournamentForResponse,
		"scoreboard":        scoreboard,
	}

	SuccessResponseV2(c, responseData)
}

func (tc *TournamentController) HandleTournamentFetchUserStatus(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		tc.InternalErrorResponse(c, "user not found")
		return
	}

	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		tc.BadRequestResponse(c, "missing input tournament id")
		return
	}

	tournament, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error getting tournament", err)
		tc.DBErrorResponse(c, "error getting tournament")
		return
	}

	status := tournament.Status

	playRecords, err := tc.DB.GetPlayRecordByTournamentAndPublicAddress(tournamentID, user.PublicAddress)
	if err != nil {
		tc.log.Error("cannot get registration record from database: ", err)
		tc.InternalErrorResponse(c, "cannot get registration record from database")
		return
	}
	isRegistered := len(playRecords) > 0

	organizerAction := "none"
	userAction := "none"
	if tournament.OrganizerID == user.Username {
		if status == "Registration" {
			organizerAction = "cancel"
		} else if status == "Ready" {
			if tournament.Status == "Completing" { // check more detailed status in DB: "Ready/Started/Completing"
				organizerAction = "complete"
			}
		}
	}
	if status == "Registration" {
		if !isRegistered {
			userAction = "register"
		} else {
			userAction = "unregister"
		}
	}

	responseData := gin.H{
		"TournamentDetails": tournament,
		"status":            status,
		"organizerAction":   organizerAction,
		"userAction":        userAction,
	}

	tc.SuccessResponse(c, "", responseData)
}

func (tc *TournamentController) HandleV2TournamentRegisterStatus(c *gin.Context) {
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

	_, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData: ", err)
		InternalErrorResponseV2(c, "error in GetTournamentData", err.Error())
		return
	}

	playRecords, err := tc.DB.GetPlayRecordByTournamentAndPublicAddress(tournamentID, user.PublicAddress)
	if err != nil {
		tc.log.Error("cannot get registration record from database: ", err)
		InternalErrorResponseV2(c, "cannot get registration record from database", err.Error())
		return
	}
	isRegistered := "NONE"
	if len(playRecords) > 0 {
		isRegistered = "JOINED"
	} else {
		oldInvitations, err := tc.tsvc.DB.FindTournamentInvite(tournamentID, user.Sub)
		if err != nil {
			tc.log.Error("error in FindTournamentInvite: ", err)
			InternalErrorResponseV2(c, "error in FindTournamentInvite", err.Error())
			return
		}
		if len(oldInvitations) > 0 {
			isRegistered = "INVITE_PENDING"
		} else {
			oldApplications, err := tc.tsvc.DB.FindTournamentApplication(tournamentID, user.Sub)
			if err != nil {
				tc.log.Error("error in FindTournamentApplication: ", err)
				InternalErrorResponseV2(c, "error in FindTournamentApplication", err.Error())
				return
			}
			if len(oldApplications) > 0 {
				isRegistered = "JOIN_PENDING"
			}
		}
	}

	SuccessResponseV2(c, gin.H{
		"tournamentId":  tournamentID,
		"hasRegistered": isRegistered,
	})
}

func (tc *TournamentController) HandleTournamentListMatch(c *gin.Context) {
	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	_, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		if err.Error() == "record not found" {
			BadRequestResponseV2(c, "tournament does not exist", "")
			return
		} else {
			tc.log.Error("error getting tournament: ", err)
			InternalErrorResponseV2(c, "error getting tournament", err.Error())
			return
		}
	}

	matches, err := tc.tsvc.DB.GetMatchesByTournament(tournamentID)
	if err != nil {
		tc.log.Error("error in GetMatchesByTournament: ", err)
		InternalErrorResponseV2(c, "error in GetMatchesByTournament", err.Error())
		return
	}

	matchesOutput, err := tc.DB.FormatMatchesForOutput(matches)
	if err != nil {
		tc.log.Error("error in FormatMatchesForOutput: ", err)
		InternalErrorResponseV2(c, "error in FormatMatchesForOutput", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentID,
		"matchList":    matchesOutput,
	})
}

func (tc *TournamentController) HandleV2TournamentListParticipantsById(c *gin.Context) {
	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	_, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", tournamentID, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}

	allParticipants, err := tc.DB.GetAllParticipantsInfo(tournamentID)
	if err != nil {
		tc.log.Error("error getting tournament participants", tournamentID, err)
		InternalErrorResponseV2(c, "error getting tournament participants", "")
		return
	}
	assignedParticipants, err := tc.DB.GetAssignedParticipants(tournamentID)
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
		"tournamentId":           tournamentID,
		"participantsList":       allParticipants,
		"assignedParticipants":   assignedParticipants,
		"unassignedParticipants": unassignedParticipants,
	})
}

func contains(array []string, element string) bool {
	for i := range array {
		if array[i] == element {
			return true
		}
	}
	return false
}

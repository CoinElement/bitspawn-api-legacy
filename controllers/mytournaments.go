/*

 */

package controllers

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/bitspawngg/bitspawn-api/models"
)

func (tc *TournamentController) HandleV2MyTournaments(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	allHostedTournaments, err := tc.DB.GetUserOrganisedTournaments(user.Username)
	if err != nil {
		tc.log.Error("error in GetUserOrganisedTournaments: ", err)
		InternalErrorResponseV2(c, "error getting my tournaments", err.Error())
		return
	}

	allRegisteredTournaments, err := tc.DB.GetRegisteredTournaments(user.Username)
	if err != nil {
		tc.log.Error("error in GetRegisteredTournaments: ", err)
		InternalErrorResponseV2(c, "error in GetRegisteredTournaments", err.Error())
		return
	}
	myRegisteredTournaments := []models.TournamentResponse{}
	myPastRegisters := []models.TournamentResponse{}
	for _, rTourney := range allRegisteredTournaments {
		if strings.ToUpper(rTourney.Status) == "COMPLETED" ||
			strings.ToUpper(rTourney.Status) == "CANCELLED" ||
			strings.ToUpper(rTourney.Status) == "PAYOUT" {
			myPastRegisters = append(myPastRegisters, rTourney)
		} else {
			myRegisteredTournaments = append(myRegisteredTournaments, rTourney)
		}
	}

	draftTournaments := []models.TournamentResponse{}
	myTournamentsData := []models.TournamentResponse{}
	myPastHosts := []models.TournamentResponse{}
	for _, hTourney := range allHostedTournaments {
		if strings.ToUpper(hTourney.Status) == "DRAFT" {
			draftTournaments = append(draftTournaments, hTourney)
		} else if strings.ToUpper(hTourney.Status) == "COMPLETED" ||
			strings.ToUpper(hTourney.Status) == "CANCELLED" ||
			strings.ToUpper(hTourney.Status) == "PAYOUT" {
			myPastHosts = append(myPastHosts, hTourney)
		} else {
			myTournamentsData = append(myTournamentsData, hTourney)
		}
	}

	SuccessResponseV2(c, gin.H{
		"hostedDraft":      draftTournaments,
		"hostedActive":     myTournamentsData,
		"hostedClosed":     myPastHosts,
		"registeredActive": myRegisteredTournaments,
		"registeredClosed": myPastRegisters,
	})
}

func (tc *TournamentController) HandleV2TournamentListRegistered(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	allRegisteredTournaments, err := tc.DB.GetRegisteredTournaments(user.Username)
	if err != nil {
		tc.log.Error("error getting my registered tournaments", err)
		InternalErrorResponseV2(c, "error getting my registered tournaments", err.Error())
		return
	}
	activeApplications, err := tc.tsvc.DB.ListActiveApplicationsByUser(user)
	if err != nil {
		tc.log.Error("error in ListActiveApplicationsByUser", err)
		InternalErrorResponseV2(c, "error in ListActiveApplicationsByUser", err.Error())
		return
	}
	activeInvites, err := tc.tsvc.DB.ListActiveInvitesByUser(user)
	if err != nil {
		tc.log.Error("error in ListActiveInvitesByUser", err)
		InternalErrorResponseV2(c, "error in ListActiveInvitesByUser", err.Error())
		return
	}
	allRegisteredTournamentIds := make([]string, 0, len(allRegisteredTournaments))
	for _, t := range allRegisteredTournaments {
		allRegisteredTournamentIds = append(allRegisteredTournamentIds, t.TournamentID)
	}
	activeApplicationIds := make([]string, 0, len(activeApplications))
	for _, v := range activeApplications {
		activeApplicationIds = append(activeApplicationIds, v.TournamentID)
	}
	activeInviteIds := make([]string, 0, len(activeInvites))
	for _, v := range activeInvites {
		activeInviteIds = append(activeInviteIds, v.TournamentID)
	}

	SuccessResponseV2(c, gin.H{
		"registeredList": allRegisteredTournamentIds,
		"joinList":       activeApplicationIds,
		"inviteList":     activeInviteIds,
	})
}

func (tc *TournamentController) HandleV2MyModeratorTournaments(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "page parameter is wrong", err.Error())
		return
	}

	perPageString := c.DefaultQuery("perPage", strconv.Itoa(PER_PAGE))
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		BadRequestResponseV2(c, "perPage parameter is wrong", err.Error())
		return
	}
	roleUser := "%" + user.Username + "%"

	totalPages, err := tc.DB.CountMyModeratorTournamentPages(roleUser, perPage)
	if err != nil {
		tc.log.Error("error in CountMyModeratorTournamentPages: ", err)
		InternalErrorResponseV2(c, "error in CountMyModeratorTournamentPages", err.Error())
		return
	}

	tournaments, err := tc.DB.GetUserModeratorTournaments(roleUser, page, perPage)
	if err != nil {
		tc.log.Error("error getting my moderator tournaments", err)
		tc.DBErrorResponse(c, "error getting my moderator tournaments")
		return
	}
	tournamentResponses := []models.TournamentResponse{}
	for _, t := range tournaments {
		tResponse, err := tc.DB.FormatTournamentForResponse(t)
		if err != nil {
			tc.log.Error("error in Format tournament for response: ", err)
			InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
			return
		}
		tournamentResponses = append(tournamentResponses, *tResponse)
	}

	SuccessResponseV2(c, gin.H{
		"tournamentList": tournamentResponses,
		"totalPages":     totalPages,
	})
}

func (tc *TournamentController) HandleV2TournamentListParticipants(c *gin.Context) {
	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "page parameter is wrong", err.Error())
		return
	}
	if page < 1 {
		BadRequestResponseV2(c, "page cannot be less than one", "")
		return
	}
	perPageString := c.DefaultQuery("perPage", strconv.Itoa(PER_PAGE_LONG))
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		BadRequestResponseV2(c, "perPage parameter is wrong", err.Error())
		return
	}
	if perPage < 1 {
		BadRequestResponseV2(c, "perPage cannot be less than one", "")
		return
	}
	gameTypeString := c.DefaultQuery("gameType", "")

	allParticipants, err := tc.DB.PlayedUserAvatar(gameTypeString)
	if err != nil {
		tc.log.Error("error in PlayedUserAvatar: ", err)
		InternalErrorResponseV2(c, "error in PlayedUserAvatar", err.Error())
		return
	}
	totalPages := math.Ceil(float64(len(allParticipants)) / float64(perPage))
	if totalPages < 1 {
		SuccessResponseV2(c, gin.H{
			"participantList": []models.UserInfo{},
			"totalPages":      1,
		})
		return
	}
	if float64(page) > totalPages {
		BadRequestResponseV2(c, fmt.Sprintf("page exceeds total page of %v", totalPages), "")
		return
	}
	outputParticipants := allParticipants[(page-1)*perPage:]
	if len(outputParticipants) > perPage {
		outputParticipants = outputParticipants[:perPage]
	}

	SuccessResponseV2(c, gin.H{
		"participantList": outputParticipants,
		"totalPages":      totalPages,
	})
}

func (tc *TournamentController) HandleV2TournamentHostedListParticipants(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "page parameter is wrong", err.Error())
		return
	}
	if page < 1 {
		BadRequestResponseV2(c, "page cannot be less than one", "")
		return
	}
	perPageString := c.DefaultQuery("perPage", strconv.Itoa(PER_PAGE_LONG))
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		BadRequestResponseV2(c, "perPage parameter is wrong", err.Error())
		return
	}
	if perPage < 1 {
		BadRequestResponseV2(c, "perPage cannot be less than one", "")
		return
	}
	gameTypeString := c.DefaultQuery("gameType", "")

	allParticipants, err := tc.DB.PlayedUserAvatar(gameTypeString)
	if err != nil {
		tc.log.Error("error in PlayedUserAvatar: ", err)
		InternalErrorResponseV2(c, "error in PlayedUserAvatar", err.Error())
		return
	}
	totalPages := math.Ceil(float64(len(allParticipants)) / float64(perPage))
	if totalPages < 1 {
		SuccessResponseV2(c, gin.H{
			"participantList": []models.UserInfo{},
			"totalPages":      1,
		})
		return
	}
	if float64(page) > totalPages {
		BadRequestResponseV2(c, fmt.Sprintf("page exceeds total page of %v", totalPages), "")
		return
	}
	outputParticipants := allParticipants[(page-1)*perPage:]
	if len(outputParticipants) > perPage {
		outputParticipants = outputParticipants[:perPage]
	}

	SuccessResponseV2(c, gin.H{
		"participantList": outputParticipants,
		"totalPages":      totalPages,
	})
}

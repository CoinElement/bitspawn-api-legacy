/*

 */

package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bitspawngg/bitspawn-api/models"
)

type TournamentOutputData struct {
	TournamentInfo models.TournamentData
	UserAvatars    []models.UserInfo
}

type ResponseV2ListTournaments struct {
	TournamentList []models.TournamentResponse `json:"tournamentList"`
}

func (tc *TournamentController) HandleV2TournamentListOpen(c *gin.Context) {

	var sortMethod string
	pageString := c.DefaultQuery("page", "1")
	perPageString := c.DefaultQuery("perPage", "50")
	filterString := c.DefaultQuery("tournamentName", "")
	sortString := c.DefaultQuery("sortBy", "tournamentDateDesc")
	gameType := c.DefaultQuery("gameType", "")
	gameSubType := c.DefaultQuery("gameSubType", "")
	totalPrizePool := c.DefaultQuery("totalPrizePool", "0")
	currentTimeString := time.Now().UTC().String()
	currentTime := strings.Split(currentTimeString, ".")
	timeToQuery := currentTime[0] + "/" + "9999-12-31 23:59:59"
	tournamentDate := c.DefaultQuery("tournamentDate", timeToQuery)
	prizePool, err := strconv.Atoi(totalPrizePool)
	if err != nil {
		BadRequestResponseV2(c, "totalPrizePool is not an integer", "")
		return
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "page is not an integer", "")
		return
	}
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		BadRequestResponseV2(c, "perPage is not an integer", "")
		return
	}
	tournamentName := "%" + filterString + "%"
	tDate := strings.Split(tournamentDate, "/")

	switch sortString {
	case "startDateDesc":
		sortMethod = "tournament_date desc"
	case "startDateAsc":
		sortMethod = "tournament_date asc"
	case "entryFeeDesc":
		sortMethod = "entry_fee desc"
	case "entryFeeAsc":
		sortMethod = "entry_fee asc"
	case "prizePoolDesc":
		sortMethod = "total_prize_pool desc"
	case "prizePoolAsc":
		sortMethod = "total_prize_pool asc"
	default:
		sortMethod = "tournament_date desc"
	}

	tournaments, err := tc.DB.ListTournamentsByStatus(tournamentName, gameType, gameSubType, prizePool, sortMethod, tDate[0], tDate[1], page, perPage, []string{"REGISTRATION", "Registration"})
	if err != nil {
		tc.log.Error("error getting open tournaments", err)
		InternalErrorResponseV2(c, "error getting open tournaments", err.Error())
		return
	}

	tournamentsBeforeCutoff := make([]models.TournamentResponse, 0)
	for _, t := range tournaments {
		if time.Now().Before(t.CutoffDate) {
			toInsert, err := tc.DB.FormatTournamentForResponse(t)
			if err != nil {
				tc.log.Error("error formatting tournament "+t.TournamentID+": ", err)
				InternalErrorResponseV2(c, "error formatting tournament "+t.TournamentID, "")
				return
			}
			tournamentsBeforeCutoff = append(tournamentsBeforeCutoff, *toInsert)
		}
	}

	SuccessResponseV2(c, ResponseV2ListTournaments{
		TournamentList: tournamentsBeforeCutoff,
	})
}

func (tc *TournamentController) HandleV2TournamentListCutoff(c *gin.Context) {

	var sortMethod string
	pageString := c.DefaultQuery("page", "1")
	perPageString := c.DefaultQuery("perPage", "50")
	filterString := c.DefaultQuery("tournamentName", "")
	sortString := c.DefaultQuery("sortBy", "tournamentDateDesc")
	gameType := c.DefaultQuery("gameType", "")
	gameSubType := c.DefaultQuery("gameSubType", "")
	totalPrizePool := c.DefaultQuery("totalPrizePool", "0")
	currentTimeString := time.Now().UTC().String()
	currentTime := strings.Split(currentTimeString, ".")
	timeToQuery := currentTime[0] + "/" + "9999-12-31 23:59:59"
	tournamentDate := c.DefaultQuery("tournamentDate", timeToQuery)
	prizePool, err := strconv.Atoi(totalPrizePool)
	if err != nil {
		BadRequestResponseV2(c, "totalPrizePool is not an integer", "")
		return
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "page is not an integer", "")
		return
	}
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		BadRequestResponseV2(c, "perPage is not an integer", "")
		return
	}
	tournamentName := "%" + filterString + "%"
	tDate := strings.Split(tournamentDate, "/")

	switch sortString {
	case "startDateDesc":
		sortMethod = "tournament_date desc"
	case "startDateAsc":
		sortMethod = "tournament_date asc"
	case "entryFeeDesc":
		sortMethod = "entry_fee desc"
	case "entryFeeAsc":
		sortMethod = "entry_fee asc"
	case "prizePoolDesc":
		sortMethod = "total_prize_pool desc"
	case "prizePoolAsc":
		sortMethod = "total_prize_pool asc"
	default:
		sortMethod = "tournament_date desc"
	}

	tournaments, err := tc.DB.ListTournamentsByStatus(tournamentName, gameType, gameSubType, prizePool, sortMethod, tDate[0], tDate[1], page, perPage,
		[]string{"REGISTRATION", "Registration"})
	if err != nil {
		tc.log.Error("error getting tournaments after cutoff", err)
		InternalErrorResponseV2(c, "error getting tournaments after cutoff", err.Error())
		return
	}

	tournamentsAfterCutoff := make([]models.TournamentResponse, 0)
	for _, t := range tournaments {
		if time.Now().After(t.CutoffDate) {
			toInsert, err := tc.DB.FormatTournamentForResponse(t)
			if err != nil {
				tc.log.Error("error formatting tournament "+t.TournamentID+": ", err)
				InternalErrorResponseV2(c, "error formatting tournament "+t.TournamentID, "")
				return
			}
			tournamentsAfterCutoff = append(tournamentsAfterCutoff, *toInsert)
		}
	}

	SuccessResponseV2(c, ResponseV2ListTournaments{
		TournamentList: tournamentsAfterCutoff,
	})
}

func (tc *TournamentController) HandleV2TournamentListActive(c *gin.Context) {

	var sortMethod string
	pageString := c.DefaultQuery("page", "1")
	perPageString := c.DefaultQuery("perPage", "50")
	filterString := c.DefaultQuery("tournamentName", "")
	sortString := c.DefaultQuery("sortBy", "tournamentDateDesc")
	gameType := c.DefaultQuery("gameType", "")
	gameSubType := c.DefaultQuery("gameSubType", "")
	totalPrizePool := c.DefaultQuery("totalPrizePool", "0")
	currentTimeString := time.Now().UTC().String()
	currentTime := strings.Split(currentTimeString, ".")
	timeToQuery := currentTime[0] + "/" + "9999-12-31 23:59:59"
	tournamentDate := c.DefaultQuery("tournamentDate", timeToQuery)
	prizePool, err := strconv.Atoi(totalPrizePool)
	if err != nil {
		BadRequestResponseV2(c, "totalPrizePool is not an integer", "")
		return
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "page is not an integer", "")
		return
	}
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		BadRequestResponseV2(c, "perPage is not an integer", "")
		return
	}
	tournamentName := "%" + filterString + "%"
	tDate := strings.Split(tournamentDate, "/")

	switch sortString {
	case "startDateDesc":
		sortMethod = "tournament_date desc"
	case "startDateAsc":
		sortMethod = "tournament_date asc"
	case "entryFeeDesc":
		sortMethod = "entry_fee desc"
	case "entryFeeAsc":
		sortMethod = "entry_fee asc"
	case "prizePoolDesc":
		sortMethod = "total_prize_pool desc"
	case "prizePoolAsc":
		sortMethod = "total_prize_pool asc"
	default:
		sortMethod = "tournament_date desc"
	}

	tournaments, err := tc.DB.ListTournamentsByStatus(tournamentName, gameType, gameSubType, prizePool, sortMethod, tDate[0], tDate[1], page, perPage,
		[]string{"STARTED", "Started", "READY", "Ready"})
	if err != nil {
		tc.log.Error("error getting active tournaments", err)
		InternalErrorResponseV2(c, "error getting active tournaments", "")
		return
	}

	tournamentsResponse := make([]models.TournamentResponse, 0)
	for _, t := range tournaments {
		toInsert, err := tc.DB.FormatTournamentForResponse(t)
		if err != nil {
			tc.log.Error("error formatting tournament "+t.TournamentID+": ", err)
			InternalErrorResponseV2(c, "error formatting tournament "+t.TournamentID, "")
			return
		}
		tournamentsResponse = append(tournamentsResponse, *toInsert)
	}

	SuccessResponseV2(c, ResponseV2ListTournaments{
		TournamentList: tournamentsResponse,
	})
}

func (tc *TournamentController) HandleV2TournamentListClosed(c *gin.Context) {

	var sortMethod string
	pageString := c.DefaultQuery("page", "1")
	perPageString := c.DefaultQuery("perPage", "50")
	filterString := c.DefaultQuery("tournamentName", "")
	sortString := c.DefaultQuery("sortBy", "tournamentDateDesc")
	gameType := c.DefaultQuery("gameType", "")
	gameSubType := c.DefaultQuery("gameSubType", "")
	totalPrizePool := c.DefaultQuery("totalPrizePool", "0")
	currentTimeString := time.Now().UTC().String()
	currentTime := strings.Split(currentTimeString, ".")
	timeToQuery := currentTime[0] + "/" + "9999-12-31 23:59:59"
	tournamentDate := c.DefaultQuery("tournamentDate", timeToQuery)
	prizePool, err := strconv.Atoi(totalPrizePool)
	if err != nil {
		BadRequestResponseV2(c, "totalPrizePool is not an integer", "")
		return
	}
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "page is not an integer", "")
		return
	}
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		BadRequestResponseV2(c, "perPage is not an integer", "")
		return
	}
	tournamentName := "%" + filterString + "%"
	tDate := strings.Split(tournamentDate, "/")

	switch sortString {
	case "startDateDesc":
		sortMethod = "tournament_date desc"
	case "startDateAsc":
		sortMethod = "tournament_date asc"
	case "entryFeeDesc":
		sortMethod = "entry_fee desc"
	case "entryFeeAsc":
		sortMethod = "entry_fee asc"
	case "prizePoolDesc":
		sortMethod = "total_prize_pool desc"
	case "prizePoolAsc":
		sortMethod = "total_prize_pool asc"
	default:
		sortMethod = "tournament_date desc"
	}

	tournaments, err := tc.DB.ListTournamentsByStatus(tournamentName, gameType, gameSubType, prizePool, sortMethod, tDate[0], tDate[1], page, perPage,
		[]string{"COMPLETED", "Completed", "CANCELLED,", "Cancelled", "PAYOUT", "Completing"})
	if err != nil {
		tc.log.Error("error getting closed tournaments", err)
		InternalErrorResponseV2(c, "error getting closed tournaments", "")
		return
	}

	tournamentsResponse := make([]models.TournamentResponse, 0)
	for _, t := range tournaments {
		toInsert, err := tc.DB.FormatTournamentForResponse(t)
		if err != nil {
			tc.log.Error("error formatting tournament "+t.TournamentID+": ", err)
			InternalErrorResponseV2(c, "error formatting tournament "+t.TournamentID, "")
			return
		}
		tournamentsResponse = append(tournamentsResponse, *toInsert)
	}

	SuccessResponseV2(c, ResponseV2ListTournaments{
		TournamentList: tournamentsResponse,
	})
}

func (tc *TournamentController) HandleV2TournamentCount(c *gin.Context) {
	gameType := c.DefaultQuery("gameType", "")
	openCount, err := tc.DB.CountOpenTournaments(gameType)
	if err != nil {
		tc.log.Error("error in CountOpenTournaments: " + err.Error())
		InternalErrorResponseV2(c, "error in CountOpenTournaments", err.Error())
		return
	}
	cutoffCount, err := tc.DB.CountCutoffTournaments(gameType)
	if err != nil {
		tc.log.Error("error in CountCutoffTournaments: " + err.Error())
		InternalErrorResponseV2(c, "error in CountCutoffTournaments", err.Error())
		return
	}
	activeCount, err := tc.DB.CountTournamentsByStatus(gameType, []string{"STARTED", "Started", "READY", "Ready"})
	if err != nil {
		tc.log.Error("error in Count active tournaments: " + err.Error())
		InternalErrorResponseV2(c, "error in Count active tournaments:", err.Error())
		return
	}
	closedCount, err := tc.DB.CountTournamentsByStatus(gameType, []string{"COMPLETED", "Completed", "CANCELLED,", "Cancelled", "PAYOUT", "Completing"})
	if err != nil {
		tc.log.Error("error in Count closed tournaments: " + err.Error())
		InternalErrorResponseV2(c, "error in Count closed tournaments", err.Error())
		return
	}
	SuccessResponseV2(c, gin.H{
		"open":   openCount,
		"cutoff": cutoffCount,
		"active": activeCount,
		"closed": closedCount,
	})
}

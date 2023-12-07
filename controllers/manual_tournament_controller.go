package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type FormTournamentMatchScore struct {
	TournamentId string `json:"tournamentId"`
	MatchId      string `json:"matchId"`
	TeamOneScore int    `json:"teamOneScore"`
	TeamTwoScore int    `json:"teamTwoScore"`
}

func (tc *TournamentController) HandleV2TournamentMatchScore(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentMatchScore{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TournamentId == "" || form.MatchId == "" {
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

	if strings.ToUpper(tournamentInfo.Status) != "STARTED" {
		BadRequestResponseV2(c, "the tournament is not in Started status", "")
		return
	}

	err = tc.DB.OrganizerUpdateMatchScoresV2(tournamentId, form.MatchId, form.TeamOneScore, form.TeamTwoScore)
	if err != nil {
		tc.log.Errorf("error in ManualUpdateMatchScoresV2 - %s: %v", form.MatchId, err)
		InternalErrorResponseV2(c, "error in ManualUpdateMatchScoresV2", err.Error())
		return
	}

	err = tc.orgSvc.PrepareManualMatchesByTournament(tournamentId)
	if err != nil {
		tc.log.Errorf("error in PrepareManualMatchesByTournament - %s: %v", tournamentId, err)
		InternalErrorResponseV2(c, "error in PrepareManualMatchesByTournament", err.Error())
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

type FormTournamentMatchGameCreate struct {
	TeamNumber   int `json:"teamNumber"`
	TeamOneScore int `json:"teamOneScore"`
	TeamTwoScore int `json:"teamTwoScore"`
}

type IMatchGame struct {
	TeamOne IMatchGameTeam `json:"teamOne"`
	TeamTwo IMatchGameTeam `json:"teamTwo"`
}

type IMatchGameTeam struct {
	Score int `json:"score "`
}

func (tc *TournamentController) HandleV2TournamentMatchGameCreate(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentId := c.Param("tournamentId")
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}
	matchId := c.Param("matchId")
	if matchId == "" {
		BadRequestResponseV2(c, "missing input match id", "")
		return
	}

	form := FormTournamentMatchGameCreate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TeamNumber == 0 {
		BadRequestResponseV2(c, "missing team number", "")
		return
	}
	if form.TeamNumber != 1 && form.TeamNumber != 2 {
		BadRequestResponseV2(c, "team number must be either 1 or 2", "")
		return
	}

	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}
	if strings.ToUpper(tournamentInfo.Status) != "STARTED" {
		BadRequestResponseV2(c, "the tournament is not in Started status", "")
		return
	}

	match, err := tc.tsvc.DB.GetMatchById(matchId)
	if err != nil {
		tc.log.Error("error in GetMatchById for match ", matchId, ": ", err)
		InternalErrorResponseV2(c, "cannot get match data", err.Error())
		return
	}
	if match.TournamentID != tournamentId {
		BadRequestResponseV2(c, "match does not belong in this tournament", "")
		return
	}

	existingMatchGames, err := tc.tsvc.DB.GetMatchGamesByMatch(matchId)
	if err != nil {
		tc.log.Error("error in GetMatchGamesByMatch ", matchId, ": ", err)
		InternalErrorResponseV2(c, "error in GetMatchGamesByMatch: ", err.Error())
		return
	}

	newMatchGame, err := tc.tsvc.DB.CreateMatchGame(match, existingMatchGames, form.TeamOneScore, form.TeamTwoScore, form.TeamNumber)
	if err != nil {
		tc.log.Errorf("error in UpdateMatchGame - %s: %v", matchId, err)
		InternalErrorResponseV2(c, "error in UpdateMatchGame", err.Error())
		return
	}
	allMatchGames := append(existingMatchGames, newMatchGame)
	for i, mg := range allMatchGames {
		if mg.ScreenshotOne != "" {
			allMatchGames[i].ScreenshotOne = fmt.Sprintf("https://%s.s3.amazonaws.com/matchGames/%s", tc.SignURLManager.BucketName, mg.ScreenshotOne)
		}
		if mg.ScreenshotTwo != "" {
			allMatchGames[i].ScreenshotTwo = fmt.Sprintf("https://%s.s3.amazonaws.com/matchGames/%s", tc.SignURLManager.BucketName, mg.ScreenshotTwo)
		}
	}

	err = tc.orgSvc.PrepareManualMatchesByTournament(tournamentId)
	if err != nil {
		tc.log.Errorf("error in PrepareManualMatchesByTournament - %s: %v", tournamentId, err)
		InternalErrorResponseV2(c, "error in PrepareManualMatchesByTournament", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"matchId":      matchId,
		"matchGames":   allMatchGames,
	})
}

func (tc *TournamentController) HandleV2TournamentMatchGameGet(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentId := c.Param("tournamentId")
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}
	matchId := c.Param("matchId")
	if matchId == "" {
		BadRequestResponseV2(c, "missing input match id", "")
		return
	}

	match, err := tc.tsvc.DB.GetMatchById(matchId)
	if err != nil {
		tc.log.Error("error in GetMatchById for match ", matchId, ": ", err)
		InternalErrorResponseV2(c, "cannot get match data", err.Error())
		return
	}
	if match.TournamentID != tournamentId {
		BadRequestResponseV2(c, "match does not belong in this tournament", "")
		return
	}

	outputMatches, err := tc.tsvc.DB.FormatMatchesForOutput([]models.Match{*match})
	if err != nil {
		tc.log.Error("error in FormatMatchesForOutput: ", err)
		InternalErrorResponseV2(c, "error in FormatMatchesForOutput", err.Error())
		return
	}
	if len(outputMatches) < 1 {
		tc.log.Error("Match details are missing")
		InternalErrorResponseV2(c, "Match details are missing", "")
		return
	}

	allMatchGames, err := tc.tsvc.DB.GetMatchGamesByMatch(matchId)
	if err != nil {
		tc.log.Error("error in GetMatchGamesByMatch ", matchId, ": ", err)
		InternalErrorResponseV2(c, "error in GetMatchGamesByMatch: ", err.Error())
		return
	}
	for i, mg := range allMatchGames {
		if mg.ScreenshotOne != "" {
			allMatchGames[i].ScreenshotOne = fmt.Sprintf("https://%s.s3.amazonaws.com/matchGames/%s", tc.SignURLManager.BucketName, mg.ScreenshotOne)
		}
		if mg.ScreenshotTwo != "" {
			allMatchGames[i].ScreenshotTwo = fmt.Sprintf("https://%s.s3.amazonaws.com/matchGames/%s", tc.SignURLManager.BucketName, mg.ScreenshotTwo)
		}
	}
	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"matchDetails": outputMatches[0],
		"matchGames":   allMatchGames,
	})
}

type FormTournamentMatchGameUpdate struct {
	PatchType    string `json:"patchType"`
	TeamNumber   int    `json:"teamNumber"`
	TeamOneScore *int   `json:"teamOneScore"`
	TeamTwoScore *int   `json:"teamTwoScore"`
}

func (tc *TournamentController) HandleV2TournamentMatchGameUpdate(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	tournamentId := c.Param("tournamentId")
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}
	matchId := c.Param("matchId")
	if matchId == "" {
		BadRequestResponseV2(c, "missing input match id", "")
		return
	}
	gameNumberStr := c.Param("gameNumber")
	if matchId == "" {
		BadRequestResponseV2(c, "missing input game number", "")
		return
	}
	gameNumber, err := strconv.Atoi(gameNumberStr)
	if err != nil {
		BadRequestResponseV2(c, "game number is not a number", "")
		return
	}
	if gameNumber < 1 {
		BadRequestResponseV2(c, "game number must be positive", "")
		return
	}

	form := FormTournamentMatchGameUpdate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}
	if form.PatchType == "" {
		BadRequestResponseV2(c, "missing patch type", "")
		return
	}
	if form.PatchType != "SCORE" && form.PatchType != "SCREENSHOT" {
		BadRequestResponseV2(c, "unsupported patch type", "")
		return
	}

	if form.TeamNumber == 0 {
		BadRequestResponseV2(c, "missing mandatory input parameter", "")
		return
	}
	if form.TeamNumber != 1 && form.TeamNumber != 2 {
		BadRequestResponseV2(c, "team number must be either 1 or 2", "")
		return
	}

	tournamentInfo, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData for tournament ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "cannot get tournament data", err.Error())
		return
	}
	if strings.ToUpper(tournamentInfo.Status) != "STARTED" {
		BadRequestResponseV2(c, "the tournament is not in Started status", "")
		return
	}

	match, err := tc.tsvc.DB.GetMatchById(matchId)
	if err != nil {
		tc.log.Error("error in GetMatchById for match ", matchId, ": ", err)
		InternalErrorResponseV2(c, "cannot get match data", err.Error())
		return
	}
	if match.TournamentID != tournamentId {
		BadRequestResponseV2(c, "match does not belong in this tournament", "")
		return
	}

	_, err = tc.tsvc.DB.GetMatchGame(matchId, gameNumber)
	if err != nil {
		tc.log.Errorf("error in GetMatchGame(%s, %d): %v", matchId, gameNumber, err)
		InternalErrorResponseV2(c, "error in GetMatchGame", err.Error())
		return
	}

	switch form.PatchType {
	case "SCORE":
		if form.TeamOneScore == nil || form.TeamTwoScore == nil {
			BadRequestResponseV2(c, "missing team score", "")
			return
		}
		_, _, err = tc.tsvc.DB.UpdateMatchGame(match, gameNumber, *form.TeamOneScore, *form.TeamTwoScore, form.TeamNumber)
		if err != nil {
			tc.log.Errorf("error in UpdateMatchGame - %s: %v", matchId, err)
			InternalErrorResponseV2(c, "error in UpdateMatchGame", err.Error())
			return
		}
	case "SCREENSHOT":
		screenshotFileName := uuid.NewV4().String()
		url, headers, err := tc.SignURLManager.GetSignedURL("matchGames/" + screenshotFileName)
		if err != nil {
			tc.log.Errorf("error in GetSignedURL - %s: %v", screenshotFileName, err)
			InternalErrorResponseV2(c, "error in GetSignedURL", err.Error())
			return
		}
		err = tc.tsvc.DB.UpdateMatchGameScreenshot(match, gameNumber, screenshotFileName, form.TeamNumber)
		if err != nil {
			tc.log.Errorf("error in UpdateMatchGameScreenshot - %s: %v", matchId, err)
			InternalErrorResponseV2(c, "error in UpdateMatchGameScreenshot", err.Error())
			return
		}
		SuccessResponseV2(c, struct {
			URL     string      `json:"url"`
			Headers http.Header `json:"headers"`
		}{
			url, headers,
		})
		return
	default:
		BadRequestResponseV2(c, "Unsupported patch type", "")
		return
	}

	allMatchGames, err := tc.tsvc.DB.GetMatchGamesByMatch(matchId)
	if err != nil {
		tc.log.Error("error in GetMatchGamesByMatch ", matchId, ": ", err)
		InternalErrorResponseV2(c, "error in GetMatchGamesByMatch", err.Error())
		return
	}

	err = tc.orgSvc.PrepareManualMatchesByTournament(tournamentId)
	if err != nil {
		tc.log.Errorf("error in PrepareManualMatchesByTournament - %s: %v", tournamentId, err)
		InternalErrorResponseV2(c, "error in PrepareManualMatchesByTournament", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentId": tournamentId,
		"matchId":      matchId,
		"gameNumber":   gameNumber,
		"matchGames":   allMatchGames,
	})
}

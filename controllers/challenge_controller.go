/*

 */

package controllers

import (
	"fmt"
	"github.com/bitspawngg/bitspawn-api/enum"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services/challenge"
	"github.com/bitspawngg/bitspawn-api/services/poa"
)

type ChallengeController struct {
	BaseController
	tsvc *challenge.ChallengeService
	bsc  *poa.BitspawnPoaClient
}

func NewChallengeController(bc *BaseController, tsvc *challenge.ChallengeService, bsc *poa.BitspawnPoaClient) *ChallengeController {

	return &ChallengeController{
		BaseController{
			Name: "challenge",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		},
		tsvc,
		bsc,
	}
}

type FormChallengeCreate struct {
	ChallengeName string         `json:"challengeName"`
	IsSponsored   bool           `json:"isSponsored"`
	ChallengeType string         `json:"challengeType"`
	Scoring       map[string]int `json:"scoring"`
	ChallengeRule string         `json:"challengeRule"`
	GameMode      string         `json:"gameMode"`
	NumberOfGames int            `json:"numberOfGames"`
	EntryOnce     bool           `json:"entryOnce"`
	FeeType       string         `json:"feeType"`

	MinParticipants     int64   `json:"minParticipants"`
	EntryFee            int64   `json:"entryFee"`
	TargetPrizePool     int64   `json:"targetPrizePool"`
	FeePercentage       int64   `json:"feePercentage"`
	OrganizerPercentage int64   `json:"organizerPercentage"`
	PrizeAllocation     []int64 `json:"prizeAllocation"`

	RegisterTimeWindow int    `json:"registerTimeWindow"`
	PlayTimeWindow     int    `json:"playTimeWindow"`
	Title              string `json:"title"`
	Authorization      string `json:"authorization"`
	MaxParticipants    int    `json:"maxParticipants"`
}

func (cc *ChallengeController) HandleAdminChallengeCreate(c *gin.Context) {
	//todo: remove this auth since we splited the admin api to correct auth already
	user, err := cc.DB.GetUserProfile(cc.tsvc.GetAdminUser())
	if err != nil {
		cc.log.Error("error getting admin user: ", err)
		cc.InternalErrorResponse(c, "error getting admin user")
		return
	}

	form := FormChallengeCreate{}

	if err := c.ShouldBindJSON(&form); err != nil {
		cc.BadRequestResponse(c, err.Error())
		return
	}
	// usernameHash := uuid.NewV5(uuid.FromStringOrNil("06d0f7f8-1a24-4f5f-9ff2-eb8d1dbf43e5"), user.Username)
	// if usernameHash.String() != form.Authorization {
	// 	cc.BadRequestResponse(c, "authorization failed")
	// 	return
	// }
	if form.ChallengeName == "" || form.ChallengeType == "" || form.FeeType == "" || form.Scoring == nil || form.MinParticipants == 0 || form.PrizeAllocation == nil || form.RegisterTimeWindow == 0 || form.PlayTimeWindow == 0 || form.GameMode == "" || form.NumberOfGames == 0 {
		cc.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	// TODO: input field validation
	switch form.FeeType {
	case enum.Credit.ToString():
		// Do nothing
	case enum.Spwn.ToString():
		// Do nothing
	case enum.Usdc.ToString():
		// Do nothing
	default:
		cc.BadRequestResponse(c, "unknown fee type")
		return
	}
	var percentCount int64
	for _, percent := range form.PrizeAllocation {
		percentCount += percent
	}
	if percentCount != 100 {
		cc.BadRequestResponse(c, "Prize allocation does not add up to 100%")
		return
	}
	// End of input field validation

	prizeString := []string{}
	for _, prize := range form.PrizeAllocation {
		p := fmt.Sprintf("%g", float64(prize))
		prizeString = append(prizeString, p)
	}

	scoring := make(map[string]*string)
	for category, multiplier := range form.Scoring {
		m := strconv.Itoa(multiplier)
		scoring[category] = &m
	}

	td := models.Challenge{ChallengeBase: models.ChallengeBase{
		ChallengeID:         uuid.NewV4(),
		ChallengeName:       form.ChallengeName,
		IsSponsored:         form.IsSponsored,
		ChallengeType:       form.ChallengeType,
		MinParticipants:     form.MinParticipants,
		ChallengeRule:       form.ChallengeRule,
		GameMode:            form.GameMode,
		NumberOfGames:       form.NumberOfGames,
		EntryOnce:           form.EntryOnce,
		EntryFee:            strconv.FormatInt(form.EntryFee, 10),
		FeeType:             form.FeeType,
		NumberOfWinners:     len(prizeString),
		PrizeAllocation:     strings.Join(prizeString, ","),
		CutoffDate:          time.Now().Add(time.Minute * time.Duration(form.RegisterTimeWindow)),
		StartDate:           time.Now(),
		Status:              "Registration",
		OrganizerID:         user.Username,
		TargetPrizePool:     form.TargetPrizePool,
		FeePercentage:       form.FeePercentage,
		OrganizerPercentage: form.OrganizerPercentage,
		MinuteTimeWindow:    form.PlayTimeWindow,
		Title:               form.Title,
		MaxParticipants:     form.MaxParticipants,
	},
		Scoring: scoring,
	}

	err = cc.tsvc.CreateStateMachine(&td, user)
	if err != nil {
		cc.log.Error("error in create challgenge: ", err)
		cc.InternalErrorResponse(c, "error in create challenge")
		return
	}

	if form.Title != "" {
		featuredObject := models.Featured{
			ObjectTitle: form.Title,
			ObjectID:    td.ChallengeID.String(),
			ObjectType:  "Challenge",
		}
		err = cc.DB.InsertFeaturedObject(featuredObject)
		if err != nil {
			cc.log.Error("error in InsertFeaturedObject: ", err)
			cc.InternalErrorResponse(c, "error in insert featured challenge")
			return
		}
	}

	t := gin.H{
		"challenge": td,
	}
	cc.SuccessResponse(c, "challenge created", t)
}

type FormChallengeComplete struct {
	ChallengeID string `json:"challengeId" binding:"required"`
}

func (cc *ChallengeController) HandleChallengeComplete(c *gin.Context) {

	user := cc.userFromContext(c)
	if user == nil {
		cc.log.Error("user not found")
		cc.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormChallengeComplete{}

	if err := c.ShouldBindJSON(&form); err != nil {
		cc.BadRequestResponse(c, err.Error())
		return
	}

	if form.ChallengeID == "" {
		cc.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}
	err := cc.tsvc.ExecuteStateMachine(form.ChallengeID, user.Username, "Complete", nil)
	if err != nil {
		cc.log.Error("error in Complete challenge ", form.ChallengeID, ": ", err)
		cc.InternalErrorResponse(c, "error in Complete challenge: "+err.Error())
		return
	}

	cc.SuccessResponse(c, "challenge completed successfully", "")
}

type FormChallengeCancel struct {
	ChallengeID string `json:"challengeId"`
}

func (cc *ChallengeController) HandleChallengeCancel(c *gin.Context) {
	user := cc.userFromContext(c)
	if user == nil {
		cc.log.Error("user not found")
		cc.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormChallengeCancel{}
	if err := c.ShouldBindJSON(&form); err != nil {
		cc.BadRequestResponse(c, err.Error())
		return
	}

	err := cc.tsvc.ExecuteStateMachine(form.ChallengeID, user.Username, "Cancel", nil)
	if err != nil {
		cc.log.Error("error in Cancel challenge ", form.ChallengeID, ": ", err)
		cc.InternalErrorResponse(c, "error in Cancel challenge: "+err.Error())
		return
	}

	cc.SuccessResponse(c, "challenge cancelled successfully", "")
}

func (cc *ChallengeController) HandleChallengeFund(c *gin.Context) {
	user := cc.userFromContext(c)
	if user == nil {
		cc.log.Error("user not found")
		cc.InternalErrorResponse(c, "user not found")
		return
	}

	form := challenge.Meta{}
	if err := c.ShouldBindJSON(&form); err != nil {
		cc.BadRequestResponse(c, err.Error())
		return
	}

	if form.ChallengeID == "" || form.TipAmount <= 0 {
		cc.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}
	err := cc.tsvc.ExecuteStateMachine(form.ChallengeID, user.Username, "Fund", &form)
	if err != nil {
		cc.log.Error("error in Fund challenge ", form.ChallengeID, ": ", err)
		cc.InternalErrorResponse(c, "error in Fund challenge: "+err.Error())
		return
	}

	cc.SuccessResponse(c, "challenge funded successfully", "")
}

func (cc *ChallengeController) HandleAdminChallengeComplete(c *gin.Context) {
	challengesToComplete, err := cc.DB.GetChallengesByStatus("Ready")
	if err != nil {
		cc.InternalErrorResponse(c, "error in GetChallengesByStatus(Ready): "+err.Error())
		return
	}

	challengesCompleted := []string{}
	challengesFailedToComplete := []string{}
	for _, challenge := range challengesToComplete {
		if time.Now().Before(challenge.CutoffDate.Add(time.Duration(challenge.MinuteTimeWindow) * time.Minute)) {
			// wait until time window passes after cutoff date
			continue
		}

		challengeId := challenge.ChallengeID.String()
		cc.log.Info("Completing challenge: ", challengeId)

		_, err := cc.DB.GetUserProfile(challenge.OrganizerID)
		if err != nil {
			cc.log.Error("error getting user ", challenge.OrganizerID, ": ", err)
			continue
		}
		err = cc.tsvc.ExecuteStateMachine(challengeId, challenge.OrganizerID, "Complete", nil)
		if err != nil {
			cc.log.Error("error in complete challenge: ", err)
			challengesFailedToComplete = append(challengesFailedToComplete, challengeId)
			continue
		}

		challengesCompleted = append(challengesCompleted, challengeId)
	}

	t := gin.H{
		"challengesCompleted":        challengesCompleted,
		"challengesFailedToComplete": challengesFailedToComplete,
	}
	cc.SuccessResponse(c, "challenges completed", t)
}

func (cc *ChallengeController) HandleAdminChallengeCancel(c *gin.Context) {
	challengesToCancel, err := cc.DB.GetChallengesByStatus("Registration")
	if err != nil {
		cc.log.Error(err)
		cc.InternalErrorResponse(c, "error in GetChallengesByStatus(Registration): "+err.Error())
		return
	}

	challengesCancelled := []string{}
	challengesFailedToCancel := []string{}
	for _, challenge := range challengesToCancel {
		if time.Now().Before(challenge.CutoffDate.Add(time.Duration(15) * time.Minute)) {
			// wait until after cutoff date + 15 min
			continue
		}

		challengeId := challenge.ChallengeID.String()
		cc.log.Info("Cancelling challenge: ", challengeId)

		_, err := cc.DB.GetUserProfile(challenge.OrganizerID)
		if err != nil {
			cc.log.Error("error getting user ", challenge.OrganizerID, ": ", err)
			continue
		}
		err = cc.tsvc.ExecuteStateMachine(challengeId, challenge.OrganizerID, "Cancel", nil)
		if err != nil {
			cc.log.Error("error in cancel challenge: ", err)
			challengesFailedToCancel = append(challengesFailedToCancel, challengeId)
			continue
		}

		challengesCancelled = append(challengesCancelled, challengeId)
	}

	t := gin.H{
		"challengesCancelled":      challengesCancelled,
		"challengesFailedToCancel": challengesFailedToCancel,
	}

	cc.SuccessResponse(c, "challenges cancelled", t)
}

func (cc *ChallengeController) HandleAdminChallengeFund(c *gin.Context) {
	challengesToFund, err := cc.DB.GetChallengesByStatus("Ready")
	if err != nil {
		cc.log.Error(err)
		cc.InternalErrorResponse(c, "error in GetChallengesByStatus(Registration): "+err.Error())
		return
	}

	challengesFunded := []string{}
	challengesFailedToFund := []string{}
	for _, ch := range challengesToFund {
		if int64(ch.ParticipantCount) < ch.MinParticipants || ch.EntryFee != "0" || ch.OrganizerContribute > 0 {
			// do not fund until participant count reaches minimum requirement
			// do not fund if players pay entry fee
			// do not fund if organizer has already funded
			continue
		}

		challengeId := ch.ChallengeID.String()
		cc.log.Info("Funding challenge: ", challengeId)

		_, err := cc.DB.GetUserProfile(ch.OrganizerID)
		if err != nil {
			cc.log.Error("error getting user ", ch.OrganizerID, ": ", err)
			continue
		}
		meta := challenge.Meta{
			TipAmount: int(ch.TargetPrizePool),
		}
		err = cc.tsvc.ExecuteStateMachine(challengeId, ch.OrganizerID, "Fund", &meta)
		if err != nil {
			cc.log.Error("error in fund challenge: ", ch.ChallengeID, ": ", err)
			challengesFailedToFund = append(challengesFailedToFund, challengeId)
			continue
		}

		challengesFunded = append(challengesFunded, challengeId)
	}

	t := gin.H{
		"challengesFunded":       challengesFunded,
		"challengesFailedToFund": challengesFailedToFund,
	}

	cc.SuccessResponse(c, "challenge funded successfully", t)
}

func (cc *ChallengeController) HandleChallengeRegister(c *gin.Context) {
	user := cc.userFromContext(c)
	if user == nil {
		cc.log.Error("user not found")
		cc.InternalErrorResponse(c, "user not found")
		return
	}
	form := challenge.Meta{}
	if err := c.ShouldBindJSON(&form); err != nil {
		cc.BadRequestResponse(c, err.Error())
		return
	}
	challengeInfo, err := cc.DB.GetChallengeData(form.ChallengeID)
	if challengeInfo.ParticipantCount >= challengeInfo.MaxParticipants {
		cc.CustomFailedResponse(c, "max participant are reached", http.StatusNotAcceptable)
		return
	}
	if err != nil {
		cc.log.Error("error in GetChallengeData for challenge ", form.ChallengeID, ": ", err)
		cc.InternalErrorResponse(c, "fail to GetChallengeData")
		return
	}
	if challengeInfo.EntryOnce {
		challengeRecords, err := cc.DB.GetChallengeRecordsByUser(form.ChallengeID, user.Sub)
		if err != nil {
			cc.log.Errorf("error in GetChallengeRecordsByUser(%s, %s): %v", form.ChallengeID, user.Sub, err)
			cc.InternalErrorResponse(c, "fail to get existing challenge records")
			return
		}
		if len(challengeRecords) > 0 {
			cc.BadRequestResponse(c, "can only join this challenge once")
			return
		}
	}
	game, err := cc.DB.GetSpecificGame(user.Sub, strings.ToUpper(challengeInfo.ChallengeType))
	if err != nil {
		cc.BadRequestResponse(c, "missing game account")
		return
	}
	playerId := ""
	for _, platform := range game.Platforms {
		if platform.PlatformName == game.SelectedPlatform {
			playerId = platform.ID
			break
		}
	}
	if playerId == "" {
		cc.BadRequestResponse(c, "missing player ID")
		return
	}

	form.Platform = game.SelectedPlatform
	form.PlayerID = playerId

	if challengeInfo.EntryFee != "0" { // check user balance if there is entry fee
		balance, err := cc.bsc.GetSPWNBalance(user.PublicAddress, challengeInfo.FeeType)
		if err != nil {
			cc.log.Error("cannot get SPWN balance")
			cc.InternalErrorResponse(c, "cannot get SPWN balance")
			return
		}
		truncatedBalance, _ := balance.Int64()

		entryFeeInt, err := strconv.ParseInt(challengeInfo.EntryFee, 10, 64)
		if err != nil {
			cc.log.Error("error in converting entryFee to int64")
			cc.InternalErrorResponse(c, "error in converting entryFee to int64")
			return
		}
		if truncatedBalance < entryFeeInt {
			cc.BadRequestResponse(c, "insufficient balance")
			return
		}
	}
	err = cc.tsvc.ExecuteStateMachine(form.ChallengeID, user.Username, "Register", &form)
	if err != nil {
		cc.log.Error("error in Register challenge ", form.ChallengeID, ": ", err)
		cc.InternalErrorResponse(c, "error in Register challenge: "+err.Error())
		return
	}

	cc.SuccessResponse(c, "challenge registered successfully", nil)
}

func (cc *ChallengeController) HandleSponsoredChallengeList(c *gin.Context) {
	sponsoredChallenges, err := cc.DB.GetSponsoredChallenges()
	if err != nil {
		cc.log.Error("error in ListSponsoredChallenges from db: ", err)
		cc.DBErrorResponse(c, "error in list sponsored challenge data from db")
		return
	}

	t := gin.H{
		"sponsoredChallenges": sponsoredChallenges,
	}
	cc.SuccessResponse(c, "", t)
}

func (cc *ChallengeController) HandleFeaturedChallengeFetch(c *gin.Context) {
	title := c.Param("title")
	featuredObject, err := cc.DB.GetFeaturedObject(title)
	if err != nil {
		cc.log.Error("there is no such featured challenge in db")
		cc.DBErrorResponse(c, "there is no such featured challenge in db")
		return
	}
	if featuredObject.ObjectType != "Challenge" {
		cc.log.Error("the object is not a challenge")
		cc.InternalErrorResponse(c, "the object is not a challenge")
		return
	}

	challenge, err := cc.DB.GetChallengeForResponse(featuredObject.ObjectID)
	if err != nil {
		if err.Error() == "record not found" {
			cc.BadRequestResponse(c, "challenge does not exist")
			return
		}
		cc.log.Error("error getting challenge: ", err)
		cc.DBErrorResponse(c, "error getting challenge")
		return
	}

	leaders, err := cc.DB.GetChallengeLeaderboard(featuredObject.ObjectID, PER_PAGE)
	if err != nil {
		cc.log.Error("error in GetChallengeLeaderboard: ", err)
		cc.DBErrorResponse(c, "error in GetChallengeLeaderboard")
		return
	}

	t := gin.H{
		"challengeDetails": challenge,
		"leaderboard":      leaders,
		"teamBoard":        cc.DB.GroupByTeam(leaders),
	}
	cc.SuccessResponse(c, "", t)
}
func (cc *ChallengeController) HandleFeaturedChallengeList(c *gin.Context) {
	featuredChallenges, err := cc.DB.GetFeaturedChallenges()
	if err != nil {
		cc.log.Error("error in get featured challenges from db", err)
		cc.DBErrorResponse(c, err.Error())
		return
	}

	t := gin.H{
		"featuredChallenges": featuredChallenges,
	}
	cc.SuccessResponse(c, "", t)
}

func (cc *ChallengeController) HandleChallengeEnum(c *gin.Context) {
	challengeTypeStruct, err := cc.DB.GetChallengeTypes()
	if err != nil {
		cc.log.Error("error get challenge type from db", err)
		cc.DBErrorResponse(c, err.Error())
		return
	}
	if len(challengeTypeStruct) == 0 {
		cc.log.Error("there are no challenge types in db")
		cc.DBErrorResponse(c, "there are no challenge types in db")
		return
	}

	challengeTypes := []string{}
	for i := range challengeTypeStruct {
		challengeType := challengeTypeStruct[i].ChallengeType
		if !contains(challengeTypes, challengeType) {
			challengeTypes = append(challengeTypes, challengeType)
		}
	}

	t := gin.H{
		"challengeTypes":    challengeTypes,
		"challengeSubtypes": challengeTypeStruct,
	}
	cc.SuccessResponse(c, "", t)
}

func (cc *ChallengeController) HandleAdminChallengeReportScore(c *gin.Context) {
	challengeRecordsToReportScore, err := cc.DB.GetChallengeRecordsToReport()
	if err != nil {
		cc.log.Error(err)
		cc.InternalErrorResponse(c, "error in GetChallengeRecordsPastDeadline: "+err.Error())
		return
	}

	counter := 0
	counterFail := 0
	for _, record := range challengeRecordsToReportScore {
		cc.log.Debug("Reporting score for user: ", record.UserId, " in challenge: ", record.ChallengeId)
		err := cc.tsvc.ReportScore(record)
		if err != nil {
			cc.log.Error(fmt.Printf("error in reporting score for user %s in challenge %s: %v", record.UserId, record.ChallengeId, err))
			counterFail++
			continue
		} else {
			counter++
		}
	}

	t := gin.H{
		"numberOfScoresReported": counter,
		"numberOfScoresFailed":   counterFail,
	}
	cc.SuccessResponse(c, "", t)
}

func (cc *ChallengeController) HandleChallengeFetch(c *gin.Context) {
	challengeID := c.Param("challengeId")
	if challengeID == "" {
		cc.BadRequestResponse(c, "missing input challenge id")
		return
	}

	challenge, err := cc.DB.GetChallengeForResponse(challengeID)
	if err != nil {
		if err.Error() == "record not found" {
			cc.BadRequestResponse(c, "challenge does not exist")
			return
		}
		cc.log.Error("error getting challenge: ", err)
		cc.DBErrorResponse(c, "error getting challenge")
		return
	}

	leaders, err := cc.DB.GetChallengeLeaderboard(challengeID, PER_PAGE)
	if err != nil {
		cc.log.Error("error in GetChallengeLeaderboard: ", err)
		cc.DBErrorResponse(c, "error in GetChallengeLeaderboard")
		return
	}

	responseData := gin.H{
		"challengeDetails": challenge,
		"leaderboard":      leaders,
		"teamBoard":        cc.DB.GroupByTeam(leaders),
	}

	cc.SuccessResponse(c, "", responseData)
}

func (cc *ChallengeController) HandleChallengeFetchUserStatus(c *gin.Context) {
	user := cc.userFromContext(c)
	if user == nil {
		cc.log.Error("user not found")
		cc.InternalErrorResponse(c, "user not found")
		return
	}
	challengeID := c.Param("challengeId")
	if challengeID == "" {
		cc.BadRequestResponse(c, "missing input challenge id")
		return
	}
	challenge, err := cc.DB.GetChallengeData(challengeID)
	if err != nil {
		if err.Error() == "record not found" {
			cc.BadRequestResponse(c, "challenge does not exist")
			return
		}
		cc.log.Error("error getting challenge: ", err)
		cc.DBErrorResponse(c, "error getting challenge")
		return
	}

	userAction := "register"
	if challenge.EntryOnce {
		records, err := cc.DB.GetChallengeRecordsByUser(challenge.ChallengeID.String(), user.Sub)
		if err != nil {
			userAction = "none"
		}
		if len(records) > 0 {
			userAction = "none"
		}
	}
	if time.Now().After(challenge.CutoffDate) {
		userAction = "none"
	}
	responseData := gin.H{
		"userAction": userAction,
	}

	cc.SuccessResponse(c, "", responseData)
}

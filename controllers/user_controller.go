/*

 */

package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bitspawngg/bitspawn-api/enum"
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services"
	"github.com/bitspawngg/bitspawn-api/services/cognito"
	"github.com/bitspawngg/bitspawn-api/services/hdkey"
	"github.com/bitspawngg/bitspawn-api/services/poa"
	"github.com/bitspawngg/bitspawn-api/services/queue"
	"github.com/bitspawngg/bitspawn-api/services/s3"
	userdata "github.com/bitspawngg/bitspawn-api/services/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/datatypes"
)

type UserController struct {
	BaseController
	bsc            *poa.BitspawnPoaClient
	AuthStore      *cognito.Auth
	S3UploadClient *s3.S3UploadClient
	SQSClient      *queue.SQSService
	usvc           *userdata.UserService
}

func NewUserController(bc *BaseController, bsc *poa.BitspawnPoaClient, authCognito *cognito.Auth, s3UploadClient *s3.S3UploadClient, sqsClient *queue.SQSService, usvc *userdata.UserService) *UserController {
	return &UserController{
		BaseController{
			Name: "tournament",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		},
		bsc,
		authCognito,
		s3UploadClient,
		sqsClient,
		usvc,
	}
}

type FormUserSignup struct {
	Username string `json:"username"`
	Password string `json:"password"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *UserController) asyncSendSPWNTokenAndETHFunction(token, publicAddress string) {

	err := u.bsc.GiftSpwn(token)
	if err != nil {
		err = u.bsc.GiftSpwn(token)
		if err != nil {
			u.log.Warn("Fail to send gift SPWN token. ", err)
		} else {
			u.log.Info("Successfully gift SPWN token to ", publicAddress)
			// log.Printf("Successfully gift SPWN token to %s", publicAddress)
		}
	} else {
		u.log.Info("Successfully gift SPWN token to ", publicAddress)
		// log.Printf("Successfully gift SPWN token to %s", publicAddress)
	}

	_, err = u.bsc.GiftEth(publicAddress)
	if err != nil {
		u.log.Error("fail to send gift eth")
		return
	}
	u.log.Info("Successfully gift eth to ", publicAddress)
	// log.Printf("Successfully gift eth to %s", publicAddress)
}

func (u *UserController) HandleSignup(c *gin.Context) {
	authCofig := *u.AuthStore
	cognitoRegion := authCofig.CognitoRegion()
	cognitoUserPoolID := authCofig.CognitoUserPoolID()
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		u.log.Error("token not found")
		u.AuthErrorResponse(c, "token not found")
		return
	}

	tokens, err := u.AuthStore.ParseJWT(token, cognitoRegion, cognitoUserPoolID)
	if err != nil || !tokens.Valid {
		// jwt is not valid
		u.log.Error("token is not valid")
		u.AuthErrorResponse(c, "token is not valid")
		return
	}

	cognitoUserName := tokens.Claims.(jwt.MapClaims)["cognito:username"]
	subfield := tokens.Claims.(jwt.MapClaims)["sub"]
	nickName, ok := tokens.Claims.(jwt.MapClaims)["custom:init_display_name"]
	if !ok {
		nickName = strings.Split(cognitoUserName.(string), "@")[0]
	}
	username := cognitoUserName.(string)
	subStr := subfield.(string)
	displayNameTemp := nickName.(string)

	displayNameCount, err := u.DB.CountDisplayNameUsage(displayNameTemp)
	if err != nil {
		u.log.Error("error counting displayName", err)
		u.DBErrorResponse(c, "error counting displayName")
		return
	}

	if displayNameCount != 0 {
		rand.Seed(time.Now().UnixNano())
		chars := []rune("0123456789")
		length := 4
		var randNum strings.Builder
		for i := 0; i < length; i++ {
			randNum.WriteRune(chars[rand.Intn(len(chars))])
		}
		randStr := randNum.String()
		displayNameTemp = displayNameTemp + "#" + randStr
	}

	publicAddress, err := hdkey.GeneratePublicAddressFromUUID(subStr)
	if err != nil {
		// jwt is not valid
		u.log.Error("fail to create publicAddress")
		u.InternalErrorResponse(c, "fail to create publicAddress")
		return
	}

	var FavouriteList json.RawMessage
	// favouriteGameList :=

	favouriteList := make(map[string]interface{})
	favouriteList["favouriteGame"] = []string{}

	favourite, err := json.Marshal(favouriteList)
	if err != nil {
		u.BadRequestResponse(c, "Wrong Favourite Game Input")
		return
	}

	_ = FavouriteList.UnmarshalJSON(favourite)

	userAccount := models.UserAccount{Username: username, DisplayName: displayNameTemp, PublicAddress: publicAddress, Sub: subStr, Favourite: datatypes.JSON(FavouriteList)}
	err = u.DB.UserSignup(username, displayNameTemp, publicAddress, subStr, FavouriteList)
	if err != nil {
		u.log.Error("error in UserSignup: ", err)
		u.DBErrorResponse(c, err.Error())
		return
	}

	go u.asyncSendSPWNTokenAndETHFunction(token, publicAddress)
	// log.Printf("async message done")

	u.SuccessResponse(c, "registration success", userAccount)
}

func (u *UserController) HandleGetAddress(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	publicAddress := user.PublicAddress

	u.SuccessResponse(c, "", gin.H{"address": publicAddress})
}

func (u *UserController) HandleGetBalance(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	creditBalance, err := u.bsc.GetSPWNBalance(user.PublicAddress, enum.Credit.ToString())
	if err != nil {
		u.log.Error("cannot get user credit balance")
		u.InternalErrorResponse(c, "cannot get user credit balance")
		return
	}
	truncatedCredit, _ := creditBalance.Int64()

	spwnBalance, err := u.bsc.GetSPWNBalance(user.PublicAddress, enum.Spwn.ToString())
	if err != nil {
		u.log.Error("cannot get user spwn balance")
		u.InternalErrorResponse(c, "cannot get user spwn balance")
		return
	}
	truncatedSPWN, _ := spwnBalance.Int64()

	u.SuccessResponse(c, "", gin.H{"balance": truncatedCredit,
		enum.Credit.ToString(): truncatedCredit, enum.Spwn.ToString(): truncatedSPWN})
}

func (u *UserController) HandleGetGameAccounts(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	gameAccounts, err := u.DB.GetUserGameAccount(user.Sub)
	if err != nil {
		if err.Error() == "record not found" {
			u.BadRequestResponse(c, "player has not connected game accounts")
			return
		}
		u.log.Error("cannot get game accounts")
		u.InternalErrorResponse(c, "cannot get game accounts")
		return
	}

	u.SuccessResponse(c, "", gin.H{"gameAccounts": gameAccounts})
}

func (u *UserController) HandleNotification(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	notes, err := u.DB.ReadNotification(user.Sub)
	if err != nil {
		u.log.Error("cannot read user notification: ", err)
		u.InternalErrorResponse(c, "cannot read user notification: "+err.Error())
		return
	}

	u.SuccessResponse(c, "", gin.H{"notification": notes})
}

type ResponseExistingUsers struct {
	UserNumber int64 `json:"userNumber"`
}

func (u *UserController) HandleFetchExistingUsers(c *gin.Context) {

	userNumbers, err := u.DB.FetchExistUserNumber()
	if err != nil {
		u.log.Error("error getting exist user number", err)
		u.DBErrorResponse(c, "error getting exist user number")
		return
	}
	u.SuccessResponse(c, "", ResponseExistingUsers{
		userNumbers,
	})
}

type ResponseUserProfile struct {
	UserProfileData *models.UserAccount `json:"userProfileData"`
}

func (u *UserController) HandleUserProfileGet(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}
	u.SuccessResponse(c, "", ResponseUserProfile{
		user,
	})
}

func (u *UserController) HandleOtherUserProfileGet(c *gin.Context) {
	displayName := c.Param("displayName")
	userInfo, err := u.DB.GetUserProfileByDisplayName(displayName)
	if err != nil {
		u.log.Error("error getting user profile", err)
		u.DBErrorResponse(c, "error getting user profile")
		return
	}
	//remove sensitive data from response
	userInfo.Username = ""
	userInfo.PhoneNumber = ""
	u.SuccessResponse(c, "", ResponseUserProfile{
		userInfo,
	})
}

type FormUserProfileUpdate struct {
	DisplayName string `json:"displayName"`
	Country     string `json:"country"`
	Timezone    string `json:"timezone"`
	PhoneNumber string `json:"phoneNumber"`
}

func (u *UserController) HandleUserProfilePut(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormUserProfileUpdate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	if user.DisplayName != form.DisplayName {
		displayNameCount, err := u.DB.CountDisplayNameUsage(form.DisplayName)
		if err != nil {
			u.log.Error("error counting displayName", err)
			u.DBErrorResponse(c, "error counting displayName")
			return
		}

		if displayNameCount != 0 {
			u.InternalErrorResponse(c, "displayName has been used")
			return
		}
	}

	users, err := u.DB.UpdateUserProfile(user.Sub, form.DisplayName)
	if err != nil {
		u.log.Error("error updating user profile", err)
		u.InternalErrorResponse(c, "updating user profile fails")
		return
	}

	u.SuccessResponse(c, "user profile update successfully", users)
}

func (u *UserController) HandleUserPhoneNumberReset(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	if user.Enabled2FA {
		u.BadRequestResponse(c, "phone number already attached to 2FA")
		return
	}

	form := struct {
		PhoneNumber string `json:"phoneNumber"`
	}{}
	if err := c.ShouldBindJSON(&form); err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	err := validatePhoneNumber(form.PhoneNumber)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
	}

	err = u.DB.UpdateUserPhoneNumber(user.Sub, form.PhoneNumber)
	if err != nil {
		u.log.Error("error updating user phone number", err)
		u.InternalErrorResponse(c, "updating user phone number fails")
		return
	}

	u.SuccessResponse(c, "success", nil)
}

func validatePhoneNumber(phoneNumber string) error {
	if strings.HasPrefix(phoneNumber, "+") {
		return fmt.Errorf("phone number must start with +")
	}
	phoneInt, err := strconv.Atoi(phoneNumber)
	if err != nil {
		return fmt.Errorf("all digits must be 0-9")
	}
	if phoneInt < 10000000 || phoneInt > 999999999999999 {
		return fmt.Errorf("phone number must be between 8 and 15 digits")
	}
	return nil
}

type FormUserFavouriteGameUpdate struct {
	FavouriteCategory string   `json:"category"`
	FavouriteList     []string `json:"favouriteList"`
}

func (u *UserController) HandleUserFavouriteGameUpdate(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormUserFavouriteGameUpdate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	var FavouriteList json.RawMessage

	// if form.FavouriteCategory == "Game" {
	favouriteGameList := form.FavouriteList
	gameTypes, err := u.DB.GetGameType()
	if err != nil {
		u.log.Error("error get game type from db", err)
		u.DBErrorResponse(c, err.Error())
		return
	}
	if len(favouriteGameList) != 0 {
		gameTypeList := []string{}
		for _, gameType := range gameTypes {
			gameTypeList = append(gameTypeList, gameType.GameType)
		}

		for _, favouriteGame := range favouriteGameList {
			if !stringInSlice(favouriteGame, gameTypeList) {
				u.BadRequestResponse(c, "Wrong Favourite Game Input")
				return
			}
		}
	}

	favouriteList := make(map[string]interface{})
	favouriteList["favouriteGame"] = favouriteGameList

	favourite, err := json.Marshal(favouriteList)
	if err != nil {
		u.BadRequestResponse(c, "Wrong Favourite Game Input")
		return
	}

	_ = FavouriteList.UnmarshalJSON(favourite)

	err = u.DB.UpdateUserFavouriteData(user.Sub, FavouriteList)
	if err != nil {
		u.log.Error("error updating user favourite games", err)
		u.InternalErrorResponse(c, "updating user favourite games fails")
		return
	}

	u.SuccessResponse(c, "user favourite games update successfully", nil)
}

type PlayRecordOutPut struct {
	GameType       string    `json:"gameType"`
	TournamentId   string    `json:"tournamentId"`
	TournamentName string    `json:"tournamentName"`
	Status         string    `json:"status"`
	EntryFee       string    `json:"entryFee"`
	PrizeEarned    int64     `json:"prizeEarned"`
	FinishDate     time.Time `json:"finishDate"`
}
type ResponsePlayHistory struct {
	PlayHistoryData []PlayRecordOutPut `json:"playHistoryData"`
	Pages           int64              `json:"pages"`
}

func (u *UserController) HandlePlayHistoryGet(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}
	perPageString := c.DefaultQuery("perPage", strconv.Itoa(PER_PAGE))
	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	numberOfCompletedPlayRecords, err := u.DB.CountFinishRecords(user.Username)
	if err != nil {
		u.log.Error("error getting CountFinishRecords: ", err)
		u.DBErrorResponse(c, "error getting CountFinishRecords: "+err.Error())
		return
	}
	pageTournaments := math.Ceil(float64(numberOfCompletedPlayRecords) / float64(perPage))

	playRecordsOutput, err := u.DB.GetFinishRecords(user.Username, page, perPage)
	if err != nil {
		u.log.Error("error getting GetFinishRecords: ", err)
		u.DBErrorResponse(c, "error getting GetFinishRecords: "+err.Error())
		return
	}

	numberOfCompletedChallengeRecords, err := u.DB.CountFinishedChallengeRecords(user.Username, perPage)
	if err != nil {
		u.log.Error("error getting CountMyFinishedChallengeRecords: ", err)
		u.DBErrorResponse(c, "error getting CountMyFinishedChallengeRecords: "+err.Error())
		return
	}
	pageChallenges := math.Ceil(float64(numberOfCompletedChallengeRecords) / float64(perPage))

	challengeRecords, err := u.DB.GetFinishedChallengeRecords(user.Username, page, perPage)
	if err != nil {
		u.log.Error("error getting GetFinishChallengeRecords: ", err)
		u.DBErrorResponse(c, "error getting GetFinishChallengeRecords: "+err.Error())
		return
	}

	t := gin.H{
		"playHistoryData":      playRecordsOutput,
		"pageTournaments":      pageTournaments,
		"challengeHistoryData": challengeRecords,
		"pageChallenges":       pageChallenges,
	}
	u.SuccessResponse(c, "", t)
}

func (u *UserController) HandlePlayHistoryByGameTypeGet(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	gameType := c.Param("gameType")
	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	pageWhole, err := u.DB.CountMyPlayRecordPages(user.Username, PER_PAGE)
	if err != nil {
		u.log.Error("error getting user play history pages", err)
		u.DBErrorResponse(c, "error getting user play history pages")
		return
	}

	playRecords, err := u.DB.GetPlayRecordByGameType(user.Username, gameType, page, PER_PAGE)
	if err != nil {
		u.log.Error("error getting user play history", err)
		u.DBErrorResponse(c, "error getting user play history")
		return
	}
	output, err := u.FormatRecordsForOutput(playRecords)
	if err != nil {
		u.log.Error(err)
		u.DBErrorResponse(c, err.Error())
		return
	}

	u.SuccessResponse(c, "", ResponsePlayHistory{
		output,
		pageWhole,
	})
}

func (u *UserController) FormatRecordsForOutput(playRecords []models.PlayRecord) ([]PlayRecordOutPut, error) {
	recordsLen := len(playRecords)
	output := make([]PlayRecordOutPut, recordsLen)
	for i := range playRecords {
		output[i].GameType = playRecords[i].GameType
		output[i].PrizeEarned = playRecords[i].PrizeEarned
		output[i].TournamentId = playRecords[i].TournamentId
		tournamentID := playRecords[i].TournamentId
		tournament, err := u.DB.GetTournamentData(tournamentID)
		if err != nil {
			return nil, fmt.Errorf("error getting tournament: %v", err)
		}
		output[i].TournamentName = tournament.TournamentName
		output[i].EntryFee = tournament.EntryFee
		output[i].Status = tournament.Status
		output[i].FinishDate = tournament.TournamentDate
	}
	return output, nil
}

type UserStats struct {
	TournamentsParticipated int64     `json:"tournamentsParticipated"`
	TournamentsWon          int64     `json:"tournamentsWon"`
	TournamentLost          int64     `json:"tournamentLost"`
	WinRate                 string    `json:"winRate"`
	RecentMatch             int       `json:"recentMatch"`
	TotalPrizeEarned        string    `json:"totalPrizeEarned"`
	MostPlayedGame          string    `json:"mostPlayedGame"`
	Displayname             string    `json:"displayname"`
	UserAvatar              string    `json:"userAvatar"`
	UserStatus              bool      `json:"userStatus"`
	RegisterTime            time.Time `json:"registerTime"`
}

type ResponseUserStats struct {
	UserPlayScore UserStats
}

func (u *UserController) HandleUserStatsGet(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}
	userStatus := false

	tournamentsParticipated, err := u.DB.CountFinishRecords(user.Username)
	if err != nil {
		u.log.Error("error counting user's finish records", err)
		u.DBErrorResponse(c, "error counting user's finish records")
		return
	}
	tournamentsWon, err := u.DB.CountWinRecords(user.Username)
	if err != nil {
		u.log.Error("error counting user's win records", err)
		u.DBErrorResponse(c, "error counting user's win records")
		return
	}
	gamePlayMost, err := u.DB.CalculateGamePlay(user.Username)
	if err != nil {
		u.log.Error("error getting user most play games", err)
		u.DBErrorResponse(c, "error getting user most play games")
		return
	}
	totalEarning := u.DB.AccumulateEarning(user.Username)
	winningRate := 0.0
	var tournamentLost int64 = 0

	if tournamentsParticipated != 0 {
		winningRate = float64(tournamentsWon) / float64(tournamentsParticipated) * 100
		tournamentLost = int64(tournamentsParticipated) - tournamentsWon
	}
	winRate := fmt.Sprintf("%0.2f %%", winningRate)
	totalPrizeEarned := fmt.Sprintf("%d", totalEarning)
	mostPlayedGame := gamePlayMost.GameType
	recentMatch, err := u.FormatMatchScore(user.Username)
	if err != nil {
		u.log.Error(err)
		u.DBErrorResponse(c, err.Error())
		return
	}

	displayname := user.DisplayName
	userAvatar := user.AvatarUrl
	time := time.Now().UTC()
	diff := time.Sub(user.OnlineTime).Seconds()
	if diff < 600 {
		userStatus = true
	}
	registerTime := user.CreatedAt
	userStats := UserStats{tournamentsParticipated, tournamentsWon, tournamentLost, winRate, recentMatch, totalPrizeEarned, mostPlayedGame, displayname, userAvatar, userStatus, registerTime}
	u.SuccessResponse(c, "", ResponseUserStats{
		userStats,
	})
}

func (u *UserController) HandleUserStatsByGameTypeGet(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}
	userStatus := false

	gameType := c.Param("gameType")
	tournamentsParticipated, err := u.DB.CountFinishRecordsByGameType(user.Username, gameType)
	if err != nil {
		u.log.Error("error getting user play history", err)
		u.DBErrorResponse(c, "error getting user play history")
		return
	}
	tournamentsWon, err := u.DB.CountWinRecordsByGameType(user.Username, gameType)
	if err != nil {
		u.log.Error("error getting user winning play history", err)
		u.DBErrorResponse(c, "error getting user winning play history")
		return
	}
	totalEarning := u.DB.AccumulateEarningByGameType(user.Username, gameType)
	winningRate := 0.0
	var tournamentLost int64 = 0

	if tournamentsParticipated != 0 {
		winningRate = float64(tournamentsWon) / float64(tournamentsParticipated) * 100
		tournamentLost = int64(tournamentsParticipated) - tournamentsWon
	}
	winRate := fmt.Sprintf("%0.2f %%", winningRate)
	totalPrizeEarned := fmt.Sprintf("%d", totalEarning)
	mostPlayedGame := ""
	recentMatch, err := u.FormatMatchScore(user.Username)
	if err != nil {
		u.log.Error(err)
		u.DBErrorResponse(c, err.Error())
		return
	}

	displayname := user.DisplayName
	userAvatar := user.AvatarUrl
	time := time.Now().UTC()
	diff := time.Sub(user.OnlineTime).Seconds()
	if diff < 600 {
		userStatus = true
	}
	registerTime := user.CreatedAt
	userStats := UserStats{tournamentsParticipated, tournamentsWon, tournamentLost, winRate, recentMatch, totalPrizeEarned, mostPlayedGame, displayname, userAvatar, userStatus, registerTime}
	u.SuccessResponse(c, "", ResponseUserStats{
		userStats,
	})
}

func (u *UserController) HandleOtherUserStatsGet(c *gin.Context) {
	displayname := c.Param("displayname")

	userStatus := false

	users, err := u.DB.GetUserProfileByDisplayName(displayname)
	if err != nil {
		u.log.Error("error getting user profile", err)
		u.DBErrorResponse(c, "error getting user profile")
		return
	}

	userId := users.Username

	tournamentsParticipated, err := u.DB.CountFinishRecords(userId)
	if err != nil {
		u.log.Error("error counting user's finish records", err)
		u.DBErrorResponse(c, "error counting user's finish records")
		return
	}
	tournamentsWon, err := u.DB.CountWinRecords(userId)
	if err != nil {
		u.log.Error("error counting user's win records", err)
		u.DBErrorResponse(c, "error counting user's win records")
		return
	}
	gamePlayMost, err := u.DB.CalculateGamePlay(userId)
	if err != nil {
		u.log.Error("error getting user most play games", err)
		u.DBErrorResponse(c, "error getting user most play games")
		return
	}
	totalEarning := u.DB.AccumulateEarning(userId)
	winningRate := 0.0
	var tournamentLost int64 = 0

	if tournamentsParticipated != 0 {
		winningRate = float64(tournamentsWon) / float64(tournamentsParticipated) * 100
		tournamentLost = int64(tournamentsParticipated) - tournamentsWon
	}
	winRate := fmt.Sprintf("%0.2f %%", winningRate)
	totalPrizeEarned := fmt.Sprintf("%d", totalEarning)
	mostPlayedGame := gamePlayMost.GameType
	recentMatch, err := u.FormatMatchScore(users.Username)
	if err != nil {
		u.log.Error(err)
		u.DBErrorResponse(c, err.Error())
		return
	}

	userAvatar := users.AvatarUrl
	time := time.Now().UTC()
	diff := time.Sub(users.OnlineTime).Seconds()
	if diff < 600 {
		userStatus = true
	}
	registerTime := users.CreatedAt
	userStats := UserStats{tournamentsParticipated, tournamentsWon, tournamentLost, winRate, recentMatch, totalPrizeEarned, mostPlayedGame, displayname, userAvatar, userStatus, registerTime}
	u.SuccessResponse(c, "", ResponseUserStats{
		userStats,
	})
}

func (u *UserController) FormatMatchScore(username string) (int, error) {
	var winningNumber int = 0
	exponentNumber := 0.5

	playerName := username
	_, err := u.DB.GetUserProfile(playerName)
	var teamId []string
	var userExist bool
	if err == nil {
		userExist = true
	} else if err.Error() == "record not found" {
		return 0, fmt.Errorf("error getting user information: %v", err)
	} else if err != nil {
		return 0, fmt.Errorf("error getting user information: %v", err)
	}

	if userExist {
		teamId, err = u.DB.FetchTeamIdsByPlayerName(playerName)
		if err != nil {
			return 0, fmt.Errorf("error getting teams' ID: %v", err)
		}

		//Fetch Recent Matches By TeamId And Staus
		teamMatches, err := u.DB.FetchRecentMatchesByTeamId(teamId, LAST_MATCH_NUMBER)
		if err != nil {
			return 0, fmt.Errorf("error getting matches: %v", err)
		}
		for _, teamMatch := range teamMatches {
			exponentNumber = exponentNumber * 2
			// log.Printf("match info: %s", teamMatch)
			if teamMatch.Status == "Finished" {
				if teamMatch.Result == 1 {
					if contains(teamId, teamMatch.TeamOne) {
						winningNumber = winningNumber + int(exponentNumber)
					}
				} else if teamMatch.Result == 2 {
					if contains(teamId, teamMatch.TeamTwo) {
						winningNumber = winningNumber + int(exponentNumber)
					}
				} else if teamMatch.Result == -1 {
					if teamMatch.TeamOne > teamMatch.TeamTwo {
						if contains(teamId, teamMatch.TeamOne) {
							winningNumber = winningNumber + int(exponentNumber)
						}
					} else {
						if contains(teamId, teamMatch.TeamTwo) {
							winningNumber = winningNumber + int(exponentNumber)
						}
					}
				}
			}
		}
	}

	return winningNumber, nil
}

type DepositRecordOutPut struct {
	TxHash     string    `json:"txHash"`
	TxSender   string    `json:"txSender"`
	TxReceiver string    `json:"txReceiver"`
	PoaAddress string    `json:"poaAddress"`
	CoinType   string    `json:"coinType"`
	CoinAmount float64   `json:"coinAmount"`
	SpwnAmount string    `json:"spwnAmount"`
	MintDate   time.Time `json:"mintDate"`
	Action     string    `json:"action"`
	Fee        string    `json:"fee"`
	Remark     string    `json:"remark"`
}
type ResponseDepositRecord struct {
	DepositRecordData []DepositRecordOutPut `json:"depositRecordData"`
	Pages             int64                 `json:"pages"`
}

func (u *UserController) HandleDepositRecordGet(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}
	pageWhole, err := u.DB.CountMyDepositRecordPages(user.PublicAddress, PER_PAGE)
	if err != nil {
		u.log.Error("error getting user deposit record pages", err)
		u.DBErrorResponse(c, "error getting user deposit record pages")
		return
	}

	depositRecord, err := u.DB.GetUserDepositRecord(user.PublicAddress, page, PER_PAGE)
	if err != nil {
		u.log.Error("error getting user deposit record", err)
		u.DBErrorResponse(c, "error getting user deposit record")
		return
	}

	recordsLen := len(depositRecord)

	depositRecords := make([]DepositRecordOutPut, recordsLen)
	for i, dr := range depositRecord {
		spwnAmount := ConvertWeiToEth(dr.SpwnAmount)
		fee := ConvertWeiToEth(dr.Fee)
		switch dr.CoinType {
		case "USD":
			coinAmountFloat64, err := strconv.ParseFloat(dr.CoinAmount, 64)
			if err != nil {
				u.log.Error("USD Amount is wrong")
				u.InternalErrorResponse(c, "USD Amount is wrong")
				return
			}
			depositRecords[i].CoinAmount = coinAmountFloat64 / float64(100)
		case "ERC20SPWN":
			coinAmount := ConvertWeiToEth(dr.CoinAmount)
			coinAmountFloat64, _ := coinAmount.Float64()
			depositRecords[i].CoinAmount = coinAmountFloat64
		default:
			coinAmount := ConvertWeiToEth(dr.CoinAmount)
			coinAmountFloat64, _ := coinAmount.Float64()
			depositRecords[i].CoinAmount = coinAmountFloat64
		}
		depositRecords[i].TxHash = dr.TxHash
		depositRecords[i].TxSender = dr.TxSender
		depositRecords[i].TxReceiver = dr.TxReceiver
		depositRecords[i].PoaAddress = dr.PoaAddress
		depositRecords[i].CoinType = dr.CoinType
		depositRecords[i].SpwnAmount = spwnAmount.String()
		depositRecords[i].MintDate = dr.MintDate
		depositRecords[i].Action = dr.Action
		depositRecords[i].Fee = fee.String()
		depositRecords[i].Remark = dr.Remark
	}

	u.SuccessResponse(c, "", ResponseDepositRecord{
		depositRecords,
		pageWhole,
	})
}

func (u *UserController) HandleMarketplaceRecordGet(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		BadRequestResponseV2(c, "empty page query parameter", err.Error())
		return
	}
	pageWhole, err := u.usvc.CountMyMarketplaceRecordPages(user.PublicAddress, PER_PAGE)
	if err != nil {
		u.log.Error("error getting user marketplace record pages", err)
		InternalErrorResponseV2(c, "error getting user deposit record pages", err.Error())
		return
	}

	marketplaceRecords, err := u.usvc.GetMarketplaceOrderRecords(user.PublicAddress, page, PER_PAGE)
	if err != nil {
		u.log.Error("error getting user marketplace record", err)
		InternalErrorResponseV2(c, "error getting user marketplace record", err.Error())
		return
	}
	SuccessResponseV2(c, struct {
		MarketplaceRecords []*models.MarketplaceOrderRecord `json:"marketplaceRecordData"`
		Pages              int64                            `json:"pages"`
	}{
		MarketplaceRecords: marketplaceRecords,
		Pages:              pageWhole,
	})
}

type FormFindFriend struct {
	DisplayName string `json:"displayName"`
}

func (u *UserController) HandleFindFriend(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormFindFriend{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	if form.DisplayName == "" {
		u.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	friendName := "%" + form.DisplayName + "%"

	friendSearchList, err := u.DB.ListFriendByDisplayName(friendName)
	if err != nil {
		u.log.Error("error getting user friend invitation data from db: ", err)
		u.DBErrorResponse(c, "error getting user friend invitation data from db")
		return
	}

	responseData := gin.H{
		"FriendSearchList": friendSearchList,
	}

	u.SuccessResponse(c, "", responseData)
}

type FormFriendInvitationCreate struct {
	DisplayName string `json:"displayName"`
}

func (u *UserController) HandleFriendInvitationCreate(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormFriendInvitationCreate{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	if form.DisplayName == "" {
		u.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	if form.DisplayName == user.DisplayName {
		u.BadRequestResponse(c, "cannot send friend request to yourself")
		return
	}

	inviteUserInfo, err := u.DB.GetUserProfileByDisplayName(form.DisplayName)
	if err != nil {
		u.log.Error("error getting invitation user profile", err)
		u.DBErrorResponse(c, "error getting invitation user profile")
		return
	}

	friendInvitations, err := u.DB.FetchFriendInviteByUserIds(user.Sub, inviteUserInfo.Sub)
	if err == nil {
		// do nothing
	} else if err.Error() == "record not found" {
		// do nothing
	} else if err != nil {
		u.log.Error("error getting user friend invitation data from db: ", err)
		u.DBErrorResponse(c, "error getting user friend invitation data from db")
		return
	}

	if len(friendInvitations) != 0 {
		u.BadRequestResponse(c, "The Friend Request has been created!")
		return
	}

	friendRequests, err := u.DB.FetchFriendInviteByUserIds(inviteUserInfo.Sub, user.Sub)
	if err == nil {
		// do nothing
	} else if err.Error() == "record not found" {
		// do nothing
	} else if err != nil {
		u.log.Error("error getting user friend invitation data from db: ", err)
		u.DBErrorResponse(c, "error getting user friend invitation data from db")
		return
	}

	if len(friendRequests) == 0 {
		friendRequest := models.Friend{
			UserIdOne:       user.Sub,
			UserIdTwo:       inviteUserInfo.Sub,
			UserOneDecision: 1,
			UserTwoDecision: 0,
		}

		err = u.DB.CreateFriendInvite(&friendRequest)
		if err != nil {
			u.log.Error("error updating user friend invitation data to db: ", err)
			u.DBErrorResponse(c, "error updating user friend invitation data to db")
			return
		}
		note := models.Notification{
			Icon:     user.AvatarUrl,
			Keyword:  user.DisplayName,
			Link:     "/otheruser/stats/" + user.DisplayName,
			Message:  user.DisplayName + " sent you a friend request",
			Type:     "Friend",
			Username: inviteUserInfo.Username,
		}
		_ = u.DB.CreateNotification(&note)
	} else if friendRequests[0].UserTwoDecision == 1 {
		u.BadRequestResponse(c, "The Friend Request has been sent!")
		return
	} else {

		err = u.DB.ApproveFriendInvitation(inviteUserInfo.Sub, user.Sub)
		if err != nil {
			u.log.Error("error updating friend list data to db: ", err)
			u.DBErrorResponse(c, "error updating friend list data to db")
			return
		}
	}

	u.SuccessResponse(c, "Friend request has been sent successfully!", nil)
}

func (u *UserController) HandleFriendInvitationList(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	friendRequest, err := u.DB.GetFriendInvitationByUserId(user.Sub)
	if err != nil {
		u.log.Error("error getting user friend invitation data from db: ", err)
		u.DBErrorResponse(c, "error getting user friend invitation data from db")
		return
	}

	responseData := gin.H{
		"FriendRequestList": friendRequest,
	}

	u.SuccessResponse(c, "", responseData)
}

type FormFriendInvitationAcceptAndReject struct {
	DisplayName string `json:"displayName"`
	IsAccept    bool   `json:"isAccept"`
}

func (u *UserController) HandleFriendInvitationAcceptAndReject(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormFriendInvitationAcceptAndReject{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}

	if form.DisplayName == "" {
		u.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	inviteUserInfo, err := u.DB.GetUserProfileByDisplayName(form.DisplayName)
	if err != nil {
		if err.Error() == "record not found" {
			u.BadRequestResponse(c, "friend requester does not exist")
			return
		} else {
			u.log.Error("error getting friend requester profile", err)
			u.DBErrorResponse(c, "error getting friend requester profile")
			return
		}
	}

	friendRequests, err := u.DB.FetchFriendInviteByUserIds(inviteUserInfo.Sub, user.Sub)
	if err == nil {
		// do nothing
	} else if err.Error() == "record not found" {
		// do nothing
	} else if err != nil {
		u.log.Error("error getting user friend invitation data from db: ", err)
		u.DBErrorResponse(c, "error getting user friend invitation data from db")
		return
	}

	if len(friendRequests) == 0 {
		u.BadRequestResponse(c, "friend request does not exist")
		return
	}

	if friendRequests[0].UserTwoDecision != 0 {
		u.BadRequestResponse(c, "You already approved or rejected this friend request")
		return
	}

	if form.IsAccept {
		err = u.DB.ApproveFriendInvitation(inviteUserInfo.Sub, user.Sub)
		if err != nil {
			u.log.Error("error updating friend list data to db: ", err)
			u.DBErrorResponse(c, "error updating friend list data to db")
			return
		}
	} else if !form.IsAccept {
		err = u.DB.RejectFriendInvitation(inviteUserInfo.Sub, user.Sub)
		if err != nil {
			u.log.Error("error updating friend list data to db: ", err)
			u.DBErrorResponse(c, "error updating friend list data to db")
			return
		}
	}

	u.SuccessResponse(c, "Friend Request Update Success!", nil)
}

type FriendInfo struct {
	DisplayName string `json:"displayName"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (u *UserController) HandleFriendListFetch(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	friends, err := u.DB.GetFriendListByUserId(user.Sub)
	if err != nil {
		u.log.Error("error getting user friend list data from db: ", err)
		u.DBErrorResponse(c, "error getting user friend list data from db")
		return
	}

	friendLists := []FriendInfo{}
	for _, friend := range friends {
		friendList := FriendInfo{}
		if friend.UserIdOne == user.Sub {
			friendInfo, err := u.DB.GetUser(friend.UserIdTwo)
			if err != nil {
				u.log.Error("error getting user ", user.Sub, ": ", err)
			}
			friendList.DisplayName = friendInfo.DisplayName
			friendList.AvatarUrl = friendInfo.AvatarUrl
			friendLists = append(friendLists, friendList)
		} else if friend.UserIdTwo == user.Sub {
			friendInfo, err := u.DB.GetUser(friend.UserIdOne)
			if err != nil {
				u.log.Error("error getting user ", user.Sub, ": ", err)
			}
			friendList.DisplayName = friendInfo.DisplayName
			friendList.AvatarUrl = friendInfo.AvatarUrl
			friendLists = append(friendLists, friendList)
		}
	}

	responseData := gin.H{
		"FriendList": friendLists,
	}

	u.SuccessResponse(c, "", responseData)
}

func (u *UserController) HandleFriendStatus(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	displayName := c.Param("displayName")

	inviteUserInfo, err := u.DB.GetUserProfileByDisplayName(displayName)
	if err != nil {
		u.BadRequestResponse(c, "error getting invitation user profile")
		return
	}

	friendInvitationsReceived, err := u.DB.FetchFriendInviteByUserIds(inviteUserInfo.Sub, user.Sub)
	if err != nil {
		u.log.Error("error getting user friend invitation data from db: ", err)
		u.DBErrorResponse(c, "error getting user friend invitation data from db")
		return
	}

	friendStatus := 0
	if len(friendInvitationsReceived) != 0 {
		if friendInvitationsReceived[0].UserOneDecision+friendInvitationsReceived[0].UserTwoDecision == 2 {
			friendStatus = 2
		}
	}

	if friendStatus == 0 {
		friendInvitationsSent, err := u.DB.FetchFriendInviteByUserIds(user.Sub, inviteUserInfo.Sub)
		if err != nil {
			u.log.Error("error getting user friend invitation data from db: ", err)
			u.DBErrorResponse(c, "error getting user friend invitation data from db")
			return
		}

		if len(friendInvitationsSent) == 0 {
			friendStatus = 0
		} else {
			if friendInvitationsSent[0].UserOneDecision+friendInvitationsSent[0].UserTwoDecision == 2 {
				friendStatus = 2
			} else if friendInvitationsSent[0].UserOneDecision == 1 {
				friendStatus = 1
			}
		}
	}

	responseData := gin.H{
		"friendStatus": friendStatus, //2-friend, 1-friend request sent, 0-not friend
	}
	u.SuccessResponse(c, "", responseData)
}

type FormUserAvatarUpload struct {
	File string `json:"file"`
}

func (u *UserController) HandleUserAvatarUpload(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}
	maxSize := int64(1024000)
	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		u.log.Error("Image too large. Max Size: 1M ", err)
		u.InternalErrorResponse(c, "Image too large. Max Size: 1M")
		return
	}
	file, fileHeader, err := c.Request.FormFile("profile_picture")
	if err != nil {
		u.log.Error("Could not get uploaded file: ", err)
		u.InternalErrorResponse(c, "Could not get uploaded file")
		return
	}
	defer file.Close()
	u.log.Info("fileHeader: ", fileHeader.Filename)
	awsConfig := u.conf.AwsConfig()
	bucketname := awsConfig.S3BucketName
	tempAvatarUrl, _ := u.S3UploadClient.HandleS3AvatarUpload(bucketname, user.PublicAddress, file)

	users, err := u.DB.UpdateUserAvatar(user.Sub, tempAvatarUrl)
	if err != nil {
		u.log.Error("error updating user avatar", err)
		u.InternalErrorResponse(c, "updating user avatar fails")
		return
	}

	u.SuccessResponse(c, "", users)
}

func (u *UserController) HandleUserProfileBannerUpload(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}
	maxSize := int64(1024000)
	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		u.log.Error("Profile Banner Image too large. Max Size: 1M ", err)
		u.InternalErrorResponse(c, "Profile Banner Image too large. Max Size: 1M")
		return
	}
	file, fileHeader, err := c.Request.FormFile("profile_banner")
	if err != nil {
		u.log.Error("Could not get uploaded profile banner file: ", err)
		u.InternalErrorResponse(c, "Could not get uploaded profile banner file")
		return
	}
	defer file.Close()
	u.log.Info("fileHeader: ", fileHeader.Filename)
	awsConfig := u.conf.AwsConfig()
	bucketname := awsConfig.S3BucketName
	tempProfileBannerUrl, _ := u.S3UploadClient.HandleS3ProfileBannerUpload(bucketname, user.PublicAddress, file)

	users, err := u.DB.UpdateUserProfileBanner(user.Sub, tempProfileBannerUrl)
	if err != nil {
		u.log.Error("error updating user profile banner", err)
		u.InternalErrorResponse(c, "updating user profile banner fails")
		return
	}

	u.SuccessResponse(c, "Profile banner upload successfully!", users)
}

type FormUserFeed struct {
	Link    string `json:"link"`
	Message string `json:"message"`
}

func (u *UserController) HandleUserFeedCreate(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	form := FormUserFeed{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		u.BadRequestResponse(c, err.Error())
		return
	}
	if form.Link == "" && form.Message == "" {
		u.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	feed := models.UserFeed{
		Sub:     user.Sub,
		Icon:    user.AvatarUrl,
		Link:    form.Link,
		Message: form.Message,
	}
	err = u.DB.CreateUserFeed(&feed)
	if err != nil {
		u.log.Error("error in CreateUserFeed: " + err.Error())
		u.InternalErrorResponse(c, "error in CreateUserFeed: "+err.Error())
		return
	}

	u.SuccessResponse(c, "", nil)
}

func (u *UserController) HandleUserFeedList(c *gin.Context) {
	displayName := c.Param("displayName")
	userInfo, err := u.DB.GetUserProfileByDisplayName(displayName)
	if err != nil {
		u.log.Error("error getting user profile", err)
		u.DBErrorResponse(c, "error getting user profile")
		return
	}

	feeds, err := u.DB.ReadUserFeed(userInfo.Sub)
	if err != nil {
		u.log.Error("error in ReadUserFeed: ", err)
		u.InternalErrorResponse(c, "error in ReadUserFeed: "+err.Error())
		return
	}

	u.SuccessResponse(c, "", gin.H{"feed": feeds})
}

type ResponseUserBadge struct {
	File *bytes.Reader `json:"file"`
}

// func (u *UserController) HandleUserBadgeGet(c *gin.Context) {
// 	user := u.userFromContext(c)
// 	if user == nil {
//      u.log.Error("user not found")
// 		u.InternalErrorResponse(c, "user not found")
// 		return
// 	}
// 	badgeURL := c.Query("badgeURL")
// 	urlStrings := strings.SplitN(badgeURL, "com/", -1)
// 	urlKey := urlStrings[len(urlStrings)-1]
// 	segments := strings.SplitN(urlKey, "/", -1)
// 	filename := segments[len(segments)-1]
// 	awsConfig := u.conf.AwsConfig()
// 	region := awsConfig.CognitoRegion
// 	bucketname := awsConfig.S3BucketName
// 	file, err := u.S3UploadClient.HandleS3Download(region, bucketname, filename, urlKey)

// 	if err != nil {
// 		u.log.Error("error Get User's Badge", err)
// 		u.InternalErrorResponse(c, "get user badge fails")
// 		return
// 	}
// 	log.Printf("file: ", file)
// 	// file, err := os.Open(filename)
// 	// if err != nil {
// 	// 	exitErrorf("Unable to open file %q, %v", err)
// 	// }

// 	c.Writer.Header().Set("Content-Type", c.Request.Header.Get("Content-Type"))
// 	c.Writer.Header().Set("Content-Length", c.Request.Header.Get("Content-Length"))

// 	//stream the body to the client without fully loading it into memory
// 	io.Copy(c.Writer, file)
// 	u.SuccessResponse(c, "", file)
// }

func (u *UserController) HandleUser2FASend(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	if user.PhoneNumber == "" {
		u.BadRequestResponse(c, "no phone number on record")
		return
	}

	code := u.usvc.Compute2FACode(user.Sub)
	if code == -1 {
		u.log.Error("error in ComputeCode")
		u.InternalErrorResponse(c, "error in ComputeCode")
		return
	}
	err := services.SendSMS("Your Bitspawn verification code: "+fmt.Sprintf("%06d", code), user.PhoneNumber, "messagebird")
	if err != nil {
		u.log.Error("Error in SendSMS: ", err)
		u.InternalErrorResponse(c, "Error in SendSMS")
		return
	}
	u.SuccessResponse(c, "", nil)
}

func (u *UserController) HandleUser2FAVerify(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}
	if !user.Enabled2FA {
		u.BadRequestResponse(c, "user 2FA is not enabled")
		return
	}
	u.SuccessResponse(c, "", nil)
}

func (u *UserController) HandleUser2FADisable(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	err := u.DB.Disable2FA(user.Sub)
	if err != nil {
		u.log.Error("fail to diable 2FA for user: ", user.Username)
		u.InternalErrorResponse(c, "fail to disable 2FA")
		return
	}
	u.SuccessResponse(c, "", nil)
}

func (u *UserController) HandleUser2FAEnable(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		u.InternalErrorResponse(c, "user not found")
		return
	}

	err := u.DB.Enable2FA(user.Sub)
	if err != nil {
		u.log.Error("fail to enable 2FA for user: ", user.Username)
		u.InternalErrorResponse(c, "fail to enable 2FA")
		return
	}
	u.SuccessResponse(c, "", nil)
}

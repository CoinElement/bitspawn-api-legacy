/*
 */

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bitspawngg/bitspawn-api/services/queue"
	"github.com/bitspawngg/bitspawn-api/services/signURL"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/datatypes"

	"github.com/bitspawngg/bitspawn-api/enum"
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services/cognito"
	organizer "github.com/bitspawngg/bitspawn-api/services/match"
	"github.com/bitspawngg/bitspawn-api/services/poa"
	"github.com/bitspawngg/bitspawn-api/services/s3"
	"github.com/bitspawngg/bitspawn-api/services/tournament"
)

type TournamentController struct {
	BaseController
	tsvc           *tournament.TournamentService
	orgSvc         *organizer.MatchService
	bsc            *poa.BitspawnPoaClient
	AuthStore      *cognito.Auth
	S3UploadClient *s3.S3UploadClient
	SignURLManager *signURL.SignURLManager
	SQSClient      *queue.SqsClient
}

func NewTournamentController(bc *BaseController, tsvc *tournament.TournamentService, organizer *organizer.MatchService, bsc *poa.BitspawnPoaClient, authCognito *cognito.Auth, s3UploadClient *s3.S3UploadClient, signURLManager *signURL.SignURLManager, sqsSvc *queue.SQSService) *TournamentController {

	return &TournamentController{
		BaseController{
			Name: "tournament",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		},
		tsvc,
		organizer,
		bsc,
		authCognito,
		s3UploadClient,
		signURLManager,
		sqsSvc.Client(bc.conf.AwsConfig().SQSNameTx, queue.WithDelay(2)),
	}
}

type FormTournamentCreate struct {
	TournamentName        string                `json:"tournamentName" binding:"required"`
	GameType              string                `json:"gameType" binding:"required"`
	GameSubtype           string                `json:"gameSubtype"`
	TournamentFormat      enum.TournamentFormat `json:"tournamentFormat" binding:"required"`
	TournamentDescription string                `json:"tournamentDescription"`
	TournamentRule        string                `json:"tournamentRule"`
	CriticalRules         string                `json:"criticalRules"`
	OrganizerPercentage   int64                 `json:"organizerPercentage"`

	MaxParticipants int    `json:"maxParticipants"`
	MinParticipants int    `json:"minParticipants"`
	MinPrizePool    int64  `json:"minPrizePool"`
	NumberOfTeams   int    `json:"numberOfTeams"`
	EntryFee        int64  `json:"entryFee"`
	FeeType         string `json:"feeType" binding:"required"`
	InviteOnly      bool   `json:"inviteOnly"`

	CutoffDate     time.Time `json:"cutoffDate"`
	TournamentDate time.Time `json:"tournamentDate" binding:"required"`

	BannerUrl    string `json:"bannerUrl"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	LogoUrl      string `json:"logoUrl"`

	Metadata                  json.RawMessage `json:"metadata"`
	Consoles                  []enum.Console  `json:"consoles"`
	MatchDurationMin          int64           `json:"matchDurationMin"`
	CheckInType               string          `json:"checkInType"`
	CheckInBeforeMin          int64           `json:"checkInBeforeMin"`
	ParticipantsPlayEachOther int             `json:"participantsPlayEachOther"`
}

func (tc *TournamentController) HandleTournamentCreate(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentCreate{}

	if err := c.BindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	// TODO: input field validation
	if !enum.FeeType(form.FeeType).IsValid() {
		BadRequestResponseV2(c, form.FeeType+" fee type is not supported", "")
		return
	}

	if form.MinParticipants < 2 {
		form.MinParticipants = 2
	}
	if form.MaxParticipants < form.MinParticipants {
		form.MaxParticipants = 2000
	}
	if form.NumberOfTeams > form.MaxParticipants {
		form.NumberOfTeams = form.MaxParticipants
	}
	if form.NumberOfTeams < 2 {
		form.NumberOfTeams = 2
	}
	if form.CutoffDate.IsZero() {
		form.CutoffDate = time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC)
	}
	if form.CheckInType == "" {
		form.CheckInType = "NONE"
	}
	if !form.TournamentFormat.IsValid() {
		BadRequestResponseV2(c, form.TournamentFormat.ToString()+" format is not supported", "")
		return
	}
	if form.TournamentFormat == enum.RoundRobin && form.ParticipantsPlayEachOther == 0 {
		BadRequestResponseV2(c, "number of round for round-robin format is mandatory", "")
		return
	}
	for _, p := range form.Consoles {
		if !p.IsValid() {
			BadRequestResponseV2(c, p.ToString()+" is not supported for this game type - "+form.GameType, "")
			return
		}
	}
	if len(form.Consoles) == 0 {
		form.Consoles = []enum.Console{enum.PC}
	}
	// End of input field validation

	organizerDetail := models.Organizer{
		DisplayName: user.DisplayName,
		AvatarUrl:   user.AvatarUrl,
	}
	t := models.TournamentResponse{
		OrganizerID: user.Username,
		Organizer:   organizerDetail,

		TournamentName:        form.TournamentName,
		GameType:              form.GameType,
		GameSubtype:           form.GameSubtype,
		TournamentFormat:      form.TournamentFormat,
		TournamentDescription: form.TournamentDescription,
		TournamentRule:        form.TournamentRule,
		CriticalRules:         form.CriticalRules,
		OrganizerPercentage:   form.OrganizerPercentage,

		MaxParticipants: form.MaxParticipants,
		MinParticipants: form.MinParticipants,
		MinPrizePool:    form.MinPrizePool,
		NumberOfTeams:   form.NumberOfTeams,
		EntryFee:        form.EntryFee,
		FeeType:         form.FeeType,
		InviteOnly:      form.InviteOnly,

		CutoffDate:     form.CutoffDate,
		TournamentDate: form.TournamentDate,

		BannerUrl:    form.BannerUrl,
		ThumbnailUrl: form.ThumbnailUrl,
		LogoUrl:      form.LogoUrl,

		Consoles:                  form.Consoles,
		MatchDurationMin:          form.MatchDurationMin,
		CheckInType:               form.CheckInType,
		CheckInBeforeMin:          form.CheckInBeforeMin,
		Metadata:                  datatypes.JSON(form.Metadata),
		ParticipantsPlayEachOther: form.ParticipantsPlayEachOther,
	}

	tResponse, err := tc.DB.CreateTournament(&t)
	if err != nil {
		tc.log.Error("error in Create Tournament: ", err)
		InternalErrorResponseV2(c, "cannot create tournament", err.Error())
		return
	}

	err = tc.orgSvc.CreateManualMatchSchedule(tResponse.TournamentID)
	if err != nil {
		tc.log.Error("error in CreateMatchSchedule for ", tResponse.TournamentID, ": ", err)
		InternalErrorResponseV2(c, "CreateMatchSchedule failed", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

type FormTournamentUpdate struct {
	TournamentID          string  `json:"tournamentId" update:"draft,registration"`
	TournamentName        string  `json:"tournamentName"  update:"draft,registration"`
	GameType              string  `json:"gameType" update:"draft"`
	GameSubtype           string  `json:"gameSubtype" update:"draft"`
	TournamentDescription *string `json:"tournamentDescription" update:""`
	TournamentRule        *string `json:"tournamentRule" update:""`
	CriticalRules         *string `json:"criticalRules" update:""`
	OrganizerPercentage   *int64  `json:"organizerPercentage" update:"draft"`

	MaxParticipants int     `json:"maxParticipants" update:"draft"`
	MinParticipants int     `json:"minParticipants" update:""`
	MinPrizePool    *int64  `json:"minPrizePool" update:"draft"`
	EntryFee        *int64  `json:"entryFee" update:"draft"`
	FeeType         string  `json:"feeType" update:"draft"`
	PrizeAllocation []int64 `json:"prizeAllocation" update:"draft"`
	InviteOnly      *bool   `json:"inviteOnly" update:"draft,registration"`

	CutoffDate     time.Time `json:"cutoffDate" update:"draft,registration"`
	TournamentDate time.Time `json:"tournamentDate" update:""`

	Consoles         []enum.Console `json:"consoles" update:"draft,registration"`
	Sponsors         []string       `json:"sponsors" update:""`
	MatchDurationMin int64          `json:"matchDurationMin" update:""`
	CheckInType      string         `json:"checkInType,omitempty" update:""`
	CheckInBeforeMin int64          `json:"checkInBeforeMin" update:""`
}

func (ft *FormTournamentUpdate) Validate(stage string) error {
	val := reflect.ValueOf(ft)
	if val.Type().Kind() == reflect.Ptr {
		val = val.Elem()
		if !val.IsValid() {
			return fmt.Errorf("type is not valid :%v", val)
		}
	}
	for i := 0; i < val.Type().NumField(); i++ {
		dstValueField := val.Field(i)
		t := val.Type().Field(i)
		if jsonTag := t.Tag.Get("update"); jsonTag != "" && jsonTag != "-" {
			tags := strings.Split(jsonTag, ",")
			for _, tag := range tags {
				if tag == stage {
					switch dstValueField.Kind() {
					case reflect.String:
						if dstValueField.String() == "" {
							return fmt.Errorf("mandatory field %s not provided", t.Name)
						}
					case reflect.Int:
					case reflect.Int64:
						if dstValueField.Int() == 0 {
							return fmt.Errorf("mandatory field %s not provided", t.Name)
						}
					case reflect.Slice:
						if dstValueField.Type().String() == "[]string" && dstValueField.Interface().([]string) == nil {
							return fmt.Errorf("mandatory field %s not provided", t.Name)
						}
						if dstValueField.Type().String() == "json.RawMessage" && dstValueField.Interface().(json.RawMessage) == nil {
							return fmt.Errorf("mandatory field %s not provided", t.Name)
						}
					case reflect.Bool:
						if dstValueField.Bool() {
							return fmt.Errorf("mandatory field %s not provided", t.Name)
						}
					case reflect.Struct:
						if dstValueField.Type().String() == "time.Time" {
							if dstValueField.Interface().(time.Time).IsZero() {
								return fmt.Errorf("mandatory field %s not provided", t.Name)
							}
						}
					case reflect.Ptr:
						if dstValueField.IsNil() {
							return fmt.Errorf("mandatory field %s not provided", t.Name)
						}
					default:
						return fmt.Errorf("filed %s with type %s not supported", t.Name, dstValueField.Kind().String())
					}
				}
			}
		}
	}

	return nil
}
func (tc *TournamentController) HandleTournamentUpdate(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentUpdate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentID := form.TournamentID
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
	if err := form.Validate(strings.ToLower(tournament.Status)); err != nil {
		tc.log.Errorf("validation failed with error:%v ", err)
		BadRequestResponseV2(c, err.Error(), "")
		return
	}
	roleUser := strings.Split(tournament.Roles, ",")
	if tournament.OrganizerID != user.Username && !stringInSlice(user.Username, roleUser) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	t, err := tc.DB.FormatTournamentForResponse(*tournament)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}

	switch strings.ToUpper(tournament.Status) {
	case "DRAFT":
		updateTournamentInDraft(t, form)
	case "REGISTRATION":
		updateTournamentInRegistration(t, form)
	case "STARTED":
		updateTournamentInStarted(t, form)
	case "PAYOUT":
		BadRequestResponseV2(c, "the tournament is being completed", "")
		return
	case "COMPLETED":
		BadRequestResponseV2(c, "the tournament has been completed", "")
		return
	case "CANCELLED":
		BadRequestResponseV2(c, "the tournament has been cancelled", "")
		return
	default:
		InternalErrorResponseV2(c, "the tournament status is corrupt", "")
		return
	}

	// TODO: input field validation
	if !enum.FeeType(t.FeeType).IsValid() {
		BadRequestResponseV2(c, form.FeeType+" fee type is not supported", "")
		return
	}

	var percentCount int64
	for _, percent := range form.PrizeAllocation {
		percentCount += percent
	}
	if percentCount != 100 {
		BadRequestResponseV2(c, "Prize allocation does not add up to 100%", "")
		return
	}

	numberOfWinners := len(t.PrizeAllocation)
	if t.MinParticipants < numberOfWinners {
		t.MinParticipants = numberOfWinners
	}
	if form.MinParticipants < 2 {
		form.MinParticipants = 2
	}
	if t.MaxParticipants < t.MinParticipants {
		t.MaxParticipants = 2000
	}
	if t.NumberOfTeams > t.MaxParticipants {
		t.NumberOfTeams = t.MaxParticipants
	}

	if !t.TournamentFormat.IsValid() {
		BadRequestResponseV2(c, t.TournamentFormat.ToString()+" format is not supported", "")
		return
	}
	for _, p := range t.Consoles {
		if !p.IsValid() {
			BadRequestResponseV2(c, p.ToString()+" is not supported for this game type - "+t.GameType, "")
			return
		}
	}
	if len(form.Consoles) == 0 {
		form.Consoles = []enum.Console{enum.PC}
	}
	// End of input field validation

	err = tc.DB.UpdateTournament(tournamentID, t)
	if err != nil {
		tc.log.Error("error in Update Tournament: ", err)
		InternalErrorResponseV2(c, "cannot update tournament", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": t,
	})
}

func updateTournamentInDraft(t *models.TournamentResponse, form FormTournamentUpdate) {
	t.TournamentName = form.TournamentName
	t.GameType = form.GameType
	t.GameSubtype = form.GameSubtype
	t.MaxParticipants = form.MaxParticipants
	t.MinParticipants = form.MinParticipants
	t.MinPrizePool = *form.MinPrizePool
	t.TournamentDescription = *form.TournamentDescription
	t.TournamentRule = *form.TournamentRule
	t.CriticalRules = *form.CriticalRules
	t.EntryFee = *form.EntryFee
	t.FeeType = form.FeeType
	t.PrizeAllocation = form.PrizeAllocation
	t.CutoffDate = form.CutoffDate
	t.TournamentDate = form.TournamentDate
	t.CheckInType = form.CheckInType
	t.CheckInBeforeMin = form.CheckInBeforeMin
	t.OrganizerPercentage = *form.OrganizerPercentage
	t.InviteOnly = *form.InviteOnly
	t.Consoles = form.Consoles
}

func updateTournamentInRegistration(t *models.TournamentResponse, form FormTournamentUpdate) {
	t.TournamentName = form.TournamentName
	t.MaxParticipants = form.MaxParticipants
	t.TournamentDescription = *form.TournamentDescription
	t.TournamentRule = *form.TournamentRule
	t.CriticalRules = *form.CriticalRules
	t.CutoffDate = form.CutoffDate
	t.TournamentDate = form.TournamentDate
	t.CheckInType = form.CheckInType
	t.CheckInBeforeMin = form.CheckInBeforeMin
	t.InviteOnly = *form.InviteOnly
	t.Consoles = form.Consoles
}

func updateTournamentInStarted(t *models.TournamentResponse, form FormTournamentUpdate) {
}

type FormTournamentModeratorUpdate struct {
	TournamentId string `json:"tournamentId"`
	Moderator    string `json:"moderator"`
}

func (tc *TournamentController) HandleTournamentModeratorAdd(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentModeratorUpdate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentID := form.TournamentId
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}
	if form.Moderator == "" {
		BadRequestResponseV2(c, "missing input moderator display name", "")
		return
	}

	tournament, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", tournamentID, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}

	moderatorUsernames := strings.Split(tournament.Roles, ",")
	if tournament.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	moderatorToAdd, err := tc.tsvc.DB.GetUserProfileByDisplayName(form.Moderator)
	if err != nil {
		BadRequestResponseV2(c, form.Moderator+" - this player does not exist", "")
		return
	}

	if stringInSlice(moderatorToAdd.Username, moderatorUsernames) {
		BadRequestResponseV2(c, form.Moderator+" - this player is already a moderator", "")
		return
	}
	moderatorUsernames = append(moderatorUsernames, moderatorToAdd.Username)
	updatedModerators := strings.Join(moderatorUsernames, ",")
	tournament.Roles = updatedModerators

	tResponse, err := tc.DB.FormatTournamentForResponse(*tournament)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}

	err = tc.DB.UpdateTournamentData(tournamentID, tournament)
	if err != nil {
		tc.log.Error("error updating tournament moderators to db: ", err)
		InternalErrorResponseV2(c, "error updating tournament moderators to db ", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

func (tc *TournamentController) HandleTournamentModeratorDelete(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentModeratorUpdate{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentID := form.TournamentId
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}
	if form.Moderator == "" {
		BadRequestResponseV2(c, "missing input moderator display name", "")
		return
	}

	tournament, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", tournamentID, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}

	moderatorUsernames := strings.Split(tournament.Roles, ",")
	if tournament.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	moderatorToDelete, err := tc.tsvc.DB.GetUserProfileByDisplayName(form.Moderator)
	if err != nil {
		BadRequestResponseV2(c, form.Moderator+" - this player does not exist", "")
		return
	}

	if !stringInSlice(moderatorToDelete.Username, moderatorUsernames) {
		BadRequestResponseV2(c, form.Moderator+" - this player is not a moderator", "")
		return
	}

	indexToRemove := -1
	for i := range moderatorUsernames {
		if moderatorToDelete.Username == moderatorUsernames[i] {
			indexToRemove = i
		}
	}
	updatedModeratorUsernames := RemoveIndex(moderatorUsernames, indexToRemove)
	updatedModerators := strings.Join(updatedModeratorUsernames, ",")
	tournament.Roles = updatedModerators

	tResponse, err := tc.DB.FormatTournamentForResponse(*tournament)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}

	err = tc.DB.UpdateTournamentData(tournamentID, tournament)
	if err != nil {
		tc.log.Error("error updating tournament moderators to db: ", err)
		InternalErrorResponseV2(c, "error updating tournament moderators to db ", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

type Team struct {
	TeamName string   `json:"teamName"`
	Players  []string `json:"players"`
}

func (tc *TournamentController) HandleTournamentTeamName(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := struct {
		TournamentId string `json:"tournamentId"`
		TeamId       string `json:"teamId"`
		TeamName     string `json:"teamName"`
	}{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TournamentId == "" || form.TeamId == "" || form.TeamName == "" {
		BadRequestResponseV2(c, "missing mandatory input parameters", "")
		return
	}

	tourney, err := tc.DB.GetTournamentData(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", form.TournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	if strings.ToUpper(tourney.Status) != "READY" && strings.ToUpper(tourney.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "tournament is not in Registration or Ready state", "")
		return
	}

	moderatorUsernames := strings.Split(tourney.Roles, ",")
	if tourney.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	err = tc.DB.UpdateTeamName(form.TeamId, form.TeamName)
	if err != nil {
		tc.log.Error("error in Update Team Name: ", err)
		InternalErrorResponseV2(c, "error in Update Team Name", err.Error())
		return
	}

	matches, err := tc.tsvc.DB.GetMatchesByTournament(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetMatchesByTournament: ", err)
		InternalErrorResponseV2(c, "error in GetMatchesByTournament", err.Error())
		return
	}
	matchesOutput, _ := tc.DB.FormatMatchesForOutput(matches)

	SuccessResponseV2(c, gin.H{
		"tournamentId": form.TournamentId,
		"matchList":    matchesOutput,
	})
}

type FormTournamentPlayersAddRemove struct {
	TournamentId string   `json:"tournamentId"`
	TeamId       string   `json:"teamId"`
	Players      []string `json:"players"`
}

func (tc *TournamentController) HandleTournamentTeamPlayersAdd(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentPlayersAddRemove{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TournamentId == "" || form.TeamId == "" || len(form.Players) == 0 {
		BadRequestResponseV2(c, "missing mandatory input parameters", "")
		return
	}

	tourney, err := tc.DB.GetTournamentData(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", form.TournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	if strings.ToUpper(tourney.Status) != "READY" && strings.ToUpper(tourney.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "tournament is not in Registration or Ready state", "")
		return
	}

	moderatorUsernames := strings.Split(tourney.Roles, ",")
	if tourney.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	teams, err := tc.DB.GetTeams(form.TeamId)
	if err != nil {
		tc.log.Error("error in GetPlayersByTeam: ", err)
		InternalErrorResponseV2(c, "error in GetPlayersByTeam", err.Error())
		return
	}
	if len(teams) == 0 {
		BadRequestResponseV2(c, "team does not exist", "")
		return
	}

	allParticipants, err := tc.DB.GetAllParticipantsInfo(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetParticipants", err.Error())
		return
	}
	if len(allParticipants) == 0 {
		BadRequestResponseV2(c, "there are no participants in this tournament", "")
		return
	}
	var allParticipantsUsernames []string
	for _, p := range allParticipants {
		allParticipantsUsernames = append(allParticipantsUsernames, p.Username)
	}

	assignedParticipants, err := tc.DB.GetAssignedParticipants(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetAssignedParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetAssignedParticipants", err.Error())
		return
	}
	var assignedParticipantsUsernames []string
	for _, p := range assignedParticipants {
		assignedParticipantsUsernames = append(assignedParticipantsUsernames, p.Username)
	}

	var teamsToAdd []models.Team
	for _, displayName := range form.Players {
		player, err := tc.tsvc.DB.GetUserProfileByDisplayName(displayName)
		if err != nil {
			BadRequestResponseV2(c, displayName+" - this player does not exist", "")
			return
		}
		if !stringInSlice(player.Username, allParticipantsUsernames) {
			BadRequestResponseV2(c, displayName+" - this player is not a participant of this tournament", "")
			return
		}
		if stringInSlice(player.Username, assignedParticipantsUsernames) {
			BadRequestResponseV2(c, displayName+" - this player is already on a team", "")
			return
		}
		playerToAdd := models.UserInfo{
			Username:    player.Username,
			DisplayName: player.DisplayName,
			AvatarUrl:   player.AvatarUrl,
		}
		assignedParticipantsUsernames = append(assignedParticipantsUsernames, playerToAdd.Username)
		assignedParticipants = append(assignedParticipants, playerToAdd)
		newTeam := models.Team{
			TournamentID: teams[0].TournamentID,
			TeamID:       teams[0].TeamID,
			TeamName:     teams[0].TeamName,
			Player:       playerToAdd.Username,
		}
		newTeam.Player = player.Username
		teamsToAdd = append(teamsToAdd, newTeam)
	}
	err = tc.DB.CreateTeams(teamsToAdd)
	if err != nil {
		tc.log.Error("error in CreateTeams: ", err)
		InternalErrorResponseV2(c, "error in CreateTeams", err.Error())
		return
	}

	unassignedParticipants := []models.UserInfo{}
	for _, p := range allParticipants {
		if !stringInSlice(p.Username, assignedParticipantsUsernames) {
			unassignedParticipants = append(unassignedParticipants, p)
		}
	}

	matches, err := tc.tsvc.DB.GetMatchesByTournament(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetMatchesByTournament: ", err)
		InternalErrorResponseV2(c, "error in GetMatchesByTournament", err.Error())
		return
	}
	matchesOutput, _ := tc.DB.FormatMatchesForOutput(matches)

	SuccessResponseV2(c, gin.H{
		"tournamentId":           form.TournamentId,
		"matchList":              matchesOutput,
		"assignedParticipants":   assignedParticipants,
		"unassignedParticipants": unassignedParticipants,
	})
}

func (tc *TournamentController) HandleTournamentTeamPlayersRemove(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentPlayersAddRemove{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	if form.TournamentId == "" || form.TeamId == "" || len(form.Players) == 0 {
		BadRequestResponseV2(c, "missing mandatory input parameters", "")
		return
	}

	tourney, err := tc.DB.GetTournamentData(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", form.TournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	if strings.ToUpper(tourney.Status) != "READY" && strings.ToUpper(tourney.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "tournament is not in Registration or Ready state", "")
		return
	}

	moderatorUsernames := strings.Split(tourney.Roles, ",")
	if tourney.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	teams, err := tc.DB.GetTeams(form.TeamId)
	if err != nil {
		tc.log.Error("error in GetPlayersByTeam: ", err)
		InternalErrorResponseV2(c, "error in GetPlayersByTeam", err.Error())
		return
	}
	if len(teams) == 0 {
		BadRequestResponseV2(c, "team does not exist", "")
		return
	}

	allParticipants, err := tc.DB.GetAllParticipantsInfo(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetParticipants", err.Error())
		return
	}
	if len(allParticipants) == 0 {
		BadRequestResponseV2(c, "there are no participants in this tournament", "")
		return
	}
	var allParticipantsUsernames []string
	for _, p := range allParticipants {
		allParticipantsUsernames = append(allParticipantsUsernames, p.Username)
	}

	assignedParticipants, err := tc.DB.GetAssignedParticipants(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetAssignedParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetAssignedParticipants", err.Error())
		return
	}
	var assignedParticipantsUsernames []string
	for _, p := range assignedParticipants {
		assignedParticipantsUsernames = append(assignedParticipantsUsernames, p.Username)
	}

	var playerUsernamesToRemove []string
	for _, displayName := range form.Players {
		player, err := tc.tsvc.DB.GetUserProfileByDisplayName(displayName)
		if err != nil {
			BadRequestResponseV2(c, displayName+" - this player does not exist", "")
			return
		}
		if !stringInSlice(player.Username, allParticipantsUsernames) {
			BadRequestResponseV2(c, displayName+" - this player is not a participant of this tournament", "")
			return
		}
		if !stringInSlice(player.Username, assignedParticipantsUsernames) {
			BadRequestResponseV2(c, displayName+" - this player is not on a team", "")
			return
		}
		indexToRemove := -1
		for i := range assignedParticipantsUsernames {
			if player.Username == assignedParticipantsUsernames[i] {
				indexToRemove = i
			}
		}
		assignedParticipantsUsernames = RemoveIndex(assignedParticipantsUsernames, indexToRemove)
		assignedParticipants = RemoveIndexUserInfo(assignedParticipants, indexToRemove)
		playerUsernamesToRemove = append(playerUsernamesToRemove, player.Username)
	}
	err = tc.DB.DeleteTeams(teams[0].TeamID, playerUsernamesToRemove)
	if err != nil {
		tc.log.Error("error in DeleteTeams: ", err)
		InternalErrorResponseV2(c, "error in DeleteTeams", err.Error())
		return
	}

	unassignedParticipants := []models.UserInfo{}
	for _, p := range allParticipants {
		if !stringInSlice(p.Username, assignedParticipantsUsernames) {
			unassignedParticipants = append(unassignedParticipants, p)
		}
	}

	matches, err := tc.tsvc.DB.GetMatchesByTournament(form.TournamentId)
	if err != nil {
		tc.log.Error("error in GetMatchesByTournament: ", err)
		InternalErrorResponseV2(c, "error in GetMatchesByTournament", err.Error())
		return
	}
	matchesOutput, _ := tc.DB.FormatMatchesForOutput(matches)

	SuccessResponseV2(c, gin.H{
		"tournamentId":           form.TournamentId,
		"matchList":              matchesOutput,
		"assignedParticipants":   assignedParticipants,
		"unassignedParticipants": unassignedParticipants,
	})
}

func RemoveIndexUserInfo(s []models.UserInfo, index int) []models.UserInfo {
	ret := make([]models.UserInfo, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func stringInSlice(username string, rolelist []string) bool {
	for _, role := range rolelist {
		if role == username {
			return true
		}
	}
	return false
}

func (tc *TournamentController) HandleTournamentPublish(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	// Begin user Account validation
	_, err := tc.getPrivKey(c)
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

	// Begin tournament data validation
	req := struct {
		TournamentID string `json:"tournamentId"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentID := req.TournamentID
	if tournamentID == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tourney, err := tc.DB.GetTournamentData(tournamentID)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", tournamentID, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}

	if tourney.OrganizerID != user.Username {
		BadRequestResponseV2(c, "No Permission to publish tournament", "")
		return
	}

	t, err := tc.DB.FormatTournamentForResponse(*tourney)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}

	if t.Status != "DRAFT" {
		BadRequestResponseV2(c, "the tournament has been published", "")
		return
	}

	if len(t.PrizeAllocation) == 0 {
		BadRequestResponseV2(c, "the tournament need to set prize distribution before publish", "")
		return
	}
	// End of tournament data validation

	meta := tournament.Meta{}
	meta.Action = string(tournament.DEPLOY)
	meta.Author = user.Sub
	meta.TournamentID = tourney.TournamentID
	formJSON, err := json.Marshal(meta)
	if err != nil {
		tc.log.Error("error in marshal meta: " + err.Error())
		InternalErrorResponseV2(c, "error in marshal meta: ", err.Error())
		return
	}
	err = tc.SQSClient.SendMsg(queue.Message{Title: string(tournament.DEPLOY), Author: user.Sub, Body: string(formJSON)})
	if err != nil {
		tc.log.Error("error in HandleSendMessage: " + err.Error())
		tc.InternalErrorResponse(c, "error in HandleSendMessage: "+err.Error())
		return
	}

	err = tc.DB.UpdateTournamentStatus(tournamentID, "REGISTRATION")
	if err != nil {
		tc.log.Error("error updating tournament status to Registration: ", err)
		tc.DBErrorResponse(c, err.Error())
		return
	}

	t.Status = "REGISTRATION"
	SuccessResponseV2(c, gin.H{
		"tournamentDetails": t,
	})
}

func (tc *TournamentController) HandleV2TournamentComplete(c *gin.Context) {
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
	tournamentId := form.TournamentID

	tournamentData, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	tResponse, err := tc.DB.FormatTournamentForResponse(*tournamentData)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}

	moderatorUsernames := strings.Split(tournamentData.Roles, ",")
	if tournamentData.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to complete tournament", user.Username)
		return
	}
	if strings.ToUpper(tournamentData.Status) != "STARTED" {
		BadRequestResponseV2(c, "tournament not in Started status", "")
		return
	}

	err = tc.orgSvc.CompleteTournament(tResponse)
	if err != nil {
		tc.log.Errorf("error in CompleteTournament: %v", err)
		InternalErrorResponseV2(c, "error in CompleteTournament", err.Error())
		return
	}
	tResponse.Status = "PAYOUT" // Already updated status in DB, also update status in response

	organizer, err := tc.DB.GetUserProfile(tournamentData.OrganizerID)
	if err != nil {
		tc.log.Error("error in GetUserProfile: ", err)
		InternalErrorResponseV2(c, "error in GetUserProfile", err.Error())
		return
	}
	err = tc.checkEth(organizer)
	if err != nil {
		tc.log.Error(err)
		InternalErrorResponseV2(c, "fail to fill gas", err.Error())
		return
	}

	form.Action = string(tournament.COMPLETE)
	form.Author = organizer.Sub
	form.TournamentID = tournamentId
	formJSON, err := json.Marshal(form)
	if err != nil {
		tc.log.Error("error in marshal form meta: " + err.Error())
		InternalErrorResponseV2(c, "error in marshal form meta", err.Error())
		return
	}
	err = tc.SQSClient.SendMsg(queue.Message{Title: string(tournament.COMPLETE), Author: organizer.Sub, Body: string(formJSON)})
	if err != nil {
		tc.log.Error("error in HandleSendMessage: " + err.Error())
		InternalErrorResponseV2(c, "error in HandleSendMessage", err.Error())
		return
	}

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

func (tc *TournamentController) HandleV2TournamentCancel(c *gin.Context) {
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
	tournamentId := form.TournamentID

	tournamentData, err := tc.tsvc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	tResponse, err := tc.DB.FormatTournamentForResponse(*tournamentData)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}

	moderatorUsernames := strings.Split(tournamentData.Roles, ",")
	if tournamentData.OrganizerID != user.Username && !stringInSlice(user.Username, moderatorUsernames) {
		BadRequestResponseV2(c, "Not Authorized to update tournament", user.Username)
		return
	}

	if strings.ToUpper(tournamentData.Status) != "DRAFT" &&
		strings.ToUpper(tournamentData.Status) != "REGISTRATION" &&
		strings.ToUpper(tournamentData.Status) != "STARTED" {
		BadRequestResponseV2(c, "tournament not in Draft, Registration or Started status", "")
		return
	}

	if tResponse.Status == "DRAFT" {
		err = tc.DB.UpdateTournamentStatus(tournamentId, "CANCELLED")
		if err != nil {
			tc.log.Error("error updating tournament status to Cancelled: ", err)
			InternalErrorResponseV2(c, "error updating tournament status to Cancelled", err.Error())
			return
		}
	} else {
		organizer, err := tc.DB.GetUserProfile(tournamentData.OrganizerID)
		if err != nil {
			tc.log.Error("error in GetUserProfile: ", err)
			InternalErrorResponseV2(c, "error in GetUserProfile", err.Error())
			return
		}
		err = tc.checkEth(organizer)
		if err != nil {
			tc.log.Error(err)
			InternalErrorResponseV2(c, "fail to fill gas", err.Error())
			return
		}

		form.Action = string(tournament.CANCEL)
		form.Author = organizer.Sub
		form.TournamentID = tournamentId
		formJSON, err := json.Marshal(form)
		if err != nil {
			tc.log.Error("error in marshal form meta: " + err.Error())
			InternalErrorResponseV2(c, "error in marshal form meta: ", err.Error())
			return
		}
		err = tc.SQSClient.SendMsg(queue.Message{Title: string(tournament.CANCEL), Author: organizer.Sub, Body: string(formJSON)})
		if err != nil {
			tc.log.Error("error in HandleSendMessage: " + err.Error())
			InternalErrorResponseV2(c, "error in HandleSendMessage: ", err.Error())
			return
		}
	}

	tResponse.Status = "CANCELLED"
	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

type FormTournamentStart struct {
	TournamentID string `json:"tournamentId"`
}

func (tc *TournamentController) HandleTournamentStart(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := FormTournamentStart{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	tournamentId := form.TournamentID
	if tournamentId == "" {
		BadRequestResponseV2(c, "missing input tournament id", "")
		return
	}

	tourney, err := tc.DB.GetTournamentData(tournamentId)
	if err != nil {
		tc.log.Error("error in GetTournamentData - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error getting tournament", err.Error())
		return
	}
	if strings.ToUpper(tourney.Status) != "READY" && strings.ToUpper(tourney.Status) != "REGISTRATION" {
		BadRequestResponseV2(c, "tournament is not in Registration or Ready state", "")
		return
	}

	if tourney.TotalPrizePool < tourney.MinPrizePool {
		BadRequestResponseV2(c, "total prize pool is insufficient", "")
		return
	}

	allParticipants, err := tc.DB.GetAllParticipantsInfo(tournamentId)
	if err != nil {
		tc.log.Error("error in GetParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetParticipants", err.Error())
		return
	}
	if len(allParticipants) == 0 {
		BadRequestResponseV2(c, "there are no participants in this tournament", "")
		return
	}

	assignedParticipants, err := tc.DB.GetAssignedParticipants(tournamentId)
	if err != nil {
		tc.log.Error("error in GetAssignedParticipants: ", err)
		InternalErrorResponseV2(c, "error in GetAssignedParticipants", err.Error())
		return
	}
	if len(assignedParticipants) < tourney.MinParticipants {
		BadRequestResponseV2(c, "number of assigned participants is insufficient", "")
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

	for _, p := range unassignedParticipants {
		deletedRecord, err := tc.tsvc.DB.DeletePlayRecord(p.Username, tourney.TournamentID)
		if err != nil {
			tc.log.Error("error in DeletePlayRecord: " + err.Error())
		}

		meta := tournament.Meta{}
		if tourney.EntryFee != "0" { // Only write blockchain if there is entry fee
			meta.Action = string(tournament.UNREGISTER)
			meta.Author = p.Sub
			meta.TournamentID = tourney.TournamentID
			formJSON, _ := json.Marshal(form)
			err = tc.SQSClient.SendMsg(queue.Message{Title: string(tournament.UNREGISTER), Author: p.Sub, Body: string(formJSON)})
			if err != nil {
				tc.log.Error("error in HandleSendMessage: " + err.Error())
				_ = tc.tsvc.DB.InsertPlayRecord(deletedRecord)
				return
			}
		}
	}

	tResponse, err := tc.DB.FormatTournamentForResponse(*tourney)
	if err != nil {
		tc.log.Error("error in Format tournament for response: ", err)
		InternalErrorResponseV2(c, "error in Format tournament for response", err.Error())
		return
	}

	err = tc.orgSvc.StartManualTournament(tournamentId)
	if err != nil {
		tc.log.Error("error in StartManualTournament - ", tournamentId, ": ", err)
		InternalErrorResponseV2(c, "error in StartManualTournament", err.Error())
		return
	}

	tResponse.Status = "STARTED"

	SuccessResponseV2(c, gin.H{
		"tournamentDetails": tResponse,
	})
}

func (tc *TournamentController) HandleAdminTournamentComplete(c *gin.Context) {
	tournamentsToComplete, err := tc.DB.GetTournamentsByStatus("Completing")
	if err != nil {
		tc.log.Error(err)
		return
	}

	tournamentsCompleted := []string{}
	for _, tournament := range tournamentsToComplete {
		if tournament.UpdatedAt.After(time.Now().Add(-480 * time.Minute).UTC()) {
			// wait 8 hours before admin completes the tournament
			continue
		}

		tournamentId := tournament.TournamentID
		tc.log.Info("Completing tournament: ", tournamentId)

		user, err := tc.DB.GetUserProfile(tournament.OrganizerID)
		if err != nil {
			tc.log.Error("error getting user ", tournament.OrganizerID, ": ", err)
			continue
		}
		err = tc.checkEth(user)
		if err != nil {
			tc.log.Error(err)
			return
		}
		err = tc.tsvc.ExecuteStateMachine(tournamentId, tournament.OrganizerID, "Complete", nil)
		if err != nil {
			tc.log.Error("error in complete tournament: ", err)
			continue
		}

		tournamentsCompleted = append(tournamentsCompleted, tournamentId)
	}

	tc.SuccessResponse(c, "tournaments completed", tournamentsCompleted)
}

func (tc *TournamentController) HandleAdminTournamentCancel(c *gin.Context) {
	tournamentsToCancel, err := tc.DB.GetTournamentsOverdue()
	if err != nil {
		tc.log.Error(err)
		tc.InternalErrorResponse(c, "error in GetTournamentsOverdue: "+err.Error())
		return
	}

	tournamentsCancelled := []string{}
	for _, tournament := range tournamentsToCancel {
		tournamentId := tournament.TournamentID
		tc.log.Info("Cancelling tournament: ", tournamentId)

		user, err := tc.DB.GetUserProfile(tournament.OrganizerID)
		if err != nil {
			tc.log.Error("error getting user ", tournament.OrganizerID, ": ", err)
			continue
		}
		err = tc.checkEth(user)
		if err != nil {
			tc.log.Error(err)
			return
		}
		err = tc.tsvc.ExecuteStateMachine(tournamentId, tournament.OrganizerID, "Cancel", nil)
		if err != nil {
			tc.log.Error("error in cancel tournament: ", err)
			continue
		}

		tournamentsCancelled = append(tournamentsCancelled, tournamentId)
	}

	tc.SuccessResponse(c, "tournaments cancelled", tournamentsCancelled)
}

func (tc *TournamentController) HandleAdminTournamentDequeue(c *gin.Context) {
	chnMessages := make(chan *sqs.Message, SQS_MAX_MESSAGES)
	// errc := make(chan error)
	errs, _ := errgroup.WithContext(context.Background())
	for i := 0; i < SQS_MAX_MESSAGES; i++ {
		errs.Go(func() error {
			return tc.SQSClient.ShortPollSqs(chnMessages)
		})
	}
	if err := errs.Wait(); err != nil {
		tc.log.Error("error in ShortPollSqs: " + err.Error())
		tc.InternalErrorResponse(c, "error in ShortPollSqs: "+err.Error())
		return
	}
	close(chnMessages) // close the channels after all sending is done

	chnErrors := make(chan error, SQS_MAX_MESSAGES)
	var wg sync.WaitGroup
	for message := range chnMessages {
		wg.Add(1)
		go func(message *sqs.Message) {
			defer wg.Done()
			fmt.Printf("message: %v\n", message)
			meta := tournament.Meta{}
			err := json.Unmarshal([]byte(*message.Body), &meta)
			if err != nil {
				tc.log.Error("error in unmarshal message body: " + err.Error())
				chnErrors <- fmt.Errorf("error in unmarshal message body: %v", err)

				return
			}
			err = tc.tsvc.ExecuteStateMachine(meta.TournamentID, meta.Author, tournament.Action(meta.Action), &meta)
			if err != nil {
				tc.log.Error("error in ExecuteStateMachine: ", err)
				chnErrors <- fmt.Errorf("error in ExecuteStateMachine: %v", err)
				return
			}
			err = tc.SQSClient.DeleteMsg(message)
			if err != nil {
				tc.log.Error("error in DeleteMsg: ", err)
				chnErrors <- fmt.Errorf("error in DeleteMsg: %v", err)
				return
			}
			tournamentInfo, err := tc.DB.GetTournamentData(meta.TournamentID)
			if err != nil {
				tc.log.Error("error in GetTournamentData for tournament ", meta.TournamentID, ": ", err)
				chnErrors <- fmt.Errorf("error in GetTournamentData: %v", err)
				return
			}
			note := models.Notification{
				Icon:     tournamentInfo.ThumbnailUrl,
				Keyword:  tournamentInfo.TournamentName,
				Link:     "/tournament/fetch/" + meta.TournamentID,
				Message:  "You have successfully " + meta.Action + "ed Tournament " + tournamentInfo.TournamentName,
				Type:     "Tournament",
				Username: meta.Author,
			}
			_ = tc.DB.CreateNotification(&note)
		}(message)
	}
	wg.Wait()
	close(chnErrors)

	var errorLogs []string
	for errorLog := range chnErrors {
		errorLogs = append(errorLogs, errorLog.Error())
	}
	tc.SuccessResponse(c, "tournament executed sccessfully", errorLogs)
}

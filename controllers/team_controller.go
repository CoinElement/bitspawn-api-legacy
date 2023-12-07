package controllers

import (
	"errors"
	"net/http"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services/hdkey"
	"github.com/bitspawngg/bitspawn-api/services/poa"
	"github.com/bitspawngg/bitspawn-api/services/s3"
	"github.com/bitspawngg/bitspawn-api/services/team"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	*BaseController
	teamService    *team.TeamService
	S3UploadClient *s3.S3UploadClient
	bsc            *poa.BitspawnPoaClient
}

func NewTeamController(bc *BaseController, s3UploadClient *s3.S3UploadClient, bsc *poa.BitspawnPoaClient) *TeamController {

	return &TeamController{
		&BaseController{
			Name: "team",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		}, team.NewTeamService(bc.DB, bc.log), s3UploadClient, bsc,
	}
}
func (tc *TeamController) HandleTeamRegister(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	team := models.TeamMachine{}

	if err := c.BindJSON(&team); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if !team.Publicity.Valid() || !team.GenrePreferred.Valid() {
		tc.BadRequestResponse(c, "mandatory fields has been missed")
		return
	}
	if err := tc.teamService.CreateTeam(&team, user.Sub); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			tc.BadRequestResponse(c, "duplicated request")
			return
		}
		InternalErrorResponseV2(c, "creating team has failed", err.Error())
		return
	}
	respBody := map[string]interface{}{
		"teamId": team.ID,
	}
	CustomSuccessResponseV2(c, respBody, http.StatusCreated)
}

func (tc *TeamController) HandleTeamUnRegister(c *gin.Context) {
	if user := tc.userFromContext(c); user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	if err := tc.teamService.DeleteTeam(teamID); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "deleting team has failed", err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (tc *TeamController) HandleTeamList(c *gin.Context) {
	if user := tc.userFromContext(c); user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	name := c.Query("name")
	var (
		teams []*models.TeamMachine
		err   error
	)
	if len(name) == 0 {
		if teams, err = tc.teamService.List(); err != nil {
			InternalErrorResponseV2(c, "fetching team list has failed", err.Error())
		}
	} else {

		if teams, err = tc.teamService.GetTeamsByName(name); err != nil {
			if errors.Is(err, models.ErrRecordNotFound) {
				tc.NotFoundResponse(c, "the team not founded")
				return
			}
			InternalErrorResponseV2(c, "fetching team data has failed", err.Error())
			return
		}
	}

	SuccessResponseV2(c, teams)
}

func (tc *TeamController) HandleGetTeam(c *gin.Context) {
	if user := tc.userFromContext(c); user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	team, err := tc.teamService.GetTeam(teamID)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "fetching team data has failed", err.Error())
		return
	}

	SuccessResponseV2(c, team)
}

func (tc *TeamController) HandleGetTeamMember(c *gin.Context) {
	if user := tc.userFromContext(c); user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	memberID := c.Param("member_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "member id is empty", "")
		return
	}
	member, err := tc.teamService.GetTeamMember(teamID, memberID)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "fetching team member data has failed", err.Error())
		return
	}

	SuccessResponseV2(c, member)
}

func (tc *TeamController) HandleRemoveMember(c *gin.Context) {
	if user := tc.userFromContext(c); user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	memberID := c.Param("member_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "member id is empty", "")
		return
	}
	if err := tc.teamService.DeleteTeamMember(teamID, memberID); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "deleting team member has failed", err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (tc *TeamController) HandleCreateTeam(c *gin.Context) {

	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	team := models.TeamMachine{}

	if err := tc.teamService.CreateTeam(&team, user.Sub); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "creating team has failed", err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

func (tc *TeamController) HandleAddTeamMember(c *gin.Context) {
	if user := tc.userFromContext(c); user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}

	request := &models.JoinRequestBody{}
	if err := c.BindJSON(request); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if !request.Action.Valid() {
		tc.BadRequestResponse(c, "mandatory fields has been missed")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	request.TeamID = teamID
	if err := tc.teamService.CreateTeamMember(request); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		if errors.Is(err, models.ErrTeamNotOpen) {
			tc.BadRequestResponse(c, "mandatory fields has been missed")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			tc.BadRequestResponse(c, "duplicated request")
			return
		}
		InternalErrorResponseV2(c, "creating team member has failed", err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

func (tc *TeamController) HandleChangeRole(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	payload := struct {
		Role   models.Role `json:"role"`
		UserID string      `json:"userId"`
	}{}

	if err := c.BindJSON(&payload); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if !payload.Role.Valid() {
		tc.BadRequestResponse(c, "mandatory fields has been missed")
		return
	}
	if err := tc.teamService.ChangeRole(teamID, payload.UserID, user.Sub, payload.Role); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			tc.BadRequestResponse(c, "duplicated request")
			return
		}
		InternalErrorResponseV2(c, "changing role has failed", err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (tc *TeamController) HandleChangeTeam(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	team := &models.TeamMachine{}

	if err := c.BindJSON(team); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if !team.Publicity.Valid() || !team.GenrePreferred.Valid() {
		tc.BadRequestResponse(c, "mandatory fields has been missed")
		return
	}
	if err := tc.teamService.UpdateTeam(teamID, team); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "updating team has failed", err.Error())
		return
	}
	c.Status(http.StatusOK)
}
func (tc *TeamController) HandleMembershipStatus(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	memberID := c.Param("member_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "member id is empty", "")
		return
	}
	payload := struct {
		Action models.RequestType `json:"action"`
	}{}

	if err := c.BindJSON(&payload); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if !payload.Action.Valid() {
		tc.BadRequestResponse(c, "mandatory fields has been missed")
		return
	}
	if payload.Action != models.Accept && payload.Action != models.Reject {
		tc.BadRequestResponse(c, "action is not supported")
	}
	if err := tc.teamService.ChangeMembershipStatus(teamID, user.Sub, memberID, payload.Action == models.Accept); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "changing membership status has failed", err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

func (tc *TeamController) HandleTeamMemberList(c *gin.Context) {
	if user := tc.userFromContext(c); user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	members, err := tc.teamService.GetTeamMembers(teamID)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "fetching team member has failed", err.Error())
		return
	}

	SuccessResponseV2(c, members)
}

func (tc *TeamController) HandleTransfer(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	payload := struct {
		Action models.TransferType `json:"action" binding:"required"`
		UserID string              `json:"userId"`
		Amount string              `json:"amount" binding:"required"`
	}{}

	if err := c.BindJSON(&payload); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if payload.Action == models.Distribute && len(payload.UserID) == 0 {
		tc.BadRequestResponse(c, "user_id need to be provided")
		return
	}
	team, err := tc.teamService.GetTeam(teamID)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "fetching team has failed", err.Error())
		return
	}
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		tc.AuthErrorResponse(c, "token is not provided")
		return
	}

	spwnAmount := ConvertEthToWei(payload.Amount)
	var (
		privateKey string
		dstAddress common.Address
	)

	switch payload.Action {
	case models.Fund:
		privateKey, err = hdkey.GeneratePrivateKeyFromUUID(user.Sub)
		if err != nil {
			tc.AuthErrorResponse(c, "fail to authorize the user ")
			return
		}
		dstAddress = common.HexToAddress(team.PublicKey)
	case models.Distribute:
		privateKey, err = hdkey.GeneratePrivateKeyFromUUID(team.ID)
		if err != nil {
			tc.AuthErrorResponse(c, "fail to authorize the user ")
			return
		}
		dstAddress = common.HexToAddress(user.PublicAddress)
	default:
		tc.BadRequestResponse(c, "action is not supported")
		return
	}

	tx, err := tc.bsc.Fund(privateKey, dstAddress, spwnAmount)
	if err != nil {
		tc.log.Error("error in funding: ", err)
		InternalErrorResponseV2(c, "funding team/individual transaction failed", err.Error())
		return
	}
	go func() {
		if err := tc.teamService.TransferLog(teamID, user.Sub, payload.UserID, payload.Amount, payload.Action); err != nil {
			tc.log.Error("error in insert into transfer log: ", err)
		}
	}()

	SuccessResponseV2(c, tx)

}
func (tc *TeamController) HandleTeamAvatarUpload(c *gin.Context) {
	user := tc.userFromContext(c)
	if user == nil {
		tc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	teamID := c.Param("team_id")
	if len(teamID) == 0 {
		BadRequestResponseV2(c, "team id is empty", "")
		return
	}
	maxSize := int64(1024000)
	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		tc.log.Error("Image too large. Max Size: 1M ", err)
		InternalErrorResponseV2(c, "Image too large. Max Size: 1M", err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("profile_picture")
	if err != nil {
		tc.log.Error("Could not get uploaded file: ", err)
		InternalErrorResponseV2(c, "Could not get uploaded file", err.Error())
		return
	}
	defer file.Close()
	tc.log.Info("fileHeader: ", fileHeader.Filename)
	awsConfig := tc.conf.AwsConfig()
	bucketname := awsConfig.S3BucketName
	tempAvatarUrl, _ := tc.S3UploadClient.HandleS3TeamAvatarUpload(bucketname, user.PublicAddress, file)
	avatarUrl := "https://" + bucketname + ".s3.amazonaws.com/" + tempAvatarUrl

	if err := tc.teamService.UpdateAvatar(teamID, avatarUrl); err != nil {
		tc.log.Error("error updating user avatar", err)
		if errors.Is(err, models.ErrRecordNotFound) {
			tc.NotFoundResponse(c, "the team not founded")
			return
		}
		InternalErrorResponseV2(c, "updating user avatar fails", err.Error())
		return
	}
	CustomSuccessResponseV2(c, nil, http.StatusNoContent)
}

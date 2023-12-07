/*

 */

package controllers

import (
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	BaseController
}

func NewGameController(bc *BaseController) *GameController {

	return &GameController{
		BaseController{
			Name: "game",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		},
	}
}

type FormGameTypeCreate struct {
	GameType string `json:"gameType"`
}

func (game *GameController) HandleGameTypeCreate(c *gin.Context) {

	form := FormGameTypeCreate{}

	if err := c.ShouldBindJSON(&form); err != nil {
		game.BadRequestResponse(c, err.Error())
		return
	}

	if form.GameType == "" {
		game.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	gameTypeSearch, err := game.DB.GetSpecificGameType(form.GameType)
	if err != nil {
		game.log.Error("error search game type from db", err)
		game.DBErrorResponse(c, err.Error())
		return
	}
	if len(gameTypeSearch) != 0 {
		game.BadRequestResponse(c, "game type exist")
		return
	}

	err = game.DB.CreateGameType(form.GameType)
	if err != nil {
		game.log.Error("error create new game type to db", err)
		game.DBErrorResponse(c, err.Error())
		return
	}

	game.SuccessResponse(c, "game type created", nil)
}

func (game *GameController) HandleGameTypeGet(c *gin.Context) {

	gameType, err := game.DB.GetGameType()
	if err != nil {
		game.log.Error("error get game type from db", err)
		game.DBErrorResponse(c, err.Error())
		return
	}

	game.SuccessResponse(c, "", gameType)
}

type FormGameSubTypeCreate struct {
	GameType    string `json:"gameType"`
	GameSubType string `json:"gameSubType"`
	TeamSize    int    `json:"teamSize"`
}

func (game *GameController) HandleGameSubTypeCreate(c *gin.Context) {

	form := FormGameSubTypeCreate{}

	if err := c.ShouldBindJSON(&form); err != nil {
		game.BadRequestResponse(c, err.Error())
		return
	}

	if form.GameType == "" || form.GameSubType == "" || form.TeamSize == 0 {
		game.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	gameSubTypeSearch, err := game.DB.GetSpecificGameSubType(form.GameType, form.GameSubType)
	if err != nil {
		game.log.Error("error search game sub type from db", err)
		game.DBErrorResponse(c, err.Error())
		return
	}
	if len(gameSubTypeSearch) != 0 {
		game.BadRequestResponse(c, "game sub type exist")
		return
	}

	err = game.DB.CreateGameSubType(form.GameType, form.GameSubType, form.TeamSize)
	if err != nil {
		game.log.Error("error create new game type to db", err)
		game.DBErrorResponse(c, err.Error())
		return
	}

	game.SuccessResponse(c, "game sub type created", nil)
}

func (game *GameController) HandleGameSubTypeGet(c *gin.Context) {
	gameType := c.Param("gameType")
	gameSubType, err := game.DB.GetGameSubType(gameType)
	if err != nil {
		game.log.Error("error get game type from db", err)
		game.DBErrorResponse(c, err.Error())
		return
	}

	game.SuccessResponse(c, "", gameSubType)
}

func (u *UserController) HandleGetUserGames(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	gamePlatformPlayerIDs, err := u.DB.GetUserGames(user.Sub)
	if err != nil {
		u.log.Error("error get games from db", err)
		InternalErrorResponseV2(c, "error get user games from db", err.Error())
		return
	}

	SuccessResponseV2(c, gamePlatformPlayerIDs)
}

func (u *UserController) HandlePutUserGame(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := struct {
		GameName string                `json:"gameName"`
		Platform models.PlatformDetail `json:"platform"`
	}{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	gameName := form.GameName
	platformName := form.Platform.PlatformName
	playerId := form.Platform.PlatformName

	if gameName == "" {
		BadRequestResponseV2(c, "game name missing from input", "")
		return
	}
	if platformName == "" {
		BadRequestResponseV2(c, "platform name missing from input", "")
		return
	}
	if playerId == "" {
		BadRequestResponseV2(c, "platform player ID missing from input", "")
		return
	}

	switch gameName {
	//TODO: move these hard codes to DB or Config
	case "CODMWBR":
		if platformName != "XBOX" && platformName != "PSN" && platformName != "BATTLENET" {
			BadRequestResponseV2(c, "unsupported Platform for this game type", "")
			return
		}
	case "APEX_LEGENDS":
		if platformName != "XBOX" && platformName != "PSN" && platformName != "STEAM" && platformName != "ORIGIN" && platformName != "NINTENDO" {
			BadRequestResponseV2(c, "unsupported Platform for this game type", "")
			return
		}
	default:
		BadRequestResponseV2(c, "unsupported game in challenge mode", "")
		return
	}

	games, err := u.DB.UpdateGamePlatformPlayerID(user.Sub, form.GameName, form.Platform.PlatformName, form.Platform.ID)
	if err != nil {
		u.log.Error("error update user game: ", err)
		InternalErrorResponseV2(c, "cannot update user game", err.Error())
		return
	}

	SuccessResponseV2(c, games)
}

func (u *UserController) HandlePutDisconnectUserGame(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := struct {
		GameName string                `json:"gameName"`
		Platform models.PlatformDetail `json:"platform"`
	}{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	switch form.GameName {
	case "CODMWBR":
		if form.Platform.PlatformName == "BATTLENET" {
			err := u.DB.DisconnectSpecificGame(user.Sub, form.GameName, form.Platform.PlatformName)
			if err != nil {
				u.log.Error("error in DisconnectSpecificGame: ", err)
				InternalErrorResponseV2(c, "error in DisconnectSpecificGame", err.Error())
				return
			}

			games, err := u.DB.GetUserGames(user.Sub)
			if err != nil {
				u.log.Error("error get games from db", err)
				InternalErrorResponseV2(c, "cannot get user games", err.Error())
				return
			}

			SuccessResponseV2(c, games)
		} else {
			BadRequestResponseV2(c, "disconnection is not supported for this platform", "")
			return
		}
	default:
		BadRequestResponseV2(c, "disconnection is not supported for this game", "")
		return
	}
}

func (u *UserController) HandlePutCheckUsername(c *gin.Context) {
	user := u.userFromContext(c)
	if user == nil {
		u.log.Error("user not found")
		InternalErrorResponseV2(c, "user not found", "")
		return
	}

	form := struct {
		GameName string                `json:"gameName"`
		Platform models.PlatformDetail `json:"platform"`
	}{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	used, err := u.DB.CheckPlatformUsernameUsed(user.Sub, form.GameName, form.Platform.PlatformName, form.Platform.ID)
	if err != nil {
		u.log.Error("error in CheckPlatformUsernameUsed: ", err)
		InternalErrorResponseV2(c, "error in CheckPlatformUsernameUsed", err.Error())
		return
	}

	SuccessResponseV2(c, used)
}

func (u *UserController) HandleReplaceUsername(c *gin.Context) {
	form := struct {
		GameName    string                `json:"gameName"`
		Platform    models.PlatformDetail `json:"platform"`
		UserAccount models.UserAccount    `json:"user"`
	}{}
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequestResponseV2(c, "error in input JSON format", err.Error())
		return
	}

	result, err := u.DB.UpdateGamePlatformPlayerID(form.UserAccount.Sub, form.GameName, form.Platform.PlatformName, form.Platform.ID)
	if err != nil {
		u.log.Error("error in UpdateGamePlatformPlayerID: ", err)
		InternalErrorResponseV2(c, "error in UpdateGamePlatformPlayerID", err.Error())
		return
	}

	SuccessResponseV2(c, result)
}

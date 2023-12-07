/*

 */

package controllers

import (
	"fmt"
	"net/http"

	"github.com/bitspawngg/bitspawn-api/config"
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	MsgPong           = "pong"
	PER_PAGE          = 15
	PER_PAGE_LONG     = 50
	CLUB_NUMBER       = 12
	CLAN_NUMBER       = 30
	LAST_MATCH_NUMBER = 5
	SQS_MAX_MESSAGES  = 2
)

// BaseController
type BaseController struct {
	Name string
	DB   *models.DB

	log  *logrus.Entry
	conf config.Config
}

func NewBaseController(name string, db *models.DB, log *logrus.Logger, conf config.Config) *BaseController {
	return &BaseController{
		name,
		db,
		log.WithField("controller", name),
		conf,
	}
}

func SuccessResponseV2(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"response": data,
		},
	)
}

func (bc *BaseController) SuccessResponse(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	c.JSON(
		http.StatusOK,
		Response{
			Msg:  msg,
			Body: data,
		},
	)
}

func BadRequestResponseV2(c *gin.Context, msg string, detail string) {
	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"error": ErrorResponse{
				Code:    14000,
				Message: msg,
				Detail:  "(14000) " + msg + ": " + detail,
			},
		},
	)
}

func InternalErrorResponseV2(c *gin.Context, msg string, detail string) {
	c.JSON(
		http.StatusInternalServerError,
		gin.H{
			"error": ErrorResponse{
				Code:    15000,
				Message: msg,
				Detail:  "(15000) " + msg + ": " + detail,
			},
		},
	)
}

func AuthErrorResponseV2(c *gin.Context, msg string, detail string) {
	c.JSON(
		http.StatusUnauthorized,
		gin.H{
			"error": ErrorResponse{
				Code:    14000,
				Message: msg,
				Detail:  "(14000) " + msg + ": " + detail,
			},
		},
	)
}

func CustomSuccessResponseV2(c *gin.Context, data interface{}, code int) {
	c.JSON(
		code,
		gin.H{
			"response": data,
		},
	)
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (bc *BaseController) BadRequestResponse(c *gin.Context, msg string) {
	c.JSON(
		http.StatusBadRequest,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}

func (bc *BaseController) NotFoundResponse(c *gin.Context, msg string) {
	c.JSON(
		http.StatusNotFound,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}

func (bc *BaseController) AuthErrorResponse(c *gin.Context, msg string) {
	c.JSON(
		http.StatusUnauthorized,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}

func (bc *BaseController) DBErrorResponse(c *gin.Context, msg string) {
	c.JSON(
		http.StatusInternalServerError,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}

func (bc *BaseController) InternalErrorResponse(c *gin.Context, msg string) {
	c.JSON(
		http.StatusInternalServerError,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}

func (bc *BaseController) BlockchainErrorResponse(c *gin.Context, msg string) {
	c.JSON(
		http.StatusOK,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}

func (bc *BaseController) ToornamentErrorResponse(c *gin.Context, msg string) {
	c.JSON(
		http.StatusFailedDependency,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}
func (bc *BaseController) CustomFailedResponse(c *gin.Context, msg string, code int) {
	c.JSON(
		code,
		Response{
			"failure",
			msg,
			nil,
		},
	)
}

// HandlePing handles the ping request for health check.
func (bc *BaseController) HandlePing(c *gin.Context) {

	bc.log.Debug("handling ping...")

	bc.SuccessResponse(c, "", MsgPong)
}

func (bc *BaseController) userFromContext(c *gin.Context) *models.UserAccount {
	u, ok := c.Get("user")
	if !ok {
		return nil
	}

	fmt.Printf("%#v \n", u)

	user, ok := u.(*models.UserAccount)
	if !ok {
		return nil
	}

	return user
}

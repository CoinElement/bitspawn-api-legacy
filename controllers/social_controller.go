package controllers

import (
	"errors"
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services/social"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SocialController struct {
	*BaseController
	socialService *social.SocialService
}

func NewSocialController(bc *BaseController) *SocialController {
	return &SocialController{
		&BaseController{
			Name: "team",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		}, social.NewSocialService(bc.DB, bc.log),
	}
}

func (sc *SocialController) HandleSocialLinkRegister(c *gin.Context) {
	user := sc.userFromContext(c)
	if user == nil {
		sc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	social := models.SocialLink{}

	if err := c.BindJSON(&social); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if !social.SocialType.Valid() {
		BadRequestResponseV2(c, "mandatory fields has been missed", "")
		return
	}
	social.UserID = user.Sub
	if err := sc.socialService.CreateSocialLink(&social); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			sc.NotFoundResponse(c, "tsocial link of this type is not found")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			BadRequestResponseV2(c, "duplicated request", err.Error())
			return
		}
		InternalErrorResponseV2(c, "creating social link has failed", err.Error())
		return
	}

	CustomSuccessResponseV2(c, social, http.StatusCreated)
}

func (sc *SocialController) HandleSocialLinkUnRegister(c *gin.Context) {
	user := sc.userFromContext(c)
	if user == nil {
		sc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	socialType := c.Param("social_type")
	if len(socialType) == 0 {
		BadRequestResponseV2(c, "social type id is empty", "")
		return
	}
	st := models.SocialType(socialType)
	if !st.Valid() {
		BadRequestResponseV2(c, "mandatory fields has been missed", "")
		return
	}
	if err := sc.socialService.DeleteSocialLink(user.Sub, st); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			sc.NotFoundResponse(c, "social link of this type is not found")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			BadRequestResponseV2(c, "duplicated request", "")
			return
		}
		InternalErrorResponseV2(c, "creating social link has failed", err.Error())
		return
	}

	CustomSuccessResponseV2(c, nil, http.StatusOK)
}

func (sc *SocialController) HandleSocialLinkGet(c *gin.Context) {
	user := sc.userFromContext(c)
	if user == nil {
		sc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	socialType := c.Param("social_type")
	if len(socialType) == 0 {
		BadRequestResponseV2(c, "social type id is empty", "")
		return
	}
	st := models.SocialType(socialType)
	if !st.Valid() {
		BadRequestResponseV2(c, "mandatory fields has been missed", "")
		return
	}
	soiallink, err := sc.socialService.GetSocialLink(user.Sub, st)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			sc.NotFoundResponse(c, "social link of this type is not found")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			BadRequestResponseV2(c, "duplicated request", "")
			return
		}
		InternalErrorResponseV2(c, "creating social link has failed", err.Error())
		return
	}

	SuccessResponseV2(c, soiallink)
}

func (sc *SocialController) HandleSocialLinkGetAll(c *gin.Context) {
	user := sc.userFromContext(c)
	if user == nil {
		sc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}

	soiallinks, err := sc.socialService.GetAllSocialLink(user.Sub)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			sc.NotFoundResponse(c, "social link of this type is not found")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			BadRequestResponseV2(c, "duplicated request", "")
			return
		}
		InternalErrorResponseV2(c, "creating social link has failed", err.Error())
		return
	}

	SuccessResponseV2(c, soiallinks)
}

func (sc *SocialController) HandleSocialLinkUpdate(c *gin.Context) {
	user := sc.userFromContext(c)
	if user == nil {
		sc.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	socialType := c.Param("social_type")
	if len(socialType) == 0 {
		BadRequestResponseV2(c, "social type id is empty", "")
		return
	}
	st := models.SocialType(socialType)
	if !st.Valid() {
		BadRequestResponseV2(c, "mandatory fields has been missed", "")
		return
	}
	social := models.SocialLink{}
	social.UserID = user.Sub
	if err := c.BindJSON(&social); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	if err := sc.socialService.UpdateSocialLink(user.Sub, social.SocialID, st); err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			sc.NotFoundResponse(c, "social link of this type is not found")
			return
		}
		if errors.Is(err, models.ErrDuplicatedEntity) {
			BadRequestResponseV2(c, "duplicated request", "")
			return
		}
		InternalErrorResponseV2(c, "creating social link has failed", err.Error())
		return
	}

	CustomSuccessResponseV2(c, nil, http.StatusOK)
}

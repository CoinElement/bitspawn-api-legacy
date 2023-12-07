package controllers

import (
	"net/http"

	"github.com/bitspawngg/bitspawn-api/services/signURL"
	"github.com/gin-gonic/gin"
)

type PreSignedController struct {
	*BaseController
	signURLMngr *signURL.SignURLManager
}

func NewPreSignedController(bc *BaseController, signURLManager *signURL.SignURLManager) *PreSignedController {
	return &PreSignedController{
		&BaseController{
			Name: "signURL",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		}, signURLManager,
	}
}

func (ps *PreSignedController) HandleURLRegister(c *gin.Context) {
	user := ps.userFromContext(c)
	if user == nil {
		ps.log.Error("user not found")
		AuthErrorResponseV2(c, "user not found", "")
		return
	}
	payload := struct {
		BucketName string `json:"bucketName"`
		FileName   string `json:"fileName"`
	}{}

	if err := c.BindJSON(&payload); err != nil {
		BadRequestResponseV2(c, "unable to parse body", err.Error())
		return
	}
	url, headers, err := ps.signURLMngr.GetSignedURL(payload.FileName)
	if err != nil {
		InternalErrorResponseV2(c, "creating pre signed url failed", err.Error())
	}
	SuccessResponseV2(c, struct {
		URL     string      `json:"url"`
		Headers http.Header `json:"headers"`
	}{
		url, headers,
	})
}

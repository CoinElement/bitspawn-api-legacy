/*

 */

package controllers

import (
	"strings"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/bitspawngg/bitspawn-api/services/s3"
	"github.com/gin-gonic/gin"
)

type BadgeController struct {
	BaseController
	S3UploadClient *s3.S3UploadClient
}

func NewBadgeController(bc *BaseController, s3UploadClient *s3.S3UploadClient) *BadgeController {

	return &BadgeController{
		BaseController{
			Name: "badge",
			DB:   bc.DB,
			log:  bc.log,
			conf: bc.conf,
		},
		s3UploadClient,
	}
}

func (badge *BadgeController) HandleBadgeUpload(c *gin.Context) {
	maxSize := int64(1024000)
	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		badge.log.Error(err)
		badge.InternalErrorResponse(c, "Image too large. Max Size: 1M")
		return
	}
	file, fileHeader, err := c.Request.FormFile("badge_picture")
	if err != nil {
		badge.log.Error(err)
		badge.InternalErrorResponse(c, "Could not get uploaded file")
		return
	}
	defer file.Close()
	badge.log.Info("fileHeader: ", fileHeader.Filename)
	awsConfig := badge.conf.AwsConfig()
	bucketname := awsConfig.S3BucketName
	tempBadgeUrl, _ := badge.S3UploadClient.HandleS3BadgeUpload(bucketname, file)
	segments := strings.SplitN(tempBadgeUrl, "/", -1)
	badgeId := segments[len(segments)-1]

	badgeInfo, err := badge.DB.StorageBadgeUrl(tempBadgeUrl, badgeId)
	if err != nil {
		badge.log.Error("error create and storage badge", err)
		badge.InternalErrorResponse(c, "create and storage badge fails")
		return
	}

	badge.SuccessResponse(c, "", badgeInfo)
}

type ResponseBadges struct {
	BadgeData []models.Badge `json:"badgeData"`
}

func (badge *BadgeController) HandleAllBadgeGet(c *gin.Context) {
	badges, err := badge.DB.GetAllBadgeUrl()
	if err != nil {
		badge.log.Error("error getting badges", err)
		badge.DBErrorResponse(c, "error getting badges")
		return
	}

	badge.SuccessResponse(c, "", ResponseBadges{
		badges,
	})
}

func (badge *BadgeController) HandleSpecificBadgeGet(c *gin.Context) {

	badgeID := c.Param("badgeId")
	badgeInfo, err := badge.DB.GetMatchBadgeUrl(badgeID)
	if err != nil {
		badge.log.Error("error getting badges", err)
		badge.DBErrorResponse(c, "error getting badges")
		return
	}

	badge.SuccessResponse(c, "", badgeInfo)
}

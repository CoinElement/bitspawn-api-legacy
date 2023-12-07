/*

 */

package controllers

import (
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/gin-gonic/gin"
)

type ResponseMyChallenges struct {
	MyChallengesData []models.Challenge `json:"myChallengesData"`
}

func (cc *ChallengeController) HandleMyChallenges(c *gin.Context) {
	user := cc.userFromContext(c)
	if user == nil {
		cc.log.Error("user not found")
		cc.InternalErrorResponse(c, "user not found")
		return
	}

	myRegisteredChallenges, err := cc.DB.GetRegisteredChallenges(user.Sub)
	if err != nil {
		cc.log.Error("error getting my registered challenges", err)
		cc.DBErrorResponse(c, "error getting my registered challenges")
		return
	}

	cc.SuccessResponse(c, "", gin.H{
		"myRegisteredChallenges": myRegisteredChallenges,
	})
}

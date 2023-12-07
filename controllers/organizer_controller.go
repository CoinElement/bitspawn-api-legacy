package controllers

import (
	"github.com/gin-gonic/gin"
)

func (tc *TournamentController) HandleStartTournament(c *gin.Context) {
	err := tc.orgSvc.StartTournament()
	if err != nil {
		tc.log.Errorf("error in StartTournament: %v", err)
		tc.InternalErrorResponse(c, "error in StartTournament: "+err.Error())
		return
	}
	tc.SuccessResponse(c, "success", nil)
}

func (tc *TournamentController) HandlePrepareManualMatches(c *gin.Context) {
	err := tc.orgSvc.PrepareManualMatches()
	if err != nil {
		tc.log.Errorf("error in PrepareManualMatches: %v", err)
		tc.InternalErrorResponse(c, "error in PrepareManualMatches: "+err.Error())
		return
	}
	tc.SuccessResponse(c, "success", nil)
}

func (tc *TournamentController) HandlePrepareMatches(c *gin.Context) {
	err := tc.orgSvc.PrepareMatches()
	if err != nil {
		tc.log.Errorf("error in PrepareMatches: %v", err)
		tc.InternalErrorResponse(c, "error in PrepareMatches: "+err.Error())
		return
	}
	tc.SuccessResponse(c, "success", nil)
}

func (tc *TournamentController) HandlePrepareRematches(c *gin.Context) {
	err := tc.orgSvc.PrepareRematches()
	if err != nil {
		tc.log.Errorf("error in PrepareRematches: %v", err)
		tc.InternalErrorResponse(c, "error in PrepareRematches: "+err.Error())
		return
	}
	tc.SuccessResponse(c, "success", nil)
}

func (tc *TournamentController) HandlePrepareRoundOne(c *gin.Context) {
	err := tc.orgSvc.PrepareRoundOne()
	if err != nil {
		tc.log.Errorf("error in PrepareRoundOne: %v", err)
		tc.InternalErrorResponse(c, "error in PrepareRoundOne: "+err.Error())
		return
	}
	tc.SuccessResponse(c, "success", nil)
}

type FormCreateMatchSchedule struct {
	TournamentID string `json:"tournamentId"`
}

func (tc *TournamentController) HandleCreateMatchSchedule(c *gin.Context) {
	form := FormCreateMatchSchedule{}
	if err := c.ShouldBindJSON(&form); err != nil {
		tc.BadRequestResponse(c, err.Error())
		return
	}

	if form.TournamentID == "" {
		tc.BadRequestResponse(c, "missing mandatory input parameter")
		return
	}

	// TODO: input field validation

	err := tc.orgSvc.CreateMatchSchedule(form.TournamentID, true)
	if err != nil {
		tc.log.Errorf("error in CreateMatchSchedule: %v", err)
		tc.InternalErrorResponse(c, "error in CreateMatchSchedule: "+err.Error())
		return
	}
	tc.SuccessResponse(c, "success", nil)
}

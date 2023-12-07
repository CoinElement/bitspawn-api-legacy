package team

import (
	"errors"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/sirupsen/logrus"
)

// TeamService is used to handler team related request services
type TeamService struct {
	th  *models.TeamHandler
	log *logrus.Entry
}

// NewTeamService create a teamservice instance
func NewTeamService(db *models.DB, log *logrus.Entry) *TeamService {
	return &TeamService{
		th:  models.NewTeamHandler(db),
		log: log.WithField("service", "team"),
	}
}

func (ts *TeamService) CreateTeam(team *models.TeamMachine, creator string) error {
	return ts.th.CreateTeamMachine(team, creator)
}

func (ts *TeamService) CreateTeamMember(request *models.JoinRequestBody) error {
	switch request.Action {
	case models.Invite:
		return ts.th.InviteToJoin(request.TeamID, request.UserID)
	case models.Apply:
		return ts.th.AskToJoin(request.TeamID, request.UserID)
	default:
		return errors.New("action is not supported")
	}
}

func (ts *TeamService) UpdateTeam(teamID string, team *models.TeamMachine) error {
	return ts.th.UpdateTeam(teamID, team)
}

func (ts *TeamService) DeleteTeam(teamID string) error {
	return ts.th.DeleteTeamMachine(teamID)
}
func (ts *TeamService) List() ([]*models.TeamMachine, error) {
	return ts.th.FetchTeamMachinesAll()
}
func (ts *TeamService) GetTeamsByName(name string) ([]*models.TeamMachine, error) {
	return ts.th.FetchTeamMachinesByName(name)
}
func (ts *TeamService) GetTeam(teamID string) (*models.TeamMachine, error) {
	return ts.th.FetchTeamMachineByID(teamID)
}
func (ts *TeamService) DeleteTeamMember(teamID, sub string) error {
	return ts.th.DeleteTeamMember(teamID, sub)
}
func (ts *TeamService) GetTeamMembers(teamID string) ([]*models.TeamMember, error) {
	team, err := ts.th.FetchTeamMachineByID(teamID)
	if err != nil {
		return nil, err
	}
	return team.Members, nil
}
func (ts *TeamService) GetTeamMember(teamID, sub string) (*models.TeamMember, error) {
	return ts.th.FetchTeamMember(teamID, sub)
}
func (ts *TeamService) ChangeRole(teamID, sub, requestedSub string, role models.Role) error {
	return ts.th.ChangeRole(teamID, sub, requestedSub, role)
}
func (ts *TeamService) UpdateAvatar(teamID, url string) error {
	return ts.th.UpdateAvatar(teamID, url)
}
func (ts *TeamService) ChangeMembershipStatus(teamID, requestedSub, sub string, accepted bool) error {
	return ts.th.ApproveMembership(teamID, requestedSub, sub, accepted)
}

func (ts *TeamService) TransferLog(teamID, from_sub, to_sub string, amount string, action models.TransferType) error {
	return ts.th.InsertTransferLog(teamID, from_sub, to_sub, amount, action)
}

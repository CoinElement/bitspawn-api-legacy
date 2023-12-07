package models

import (
	"time"
)

// TransferType is used for defining type of tranfer
type TransferType string

const (
	Fund       TransferType = "FUND"
	Distribute TransferType = "DISTRIBUTE"
)

func (t TransferType) Valid() bool {
	if t != Fund && t != Distribute {
		return false
	}
	return true
}

// Role is define which kind of membership user has
type Role string

const (
	owner   Role = "OWNER"
	manager Role = "MANAGER"
	member  Role = "MEMBER"
)

func (r Role) Valid() bool {
	if r != owner && r != manager && r != member {
		return false
	}
	return true
}

// Status is define the status of the request to join team
type Status string

const (
	Pending  Status = "PENDING"
	Approved Status = "APPROVED"
	Rejected Status = "REJECTED"
)

func (s Status) Status() bool {
	if s != Pending && s != Approved && s != Rejected {
		return false
	}
	return true
}

type RequestType string

const (
	Invite RequestType = "INVITE" //invite by a member of the team
	Apply  RequestType = "APPLY"  // asked to join by the user itself
	Accept RequestType = "ACCEPT" // accept the request to join the team
	Reject RequestType = "REJECT" // reject the request to join the ream , // also it can be called by delete endpoint
)

func (r RequestType) Valid() bool {
	if r != Invite && r != Apply && r != Accept && r != Reject {
		return false
	}
	return true
}

// TeamMemberHandler is used to handle membership data handler and maniupulate team data in database
type TeamMemberHandler struct {
	*DB
}

// request body of AskToJoin and invite
type JoinRequestBody struct {
	TeamID string      `json:"teamId"`
	Action RequestType `json:"action"`
	UserID string      `json:"userId"` //equal to sub
}

// NewTeamHandler is used to create a new TeamMemberHandler
func NewTeamMemberHandler(db *DB) *TeamMemberHandler {
	return &TeamMemberHandler{db}
}

// TeamMember is used to store and retrieve data from and to database
type TeamMember struct {
	TeamID      string       `json:"teamId"  gorm:"primaryKey;"`
	Sub         string       `json:"sub"  gorm:"primaryKey;"`
	User        *UserAccount `gorm:"-" json:"user"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Role        Role         `json:"role"`
	Status      Status       `json:"status"`
	RequestType RequestType  `json:"requestType"`
}

// CreateTeamMember is used to add a user as member of the teams
func (tm *TeamMemberHandler) CreateTeamMember(teamID, sub string, role Role, requestType RequestType) error {
	// default value for role for new member should be 'member'
	if _, err := tm.GetUser(sub); err != nil {
		return err
	}
	status := Pending
	if role == owner {
		status = Approved
	}
	one := &TeamMember{
		TeamID:      teamID,
		Sub:         sub,
		Role:        role,
		RequestType: requestType,
		Status:      status, // default value is pending. Acquired action needed to make the requested user a member
	}
	return convError(tm.Create(one).Error)
}

// CreateTeamMember is used to add a user as member of the teams
func (tm *TeamMemberHandler) ChangeRole(teamID, sub string, role Role) error {
	team, err := tm.FindTeamMember(teamID, sub)
	if err != nil {
		return err
	}
	return tm.Model(team).Update("role", role).Error
}

// Approve is used to change the status from pending to approve
func (tm *TeamMemberHandler) Approve(teamID, sub string) error {
	team, err := tm.FindTeamMember(teamID, sub)
	if err != nil {
		return err
	}
	return convError(tm.Model(team).Update("status", Approved).Error)
}

// Reject is used to change the status from pending to reject
func (tm *TeamMemberHandler) Reject(teamID, sub string) error {
	team, err := tm.FindTeamMember(teamID, sub)
	if err != nil {
		return err
	}
	return convError(tm.Model(team).Update("status", Rejected).Error)
}

// DeleteTeamMember is used to remove a user as member of the teams
func (tm *TeamMemberHandler) DeleteTeamMember(teamID, sub string) error {
	team, err := tm.FindTeamMember(teamID, sub)
	if err != nil {
		return err
	}
	return convError(tm.Delete(team).Error)
}

// DeleteTeamMembers is used to delete all of the member of the given team id
func (tm *TeamMemberHandler) DeleteTeamMembers(teamID string) error {
	return convError(tm.Where(&TeamMember{TeamID: teamID}).Delete(&TeamMember{}).Error)
}

// FindTeamMember is used to find membership according to the given team_id and user_sub
func (tm *TeamMemberHandler) FindTeamMember(teamID, sub string) (*TeamMember, error) {
	team := &TeamMember{}
	if err := tm.Where(&TeamMember{TeamID: teamID, Sub: sub}).First(&team).Error; err != nil {
		return nil, convError(err)
	}
	user, err := tm.GetUser(sub)
	if err != nil {
		return nil, err
	}
	team.User = user
	return team, nil
}

// FindTeamMember is used to find membership according to the given team_id and user_sub
func (tm *TeamMemberHandler) FindTeamMembers(teamID string) ([]*TeamMember, error) {
	teamMembers := []*TeamMember{}
	if err := tm.Where(&TeamMember{TeamID: teamID}).Find(&teamMembers).Order("created_at asc").Error; err != nil {
		return nil, convError(err)
	}
	for _, teamMember := range teamMembers {
		user := &UserAccount{}
		if err := tm.Where(&UserAccount{Sub: teamMember.Sub}).First(user).Error; err != nil {
			return nil, convError(err)
		}
		teamMember.User = user
	}
	return teamMembers, nil
}

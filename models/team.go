package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/bitspawngg/bitspawn-api/services/hdkey"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var (
	ErrDuplicatedEntity = errors.New("Duplicate key found.")
	ErrRecordNotFound   = errors.New("Record not found.")
	ErrNotAllowed       = errors.New("Operation is not allowed.")
	ErrWrongValues      = errors.New("provided values are not valid.")
)

// TeamHandler is used to handle team data handler and maniupulate team data in database
type TeamHandler struct {
	tmh *TeamMemberHandler
	db  *DB
}

// NewTeamHandler is used to create a new TeamHandler
func NewTeamHandler(db *DB) *TeamHandler {
	return &TeamHandler{db: db, tmh: NewTeamMemberHandler(db)}
}

type Publicity string

const (
	Open       Publicity = "OPEN"
	InviteOnly Publicity = "INVITE_ONLY"
)

func (p Publicity) Valid() bool {
	if p != Open && p != InviteOnly {
		return false
	}
	return true
}

type GenrePreferred string

const (
	BattleRoyale GenrePreferred = "BATTLE ROYALE"
	Sports       GenrePreferred = "SPORTS"
	FPS          GenrePreferred = "FPS"
)

func (g GenrePreferred) Valid() bool {
	if g != BattleRoyale && g != Sports && g != FPS {
		return false
	}
	return true
}

// TeamMachine is used to store and retrieve data from and to database
type TeamMachine struct {
	ID             string         `json:"id,omitempty"  gorm:"primaryKey;"`
	Name           string         `json:"name" gorm:"uniqueIndex"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
	Publicity      Publicity      `json:"publicity" gorm:"type:text"`
	GenrePreferred GenrePreferred `json:"genrePreferred" gorm:"type:text"`
	AvatarURL      string         `json:"avatarURL,omitempty" gorm:"type:text"`
	PublicKey      string         `json:"publicKey,omitempty" gorm:"type:text"`
	Members        []*TeamMember  `gorm:"-" json:"members,omitempty"`
}

var (
	ErrTeamNotOpen = errors.New("the team is not open for new members")
)

// ManipulateStatus is used to approve or reject the membership of the given user
func (th *TeamHandler) ApproveMembership(teamID, requestedSub, sub string, accepted bool) error {
	requestedMember, err := th.tmh.FindTeamMember(teamID, requestedSub)
	if err != nil {
		return err
	}
	if requestedMember.Role != manager && requestedMember.Role != owner {
		return ErrNotAllowed
	}
	if accepted {
		return th.tmh.Approve(teamID, sub)
	}
	return th.tmh.Reject(teamID, sub)
}

// AskToJoin is used to create a pending member, need to be approved by owner/managers
// the team should be publicly open to join
func (th *TeamHandler) AskToJoin(teamID, sub string) error {
	team, err := th.FetchTeamMachineByID(teamID)
	if err != nil {
		return err
	}
	if team.Publicity != Open {
		return ErrTeamNotOpen
	}
	return th.tmh.CreateTeamMember(teamID, sub, member, Apply)
}

// InviteToJoin is used to create a pending member, need to be approved by the user
func (th *TeamHandler) InviteToJoin(teamID, sub string) error {
	if _, err := th.FetchTeamMachineByID(teamID); err != nil {
		return err
	}
	return th.tmh.CreateTeamMember(teamID, sub, member, Invite)
}

// CreateTeamMachine is used to create a team item in database
func (th *TeamHandler) ChangeRole(teamID, sub, requestedSub string, role Role) error {
	requestedMember, err := th.tmh.FindTeamMember(teamID, requestedSub)
	if err != nil {
		return err
	}
	switch role {
	case member:
		if requestedMember.Role != owner && requestedMember.Role != manager {
			return ErrNotAllowed
		}
	case manager:
		if requestedMember.Role != owner {
			return ErrNotAllowed
		}
	case owner:
		if requestedMember.Role != owner {
			return ErrNotAllowed
		}
		// one owner is allowed , it should be replaced with current owner
		// current one will be manager
		if err := th.tmh.ChangeRole(teamID, requestedSub, manager); err != nil {
			return err
		}
	default:
		return ErrNotAllowed
	}
	return th.tmh.ChangeRole(teamID, sub, role)
}

// CreateTeamMachine is used to create a team item in database
func (th *TeamHandler) CreateTeamMachine(team *TeamMachine, creator string) error {
	team.ID = uuid.NewV4().String()
	var err error
	if team.PublicKey, err = hdkey.GeneratePrivateKeyFromUUID(team.ID); err != nil {
		return err
	}
	if err := th.tmh.CreateTeamMember(team.ID, creator, owner, Invite); err != nil {
		return err
	}

	if err := th.db.Create(team).Error; err != nil {
		return convError(err)
	}
	return nil
}

// FetchTeamMachineByID fetches the team based on the given id
func (th *TeamHandler) FetchTeamMachineByID(id string) (*TeamMachine, error) {
	team := &TeamMachine{}
	if err := th.db.Where(&TeamMachine{ID: id}).First(&team).Error; err != nil {
		return nil, convError(err)
	}

	members, err := th.tmh.FindTeamMembers(id)
	if err != nil {
		return nil, err
	}
	team.Members = members
	return team, nil
}

// FetchTeamMachineByID fetches the team based on the given id
func (th *TeamHandler) FetchTeamMember(teamID, sub string) (*TeamMember, error) {
	return th.tmh.FindTeamMember(teamID, sub)
}

// FetchTeamMachineByID fetchs team based on the given name
func (th *TeamHandler) FetchTeamMachinesByName(name string) ([]*TeamMachine, error) {
	teams := []*TeamMachine{}
	if err := th.db.Where("name like ? ", "%"+name+"%").Find(&teams).Error; err != nil {
		return nil, convError(err)
	}
	return teams, nil
}

// FetchTeamMachineByID fetchs team based on the given name
func (th *TeamHandler) FetchTeamMachinesAll() ([]*TeamMachine, error) {
	teams := []*TeamMachine{}
	if err := th.db.Where("name!='1'").Find(&teams).Error; err != nil {
		return nil, convError(err)
	}
	return teams, nil
}

// UpdateTeam updates team based on the given info
func (th *TeamHandler) UpdateTeam(id string, team *TeamMachine) error {
	err := th.db.Model(&TeamMachine{}).Where("id = ?", id).Updates(map[string]interface{}{"name": team.Name, "genre_preferred": team.GenrePreferred, "publicity": team.Publicity}).Error
	return convError(err)
}

// FetchTeamMachine fetchs team based on the given id
func (th *TeamHandler) DeleteTeamMachine(id string) error {
	team := &TeamMachine{}
	if err := th.db.Where(&TeamMachine{ID: id}).First(&team).Error; err != nil {
		return err
	}
	// it should clean all members created in the team
	if err := th.tmh.DeleteTeamMembers(id); err != nil {
		return err
	}
	return convError(th.db.Delete(team).Error)
}

func (th *TeamHandler) DeleteTeamMember(teamID, sub string) error {
	return th.tmh.DeleteTeamMember(teamID, sub)
}

func (th *TeamHandler) UpdateAvatar(teamID, url string) error {
	team, err := th.FetchTeamMachineByID(teamID)
	if err != nil {
		return err
	}
	return convError(th.db.Model(team).Update("avatar_url", url).Error)
}

type TransferLog struct {
	ID        string       `json:"id,omitempty"  gorm:"primaryKey;"`
	TeamID    string       `json:"team_id,omitempty"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	TrxType   TransferType `json:"trx_type" gorm:"type:text"`
	From      string       `json:"from" gorm:"type:text"`
	To        string       `json:"to" gorm:"type:text"`
	Amount    string       `json:"amount"`
}

// InsertTrx is used to insert a transaction into database
func (th *TeamHandler) InsertTransferLog(teamID, from, to, amount string, trxType TransferType) error {
	one := &TransferLog{
		ID:      uuid.NewV4().String(),
		TeamID:  teamID,
		From:    from,
		To:      to,
		Amount:  amount,
		TrxType: trxType,
	}
	return convError(th.db.Create(one).Error)
}

func IsDuplicateKeyError(err error) bool {
	duplicate := regexp.MustCompile(`\(SQLSTATE 23505\)$`)
	return duplicate.MatchString(err.Error())
}

// convError will convert gorm error to our custom error
func convError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrRecordNotFound
	case IsDuplicateKeyError(err):
		return ErrDuplicatedEntity
	default:
		return err
	}
}

/*

 */

package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TournamentRequest struct {
	RequestID    string           `gorm:"primaryKey" json:"requestId"`
	TournamentID string           `gorm:"type:uuid" json:"tournamentId"`
	Sub          string           `json:"-"`
	UserAccount  *UserAccount     `gorm:"-" json:"-"`
	User         *UserInfo        `gorm:"-" json:"user"`
	CreatedAt    time.Time        `json:"createdAt"`
	UpdatedAt    time.Time        `json:"updatedAt"`
	Category     TRequestCategory `json:"-"`
}

type TRequestCategory string

const (
	TApplication TRequestCategory = "APPLICATION"
	TInvite      TRequestCategory = "INVITE"
)

func (db *DB) CreateTournamentApplication(tournamentId, sub string) error {
	requestId := uuid.NewV4().String()
	application := TournamentRequest{RequestID: requestId, TournamentID: tournamentId, Sub: sub, Category: TApplication}
	return db.Create(&application).Error
}

func (db *DB) CreateTournamentInvite(tournamentId, sub string) error {
	requestId := uuid.NewV4().String()
	invite := TournamentRequest{RequestID: requestId, TournamentID: tournamentId, Sub: sub, Category: TInvite}
	return db.Create(&invite).Error
}

func (db *DB) FetchTournamentApplication(requestId string) (*TournamentRequest, error) {
	application := TournamentRequest{}
	err := db.Where(&TournamentRequest{RequestID: requestId, Category: TApplication}).
		First(&application).Error
	if err != nil {
		return nil, err
	}
	userAccount := UserAccount{}
	err = db.Where(&UserAccount{Sub: application.Sub}).First(&userAccount).Error
	if err != nil {
		return nil, err
	}
	application.UserAccount = &userAccount
	return &application, nil
}

func (db *DB) FetchTournamentInvite(requestId string) (*TournamentRequest, error) {
	invite := TournamentRequest{}
	err := db.Where(&TournamentRequest{RequestID: requestId, Category: TInvite}).
		First(&invite).Error
	if err != nil {
		return nil, err
	}
	userAccount := UserAccount{}
	err = db.Where(&UserAccount{Sub: invite.Sub}).First(&userAccount).Error
	if err != nil {
		return nil, err
	}
	invite.UserAccount = &userAccount
	return &invite, nil
}

func (db *DB) FindTournamentApplication(tournamentId, sub string) ([]TournamentRequest, error) {
	applications := []TournamentRequest{}
	err := db.Where(&TournamentRequest{TournamentID: tournamentId, Sub: sub, Category: TApplication}).
		Find(&applications).Error
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (db *DB) FindTournamentInvite(tournamentId, sub string) ([]TournamentRequest, error) {
	invites := []TournamentRequest{}
	err := db.Where(&TournamentRequest{TournamentID: tournamentId, Sub: sub, Category: TInvite}).
		Find(&invites).Error
	if err != nil {
		return nil, err
	}
	return invites, err
}

func (db *DB) ListActiveApplicationsByUser(user *UserAccount) ([]TournamentRequest, error) {
	applications := []TournamentRequest{}
	err := db.Model(&TournamentRequest{}).
		Joins("Left Join tournament_data on tournament_requests.tournament_id = tournament_data.tournament_id").
		Where("sub = ?", user.Sub).
		Where("category = ?", TApplication).
		Where("tournament_data.status = ?", "REGISTRATION").
		Find(&applications).Error
	if err != nil {
		return nil, err
	}
	for i := range applications {
		userInfo := UserInfo{
			AvatarUrl:   user.AvatarUrl,
			DisplayName: user.DisplayName,
			Sub:         user.Sub,
			Username:    user.Username,
		}
		applications[i].User = &userInfo
	}
	return applications, nil
}

func (db *DB) ListActiveInvitesByUser(user *UserAccount) ([]TournamentRequest, error) {
	invites := []TournamentRequest{}
	err := db.Model(&TournamentRequest{}).
		Joins("Left Join tournament_data on tournament_requests.tournament_id = tournament_data.tournament_id").
		Where("sub = ?", user.Sub).
		Where("category = ?", TInvite).
		Where("tournament_data.status = ?", "REGISTRATION").
		Find(&invites).Error
	if err != nil {
		return nil, err
	}
	for i := range invites {
		userInfo := UserInfo{
			AvatarUrl:   user.AvatarUrl,
			DisplayName: user.DisplayName,
			Sub:         user.Sub,
			Username:    user.Username,
		}
		invites[i].User = &userInfo
	}
	return invites, nil
}

func (db *DB) ListApplicationByTournament(tournamentId string) ([]TournamentRequest, error) {
	rows, err := db.Model(&TournamentRequest{}).
		Joins("Left Join user_accounts on tournament_requests.sub = user_accounts.sub").
		Where("tournament_id = ?", tournamentId).
		Where("category = ?", TApplication).
		Select("tournament_requests.*, user_accounts.*").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	applications := []TournamentRequest{}
	for rows.Next() {
		var application TournamentRequest
		var user UserInfo
		_ = db.ScanRows(rows, &application)
		_ = db.ScanRows(rows, &user)
		application.User = &user
		applications = append(applications, application)
	}
	return applications, nil
}

func (db *DB) ListInviteByTournament(tournamentId string) ([]TournamentRequest, error) {
	rows, err := db.Model(&TournamentRequest{}).
		Joins("Left Join user_accounts on tournament_requests.sub = user_accounts.sub").
		Where("tournament_id = ?", tournamentId).
		Where("category = ?", TInvite).
		Select("tournament_requests.*, user_accounts.*").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invites := []TournamentRequest{}
	for rows.Next() {
		var invite TournamentRequest
		var user UserInfo
		_ = db.ScanRows(rows, &invite)
		_ = db.ScanRows(rows, &user)
		invite.User = &user
		invites = append(invites, invite)
	}
	return invites, nil
}

func (db *DB) DeleteTournamentApplication(requestId string) error {
	application := TournamentRequest{}
	err := db.Where(&TournamentRequest{RequestID: requestId, Category: TApplication}).
		First(&application).Error
	if err != nil {
		return err
	}
	return db.Delete(application).Error
}

func (db *DB) DeleteTournamentInvite(requestId string) error {
	invite := TournamentRequest{}
	err := db.Where(&TournamentRequest{RequestID: requestId, Category: TInvite}).
		First(&invite).Error
	if err != nil {
		return err
	}
	return db.Delete(invite).Error
}

func (db *DB) DeleteInviteByTournamentId(tournamentId string) error {
	return db.Where(&TournamentRequest{TournamentID: tournamentId, Category: TInvite}).
		Delete(&TournamentRequest{}).Error
}

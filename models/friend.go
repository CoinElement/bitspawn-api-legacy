/*

 */

package models

import (
	"database/sql"
	"time"
)

type Friend struct {
	UserIdOne       string    `json:"userIdOne" gorm:"primary_key"`
	UserIdTwo       string    `json:"userIdTwo" gorm:"primary_key"`
	UserOneDecision int       `json:"userOneDecision"` // 1: approve, 0: null, -1: reject
	UserTwoDecision int       `json:"userTwoDecision"` // 1: approve, 0: null, -1: reject
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updateAt"`
}

func (db *DB) CreateFriendInvite(friendInvite *Friend) error {
	return db.Create(&friendInvite).Error
}

func (db *DB) FetchFriendInviteByUserIds(applicantId, acceptUserId string) ([]Friend, error) {
	friendInvite := []Friend{}
	err := db.Where(&Friend{UserIdOne: applicantId, UserIdTwo: acceptUserId}).Find(&friendInvite).Error
	return friendInvite, err
}

type InvitationRequestUser struct {
	DisplayName     string `json:"displayName"`
	AvatarUrl       string `json:"avatarUrl"`
	UserTwoDecision int    `json:"userTwoDecision"`
}

func (db *DB) GetFriendInvitationByUserId(userId string) ([]InvitationRequestUser, error) {
	var rows *sql.Rows
	var err error
	invitations := []InvitationRequestUser{}

	rows, err = db.Table("friends").
		Select("friends.user_id_one, users.display_name, users.avatar_url, friends.user_two_decision").
		Joins("LEFT JOIN user_accounts AS users ON friends.user_id_one = users.sub").
		Where("friends.user_id_two = ? AND friends.user_two_decision = 0", userId).
		Rows()
	if err != nil {
		return invitations, err
	}
	defer rows.Close()

	for rows.Next() {
		var invitation InvitationRequestUser
		_ = db.ScanRows(rows, &invitation)
		invitations = append(invitations, invitation)
	}
	return invitations, nil
}

func (db *DB) GetFriendListByUserId(userId string) ([]Friend, error) {
	friends := []Friend{}
	err := db.Table("friends").
		Select("user_id_one, user_id_two, (user_one_decision + user_two_decision) AS total").
		Where(`user_id_one = (?) OR user_id_two = (?)`, userId, userId).
		Where(` (user_one_decision + user_two_decision) = 2`).
		Find(&friends).
		Error
	if err != nil {
		return nil, err
	}
	return friends, err
}

func (db *DB) ApproveFriendInvitation(applicantId, userId string) error {
	return db.Model(&Friend{}).
		Where(&Friend{UserIdOne: applicantId, UserIdTwo: userId}).
		Updates(map[string]interface{}{"user_two_decision": 1}).
		Error

}

func (db *DB) RejectFriendInvitation(applicantId, userId string) error {
	return db.Model(&Friend{}).
		Where(&Friend{UserIdOne: applicantId, UserIdTwo: userId}).
		Updates(map[string]interface{}{"user_two_decision": -1}).
		Error

}

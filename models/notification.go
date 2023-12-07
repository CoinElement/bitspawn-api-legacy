/*

 */

package models

import (
	"time"
)

type Notification struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updateAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
	Icon      string     `json:"icon"`
	Keyword   string     `json:"keyword"`
	Link      string     `json:"link"`
	Message   string     `json:"message"`
	Type      string     `json:"type"`
	Username  string     `json:"username"`
}

func (db *DB) CreateNotification(note *Notification) error {
	return db.Create(&note).Error
}

func (db *DB) ReadNotification(username string) ([]Notification, error) {
	notes := []Notification{}
	err := db.Where(&Notification{Username: username}).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}

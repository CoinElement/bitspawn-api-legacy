/*

 */

package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Base contains generated columns for tournament table only.
type UUIDBase struct {
	TournamentID uuid.UUID  `json:"tournamentId" gorm:"type:uuid;primary_key;"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updateAt"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *UUIDBase) BeforeCreate(tx *gorm.DB) error {

	if len(base.TournamentID.Bytes()) == 0 {

		uuid := uuid.NewV4()

		tx.Statement.SetColumn("TournamentID", uuid)
	}

	return nil
}

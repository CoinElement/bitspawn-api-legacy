/*

 */

package models

import (
	"time"
)

type Featured struct {
	ObjectTitle string    `json:"objectTitle" gorm:"primary_key"`
	ObjectID    string    `json:"objectId"`
	ObjectType  string    `json:"objectType"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (db *DB) InsertFeaturedObject(featured Featured) error {
	return db.Save(&featured).Error
}

func (db *DB) GetFeaturedObject(objectTitle string) (*Featured, error) {
	featured := Featured{}
	err := db.Where(&Featured{ObjectTitle: objectTitle}).First(&featured).Error
	if err != nil {
		return nil, err
	}
	return &featured, nil
}

func (db *DB) ListAllFeatured(objectType string) ([]Featured, error) {
	allFeatured := []Featured{}
	var err error
	if objectType != "" {
		err = db.Where(&Featured{ObjectType: objectType}).Find(&allFeatured).Error
	} else {
		err = db.Find(&allFeatured).Error
	}
	if err != nil {
		return nil, err
	}
	return allFeatured, nil
}

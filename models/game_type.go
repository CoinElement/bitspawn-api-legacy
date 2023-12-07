package models

import (
	uuid "github.com/satori/go.uuid"
)

// PlayRecord maps the users play record table in the database.
type GameType struct {
	GameTypeID uuid.UUID `json:"gameTypeId" gorm:"type:uuid;primary_key;"`
	GameType   string    `json:"gameType"`
}

func (db *DB) CreateGameType(game_type string) error {
	uuid := uuid.NewV4()
	game := GameType{
		GameTypeID: uuid,
		GameType:   game_type,
	}

	db.log.Debug(game)

	return db.Save(&game).Error
}

func (db *DB) GetSpecificGameType(game_type string) ([]GameType, error) {
	gameType := []GameType{}

	err := db.Where(&GameType{GameType: game_type}).Find(&gameType).Error
	if err != nil {
		return nil, err
	}

	return gameType, nil
}

func (db *DB) GetGameType() ([]GameType, error) {
	gameType := []GameType{}

	err := db.Find(&gameType).Error
	if err != nil {
		return nil, err
	}

	return gameType, nil
}

package models

// PlayRecord maps the users play record table in the database.
type GameSubType struct {
	GameType    string `json:"gameType" gorm:"type:text;primary_key"`
	GameSubType string `json:"gameSubType" gorm:"type:text;primary_key"`
	TeamSize    int    `json:"teamSize"`
}

func (db *DB) CreateGameSubType(game_type, game_sub_type string, teamSize int) error {

	game := GameSubType{
		GameType:    game_type,
		GameSubType: game_sub_type,
		TeamSize:    teamSize,
	}

	db.log.Debug(game)

	return db.Save(&game).Error
}

func (db *DB) GetSpecificGameSubType(game_type, game_sub_type string) ([]GameSubType, error) {
	gameSubType := []GameSubType{}

	err := db.Where(&GameSubType{GameType: game_type, GameSubType: game_sub_type}).Find(&gameSubType).Error
	if err != nil {
		return nil, err
	}

	return gameSubType, nil
}

func (db *DB) GetGameSubType(game_type string) ([]GameSubType, error) {
	gameSubType := []GameSubType{}

	err := db.Where(&GameSubType{GameType: game_type}).Find(&gameSubType).Error
	if err != nil {
		return nil, err
	}

	return gameSubType, nil
}

func (db *DB) DeleteSubtype(game_type, game_sub_type string) error {
	gameSubType := GameSubType{}
	err := db.Where(&GameSubType{GameType: game_type, GameSubType: game_sub_type}).First(&gameSubType).Error
	if err != nil {
		return err
	}
	return db.Delete(&gameSubType).Error
}

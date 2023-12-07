package models

import (
	"time"
)

type GamePlatform struct {
	UpdatedAt    time.Time
	Sub          string `gorm:"primaryKey"`
	GameName     string `gorm:"primaryKey"`
	PlatformName string `gorm:"primaryKey"`
	PlayerID     string
}

type GameDTO struct {
	Sub              string           `json:"sub"`
	GameName         string           `json:"gameName"`
	SelectedPlatform string           `json:"selectedPlatform"`
	Platforms        []PlatformDetail `json:"platforms"`
}

type PlatformDetail struct {
	PlatformName string `json:"platformName"`
	ID           string `json:"id"`
}

func (db *DB) GetUserGames(sub string) ([]GameDTO, error) {
	gamePlatforms := []GamePlatform{}
	err := db.Where("sub = ?", sub).Order("updated_at").Find(&gamePlatforms).Error
	if err != nil {
		return nil, err
	}
	gamePlatformMap := initGamePlatform()
	for _, g := range gamePlatforms {
		gamePlatformMap[g.GameName][g.PlatformName] = g.PlayerID
	}
	outputGameDTO := gamePlatformMapToDTO(gamePlatformMap)
	for i := range outputGameDTO {
		outputGameDTO[i].Sub = sub
		for _, g := range gamePlatforms {
			if g.GameName == outputGameDTO[i].GameName {
				outputGameDTO[i].SelectedPlatform = g.PlatformName
			}
		}
	}
	return outputGameDTO, nil
}

func (db *DB) GetSpecificGame(sub string, gameName string) (*GameDTO, error) {
	allGames, err := db.GetUserGames(sub)
	if err != nil {
		return nil, err
	}
	for _, g := range allGames {
		if g.GameName == gameName {
			return &g, nil
		}
	}
	return nil, err
}

func (db *DB) UpdateGamePlatformPlayerID(sub, gameName, platformName, playerId string) ([]GameDTO, error) {
	gamePlatformPlayerId := GamePlatform{
		Sub:          sub,
		GameName:     gameName,
		PlatformName: platformName,
		PlayerID:     playerId,
	}
	err := db.Save(&gamePlatformPlayerId).Error
	if err != nil {
		return nil, err
	}
	return db.GetUserGames(sub)
}

func (db *DB) DisconnectSpecificGame(sub string, gameName string, platformName string) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	gamePlatformPlayerId := GamePlatform{
		Sub:          sub,
		GameName:     gameName,
		PlatformName: platformName,
	}
	err := tx.Delete(&gamePlatformPlayerId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Exec("UPDATE user_game_accounts SET game_account = game_account::jsonb - lower(?) WHERE sub = ?", platformName, sub).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) CheckPlatformUsernameUsed(sub string, gameName string, platformName string, platformUserName string) (bool, error) {
	var used bool
	gamePlatforms := []GamePlatform{}
	err := db.Where(&GamePlatform{GameName: gameName, PlatformName: platformName, PlayerID: platformUserName}).
		Not(&GamePlatform{Sub: sub}).Find(&gamePlatforms).Error
	if err != nil {
		return false, err
	}
	if len(gamePlatforms) > 0 {
		used = true
	} else {
		used = false
	}
	return used, nil
}

func (db *DB) ReplacePlatformUsername(sub string, gameName string, platformName string, platformUserName string) (bool, error) {
	var done bool
	err := db.Exec("CALL replace_platform_username('" + sub + "','" + gameName + "','" + platformName + ",'" + platformUserName + "')").Scan(&done).Error
	if err != nil {
		return false, err
	}
	return done, nil
}

func gamePlatformMapToDTO(m map[string]map[string]string) []GameDTO {
	gameDTOs := make([]GameDTO, 0, len(m))
	for gameName, m2 := range m {
		platforms := make([]PlatformDetail, 0, len(m2))
		for platformName, playerID := range m2 {
			p := PlatformDetail{
				PlatformName: platformName,
				ID:           playerID,
			}
			platforms = append(platforms, p)
		}
		g := GameDTO{
			GameName:  gameName,
			Platforms: platforms,
		}
		gameDTOs = append(gameDTOs, g)
	}
	return gameDTOs
}

func initGamePlatform() map[string]map[string]string {
	return map[string]map[string]string{
		"CODMWBR": {
			"XBOX":      "",
			"PSN":       "",
			"BATTLENET": "",
		},
		"APEX_LEGENDS": {
			"XBOX":     "",
			"PSN":      "",
			"STEAM":    "",
			"ORIGIN":   "",
			"NINTENDO": "",
		},
	}
}

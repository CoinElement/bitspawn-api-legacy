/*

 */

package models

type Console struct {
	GameType    string `gorm:"type:text;primary_key"`
	ConsoleName string `gorm:"type:text;primary_key"`
}

func (db *DB) GetConsolesForGameType(gameType string) ([]string, error) {
	consoles := []Console{}
	err := db.Where(Console{GameType: gameType}).Find(&consoles).Error
	if err != nil {
		return nil, err
	}
	var consoleNames []string
	for _, p := range consoles {
		consoleNames = append(consoleNames, p.ConsoleName)
	}
	return consoleNames, nil
}

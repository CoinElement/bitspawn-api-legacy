package models

// PlayRecord maps the users play record table in the database.
type TournamentFormat struct {
	Format     string `json:"tournamentFormat" gorm:"primary_key;"`
	Label      string `json:"formatLabel"`
	MaxWinners int    `json:"maxWinners"`
}

type BestOfNFormat struct {
	AllRounds  string `json:"allRounds"`
	SemiFinals string `json:"semiFinals"`
	Finals     string `json:"finals"`
}

func (db *DB) CreateFormat(format, label string, maxWinners int) error {
	tournamentFormat := TournamentFormat{
		Format:     format,
		Label:      label,
		MaxWinners: maxWinners,
	}
	return db.Save(&tournamentFormat).Error
}

func (db *DB) GetFormat(format string) ([]TournamentFormat, error) {
	tournamentFormat := []TournamentFormat{}
	err := db.Where(&TournamentFormat{Format: format}).Find(&tournamentFormat).Error
	if err != nil {
		return nil, err
	}
	return tournamentFormat, nil
}

func (db *DB) ListFormat() ([]TournamentFormat, error) {
	tournamentFormat := []TournamentFormat{}
	err := db.Find(&tournamentFormat).Error
	if err != nil {
		return nil, err
	}
	return tournamentFormat, nil
}

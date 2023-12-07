package models

import "fmt"

type MatchGame struct {
	TournamentID  string `gorm:"type:uuid;primaryKey" json:"tournamentId"`
	MatchID       string `gorm:"type:uuid;primaryKey" json:"matchId" `
	GameNumber    int    `gorm:"primaryKey" json:"gameNumber"`
	TeamOneScore  int    `json:"teamOneScore"`
	TeamTwoScore  int    `json:"teamTwoScore"`
	ScreenshotOne string `json:"screenshotOne"`
	ScreenshotTwo string `json:"screenshotTwo"`
	ReportedBy    int    `json:"reportedBy"`
}

// func (db *DB) CreateMatchGame(match *Match, teamOneScore, teamTwoScore, reportedBy int) (*MatchGame, error) {
// 	existingMatchGames, err := db.GetMatchGamesByMatch(match.MatchID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	matchGame := MatchGame{
// 		TournamentID: match.TournamentID,
// 		MatchID:      match.MatchID,
// 		GameNumber:   len(existingMatchGames) + 1,
// 		TeamOneScore: teamOneScore,
// 		TeamTwoScore: teamTwoScore,
// 		ReportedBy: reportedBy,
// 	}
// }

func (db *DB) GetMatchGame(matchId string, gameNumber int) (*MatchGame, error) {
	matchGame := MatchGame{}
	err := db.DB.Where(&MatchGame{MatchID: matchId, GameNumber: gameNumber}).First(&matchGame).Error
	if err != nil {
		return nil, err
	}
	return &matchGame, nil
}

func (db *DB) GetMatchGamesByMatch(matchId string) ([]MatchGame, error) {
	matchGames := []MatchGame{}
	err := db.DB.Where(&MatchGame{MatchID: matchId}).Order("game_number").Find(&matchGames).Error
	if err != nil {
		return nil, err
	}
	return matchGames, nil
}

func (db *DB) CreateMatchGame(match *Match, existingMatchGames []MatchGame, teamOneScore, teamTwoScore, reportedBy int) (MatchGame, error) {
	largestGameNumber := 0
	for _, m := range existingMatchGames {
		if m.GameNumber > largestGameNumber {
			largestGameNumber = m.GameNumber
		}
	}
	mg := MatchGame{
		TournamentID: match.TournamentID,
		MatchID:      match.MatchID,
		GameNumber:   largestGameNumber + 1,
		TeamOneScore: teamOneScore,
		TeamTwoScore: teamTwoScore,
		ReportedBy:   reportedBy,
	}
	err := db.Create(mg).Error
	return mg, err
}

func (db *DB) UpdateMatchGame(match *Match, gameNumber, teamOneScore, teamTwoScore, reportedBy int) (*MatchGame, *Match, error) {
	mg := MatchGame{
		TournamentID: match.TournamentID,
		MatchID:      match.MatchID,
		GameNumber:   gameNumber,
		TeamOneScore: teamOneScore,
		TeamTwoScore: teamTwoScore,
		ReportedBy:   reportedBy,
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, nil, err
	}

	err := tx.Save(mg).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	matchGames, err := db.GetMatchGamesByMatch(match.MatchID)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	if len(matchGames) < gameNumber {
		tx.Rollback()
		return nil, nil, err
	}
	// Update minor scores locally, before the update is committed to db
	matchGames[gameNumber-1].TeamOneScore = mg.TeamOneScore
	matchGames[gameNumber-1].TeamTwoScore = mg.TeamTwoScore

	majorScorePlayerOne := 0
	majorScorePlayerTwo := 0
	for _, g := range matchGames {
		if g.TeamOneScore > g.TeamTwoScore {
			majorScorePlayerOne += 1
		} else {
			majorScorePlayerTwo += 1
		}
	}
	match.TeamOneScore = majorScorePlayerOne
	match.TeamTwoScore = majorScorePlayerTwo
	if teamOneScore > match.BestOfN/2 || teamTwoScore > match.BestOfN/2 {
		match.Status = "Finished"
		if teamOneScore > teamTwoScore {
			match.Result = 1
		} else {
			match.Result = 2
		}
	}

	err = tx.Save(match).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	return &mg, match, nil
}

func (db *DB) UpdateMatchGameScreenshot(match *Match, gameNumber int, screenshotFileName string, reportedBy int) error {
	var err error
	var screenshot_column string
	if reportedBy == 1 {
		screenshot_column = "screenshot_one"
	} else if reportedBy == 2 {
		screenshot_column = "screenshot_two"
	} else {
		return fmt.Errorf("unsupported reportedBy value - %d", reportedBy)
	}
	err = db.DB.Model(&MatchGame{}).
		Where(&MatchGame{TournamentID: match.TournamentID, MatchID: match.MatchID, GameNumber: gameNumber}).
		Update(screenshot_column, screenshotFileName).Error
	return err
}

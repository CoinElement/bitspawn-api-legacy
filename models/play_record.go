/*

 */

package models

import (
	"database/sql"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

// PlayRecord maps the users play record table in the database.
type PlayRecord struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	UserId        string `json:"userId" gorm:"primary_key"`
	TournamentId  string `json:"tournamentId" gorm:"primary_key;type:uuid"`
	PublicAddress string `json:"publicAddress"`
	Clan          string `json:"clan"`
	Club          string `json:"club"`
	GameType      string `json:"gameType"`
	PrizeEarned   int64  `json:"prizeEarned"`
	RankingNumber int    `json:"rankingNumber"`
	RecordType    string `json:"-"`
}

type PlayRecordOutput struct {
	GameType       string    `json:"gameType"`
	UserId         string    `json:"userId" `
	TournamentId   string    `json:"tournamentId"`
	TournamentName string    `json:"tournamentName"`
	Status         string    `json:"status"`
	EntryFee       string    `json:"entryFee"`
	PrizeEarned    int64     `json:"prizeEarned"`
	FinishDate     time.Time `json:"finishDate"`
}

func (db *DB) InsertPlayRecord(playRecord *PlayRecord) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(playRecord).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&TournamentData{}).Where("tournament_id = ?", playRecord.TournamentId).
		Update("participant_count", gorm.Expr("participant_count + 1")).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) BatchInsertPlayRecords(playRecords []PlayRecord) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, pr := range playRecords {
		if err := tx.Create(&pr).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Model(&TournamentData{}).Where("tournament_id = ?", playRecords[0].TournamentId).
		Update("participant_count", gorm.Expr("participant_count + ?", len(playRecords))).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) GetPlayRecordByTournament(tournamentId string) ([]PlayRecord, error) {
	playRecord := []PlayRecord{}
	err := db.Where("tournament_id = ?", tournamentId).Order("ranking_number").Find(&playRecord).Error
	if err != nil {
		return nil, err
	}
	return playRecord, nil
}

func (db *DB) GetPlayRecordByTournamentAndPublicAddress(tournamentId, publicAddress string) ([]PlayRecord, error) {
	playRecord := []PlayRecord{}
	err := db.Where(&PlayRecord{PublicAddress: publicAddress, TournamentId: tournamentId}).Find(&playRecord).Error
	if err != nil {
		return nil, err
	}
	return playRecord, nil
}

func (db *DB) GetPlayRecord(username string, page, per_page int) ([]PlayRecord, error) {
	playRecord := []PlayRecord{}

	offset := (page - 1) * per_page
	limit := per_page

	err := db.Where("user_id = ?", username).Limit(limit).Offset(offset).Find(&playRecord).Error
	if err != nil {
		return nil, err
	}

	return playRecord, nil
}

func (db *DB) GetPlayRecordByGameType(username, game_type string, page, per_page int) ([]PlayRecord, error) {
	playRecord := []PlayRecord{}

	offset := (page - 1) * per_page
	limit := per_page

	err := db.Where(&PlayRecord{UserId: username, GameType: game_type}).Limit(limit).Offset(offset).Find(&playRecord).Error
	if err != nil {
		return nil, err
	}

	return playRecord, nil
}

func (db *DB) GetRegisteredTournaments(username string) ([]TournamentResponse, error) {
	tournaments := []TournamentData{}
	err := db.Joins("Inner Join play_records On tournament_data.tournament_id=play_records.tournament_id").
		Where("user_id = ? AND status != 'Cancelled'", username).Order("tournament_date").
		Find(&tournaments).
		Error
	if err != nil {
		return nil, err
	}
	trs := []TournamentResponse{}
	for _, t := range tournaments {
		tr, err := db.FormatTournamentForResponse(t)
		if err != nil {
			return nil, err
		}
		trs = append(trs, *tr)
	}
	return trs, nil
}

func (db *DB) GetPlayRecordsByGameType(username, game_type string) ([]PlayRecord, error) {
	playRecord := []PlayRecord{}

	err := db.Where(&PlayRecord{UserId: username, GameType: game_type}).Find(&playRecord).Error
	if err != nil {
		return nil, err
	}

	return playRecord, nil
}

func (db *DB) GetClubTournaments(clubName, gameType string, page, per_page int) ([]TournamentResponse, error) {
	offset := (page - 1) * per_page
	limit := per_page

	tournaments := []TournamentData{}
	err := db.Joins("LEFT JOIN play_records ON tournament_data.tournament_id=play_records.tournament_id").
		Where("club = ? AND play_records.game_type = ?", clubName, gameType).
		Where("tournament_data.status != ?", "Cancelled").
		Order("tournament_date").
		Select("DISTINCT play_records.tournament_id, tournament_data.*").
		Limit(limit).Offset(offset).
		Find(&tournaments).Error
	if err != nil {
		return nil, err
	}

	trs := []TournamentResponse{}
	for _, t := range tournaments {
		tr, err := db.FormatTournamentForResponse(t)
		if err != nil {
			return nil, err
		}
		trs = append(trs, *tr)
	}
	return trs, nil
}

func (db *DB) GetClubFinishRecords(clubName, gameType string, page, per_page int) ([]PlayRecordOutput, error) {
	offset := (page - 1) * per_page
	limit := per_page

	rows, err := db.Model(&PlayRecord{}).
		Joins("Left Join tournament_data on play_records.tournament_id = tournament_data.tournament_id").
		Where("status = 'Completed'").
		Where("club = ? AND play_records.game_type = ?", clubName, gameType).
		Group("play_records.tournament_id").
		Select("play_records.tournament_id,SUM(play_records.prize_earned) AS prize_earned,MAX(tournament_data.tournament_name) AS tournament_name,SUM(tournament_data.entry_fee::INTEGER) AS entry_fee,MAX(tournament_data.tournament_date) AS finish_date").
		Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	playRecordsOutput := []PlayRecordOutput{}
	for rows.Next() {
		var playRecordOutput PlayRecordOutput
		_ = db.ScanRows(rows, &playRecordOutput)
		playRecordsOutput = append(playRecordsOutput, playRecordOutput)
	}
	return playRecordsOutput, nil
}

func (db *DB) CountFinishRecords(username string) (int64, error) {
	var count int64
	err := db.Table("play_records").
		Joins("Left Join tournament_data on play_records.tournament_id = tournament_data.tournament_id").
		Where("status = 'Completed' AND user_id = ?", username).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) GetFinishRecords(username string, page, per_page int) ([]PlayRecordOutput, error) {
	offset := (page - 1) * per_page
	limit := per_page

	rows, err := db.Model(&PlayRecord{}).
		Joins("Left Join tournament_data on play_records.tournament_id = tournament_data.tournament_id").
		Where("status = 'Completed' AND user_id = ?", username).
		Select("play_records.*,tournament_data.tournament_name, tournament_data.status, tournament_data.entry_fee, tournament_data.tournament_date AS finish_date").
		Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	playRecordsOutput := []PlayRecordOutput{}
	for rows.Next() {
		var playRecordOutput PlayRecordOutput
		_ = db.ScanRows(rows, &playRecordOutput)
		playRecordsOutput = append(playRecordsOutput, playRecordOutput)
	}
	return playRecordsOutput, nil
}

func (db *DB) CountWinRecords(username string) (int64, error) {
	var count int64
	err := db.Table("play_records").
		Where("prize_earned != '0' AND user_id = ?", username).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) AccumulateEarning(username string) int64 {
	var sum int64
	row := db.Table("play_records").
		Where("user_id = ?", username).
		Select("sum(prize_earned)").
		Row()
	_ = row.Scan(&sum)
	return sum
}

type Result struct {
	GameType string
	Count    int64
}

func (db *DB) CalculateGamePlay(username string) (*Result, error) {
	var count int64
	result := Result{}

	err := db.Table("play_records").
		Where("user_id = ?", username).
		Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count != 0 {
		err = db.Table("play_records").
			Where("user_id = ?", username).
			Select("game_type, count(game_type) as count").
			Group("game_type").
			Count(&count).
			Order("count desc").
			First(&result).
			Error
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}
func (db *DB) CountFinishRecordsByGameType(username string, game_type string) (int64, error) {
	var count int64
	err := db.Table("play_records").
		Joins("Left Join tournament_data on play_records.tournament_id = tournament_data.tournament_id").
		Where("status = 'Completed' AND user_id = ? AND play_records.game_type = ?", username, game_type).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) CountWinRecordsByGameType(username string, game_type string) (int64, error) {
	var count int64
	err := db.Table("play_records").
		Where("prize_earned != '0' AND user_id = ? AND game_type = ?", username, game_type).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) AccumulateEarningByGameType(username string, game_type string) int64 {
	var sum int64
	row := db.Table("play_records").
		Where("user_id = ? AND game_type = ?", username, game_type).
		Select("sum(prize_earned)").
		Row()
	_ = row.Scan(&sum)
	return sum
}

func (db *DB) DeletePlayRecordByTournamentAndPublicAddress(tournamentId, publicAddress string) ([]PlayRecord, error) {
	playRecords := []PlayRecord{}
	err := db.Where(&PlayRecord{PublicAddress: publicAddress, TournamentId: tournamentId}).Find(&playRecords).Error
	if err != nil {
		return nil, err
	}
	return playRecords, db.Delete(playRecords).Error
}

func (db *DB) DeletePlayRecord(username string, tournamentId string) (*PlayRecord, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	playRecord := PlayRecord{}
	err := db.Where(&PlayRecord{UserId: username, TournamentId: tournamentId}).First(&playRecord).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = db.Delete(playRecord).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Model(&TournamentData{}).Where("tournament_id = ?", playRecord.TournamentId).
		Update("participant_count", gorm.Expr("participant_count - 1")).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Where("tournament_id = ? AND player = ?", tournamentId, username).Delete(&Team{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &playRecord, tx.Commit().Error
}

func (db *DB) BatchDeletePlayRecords(playRecords []PlayRecord) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, pr := range playRecords {
		if err := tx.Delete(&pr).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (db *DB) CountMyPlayRecordPages(username string, per_page int) (int64, error) {
	var count int64
	err := db.Table("play_records").
		Where("user_id = ?", username).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	pagefloat := math.Ceil(float64(count) / float64(per_page))
	page := int64(pagefloat)

	return page, nil
}

type UserInfo struct {
	Sub         string `json:"sub"`
	Username    string `json:"-"`
	DisplayName string `json:"displayName"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (db *DB) CountPlayedUser(game_type string) (int64, error) {
	var count int64
	var err error
	if game_type != "" {
		err = db.Table("play_records").
			Where("game_type = ? ", game_type).
			Group("user_id").
			Count(&count).Error
		// err = db.Table("tournament_data").Joins("LEFT JOIN play_records AS records ON tournament_data.tournament_id = match.tournament_id").
		// 	Where("tournament_data.status = ? AND game_type = ? ", "Completed", game_type).
		// 	Count(&count).Error
	} else {
		err = db.Table("play_records").
			Group("user_id").
			Count(&count).Error
		// err = db.Table("tournament_data").Joins("LEFT JOIN play_records AS records ON tournament_data.tournament_id = match.tournament_id").
		// 	Where("tournament_data.status = ?", "Completed").
		// 	Count(&count).Error
	}
	if err != nil {
		return 0, err
	}

	return count, nil
}

type PlayedGameAttendUser struct {
	UserInfo UserInfo `json:"userInfo"`
}

func (db *DB) PlayedUserAvatar(game_type string) ([]UserInfo, error) {
	attendUsers := []UserInfo{}
	var rows *sql.Rows
	var err error
	if game_type != "" {
		rows, err = db.Table("play_records").
			Select("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url").
			Joins("LEFT JOIN user_accounts ON play_records.user_id = user_accounts.username").
			Where("play_records.game_type = ? ", game_type).
			Group("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url").
			Rows()
	} else {
		rows, err = db.Table("play_records").
			Select("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url").
			Joins("LEFT JOIN user_accounts ON play_records.user_id = user_accounts.username").
			Group("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url,play_records.updated_at").
			Order("play_records.updated_at DESC").
			Rows()
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var attendUser UserInfo
		_ = db.ScanRows(rows, &attendUser)
		attendUsers = append(attendUsers, attendUser)
	}
	return attendUsers, nil
}

func (db DB) GetParticipants(tournamentId string) ([]PlayRecord, error) {
	playRecord := []PlayRecord{}
	err := db.DB.Where("tournament_id = ?", tournamentId).Find(&playRecord).Error
	if err != nil {
		return nil, err
	}
	return playRecord, nil
}

func (db DB) GetAllParticipantsInfo(tournamentId string) ([]UserInfo, error) {
	allParticipants := []UserInfo{}
	var rows *sql.Rows
	var err error
	rows, err = db.Table("play_records").
		Select("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url").
		Joins("LEFT JOIN user_accounts ON play_records.user_id = user_accounts.username").
		Where("play_records.tournament_id = ?", tournamentId).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var participant UserInfo
		_ = db.ScanRows(rows, &participant)
		allParticipants = append(allParticipants, participant)
	}
	return allParticipants, nil
}

func (db DB) ReportWinner(tournamentId string, winner string, rank int, prizeWon int) error {
	playRecord := PlayRecord{}
	err := db.DB.Where(`"tournament_id" = ? AND "user_id" = ?`, tournamentId, winner).First(&playRecord).Error
	if err != nil {
		return err
	}
	err = db.DB.Model(&playRecord).Updates(map[string]interface{}{"ranking_number": rank, "prize_earned": prizeWon}).Error
	if err != nil {
		return fmt.Errorf("Error in Update rank for tournament %s, winner %s, rank %d: %v", tournamentId, winner, rank, err)
	}
	return nil
}

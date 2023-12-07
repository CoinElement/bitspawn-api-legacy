/*

 */

package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ChallengeRecord maps the users play record table in the database.
type ChallengeRecord struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	ID                uint      `gorm:"primary_key"`
	UserId            string    `json:"userId"`
	ChallengeId       string    `json:"challengeId" gorm:"type:uuid"`
	Platform          string    `json:"platform"`
	PlayerId          string    `json:"playerId"`
	PublicAddress     string    `json:"publicAddress"`
	RegisterDate      time.Time `json:"registerDate"`
	ChallengeDeadline time.Time `json:"challengeDeadline"`
	ChallengeType     string    `json:"challengeType"`
	Score             float64   `json:"score"`
	RankingNumber     int       `json:"rankingNumber"`
	ScoreReported     bool      `json:"scoreReported"`
	PrizeEarned       int64     `json:"prizeEarned"`

	Scoring  map[string]*string `json:"scoring"`
	ScoreMap map[string]*string `json:"scoreMap"`
	TeamName string             `json:"TeamName"`
}

func (t ChallengeRecord) toGorm() *ChallengeRecordGorm {
	scoring_json, _ := json.Marshal(t.Scoring)
	scoreMap_json, _ := json.Marshal(t.ScoreMap)
	challengeRecordGorm := ChallengeRecordGorm{
		t.CreatedAt, t.UpdatedAt,
		t.ID, t.UserId, t.ChallengeId, t.Platform, t.PlayerId, t.PublicAddress,
		t.RegisterDate, t.ChallengeDeadline, t.ChallengeType,
		t.Score, t.RankingNumber, t.ScoreReported, t.PrizeEarned,
		datatypes.JSON(scoring_json), datatypes.JSON(scoreMap_json), t.TeamName,
	}
	return &challengeRecordGorm
}

type ChallengeRecordGorm struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	ID                uint      `gorm:"primary_key"`
	UserId            string    `json:"userId"`
	ChallengeId       string    `json:"challengeId" gorm:"type:uuid"`
	Platform          string    `json:"platform"`
	PlayerId          string    `json:"playerId"`
	PublicAddress     string    `json:"publicAddress"`
	RegisterDate      time.Time `json:"registerDate"`
	ChallengeDeadline time.Time `json:"challengeDeadline"`
	ChallengeType     string    `json:"challengeType"`
	Score             float64   `json:"score"`
	RankingNumber     int       `json:"rankingNumber"`
	ScoreReported     bool      `json:"scoreReported"`
	PrizeEarned       int64     `json:"prizeEarned"`

	Scoring  datatypes.JSON `json:"scoring"`
	ScoreMap datatypes.JSON `json:"scoreMap"`
	TeamName string         `json:"teamName"`
}

func (t ChallengeRecordGorm) fromGorm() ChallengeRecord {
	var scoring, scoreMap map[string]*string
	_ = json.Unmarshal([]byte(t.Scoring), &scoring)
	_ = json.Unmarshal([]byte(t.ScoreMap), &scoreMap)
	challengeRecord := ChallengeRecord{
		t.CreatedAt, t.UpdatedAt,
		t.ID, t.UserId, t.ChallengeId, t.Platform, t.PlayerId, t.PublicAddress,
		t.RegisterDate, t.ChallengeDeadline, t.ChallengeType,
		t.Score, t.RankingNumber, t.ScoreReported, t.PrizeEarned,
		scoring, scoreMap, t.TeamName,
	}
	return challengeRecord
}

func (ChallengeRecordGorm) TableName() string {
	return "challenge_records"
}

type ChallengeRecordOutput struct {
	UserId        string    `json:"userId" `
	ChallengeId   string    `json:"challengeId"`
	ChallengeName string    `json:"challengeName"`
	ChallengeType string    `json:"challengeType"`
	Status        string    `json:"status"`
	EntryFee      string    `json:"entryFee"`
	PrizeEarned   int64     `json:"prizeEarned"`
	FinishDate    time.Time `json:"finishDate"`
}

type ChallengeLeader struct {
	AvatarUrl   string          `json:"avatarUrl"`
	DisplayName string          `json:"displayName"`
	Score       float64         `json:"score"`
	TeamName    string          `json:"teamName"`
	ScoreMap    json.RawMessage `json:"scoreMap"`
}

func (db *DB) InsertChallengeRecord(challengeRecord *ChallengeRecord) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(challengeRecord.toGorm()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&ChallengeData{}).Where("challenge_id = ?", challengeRecord.ChallengeId).
		Update("participant_count", gorm.Expr("participant_count + 1")).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) BatchInsertChallengeRecords(challengeRecords []ChallengeRecord) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, cr := range challengeRecords {
		if err := tx.Create(cr.toGorm()).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (db *DB) UpdateChallengeRecord(challengeRecord *ChallengeRecord) error {
	err := db.DB.Save(challengeRecord.toGorm()).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetChallengeRecordsToReport() ([]ChallengeRecord, error) {
	challengeRecordsGorm := []ChallengeRecordGorm{}
	err := db.Where("score_reported = ?", false).Find(&challengeRecordsGorm).Error
	if err != nil {
		return nil, err
	}
	challengeRecords := []ChallengeRecord{}
	for _, cr := range challengeRecordsGorm {
		challengeRecords = append(challengeRecords, cr.fromGorm())
	}
	return challengeRecords, nil
}

func (db *DB) GetChallengeRecordByChallenge(challengeId string) ([]ChallengeRecord, error) {
	challengeRecordsGorm := []ChallengeRecordGorm{}
	err := db.Where(&ChallengeRecordGorm{ChallengeId: challengeId}).Order("score desc").Find(&challengeRecordsGorm).Error
	if err != nil {
		return nil, err
	}
	challengeRecords := []ChallengeRecord{}
	for _, cr := range challengeRecordsGorm {
		challengeRecords = append(challengeRecords, cr.fromGorm())
	}
	return challengeRecords, nil
}

func (db *DB) GetChallengeRecordsByUser(challengeId, userId string) ([]ChallengeRecord, error) {
	challengeRecordsGorm := []ChallengeRecordGorm{}
	err := db.Where(&ChallengeRecordGorm{ChallengeId: challengeId, UserId: userId}).Order("score desc").Find(&challengeRecordsGorm).Error
	if err != nil {
		return nil, err
	}
	challengeRecords := []ChallengeRecord{}
	for _, cr := range challengeRecordsGorm {
		challengeRecords = append(challengeRecords, cr.fromGorm())
	}
	return challengeRecords, nil
}

type TeamScore struct {
	TeamName  string            `json:"teamName"`
	TeamScore float64           `json:"teamScore"`
	Players   []ChallengeLeader `json:"players"`
}

// GroupByTeam is used to group by the result by team name
func (db *DB) GroupByTeam(challengeLeaders []ChallengeLeader) []TeamScore {
	teamScores := make(map[string]*TeamScore)
	for _, leader := range challengeLeaders {
		if len(leader.TeamName) == 0 {
			continue
		}
		if _, ok := teamScores[leader.TeamName]; !ok {
			teamScores[leader.TeamName] = &TeamScore{
				TeamName:  leader.TeamName,
				TeamScore: 0,
			}
		}
		t := teamScores[leader.TeamName]
		t.Players = append(t.Players, leader)
		t.TeamScore += leader.Score
	}
	result := []TeamScore{}
	for _, score := range teamScores {
		result = append(result, *score)
	}
	// sorting result based on the team scores
	sort.Slice(result, func(i, j int) bool {
		return result[i].TeamScore > result[j].TeamScore
	})
	return result
}

func (db *DB) GetChallengeLeaderboard(challengeId string, per_page int) ([]ChallengeLeader, error) {
	challengeRecordsGorm := []ChallengeRecordGorm{}
	err := db.Where(&ChallengeRecordGorm{ChallengeId: challengeId}).Order("score desc").Limit(per_page).Find(&challengeRecordsGorm).Error
	if err != nil {
		return nil, err
	}
	challengeLeaders := []ChallengeLeader{}
	for _, record := range challengeRecordsGorm {
		user, err := db.GetUser(record.UserId)
		if err != nil {
			return nil, err
		}
		challengeLeader := ChallengeLeader{
			AvatarUrl:   user.AvatarUrl,
			DisplayName: user.DisplayName,
			Score:       record.Score,
			ScoreMap:    json.RawMessage(record.ScoreMap),
			TeamName:    record.TeamName,
		}
		challengeLeaders = append(challengeLeaders, challengeLeader)
	}
	return challengeLeaders, nil
}

func (db *DB) GetChallengeRecordWithPage(username string, page, per_page int) ([]ChallengeRecord, error) {
	challengeRecordsGorm := []ChallengeRecordGorm{}
	offset := (page - 1) * per_page
	limit := per_page
	err := db.Where("user_id = ?", username).Limit(limit).Offset(offset).Find(&challengeRecordsGorm).Error
	if err != nil {
		return nil, err
	}
	challengeRecords := []ChallengeRecord{}
	for _, cr := range challengeRecordsGorm {
		challengeRecords = append(challengeRecords, cr.fromGorm())
	}
	return challengeRecords, nil
}

func (db *DB) GetChallengeRecordByChallengeTypeWithPage(username, challenge_type string, page, per_page int) ([]ChallengeRecord, error) {
	challengeRecordsGorm := []ChallengeRecordGorm{}
	offset := (page - 1) * per_page
	limit := per_page
	err := db.Where(&ChallengeRecordGorm{UserId: username, ChallengeType: challenge_type}).Limit(limit).Offset(offset).Find(&challengeRecordsGorm).Error
	if err != nil {
		return nil, err
	}
	challengeRecords := []ChallengeRecord{}
	for _, cr := range challengeRecordsGorm {
		challengeRecords = append(challengeRecords, cr.fromGorm())
	}
	return challengeRecords, nil
}

func (db *DB) GetChallengeRecordsByChallengeType(username, challenge_type string) ([]ChallengeRecord, error) {
	challengeRecordsGorm := []ChallengeRecordGorm{}
	err := db.Where(&ChallengeRecordGorm{UserId: username, ChallengeType: challenge_type}).Find(&challengeRecordsGorm).Error
	if err != nil {
		return nil, err
	}
	challengeRecords := []ChallengeRecord{}
	for _, cr := range challengeRecordsGorm {
		challengeRecords = append(challengeRecords, cr.fromGorm())
	}
	return challengeRecords, nil
}

func (db *DB) DeleteChallengeRecord(challengeRecord *ChallengeRecord) error {
	return db.Delete(challengeRecord.toGorm()).Error
}

func (db *DB) BatchDeleteChallengeRecords(challengeRecords []ChallengeRecord) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, cr := range challengeRecords {
		if err := tx.Delete(cr.toGorm()).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (db *DB) CountFinishedChallengeRecords(username string, per_page int) (int64, error) {
	var count int64
	err := db.Model(&ChallengeRecordGorm{}).
		Joins("Left Join challenge_data on challenge_records.challenge_id = challenge_data.challenge_id").
		Where("status = 'Completed' AND user_id = ?", username).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetFinishedChallengeRecords(username string, page, per_page int) ([]ChallengeRecordOutput, error) {
	offset := (page - 1) * per_page
	limit := per_page

	rows, err := db.Model(&ChallengeRecordGorm{}).
		Joins("Left Join challenge_data on challenge_records.challenge_id = challenge_data.challenge_id").
		Where("status = 'Completed' AND user_id = ?", username).
		Select("challenge_records.*, challenge_data.challenge_name, challenge_data.status, challenge_data.entry_fee, challenge_data.cutoff_date AS finish_date").
		Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	challengeRecordsOutput := []ChallengeRecordOutput{}
	for rows.Next() {
		var challengeRecordOutput ChallengeRecordOutput
		_ = db.ScanRows(rows, &challengeRecordOutput)
		challengeRecordsOutput = append(challengeRecordsOutput, challengeRecordOutput)
	}
	return challengeRecordsOutput, nil
}

func (db *DB) ReportChallengeWinners(challenge *Challenge) ([]ChallengeRecord, error) {
	challengeRecords, err := db.GetChallengeRecordByChallenge(challenge.ChallengeID.String())
	if err != nil {
		return nil, fmt.Errorf("error in GetChallengeRecordByChallenge: %v", err)
	}
	if len(challengeRecords) < challenge.NumberOfWinners {
		return nil, fmt.Errorf("not enough winners")
	}
	prizeAllocation := strings.Split(challenge.PrizeAllocation, ",")
	if len(prizeAllocation) != challenge.NumberOfWinners {
		return nil, fmt.Errorf("prize allocation string does not match number of winners")
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}

	for i := 0; i < int(challenge.NumberOfWinners); i++ {
		pa, _ := strconv.ParseInt(prizeAllocation[i], 10, 64)
		err := tx.Model(&ChallengeRecordGorm{}).
			Where("challenge_records.challenge_id = ?", challenge.ChallengeID).
			Where("challenge_records.user_id = ?", challengeRecords[i].UserId).
			Updates(ChallengeRecordGorm{
				RankingNumber: i + 1,
				PrizeEarned:   pa * challenge.TargetPrizePool / 100,
			}).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return challengeRecords, nil
}

func (db *DB) ChallengePlayerAvatar(game_type string) ([]UserInfo, error) {
	attendUsers := []UserInfo{}
	var rows *sql.Rows
	var err error
	if game_type != "" {
		rows, err = db.Table("challenge_records").
			Select("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url").
			Joins("LEFT JOIN user_accounts ON challenge_records.user_id = user_accounts.username").
			Where("challenge_records.challenge_type = ? ", game_type).
			Group("challenge_records.user_id, user_accounts.display_name, user_accounts.avatar_url").
			Rows()
	} else {
		rows, err = db.Table("challenge_records").
			Select("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url").
			Joins("LEFT JOIN user_accounts ON challenge_records.user_id = user_accounts.username").
			Group("challenge_records.user_id, user_accounts.display_name, user_accounts.avatar_url").
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

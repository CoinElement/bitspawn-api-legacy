/*

 */

package models

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/vmihailenco/msgpack"
	"gorm.io/datatypes"
)

type ChallengeBase struct {
	ChallengeID uuid.UUID  `json:"challengeId" gorm:"type:uuid;primary_key;"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `sql:"index" json:"-"`

	ChallengeName    string `json:"challengeName" msgpack:"1"`
	IsSponsored      bool   `json:"isSponsored" msgpack:"2"`
	ChallengeType    string `json:"challengeType" msgpack:"3"`
	FeeType          string `json:"feeType" msgpack:"4"`
	MinuteTimeWindow int    `json:"minuteTimeWindow" msgpack:"5"`

	MinParticipants int64     `json:"minParticipants" msgpack:"6"`
	ChallengeRule   string    `json:"challengeRule" msgpack:"7" gorm:"size:65535"`
	EntryFee        string    `json:"entryFee" msgpack:"8"`
	NumberOfWinners int       `json:"numberOfWinners" msgpack:"9"`
	PrizeAllocation string    `json:"prizeAllocation" msgpack:"10"`
	CutoffDate      time.Time `json:"cutoffDate" msgpack:"11"`
	StartDate       time.Time `json:"startDate" msgpack:"12"`

	OrganizerID         string `json:"-" msgpack:"13"`
	BannerUrl           string `json:"bannerUrl" msgpack:"14"`
	ContractAddress     string `json:"-" msgpack:"15"`
	Status              string `json:"status" msgpack:"16"`
	ParticipantCount    int    `json:"participantCount" msgpack:"17"`
	TargetPrizePool     int64  `json:"targetPrizePool" msgpack:"18"`
	FeePercentage       int64  `json:"feePercentage" msgpack:"19"`
	OrganizerPercentage int64  `json:"organizerPercentage" msgpack:"20"`
	GameMode            string `json:"gameMode" msgpack:"21"`
	NumberOfGames       int    `json:"numberOfGames" msgpack:"22"`
	EntryOnce           bool   `json:"entryOnce" msgpack:"23"`
	Title               string `json:"title" msgpack:"24"`
	OrganizerContribute int64  `json:"organizerContribute" msgpack:"25"`
	MaxParticipants     int    `json:"maxParticipants" msgpack:"26"`
}

type Challenge struct {
	ChallengeBase
	Scoring map[string]*string `json:"scoring" msgpack:"4"`
}

type ChallengeData struct {
	ChallengeBase
	Scoring datatypes.JSON
}

type ChallengeResponse struct {
	ChallengeID      string    `json:"challengeId"`
	ChallengeName    string    `json:"challengeName"`
	IsSponsored      bool      `json:"isSponsored"`
	ChallengeType    string    `json:"challengeType"`
	MinuteTimeWindow int       `json:"minuteTimeWindow"`
	MinParticipants  int64     `json:"minParticipants"`
	ChallengeRule    string    `json:"challengeRule"`
	EntryFee         int64     `json:"entryFee"`
	FeeType          string    `json:"feeType"`
	PrizeAllocation  []int64   `json:"prizeAllocation"`
	EndDate          time.Time `json:"endDate"`
	StartDate        time.Time `json:"startDate"`
	BannerUrl        string    `json:"bannerUrl"`
	Status           string    `json:"status"`
	ParticipantCount int       `json:"participantCount"`
	MinPrizePool     int64     `json:"minPrizePool"`
	GameMode         string    `json:"gameMode"`
	NumberOfGames    int       `json:"numberOfGames"`
	EntryOnce        bool      `json:"entryOnce"`
	Url              string    `json:"url"`

	FeePercentage       int64           `json:"feePercentage"`
	OrganizerPercentage int64           `json:"organizerPercentage"`
	Scoring             json.RawMessage `json:"scoring"`
	FundContribute      int64           `json:"fundContribute"`
	MaxParticipants     int             `json:"maxParticipants"`
}

func (td *Challenge) GetData() ([]byte, error) {
	return msgpack.Marshal(td)
}

func NewChallengeData(d []byte) (*Challenge, error) {
	td := Challenge{}
	err := msgpack.Unmarshal(d, &td)
	return &td, err
}

func (db *DB) CreateChallengeData(td *Challenge) error {
	scoring_json, err := json.Marshal(td.Scoring)
	if err != nil {
		return err
	}
	challengeData := ChallengeData{
		ChallengeBase: td.ChallengeBase,
		Scoring:       datatypes.JSON(scoring_json),
	}
	return db.Create(&challengeData).Error
}

func (db *DB) GetChallengeData(id string) (*Challenge, error) {
	challengeID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	t := ChallengeData{}
	err = db.Where("challenge_id = ?", challengeID).First(&t).Error
	if err != nil {
		return nil, err
	}
	scoring := make(map[string]*string)
	err = json.Unmarshal([]byte(t.Scoring), &scoring)
	if err != nil {
		return nil, err
	}
	challenge := Challenge{
		ChallengeBase: t.ChallengeBase,
		Scoring:       scoring,
	}
	return &challenge, err
}

func (db *DB) GetChallengeForResponse(id string) (*ChallengeResponse, error) {
	challengeID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	t := ChallengeData{}
	err = db.Where("challenge_id = ?", challengeID).First(&t).Error
	if err != nil {
		return nil, err
	}
	challengeResponse, err := db.formatChallengesForResponse([]ChallengeData{t})
	if err != nil {
		return nil, err
	}
	return &challengeResponse[0], nil
}

func (db *DB) GetSponsoredChallenges() ([]ChallengeResponse, error) {
	challenges := []ChallengeData{}
	err := db.Where("is_sponsored = ?", true).
		Where("status = ? OR status = ?", "Registration", "Ready").
		Order("start_date").Find(&challenges).Error
	if err != nil {
		return nil, err
	}
	return db.formatChallengesForResponse(challenges)
}

func (db *DB) GetFeaturedChallenges() ([]ChallengeResponse, error) {
	challenges := []ChallengeData{}
	err := db.Table("challenge_data").
		Joins("Inner Join featureds ON challenge_data.challenge_id::TEXT=featureds.object_id").
		Where("status = ? OR status = ?", "Registration", "Ready").
		Order("start_date").Find(&challenges).Error
	if err != nil {
		return nil, err
	}
	return db.formatChallengesForResponse(challenges)
}

func (db *DB) GetRegisteredChallenges(sub string) ([]ChallengeResponse, error) {
	challenges := []ChallengeData{}
	err := db.Table("challenge_data").
		Joins("Inner Join challenge_records ON challenge_data.challenge_id=challenge_records.challenge_id").
		Where("user_id = ? AND status != 'Cancelled'", sub).Order("start_date").
		Find(&challenges).
		Error
	if err != nil {
		return nil, err
	}
	return db.formatChallengesForResponse(challenges)
}

func (db *DB) GetChallengesByStatus(status string) ([]Challenge, error) {
	ts := []ChallengeData{}
	err := db.DB.Where("status = ?", status).Find(&ts).Error
	if err != nil {
		return nil, err
	}
	challenges := []Challenge{}
	for _, t := range ts {
		scoring := make(map[string]*string)
		err = json.Unmarshal([]byte(t.Scoring), &scoring)
		if err != nil {
			return nil, err
		}
		challenge := Challenge{
			ChallengeBase: t.ChallengeBase,
			Scoring:       scoring,
		}
		challenges = append(challenges, challenge)
	}
	return challenges, err
}

func (db *DB) UpdateChallengeOrganizerContribute(id uuid.UUID, organizerContribute int64) error {
	return db.Model(&ChallengeData{}).
		Where("challenge_id = ?", id).
		Update("organizer_contribute", organizerContribute).
		Error
}

func (db *DB) UpdateChallengeState(id uuid.UUID, status string) error {
	return db.Model(&ChallengeData{}).
		Where("challenge_id = ?", id).
		Update("status", status).
		Error
}

func (db *DB) UpdateChallengeBanner(challenge_id, banner_url string) (*Challenge, error) {
	challengeInfo, err := db.GetChallengeData(challenge_id)
	if err != nil {
		return nil, err
	}
	return nil, db.Model(challengeInfo).Updates(map[string]interface{}{"banner_url": banner_url}).Error
}

func (db *DB) formatChallengesForResponse(challengeData []ChallengeData) ([]ChallengeResponse, error) {
	var challengeResponse []ChallengeResponse
	for _, td := range challengeData {
		entryFeeInt, err := strconv.ParseInt(td.EntryFee, 10, 64)
		if err != nil {
			return nil, err
		}

		var prizeAllocationIntArray []int64
		if td.PrizeAllocation == "" {
			// do nothing
		} else {
			s := strings.Split(td.PrizeAllocation, ",")
			for _, prize := range s {
				prizeAllocationInt, err := strconv.ParseInt(prize, 10, 64)
				if err != nil {
					return nil, err
				}
				prizeAllocationIntArray = append(prizeAllocationIntArray, prizeAllocationInt)
			}
		}

		cr := ChallengeResponse{
			ChallengeID:      td.ChallengeID.String(),
			ChallengeName:    td.ChallengeName,
			IsSponsored:      td.IsSponsored,
			ChallengeType:    strings.ToUpper(td.ChallengeType),
			MinuteTimeWindow: td.MinuteTimeWindow,
			MinParticipants:  td.MinParticipants,
			ChallengeRule:    td.ChallengeRule,
			EntryFee:         entryFeeInt,
			FeeType:          td.FeeType,
			PrizeAllocation:  prizeAllocationIntArray,
			EndDate:          td.CutoffDate,
			StartDate:        td.StartDate,
			BannerUrl:        td.BannerUrl,
			Status:           strings.ToUpper(td.Status),
			ParticipantCount: td.ParticipantCount,
			MinPrizePool:     td.TargetPrizePool,
			GameMode:         strings.ToUpper(td.GameMode),
			NumberOfGames:    td.NumberOfGames,
			EntryOnce:        td.EntryOnce,
			Url:              td.Title,

			FeePercentage:       td.FeePercentage,
			OrganizerPercentage: td.OrganizerPercentage,
			Scoring:             json.RawMessage(td.Scoring),
			FundContribute:      td.OrganizerContribute,
			MaxParticipants:     td.MaxParticipants,
		}
		challengeResponse = append(challengeResponse, cr)
	}
	return challengeResponse, nil
}

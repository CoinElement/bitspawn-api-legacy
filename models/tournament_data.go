/*

 */

package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/bitspawngg/bitspawn-api/enum"
	"math"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/vmihailenco/msgpack"
	"gorm.io/datatypes"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

type TournamentData struct {
	TournamentID string     `json:"tournamentId" gorm:"type:uuid;primary_key;"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updateAt"`
	DeletedAt    *time.Time `sql:"index" json:"-"`

	TournamentName        string                `json:"tournamentName" msgpack:"1"`
	IsSponsored           bool                  `json:"isFeatured" msgpack:"2"`
	GameType              string                `json:"gameType" msgpack:"3"`
	GameSubtype           string                `json:"gameSubtype" msgpack:"4"`
	TournamentFormat      enum.TournamentFormat `json:"tournamentFormat" msgpack:"5"`
	MaxParticipants       int                   `json:"maxParticipants" msgpack:"6"`
	TournamentDescription string                `json:"tournamentDescription" msgpack:"7" gorm:"size:65535"`
	TournamentRule        string                `json:"tournamentRule" msgpack:"8" gorm:"size:65535"`
	EntryFee              string                `json:"entryFee" msgpack:"9"`
	NumberOfWinners       int64                 `json:"numberOfWinners" msgpack:"10"`
	PrizeAllocation       string                `json:"prizeAllocation" msgpack:"11"`
	CutoffDate            time.Time             `json:"cutoffDate" msgpack:"12"`
	TournamentDate        time.Time             `json:"tournamentDate" msgpack:"13"`

	OrganizerID      string         `json:"-" msgpack:"14"`
	Roles            string         `json:"moderators" msgpack:"15"`
	BannerUrl        string         `json:"bannerUrl" msgpack:"16"`
	ThumbnailUrl     string         `json:"thumbnailUrl" msgpack:"17"`
	ContractAddress  string         `json:"-" msgpack:"18"`
	Status           string         `json:"status" msgpack:"19"`
	ParticipantCount int64          `json:"participantCount" msgpack:"20"`
	TotalPrizePool   int64          `json:"totalPrizePool" msgpack:"21"`
	Blocked          datatypes.JSON `json:"blocked" sql:"type:jsonb" msgpack:"22"`
	InviteOnly       bool           `json:"inviteOnly" msgpack:"23"`
	IsManual         bool           `json:"isManual" msgpack:"24"`

	RoundsFormat datatypes.JSON `json:"roundsFormat" msgpack:"25"`
	Metadata     datatypes.JSON `json:"metadata" msgpack:"26"`

	Organizer           Organizer `json:"organizer,omitempty" msgpack:"28" gorm:"-"`
	OrganizerContribute int64     `json:"fundContribute" msgpack:"29"`

	Sponsors            string `json:"sponsors"`
	Consoles            string `json:"consoles"`
	MinPrizePool        int64  `json:"minPrizePool"`
	FeePercentage       int64  `json:"feePercentage"`
	OrganizerPercentage int64  `json:"organizerPercentage"`
	MinParticipants     int    `json:"minParticipants"`
	NumberOfTeams       int    `json:"numberOfTeams"`
	LogoUrl             string `json:"logoUrl"`
	FeeType             string `json:"feeType"`
	CriticalRules       string `json:"criticalRules"`
	RoundsBestOfN       string `json:"roundsBestOfN"`
	CheckInType         string `json:"checkInType"`
	CheckInBeforeMin    int64  `json:"checkInBeforeMin"`
	MatchDurationMin    int64  `json:"matchDurationMin"`
	// this is used for round-robin
	ParticipantsPlayEachOther int `json:"participantsPlayEachOther"`
}

type TournamentResponse struct {
	TournamentID string    `json:"tournamentId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	OrganizerID  string    `json:"-"`

	TournamentName        string                `json:"tournamentName"`
	GameType              string                `json:"gameType"`
	GameSubtype           string                `json:"gameSubtype"`
	TournamentFormat      enum.TournamentFormat `json:"tournamentFormat"`
	TournamentDescription string                `json:"tournamentDescription"`
	TournamentRule        string                `json:"tournamentRule"`
	CriticalRules         string                `json:"criticalRules"`
	OrganizerPercentage   int64                 `json:"organizerPercentage"`
	FeePercentage         int64                 `json:"feePercentage"`

	MaxParticipants int    `json:"maxParticipants"`
	MinParticipants int    `json:"minParticipants"`
	MinPrizePool    int64  `json:"minPrizePool"`
	NumberOfTeams   int    `json:"numberOfTeams"`
	EntryFee        int64  `json:"entryFee"`
	FeeType         string `json:"feeType"`
	InviteOnly      bool   `json:"inviteOnly"`

	CutoffDate     time.Time `json:"cutoffDate"`
	TournamentDate time.Time `json:"tournamentDate"`

	BannerUrl    string `json:"bannerUrl"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	LogoUrl      string `json:"logoUrl"`

	Organizer      Organizer `json:"organizer,omitempty"`
	FundContribute int64     `json:"fundContribute,omitempty"`

	Consoles         []enum.Console `json:"consoles"`
	MatchDurationMin int64          `json:"matchDurationMin"`
	CheckInType      string         `json:"checkInType"`
	CheckInBeforeMin int64          `json:"checkInBeforeMin"`
	Metadata         datatypes.JSON `json:"metadata"`

	IsFeatured      bool        `json:"isFeatured"`
	Sponsors        []string    `json:"sponsors"`
	Moderators      []Organizer `json:"moderators"`
	PrizeAllocation []int64     `json:"prizeAllocation"`

	Status                    string `json:"status"`
	ParticipantCount          int64  `json:"participantCount"`
	TotalPrizePool            int64  `json:"totalPrizePool"`
	ParticipantsPlayEachOther int    `json:"participantsPlayEachOther"`
}

type Organizer struct {
	DisplayName string `json:"displayName"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (td *TournamentData) GetData() ([]byte, error) {
	return msgpack.Marshal(td)
}

func NewTournamentData(d []byte) (*TournamentData, error) {

	td := TournamentData{}
	err := msgpack.Unmarshal(d, &td)

	return &td, err
}

func (db *DB) CreateTournament(t *TournamentResponse) (*TournamentResponse, error) {
	t.TournamentID = uuid.NewV4().String()
	t.CreatedAt = time.Now().UTC()
	t.UpdatedAt = time.Now().UTC()
	t.IsFeatured = false
	t.PrizeAllocation = []int64{100}
	t.Moderators = []Organizer{}
	t.Sponsors = []string{}
	t.FeePercentage = 5
	t.TotalPrizePool = 0
	t.FundContribute = 0
	t.Status = "DRAFT"
	td := TournamentData{
		TournamentID:              t.TournamentID,
		TournamentName:            t.TournamentName,
		IsSponsored:               t.IsFeatured,
		GameType:                  t.GameType,
		GameSubtype:               t.GameSubtype,
		TournamentFormat:          t.TournamentFormat,
		TournamentDescription:     t.TournamentDescription,
		TournamentRule:            t.TournamentRule,
		CriticalRules:             t.CriticalRules,
		FeePercentage:             t.FeePercentage,
		OrganizerPercentage:       t.OrganizerPercentage,
		MaxParticipants:           t.MaxParticipants,
		MinParticipants:           t.MinParticipants,
		MinPrizePool:              t.MinPrizePool,
		NumberOfTeams:             t.NumberOfTeams,
		EntryFee:                  strconv.FormatInt(t.EntryFee, 10),
		FeeType:                   t.FeeType,
		PrizeAllocation:           "100",
		CutoffDate:                t.CutoffDate,
		TournamentDate:            t.TournamentDate,
		OrganizerID:               t.OrganizerID,
		TotalPrizePool:            t.TotalPrizePool,
		OrganizerContribute:       t.FundContribute,
		InviteOnly:                t.InviteOnly,
		IsManual:                  true,
		Status:                    t.Status,
		Metadata:                  t.Metadata,
		Consoles:                  enum.ConsoleJoin(t.Consoles),
		MatchDurationMin:          t.MatchDurationMin,
		CheckInType:               t.CheckInType,
		CheckInBeforeMin:          t.CheckInBeforeMin,
		BannerUrl:                 t.BannerUrl,
		ThumbnailUrl:              t.ThumbnailUrl,
		LogoUrl:                   t.LogoUrl,
		ParticipantsPlayEachOther: t.ParticipantsPlayEachOther,
	}
	err := db.Create(&td).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (db *DB) UpdateTournamentData(tournamentId string, td *TournamentData) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	err := tx.Model(&TournamentData{}).Where("tournament_id = ?", tournamentId).Updates(&td).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&TournamentData{}).
		Where("tournament_id = ?", tournamentId).
		Select("invite_only").Updates(&td).
		Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) UpdateTournament(tournamentId string, t *TournamentResponse) error {
	prizeString := []string{}
	for _, teamPrize := range t.PrizeAllocation {
		p := fmt.Sprintf("%d", teamPrize)
		prizeString = append(prizeString, p)
	}
	td := TournamentData{
		TournamentName:        t.TournamentName,
		GameType:              t.GameType,
		GameSubtype:           t.GameSubtype,
		TournamentFormat:      t.TournamentFormat,
		TournamentDescription: t.TournamentDescription,
		TournamentRule:        t.TournamentRule,
		CriticalRules:         t.CriticalRules,
		OrganizerPercentage:   t.OrganizerPercentage,

		MaxParticipants: t.MaxParticipants,
		MinParticipants: t.MinParticipants,
		MinPrizePool:    t.MinPrizePool,
		NumberOfTeams:   t.NumberOfTeams,
		EntryFee:        strconv.FormatInt(t.EntryFee, 10),
		FeeType:         t.FeeType,
		PrizeAllocation: strings.Join(prizeString, ","),
		InviteOnly:      t.InviteOnly,

		CutoffDate:     t.CutoffDate,
		TournamentDate: t.TournamentDate,

		Metadata:         t.Metadata,
		Consoles:         enum.ConsoleJoin(t.Consoles),
		Sponsors:         strings.Join(t.Sponsors, ","),
		MatchDurationMin: t.MatchDurationMin,
		CheckInType:      t.CheckInType,
		CheckInBeforeMin: t.CheckInBeforeMin,
	}
	return db.UpdateTournamentData(tournamentId, &td)
}

func (db *DB) UpdateTournamentStatus(tournamentId string, status string) error {
	return db.Model(&TournamentData{}).
		Where("tournament_id = ?", tournamentId).
		Update("status", status).
		Error
}

func (db *DB) UpdateTournamentBlockedData(tournamentId string, blocked JSONB) error {
	return db.Exec("UPDATE public.tournament_data SET blocked = jsonb_set(blocked, , $1) WHERE tournament_id = $2", blocked, tournamentId).Error
}

func (db *DB) GetTournamentData(id string) (*TournamentData, error) {
	tournamentID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	t := TournamentData{}
	err = db.Where("tournament_id = ?", tournamentID).First(&t).Error
	if err != nil {
		return nil, err
	}
	user, err := db.GetUserProfile(t.OrganizerID)
	if err != nil {
		return nil, err
	}
	t.Organizer.DisplayName = user.DisplayName
	t.Organizer.AvatarUrl = user.AvatarUrl
	return &t, nil
}

func (db *DB) GetTournamentResponse(id string) (*TournamentResponse, error) {
	td, err := db.GetTournamentData(id)
	if err != nil {
		return nil, err
	}
	return db.FormatTournamentForResponse(*td)
}

func (db *DB) ListTournamentsByStatus(tournament_name string, game_type string, game_subtype string, total_prize_pool int, sortMethod string, start_time string, end_time string, page, per_page int, status []string) ([]TournamentData, error) {
	ts := []TournamentData{}
	offset := (page - 1) * per_page
	limit := per_page

	var rows *sql.Rows
	var err error
	if game_type != "" {
		if tournament_name != "%%" && game_subtype != "" {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND (tournament_name LIKE (?)) AND game_type = ? AND game_subtype = ? AND total_prize_pool >= ?", status, tournament_name, game_type, game_subtype, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		} else if tournament_name != "%%" && game_subtype == "" {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND (tournament_name LIKE (?)) AND game_type = ? AND total_prize_pool >= ?", status, tournament_name, game_type, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		} else if tournament_name == "%%" && game_subtype != "" {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND game_type = ? AND game_subtype = ? AND total_prize_pool >= ?", status, game_type, game_subtype, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		} else {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND game_type = ?  AND total_prize_pool >= ?", status, game_type, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		}
	} else {
		if tournament_name != "%%" && game_subtype != "" {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND (tournament_name LIKE (?)) AND game_subtype = ? AND total_prize_pool >= ?", status, tournament_name, game_subtype, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		} else if tournament_name != "%%" && game_subtype == "" {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND (tournament_name LIKE (?)) AND total_prize_pool >= ?", status, tournament_name, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		} else if tournament_name == "%%" && game_subtype != "" {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND game_subtype = ? AND total_prize_pool >= ?", status, game_subtype, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		} else {
			rows, err = db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
				Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
				Where("status IN (?) AND total_prize_pool >= ?", status, total_prize_pool).
				Order(sortMethod).
				Limit(limit).
				Offset(offset).
				Rows()
		}
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tournament TournamentData
		var organizer Organizer
		_ = db.ScanRows(rows, &tournament)
		_ = db.ScanRows(rows, &organizer)
		tournament.Organizer = organizer
		ts = append(ts, tournament)
	}
	return ts, nil
}

func (db *DB) CountTournamentsByStatus(game_type string, status []string) (int64, error) {
	var count int64
	var err error
	if game_type != "" {
		err = db.Table("tournament_data").
			Where("tournament_data.status IN (?) AND game_type = ? ", status, game_type).
			Count(&count).Error
	} else {
		err = db.Table("tournament_data").
			Where("tournament_data.status IN (?)", status).
			Count(&count).Error
	}
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) CountOpenTournaments(game_type string) (int64, error) {
	var count int64
	var err error
	statusRegistration := []string{"REGISTRATION", "Registration"}
	if game_type != "" {
		err = db.Table("tournament_data").
			Where("tournament_data.status IN (?) AND game_type = ? AND cutoff_date > now()", statusRegistration, game_type).
			Count(&count).Error
	} else {
		err = db.Table("tournament_data").
			Where("tournament_data.status IN (?) AND cutoff_date > now()", statusRegistration).
			Count(&count).Error
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) CountCutoffTournaments(game_type string) (int64, error) {
	var count int64
	var err error
	statusRegistration := []string{"REGISTRATION", "Registration"}
	if game_type != "" {
		err = db.Table("tournament_data").
			Where("tournament_data.status IN (?) AND game_type = ? AND cutoff_date < now()", statusRegistration, game_type).
			Count(&count).Error
	} else {
		err = db.Table("tournament_data").
			Where("tournament_data.status IN (?) AND cutoff_date < now()", statusRegistration).
			Count(&count).Error
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetUserOrganisedTournaments(username string) ([]TournamentResponse, error) {
	ts := []TournamentResponse{}

	rows, err := db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
		Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
		Where("status != ?", "Cancelled").
		Where("organizer_id = ?", username).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tournament TournamentData
		var organizer Organizer
		_ = db.ScanRows(rows, &tournament)
		_ = db.ScanRows(rows, &organizer)
		tournamentResponse, err := db.FormatTournamentForResponse(tournament)
		if err != nil {
			return nil, err
		}
		tournamentResponse.Organizer = organizer
		ts = append(ts, *tournamentResponse)
	}
	return ts, nil
}

func (db *DB) GetUserModeratorTournaments(username string, page, per_page int) ([]TournamentData, error) {
	ts := []TournamentData{}
	offset := (page - 1) * per_page
	limit := per_page

	rows, err := db.Table("tournament_data").Joins("LEFT JOIN user_accounts ON tournament_data.organizer_id = user_accounts.username").
		Select("tournament_data.*, user_accounts.display_name, user_accounts.avatar_url").
		Where("(roles LIKE (?))", username).Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tournament TournamentData
		var organizer Organizer
		_ = db.ScanRows(rows, &tournament)
		_ = db.ScanRows(rows, &organizer)
		tournament.Organizer = organizer
		ts = append(ts, tournament)
	}
	return ts, nil
}

func (db *DB) GetTournamentsByStatus(status string) ([]TournamentData, error) {
	td := []TournamentData{}
	err := db.DB.Where("status = ?", status).Order("tournament_date").Find(&td).Error
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (db *DB) GetTournamentsOverdue() ([]TournamentData, error) {
	td := []TournamentData{}
	err := db.DB.Where("(status) = ? AND (cutoff_date) < now() AND is_manual = ?", "REGISTRATION", false).Order("tournament_date").Find(&td).Error
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (db *DB) UpdateTournamentState(id, status string, participantCount, totalPrizePool int64) error {
	return db.Model(&TournamentData{}).
		Where("tournament_id = ?", id).
		Updates(map[string]interface{}{"status": status, "participant_count": participantCount, "total_prize_pool": totalPrizePool}).
		Error
}

func (db *DB) UpdateTournamentBanner(tournament_id, banner_url string) (*TournamentResponse, error) {
	tournamentInfo, err := db.GetTournamentData(tournament_id)
	if err != nil {
		return nil, err
	}
	tournamentInfo.BannerUrl = banner_url
	err = db.Model(tournamentInfo).Updates(map[string]interface{}{"banner_url": banner_url}).Error
	if err != nil {
		return nil, err
	}
	return db.FormatTournamentForResponse(*tournamentInfo)
}

func (db *DB) UpdateTournamentLogo(tournament_id, logo_url string) (*TournamentResponse, error) {
	tournamentInfo, err := db.GetTournamentData(tournament_id)
	if err != nil {
		return nil, err
	}
	tournamentInfo.LogoUrl = logo_url
	err = db.Model(tournamentInfo).Updates(map[string]interface{}{"logo_url": logo_url}).Error
	if err != nil {
		return nil, err
	}
	return db.FormatTournamentForResponse(*tournamentInfo)
}

func (db *DB) UpdateTournamentThumbnail(tournament_id, thumbnail_url string) (*TournamentResponse, error) {
	tournamentInfo, err := db.GetTournamentData(tournament_id)
	if err != nil {
		return nil, err
	}
	tournamentInfo.ThumbnailUrl = thumbnail_url
	err = db.Model(tournamentInfo).Updates(map[string]interface{}{"thumbnail_url": thumbnail_url}).Error
	if err != nil {
		return nil, err
	}
	return db.FormatTournamentForResponse(*tournamentInfo)
}

func (db *DB) GetMyDraftTournaments(username string) ([]TournamentResponse, error) {
	tournaments := []TournamentData{}
	err := db.Where(&TournamentData{OrganizerID: username, Status: "DRAFT"}).Find(&tournaments).Error
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

func (db *DB) CountMyTournamentPages(username string, per_page int) (int64, error) {
	var count int64
	err := db.Table("tournament_data").
		Where("organizer_id = ?", username).
		Where("draft_mode = ?", false).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	pagefloat := math.Ceil(float64(count) / float64(per_page))
	page := int64(pagefloat)

	return page, nil
}

func (db *DB) CountMyModeratorTournamentPages(username string, per_page int) (int64, error) {
	var count int64
	err := db.Table("tournament_data").
		Where("(roles LIKE (?))", username).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	pagefloat := math.Ceil(float64(count) / float64(per_page))
	page := int64(pagefloat)

	return page, nil
}

func (db *DB) CountPagesByStatus(tournament_name string, game_type string, per_page int, status []string) (int64, error) {
	var count int64
	var err error
	if game_type != "" {
		if tournament_name != "%%" {
			err = db.Table("tournament_data").
				Where("status IN (?) AND (tournament_name LIKE (?)) AND game_type = ? ", status, tournament_name, game_type).
				Count(&count).Error
		} else {
			err = db.Table("tournament_data").
				Where("status IN (?) AND game_type = ? ", status, game_type).
				Count(&count).Error
		}
	} else {
		if tournament_name != "%%" {
			err = db.Table("tournament_data").
				Where("status IN (?) AND (tournament_name LIKE (?))", status, tournament_name).
				Count(&count).Error
		} else {
			err = db.Table("tournament_data").
				Where("status IN (?)", status).
				Count(&count).Error
		}
	}

	if err != nil {
		return 0, err
	}
	pagefloat := math.Ceil(float64(count) / float64(per_page))
	page := int64(pagefloat)

	return page, nil
}

func (db DB) FormatTournamentForResponse(td TournamentData) (*TournamentResponse, error) {
	entryFeeInt, err := strconv.ParseInt(td.EntryFee, 10, 64)
	if err != nil {
		return nil, err
	}

	var prizeAllocationIntArray []int64
	if td.PrizeAllocation == "" {
		prizeAllocationIntArray = append(prizeAllocationIntArray, 100)
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

	sponsorSlice := []string{}
	if td.Sponsors != "" {
		sponsorSlice = strings.Split(td.Sponsors, ",")
	}

	consoleSlice := []enum.Console{}
	if td.Consoles != "" {

		for _, Consoles := range strings.Split(td.Consoles, ",") {
			enumConsole := enum.Console(Consoles)
			if enumConsole.IsValid() {
				consoleSlice = append(consoleSlice, enumConsole)
			}
		}
	}

	moderatorSlice := []Organizer{}
	if td.Roles != "" {
		for _, username := range strings.Split(td.Roles, ",") {
			if username == "" {
				continue
			}
			player, err := db.GetUserProfile(username)
			if err != nil {
				return nil, err
			}
			moderator := Organizer{
				DisplayName: player.DisplayName,
				AvatarUrl:   player.AvatarUrl,
			}
			moderatorSlice = append(moderatorSlice, moderator)
		}
	}
	tr := TournamentResponse{
		TournamentID:              td.TournamentID,
		CreatedAt:                 td.CreatedAt,
		UpdatedAt:                 td.UpdatedAt,
		TournamentName:            td.TournamentName,
		IsFeatured:                td.IsSponsored,
		Sponsors:                  sponsorSlice,
		GameType:                  td.GameType,
		GameSubtype:               td.GameSubtype,
		TournamentFormat:          td.TournamentFormat,
		MaxParticipants:           td.MaxParticipants,
		TournamentDescription:     td.TournamentDescription,
		TournamentRule:            td.TournamentRule,
		CriticalRules:             td.CriticalRules,
		EntryFee:                  entryFeeInt,
		FeeType:                   td.FeeType,
		PrizeAllocation:           prizeAllocationIntArray,
		CutoffDate:                td.CutoffDate,
		TournamentDate:            td.TournamentDate,
		Moderators:                moderatorSlice,
		BannerUrl:                 td.BannerUrl,
		ThumbnailUrl:              td.ThumbnailUrl,
		LogoUrl:                   td.LogoUrl,
		Status:                    td.Status,
		ParticipantCount:          td.ParticipantCount,
		MinParticipants:           td.MinParticipants,
		TotalPrizePool:            td.TotalPrizePool,
		MinPrizePool:              td.MinPrizePool,
		FeePercentage:             td.FeePercentage,
		OrganizerPercentage:       td.OrganizerPercentage,
		Consoles:                  consoleSlice,
		InviteOnly:                td.InviteOnly,
		Metadata:                  td.Metadata,
		Organizer:                 td.Organizer,
		FundContribute:            td.OrganizerContribute,
		NumberOfTeams:             td.NumberOfTeams,
		CheckInType:               td.CheckInType,
		CheckInBeforeMin:          td.CheckInBeforeMin,
		ParticipantsPlayEachOther: td.ParticipantsPlayEachOther,
	}

	return &tr, nil
}

func (db DB) GetTournament(tournamentID string) (*TournamentData, error) {
	td := TournamentData{}
	err := db.DB.Where("tournament_id = ?", tournamentID).First(&td).Error
	if err != nil {
		return nil, err
	}
	return &td, err
}

func (db DB) GetTournamentsToStart() ([]string, error) {
	td := []TournamentData{}
	statusReady := "READY"
	err := db.DB.Where("status = ?", statusReady).Where("is_manual = ?", false).Order("tournament_date").Find(&td).Error
	if err != nil {
		return nil, err
	}
	var tournamentsToStart []string
	for _, t := range td {
		tournamentsToStart = append(tournamentsToStart, t.TournamentID)
	}
	return tournamentsToStart, nil
}

func (db DB) GetTournamentsToComplete() ([]string, error) {
	td := []TournamentData{}
	statusStarted := "Started"
	err := db.DB.Where("status = ?", statusStarted).Order("tournament_date").Find(&td).Error
	if err != nil {
		return nil, err
	}
	var tournamentsToComplete []string
	for _, t := range td {
		tournamentsToComplete = append(tournamentsToComplete, t.TournamentID)
	}
	return tournamentsToComplete, nil
}

func (db DB) StartTournament(tournamentID string) error {
	td := TournamentData{}
	err := db.DB.Where("tournament_id = ?", tournamentID).First(&td).Error
	if err != nil {
		return err
	}
	return db.DB.Model(&td).Update("status", "Started").Error
}

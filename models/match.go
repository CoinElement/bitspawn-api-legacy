/*
 */

package models

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type MatchRow struct {
}

type Match struct {
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	MatchID      string    `json:"matchId" gorm:"type:uuid;primarykey"`
	TournamentID string    `json:"tournamentId" gorm:"type:uuid"`
	Round        int       `json:"round"`
	Table        int       `json:"table"`
	TeamOne      string    `json:"teamOne"`
	TeamTwo      string    `json:"teamTwo"`
	Status       string    `json:"status"`
	Result       int       `json:"result"` // 1 if Player One wins, 2 if Player Two wins, -1 if no winner
	MatchDate    time.Time `json:"matchDate"`
	BotID        string    `json:"botId"`
	BestOfN      int       `json:"bestOfN"`
	TeamOneScore int       `json:"teamOneScore"`
	TeamTwoScore int       `json:"teamTwoScore"`

	Teams1 Team `json:"teams1" gorm:"-"`
	Teams2 Team `json:"teams2" gorm:"-"`
}

type MatchOutput struct {
	MatchID      string     `json:"matchId"`
	TournamentID string     `json:"tournamentId"`
	Round        int        `json:"round"`
	Table        int        `json:"table"`
	Status       string     `json:"-"`
	Result       int        `json:"result"` // 1 if Player One wins, 2 if Player Two wins, -1 if no winner
	MatchDate    time.Time  `json:"matchDate"`
	Winner       string     `json:"-"`
	TeamOne      TeamOutput `json:"teamOne"`
	TeamTwo      TeamOutput `json:"teamTwo"`
	BestOfN      int        `json:"bestofN"`
	TeamOneScore int        `json:"teamOneScore"`
	TeamTwoScore int        `json:"teamTwoScore"`
}

type Team struct {
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	TournamentID string    `json:"tournamentId" gorm:"type:uuid"`
	ID           uint      `gorm:"primarykey"`
	TeamID       string    `json:"teamId"`
	TeamName     string    `json:"teamName"`
	Player       string    `json:"player"`
}

type TeamDTO struct {
	TeamID  string
	Name    string
	Members []string
	Slot    int
}

type Bot struct {
	SteamID      string `json:"steamId" gorm:"primarykey"`
	SteamAccount string `json:"steamAccount"`
	Password     string `json:"password"`
	Status       string `json:"status"`
}

func (db DB) CreateMatches(matches []Match) error {
	return db.DB.Create(matches).Error
}

func (db DB) CreateTeams(teams []Team) error {
	return db.DB.Create(teams).Error
}

func (db DB) GetTeams(teamId string) ([]Team, error) {
	teams := []Team{}
	err := db.DB.Where("team_id = ?", teamId).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (db DB) DeleteTeams(teamId string, playerUsernames []string) error {
	return db.DB.Where("team_id = ? AND player in ?", teamId, playerUsernames).Delete(&Team{}).Error
}

func (db DB) GetMatch(tournamentId string, round, table int) (*Match, error) {
	match := Match{}
	err := db.DB.Where(`"tournament_id" = ? AND "round" = ? AND "table" = ?`, tournamentId, round, table).First(&match).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

func (db DB) GetMatchById(matchId string) (*Match, error) {
	match := Match{}
	err := db.DB.Where(Match{MatchID: matchId}).First(&match).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

func (db DB) GetMatchesByTournament(tournamentId string) ([]Match, error) {
	matches := make([]Match, 0)
	err := db.DB.Order(`round,"table"`).Where("tournament_id = ?", tournamentId).Find(&matches).Error
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (db *DB) DeleteMatchesByTournamentId(tournamentId string) error {
	matches, err := db.GetMatchesByTournament(tournamentId)
	if err != nil {
		return err
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	for _, m := range matches {
		if _, err := uuid.FromString(m.TeamOne); err == nil {
			if err := tx.Where("team_id = ?", m.TeamOne).Delete(Team{}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		if _, err := uuid.FromString(m.TeamTwo); err == nil {
			if err := tx.Where("team_id = ?", m.TeamTwo).Delete(Team{}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if err := tx.Table("matches").Where("tournament_id = ?", tournamentId).Delete(Match{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) UpdateMatch(match *Match) error {
	err := db.Model(&match).
		Where(`tournament_id = ? AND round = ? AND "table" = ?`, match.TournamentID, match.Round, match.Table).
		Updates(&match).Error
	return err
}

func (db *DB) UpdateRoundBestOfN(tournamentId string, round int, bestOfN int) error {
	err := db.Model(Match{}).
		Where(`tournament_id = ? AND round = ?`, tournamentId, round).
		Update("best_of_n", bestOfN).Error
	return err
}

func (db *DB) UpdateRoundMatchDate(tournamentId string, round int, matchDate time.Time) error {
	err := db.Model(Match{}).
		Where(`tournament_id = ? AND round = ?`, tournamentId, round).
		Update("match_date", matchDate).Error
	return err
}

func (db DB) GetPlayersByTeam(teamId string) ([]string, error) {
	teams := []Team{}
	players := []string{}
	err := db.DB.Where("team_id = ?", teamId).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	for _, team := range teams {
		if team.Player != "" {
			players = append(players, team.Player)
		}
	}
	return players, nil
}

func (db DB) GetAssignedParticipants(tournamentId string) ([]UserInfo, error) {
	assignedParticipants := []UserInfo{}
	var rows *sql.Rows
	var err error
	rows, err = db.Table("teams").
		Select("user_accounts.sub,user_accounts.username,user_accounts.display_name,user_accounts.avatar_url").
		Joins("LEFT JOIN user_accounts ON teams.player = user_accounts.username").
		Where("teams.tournament_id = ? AND teams.player != ''", tournamentId).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var assignedParticipant UserInfo
		_ = db.ScanRows(rows, &assignedParticipant)
		assignedParticipants = append(assignedParticipants, assignedParticipant)
	}
	return assignedParticipants, nil
}

type PlayerOnTeam struct {
	TeamID      string `json:"teamId"`
	TeamName    string `json:"teamName"`
	DisplayName string `json:"displayName"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (db DB) GetPlayersByTeamId(teamId string) ([]PlayerOnTeam, error) {
	var rows *sql.Rows
	var err error
	players := []PlayerOnTeam{}

	rows, err = db.Table("teams").
		Select("teams.team_id, teams.team_name, users.display_name, users.avatar_url").
		Joins("LEFT JOIN user_accounts AS users ON teams.player = users.username").
		Where("teams.team_id = ?", teamId).
		Rows()
	if err != nil {
		return players, err
	}
	defer rows.Close()

	for rows.Next() {
		var player PlayerOnTeam
		_ = db.ScanRows(rows, &player)
		players = append(players, player)
	}
	return players, nil
}

func (db *DB) OrganizerUpdateMatchScoresV2(tournamentId string, matchId string, teamOneScore, teamTwoScore int) error {
	match, err := db.GetMatchById(matchId)
	if err != nil {
		return fmt.Errorf("error in GetMatch: %v", err)
	}
	if match.TournamentID != tournamentId {
		return fmt.Errorf("match does not belong in this tournament")
	}
	if teamOneScore+teamTwoScore > match.BestOfN {
		return fmt.Errorf("combined scores exceed best of N limit")
	}
	match.TeamOneScore = teamOneScore
	match.TeamTwoScore = teamTwoScore
	if teamOneScore > match.BestOfN/2 || teamTwoScore > match.BestOfN/2 {
		if match.Status == "Finished" { // if the match was previously finished, adjust next round teamID
			if (teamOneScore > teamTwoScore && match.Result == 2) || (teamOneScore < teamTwoScore && match.Result == 1) { // winner changed
				var winnerTeamID string
				if teamOneScore > teamTwoScore {
					winnerTeamID = match.TeamOne
				} else {
					winnerTeamID = match.TeamTwo
				}
				nextRoundMatch, err := db.GetMatch(tournamentId, match.Round+1, (match.Table+1)/2)
				if err == nil {
					if nextRoundMatch.Status == "ManualReady" { // only adjust if next round match is still in Ready status
						if match.Table%2 == 1 { // this table number is odd, need to adjust team one of next round match
							nextRoundMatch.TeamOne = winnerTeamID
						} else {
							nextRoundMatch.TeamTwo = winnerTeamID
						}
					}
					err = db.UpdateMatch(nextRoundMatch)
					if err != nil {
						return fmt.Errorf("error in UpdateMatch: %v", err)
					}
				}
			}
		}
		match.Status = "Finished"
		if teamOneScore > teamTwoScore {
			match.Result = 1
		} else {
			match.Result = 2
		}
	}
	err = db.Save(match).Error
	if err != nil {
		return fmt.Errorf("error in Save Match: %v", err)
	}
	return nil
}

func (db *DB) SwapTeams(tournamentId string, round int, table int, teamNumber int, team string) error {
	var err error
	if teamNumber == 1 {
		err = db.Model(&Match{}).
			Where(`"tournament_id" = ? AND "round" = ? AND "table" = ?`, tournamentId, round, table).
			Updates(map[string]interface{}{"team_one": team}).
			Error
	} else if teamNumber == 2 {
		err = db.Model(&Match{}).
			Where(`"tournament_id" = ? AND "round" = ? AND "table" = ?`, tournamentId, round, table).
			Updates(map[string]interface{}{"team_two": team}).
			Error
	}
	return err
}

func (db *DB) SwapTeamsV2(tournamentId string, teamAId string, teamBId string) error {
	var err error
	var matchA, matchB Match
	err = db.Where("tournament_id = ? AND team_one = ? OR team_two = ?", tournamentId, teamAId, teamAId).First(&matchA).Error
	if err != nil {
		return err
	}
	err = db.Where("tournament_id = ? AND team_one = ? OR team_two = ?", tournamentId, teamBId, teamBId).First(&matchB).Error
	if err != nil {
		return err
	}
	if matchA.TeamOne == teamAId {
		matchA.TeamOne = teamBId
	} else {
		matchA.TeamTwo = teamBId
	}
	if matchB.TeamOne == teamBId {
		matchB.TeamOne = teamAId
	} else {
		matchB.TeamTwo = teamAId
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Save(&matchA).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&matchB).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) UpdateTeamName(teamId, teamName string) error {
	return db.Model(Team{}).Where("team_id = ?", teamId).Update("team_name", teamName).Error
}

func (db *DB) UpdateTeam(teamId string, teamName string, players []string) error {
	teams := []Team{}
	err := db.Where("team_id = ?", teamId).Find(&teams).Error
	if err != nil {
		return err
	}
	if len(players) != len(teams) {
		return fmt.Errorf("team records length do not match")
	}
	for i := range teams {
		user, err := db.GetUserProfileByDisplayName(players[i])
		if err != nil {
			return fmt.Errorf("user %s does not exist", players[i])
		}
		teams[i].TeamName = teamName
		teams[i].Player = user.Username
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	for _, t := range teams {
		if err := tx.Save(&t).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (db DB) FetchRecentMatchesByTeamId(teamId []string, lastMatchNumber int) ([]Match, error) {
	matches := []Match{}
	var err error
	limit := lastMatchNumber
	err = db.Table("matches").
		Order("match_date desc").
		Where(`team_one IN (?) OR team_two IN (?) AND "status" = ?`, teamId, teamId, "Finished").
		Limit(limit).
		Find(&matches).
		Error
	if err != nil {
		return nil, err
	}
	return matches, nil
} //SELECT columnA FROM tableA WHERE ;

func (db DB) FetchTeamInMatchesByTournamentId(tournamentId string) ([]Team, error) {
	teams := []Team{}
	err := db.Table("teams").
		Select("teams.team_id, teams.team_name").
		Group("team_id, team_name").
		Where("teams.tournament_id = ?", tournamentId).
		Find(&teams).
		Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (db DB) FetchIndividualInMatchesByTournamentId(tournamentId string) ([]Team, error) {
	teams := []Team{}
	err := db.Table("teams").
		Select("teams.team_id, teams.player").
		Group("team_id, player").
		Where("teams.tournament_id = ? AND team_name IS NULL", tournamentId).
		Find(&teams).
		Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (db DB) FetchTeamIdsByTeamName(teamName string) ([]string, error) {
	teams := []Team{}
	teamIds := []string{}
	err := db.Table("teams").
		Select("teams.team_id, teams.team_name").
		Group("team_id, team_name").
		Where("teams.team_name = ?", teamName).
		Find(&teams).
		Error
	if err != nil {
		return nil, err
	}
	for _, team := range teams {
		teamId := team.TeamID
		teamIds = append(teamIds, teamId)
	}
	return teamIds, nil
}

func (db DB) FetchTeamIdsByPlayerName(playerName string) ([]string, error) {
	players := []Team{}
	teamIds := []string{}
	err := db.Table("teams").
		Select("teams.team_id, teams.player").
		Group("team_id, player").
		Where("teams.player = ?", playerName).
		Find(&players).
		Error
	if err != nil {
		return nil, err
	}
	for _, player := range players {
		teamId := player.TeamID
		teamIds = append(teamIds, teamId)
	}
	return teamIds, nil
}

func (db DB) CountUnfinishedMatches(tournamentId string) (int, error) {
	var matches []Match
	statusFinished := "Finished"
	err := db.DB.Table("matches").Where(`"tournament_id" = ? AND "status" != ?`, tournamentId, statusFinished).
		Find(&matches).Error
	if err != nil {
		return -1, err
	}
	return len(matches), nil
}

func (db DB) GetMatchesByStatus(status string) ([]Match, error) {
	matches := make([]Match, 0)
	err := db.DB.Where("status = ?", status).Find(&matches).Error
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (db DB) PrepareManualMatch(match Match) error {
	if match.Status != "ManualWaiting" {
		return fmt.Errorf("Match not in ManualWaiting status, tournament: %s, round: %d, table: %d", match.TournamentID, match.Round, match.Table)
	}

	newMatch, err := db.ParsePlayer(match)
	if err != nil {
		return fmt.Errorf("error in Parse Player: %v, tournament: %s, round: %d, table: %d", err, match.TournamentID, match.Round, match.Table)
	}
	err = db.DB.Model(&newMatch).Updates(Match{TeamOne: newMatch.TeamOne, TeamTwo: newMatch.TeamTwo, Status: "ManualReady"}).Error
	if err != nil {
		return fmt.Errorf("error in Update match: %v, tournament: %s, round: %d, table: %d", err, match.TournamentID, match.Round, match.Table)
	}
	return nil
}

func (db DB) PrepareManualTournamentStart(tournamentId string) error {
	t, err := db.GetTournament(tournamentId)
	if err != nil {
		return fmt.Errorf("Error in GetTournament(%s): %v", tournamentId, err)
	}
	if t.Status != "REGISTRATION" {
		return fmt.Errorf("tournament %s not in REGISTRATION status", tournamentId)
	}
	matches, err := db.GetMatchesByTournament(tournamentId)
	if err != nil {
		return fmt.Errorf("Error in GetMatchesByTournament(%s): %v", tournamentId, err)
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	for _, match := range matches {
		if match.Round == 1 {
			if match.Status != "ManualRegistration" {
				tx.Rollback()
				return fmt.Errorf("Match not in ManualRegistration status, tournament: %s, round: %d, table: %d", match.TournamentID, match.Round, match.Table)
			}
			teamOnePlayers, err := db.GetPlayersByTeam(match.TeamOne)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("Error in GetPlayersByTeam - %s: %v", match.TeamOne, err)
			}
			teamTwoPlayers, err := db.GetPlayersByTeam(match.TeamTwo)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("Error in GetPlayersByTeam - %s: %v", match.TeamOne, err)
			}
			if len(teamOnePlayers) < 1 {
				match.Result = 2
				match.Status = "Finished"
				if err := tx.Save(&match).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else if len(teamTwoPlayers) < 1 {
				match.Result = 1
				match.Status = "Finished"
				if err := tx.Save(&match).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	t.Status = "STARTED"
	if err := tx.Save(&t).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db DB) AutoAdvance(matches []Match) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	for _, match := range matches {
		if match.Status == "ManualRegistration" || match.Status == "ManualReady" {
			teamOnePlayers, err := db.GetPlayersByTeam(match.TeamOne)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("Error in GetPlayersByTeam - %s: %v", match.TeamOne, err)
			}
			teamTwoPlayers, err := db.GetPlayersByTeam(match.TeamTwo)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("Error in GetPlayersByTeam - %s: %v", match.TeamOne, err)
			}
			if len(teamOnePlayers) < 1 {
				match.Result = 2
				match.Status = "Finished"
				if err := tx.Save(&match).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else if len(teamTwoPlayers) < 1 {
				match.Result = 1
				match.Status = "Finished"
				if err := tx.Save(&match).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	return tx.Commit().Error
}

func (db DB) PrepareMatch(match Match) error {
	if match.Status != "Waiting" {
		return fmt.Errorf("Match not in Waiting status, tournament: %s, round: %d, table: %d", match.TournamentID, match.Round, match.Table)
	}

	newMatch, err := db.ParsePlayer(match)
	if err != nil {
		return fmt.Errorf("error in Parse Player: %v, tournament: %s, round: %d, table: %d", err, match.TournamentID, match.Round, match.Table)
	}
	err = db.DB.Model(&newMatch).Updates(Match{TeamOne: newMatch.TeamOne, TeamTwo: newMatch.TeamTwo, Status: "Ready", MatchDate: time.Now().Add(6 * time.Minute)}).Error
	if err != nil {
		return fmt.Errorf("error in Update match: %v, tournament: %s, round: %d, table: %d", err, match.TournamentID, match.Round, match.Table)
	}
	return nil
}

func (db DB) PrepareRematch(match Match) error {
	if match.Status != "Break" {
		return fmt.Errorf("Match not in Break status, tournament: %s, round: %d, table: %d", match.TournamentID, match.Round, match.Table)
	}

	var err error
	switch match.Result {
	case -1:
		err = db.DB.Model(&match).Updates(Match{Status: "Finished"}).Error
	case 1:
		if match.TeamOneScore+1 > match.BestOfN/2 {
			err = db.DB.Model(&match).Updates(Match{Status: "Finished", TeamOneScore: match.TeamOneScore + 1}).Error
		} else {
			err = db.DB.Model(&match).Updates(Match{Status: "Ready", TeamOneScore: match.TeamOneScore + 1}).Error
		}
	case 2:
		if match.TeamTwoScore+1 > match.BestOfN/2 {
			err = db.DB.Model(&match).Updates(Match{Status: "Finished", TeamTwoScore: match.TeamTwoScore + 1}).Error
		} else {
			err = db.DB.Model(&match).Updates(Match{Status: "Ready", TeamTwoScore: match.TeamTwoScore + 1}).Error
		}
	default:
		err = fmt.Errorf("invalid match result")
	}
	if err != nil {
		return fmt.Errorf("error in Update match: %v, tournament: %s, round: %d, table: %d", err, match.TournamentID, match.Round, match.Table)
	}

	return nil
}

func (db DB) PrepareRoundOne(match Match) error {
	if match.Status != "Registration" {
		return fmt.Errorf("Match not in Registration status, tournament: %s, round: %d, table: %d", match.TournamentID, match.Round, match.Table)
	}
	err := db.DB.Model(&match).Updates(Match{Status: "Ready"}).Error
	if err != nil {
		return fmt.Errorf("error in Update match: %v, tournament: %s, round: %d, table: %d", err, match.TournamentID, match.Round, match.Table)
	}
	return nil
}

func (db DB) ParsePlayer(match Match) (Match, error) {
	newMatch := match
	var pOne, pTwo []string
	delimiter := regexp.MustCompile(`,|\(|\)`)
	pOne = delimiter.Split(match.TeamOne, 4)
	pTwo = delimiter.Split(match.TeamTwo, 4)

	switch strings.ToLower(pOne[0]) {
	case "winner":
		winner, err := db.Winner(match.TournamentID, pOne[1], pOne[2])
		if err != nil {
			return match, fmt.Errorf("error calling winner(): %v", err)
		}
		newMatch.TeamOne = winner
	case "loser":
		loser, err := db.Loser(match.TournamentID, pOne[1], pOne[2])
		if err != nil {
			return match, fmt.Errorf("error calling loser(): %v", err)
		}
		newMatch.TeamOne = loser
	default:
		// do nothing
	}

	switch strings.ToLower(pTwo[0]) {
	case "winner":
		winner, err := db.Winner(match.TournamentID, pTwo[1], pTwo[2])
		if err != nil {
			return match, fmt.Errorf("error calling winner(): %v", err)
		}
		newMatch.TeamTwo = winner
	case "loser":
		loser, err := db.Loser(match.TournamentID, pTwo[1], pTwo[2])
		if err != nil {
			return match, fmt.Errorf("error calling loser(): %v", err)
		}
		newMatch.TeamTwo = loser
	default:
		// do nothing
	}

	return newMatch, nil
}

func (db DB) DeleteMatch(tournamentId string, round, table int) error {
	match := Match{}
	err := db.DB.Where(`"tournament_id" = ? AND "round" = ? AND "table" = ?`, tournamentId, round, table).First(&match).Error
	if err != nil {
		return err
	}
	err = db.DB.Delete(&match).Error
	if err != nil {
		return err
	}
	return nil
}

func (db DB) Winner(tournamentId, roundString, tableString string) (string, error) {
	round, _ := strconv.Atoi(roundString)
	table, _ := strconv.Atoi(tableString)
	match, err := db.GetMatch(tournamentId, round, table)
	if err != nil {
		return "", fmt.Errorf("failed to call GetMatch(tournamentId: %s, round: %d, table: %d), err: %v", tournamentId, round, table, err)
	}
	if match.Status != "Finished" {
		return "", errors.New("match is not finished")
	}
	var winner string
	if match.Result == 1 {
		winner = match.TeamOne
	} else if match.Result == 2 {
		winner = match.TeamTwo
	} else {
		if match.TeamOne > match.TeamTwo {
			winner = match.TeamOne
		} else {
			winner = match.TeamTwo
		}
	}
	return winner, nil
}

func (db DB) Loser(tournamentId, roundString, tableString string) (string, error) {
	round, _ := strconv.Atoi(roundString)
	table, _ := strconv.Atoi(tableString)
	match, err := db.GetMatch(tournamentId, round, table)
	if err != nil {
		return "", err
	}
	if match.Status != "Finished" {
		return "", errors.New("match is not finished")
	}
	var loser string
	if match.Result == 1 {
		loser = match.TeamTwo
	} else if match.Result == 2 {
		loser = match.TeamOne
	} else {
		if match.TeamOne > match.TeamTwo {
			loser = match.TeamTwo
		} else {
			loser = match.TeamOne
		}
	}
	return loser, nil
}

type TeamOutput struct {
	TeamID   string     `json:"teamId" gorm:"type:uuid"`
	TeamName string     `json:"teamName"`
	Players  []UserInfo `json:"players"`
}

func (db DB) GetTeamUserInfo(teamId string) (*TeamOutput, error) {
	var rows *sql.Rows
	var err error
	players := []UserInfo{}
	team := TeamOutput{TeamID: teamId}

	rows, err = db.Table("teams").
		Select("teams.team_id,teams.team_name,users.sub,users.username,users.display_name, users.avatar_url").
		Joins("LEFT JOIN user_accounts AS users ON teams.player = users.username").
		Where("teams.team_id = ?", teamId).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var player UserInfo
		_ = db.ScanRows(rows, &player)
		if player.Username != "" {
			players = append(players, player)
		} else {
			_ = db.ScanRows(rows, &team)
		}
	}
	team.Players = players
	return &team, nil
}

func (db DB) FormatMatchesForOutput(matches []Match) ([]MatchOutput, error) {
	matchesOutput := make([]MatchOutput, len(matches))

	for i := range matches {
		matchesOutput[i].MatchID = matches[i].MatchID
		matchesOutput[i].TournamentID = matches[i].TournamentID
		matchesOutput[i].Round = matches[i].Round
		matchesOutput[i].Table = matches[i].Table
		matchesOutput[i].Status = matches[i].Status
		matchesOutput[i].Result = matches[i].Result
		matchesOutput[i].MatchDate = matches[i].MatchDate
		if _, err := uuid.FromString(matches[i].TeamOne); err == nil {
			teamOne, err := db.GetTeamUserInfo(matches[i].TeamOne)
			if err != nil {
				return nil, err
			}
			matchesOutput[i].TeamOne = *teamOne
		}
		if _, err := uuid.FromString(matches[i].TeamTwo); err == nil {
			teamTwo, err := db.GetTeamUserInfo(matches[i].TeamTwo)
			if err != nil {
				return nil, err
			}
			matchesOutput[i].TeamTwo = *teamTwo
		}
		// if matches[i].Result == 1 && matches[i].Status == "Finished" {
		// 	matchesOutput[i].Winner = matches[i].TeamOne
		// } else if matches[i].Result == 2 && matches[i].Status == "Finished" {
		// 	matchesOutput[i].Winner = matches[i].TeamTwo
		// } else if matches[i].Result == -1 && matches[i].Status == "Finished" {
		// 	if matches[i].TeamOne > matches[i].TeamTwo {
		// 		matchesOutput[i].Winner = matches[i].TeamOne
		// 	} else {
		// 		matchesOutput[i].Winner = matches[i].TeamTwo
		// 	}
		// } else {
		// 	// match is not finished, do nothing
		// }
		// if matches[i].Status == "Waiting" {
		// 	matchesOutput[i].TeamOne = teamTemp
		// 	matchesOutput[i].TeamOne.TeamName = matches[i].TeamOne
		// } else {
		// 	matchesOutput[i].TeamOne, err = tc.generateTeamOutput(matches[i].TeamOne)
		// 	if err != nil {
		// 		tc.log.Error("error getting tournament team information: ", err)
		// 		tc.DBErrorResponse(c, "error getting tournament team information")
		// 		return
		// 	}
		// }
		// if matches[i].Status == "Waiting" {
		// 	matchesOutput[i].TeamTwo = teamTemp
		// 	matchesOutput[i].TeamTwo.TeamName = matches[i].TeamTwo
		// } else {
		// 	matchesOutput[i].TeamTwo, err = tc.generateTeamOutput(matches[i].TeamTwo)
		// 	if err != nil {
		// 		tc.log.Error("error getting tournament team information: ", err)
		// 		tc.DBErrorResponse(c, "error getting tournament team information")
		// 		return
		// 	}
		// }
		matchesOutput[i].BestOfN = matches[i].BestOfN
		matchesOutput[i].TeamOneScore = matches[i].TeamOneScore
		matchesOutput[i].TeamTwoScore = matches[i].TeamTwoScore

		// if matches[i].Round == 1 {
		// 	if strings.ToUpper(tournament.GameSubtype) == "1V1" || tournament.GameSubtype == "SOLO" {
		// 		var p1, p2 PlayerOutput
		// 		for _, participant := range matchesOutput[i].TeamOne.Players {
		// 			p1.AvatarUrl = participant.AvatarUrl
		// 			p1.DisplayName = participant.DisplayName
		// 			participants = append(participants, p1)
		// 		}

		// 		for _, participant := range matchesOutput[i].TeamTwo.Players {
		// 			p2.AvatarUrl = participant.AvatarUrl
		// 			p2.DisplayName = participant.DisplayName
		// 			participants = append(participants, p2)
		// 		}
		// 	} else {
		// 		var p1, p2 PlayerOutput
		// 		p1.DisplayName = matchesOutput[i].TeamOne.TeamName
		// 		participants = append(participants, p1)

		// 		p2.DisplayName = matchesOutput[i].TeamTwo.TeamName
		// 		participants = append(participants, p2)
		// 	}
		// }
	}
	return matchesOutput, nil
}

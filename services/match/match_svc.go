/*

 */

package organizer

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/bitspawngg/bitspawn-api/enum"
	"github.com/bitspawngg/bitspawn-api/models"
)

type MatchService struct {
	log *logrus.Entry
	db  *models.DB
}

func NewMatchService(log *logrus.Logger, db *models.DB) *MatchService {
	return &MatchService{
		log: log.WithField("services", "Match"),
		db:  db,
	}
}

func GetMatchSchedule(teams []models.TeamDTO, format enum.TournamentFormat, participantsPlayEachOther int) ([]models.Match, error) {
	var matches []models.Match
	nTeams := len(teams)
	// Create matches in Round 1
	for table := 1; table <= nTeams/2; table++ {
		match := models.Match{
			Round:   1,
			Table:   table,
			Status:  "Registration",
			TeamOne: teams[2*table-2].TeamID,
			TeamTwo: teams[2*table-1].TeamID,
		}
		matches = append(matches, match)
	}
	switch format {
	case enum.SingleElimination:
		nTeams = nTeams / 2
		// Create matches for later rounds
		for round := 2; nTeams > 1; round++ {
			for table := 1; table <= nTeams/2; table++ {
				match := models.Match{
					Round:   round,
					Table:   table,
					TeamOne: "Winner(" + strconv.Itoa(round-1) + "," + strconv.Itoa(table*2-1) + ")",
					TeamTwo: "Winner(" + strconv.Itoa(round-1) + "," + strconv.Itoa(table*2) + ")",
					Status:  "Waiting",
				}
				matches = append(matches, match)
			}
			if nTeams == 2 {
				matchForThird := models.Match{
					Round:   round,
					Table:   2,
					TeamOne: "Loser(" + strconv.Itoa(round-1) + "," + "1)",
					TeamTwo: "Loser(" + strconv.Itoa(round-1) + "," + "2)",
					Status:  "Waiting",
				}
				matches = append(matches, matchForThird)
			}
			nTeams = nTeams / 2
		}
	case enum.Consolation:
		for round := 2; 1<<round <= nTeams; round++ {
			for table := 1; table <= nTeams/4; table++ {
				match := models.Match{
					Round:   round,
					Table:   table,
					TeamOne: "Winner(" + strconv.Itoa(round-1) + "," + strconv.Itoa(table*2-1) + ")",
					TeamTwo: "Winner(" + strconv.Itoa(round-1) + "," + strconv.Itoa(table*2) + ")",
					Status:  "Waiting",
				}
				matches = append(matches, match)
			}
			for table := nTeams/4 + 1; table <= nTeams/2; table++ {
				match := models.Match{
					Round:   round,
					Table:   table,
					TeamOne: "Loser(" + strconv.Itoa(round-1) + "," + strconv.Itoa((table-nTeams/4)*2-1) + ")",
					TeamTwo: "Loser(" + strconv.Itoa(round-1) + "," + strconv.Itoa((table-nTeams/4)*2) + ")",
					Status:  "Waiting",
				}
				matches = append(matches, match)
			}
		}
	case enum.DoubleElimination:
		if nTeams == 2 {
			// do nothing
		} else {
			for table := 1; table <= nTeams/4; table++ {
				match := models.Match{
					Round:   2,
					Table:   table,
					TeamOne: "Winner(1," + strconv.Itoa(table*2-1) + ")",
					TeamTwo: "Winner(1," + strconv.Itoa(table*2) + ")",
					Status:  "Waiting",
				}
				matches = append(matches, match)
			}
			for table := nTeams/4 + 1; table <= nTeams/2; table++ {
				match := models.Match{
					Round:   2,
					Table:   table,
					TeamOne: "Loser(1," + strconv.Itoa((table-nTeams/4)*2-1) + ")",
					TeamTwo: "Loser(1," + strconv.Itoa((table-nTeams/4)*2) + ")",
					Status:  "Waiting",
				}
				matches = append(matches, match)
			}
			matches = doubleElim(matches, nTeams, 3)
		}
	case enum.RoundRobin:
		table := 1
		for i := 1; i < participantsPlayEachOther; i++ {
			for _, teamA := range teams {
				for _, teamB := range teams {
					if teamA.TeamID == teamB.TeamID {
						continue
					}
					match := models.Match{
						Round:   i,
						Table:   table,
						TeamOne: teamA.TeamID,
						TeamTwo: teamB.TeamID,
						Status:  "Registration",
					}
					matches = append(matches, match)
					table++
				}
			}
		}
	default:
		return nil, fmt.Errorf("Unsupported tournament format [%s]", format)
	}

	return matches, nil
}

func doubleElim(matches []models.Match, nTeams, round int) []models.Match {
	if nTeams <= 4 {
		match := models.Match{
			Round:   round,
			Table:   1,
			TeamOne: "Loser(" + strconv.Itoa(round-1) + ",1)",
			TeamTwo: "Winner(" + strconv.Itoa(round-1) + ",2)",
			Status:  "Waiting",
		}
		matchFinal := models.Match{
			Round:   round + 1,
			Table:   1,
			TeamOne: "Winner(" + strconv.Itoa(round-1) + ",1)",
			TeamTwo: "Winner(" + strconv.Itoa(round) + ",1)",
			Status:  "Waiting",
		}
		matches = append(matches, match)
		matches = append(matches, matchFinal)
		return matches
	}
	for table := 1; table <= nTeams/4; table++ {
		match := models.Match{
			Round:   round,
			Table:   table,
			TeamOne: "Loser(" + strconv.Itoa(round-1) + "," + strconv.Itoa(table) + ")",
			TeamTwo: "Winner(" + strconv.Itoa(round-1) + "," + strconv.Itoa(table+nTeams/4) + ")",
			Status:  "Waiting",
		}
		matches = append(matches, match)
	}
	for table := 1; table <= nTeams/8; table++ {
		match := models.Match{
			Round:   round + 1,
			Table:   table,
			TeamOne: "Winner(" + strconv.Itoa(round-1) + "," + strconv.Itoa(2*table-1) + ")",
			TeamTwo: "Winner(" + strconv.Itoa(round-1) + "," + strconv.Itoa(2*table) + ")",
			Status:  "Waiting",
		}
		matches = append(matches, match)
	}
	for table := nTeams/8 + 1; table <= nTeams/4; table++ {
		match := models.Match{
			Round:   round + 1,
			Table:   table,
			TeamOne: "Winner(" + strconv.Itoa(round) + "," + strconv.Itoa((table-nTeams/8)*2-1) + ")",
			TeamTwo: "Winner(" + strconv.Itoa(round) + "," + strconv.Itoa((table-nTeams/8)*2) + ")",
			Status:  "Waiting",
		}
		matches = append(matches, match)
	}
	return doubleElim(matches, nTeams/2, round+2)
}

func (ms *MatchService) CreateMatchSchedule(tournamentId string, isManual bool) error {
	existingMatches, err := ms.db.GetMatchesByTournament(tournamentId)
	if err != nil {
		return fmt.Errorf("Error in GetMatchesByTournament %s: %v", tournamentId, err)
	}
	if len(existingMatches) > 0 {
		return errors.New("match schedule already exists")
	}

	tournament, err := ms.db.GetTournament(tournamentId)
	if err != nil {
		return fmt.Errorf("Error in GetTournament %s: %v", tournamentId, err)
	}

	participants, err := ms.db.GetParticipants(tournamentId)
	if err != nil {
		return fmt.Errorf("Error in GetParticipants of %s: %v", tournamentId, err)
	}
	if int64(len(participants)) != tournament.ParticipantCount {
		return errors.New("Number of play records does not equal participant count")
	}

	teamSize, err := ms.getTeamSize(tournament.GameType, tournament.GameSubtype)
	if err != nil {
		return err
	}

	if !tournament.TournamentFormat.IsValid() {
		return fmt.Errorf("Unsupported tournament format [%s]", tournament.TournamentFormat.ToString())
	}

	teams, err := GetTeams(participants, teamSize)
	if err != nil {
		return fmt.Errorf("Error in getTeams: %v", err)
	}
	// it could be used for ROUND-ROBIN format to define the number of rounds
	matches, err := GetMatchSchedule(teams, tournament.TournamentFormat, tournament.ParticipantsPlayEachOther)
	if err != nil {
		return fmt.Errorf("Error in getMatchSchedule: %v", err)
	}

	numberOfRounds := 1
	for i, match := range matches {
		match.TournamentID = tournamentId
		match.MatchID = uuid.NewV4().String()
		match.BotID = uuid.NewV4().String()
		if match.Round == 1 {
			match.MatchDate = tournament.TournamentDate
		}
		match.BestOfN = 1
		if isManual {
			match.Status = "Manual" + match.Status
		}
		matches[i] = match
		if match.Round > numberOfRounds {
			numberOfRounds = match.Round
		}
	}
	roundsFormat := models.BestOfNFormat{}
	_ = json.Unmarshal(tournament.RoundsFormat, &roundsFormat)
	matches = updateBestOfN(matches, numberOfRounds, roundsFormat)

	teamRecords := []models.Team{}
	for _, team := range teams {
		for _, member := range team.Members {
			teamRecord := models.Team{
				TournamentID: tournamentId,
				TeamID:       team.TeamID,
				TeamName:     team.Name,
				Player:       member,
			}
			teamRecords = append(teamRecords, teamRecord)
		}
	}
	err = ms.db.CreateTeams(teamRecords)
	if err != nil {
		return fmt.Errorf("CreateTeams failed: %v", err)
	}

	err = ms.db.CreateMatches(matches)
	if err != nil {
		return fmt.Errorf("CreateMatches failed: %v", err)
	}

	ms.log.Info("Successfully Created Match Schedule for tournament: ", tournamentId)
	return nil
}

func (ms *MatchService) CreateManualMatchSchedule(tournamentId string) error {
	existingMatches, err := ms.db.GetMatchesByTournament(tournamentId)
	if err != nil {
		return fmt.Errorf("Error in GetMatchesByTournament %s: %v", tournamentId, err)
	}
	if len(existingMatches) > 0 {
		return errors.New("match schedule already exists")
	}

	tourney, err := ms.db.GetTournament(tournamentId)
	if err != nil {
		return fmt.Errorf("Error in GetTournament %s: %v", tournamentId, err)
	}

	if !tourney.TournamentFormat.IsValid() {
		return fmt.Errorf("Unsupported tournament format [%s]", tourney.TournamentFormat.ToString())
	}

	numberOfTeams := nextPow2(tourney.NumberOfTeams)
	teams := make([]models.TeamDTO, numberOfTeams)
	for i := range teams {
		teams[i].TeamID = uuid.NewV4().String()
		teams[i].Name = fmt.Sprintf("Team%d", i+1)
	}

	matches, err := GetMatchSchedule(teams, tourney.TournamentFormat, tourney.ParticipantsPlayEachOther)
	if err != nil {
		return fmt.Errorf("Error in getMatchSchedule: %v", err)
	}

	numberOfRounds := 1

	for i, match := range matches {
		matches[i].TournamentID = tournamentId
		matches[i].BotID = uuid.NewV4().String()
		matches[i].MatchID = uuid.NewV4().String()
		matches[i].MatchDate = tourney.TournamentDate.Add(time.Hour * time.Duration(match.Round-1))
		matches[i].BestOfN = 1
		matches[i].Status = "Manual" + match.Status
		if match.Round > numberOfRounds {
			numberOfRounds = match.Round
		}
	}

	teamRecords := []models.Team{}
	for _, team := range teams {
		teamRecord := models.Team{
			TournamentID: tournamentId,
			TeamID:       team.TeamID,
			TeamName:     team.Name,
			Player:       "",
		}
		teamRecords = append(teamRecords, teamRecord)
	}
	err = ms.db.CreateTeams(teamRecords)
	if err != nil {
		return fmt.Errorf("CreateTeams failed: %v", err)
	}

	err = ms.db.CreateMatches(matches)
	if err != nil {
		return fmt.Errorf("CreateMatches failed: %v", err)
	}

	ms.log.Info("Successfully Created Match Schedule for tournament: ", tournamentId)
	return nil
}

func (ms *MatchService) StartTournament() error {
	tournaments, err := ms.db.GetTournamentsToStart()
	if err != nil {
		if err.Error() == "record not found" {
			ms.log.Debug("No tournaments Ready to start")
			return nil
		} else {
			ms.log.Error("Error in GetTournamentToStart: ", err)
			return fmt.Errorf("Error in GetTournamentToStart: %v", err)
		}
	}
	for _, tid := range tournaments {
		ms.log.Info("Starting next tournament: ", tid)
		err = ms.CreateMatchSchedule(tid, false)
		if err != nil {
			ms.log.Error("Error in CreateMatchSchedule for Tournament ", "tid", ": ", err)
			continue
		}
		err = ms.db.StartTournament(tid)
		if err != nil {
			ms.log.Error("Error in StartTournament ", "tid", ": ", err)
			continue
		}
	}
	return nil
}

func (ms *MatchService) CompleteTournament(t *models.TournamentResponse) error {
	tournamentId := t.TournamentID
	if t.TournamentFormat != enum.RoundRobin {
		unfinishedMatchCount, err := ms.db.CountUnfinishedMatches(tournamentId)
		if err != nil {
			ms.log.Errorf("Error in CountUnfinishedMatches in %s: %v", tournamentId, err)
			return fmt.Errorf("Error in CountUnfinishedMatches in %s: %v", tournamentId, err)
		}
		if unfinishedMatchCount != 0 {
			return fmt.Errorf("There are still unfinished matches in tournament %s", tournamentId)
		}
	}

	winnerTeams, err := ms.getWinners(tournamentId)
	if err != nil {
		ms.log.Error("Error in getWinners for Tournament ", tournamentId, ": ", err)
		return fmt.Errorf("Error in getWinners for Tournament %s: %v", tournamentId, err)
	}
	if len(winnerTeams) < len(t.PrizeAllocation) {
		ms.log.Error("Winners less than prize allocation for Tournament ", tournamentId)
		return fmt.Errorf("Winners less than prize allocation for Tournament %s", tournamentId)
	}
	for i := 0; i < len(t.PrizeAllocation); i++ {
		playersOnWinnerTeam, err := ms.db.GetPlayersByTeam(winnerTeams[i])
		if err != nil {
			return fmt.Errorf("Error in GetPlayersByTeam(%s): %v", winnerTeams[i], err)
		}
		for _, winner := range playersOnWinnerTeam {
			prizeWon := float64(t.TotalPrizePool*t.PrizeAllocation[i]*(100-t.OrganizerPercentage-t.FeePercentage)) /
				float64(len(playersOnWinnerTeam)*10000)
			prizeWonInt := int(prizeWon)
			err = ms.db.ReportWinner(tournamentId, winner, i+1, prizeWonInt)
			if err != nil {
				ms.log.Errorf("Error in ReportWinners for Tournament %s: %v", tournamentId, err)
				return fmt.Errorf("Error in ReportWinners for Tournament %s: %v", tournamentId, err)
			}
		}
	}
	err = ms.db.UpdateTournamentStatus(tournamentId, "PAYOUT")
	if err != nil {
		ms.log.Errorf("Error updating tournament %s to Payout status: %v", tournamentId, err)
		return fmt.Errorf("Error updating tournament %s to Payout status: %v", tournamentId, err)
	}
	return nil
}

func (ms *MatchService) PrepareRoundOne() error {
	matchesInRegistration, err := ms.db.GetMatchesByStatus("Registration")
	if err != nil {
		ms.log.Error("Error in GetMatchesByStatus(Registration): ", err)
		return fmt.Errorf("Error in GetMatchesByStatus(Registration): %v", err)
	}
	for _, match := range matchesInRegistration {
		if match.MatchDate.Before(time.Now().Add(10 * time.Minute)) {
			err = ms.db.PrepareRoundOne(match)
			if err != nil {
				ms.log.Debug("Unable to Prepare Round One in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
				continue
			}
			ms.log.Info("Successfully prepared Round One in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
		}
	}
	return nil
}

func (ms *MatchService) PrepareManualMatches() error {
	matchesInManualWaiting, err := ms.db.GetMatchesByStatus("ManualWaiting")
	if err != nil {
		ms.log.Error("Error in GetMatchesByStatus(ManualWaiting): ", err)
		return fmt.Errorf("Error in GetMatchesByStatus(ManualWaiting): %v", err)
	}
	for _, match := range matchesInManualWaiting {
		err = ms.db.PrepareManualMatch(match)
		if err != nil {
			ms.log.Debug("Unable to prepare match in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
			continue
		}
		ms.log.Info("Successfully prepared match in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
	}
	return nil
}

func (ms *MatchService) PrepareManualMatchesByTournament(tournamentId string) error {
	allMatchesInTournament, err := ms.db.GetMatchesByTournament(tournamentId)
	if err != nil {
		ms.log.Error("Error in GetMatchesByTournament: ", err)
		return fmt.Errorf("Error in GetMatchesByTournament: %v", err)
	}
	err = ms.db.AutoAdvance(allMatchesInTournament)
	if err != nil {
		ms.log.Errorf("Error in AutoAdvance teams in tournament %s: %v", tournamentId, err)
		return fmt.Errorf("Error in AutoAdvance teams in tournament %s: %v", tournamentId, err)
	}
	for _, match := range allMatchesInTournament {
		if match.Status == "ManualWaiting" {
			err = ms.db.PrepareManualMatch(match)
			if err != nil {
				ms.log.Errorf("Unable to prepare match in tournament: %s, round: %d, table: %d", match.TournamentID, match.Round, match.Table)
				continue
			}
		}
	}
	return nil
}

func (ms *MatchService) StartManualTournament(tournamentId string) error {
	err := ms.db.PrepareManualTournamentStart(tournamentId)
	if err != nil {
		ms.log.Debug("error in PrepareManualTournamentStart - ", tournamentId, ": ", err)
		return fmt.Errorf("error in PrepareManualTournamentStart - %s: %v", tournamentId, err)
	}
	return nil
}

func (ms *MatchService) PrepareMatches() error {
	matchesInWaiting, err := ms.db.GetMatchesByStatus("Waiting")
	if err != nil {
		ms.log.Error("Error in GetMatchesByStatus(Waiting): ", err)
		return fmt.Errorf("Error in GetMatchesByStatus(Waiting): %v", err)
	}
	for _, match := range matchesInWaiting {
		err = ms.db.PrepareMatch(match)
		if err != nil {
			ms.log.Debug("Unable to prepare match in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
			continue
		}
		ms.log.Info("Successfully prepared match in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
	}
	return nil
}

func (ms *MatchService) PrepareRematches() error {
	matchesInBreak, err := ms.db.GetMatchesByStatus("Break")
	if err != nil {
		ms.log.Error("Error in GetMatchesByStatus(Break): ", err)
		return fmt.Errorf("Error in GetMatchesByStatus(Break): %v", err)
	}
	for _, match := range matchesInBreak {
		err = ms.db.PrepareRematch(match)
		if err != nil {
			ms.log.Debug("Unable to prepare rematch in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
			continue
		}
		ms.log.Info("Successfully prepared rematch in tournament: ", match.TournamentID, ", round: ", match.Round, ", table: ", match.Table)
	}
	return nil
}

func (ms *MatchService) UpdateRoundBestOfN(tournamentId string, round int, bestOfN int) ([]models.Match, error) {
	matches, err := ms.db.GetMatchesByTournament(tournamentId)
	if err != nil {
		return nil, err
	}
	for i := range matches {
		if matches[i].Round == round {
			matches[i].BestOfN = bestOfN
		}
	}
	err = ms.db.UpdateRoundBestOfN(tournamentId, round, bestOfN)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (ms *MatchService) UpdateRoundTime(tournamentId string, round int, matchDate time.Time) ([]models.Match, error) {
	matches, err := ms.db.GetMatchesByTournament(tournamentId)
	if err != nil {
		return nil, err
	}
	for i := range matches {
		if matches[i].Round == round {
			matches[i].MatchDate = matchDate
		}
	}
	err = ms.db.UpdateRoundMatchDate(tournamentId, round, matchDate)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func GetScoreboard(matches []models.Match) map[string]int {
	scoreboard := make(map[string]int)
	for _, match := range matches {
		if match.Status == "Finished" {
			if match.Result == 1 {
				scoreboard[match.TeamOne] = scoreboard[match.TeamOne] + 3
				scoreboard[match.TeamTwo] = scoreboard[match.TeamTwo] + 0
			} else if match.Result == 2 {
				scoreboard[match.TeamTwo] = scoreboard[match.TeamTwo] + 3
				scoreboard[match.TeamOne] = scoreboard[match.TeamOne] + 0
			} else {
				// no winner, match ended in a tie
				scoreboard[match.TeamOne] = scoreboard[match.TeamOne] + 1
				scoreboard[match.TeamTwo] = scoreboard[match.TeamTwo] + 1
			}
		}
	}
	return scoreboard
}

func GetTeams(playRecords []models.PlayRecord, teamSize int) ([]models.TeamDTO, error) {
	existingTeams := []models.TeamDTO{}
	if len(playRecords)%teamSize != 0 {
		return nil, fmt.Errorf("%d participants cannot be evenly broken into teams of %d", len(playRecords), teamSize)
	}
	nTeams := len(playRecords) / teamSize
	if nTeams < 2 {
		return nil, errors.New("not enough players to start with")
	}
	if nTeams&(nTeams-1) != 0 {
		return nil, errors.New("number of teams not a power of 2")
	}
	clubBucket := make(map[string][]int)
	clanSize := make(map[string]int)
	clanBucket := make(map[string][]int)
	noClanSize := 0
	noClanBucket := []int{}
	for i, playRecord := range playRecords {
		club := playRecord.Club
		if club != "" {
			clubBucket[club] = append(clubBucket[club], i)
		} else {
			clan := playRecord.Clan
			if clan == "" {
				noClanSize += 1
				noClanBucket = append(noClanBucket, i)
			} else {
				clanSize[clan] += 1
				clanBucket[clan] = append(clanBucket[clan], i)
			}
		}
	}
	for clubName, clubMembers := range clubBucket {
		members := []string{}
		for _, index := range clubMembers {
			members = append(members, playRecords[index].UserId)
		}
		existingTeams = append(existingTeams, newTeam(clubName, members, 0))
	}
	for len(clanSize) > 0 {
		largestClan := sortKeyByValue(clanSize)[0]
		if clanSize[largestClan] < 1 {
			break
		}
		// form teams
		if clanSize[largestClan] >= teamSize {
			var x []int
			x, clanBucket[largestClan] = clanBucket[largestClan][0:teamSize], clanBucket[largestClan][teamSize:]
			// pop n elements: x, a = a[0:n], a[n:]
			clanSize[largestClan] -= teamSize
			members := []string{}
			for _, index := range x {
				members = append(members, playRecords[index].UserId)
			}
			existingTeams = append(existingTeams, newTeam(largestClan, members, 0))
		} else {
			x := clanBucket[largestClan]
			clanBucket[largestClan] = nil
			clanSize[largestClan] = 0
			members := []string{}
			for _, index := range x {
				members = append(members, playRecords[index].UserId)
			}
			teamFilled := false
			sort.Slice(existingTeams, func(i, j int) bool { return existingTeams[i].Slot <= existingTeams[j].Slot })
			for j, t := range existingTeams {
				if t.Slot >= len(x) {
					existingTeams[j] = fillTeam(t, members)
					teamFilled = true
					break
				}
			}
			if !teamFilled {
				if len(existingTeams) < nTeams {
					existingTeams = append(existingTeams, newTeam(largestClan, members, teamSize-len(x)))
				} else {
					// remove a member from the largest clan
					clanBucket[largestClan] = x[1:]
					clanSize[largestClan] = len(x) - 1
					noClanBucket = append(noClanBucket, x[0])
					noClanSize += 1
				}
			}
		}
	}
	for noClanSize > 0 {
		var index int
		index, noClanBucket = noClanBucket[0], noClanBucket[1:]
		noClanSize -= 1
		members := []string{playRecords[index].UserId}
		teamFilled := false
		sort.Slice(existingTeams, func(i, j int) bool { return existingTeams[i].Slot <= existingTeams[j].Slot })
		for j, t := range existingTeams {
			if t.Slot > 0 {
				existingTeams[j] = fillTeam(t, members)
				teamFilled = true
				break
			}
		}
		if !teamFilled {
			if len(existingTeams) < nTeams {
				existingTeams = append(existingTeams, newTeam(petname.Name(), members, teamSize-1))
			}
		}
	}
	return existingTeams, nil
}

func fillTeam(unfinishedTeam models.TeamDTO, members []string) models.TeamDTO {
	unfinishedTeam.Members = append(unfinishedTeam.Members, members...)
	unfinishedTeam.Slot -= len(members)
	return unfinishedTeam
}

func newTeam(teamName string, members []string, remainingSlot int) models.TeamDTO {
	return models.TeamDTO{
		TeamID:  uuid.NewV4().String(),
		Name:    teamName,
		Members: members,
		Slot:    remainingSlot,
	}
}

func (ms *MatchService) getTeamSize(gameType, gameSubtype string) (int, error) {
	gameSubType, err := ms.db.GetSpecificGameSubType(gameType, gameSubtype)
	if err != nil || len(gameSubType) == 0 {
		return 0, fmt.Errorf("Unsupported game subtype [%s] in [%s]", gameSubtype, gameType)
	}
	return gameSubType[0].TeamSize, nil
}

func (ms *MatchService) getWinners(tournamentId string) ([]string, error) {
	tournament, err := ms.db.GetTournament(tournamentId)
	if err != nil {
		return nil, fmt.Errorf("Error in GetTournament(%s): %v", tournamentId, err)
	}
	matches, err := ms.db.GetMatchesByTournament(tournamentId)
	if err != nil {
		return nil, fmt.Errorf("Error in GetMatchesByTournament(%s): %v", tournamentId, err)
	}
	format := tournament.TournamentFormat
	if format == enum.DoubleElimination {
		winners := []string{}
		for i := len(matches) - 1; i >= 0; i-- {
			match := matches[i]
			winnerTeam, err := ms.db.Winner(tournamentId, strconv.Itoa(match.Round), strconv.Itoa(match.Table))
			if err != nil {
				return nil, fmt.Errorf("Error in Winner(%s,%d,%d): %v", tournamentId, match.Round, match.Table, err)
			}
			loserTeam, err := ms.db.Loser(tournamentId, strconv.Itoa(match.Round), strconv.Itoa(match.Table))
			if err != nil {
				return nil, fmt.Errorf("Error in Loser(%s,%d,%d): %v", tournamentId, match.Round, match.Table, err)
			}
			winners = appendIfMissing(winners, winnerTeam)
			winners = appendIfMissing(winners, loserTeam)
		}
		return winners, nil
	} else if format == enum.SingleElimination || format == enum.Consolation {
		scoreMap := make(map[string]int)
		for _, match := range matches {
			winnerTeam, err := ms.db.Winner(tournamentId, strconv.Itoa(match.Round), strconv.Itoa(match.Table))
			if err != nil {
				return nil, fmt.Errorf("Error in Winner(%s,%d,%d): %v", tournamentId, match.Round, match.Table, err)
			}
			loserTeam, err := ms.db.Loser(tournamentId, strconv.Itoa(match.Round), strconv.Itoa(match.Table))
			if err != nil {
				return nil, fmt.Errorf("Error in Loser(%s,%d,%d): %v", tournamentId, match.Round, match.Table, err)
			}
			scoreMap[winnerTeam] = scoreMap[winnerTeam]*2 + 1
			scoreMap[loserTeam] = scoreMap[loserTeam] * 2
		}
		return sortKeyByValue(scoreMap), nil
	} else if format == enum.RoundRobin {
		scoreMap := GetScoreboard(matches)
		return sortKeyByValue(scoreMap), nil
	}
	return nil, fmt.Errorf("Unsupported tournament format [%s]", format)
}

func sortKeyByValue(m map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	output := []string{}
	for _, kv := range ss {
		output = append(output, kv.Key)
	}
	return output
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func updateBestOfN(matches []models.Match, numberOfRounds int, roundsFormat models.BestOfNFormat) []models.Match {
	for i, match := range matches {
		matches[i].BestOfN = bestOfN_to_int(roundsFormat.AllRounds)
		if match.Round == numberOfRounds {
			if roundsFormat.Finals != "" {
				matches[i].BestOfN = bestOfN_to_int(roundsFormat.Finals)
			}
		}
		if numberOfRounds > 1 {
			if match.Round == numberOfRounds-1 {
				if roundsFormat.SemiFinals != "" {
					matches[i].BestOfN = bestOfN_to_int(roundsFormat.SemiFinals)
				}
			}
		}
	}
	return matches
}

func bestOfN_to_int(s string) int {
	if s == "BO5" {
		return 5
	} else if s == "BO3" {
		return 3
	} else if s == "BO1" {
		return 1
	} else {
		return 1
	}
}

func nextPow2(n int) int {
	// Assumes a non-negative integer
	if n <= 1 {
		return n
	}
	if n&(n-1) == 0 {
		return n
	}

	count := 0
	for n != 0 {
		n >>= 1
		count += 1
	}

	return 1 << count
}

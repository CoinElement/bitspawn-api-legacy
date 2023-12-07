package models

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupByTeamNameEmpty(t *testing.T) {
	db := NewDB("", "", logrus.New())
	leaders := []ChallengeLeader{
		{AvatarUrl: "", DisplayName: "p1", Score: 10},
		{AvatarUrl: "", DisplayName: "p2", Score: 11},
		{AvatarUrl: "", DisplayName: "p3", Score: 20},
	}
	teamScores := db.GroupByTeam(leaders)
	assert.Equal(t, 0, len(teamScores))

}

func TestGroupByTeam(t *testing.T) {
	db := NewDB("", "", logrus.New())
	leaders := []ChallengeLeader{
		{AvatarUrl: "", DisplayName: "p1", Score: 10, TeamName: "team1"},
		{AvatarUrl: "", DisplayName: "p2", Score: 11, TeamName: "team1"},
		{AvatarUrl: "", DisplayName: "p3", Score: 20, TeamName: "team2"},
		{AvatarUrl: "", DisplayName: "p4", Score: 93, TeamName: "team1"},
		{AvatarUrl: "", DisplayName: "p5", Score: 143, TeamName: "team3"},
		{AvatarUrl: "", DisplayName: "p6", Score: 1, TeamName: "team1"},
		{AvatarUrl: "", DisplayName: "p7", Score: 109, TeamName: "team2"},
		{AvatarUrl: "", DisplayName: "p8", Score: 11, TeamName: "team3"},
	}
	teamScores := db.GroupByTeam(leaders)
	assert.Equal(t, 3, len(teamScores))
	assert.Equal(t, "team3", teamScores[0].TeamName)
	assert.Equal(t, 2, len(teamScores[0].Players))
	assert.Equal(t, float64(154), teamScores[0].TeamScore)

	assert.Equal(t, "team2", teamScores[1].TeamName)
	assert.Equal(t, 2, len(teamScores[1].Players))
	assert.Equal(t, float64(129), teamScores[1].TeamScore)

	assert.Equal(t, "team1", teamScores[2].TeamName)
	assert.Equal(t, 4, len(teamScores[2].Players))
	assert.Equal(t, float64(115), teamScores[2].TeamScore)
}

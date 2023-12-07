package controllers

import (
	"errors"
	"github.com/bitspawngg/bitspawn-api/enum"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		description string
		payload     FormTournamentUpdate
		expectedErr error
	}{
		{"TournamentRule missed", FormTournamentUpdate{
			TournamentID:          "club",
			TournamentDescription: toString("desc"),
			TournamentRule:        toString("rule"),
			CriticalRules:         toString("critical"),
			TournamentName:        "name",
			GameType:              "otl",
			GameSubtype:           "sub",
			Consoles:              []enum.Console{enum.PC},
			CutoffDate:            time.Now(),
			EntryFee:              toInt64(13),
		}, errors.New("mandatory field OrganizerPercentage not provided")},
		{"TournamentRule missed", FormTournamentUpdate{
			TournamentID:          "club",
			TournamentDescription: toString("desc"),
			TournamentRule:        toString("rule"),
			CriticalRules:         toString("critical"),
			OrganizerPercentage:   toInt64(10),
			TournamentName:        "name",
			GameType:              "otl",
			GameSubtype:           "sub",
			Consoles:              []enum.Console{enum.PC},
			CutoffDate:            time.Now(),
			EntryFee:              toInt64(13),
		}, errors.New("mandatory field MinPrizePool not provided")},
		{"TournamentRule missed", FormTournamentUpdate{
			TournamentID:          "club",
			TournamentDescription: toString("desc"),
			TournamentRule:        toString("rule"),
			CriticalRules:         toString("critical"),
			OrganizerPercentage:   toInt64(10),
			TournamentName:        "name",
			MinPrizePool:          toInt64(100),
			FeeType:               "fee",
			GameType:              "otl",
			GameSubtype:           "sub",
			CutoffDate:            time.Now(),
			EntryFee:              toInt64(13),
			TournamentDate:        time.Now(),
		}, errors.New("mandatory field InviteOnly not provided")},
		{"TournamentRule missed", FormTournamentUpdate{
			TournamentID:          "club",
			TournamentDescription: toString("desc"),
			TournamentRule:        toString("rule"),
			CriticalRules:         toString("critical"),
			OrganizerPercentage:   toInt64(10),
			TournamentName:        "name",
			MinPrizePool:          toInt64(100),
			FeeType:               "fee",
			GameType:              "otl",
			GameSubtype:           "sub",
			Consoles:              []enum.Console{enum.PC},
			CutoffDate:            time.Now(),
			EntryFee:              toInt64(13),
			InviteOnly:            toBool(true),
			TournamentDate:        time.Now(),
		}, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, tc.payload.Validate("draft"), tc.description)
		})
	}

}

func toString(s string) *string {
	return &s
}

func toInt64(i int64) *int64 {
	return &i
}

func toBool(b bool) *bool {
	return &b
}

package enum

type TournamentFormat string

const (
	SingleElimination TournamentFormat = "SINGLE_ELIM"
	DoubleElimination TournamentFormat = "DOUBLE_ELIM"
	RoundRobin        TournamentFormat = "ROUND_ROBIN"
	Consolation       TournamentFormat = "CONSOLATION"
)

func (tf TournamentFormat) IsValid() bool {
	if tf != SingleElimination && tf != RoundRobin {
		return false
	}
	return true
}

func (tf TournamentFormat) ToString() string {
	return string(tf)
}

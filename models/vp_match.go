package models

import "time"

type VpMatch struct {
	Id             int
	TeamAID        string
	TeamBID        string
	TeamA          string
	TeamB          string
	LogoA          string
	LogoB          string
	TeamAShort     string
	TeamBShort     string
	Tournament     string
	TournamentLogo string
	Game           string
	BestOf         string
	// sub match specific
	MatchID        string
	URL            string
	Time           *time.Time
	MatchName      string
	MatchType      []string
	ModeName       string
	ModeDesc       string
	HandicapAmount string
	HandicapTeam   string
	RatioA         float64
	RatioB         float64
	Winner         string
	Status         string
	ScoreA         float64
	ScoreB         float64
	Note           string
	SeriesID       string
}

package models

import (
	_ "github.com/lib/pq"

	"poker/gamelogic"
)

type PageData struct {
	SiteRoot    string
	Session     *Session
	UserPage    *UserPage
	GameListing *GameListing
	Lobby       *Lobby
	Leaderboard *Leaderboard
}

type UserPage struct {
	MatchesSession bool
	Username       string
	Name           string
	Email          string
	Description    string
	PictureSlug    string
}

type Lobby struct {
	Empty bool
	Games []*GameListing
}

type GameListing struct {
	Name        string
	Slug        string
	Status      string
	Ante        int
	MinBet      int
	MaxBet      int
	PlayerCount int
	Game        *gamelogic.Game
}

type LeaderboardEntry struct {
	Username string
	Cash     int64
}

type Leaderboard struct {
	Empty   bool
	Entries []*LeaderboardEntry
}

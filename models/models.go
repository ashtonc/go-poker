package models

import (
	"database/sql"
	"html/template"

	_ "github.com/lib/pq"

	"poker/gamelogic"
)

/*
 *    Session management
 */

// Env serves as an environment that contains "global" variables. See
// http://www.alexedwards.net/blog/organising-database-access for the idea
type Env struct {
	Database  *sql.DB
	Port      string
	Templates map[string]*template.Template
	SiteRoot  string
	Games     map[string]*gamelogic.Game
	// authentication middleware
	// logger middleware
}

type Session struct {
	LoggedIn        bool
	Username        string
	Name            string
	PageHome        bool
	PageLogin       bool
	PageRegister    bool
	PageUser        bool
	PageGame        bool
	PageLeaderboard bool
}

type UserPage struct {
	MatchesSession bool
	Username       string
	Name           string
	Email          string
	PictureUrl     string
}

/*
 *    Template models
 */

type PageData struct {
	SiteRoot    string
	Session     Session
	UserPage    UserPage
	Game        Game
	Lobby       Lobby
	Leaderboard Leaderboard
	Games       map[string]Game
}

type Lobby struct {
	Empty bool
	Games []LobbyListing
}

type LobbyListing struct {
	Name    string
	Stakes  GameStakes
	Players int
	Private bool
}

type LeaderboardEntry struct {
	Username string
	Cash     int64
}

type Leaderboard struct {
	Empty   bool
	Entries []LeaderboardEntry
}

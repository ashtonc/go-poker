package models

import (
	"database/sql"
	"html/template"

	"github.com/gorilla/websocket"
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
	Games     map[string]*GameListing
	Upgrader  *websocket.Upgrader
	// authentication middleware
	//env.Games["slug"].Game gete you a game 
}

type UserAccount struct {
	Username string
	Name     string
	Email    string
	Password string
}

/*
 *    Template models
 */

type PageData struct {
	SiteRoot    string
	Session     *Session
	UserPage    *UserPage
	GameListing *GameListing
	Lobby       *Lobby
	Leaderboard *Leaderboard
}

type Session struct {
	LoggedIn        bool
	Username        string
	Name            string
	PageHome        bool
	PageUser        bool
	PageGame        bool
	PageLeaderboard bool
	PageLogin       bool
	PageRegister    bool
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

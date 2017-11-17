package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Env will serve as an environment that contains "global" variables. See
// http://www.alexedwards.net/blog/organising-database-access for more details.
type Env struct {
	Database *sql.DB
	// template cache
	// logger middleware
	// authentication middleware
}

type UserPage struct {
	Username string
	Name string
	Email string
	PictureUrl string
}

type Session struct {
	LoggedIn bool
	Username string
	Name string
	PageHome bool
	PageGame bool
	PageUser bool
	PageLeaderboards bool
}

type Game struct {
	
}

type Lobby struct {
	
}

type Leaderboard struct {
	
}

type PageData struct {
	Session Session
	UserPage UserPage
	Game Game
	Lobby Lobby
	Leaderboard Leaderboard
}

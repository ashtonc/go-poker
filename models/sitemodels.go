package models

import (
	"database/sql"
	"html/template"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"

	"poker/connection"
	"poker/gamelogic"
)

// Env serves as an environment that contains "global" variables. See
// http://www.alexedwards.net/blog/organising-database-access for the idea
type Env struct {
	Database      *sql.DB
	Port          string
	Templates     map[string]*template.Template
	SiteRoot      string
	Games         map[string]*GameListing
	Upgrader      *websocket.Upgrader
	CookieHandler *securecookie.SecureCookie
}

// GameListing lists the information about an instantiated game
type GameListing struct {
	Name        string
	Slug        string
	Status      string
	Ante        int
	MinBet      int
	MaxBet      int
	PlayerCount int
	Game        *gamelogic.Game
	Hub         *connection.Hub
}

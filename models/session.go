package models

import (
	"database/sql"
	"html/template"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

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
	//password here now refers to the password hash
}

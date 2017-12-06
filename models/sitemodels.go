package models

import (
	"database/sql"
	"html/template"
	"time"

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
}

type UserAccount struct {
	Username string
	Name     string
	Email    string
	Password string
	//password here now refers to the password hash
}

type Session struct {
	Id				int
	Uuid			string
	Email			string
	UserId			int
	Username        string
	Name			string
	Expiry          time.Time
	LoggedIn        bool
	PageHome        bool
	PageUser        bool
	PageGame        bool
	PageLeaderboard bool
	PageLogin       bool
	PageRegister    bool
}

type UserIDSearch struct {
	UserId			int
}

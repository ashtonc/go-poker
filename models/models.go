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
	PictureUrl     string
}

type UserAccount struct {
	Username 		string
	Name			string
	Email			string
	Password		string
}

/*
 *    Game models
 */

type Game struct {
	Name    string
	Stakes  GameStakes
	Phase   GamePhase
	Deck    GameDeck
	Player1 GamePlayer
	Player2 GamePlayer
	Player3 GamePlayer
	Player4 GamePlayer
	Player5 GamePlayer
	Player6 GamePlayer
}

type GamePhase struct {
	Phase int
}

type GameDeck struct {
	Cards [52]Card
}

type Card struct {
	Face string
	Suit string
	Rank int
}

type GameStakes struct {
	Ante   int64
	MaxBet int64
	MinBet int64
}

type GamePlayer struct {
	Cash     int64
	Hand     GameHand
	Username string
}

type GameHand struct {
	Card1 Card
	Card2 Card
	Card3 Card
	Card4 Card
	Card5 Card
	PictureURL     string
}

/*
 *    Template models
 */

type PageData struct {
	SiteRoot    string
	Session     Session
	UserPage    UserPage
	Game        gamelogic.Game
	Lobby       Lobby
	Leaderboard Leaderboard
}

type Lobby struct {
	Empty bool
	Games []LobbyListing
}

type LobbyListing struct {
	Name    string
	Stakes  gamelogic.GameStakes
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

package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Env will serve as an environment that contains "global" variables. See
// http://www.alexedwards.net/blog/organising-database-access for the idea
type Env struct {
	Database *sql.DB
	// template cache middleware
	// logger middleware
	// authentication middleware
}

type Session struct {
	LoggedIn bool
	Username string
	Name string
	PageHome bool
	PageLogin bool
	PageRegister bool
	PageUser bool
	PageGame bool
	PageLeaderboard bool
}

type UserPage struct {
	MatchesSession bool
	Username string
	Name string
	Email string
	PictureUrl string
}

type Card struct {
	// each type of card or a string or whatever
}

type GameDeck struct {
	Cards [52]Card
}

type GameHand struct {
	Card1 Card
	Card2 Card
	Card3 Card
	Card4 Card
	Card5 Card
}

type GameSeat struct {
	Cash int64
	Hand GameHand
	// User that they're connected to
}

type GameStakes struct {
	Ante int64
	MaxBet int64
	MinBet int64
}

type GamePhase struct {
	
}

type Game struct {
	Name string
	Stakes GameStakes
	Phase GamePhase
	Deck GameDeck
	Player1 GameSeat
	Player2 GameSeat
	Player3 GameSeat
	Player4 GameSeat
	Player5 GameSeat
	Player6 GameSeat
}

type GameListing struct {
	Name string
	Stakes GameStakes
	Players int
}

type Lobby struct {
	Games []GameListing
}

type LeaderboardEntry struct {
	Username string
	Cash int64
}

type Leaderboard struct {
	Entries []LeaderboardEntry
}

type PageData struct {
	Session Session
	UserPage UserPage
	Game Game
	Lobby Lobby
	Leaderboard Leaderboard
}

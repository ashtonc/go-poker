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
	
}

type GameDeck struct {
	Cards []Card
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
}

type GameStakes struct {
	Ante int64
	MaxBet int64
	MinBet int64
}

type GamePhase struct {
	
}

type Game struct {
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

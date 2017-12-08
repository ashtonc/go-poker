package models

import (
	"poker/gamelogic"
)

const (
	// Account type constants
	TYPE_ADMIN_ACCOUNT = 1
	TYPE_USER_ACCOUNT  = 2

	// Navigation bar constants
	NAVIGATION_HOME        = 1
	NAVIGATION_GAME        = 2
	NAVIGATION_LEADERBOARD = 3
	NAVIGATION_USER        = 4
	NAVIGATION_LOGIN       = 5
	NAVIGATION_ADMIN       = 6
)

// The pagedata struct is used to fill data into the templates
type PageData struct {
	SiteRoot        string
	Identity        *Identity
	UserPage        *UserPage
	Lobby           *Lobby
	GamePage        *GameListing
	Leaderboard     *Leaderboard
	NavigationLevel int
}

// An identity is almost like a session. It is used to populate the template with information specific to the the user. It can be used to personalize the contents of pages.
type Identity struct {
	LoggedIn    bool
	AccountType int
	Username    string
	PictureSlug string
	Name        string
}

// UserPage contains the information that will be displayed on the individual user pages
type UserPage struct {
	MatchesIdentity bool
	Username        string
	Name            string
	Email           string
	Description     string
	PictureSlug     string
	HandsPlayed     int
	BestHand        []*gamelogic.Card
}

// Lobby is a wrapper for game listings to make code easier to read (rather than len(games))
type Lobby struct {
	Empty bool
	Games []*GameListing
}

// Leaderboard wraps LeaderboardEntry
type Leaderboard struct {
	Empty   bool
	Entries []*LeaderboardEntry
}

// LeaderboardEntry is a single entry on the leaderboard
type LeaderboardEntry struct {
	Username string
	Cash     int
	BestHand []*gamelogic.Card
}

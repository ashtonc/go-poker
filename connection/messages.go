package connection

import (
	"poker/gamelogic"
)

const (
	KindGameState = iota + 1

	KindPlayerSits
	KindPlayerLeaves
	KindTimedOut

	KindTakeSeat
	KindLeaveSeat

	KindCheck
	KindBet
	KindCall
	KindFold
	KindDiscard
)

// Server to players

type GamestateMessage struct {
	Kind  int             `json:"kind"`
	State *gamelogic.Game `json:"game"`
}

func GetGamestate(game *gamelogic.Game) *GamestateMessage {
	return &GamestateMessage{
		Kind:  KindGameState,
		State: game,
	}
}

type PlayerSitMessage struct {
	Kind int             `json:"kind"`
	Seat *gamelogic.Seat `json:"seatid"`
}

func NewPlayerSits(buyin int) *PlayerSitMessage {

	// Add player to the game

	return &PlayerSitMessage{
		Kind: KindPlayerSits,
		// Seat from the updated game object
	}

}

type PlayerLeaveMessage struct {
	Kind int `json:"kind"`
	Seat int `json:"seatid"`
}

type TimedOutMessage struct {
	Kind int `json:"kind"`
	Seat int `json:"seatid"`
}

// Player to server

type SitMessage struct {
	Kind  int `json:"kind"`
	Seat  int `json:"seatid"`
	BuyIn int `json:"buyin"`
}

type LeaveMessage struct {
	Kind int `json:"kind"`
}

type CheckMessage struct {
	Kind int `json:"kind"`
}

type BetMessage struct {
	Kind int `json:"kind"`
	Bet  int `json:"bet"`
}

type CallMessage struct {
	Kind int `json:"kind"`
}

type FoldMessage struct {
	Kind int `json:"kind"`
}

type DiscardMessage struct {
	Kind  int  `json:"kind"`
	Card1 bool `json:"discard1"`
	Card2 bool `json:"discard2"`
	Card3 bool `json:"discard3"`
	Card4 bool `json:"discard4"`
	Card5 bool `json:"discard5"`
}

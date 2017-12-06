package gamelogic

import (
	"encoding/json"
)

func (g Game) GetJsonBySeat(seat int) json.RawMessage {

	return nil
}

type JsonGame struct {
	Pot   int     `json: "pot"`
	Seats [6]Seat `json: "seats"`
}

type JsonSeat struct {
	Number   int     `json: "number"`
	Occupied bool    `json: "occupied"`
	Player   *Player `json: "player"`
}

type JsonPlayer struct {
	Username    string `json: "username"`
	PictureSlug string `json: "pictureslug"`
	Money       int    `json: "money"`
	Hand        []Card `json: "hand"`
	Folded      bool   `json: "folded"`
}

type JsonCard struct {
	Face string `json: "face"`
	Suit string `json: "suit"`
}

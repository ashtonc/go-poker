package gamelogic

import (
	"encoding/json"
)

func (g Game) GetJsonBySeat(seat int) ([]byte, error) {

	gameJson, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	return gameJson, err
}

package gamelogic

import (
	"poker/database"
	"poker/models"
)

func InitializeGames(env *models.Env) {
	games := database.GetGames()
	gameMap := make(map[string]*gamelogic.Game)

	for game := range games {
		game.Init()
		m[game.Slug] = game
	}

	env.Games = gameMap
}

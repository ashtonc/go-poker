package database

import (
	_ "github.com/lib/pq"

	"poker/models"
)

func GetUserPage(env *models.Env, userName string) (*models.UserPage, error) {
	var page models.UserPage
	page.Username = userName

	sqlStatement := `SELECT name, email, picture FROM player WHERE username=$1;`

	row := env.Database.QueryRow(sqlStatement, page.Username)
	err := row.Scan(&page.Name, &page.Email, &page.Email, &page.PictureUrl)

	return &page, err
}

/*
func GetGame(env *models.Env, gameId int) (*models.Game, error) {

}

func GetLeaderboard(env *models.Env) (*models.Leaderboard, error) {

}

func GetLobby(env *models.Env) (*models.Lobby, error) {

}
*/

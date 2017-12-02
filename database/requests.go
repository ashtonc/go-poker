package database

import (
	_ "github.com/lib/pq"
	"log"

	"poker/models"
)

func GetUserPage(env *models.Env, userName string) (*models.UserPage, error) {
	var page models.UserPage
	page.Username = userName

	sqlStatement := `SELECT name, email, picture FROM user WHERE username=$1;`

	row := env.Database.QueryRow(sqlStatement, page.Username)
	err := row.Scan(&page.Name, &page.Email, &page.Email, &page.PictureUrl)

	return &page, err
}

/*
func GetGame(env *models.Env, gameId int) (*models.Game, error) {

}
*/

func GetLeaderboard(env *models.Env) (*models.Leaderboard, error) {
	var leaderboard models.Leaderboard

	sqlStatement := `SELECT username, total_cash FROM player_stats, account WHERE player_stats.user_id = account.id;`

	rows, err := env.Database.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var username string
		var cash int64
		err = rows.Scan(&username, &cash)
		if err != nil {
			log.Fatal(err)
		}
		leaderboard.Entries = append(leaderboard.Entries, models.LeaderboardEntry{Username: username, Cash: cash})
	}

	if len(leaderboard.Entries) > 0 {
		leaderboard.Empty = false
	} else {
		leaderboard.Empty = true
	}

	return &leaderboard, err
}

func GetLobby(env *models.Env) (*models.Lobby, error) {
	var lobby models.Lobby

	sqlStatement := `SELECT name, players FROM game;`
	// , game_status WHERE game_status.description = 'open'

	rows, err := env.Database.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var players int
		err = rows.Scan(&name, &players)
		if err != nil {
			log.Fatal(err)
		}
		lobby.Games = append(lobby.Games, models.LobbyListing{Name: name, Players: players})
	}

	if len(lobby.Games) > 0 {
		lobby.Empty = false
	} else {
		lobby.Empty = true
	}

	return &lobby, err
}

// func UserLogin(env *models.Env, userName string) (*models.UserPage, error) {
// 	var users models.UserPage
// 	users.Username = userName
// 	// page.Password = password

// 	sqlStatement := `SELECT * FROM user;`

// 	row := env.Database.QueryRow(sqlStatement, "ghth")
// 	err := row.Scan(&users.Username, &users.Name, &users.Email, &users.PictureUrl)

// 	return &users, err
// }

func UserRegister(env *models.Env, userName string) (*models.UserPage, error) {
	var users models.UserPage
	users.Username = userName

	sqlStatement := `  
	INSERT INTO account (username, name, email) 
	VALUES ($1, $2, $3)`
	_, err := env.Database.Exec(sqlStatement, "username!", "Jonathan", "fff@f.com")
	if err != nil {
		panic(err)
	}

	return &users, err
}

// Temporary function that adds entries to the game database
func CreateLobbyEntries(env *models.Env) error {
	// var leaderboard models.Leaderboard

	sqlStatement := `  
	INSERT INTO game (name) 
	VALUES ($1)`
	_, err := env.Database.Exec(sqlStatement, "my name")
	if err != nil {
		panic(err)
	}

	return err
}

// Temporary function that adds entries to the game database
func CreateLeaderboardEntries(env *models.Env) error {
	// var leaderboard models.Leaderboard

	sqlStatement := `  
	INSERT INTO player_stats (total_hands) 
	VALUES ($1)`
	_, err := env.Database.Exec(sqlStatement, 1000)
	if err != nil {
		panic(err)
	}

	return err
}

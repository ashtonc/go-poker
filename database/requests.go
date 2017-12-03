package database

import (
	// "database/sql"
	"log"

	// Wraps database/sql for our postgres database
	_ "github.com/lib/pq"

	"poker/gamelogic"
	"poker/models"
)

func GetUserPage(env *models.Env, userName string) (*models.UserPage, error) {
	var page models.UserPage
	page.Username = userName

	sqlStatement := `SELECT name, email FROM user WHERE username=$1;`

	row := env.Database.QueryRow(sqlStatement, page.Username)
	err := row.Scan(&page.Name, &page.Email, &page.Email, &page.PictureUrl)

	return &page, err
}

func GetGames(env *models.Env) ([]*gamelogic.Game, error) {
	var games []*gamelogic.Game

	/*
		sqlStatement := `SELECT game.name, game_stakes.ante, etc FROM game, game_stakes, game_status WHERE ...;`

		rows, err := env.Database.Query(sqlStatement)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			//create some vars here
			err = rows.Scan(&*var*, &*var*)
			if err != nil {
				log.Fatal(err)
			}
			games = append(games, *game object*)
		}
	*/

	return games, nil
}

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

	sqlStatement := `SELECT name FROM game;`
	// , game_status WHERE game_status.description = 'open'

	rows, err := env.Database.Query(sqlStatement)
	if err != nil {
		//	log.Fatal(err)
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

func UserRegister(env *models.Env, username string, name string, email string, password string) error {

	sqlStatement := `  
	INSERT INTO account (username, name, email, password) 
	VALUES ($1, $2, $3, $4)`
	_, err := env.Database.Exec(sqlStatement, username, name, email, password)
	if err != nil {
		panic(err)
	}
	return err
}

func FindByUsername(env *models.Env, inputUsername string) (models.UserAccount) {
	var userAccount models.UserAccount

	sqlStatement := `SELECT username, name, email, password FROM account WHERE username=$1`

	rows, err := env.Database.Query(sqlStatement, inputUsername)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&userAccount.Username, &userAccount.Name, &userAccount.Email, &userAccount.Password)
		if err != nil {
			panic(err)
		}
	}

	return userAccount
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

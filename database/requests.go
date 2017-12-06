package database

import (
	"log"
	"time"

	_ "github.com/lib/pq"

	"poker/models"
)

func GetUserPage(env *models.Env, userName string) (*models.UserPage, error) {
	var page models.UserPage

	sqlStatement := `SELECT name, email, description, picture_slug FROM account WHERE username = $1;`

	row := env.Database.QueryRow(sqlStatement, userName)
	err := row.Scan(&page.Name, &page.Email, &page.Description, &page.PictureSlug)
	page.Username = userName

	return &page, err
}

func GetGames(env *models.Env) (map[string]*models.GameListing, error) {
	gameMap := make(map[string]*models.GameListing)

	sqlStatement := `SELECT game.name, game.slug, game_stakes.ante, game_stakes.min_bet, game_stakes.max_bet, game_status.description FROM game, game_stakes, game_status WHERE game.game_status = game_status.id AND game.stakes = game_stakes.id;`

	rows, err := env.Database.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var gameListing models.GameListing
		gameListing.PlayerCount = 0

		err = rows.Scan(&gameListing.Name, &gameListing.Slug, &gameListing.Ante, &gameListing.MinBet, &gameListing.MaxBet, &gameListing.Status)
		if err != nil {
			log.Fatal(err)
		}

		gameMap[gameListing.Slug] = &gameListing
	}

	return gameMap, err
}

func GetSession(env *models.Env, sessionid string) (*models.Session, error) {
	var session models.Session

	sqlStatement := `SELECT account.username, user_session.expiry_time FROM account, user_session WHERE account.id = user_session.user_id AND user_session.token = $1;`

	row := env.Database.QueryRow(sqlStatement, sessionid)
	err := row.Scan(&session.Username, &session.Expiry)

	return &session, err
}

func GetLeaderboard(env *models.Env) (*models.Leaderboard, error) {
	var leaderboard models.Leaderboard

	sqlStatement := `SELECT username, total_cash FROM account, player_stats WHERE player_stats.user_id = account.id;`

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
		leaderboard.Entries = append(leaderboard.Entries, &models.LeaderboardEntry{Username: username, Cash: cash})
	}

	if len(leaderboard.Entries) > 0 {
		leaderboard.Empty = false
	} else {
		leaderboard.Empty = true
	}

	return &leaderboard, err
}

func UserRegister(env *models.Env, username string, name string, email string, password_hash string) error {

	sqlStatement := `  
	INSERT INTO account (username, name, email, password_hash) 
	VALUES ($1, $2, $3, $4)`
	_, err := env.Database.Exec(sqlStatement, username, name, email, password_hash)
	if err != nil {
		panic(err)
	}
	return err
}

func FindByUsername(env *models.Env, inputUsername string) models.UserAccount {
	var userAccount models.UserAccount

	sqlStatement := `SELECT username, name, email, password_hash FROM account WHERE username=$1`

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

// Saves the already created sessions object in the database
func AddSessionData(env *models.Env, session models.Session) error {

	

	/*
	This function should only create an entry in this table:

	CREATE TABLE user_session (
		id SERIAL PRIMARY KEY,
		token VARCHAR(256),
		expiry_time TIMESTAMP,
		user_id INTEGER REFERENCES account (id)
	);
	*/

	return nil
}

func GetUserID(env *models.Env, inputUsername string) models.UserIDSearch {
	var userData models.UserIDSearch

	sqlStatement := `SELECT id FROM account WHERE username=$1`

	rows, err := env.Database.Query(sqlStatement, inputUsername)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&userData.UserId)
		if err != nil {
			panic(err)
		}
	}

	return userData
}

func NewSession(env *models.Env, token string, expiry_time time.Time, user_id int) error {
	sqlStatement := `  
	INSERT INTO user_session (token, expiry_time, user_id) 
	VALUES ($1, $2, $3)`
	_, err := env.Database.Exec(sqlStatement, token, expiry_time, user_id)
	if err != nil {
		panic(err)
	}
	return err
}
package database

import (
	"log"

	_ "github.com/lib/pq"

	"poker/models"
)

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
		var cash int
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

func GetUserPage(env *models.Env, userName string) (*models.UserPage, error) {
	var page models.UserPage

	sqlStatement := `SELECT name, email, description, picture_slug FROM account WHERE username = $1;`

	row := env.Database.QueryRow(sqlStatement, userName)
	err := row.Scan(&page.Name, &page.Email, &page.Description, &page.PictureSlug)
	page.Username = userName

	return &page, err
}

func GetUser(env *models.Env, username string) (*models.User, error) {
	var user models.User
	user.Username = username

	sqlStatement := `SELECT name, email, description, picture_slug, password_hash FROM account WHERE username = $1;`

	row := env.Database.QueryRow(sqlStatement, username)
	err := row.Scan(&user.Name, &user.Email, &user.PictureSlug, &user.Description, &user.HashedPassword)

	return &user, err
}

func AddUser(env *models.Env, user *models.User) error {
	sqlStatement := `INSERT INTO account (username, name, email, picture_slug, description, password_hash) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := env.Database.Exec(sqlStatement, user.Username, user.Name, user.Email, user.PictureSlug, user.Description, user.HashedPassword)
	if err != nil {
		log.Print(err)
	}

	return err
}

func UpdateUser(env *models.Env, user *models.User) error {
	sqlStatement := `UPDATE account SET (name, email, description, password_hash) = ($1, $2, $3, $4) WHERE username = $5`

	_, err := env.Database.Exec(sqlStatement, user.Name, user.Email, user.Description, user.HashedPassword, user.Username)
	if err != nil {
		log.Print(err)
	}

	return err
}

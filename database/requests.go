package database

import (
	_ "github.com/lib/pq"
	// "log"

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

/*func GetGame(env *models.Env, gameId int) (*models.Game, error) {

}

func GetLeaderboard(env *models.Env) (*models.Leaderboard, error) {

}*/

// commented out unfinished work
// func GetLobby(env *models.Env) (*models.Lobby, error) {
// 	var lobby models.Lobby

// 	sqlStatement := `SELECT name, users FROM game;`

// 	rows, err := env.Database.Query(sqlStatement)
// 	if err != nil {
// 		log.Fatal(err
// )	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var name string
// 		var players int
// 		err = rows.Scan(&name, &players)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		lobby.Games = append(models.LobbyListing{Name: name, Players: players})
// 	}

// 	if true {
// 		lobby.Empty = False
// 	}

// 	return lobby, err
// }


func UserLogin(env *models.Env, userName string) (*models.UserPage, error) {
	var users models.UserPage
	users.Username = userName
	// page.Password = password

	sqlStatement := `SELECT * FROM user WHERE Username=$1;`

	row := env.Database.QueryRow(sqlStatement, "ghth")
	err := row.Scan(&users.Username, &users.Name, &users.Email, &users.PictureUrl)

	return &users, err
}


// func UserRegister(env *models.Env, userName string, password string) (*models.user, error) {
// 	var users models.users
// 	user.Username = userName

// 	sqlStatement := `INSERT INTO users ();`

// 	row := env.Database.QueryRow(sqlStatement, page.Username)
// 	err := row.Scan(&page.Name, &page.Email, &page.Email, &page.PictureUrl)

// 	return &page, err
// }
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
	var page models.UserPage
	page.Username = userName

	sqlStatement := `SELECT name, email, picture FROM user WHERE username=$1;`

	row := env.Database.QueryRow(sqlStatement, page.Username)
	err := row.Scan(&page.Name, &page.Email, &page.Email, &page.PictureUrl)

	return &page, err
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

	// sqlStatement := `
	// INSERT INTO account(username, name, email)  
	// VALUES ('username', 'name', 'email');  
	// RETURNING id;`  
	// id := 0 	
	// err := env.Database.Query(sqlStatement)
	// // err := env.Database.QueryRow(sqlStatement, 'username!', 'name!', 'email!').Scan(&id)
	// if err == nil {
	// panic(err)
	// }
	

	//row := env.Database.QueryRow(sqlStatement, page.Username)
	//err := row.Scan(&page.Name, &page.Email, &page.Email, &page.PictureUrl)

	return &users, err
}
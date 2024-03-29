package database

import (
	"database/sql"
	"log"

	// Wraps database/sql for postgres
	_ "github.com/lib/pq"

	"poker/connection"
	"poker/gamelogic"
	"poker/models"
)

func CreateDatabase(username string, password string, name string) (database *sql.DB, err error) {
	log.Print("Connecting to the database server...")

	// Save information about our database in a string
	dbinfo := "user=" + username + " password=" + password + " dbname=" + name + " sslmode=disable"

	// Open the database (using the postgres driver) and pass in the database info we saved earlier
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	// Check whether or not the database is running (db.Open only validates arguments)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeGames(env *models.Env, gameMap map[string]*models.GameListing) {
	log.Print("Initializing game objects from database...")

	for _, listing := range gameMap {
		game, err := gamelogic.GameInit(listing.Ante, listing.MinBet, listing.MaxBet)
		if err != nil {
			log.Fatal(err)
		}

		listing.Game = game

		// Create a new websockets hub
		hub := connection.NewHub()
		listing.Hub = hub
	}

	env.Games = gameMap
}

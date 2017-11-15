package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"poker/database"
	"poker/handlers"
	"poker/models"
	//"poker-project/gamelogic"
)

func main() {
	// Session variables. Should probably be included in the environment or
	// taken from a config file somewhere.
	db_user := "postgres"
	db_password := "postgres"
	db_name := "pokerdb"
	server_port := "8000"

	// Populate our environment
	database, err := database.CreateDatabase(db_user, db_password, db_name)
	if err != nil {
		//log.Fatal(err)
	}

	env := &models.Env{Database: database}

	// Create a new router and initialize the handlers
	router := mux.NewRouter()

	router.Handle("/", handlers.HomeRedirect(env))
	router.Handle("/poker/", handlers.Home(env))
	router.Handle("/poker/user/{username:[0-9]+}", handlers.User(env))
	router.Handle("/poker/game/", handlers.Game(env))

	// These functions are deferred until main finishes
	defer env.Database.Close()

	// Start the server
	log.Print("Running server on port " + server_port + ".")
	log.Fatal(http.ListenAndServe(":"+server_port, router))
}

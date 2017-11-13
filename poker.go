package main

import (
	"log"
	"net/http"

	"poker-project/database"
	"poker-project/handlers"
	"poker-project/models"
	//"poker-project/gamelogic"

	"github.com/gorilla/mux"
)

func main() {
	// Session variables. Should probably be included in the environment or
	// taken from a config file somewhere.
	var db_user = "postgres"
	var db_password = "postgres"
	var db_name = "pokerdb"
	var server_port = "8000"

	// Populate environment
	database, err := database.CreateDatabase(db_user, db_password, db_name)
	if err != nil {
		log.Fatal(err)
	}

	env := &models.Env{Database: database}
	
	// Create a new router and initialize the handlers
	router := mux.NewRouter()
	router.Handle("/", handlers.HomeRedirect(env))
	
	// These functions will run when main finishes
	defer env.Database.Close()

	// Start the server
	log.Print("Running server on port " + server_port + ".")
	log.Fatal(http.ListenAndServe(":"+server_port, router))
}

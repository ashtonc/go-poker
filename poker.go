package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/gorilla/securecookie"
	//"github.com/gorilla/websockets"

	"poker/database"
	"poker/handlers"
	"poker/models"
	"poker/templates"
	//"poker/gamelogic"
)

func main() {
	// Connect to the database
	db_user := "postgres"
	db_password := "postgres"
	db_name := "pokerdb"

	database, err := database.CreateDatabase(db_user, db_password, db_name)
	if err != nil {
		log.Fatal(err)
	}

	// Generate our templates
	templates := templates.BuildTemplates()

	// Populate our environment
	env := &models.Env{
		Database:  database,
		Port:      "8000",
		Templates: templates,
	}

	// Create a new router and initialize the handlers
	router := mux.NewRouter()

	router.Handle("/", handlers.HomeRedirect(env))
	router.Handle("/poker/", handlers.Home(env))
	router.Handle("/poker/login/", handlers.Login(env))
	router.Handle("/poker/logout/", handlers.Logout(env))
	router.Handle("/poker/register/", handlers.Register(env))
	router.Handle("/poker/user/{username:[A-Za-z0-9-_.]+}", handlers.ViewUser(env))
	router.Handle("/poker/user/{username:[A-Za-z0-9-_.]+}/edit", handlers.EditUser(env))
	router.Handle("/poker/game/", handlers.RedirectGame(env))
	router.Handle("/poker/game/play", handlers.PlayGame(env))
	router.Handle("/poker/game/lobby", handlers.ViewLobby(env))
	router.Handle("/poker/game/watch", handlers.WatchGame(env))
	router.Handle("/poker/leaderboard/", handlers.Leaderboard(env))

	// These functions are deferred until main finishes
	defer env.Database.Close()

	// Start the server
	log.Print("Running server on port " + env.Port + ".")
	log.Fatal(http.ListenAndServe(":"+env.Port, router))
}

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

	// Create a template cache
	templates := templates.BuildTemplateCache()

	// Populate our environment
	env := &models.Env{
		Database:  database,
		Port:      ":8000",
		Templates: templates,
		SiteRoot:  "/poker",
	}

	// Deferred until main finishes, idiomatic
	defer env.Database.Close()

	// Create a new router and initialize the handlers
	router := mux.NewRouter()

	router.Handle("/", handlers.HomeRedirect(env))
	router.Handle(env.SiteRoot+"/", handlers.Home(env))
	router.Handle(env.SiteRoot+"/login/", handlers.Login(env))
	router.Handle(env.SiteRoot+"/logout/", handlers.Logout(env))
	router.Handle(env.SiteRoot+"/register/", handlers.Register(env))
	router.Handle(env.SiteRoot+"/user/{username:[A-Za-z0-9-_.]+}", handlers.ViewUser(env))
	router.Handle(env.SiteRoot+"/user/{username:[A-Za-z0-9-_.]+}/edit", handlers.EditUser(env))
	router.Handle(env.SiteRoot+"/game/", handlers.RedirectGame(env))
	router.Handle(env.SiteRoot+"/game/play", handlers.PlayGame(env))
	router.Handle(env.SiteRoot+"/game/lobby", handlers.ViewLobby(env))
	router.Handle(env.SiteRoot+"/game/watch", handlers.WatchGame(env))
	router.Handle(env.SiteRoot+"/leaderboard/", handlers.Leaderboard(env))

	// Start the server
	log.Print("Running server on port " + env.Port + ".")
	log.Fatal(http.ListenAndServe(env.Port, router))
}

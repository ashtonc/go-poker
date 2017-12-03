package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"poker/database"
	"poker/gamelogic"
	"poker/handlers"
	"poker/models"
	"poker/templates"
)

func main() {
	// Connect to the database
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "pokerdb"

	database, err := database.CreateDatabase(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer env.Database.Close()

	// Create a template cache
	templates := templates.BuildTemplateCache()

	// Populate our environment
	env := &models.Env{
		Database:  database,
		Port:      ":8000",
		Templates: templates,
		SiteRoot:  "/poker",
	}

	// Initialize the games found in the database (imagine these as poker tables)
	gamelogic.InitializeGames(env)

	// Create a new router and initialize the handlers
	router := mux.NewRouter()

	router.Handle("/", handlers.HomeRedirect(env))
	router.Handle(env.SiteRoot+"/", handlers.Home(env))
	router.Handle(env.SiteRoot+"/login/", handlers.Login(env))
	router.Handle(env.SiteRoot+"/logout/", handlers.Logout(env))
	router.Handle(env.SiteRoot+"/register/", handlers.Register(env))
	router.Handle(env.SiteRoot+"/user/{username:[A-Za-z0-9-_.]+}/view", handlers.ViewUser(env))
	router.Handle(env.SiteRoot+"/user/{username:[A-Za-z0-9-_.]+}/edit", handlers.EditUser(env))
	router.Handle(env.SiteRoot+"/lobby/", handlers.ViewLobby(env))
	router.Handle(env.SiteRoot+"/game/", handlers.RedirectGame(env))
	router.Handle(env.SiteRoot+"/game/{gameid:[a-z0-9-]+}/play", handlers.PlayGame(env))
	router.Handle(env.SiteRoot+"/game/{gameid:[a-z0-9-]+}/watch", handlers.WatchGame(env))
	router.Handle(env.SiteRoot+"/leaderboard/", handlers.Leaderboard(env))

	// Start the server
	log.Print("Running server at " + env.SiteRoot + " on port " + env.Port + ".")
	log.Fatal(http.ListenAndServe(env.Port, router))
}

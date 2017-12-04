package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"poker/database"
	"poker/handlers"
	"poker/models"
	"poker/templates"
)

func main() {
	// Connect to the database
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "pokerdb"

	db, err := database.CreateDatabase(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}

	// Create a template cache
	templates := templates.BuildTemplateCache()

	// Create a websockets upgrader
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Populate our environment
	env := &models.Env{
		Database:  db,
		Port:      ":8000",
		Templates: templates,
		SiteRoot:  "/poker",
		Upgrader:  &upgrader,
	}

	// Close the database after main finishes
	defer env.Database.Close()

	// Initialize the games found in the database (imagine these as tables)
	games, err := database.GetGames(env)
	if err != nil {
		log.Fatal(err)
	}

	database.InitializeGames(env, games)

	// Create a new router and initialize the handlers
	router := mux.NewRouter()

	router.Handle("/", handlers.HomeRedirect(env))
	router.Handle(env.SiteRoot+"/", handlers.Home(env))
	router.Handle(env.SiteRoot+"/login/", handlers.Login(env))
	router.Handle(env.SiteRoot+"/logout/", handlers.Logout(env))
	router.Handle(env.SiteRoot+"/register/", handlers.Register(env))
	router.Handle(env.SiteRoot+"/user/{username:[A-Za-z0-9-_.]+}/{action:view|edit}", handlers.User(env))
	router.Handle(env.SiteRoot+"/lobby/", handlers.ViewLobby(env))
	router.Handle(env.SiteRoot+"/game/", handlers.RedirectGame(env))
	router.Handle(env.SiteRoot+"/game/{gameslug:[a-z0-9-]+}/{action:play|watch}", handlers.Game(env))
	router.Handle(env.SiteRoot+"/leaderboard/", handlers.Leaderboard(env))
	router.Handle(env.SiteRoot+"/ws", handlers.WebsocketConnection(env))

	// Start the server
	log.Print("Running server at " + env.SiteRoot + " on port " + env.Port + ".")
	log.Fatal(http.ListenAndServe(env.Port, router))
}

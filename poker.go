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
    //"poker/gamelogic"
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
        log.Fatal(err)
    }

    env := &models.Env{Database: database}

    // Create a new router and initialize the handlers
    router := mux.NewRouter()

    router.Handle("/", handlers.HomeRedirect(env))
    router.Handle("/poker/", handlers.Home(env))
    router.Handle("/poker/register/", handlers.Register(env))
    router.Handle("/poker/user/{username:[A-Za-z0-9-_.]+}", handlers.User(env))
    router.Handle("/poker/user/{username:[A-Za-z0-9-_.]+}/edit", handlers.UserEdit(env))
	router.Handle("/poker/game/", handlers.Lobby(env))
	router.Handle("/poker/game/lobby", handlers.Lobby(env))
	router.Handle("/poker/game/play", handlers.PlayGame(env))
	router.Handle("/poker/game/watch", handlers.ViewGame(env))
    router.Handle("/poker/leaderboard/", handlers.Leaderboard(env))
    //router.HandleFunc("/poker/home", indexPageHandler)
    //router.HandleFunc("/poker/internal", internalPageHandler)
   //router.Handle("/poker/login/", handlers.Login(env))
    router.HandleFunc("/poker/login/", handlers.LoginGET).Methods("GET")
    router.HandleFunc("/poker/login/", handlers.LoginPOST).Methods("POST")
    //router.HandleFunc("/logout", logoutHandler).Methods("POST")

    // These functions are deferred until main finishes
    defer env.Database.Close()

    // Start the server
    log.Print("Running server on port " + server_port + ".")
    log.Fatal(http.ListenAndServe(":"+server_port, router))
}

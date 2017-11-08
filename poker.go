package main

import (
	"log"
	"net/http"

	"poker-project/database"
	"poker-project/gamelogic"

	"github.com/gorilla/mux"
)

const (
	PORT = "8000"
)

func main() {
	database.InitializeDatabase()

	// Create a new router
	r := mux.NewRouter()

	// -- Handlers here --

	// Start the server
	log.Print("Running server on port " + PORT + ".")
	log.Fatal(http.ListenAndServe(":"+PORT, r))

	gamelogic.ConnectToGame()
}

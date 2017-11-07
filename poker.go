package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

const (
	PORT        = "8000"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "pokerdb"
)

func main() {
	r := mux.NewRouter()
	
	// Create a new database and save information about the database in a string
	var db *sql.DB
	dbinfo := "user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " sslmode=disable"

	// Open the database (using the postgres driver) and pass in the database info we saved earlier
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	
	// Close the database when main() finishes
	defer db.Close()

	// Check whether or not the database is running (db.Open only validates arguments)
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	
	// -- Handlers here --
	
	
	// Start the server
	log.Print("Running server on port " + PORT + ".")
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
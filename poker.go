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
	
	var db *sql.DB
	dbinfo := "user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " sslmode=disable"

	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	
	// Handlers here
	
	log.Print("Running server on port " + PORT + ".")
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "pokerdb"
)

func InitializeDatabase() {
	// todo: don't use DB globals: http://www.alexedwards.net/blog/organising-database-access

	// Create a new database and save information about the database in a string
	var db *sql.DB
	dbinfo := "user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " sslmode=disable"

	// Open the database (using the postgres driver) and pass in the database info we saved earlier
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	// Close the database when this process finishes
	defer db.Close()

	// Check whether or not the database is running (db.Open only validates arguments)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

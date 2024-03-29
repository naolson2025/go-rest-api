package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "api.db"
	}

	DB, err = sql.Open("sqlite3", dbName)

	if err != nil {
		panic("Could not connect to database!")
	}

	DB.SetMaxOpenConns(10)
	// keep 5 idle connections in the pool
	// that way when a request comes in, we don't have to wait for a new connection to be established
	DB.SetMaxIdleConns(5)
	fmt.Println("Connected to database!")
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userID INTEGER,
		FOREIGN KEY(userID) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		eventID INTEGER,
		userID INTEGER,
		FOREIGN KEY(eventID) REFERENCES events(id),
		FOREIGN KEY(userID) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create registrations table")
	}
}

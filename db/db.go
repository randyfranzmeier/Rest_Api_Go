package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Cannot open the database driver!!!")
	}

	DB.SetMaxOpenConns(10) //limits number of connections
	DB.SetMaxIdleConns(5)
	CreateTables()
}

func CreateTables() {
	createUsersTable := `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL 
                                 )
`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
    )
    `
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationTable := `
CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    eventID INTEGER,
    userID Integer,
    FOREIGN KEY(eventID) REFERENCES events(id),
    FOREIGN KEY(userID) REFERENCES users(id)
    )
`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("registration table creation failed")
	}

}

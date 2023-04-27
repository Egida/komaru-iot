package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"server/core/config"
)

type Database struct {
	Instance *sql.DB
}

var Connection *Database

func Serve() {
	conn, err := sql.Open("sqlite3", "cnc.sqlite")
	// open connection to the database
	if err != nil {
		config.Logger.Error("Failed to open database", "error", err)
		return
	}

	// set the database connection
	Connection = &Database{
		Instance: conn,
	}

	// create user table
	err = Connection.CreateUserTable()
	if err != nil {
		config.Logger.Error("Failed to create user table", "error", err)
		return
	}

	exists, err := Connection.UserExists("admin")
	if err != nil {
		config.Logger.Error("Failed to check user table", "error", err)
		return
	}

	if !exists {
		err := Connection.CreateUser(&User{
			Name:     "admin",
			Password: "securepassword123",
		})
		if err != nil {
			config.Logger.Error("Failed to insert in user table", "error", err)
			return
		}
		config.Logger.Info("Created default user.", "username", "admin", "password", "securepassword123")
	}

	config.Logger.Info("Opened SQLite3 database.")
}

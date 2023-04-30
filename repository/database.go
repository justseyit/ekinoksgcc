package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

var closed bool = true

const (
	user     = "ekinoks"
	password = "ekinoksgccdb"
	dbname   = "ekinoksdb"
)

func InitDB() {
	if closed {
		var err error
		DB, err = sql.Open("postgres", "user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		closed = false
		log.Println("Database initialized")
	} else {
		log.Println("Database already initialized")
	}
}

func DisposeDB() {
	if !closed {
		DB.Close()
		closed = true
	}
}

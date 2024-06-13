package model

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Piyush@99@(tcp:localhost:6375Piyush299)")

	if err != nil {

		log.Fatal(err)
	}

	return db
}

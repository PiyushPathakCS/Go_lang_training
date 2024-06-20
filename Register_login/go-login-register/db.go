package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() {
	var err error
	Db, err = sql.Open("mysql", "root:Piyush@99@tcp(localhost:3306)/user")
	if err != nil {
		log.Fatal(err)
	} else {
		err = Db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to database!")
	}
}

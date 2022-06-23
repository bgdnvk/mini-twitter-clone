package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Connect() error{
	var err error
	DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitterdb")

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	return nil
	
}
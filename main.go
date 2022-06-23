package main

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id         int
	user       string
	passhash  string
	email      string
	first_name  string
	last_name  string
	dob		string
}


func main() {


	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitterdb")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	var users []User

	for res.Next() {

		var user User
		err := res.Scan(&user.id, &user.user, &user.passhash, &user.email, &user.first_name, &user.last_name, &user.dob)


		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", user)
		//storing the data in a slice
		users = append(users, user)
	}
	
	fmt.Printf("%v\n", users)
	
	json_data, err := json.Marshal(users)

	if  err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%v\n", json.Unmarshal(json_data, &users))

}
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goexample/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)



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

	var users []models.User

	
	for res.Next() {
		
		
		user := models.User{}

		// err := res.Scan(&user.id, &user.user, &user.passhash, &user.email, &user.first_name, &user.last_name, &user.dob)
		err := res.Scan( &user.Id, &user.User, &user.Passhash, &user.Email, &user.First_name, &user.Last_name, &user.Dob)


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
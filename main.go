package main

import (
	// "database/sql"
	// "encoding/json"
	// "fmt"
	// "goexample/models"
	// "log"

	"goexample/services"

	"github.com/gofiber/fiber/v2"
	"goexample/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/users", services.GetUsers)



	// app.Get("/", func(c *fiber.Ctx) error {
    //     return c.JSON(fiber.Map{
    //         "message": "Hello World",
    //     })
    // })

    log.Fatal(app.Listen(":3000"))


	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitterdb")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	// res, err := db.Query("SELECT * FROM users")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer res.Close()

	// var users []models.User

	// for res.Next() {

	// 	user := models.User{}

	// 	err := res.Scan( &user.Id, &user.User, &user.Passhash, &user.Email, &user.First_name, &user.Last_name, &user.Dob)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("%v\n", user)
	// 	//storing the data in a slice
	// 	users = append(users, user)
	// }

	// fmt.Printf("%v\n", users)

	// json_data, err := json.Marshal(users)

	// if  err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%v\n", json.Unmarshal(json_data, &users))

}

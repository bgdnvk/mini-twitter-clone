package services

import(
	"goexample/database"
	"goexample/models"
	"log"
	// "database/sql"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

//get all users from db
func GetUsers(c *fiber.Ctx) error {

	res, err := database.DB.Query("SELECT * FROM users")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		  })
		return nil
	}

	defer res.Close()

	var users []models.User

	for res.Next() {

		user := models.User{}

		err := res.Scan( &user.Id, &user.User, &user.Passhash, &user.Email, &user.First_name, &user.Last_name, &user.Dob)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", user)
		//storing the data in a slice
		users = append(users, user)
	}

	fmt.Printf("%v\n", users)
	if users != nil {
		c.Status(200).JSON(&fiber.Map{
			"success": true,
			"users": users,
		  })
	} else {
		c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error": "No users found",
		})}

	return nil

	// json_data, err := json.Marshal(users)

	// if  err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%v\n", json.Unmarshal(json_data, &users))

	
}
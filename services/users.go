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

	//get the users from db
	res, err := database.DB.Query("SELECT * FROM users")
	//check for errors
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		  })
		return nil
	}
	//close the result set
	defer res.Close()
	//create a slice of users
	var users []models.User
	//loop through the result set
	for res.Next() {

		user := models.User{}

		err := res.Scan( &user.Id, &user.User, &user.Passhash, &user.Email, &user.First_name, &user.Last_name, &user.Dob)

		if err != nil {
			log.Fatal(err)
		}
		//best logger amirite
		fmt.Printf("%v\n", user)
		//storing the data in a slice
		users = append(users, user)
	}

	// fmt.Printf("%v\n", users)
	
	//send the users to the client
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
}
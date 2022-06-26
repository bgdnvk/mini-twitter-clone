package services

import (
	"goexample/database"
	"goexample/models"
	"log"

	// "database/sql"
	"fmt"
	"goexample/services/utils"

	"github.com/gofiber/fiber/v2"
)

//get all users from db
func GetUsers(c *fiber.Ctx) error {

	//get the users from db
	res, err := database.DB.Query("SELECT * FROM users")
	// res, err := database.DB.Query("SELECT user, dob FROM users")

	// check for errors
	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}

	//close the result set
	defer res.Close()
	//create a slice of users
	var users []models.User
	//loop through the result set
	for res.Next() {

		user := models.User{}

		err := res.Scan(&user.Id, &user.User, &user.Passhash, &user.Email, &user.First_name, &user.Last_name, &user.Dob)

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
	utils.ResponseHelperJSON(c, users, "users", "No users found")

	return err
}

func GetUsersByAgeAsc(c *fiber.Ctx) error {
	//get the users from db
	res, err := database.DB.Query("SELECT user_id, user, passhash, email, first_name, last_name, TIMESTAMPDIFF(YEAR, dob, CURDATE()) AS age FROM users ORDER BY age ASC")
	// res, err := database.DB.Query("SELECT user, dob FROM users")

	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}

	//close the result set
	defer res.Close()
	//create a slice of users
	var users []models.UserWithAge
	//loop through the result set
	for res.Next() {

		user := models.UserWithAge{}

		err := res.Scan(&user.Id, &user.User, &user.Passhash, &user.Email, &user.First_name, &user.Last_name, &user.Age)

		if err != nil {
			log.Fatal(err)
		}
		//best logger amirite
		fmt.Printf("%v\n", user)
		//storing the data in a slice
		users = append(users, user)
	}
	//testing the new func
	utils.AtLeastTwice(users)
	utils.AtLEastTwiceAlt(users)
	utils.ExactlyTwice(users)

	//get users within limits
	usersWithinLimits := utils.GetUsersWithinLimits(users, 18, 80)
	fmt.Printf("%v\n", usersWithinLimits)

	//binary search
	utils.ConstrainedExactlyTwice(usersWithinLimits)

	//----end tests
	utils.ResponseHelperJSON(c, users, "users", "No users found")

	return err
}

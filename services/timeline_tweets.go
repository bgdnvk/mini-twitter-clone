package services

import (
	"goexample/database"
	"goexample/models"
	"log"

	// "database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)



func GetTimelineTweets(c *fiber.Ctx) error {

	//Message: "You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near '%!d(string=)' at line 1"
	dbQuery := fmt.Sprintf("SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = %d;", c.Params("id"))
	//WHERE followers.id_follower = %!d(string=1)
	fmt.Print(dbQuery)
	res, err := database.DB.Query(dbQuery)

	//check for errors
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return nil
	}
	//close the result set
	defer res.Close()

	//create a slice of tweets
	var timelineTweets []models.TimelineTweet
	//loop through the result set
	for res.Next() {
		timelineTweet := models.TimelineTweet{}
		err := res.Scan(&timelineTweet.User_id, &timelineTweet.User, &timelineTweet.First_name, &timelineTweet.Last_name, &timelineTweet.Tweet, &timelineTweet.Date_tweet)
		if err != nil {
			log.Fatal(err)
		}
		timelineTweets = append(timelineTweets, timelineTweet)
	}
	fmt.Print(timelineTweets)

	return nil
}
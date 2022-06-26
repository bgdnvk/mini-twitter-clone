package services

import (
	"goexample/database"
	"goexample/models"
	"goexample/services/utils"
	"log"

	// "database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetTimelineTweets(c *fiber.Ctx) error {
	//TODO: pagination w/ limit n offset
	dbQuery := fmt.Sprintf("SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = %s;", c.Params("id"))
	// dbQuery := fmt.Sprintf("SELECT users.user_od, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = %s;", c.Params("id"))
	rows, err := database.DB.Query(dbQuery)

	//check for errors
	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}
	//close db connection
	defer rows.Close()

	//create a slice of tweets
	var timelineTweets []models.TimelineTweet
	//loop through the result set
	for rows.Next() {
		timelineTweet := models.TimelineTweet{}
		err := rows.Scan(&timelineTweet.User_id, &timelineTweet.User, &timelineTweet.First_name, &timelineTweet.Last_name, &timelineTweet.Tweet, &timelineTweet.Date_tweet)
		if err != nil {
			log.Fatal(err)
		}
		timelineTweets = append(timelineTweets, timelineTweet)
	}
	fmt.Print(timelineTweets)

	// if timelineTweets != nil {
	// 	c.Status(200).JSON(&fiber.Map{
	// 		"success": true,
	// 		"tweets":  timelineTweets,
	// 	})
	// } else {
	// 	c.Status(404).JSON(&fiber.Map{
	// 		"success": false,
	// 		"error":   "No timeline found",
	// 	})
	// }
	utils.ResponseHelperJSON(c, timelineTweets, "timeline", "No timeline found")

	return err
}

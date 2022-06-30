package services

import (
	"fmt"
	"goexample/database"
	"goexample/models"
	"goexample/services/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetFeedTweets(c *fiber.Ctx) error {

	//you shouldn't do this by the way, but it's just a demo
	// dbQuery := fmt.Sprintf("SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = %s ORDER BY tweets.date_tweet DESC;", c.Params("id"))
	// rows, err := database.DB.Query(dbQuery)

	//avoid the SQL injection by rewriting it like
	dbQuery := "SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = ? ORDER BY tweets.date_tweet DESC;"
	rows, err := database.DB.Query(dbQuery, c.Params("id"))

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

	utils.ResponseHelperJSON(c, timelineTweets, "timeline", "No timeline found")

	return err
}

func GetFeedTweetsPaginated(c *fiber.Ctx) error {

	dbQuery := fmt.Sprintf("SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = %s ORDER BY tweets.date_tweet DESC LIMIT %s OFFSET %s;", c.Params("id"), c.Params("limit"), c.Params("offset"))

	rows, err := database.DB.Query(dbQuery)
	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}

	defer rows.Close()

	var timelineTweets []models.TimelineTweet
	for rows.Next() {
		timelineTweet := models.TimelineTweet{}
		err := rows.Scan(&timelineTweet.User_id, &timelineTweet.User, &timelineTweet.First_name, &timelineTweet.Last_name, &timelineTweet.Tweet, &timelineTweet.Date_tweet)
		if err != nil {
			log.Fatal(err)
		}
		timelineTweets = append(timelineTweets, timelineTweet)
	}
	//TODO: implement a response with pages and all that pagination jazz
	utils.ResponseHelperJSON(c, timelineTweets, "timeline", "No timeline found")

	return err
}

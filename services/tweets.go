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

func GetTweets(c *fiber.Ctx) error {

	//get the tweets from db
	res, err := database.DB.Query("SELECT * FROM tweets")
	//check for errors
	// if err != nil {
	// 	c.Status(500).JSON(&fiber.Map{
	// 		"success": false,
	// 		"error":   err,
	// 	})
	// 	return nil
	// }
	if err != nil {
		return utils.DefaultErrorHandler(c, err)
	}

	//close the result set
	defer res.Close()
	//create a slice of tweets
	var tweets []models.Tweet
	//loop through the result set
	for res.Next() {

		tweet := models.Tweet{}

		err := res.Scan(&tweet.Id, &tweet.User_id, &tweet.Tweet, &tweet.Date_tweet)

		if err != nil {
			log.Fatal(err)
		}
		//best logger amirite
		fmt.Printf("%v\n", tweet)
		//storing the data in a slice
		tweets = append(tweets, tweet)
	}

	// fmt.Printf("%v\n", tweets)

	//send the tweets to the client
	// if tweets != nil {
	// 	c.Status(200).JSON(&fiber.Map{
	// 		"success": true,
	// 		"tweets":  tweets,
	// 	})
	// } else {
	// 	c.Status(404).JSON(&fiber.Map{
	// 		"success": false,
	// 		"error":   "No tweets found",
	// 	})
	// }
	utils.ResponseHelperJSON(c, tweets, "tweets", "No tweets found")

	return err
}

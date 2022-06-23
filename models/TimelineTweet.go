package models

type TimelineTweet struct {
	User_id int `json:"user_id"`
	User string `json:"user"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Tweet string `json:"tweet"`
	Date_tweet string `json:"date_tweet"`
}

type TimelineTweets struct {
	Timeline_tweets []TimelineTweet `json:"timeline_tweets"`
}
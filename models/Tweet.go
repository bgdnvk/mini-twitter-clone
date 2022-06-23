package models

type Tweet struct {
	Id        int    `json:"id"`
	User_id   int    `json:"user_id"`
	Tweet     string `json:"tweet"`
	Date_tweet string `json:"date_tweet"`
}

type Tweets struct {
	Tweets []Tweet `json:"tweets"`
}
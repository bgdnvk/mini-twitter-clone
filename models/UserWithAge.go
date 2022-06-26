package models

type UserWithAge struct{
	Id         int   `json:"id"`
	User       string `json:"user"`
	Passhash   string `json:"passhash"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Age        int    `json:"age"`
}
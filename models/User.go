package models

type User struct {
	Id         int   `json:"id"`
	User       string `json:"user"`
	Passhash   string `json:"passhash"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Dob        string `json:"dob"`
}

type Users struct {
	Users []User `json:"users"`
}

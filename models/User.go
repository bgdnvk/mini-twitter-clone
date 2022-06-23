package models

type User struct {
	Id         int
	User       string
	Passhash   string
	Email      string
	First_name string
	Last_name  string
	Dob        string
}

// func NewUser() *user {
// 	return &user{}
// }

func NewUser() *User {
	return &User{}
}

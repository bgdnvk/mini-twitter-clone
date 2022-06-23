package models

type Followers struct {
	Id_user int `json:"id_user"`
	Id_follower int `json:"id_follower"`
}

type Followerss struct {
	Followers []Followers `json:"followers"`
}
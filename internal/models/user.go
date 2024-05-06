package models

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsPrivate bool   `json:"is_private"`
}

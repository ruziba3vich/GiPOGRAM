package models

type Post struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Content int    `json:"content"`
	Likes   int    `json:"likes"`
}

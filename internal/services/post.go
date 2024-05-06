package services

import (
	"github.com/ruziba3vich/GiPOGRAM/internal/models"
	"database/sql"
)

type PostRequest struct {
	UserId int         `json:"user_id"`
	Post   models.Post `json:"post"`
}

func CreatePost(postReq PostRequest, db *sql.DB) (*models.Post, error) {
	query := `
		INSERT INTO Posts(
			user_id,
			title,
			content,
			likes
		)
		VALUES (
			$1, $2, $3, $4
		)
		RETURNING 
			id,
			user_id,
			title,
			content,
			likes;
	`

	row, err := db.Query(
		query,
		postReq.UserId,
		postReq.Post.Title,
		postReq.Post.Content,
		postReq.Post.Likes)
	if err != nil {
		return nil, err
	}
	var newPost models.Post
	if err := row.Scan(
		&newPost.Id,
		&newPost.UserId,
		&newPost.Title,
		&newPost.Content,
		&newPost.Likes); err != nil {
		return nil, err
	}

	return &newPost, nil
}

/*
type Post struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Content int    `json:"content"`
	Likes   int    `json:"likes"`
}
*/

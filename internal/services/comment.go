package services

import (
	"github.com/ruziba3vich/GiPOGRAM/internal/models"
	"database/sql"
)

/*
type Comment struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	PostId  int    `json:"post_id"`
	Content string `json:"content"`
}
*/

type CommentRequest struct {
	PostId  int            `json:"post_id"`
	UserId  int            `json:"user_id"`
	Comment models.Comment `json:"comment"`
}

func MakeComment(comtReq CommentRequest, db *sql.DB) (*models.Comment, error) {
	query := `
		INSERT INTO Comments(
			user_id,
			post_id,
			content
		)
		VALUES (
			$1, $2, $3
		)
		RETURNING id;
	`
	rows, err := db.Query(
		query,
		comtReq.PostId,
		comtReq.UserId,
		comtReq.Comment)
	if err != nil {
		return nil, err
	}
	var newComment models.Comment
	if err := rows.Scan(
		&newComment.UserId,
		&newComment.UserId,
		&newComment.PostId,
		&newComment.Content); err != nil {
		return nil, err
	}
	return &newComment, nil
}

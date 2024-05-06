package services

import (
	"database/sql"

	"github.com/ruziba3vich/GiPOGRAM/internal/models"
)

/*
type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsPrivate bool   `json:"is_private"`
}
*/

func CreateUser(user models.User, db *sql.DB) (*models.User, error) {
	query := `
		INSERT INTO Users(
			username,
			password,
			is_private
		)
		VALUES (
			$1, $2, $3
		)
		RETURNING id, username, password, is_private;
	`
	var newUser models.User
	row, err := db.Query(
		query,
		user.Username,
		user.Password,
		user.IsPrivate)
	if err != nil {
		return nil, err
	}
	if err := row.Scan(
		&newUser.Id,
		&newUser.Username,
		&newUser.Password,
		&newUser.IsPrivate); err != nil {
		return nil, err
	}
	return &newUser, nil
}

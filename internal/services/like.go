package services

import "database/sql"

func MakeLike(userId, postId int, db *sql.DB) (bool, error) {
	query := `
	SELECT EXISTS (
		SELECT 1
		FROM likes
		WHERE user_id = $1 AND post_id = $2;
	);
	`
	rows, err := db.Query(query, userId, postId)
	if err != nil {
		return false, err
	}
	var exists bool
	if err := rows.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

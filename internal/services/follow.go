package services

import (
	"database/sql"
)

// FollowRequests

// Connections

type UserDTO struct {
	Username string `json:"username"`
}

func SendFollowRequest(senderId, reciepantId int, db *sql.DB) (bool, error) {
	getReciepantQuery := `
		SELECT is_private FROM Users
		WHERE user_id = $1;
	`
	var isPrivate bool
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	row, err := tx.Query(getReciepantQuery, reciepantId)
	if err != nil {
		return false, err
	}
	if err := row.Scan(&isPrivate); err != nil {
		return false, err
	}
	if isPrivate {
		query := `INSERT INTO FollowRequests(
					reciepant_id,
					sender_id
				) VALUES (
					$1, $2
				);`
		_, err := tx.Query(query, reciepantId, senderId)
		if err != nil {
			return false, err
		}
		tx.Commit()
		return true, nil
	} else {
		check := `
		SELECT EXISTS (
			SELECT 1
			FROM Connections
			WHERE getter = $1 AND sender = $2;
		);
		`
		row, err := tx.Query(check, reciepantId, senderId)
		if err != nil {
			return false, nil
		}
		var follows bool
		if err := row.Scan(&follows); err != nil {
			return false, err
		}
		if !follows {
			query := `
			INSERT INTO Connections (
				getter,
				sender
			)
			VALUES (
				$1, $2
			);
		`
			_, err := tx.Query(query, reciepantId, senderId)
			if err != nil {
				return false, err
			}
			tx.Commit()
			return true, nil
		} else {
			query := `
			DELETE FROM Connections
			WHERE getter = $1 AND sender = $2;
		`
			_, err := tx.Query(query, reciepantId, senderId)
			if err != nil {
				return false, err
			}
			tx.Commit()
			return false, nil
		}
	}
}

func GetAllRequests(userId int, db *sql.DB) ([]UserDTO, error) {
	query := `
		SELECT username FROM FollowRequests
		WHERE reciepant_id = $1;
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	var users []UserDTO
	for rows.Next() {
		var userDto UserDTO
		err := rows.Scan(&userDto.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, userDto)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// func AcceptFollowRequest(reciepantId, senderId int, db *sql.DB) (bool, error) {

// }

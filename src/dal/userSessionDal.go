package dal

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"house-manager-backend/models"
	"time"
)

func CreateUserSession(userId int) error {

	connection := InitDB()

	query := "INSERT INTO user_session (user_id,expires_at) VALUES (?,?)"
	expireTime := time.Now().UTC().Add(15 * time.Minute)

	_, err := connection.Exec(query, userId, expireTime)
	if err != nil {
		return err
	}
	return nil

}
func GetUserSession(userid int) (models.UserSession, error) {

	connection := InitDB()

	var userSession models.UserSession

	query := "SELECT session_id,user_id,created_at,expires_at FROM user_session WHERE user_id = ? ORDER BY created_at DESC LIMIT 1"
	value := connection.QueryRow(query, userid).Scan(
		&userSession.SessionId,
		&userSession.UserId,
		&userSession.CreatedAt,
		&userSession.ExpiresAt)

	if value == sql.ErrNoRows {
		return userSession, errors.New("user session not found")
	}

	return userSession, nil

}

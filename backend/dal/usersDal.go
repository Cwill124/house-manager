package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"house-manager-backend/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string, email string) error {

	connection := InitDB()

	saltPassword, _ := SaltPassword(password)
	query := "INSERT INTO users (username,password,email) VALUES (?,?,?)"
	_, err := connection.Exec(query, username, saltPassword, email)
	if err != nil {
		return err
	}
	fmt.Println("User Created successfully")
	return nil
}

func UserLogin(username string, password string) (models.User, error) {

	connection := InitDB()

	var user models.User

	query := "SELECT username, password, email FROM users WHERE username = ?"

	value := connection.QueryRow(query, strings.TrimSpace(username)).Scan(
		&user.Username,
		&user.Password,
		&user.Email)

	if value == sql.ErrNoRows {
		return user, errors.New("user not found, username or password is incorrect")
	}

	comparePass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if comparePass != nil {
		return user, errors.New("password is incorrect")
	}
	fmt.Printf("User found username %s", user.Username)

	return user, nil

}

func SaltPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

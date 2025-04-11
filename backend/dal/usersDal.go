package dal

import (
	"fmt"
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



func SaltPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

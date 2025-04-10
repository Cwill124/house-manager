package dal

import (
  "fmt"
)

func CreateUser(username string, password string, email string) error {

	connection := InitDB()

	query := "INSERT INTO users (username,password,email) VALUES (?,?,?)"
	_, err := connection.Exec(query, username, password, email)
	if err != nil {
		return err
	}
	fmt.Println("User Created successfully")
	return nil
}

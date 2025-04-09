package main

import (
	"fmt"
	"house-manager-backend/handlers/auth"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.Login) // Set up a handler for the root URL
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // Start the server on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

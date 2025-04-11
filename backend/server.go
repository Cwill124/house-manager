package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	auth "house-manager-backend/restController/auth"
	"net/http"
)

func main() {
	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	http.HandleFunc("/login", auth.GenerateJWT)
	http.HandleFunc("/refresh", auth.RefreshToken)
	http.HandleFunc("/createUser", auth.CreateUser)

	fmt.Println("Server is running at http://localhost:8000")

	err := http.ListenAndServe(":8000", handlers.CORS(corsOptions, corsHeaders, corsMethods)(http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

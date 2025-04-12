package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	auth "house-manager-backend/restController/auth"
	authView "house-manager-backend/viewController/auth"
	"net/http"
)

func initRestApi() {

	http.HandleFunc("/api/login", auth.GenerateJWT)
	http.HandleFunc("/api/refresh", auth.RefreshToken)
	http.HandleFunc("/api/createUser", auth.CreateUser)

}
func initViews() {
	http.HandleFunc("/", authView.LoginViewHandler)
	http.HandleFunc("/sign-up", authView.SignUpViewHandler)

}

func main() {
	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	initRestApi()
	initViews()
fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server is running at http://localhost:8000")

	err := http.ListenAndServe(":8000", handlers.CORS(corsOptions, corsHeaders, corsMethods)(http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

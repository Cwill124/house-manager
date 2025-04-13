package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	auth "house-manager-backend/restController/auth"
	authView "house-manager-backend/viewController/auth"
	"net/http"
)

func initRestApi(r *mux.Router) {
	r.HandleFunc("/api/login", auth.UserLogin).Methods("POST")
	r.HandleFunc("/api/createUser", auth.CreateUser).Methods("POST")
}

func initViews(r *mux.Router) {
	r.HandleFunc("/", authView.LoginViewHandler).Methods("GET")
	r.HandleFunc("/sign-up", authView.SignUpViewHandler).Methods("GET")
}

func main() {
	r := mux.NewRouter()
	initRestApi(r)
	initViews(r)
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	fmt.Println("Server is running at http://localhost:8000")
	err := http.ListenAndServe(":8000", handlers.CORS(corsOptions, corsHeaders, corsMethods)(r))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

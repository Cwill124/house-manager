package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
} 

func Login(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "This is a successful response",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Optional: Status OK (200)
	json.NewEncoder(w).Encode(response)
}

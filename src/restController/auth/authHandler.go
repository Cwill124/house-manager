package handlers

import (
	"encoding/json"
	"fmt"
	"house-manager-backend/dal"
	"net/http"
)

type reDirectResponse struct {
	RedirectTo string `json:"redirectTo,omitempty"`
	Error      string `json:"error,omitempty"`
}
type loginResponse struct {
	RedirectTo string `json:"redirectTo,omitempty"`
	UserId     int    `json:"userId"`
	Error      string `json:"error,omitempty"`
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.Username == "" {
		http.Error(w, "Missing or invalid Username", http.StatusBadRequest)
		return
	} else if requestBody.Password == "" {
		http.Error(w, "Missing or invalid Password", http.StatusBadRequest)
		return
	}

	user, loginError := dal.UserLogin(requestBody.Username, requestBody.Password)
	if loginError != nil {
		http.Error(w, "Username or password is incorrect", http.StatusBadRequest)
		return
	}

	createUserSessionError := dal.CreateUserSession(user.Id)

	if createUserSessionError != nil {
		http.Error(w, "failed to create user session", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	houseSelectionUrl := "/house-selection-dashboard"
	fmt.Println(houseSelectionUrl)
	resp := loginResponse{RedirectTo: houseSelectionUrl, UserId: user.Id}
	json.NewEncoder(w).Encode(resp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.Username == "" {
		http.Error(w, "Missing or invalid Username", http.StatusBadRequest)
		return
	} else if requestBody.Password == "" {
		http.Error(w, "Missing or invalid Password", http.StatusBadRequest)
		return
	} else if requestBody.Email == "" {
		http.Error(w, "Missing or invalid Email", http.StatusBadRequest)
		return
	}
	createUser := dal.CreateUser(requestBody.Username, requestBody.Password, requestBody.Email)

	if createUser != nil {
		fmt.Println(createUser.Error())
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Restricted\"")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error ")
		return
	}
	resp := reDirectResponse{RedirectTo: "/"}
	json.NewEncoder(w).Encode(resp)
}

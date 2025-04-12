package handlers

import (
	"encoding/json"
	"fmt"
	"house-manager-backend/dal"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret keys for signing the JWT (use a secure method for production)
var jwtSecret = []byte("my_secret_key")
var refreshSecret = []byte("my_refresh_secret_key")

// Response structure for returning the JWT and Refresh token
type Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type SignupResponse struct {
	RedirectTo string `json:"redirectTo,omitempty"`
	Error      string `json:"error,omitempty"`
}

func GenerateJWT(w http.ResponseWriter, r *http.Request) {

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
	username := user.Username

	accessClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Access token expires in 15 minutes
	}

	refreshClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Could not create access token", http.StatusInternalServerError)
		return
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		http.Error(w, "Could not create refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})
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
	resp := SignupResponse{RedirectTo: "/"}
	json.NewEncoder(w).Encode(resp)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Get the Refresh Token from the request body
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.RefreshToken == "" {
		http.Error(w, "Missing or invalid refresh token", http.StatusBadRequest)
		return
	}

	// Now requestBody.RefreshToken contains the refresh token
	refreshToken := requestBody.RefreshToken

	// Parse and validate the refresh token (same as before)
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return refreshSecret, nil
	})

	// If the token is invalid or expired
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	// Generate a new Access Token
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	// Create a new Access Token
	accessClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Access token expires in 15 minutes
	}

	// Create new Access Token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Could not create new access token", http.StatusInternalServerError)
		return
	}

	// Send the new access token back in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		AccessToken: accessTokenString,
	})
}

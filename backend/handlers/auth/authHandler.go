package handlers

import (
	"encoding/json"
	"fmt"
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

// GenerateJWT generates both the Access Token and Refresh Token
func GenerateJWT(w http.ResponseWriter, r *http.Request) {
	// Simulate login and generate a username (in real apps, authenticate the user)
	username := "user1"

	// Claims for Access Token (short expiry time)
	accessClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Access token expires in 15 minutes
	}

	// Claims for Refresh Token (long expiry time)
	refreshClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
	}

	// Create Access Token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Could not create access token", http.StatusInternalServerError)
		return
	}

	// Create Refresh Token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		http.Error(w, "Could not create refresh token", http.StatusInternalServerError)
		return
	}

	// Send both tokens in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})
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

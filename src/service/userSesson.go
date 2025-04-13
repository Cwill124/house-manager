package service

import (
	"house-manager-backend/dal"
	"time"
)

func CheckUserSession(userId int) bool {

	userSession, userSessionError := dal.GetUserSession(userId)

	if userSessionError != nil {
		return false
	}

	if time.Now().UTC().After(userSession.ExpiresAt.Add(3 * time.Minute)){
		return false
	} else {
		dal.CreateUserSession(userId)
		return true
	}
}

package models

import (
	"time"
)

type UserSession struct {
	SessionId int
	UserId    int
	CreatedAt time.Time
	ExpiresAt time.Time
}

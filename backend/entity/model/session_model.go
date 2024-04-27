package model

import (
	"github.com/google/uuid"
	"time"
)

const (
	SessionTokenDuration         = time.Hour * 24 * 7     // 1 week
	SessionTokenAbsoluteDuration = time.Hour * 24 * 7 * 4 // 4 weeks
)

type Session struct {
	ID         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	CSRF       uuid.UUID  `json:"csrf"`
	Token      uuid.UUID  `json:"token"`
	Expiration time.Time  `json:"expiration"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type SessionU struct {
	Expiration *time.Time
}

type SessionF struct {
	ID         *uuid.UUID
	UserID     *uuid.UUID
	CSRF       *uuid.UUID
	Token      *uuid.UUID
	Expiration *time.Time
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

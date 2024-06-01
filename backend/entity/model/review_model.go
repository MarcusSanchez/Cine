package model

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	MediaID   uuid.UUID  `json:"media_id"`
	Content   string     `json:"content"`
	Rating    int        `json:"rating"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ReviewU struct {
	Content *string
	Rating  *int
}

type ReviewF struct {
	ID        *uuid.UUID
	UserID    *uuid.UUID
	MediaID   *uuid.UUID
	Content   *string
	Rating    *int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type DetailedReview struct {
	Review *Review `json:"review"`
	User   *User   `json:"user"`
}

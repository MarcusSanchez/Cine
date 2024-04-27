package model

import (
	"github.com/google/uuid"
	"time"
)

type Like struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	CommentID uuid.UUID  `json:"comment_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LikeU struct{}

type LikeF struct {
	ID        *uuid.UUID
	UserID    *uuid.UUID
	CommentID *uuid.UUID
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID           uuid.UUID  `json:"id"`
	UserID       *uuid.UUID `json:"user_id"`
	MediaID      uuid.UUID  `json:"media_id"`
	ReplyingToID *uuid.UUID `json:"replying_to_id"`
	Content      string     `json:"content"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type CommentU struct {
	Content *string
}

type CommentF struct {
	ID           *uuid.UUID
	UserID       *uuid.UUID
	MediaID      *uuid.UUID
	ReplyingToID *uuid.UUID
	Content      *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

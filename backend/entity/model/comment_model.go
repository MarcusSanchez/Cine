package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
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

type DetailedComment struct {
	Comment      *Comment `json:"comment"`
	User         *User    `json:"user"`
	RepliesCount int      `json:"replies_count"`
	LikesCount   int      `json:"likes_count"`
	LikedByUser  bool     `json:"liked_by_user"`
}

package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID  `json:"id"`
	DisplayName    string     `json:"display_name"`
	Username       string     `json:"username"`
	Email          string     `json:"-"`
	Password       string     `json:"-"`
	ProfilePicture string     `json:"profile_picture"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type UserU struct {
	DisplayName    *string
	Username       *string
	Email          *string
	Password       *string
	ProfilePicture *string
}

type UserF struct {
	ID             *uuid.UUID
	DisplayName    *string
	Username       *string
	UsernameNotIn  *[]string
	Email          *string
	Password       *string
	ProfilePicture *string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type DetailedUser struct {
	User           *User `json:"user"`
	FollowingCount int   `json:"following_count"`
	FollowersCount int   `json:"followers_count"`
	LikesCount     int   `json:"likes_count"`
	CommentsCount  int   `json:"comments_count"`
	ReviewsCount   int   `json:"reviews_count"`
	ListsCount     int   `json:"lists_count"`
	Followed       bool  `json:"followed"`
}

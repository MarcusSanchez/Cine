package model

import (
	"github.com/google/uuid"
	"time"
)

type List struct {
	ID        uuid.UUID  `json:"id"`
	OwnerID   uuid.UUID  `json:"owner_id"`
	Title     string     `json:"name"`
	Public    bool       `json:"is_public"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ListU struct {
	Title  *string
	Public *bool
}

type ListF struct {
	ID        *uuid.UUID
	OwnerID   *uuid.UUID
	Title     *string
	Public    *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time

	HasMember *uuid.UUID
	HasMedia  *uuid.UUID
}

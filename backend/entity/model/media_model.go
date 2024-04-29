package model

import (
	"github.com/google/uuid"
	"time"
)

type MediaType string

const (
	MediaTypeMovie MediaType = "movie"
	MediaTypeShow  MediaType = "show"
)

type Media struct {
	ID           uuid.UUID  `json:"id"`
	Ref          int        `json:"ref"`
	MediaType    MediaType  `json:"media_type"`
	Overview     string     `json:"overview"`
	BackdropPath *string    `json:"backdrop_path"`
	Language     string     `json:"language"`
	PosterPath   *string    `json:"poster_path"`
	ReleaseDate  *string    `json:"release_date"`
	Title        string     `json:"title"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type MediaU struct{}

type MediaF struct {
	ID           *uuid.UUID
	Ref          *int
	MediaType    *MediaType
	Overview     *string
	BackdropPath *string
	Language     *string
	PosterPath   *string
	ReleaseDate  *string
	Title        *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

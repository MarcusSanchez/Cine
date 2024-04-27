package datastore

import (
	"cine/repository"
	"context"
)

type Store interface {
	Users() repository.UserRepository
	Sessions() repository.SessionRepository
	Comments() repository.CommentRepository
	Likes() repository.LikeRepository
	Reviews() repository.ReviewRepository
	Medias() repository.MediaRepository
	Lists() repository.ListRepository

	Transaction(ctx context.Context) (Transaction, error)
}

type Transaction interface {
	Users() repository.UserRepository
	Sessions() repository.SessionRepository
	Comments() repository.CommentRepository
	Likes() repository.LikeRepository
	Reviews() repository.ReviewRepository
	Medias() repository.MediaRepository
	Lists() repository.ListRepository

	Commit() error
	Rollback() error
}

package mocks

import (
	"cine/datastore"
	"cine/repository"
	"context"
)

type Store struct {
	User    *UserRepository
	Session *SessionRepository
	Comment *CommentRepository
	Like    *LikeRepository
	Review  *ReviewRepository
	Media   *MediaRepository
	List    *ListRepository
}

var _ datastore.Store = (*Store)(nil)

func NewStore() *Store {
	return &Store{
		User:    NewUserRepository(),
		Session: NewSessionRepository(),
		Comment: NewCommentRepository(),
		Like:    NewLikeRepository(),
		Review:  NewReviewRepository(),
		Media:   NewMediaRepository(),
		List:    NewListRepository(),
	}
}

func (s Store) Users() repository.UserRepository       { return s.User }
func (s Store) Sessions() repository.SessionRepository { return s.Session }
func (s Store) Comments() repository.CommentRepository { return s.Comment }
func (s Store) Likes() repository.LikeRepository       { return s.Like }
func (s Store) Reviews() repository.ReviewRepository   { return s.Review }
func (s Store) Medias() repository.MediaRepository     { return s.Media }
func (s Store) Lists() repository.ListRepository       { return s.List }

type transaction struct {
	store *Store
}

func (s Store) Transaction(ctx context.Context) (datastore.Transaction, error) {
	return &transaction{store: &s}, nil
}

func (t transaction) Users() repository.UserRepository       { return t.store.User }
func (t transaction) Sessions() repository.SessionRepository { return t.store.Session }
func (t transaction) Comments() repository.CommentRepository { return t.store.Comment }
func (t transaction) Likes() repository.LikeRepository       { return t.store.Like }
func (t transaction) Reviews() repository.ReviewRepository   { return t.store.Review }
func (t transaction) Medias() repository.MediaRepository     { return t.store.Media }
func (t transaction) Lists() repository.ListRepository       { return t.store.List }
func (t transaction) Commit() error                          { return nil }
func (t transaction) Rollback() error                        { return nil }

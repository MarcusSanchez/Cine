package ent

import (
	"cine/datastore"
	"cine/datastore/ent/ent"
	"cine/repository"
	"context"
	"database/sql"
	"errors"
)

type transaction struct {
	tx          *ent.Tx
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	commentRepo repository.CommentRepository
	likeRepo    repository.LikeRepository
	reviewRepo  repository.ReviewRepository
	mediaRepo   repository.MediaRepository
	listRepo    repository.ListRepository
}

func (s *store) Transaction(ctx context.Context) (datastore.Transaction, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, new(transaction).txError(err)
	}
	client := tx.Client()
	return &transaction{
		tx:          tx,
		userRepo:    newUserRepository(client),
		sessionRepo: newSessionRepository(client),
		commentRepo: newCommentRepository(client),
		likeRepo:    newLikeRepository(client),
		reviewRepo:  newReviewRepository(client),
		mediaRepo:   newMediaRepository(client),
		listRepo:    newListRepository(client),
	}, nil
}

func (t *transaction) Users() repository.UserRepository       { return t.userRepo }
func (t *transaction) Sessions() repository.SessionRepository { return t.sessionRepo }
func (t *transaction) Comments() repository.CommentRepository { return t.commentRepo }
func (t *transaction) Likes() repository.LikeRepository       { return t.likeRepo }
func (t *transaction) Reviews() repository.ReviewRepository   { return t.reviewRepo }
func (t *transaction) Medias() repository.MediaRepository     { return t.mediaRepo }
func (t *transaction) Lists() repository.ListRepository       { return t.listRepo }

func (t *transaction) Commit() error {
	err := t.tx.Commit()
	return t.txError(err)
}

func (t *transaction) Rollback() error {
	err := t.tx.Rollback()
	return t.txError(err)
}

func (*transaction) txError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, sql.ErrTxDone):
		return datastore.Wrap(datastore.ErrTxDone, "rollback or commit already called")
	case errors.Is(err, sql.ErrConnDone):
		return datastore.Wrap(datastore.ErrInternal, "connection closed")
	case errors.Is(err, ent.ErrTxStarted):
		return datastore.Wrap(datastore.ErrTxNested, "attempted to start a transaction from within a transaction")
	default:
		return datastore.Wrap(datastore.ErrInternal, err.Error())
	}
}

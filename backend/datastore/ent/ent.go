package ent

import (
	"cine/config"
	"cine/datastore"
	"cine/datastore/ent/ent"
	"cine/pkg/logger"
	"cine/repository"
	"context"
	"entgo.io/ent/dialect"
	"go.uber.org/fx"
)

type store struct {
	client      *ent.Client
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	commentRepo repository.CommentRepository
	likeRepo    repository.LikeRepository
	reviewRepo  repository.ReviewRepository
	mediaRepo   repository.MediaRepository
	listRepo    repository.ListRepository
}

func NewStore(
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	config *config.Config,
	logger logger.Logger,
) datastore.Store {
	client, err := ent.Open(dialect.Postgres, config.Datasource)
	if err != nil {
		logger.Error("failed connecting to postgresql", err)
		_ = shutdowner.Shutdown()
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		logger.Error("failed creating schema resources", err)
		_ = shutdowner.Shutdown()
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return client.Close()
		},
	})

	return &store{
		client:      client,
		userRepo:    newUserRepository(client),
		sessionRepo: newSessionRepository(client),
		commentRepo: newCommentRepository(client),
		likeRepo:    newLikeRepository(client),
		reviewRepo:  newReviewRepository(client),
		mediaRepo:   newMediaRepository(client),
		listRepo:    newListRepository(client),
	}
}

func (s *store) Users() repository.UserRepository       { return s.userRepo }
func (s *store) Sessions() repository.SessionRepository { return s.sessionRepo }
func (s *store) Comments() repository.CommentRepository { return s.commentRepo }
func (s *store) Likes() repository.LikeRepository       { return s.likeRepo }
func (s *store) Reviews() repository.ReviewRepository   { return s.reviewRepo }
func (s *store) Medias() repository.MediaRepository     { return s.mediaRepo }
func (s *store) Lists() repository.ListRepository       { return s.listRepo }

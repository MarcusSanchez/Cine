package mocks

import (
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
)

var _ repository.LikeRepository = (*LikeRepository)(nil)

type LikeRepository struct {
	OneFn        func(ctx context.Context, filters ...*model.LikeF) (*model.Like, error)
	AllFn        func(ctx context.Context, filters ...*model.LikeF) ([]*model.Like, error)
	ExistsFn     func(ctx context.Context, filters ...*model.LikeF) (bool, error)
	CountFn      func(ctx context.Context, filters ...*model.LikeF) (int, error)
	InsertFn     func(ctx context.Context, entity *model.Like) (*model.Like, error)
	InsertBulkFn func(ctx context.Context, entities []*model.Like) ([]*model.Like, error)
	UpdateFn     func(ctx context.Context, id uuid.UUID, updater *model.LikeU) (*model.Like, error)
	UpdateExecFn func(ctx context.Context, updater *model.LikeU, filters ...*model.LikeF) (int, error)
	DeleteFn     func(ctx context.Context, id uuid.UUID) error
	DeleteExecFn func(ctx context.Context, filters ...*model.LikeF) (int, error)
}

func NewLikeRepository() *LikeRepository {
	return &LikeRepository{}
}

func (l *LikeRepository) One(ctx context.Context, filters ...*model.LikeF) (*model.Like, error) {
	if l.OneFn != nil {
		return l.OneFn(ctx, filters...)
	}
	return &model.Like{}, nil
}

func (l *LikeRepository) All(ctx context.Context, filters ...*model.LikeF) ([]*model.Like, error) {
	if l.AllFn != nil {
		return l.AllFn(ctx, filters...)
	}
	return []*model.Like{}, nil
}

func (l *LikeRepository) Exists(ctx context.Context, filters ...*model.LikeF) (bool, error) {
	if l.ExistsFn != nil {
		return l.ExistsFn(ctx, filters...)
	}
	return false, nil
}

func (l *LikeRepository) Count(ctx context.Context, filters ...*model.LikeF) (int, error) {
	if l.CountFn != nil {
		return l.CountFn(ctx, filters...)
	}
	return 0, nil
}

func (l *LikeRepository) Insert(ctx context.Context, entity *model.Like) (*model.Like, error) {
	if l.InsertFn != nil {
		return l.InsertFn(ctx, entity)
	}
	return &model.Like{}, nil
}

func (l *LikeRepository) InsertBulk(ctx context.Context, entities []*model.Like) ([]*model.Like, error) {
	if l.InsertBulkFn != nil {
		return l.InsertBulkFn(ctx, entities)
	}
	return []*model.Like{}, nil
}

func (l *LikeRepository) Update(ctx context.Context, id uuid.UUID, updater *model.LikeU) (*model.Like, error) {
	if l.UpdateFn != nil {
		return l.UpdateFn(ctx, id, updater)
	}
	return &model.Like{}, nil
}

func (l *LikeRepository) UpdateExec(ctx context.Context, updater *model.LikeU, filters ...*model.LikeF) (int, error) {
	if l.UpdateExecFn != nil {
		return l.UpdateExecFn(ctx, updater, filters...)
	}
	return 0, nil
}

func (l *LikeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if l.DeleteFn != nil {
		return l.DeleteFn(ctx, id)
	}
	return nil
}

func (l *LikeRepository) DeleteExec(ctx context.Context, filters ...*model.LikeF) (int, error) {
	if l.DeleteExecFn != nil {
		return l.DeleteExecFn(ctx, filters...)
	}
	return 0, nil
}

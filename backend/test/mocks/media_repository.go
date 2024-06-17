package mocks

import (
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
)

var _ repository.MediaRepository = (*MediaRepository)(nil)

type MediaRepository struct {
	OneFn        func(ctx context.Context, filters ...*model.MediaF) (*model.Media, error)
	AllFn        func(ctx context.Context, filters ...*model.MediaF) ([]*model.Media, error)
	ExistsFn     func(ctx context.Context, filters ...*model.MediaF) (bool, error)
	CountFn      func(ctx context.Context, filters ...*model.MediaF) (int, error)
	InsertFn     func(ctx context.Context, entity *model.Media) (*model.Media, error)
	InsertBulkFn func(ctx context.Context, entities []*model.Media) ([]*model.Media, error)
	UpdateFn     func(ctx context.Context, id uuid.UUID, updater *model.MediaU) (*model.Media, error)
	UpdateExecFn func(ctx context.Context, updater *model.MediaU, filters ...*model.MediaF) (int, error)
	DeleteFn     func(ctx context.Context, id uuid.UUID) error
	DeleteExecFn func(ctx context.Context, filters ...*model.MediaF) (int, error)
}

func NewMediaRepository() *MediaRepository {
	return &MediaRepository{}
}

func (m *MediaRepository) One(ctx context.Context, filters ...*model.MediaF) (*model.Media, error) {
	if m.OneFn != nil {
		return m.OneFn(ctx, filters...)
	}
	return &model.Media{}, nil
}

func (m *MediaRepository) All(ctx context.Context, filters ...*model.MediaF) ([]*model.Media, error) {
	if m.AllFn != nil {
		return m.AllFn(ctx, filters...)
	}
	return []*model.Media{}, nil
}

func (m *MediaRepository) Exists(ctx context.Context, filters ...*model.MediaF) (bool, error) {
	if m.ExistsFn != nil {
		return m.ExistsFn(ctx, filters...)
	}
	return false, nil
}

func (m *MediaRepository) Count(ctx context.Context, filters ...*model.MediaF) (int, error) {
	if m.CountFn != nil {
		return m.CountFn(ctx, filters...)
	}
	return 0, nil
}

func (m *MediaRepository) Insert(ctx context.Context, entity *model.Media) (*model.Media, error) {
	if m.InsertFn != nil {
		return m.InsertFn(ctx, entity)
	}
	return &model.Media{}, nil
}

func (m *MediaRepository) InsertBulk(ctx context.Context, entities []*model.Media) ([]*model.Media, error) {
	if m.InsertBulkFn != nil {
		return m.InsertBulkFn(ctx, entities)
	}
	return []*model.Media{}, nil
}

func (m *MediaRepository) Update(ctx context.Context, id uuid.UUID, updater *model.MediaU) (*model.Media, error) {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, id, updater)
	}
	return &model.Media{}, nil
}

func (m *MediaRepository) UpdateExec(ctx context.Context, updater *model.MediaU, filters ...*model.MediaF) (int, error) {
	if m.UpdateExecFn != nil {
		return m.UpdateExecFn(ctx, updater, filters...)
	}
	return 0, nil
}

func (m *MediaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

func (m *MediaRepository) DeleteExec(ctx context.Context, filters ...*model.MediaF) (int, error) {
	if m.DeleteExecFn != nil {
		return m.DeleteExecFn(ctx, filters...)
	}
	return 0, nil
}

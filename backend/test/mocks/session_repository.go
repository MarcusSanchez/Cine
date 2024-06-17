package mocks

import (
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
)

var _ repository.SessionRepository = (*SessionRepository)(nil)

type SessionRepository struct {
	OneFn        func(ctx context.Context, filters ...*model.SessionF) (*model.Session, error)
	AllFn        func(ctx context.Context, filters ...*model.SessionF) ([]*model.Session, error)
	ExistsFn     func(ctx context.Context, filters ...*model.SessionF) (bool, error)
	CountFn      func(ctx context.Context, filters ...*model.SessionF) (int, error)
	InsertFn     func(ctx context.Context, entity *model.Session) (*model.Session, error)
	InsertBulkFn func(ctx context.Context, entities []*model.Session) ([]*model.Session, error)
	UpdateFn     func(ctx context.Context, id uuid.UUID, updater *model.SessionU) (*model.Session, error)
	UpdateExecFn func(ctx context.Context, updater *model.SessionU, filters ...*model.SessionF) (int, error)
	DeleteFn     func(ctx context.Context, id uuid.UUID) error
	DeleteExecFn func(ctx context.Context, filters ...*model.SessionF) (int, error)
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (s *SessionRepository) One(ctx context.Context, filters ...*model.SessionF) (*model.Session, error) {
	if s.OneFn != nil {
		return s.OneFn(ctx, filters...)
	}
	return &model.Session{}, nil
}

func (s *SessionRepository) All(ctx context.Context, filters ...*model.SessionF) ([]*model.Session, error) {
	if s.AllFn != nil {
		return s.AllFn(ctx, filters...)
	}
	return []*model.Session{}, nil
}

func (s *SessionRepository) Exists(ctx context.Context, filters ...*model.SessionF) (bool, error) {
	if s.ExistsFn != nil {
		return s.ExistsFn(ctx, filters...)
	}
	return false, nil
}

func (s *SessionRepository) Count(ctx context.Context, filters ...*model.SessionF) (int, error) {
	if s.CountFn != nil {
		return s.CountFn(ctx, filters...)
	}
	return 0, nil
}

func (s *SessionRepository) Insert(ctx context.Context, entity *model.Session) (*model.Session, error) {
	if s.InsertFn != nil {
		return s.InsertFn(ctx, entity)
	}
	return &model.Session{}, nil
}

func (s *SessionRepository) InsertBulk(ctx context.Context, entities []*model.Session) ([]*model.Session, error) {
	if s.InsertBulkFn != nil {
		return s.InsertBulkFn(ctx, entities)
	}
	return []*model.Session{}, nil
}

func (s *SessionRepository) Update(ctx context.Context, id uuid.UUID, updater *model.SessionU) (*model.Session, error) {
	if s.UpdateFn != nil {
		return s.UpdateFn(ctx, id, updater)
	}
	return &model.Session{}, nil
}

func (s *SessionRepository) UpdateExec(ctx context.Context, updater *model.SessionU, filters ...*model.SessionF) (int, error) {
	if s.UpdateExecFn != nil {
		return s.UpdateExecFn(ctx, updater, filters...)
	}
	return 0, nil
}

func (s *SessionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if s.DeleteFn != nil {
		return s.DeleteFn(ctx, id)
	}
	return nil
}

func (s *SessionRepository) DeleteExec(ctx context.Context, filters ...*model.SessionF) (int, error) {
	if s.DeleteExecFn != nil {
		return s.DeleteExecFn(ctx, filters...)
	}
	return 0, nil
}

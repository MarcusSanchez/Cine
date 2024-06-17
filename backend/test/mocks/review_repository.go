package mocks

import (
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
)

var _ repository.ReviewRepository = (*ReviewRepository)(nil)

type ReviewRepository struct {
	OneFn         func(ctx context.Context, filters ...*model.ReviewF) (*model.Review, error)
	AllFn         func(ctx context.Context, filters ...*model.ReviewF) ([]*model.Review, error)
	ExistsFn      func(ctx context.Context, filters ...*model.ReviewF) (bool, error)
	CountFn       func(ctx context.Context, filters ...*model.ReviewF) (int, error)
	InsertFn      func(ctx context.Context, entity *model.Review) (*model.Review, error)
	InsertBulkFn  func(ctx context.Context, entities []*model.Review) ([]*model.Review, error)
	UpdateFn      func(ctx context.Context, id uuid.UUID, updater *model.ReviewU) (*model.Review, error)
	UpdateExecFn  func(ctx context.Context, updater *model.ReviewU, filters ...*model.ReviewF) (int, error)
	DeleteFn      func(ctx context.Context, id uuid.UUID) error
	DeleteExecFn  func(ctx context.Context, filters ...*model.ReviewF) (int, error)
	AllWithUserFn func(ctx context.Context, reviewFs ...*model.ReviewF) ([]*model.DetailedReview, error)
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{}
}

func (r *ReviewRepository) One(ctx context.Context, filters ...*model.ReviewF) (*model.Review, error) {
	if r.OneFn != nil {
		return r.OneFn(ctx, filters...)
	}
	return &model.Review{}, nil
}

func (r *ReviewRepository) All(ctx context.Context, filters ...*model.ReviewF) ([]*model.Review, error) {
	if r.AllFn != nil {
		return r.AllFn(ctx, filters...)
	}
	return []*model.Review{}, nil
}

func (r *ReviewRepository) Exists(ctx context.Context, filters ...*model.ReviewF) (bool, error) {
	if r.ExistsFn != nil {
		return r.ExistsFn(ctx, filters...)
	}
	return false, nil
}

func (r *ReviewRepository) Count(ctx context.Context, filters ...*model.ReviewF) (int, error) {
	if r.CountFn != nil {
		return r.CountFn(ctx, filters...)
	}
	return 0, nil
}

func (r *ReviewRepository) Insert(ctx context.Context, entity *model.Review) (*model.Review, error) {
	if r.InsertFn != nil {
		return r.InsertFn(ctx, entity)
	}
	return &model.Review{}, nil
}

func (r *ReviewRepository) InsertBulk(ctx context.Context, entities []*model.Review) ([]*model.Review, error) {
	if r.InsertBulkFn != nil {
		return r.InsertBulkFn(ctx, entities)
	}
	return []*model.Review{}, nil
}

func (r *ReviewRepository) Update(ctx context.Context, id uuid.UUID, updater *model.ReviewU) (*model.Review, error) {
	if r.UpdateFn != nil {
		return r.UpdateFn(ctx, id, updater)
	}
	return &model.Review{}, nil
}

func (r *ReviewRepository) UpdateExec(ctx context.Context, updater *model.ReviewU, filters ...*model.ReviewF) (int, error) {
	if r.UpdateExecFn != nil {
		return r.UpdateExecFn(ctx, updater, filters...)
	}
	return 0, nil
}

func (r *ReviewRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if r.DeleteFn != nil {
		return r.DeleteFn(ctx, id)
	}
	return nil
}

func (r *ReviewRepository) DeleteExec(ctx context.Context, filters ...*model.ReviewF) (int, error) {
	if r.DeleteExecFn != nil {
		return r.DeleteExecFn(ctx, filters...)
	}
	return 0, nil
}

func (r *ReviewRepository) AllWithUser(ctx context.Context, reviewFs ...*model.ReviewF) ([]*model.DetailedReview, error) {
	if r.AllWithUserFn != nil {
		return r.AllWithUserFn(ctx, reviewFs...)
	}
	return []*model.DetailedReview{}, nil
}

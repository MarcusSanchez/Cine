package mocks

import (
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
)

var _ repository.CommentRepository = (*CommentRepository)(nil)

type CommentRepository struct {
	OneFn                  func(ctx context.Context, filters ...*model.CommentF) (*model.Comment, error)
	AllFn                  func(ctx context.Context, filters ...*model.CommentF) ([]*model.Comment, error)
	ExistsFn               func(ctx context.Context, filters ...*model.CommentF) (bool, error)
	CountFn                func(ctx context.Context, filters ...*model.CommentF) (int, error)
	InsertFn               func(ctx context.Context, entity *model.Comment) (*model.Comment, error)
	InsertBulkFn           func(ctx context.Context, entities []*model.Comment) ([]*model.Comment, error)
	UpdateFn               func(ctx context.Context, id uuid.UUID, updater *model.CommentU) (*model.Comment, error)
	UpdateExecFn           func(ctx context.Context, updater *model.CommentU, filters ...*model.CommentF) (int, error)
	DeleteFn               func(ctx context.Context, id uuid.UUID) error
	DeleteExecFn           func(ctx context.Context, filters ...*model.CommentF) (int, error)
	AllAsDetailedFn        func(ctx context.Context, mediaID uuid.UUID, userID uuid.UUID) ([]*model.DetailedComment, error)
	AllRepliesAsDetailedFn func(ctx context.Context, comment *model.Comment, userID uuid.UUID) ([]*model.DetailedComment, error)
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{}
}

func (c *CommentRepository) One(ctx context.Context, filters ...*model.CommentF) (*model.Comment, error) {
	if c.OneFn != nil {
		return c.OneFn(ctx, filters...)
	}
	return &model.Comment{}, nil
}

func (c *CommentRepository) All(ctx context.Context, filters ...*model.CommentF) ([]*model.Comment, error) {
	if c.AllFn != nil {
		return c.AllFn(ctx, filters...)
	}
	return []*model.Comment{}, nil
}

func (c *CommentRepository) Exists(ctx context.Context, filters ...*model.CommentF) (bool, error) {
	if c.ExistsFn != nil {
		return c.ExistsFn(ctx, filters...)
	}
	return false, nil
}

func (c *CommentRepository) Count(ctx context.Context, filters ...*model.CommentF) (int, error) {
	if c.CountFn != nil {
		return c.CountFn(ctx, filters...)
	}
	return 0, nil
}

func (c *CommentRepository) Insert(ctx context.Context, entity *model.Comment) (*model.Comment, error) {
	if c.InsertFn != nil {
		return c.InsertFn(ctx, entity)
	}
	return &model.Comment{}, nil
}

func (c *CommentRepository) InsertBulk(ctx context.Context, entities []*model.Comment) ([]*model.Comment, error) {
	if c.InsertBulkFn != nil {
		return c.InsertBulkFn(ctx, entities)
	}
	return []*model.Comment{}, nil
}

func (c *CommentRepository) Update(ctx context.Context, id uuid.UUID, updater *model.CommentU) (*model.Comment, error) {
	if c.UpdateFn != nil {
		return c.UpdateFn(ctx, id, updater)
	}
	return &model.Comment{}, nil
}

func (c *CommentRepository) UpdateExec(ctx context.Context, updater *model.CommentU, filters ...*model.CommentF) (int, error) {
	if c.UpdateExecFn != nil {
		return c.UpdateExecFn(ctx, updater, filters...)
	}
	return 0, nil
}

func (c *CommentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if c.DeleteFn != nil {
		return c.DeleteFn(ctx, id)
	}
	return nil
}

func (c *CommentRepository) DeleteExec(ctx context.Context, filters ...*model.CommentF) (int, error) {
	if c.DeleteExecFn != nil {
		return c.DeleteExecFn(ctx, filters...)
	}
	return 0, nil
}

func (c *CommentRepository) AllAsDetailed(ctx context.Context, mediaID uuid.UUID, userID uuid.UUID) ([]*model.DetailedComment, error) {
	if c.AllAsDetailedFn != nil {
		return c.AllAsDetailedFn(ctx, mediaID, userID)
	}
	return []*model.DetailedComment{}, nil
}

func (c *CommentRepository) AllRepliesAsDetailed(ctx context.Context, comment *model.Comment, userID uuid.UUID) ([]*model.DetailedComment, error) {
	if c.AllRepliesAsDetailedFn != nil {
		return c.AllRepliesAsDetailedFn(ctx, comment, userID)
	}
	return []*model.DetailedComment{}, nil
}

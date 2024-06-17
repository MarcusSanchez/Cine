package mocks

import (
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
)

var _ repository.ListRepository = (*ListRepository)(nil)

type ListRepository struct {
	OneFn          func(ctx context.Context, filters ...*model.ListF) (*model.List, error)
	AllFn          func(ctx context.Context, filters ...*model.ListF) ([]*model.List, error)
	ExistsFn       func(ctx context.Context, filters ...*model.ListF) (bool, error)
	CountFn        func(ctx context.Context, filters ...*model.ListF) (int, error)
	InsertFn       func(ctx context.Context, entity *model.List) (*model.List, error)
	InsertBulkFn   func(ctx context.Context, entities []*model.List) ([]*model.List, error)
	UpdateFn       func(ctx context.Context, id uuid.UUID, updater *model.ListU) (*model.List, error)
	UpdateExecFn   func(ctx context.Context, updater *model.ListU, filters ...*model.ListF) (int, error)
	DeleteFn       func(ctx context.Context, id uuid.UUID) error
	DeleteExecFn   func(ctx context.Context, filters ...*model.ListF) (int, error)
	AllWithMediaFn func(ctx context.Context, filters ...*model.ListF) ([]*model.ListWithMedia, error)
	OneWithMediaFn func(ctx context.Context, filters ...*model.ListF) (*model.ListWithMedia, error)
	AllMembersFn   func(ctx context.Context, list *model.List) ([]*model.User, error)
	AddMemberFn    func(ctx context.Context, list *model.List, userID uuid.UUID) error
	RemoveMemberFn func(ctx context.Context, list *model.List, userID uuid.UUID) error
	AddMediaFn     func(ctx context.Context, list *model.List, mediaID uuid.UUID) error
	RemoveMediaFn  func(ctx context.Context, list *model.List, mediaID uuid.UUID) error
	AllMediaFn     func(ctx context.Context, list *model.List) ([]*model.Media, error)
}

func NewListRepository() *ListRepository {
	return &ListRepository{}
}

func (l *ListRepository) One(ctx context.Context, filters ...*model.ListF) (*model.List, error) {
	if l.OneFn != nil {
		return l.OneFn(ctx, filters...)
	}
	return &model.List{}, nil
}

func (l *ListRepository) All(ctx context.Context, filters ...*model.ListF) ([]*model.List, error) {
	if l.AllFn != nil {
		return l.AllFn(ctx, filters...)
	}
	return []*model.List{}, nil
}

func (l *ListRepository) Exists(ctx context.Context, filters ...*model.ListF) (bool, error) {
	if l.ExistsFn != nil {
		return l.ExistsFn(ctx, filters...)
	}
	return false, nil
}

func (l *ListRepository) Count(ctx context.Context, filters ...*model.ListF) (int, error) {
	if l.CountFn != nil {
		return l.CountFn(ctx, filters...)
	}
	return 0, nil
}

func (l *ListRepository) Insert(ctx context.Context, entity *model.List) (*model.List, error) {
	if l.InsertFn != nil {
		return l.InsertFn(ctx, entity)
	}
	return &model.List{}, nil
}

func (l *ListRepository) InsertBulk(ctx context.Context, entities []*model.List) ([]*model.List, error) {
	if l.InsertBulkFn != nil {
		return l.InsertBulkFn(ctx, entities)
	}
	return []*model.List{}, nil
}

func (l *ListRepository) Update(ctx context.Context, id uuid.UUID, updater *model.ListU) (*model.List, error) {
	if l.UpdateFn != nil {
		return l.UpdateFn(ctx, id, updater)
	}
	return &model.List{}, nil
}

func (l *ListRepository) UpdateExec(ctx context.Context, updater *model.ListU, filters ...*model.ListF) (int, error) {
	if l.UpdateExecFn != nil {
		return l.UpdateExecFn(ctx, updater, filters...)
	}
	return 0, nil
}

func (l *ListRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if l.DeleteFn != nil {
		return l.DeleteFn(ctx, id)
	}
	return nil
}

func (l *ListRepository) DeleteExec(ctx context.Context, filters ...*model.ListF) (int, error) {
	if l.DeleteExecFn != nil {
		return l.DeleteExecFn(ctx, filters...)
	}
	return 0, nil
}

func (l *ListRepository) AllWithMedia(ctx context.Context, filters ...*model.ListF) ([]*model.ListWithMedia, error) {
	if l.AllWithMediaFn != nil {
		return l.AllWithMediaFn(ctx, filters...)
	}
	return []*model.ListWithMedia{}, nil
}

func (l *ListRepository) OneWithMedia(ctx context.Context, filters ...*model.ListF) (*model.ListWithMedia, error) {
	if l.OneWithMediaFn != nil {
		return l.OneWithMediaFn(ctx, filters...)
	}
	return &model.ListWithMedia{}, nil
}

func (l *ListRepository) AllMembers(ctx context.Context, list *model.List) ([]*model.User, error) {
	if l.AllMembersFn != nil {
		return l.AllMembersFn(ctx, list)
	}
	return []*model.User{}, nil
}

func (l *ListRepository) AddMember(ctx context.Context, list *model.List, userID uuid.UUID) error {
	if l.AddMemberFn != nil {
		return l.AddMemberFn(ctx, list, userID)
	}
	return nil
}

func (l *ListRepository) RemoveMember(ctx context.Context, list *model.List, userID uuid.UUID) error {
	if l.RemoveMemberFn != nil {
		return l.RemoveMemberFn(ctx, list, userID)
	}
	return nil
}

func (l *ListRepository) AddMedia(ctx context.Context, list *model.List, mediaID uuid.UUID) error {
	if l.AddMediaFn != nil {
		return l.AddMediaFn(ctx, list, mediaID)
	}
	return nil
}

func (l *ListRepository) RemoveMedia(ctx context.Context, list *model.List, mediaID uuid.UUID) error {
	if l.RemoveMediaFn != nil {
		return l.RemoveMediaFn(ctx, list, mediaID)
	}
	return nil
}

func (l *ListRepository) AllMedia(ctx context.Context, list *model.List) ([]*model.Media, error) {
	if l.AllMediaFn != nil {
		return l.AllMediaFn(ctx, list)
	}
	return []*model.Media{}, nil
}

package mocks

import (
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	OneFn          func(ctx context.Context, filters ...*model.UserF) (*model.User, error)
	AllFn          func(ctx context.Context, filters ...*model.UserF) ([]*model.User, error)
	ExistsFn       func(ctx context.Context, filters ...*model.UserF) (bool, error)
	CountFn        func(ctx context.Context, filters ...*model.UserF) (int, error)
	InsertFn       func(ctx context.Context, entity *model.User) (*model.User, error)
	InsertBulkFn   func(ctx context.Context, entities []*model.User) ([]*model.User, error)
	UpdateFn       func(ctx context.Context, id uuid.UUID, updater *model.UserU) (*model.User, error)
	UpdateExecFn   func(ctx context.Context, updater *model.UserU, filters ...*model.UserF) (int, error)
	DeleteFn       func(ctx context.Context, id uuid.UUID) error
	DeleteExecFn   func(ctx context.Context, filters ...*model.UserF) (int, error)
	OneDetailedFn  func(ctx context.Context, id, userID uuid.UUID) (*model.DetailedUser, error)
	OneFollowedFn  func(ctx context.Context, user *model.User, followedID uuid.UUID) (*model.User, error)
	AllFollowedFn  func(ctx context.Context, user *model.User) ([]*model.User, error)
	FollowUserFn   func(ctx context.Context, user *model.User, userToFollowID uuid.UUID) error
	UnfollowUserFn func(ctx context.Context, user *model.User, followedID uuid.UUID) error
	OneFollowerFn  func(ctx context.Context, user *model.User, followerID uuid.UUID) (*model.User, error)
	AllFollowersFn func(ctx context.Context, user *model.User) ([]*model.User, error)
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) One(ctx context.Context, filters ...*model.UserF) (*model.User, error) {
	if u.OneFn != nil {
		return u.OneFn(ctx, filters...)
	}
	return &model.User{}, nil
}

func (u *UserRepository) All(ctx context.Context, filters ...*model.UserF) ([]*model.User, error) {
	if u.AllFn != nil {
		return u.AllFn(ctx, filters...)
	}
	return []*model.User{}, nil
}

func (u *UserRepository) Exists(ctx context.Context, filters ...*model.UserF) (bool, error) {
	if u.ExistsFn != nil {
		return u.ExistsFn(ctx, filters...)
	}
	return false, nil
}

func (u *UserRepository) Count(ctx context.Context, filters ...*model.UserF) (int, error) {
	if u.CountFn != nil {
		return u.CountFn(ctx, filters...)
	}
	return 0, nil
}

func (u *UserRepository) Insert(ctx context.Context, entity *model.User) (*model.User, error) {
	if u.InsertFn != nil {
		return u.InsertFn(ctx, entity)
	}
	return &model.User{}, nil
}

func (u *UserRepository) InsertBulk(ctx context.Context, entities []*model.User) ([]*model.User, error) {
	if u.InsertBulkFn != nil {
		return u.InsertBulkFn(ctx, entities)
	}
	return []*model.User{}, nil
}

func (u *UserRepository) Update(ctx context.Context, id uuid.UUID, updater *model.UserU) (*model.User, error) {
	if u.UpdateFn != nil {
		return u.UpdateFn(ctx, id, updater)
	}
	return &model.User{}, nil
}

func (u *UserRepository) UpdateExec(ctx context.Context, updater *model.UserU, filters ...*model.UserF) (int, error) {
	if u.UpdateExecFn != nil {
		return u.UpdateExecFn(ctx, updater, filters...)
	}
	return 0, nil
}

func (u *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if u.DeleteFn != nil {
		return u.DeleteFn(ctx, id)
	}
	return nil
}

func (u *UserRepository) DeleteExec(ctx context.Context, filters ...*model.UserF) (int, error) {
	if u.DeleteExecFn != nil {
		return u.DeleteExecFn(ctx, filters...)
	}
	return 0, nil
}

func (u *UserRepository) OneDetailed(ctx context.Context, id, userID uuid.UUID) (*model.DetailedUser, error) {
	if u.OneDetailedFn != nil {
		return u.OneDetailedFn(ctx, id, userID)
	}
	return &model.DetailedUser{}, nil
}

func (u *UserRepository) OneFollowed(ctx context.Context, user *model.User, followedID uuid.UUID) (*model.User, error) {
	if u.OneFollowedFn != nil {
		return u.OneFollowedFn(ctx, user, followedID)
	}
	return &model.User{}, nil
}

func (u *UserRepository) AllFollowed(ctx context.Context, user *model.User) ([]*model.User, error) {
	if u.AllFollowedFn != nil {
		return u.AllFollowedFn(ctx, user)
	}
	return []*model.User{}, nil
}

func (u *UserRepository) FollowUser(ctx context.Context, user *model.User, userToFollowID uuid.UUID) error {
	if u.FollowUserFn != nil {
		return u.FollowUserFn(ctx, user, userToFollowID)
	}
	return nil
}

func (u *UserRepository) UnfollowUser(ctx context.Context, user *model.User, followedID uuid.UUID) error {
	if u.UnfollowUserFn != nil {
		return u.UnfollowUserFn(ctx, user, followedID)
	}
	return nil
}

func (u *UserRepository) OneFollower(ctx context.Context, user *model.User, followerID uuid.UUID) (*model.User, error) {
	if u.OneFollowerFn != nil {
		return u.OneFollowerFn(ctx, user, followerID)
	}
	return &model.User{}, nil
}

func (u *UserRepository) AllFollowers(ctx context.Context, user *model.User) ([]*model.User, error) {
	if u.AllFollowersFn != nil {
		return u.AllFollowersFn(ctx, user)
	}
	return []*model.User{}, nil
}

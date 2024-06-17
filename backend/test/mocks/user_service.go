package mocks

import (
	"cine/entity/model"
	"cine/service"
	"context"
	"github.com/google/uuid"
)

var _ service.UserService = (*UserServiceMock)(nil)

type UserServiceMock struct {
	GetUserFn         func(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetDetailedUserFn func(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*model.DetailedUser, error)
	UpdateUserFn      func(ctx context.Context, id uuid.UUID, userU *model.UserU) (*model.User, error)
	DeleteUserFn      func(ctx context.Context, id uuid.UUID) error
	FollowUserFn      func(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error
	UnfollowUserFn    func(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error
}

func NewUserService() *UserServiceMock {
	return &UserServiceMock{}
}

func (m *UserServiceMock) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	if m.GetUserFn != nil {
		return m.GetUserFn(ctx, id)
	}
	return &model.User{}, nil
}

func (m *UserServiceMock) GetDetailedUser(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*model.DetailedUser, error) {
	if m.GetDetailedUserFn != nil {
		return m.GetDetailedUserFn(ctx, id, userID)
	}
	return &model.DetailedUser{}, nil
}

func (m *UserServiceMock) UpdateUser(ctx context.Context, id uuid.UUID, userU *model.UserU) (*model.User, error) {
	if m.UpdateUserFn != nil {
		return m.UpdateUserFn(ctx, id, userU)
	}
	return &model.User{}, nil
}

func (m *UserServiceMock) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if m.DeleteUserFn != nil {
		return m.DeleteUserFn(ctx, id)
	}
	return nil
}

func (m *UserServiceMock) FollowUser(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error {
	if m.FollowUserFn != nil {
		return m.FollowUserFn(ctx, followerID, followeeID)
	}
	return nil
}

func (m *UserServiceMock) UnfollowUser(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error {
	if m.UnfollowUserFn != nil {
		return m.UnfollowUserFn(ctx, followerID, followeeID)
	}
	return nil
}

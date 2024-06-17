package unit

import (
	"cine/datastore"
	"cine/entity/model"
	"cine/pkg/fault"
	"cine/service"
	"cine/test/mocks"
	"context"
	"github.com/google/uuid"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {} // redundant test

func TestUserService_GetDetailedUser(t *testing.T) {} // redundant test

func TestUserService_UpdateUser(t *testing.T) {
	assert := testify.New(t)
	ctx := context.Background()
	store := mocks.NewStore()
	us := service.NewUserService(store, mocks.NopLogger{})

	t.Run("success", func(t *testing.T) {
		store.User.ExistsFn = func(ctx context.Context, filters ...*model.UserF) (bool, error) {
			if filters[0].ID != nil {
				return true, nil
			}
			return false, nil
		}

		username := "username"
		email := "email"
		password := "password"

		_, err := us.UpdateUser(ctx, uuid.UUID{}, &model.UserU{
			Username: &username,
			Email:    &email,
			Password: &password,
		})

		assert.Nil(err, "error should be nil")
	})

	t.Run("username already exists", func(t *testing.T) {
		username := "username"

		store.User.ExistsFn = func(ctx context.Context, filters ...*model.UserF) (bool, error) {
			if filters[0].ID != nil || filters[0].Username != nil {
				return true, nil
			}
			return false, nil
		}

		_, err := us.UpdateUser(ctx, uuid.UUID{}, &model.UserU{Username: &username})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeConflict, "error code should be conflict")
	})

	t.Run("email already exists", func(t *testing.T) {
		email := "email"

		store.User.ExistsFn = func(ctx context.Context, filters ...*model.UserF) (bool, error) {
			if filters[0].ID != nil || filters[0].Email != nil {
				return true, nil
			}
			return false, nil
		}

		_, err := us.UpdateUser(ctx, uuid.UUID{}, &model.UserU{Email: &email})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeConflict, "error code should be conflict")
	})

	t.Run("no fields to update", func(t *testing.T) {
		_, err := us.UpdateUser(ctx, uuid.UUID{}, &model.UserU{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeBadRequest, "error code should be bad request")
	})
}

func TestUserService_DeleteUser(t *testing.T) {} // redundant test

func TestUserService_FollowUser(t *testing.T) {
	assert := testify.New(t)
	ctx := context.Background()
	store := mocks.NewStore()
	us := service.NewUserService(store, mocks.NopLogger{})

	t.Run("success", func(t *testing.T) {
		err := us.FollowUser(ctx, uuid.UUID{}, uuid.UUID{})
		assert.Nil(err, "error should be nil")
	})

	t.Run("user not found", func(t *testing.T) {
		store.User.OneFn = func(ctx context.Context, filters ...*model.UserF) (*model.User, error) {
			return nil, datastore.ErrNotFound
		}

		err := us.FollowUser(ctx, uuid.UUID{}, uuid.UUID{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeNotFound, "error code should be not found")

		store.User.OneFn = nil
	})

	t.Run("user already followed", func(t *testing.T) {
		store.User.FollowUserFn = func(ctx context.Context, user *model.User, userToFollowID uuid.UUID) error {
			return datastore.ErrConstraint
		}

		err := us.FollowUser(ctx, uuid.UUID{}, uuid.UUID{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeConflict, "error code should be conflict")
	})
}

func TestUserService_UnfollowUser(t *testing.T) {
	assert := testify.New(t)
	ctx := context.Background()
	store := mocks.NewStore()
	us := service.NewUserService(store, mocks.NopLogger{})

	t.Run("success", func(t *testing.T) {
		err := us.UnfollowUser(ctx, uuid.UUID{}, uuid.UUID{})
		assert.Nil(err, "error should be nil")
	})

	t.Run("user not found", func(t *testing.T) {
		store.User.OneFn = func(ctx context.Context, filters ...*model.UserF) (*model.User, error) {
			return nil, datastore.ErrNotFound
		}

		err := us.UnfollowUser(ctx, uuid.UUID{}, uuid.UUID{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeNotFound, "error code should be not found")

		store.User.OneFn = nil
	})

	t.Run("user already unfollowed", func(t *testing.T) {
		store.User.UnfollowUserFn = func(ctx context.Context, user *model.User, userToFollowID uuid.UUID) error {
			return datastore.ErrConstraint
		}

		err := us.UnfollowUser(ctx, uuid.UUID{}, uuid.UUID{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeConflict, "error code should be conflict")
	})
}

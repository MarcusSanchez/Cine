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
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestAuthService_Register(t *testing.T) {
	assert := testify.New(t)
	ctx := context.Background()
	store := mocks.NewStore()
	as := service.NewAuthService(store, mocks.NopLogger{})

	t.Run("success", func(t *testing.T) {
		_, _, err := as.Register(ctx, &service.RegisterInput{})
		assert.Nil(err, "error should be nil")
	})

	t.Run("username already exists", func(t *testing.T) {
		store.User.ExistsFn = func(ctx context.Context, userFs ...*model.UserF) (bool, error) { return true, nil }
		_, _, err := as.Register(ctx, &service.RegisterInput{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeConflict, "error code should be conflict")
	})

	t.Run("email already exists", func(t *testing.T) {
		store.User.ExistsFn = func(ctx context.Context, userFs ...*model.UserF) (bool, error) {
			if userFs[0].Email != nil {
				return true, nil
			}
			return false, nil
		}

		_, _, err := as.Register(ctx, &service.RegisterInput{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeConflict, "error code should be conflict")
	})
}

func TestAuthService_Login(t *testing.T) {
	assert := testify.New(t)
	ctx := context.Background()
	store := mocks.NewStore()
	as := service.NewAuthService(store, mocks.NopLogger{})

	t.Run("success", func(t *testing.T) {
		store.User.OneFn = func(ctx context.Context, filters ...*model.UserF) (*model.User, error) {
			password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
			return &model.User{Password: string(password)}, nil
		}

		_, _, err := as.Login(ctx, "username", "password")
		assert.Nil(err, "error should be nil")
	})

	t.Run("user not found", func(t *testing.T) {
		store.User.OneFn = func(ctx context.Context, filters ...*model.UserF) (*model.User, error) {
			return nil, datastore.ErrNotFound
		}

		_, _, err := as.Login(ctx, "username", "password")
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeNotFound, "error code should be not found")
	})

	t.Run("password mismatch", func(t *testing.T) {
		store.User.OneFn = func(ctx context.Context, filters ...*model.UserF) (*model.User, error) {
			password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
			return &model.User{Password: string(password)}, nil
		}

		_, _, err := as.Login(ctx, "username", "wrong-password")
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeUnauthorized, "error code should be unauthorized")
	})
}

func TestAuthService_Logout(t *testing.T) {} // Redundant test

func TestAuthService_Authenticate(t *testing.T) {
	assert := testify.New(t)
	ctx := context.Background()
	store := mocks.NewStore()
	as := service.NewAuthService(store, mocks.NopLogger{})

	t.Run("success", func(t *testing.T) {
		_, _, err := as.Authenticate(ctx, &model.Session{Expiration: time.Now().Add(24 * time.Hour)})
		assert.Nil(err, "error should be nil")
	})

	t.Run("rejects expired session", func(t *testing.T) {
		_, _, err := as.Authenticate(ctx, &model.Session{Expiration: time.Now().Add(-24 * time.Hour)})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeUnauthorized, "error code should be unauthorized")
	})

	t.Run("user not found", func(t *testing.T) {
		store.User.OneFn = func(ctx context.Context, filters ...*model.UserF) (*model.User, error) {
			return nil, datastore.ErrNotFound
		}

		_, _, err := as.Authenticate(ctx, &model.Session{Expiration: time.Now().Add(24 * time.Hour)})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeNotFound, "error code should be not found")
	})
}

func TestAuthService_Session(t *testing.T) {
	assert := testify.New(t)
	ctx := context.Background()
	store := mocks.NewStore()
	as := service.NewAuthService(store, mocks.NopLogger{})

	t.Run("success", func(t *testing.T) {
		store.Session.OneFn = func(ctx context.Context, filters ...*model.SessionF) (*model.Session, error) {
			return &model.Session{Expiration: time.Now().Add(24 * time.Hour)}, nil
		}

		_, err := as.Session(ctx, uuid.UUID{})
		assert.Nil(err, "error should be nil")
	})

	t.Run("session not found", func(t *testing.T) {
		store.Session.OneFn = func(ctx context.Context, filters ...*model.SessionF) (*model.Session, error) {
			return nil, datastore.ErrNotFound
		}

		_, err := as.Session(ctx, uuid.UUID{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeNotFound, "error code should be not found")
	})

	t.Run("rejects expired session", func(t *testing.T) {
		store.Session.OneFn = func(ctx context.Context, filters ...*model.SessionF) (*model.Session, error) {
			return &model.Session{Expiration: time.Now().Add(-24 * time.Hour)}, nil
		}

		_, err := as.Session(ctx, uuid.UUID{})
		assert.NotNil(err, "error should be not nil")

		e, _ := fault.As(err)
		assert.Equal(e.Code, fault.CodeUnauthorized, "error code should be unauthorized")
	})
}

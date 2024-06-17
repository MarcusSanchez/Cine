package mocks

import (
	"cine/entity/model"
	"cine/service"
	"context"
	"github.com/google/uuid"
)

var _ service.AuthService = (*AuthServiceMock)(nil)

type AuthServiceMock struct {
	RegisterFn     func(ctx context.Context, input *service.RegisterInput) (*model.User, *model.Session, error)
	LoginFn        func(ctx context.Context, username, password string) (*model.User, *model.Session, error)
	LogoutFn       func(ctx context.Context, session *model.Session) error
	AuthenticateFn func(ctx context.Context, session *model.Session) (*model.User, *model.Session, error)
	SessionFn      func(ctx context.Context, access uuid.UUID) (*model.Session, error)
}

func NewAuthService() *AuthServiceMock {
	return &AuthServiceMock{}
}

func (m *AuthServiceMock) Register(ctx context.Context, input *service.RegisterInput) (*model.User, *model.Session, error) {
	if m.RegisterFn != nil {
		return m.RegisterFn(ctx, input)
	}
	return &model.User{}, &model.Session{}, nil
}

func (m *AuthServiceMock) Login(ctx context.Context, username, password string) (*model.User, *model.Session, error) {
	if m.LoginFn != nil {
		return m.LoginFn(ctx, username, password)
	}
	return &model.User{}, &model.Session{}, nil
}

func (m *AuthServiceMock) Logout(ctx context.Context, session *model.Session) error {
	if m.LogoutFn != nil {
		return m.LogoutFn(ctx, session)
	}
	return nil
}

func (m *AuthServiceMock) Authenticate(ctx context.Context, session *model.Session) (*model.User, *model.Session, error) {
	if m.AuthenticateFn != nil {
		return m.AuthenticateFn(ctx, session)
	}
	return &model.User{}, &model.Session{}, nil
}

func (m *AuthServiceMock) Session(ctx context.Context, access uuid.UUID) (*model.Session, error) {
	if m.SessionFn != nil {
		return m.SessionFn(ctx, access)
	}
	return &model.Session{}, nil
}

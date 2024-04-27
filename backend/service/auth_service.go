package service

import (
	"cine/datastore"
	"cine/entity/model"
	"cine/pkg/fault"
	"cine/pkg/logger"
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, input *RegisterInput) (*model.User, *model.Session, error)
	Login(ctx context.Context, username, password string) (*model.Session, error)
	Logout(ctx context.Context, session *model.Session) error
	Authenticate(ctx context.Context, session *model.Session) (*model.User, *model.Session, error)
	Session(ctx context.Context, access uuid.UUID) (*model.Session, error)
}

type authService struct {
	store  datastore.Store
	logger logger.Logger
}

func NewAuthService(store datastore.Store, logger logger.Logger) AuthService {
	return &authService{store: store, logger: logger}
}

type RegisterInput struct {
	DisplayName    string
	Email          string
	Username       string
	Password       string
	ProfilePicture string
}

func (as authService) Register(ctx context.Context, input *RegisterInput) (*model.User, *model.Session, error) {
	exists, err := as.store.Users().Exists(ctx, &model.UserF{Username: &input.Username})
	if err != nil {
		as.logger.Error("exists check on username failed", err)
		return nil, nil, fault.Internal("error registering user")
	} else if exists {
		return nil, nil, fault.Conflict("username already exists")
	}

	exists, err = as.store.Users().Exists(ctx, &model.UserF{Email: &input.Email})
	if err != nil {
		as.logger.Error("exists check on email failed", err)
		return nil, nil, fault.Internal("error registering user")
	} else if exists {
		return nil, nil, fault.Conflict("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		as.logger.Error("password hashing failed", err)
		return nil, nil, fault.Internal("error registering user")
	}

	tx, err := as.store.Transaction(ctx)
	if err != nil {
		as.logger.Error("transaction creation failed", err)
		return nil, nil, fault.Internal("error registering user")
	}
	defer tx.Rollback()

	user, err := tx.Users().Insert(
		ctx, &model.User{
			DisplayName:    input.DisplayName,
			Username:       input.Username,
			Password:       string(hashedPassword),
			Email:          input.Email,
			ProfilePicture: input.ProfilePicture,
		},
	)
	if err != nil {
		as.logger.Error("user creation failed", err)
		return nil, nil, fault.Internal("error registering user")
	}

	session, err := tx.Sessions().Insert(
		ctx, &model.Session{
			UserID:     user.ID,
			CSRF:       uuid.New(),
			Token:      uuid.New(),
			Expiration: time.Now().Add(model.SessionTokenDuration),
		},
	)
	if err != nil {
		as.logger.Error("session creation failed", err)
		return nil, nil, fault.Internal("error registering user")
	}

	if err = tx.Commit(); err != nil {
		as.logger.Error("transaction commit failed", err)
		return nil, nil, fault.Internal("error registering user")
	}

	return user, session, nil
}

func (as authService) Login(ctx context.Context, username, password string) (*model.Session, error) {
	user, err := as.store.Users().One(ctx, &model.UserF{Username: &username})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("user not found")
		}
		as.logger.Error("user retrieval failed", err)
		return nil, fault.Internal("error logging in")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, fault.Unauthorized("mismatch username and password")
		}
		as.logger.Error("password comparison failed", err)
		return nil, fault.Internal("error logging in")
	}

	session, err := as.store.Sessions().Insert(
		ctx, &model.Session{
			UserID:     user.ID,
			CSRF:       uuid.New(),
			Token:      uuid.New(),
			Expiration: time.Now().Add(model.SessionTokenDuration),
		},
	)
	if err != nil {
		as.logger.Error("session creation failed", err)
		return nil, fault.Internal("error logging in")
	}

	return session, nil
}

func (as authService) Logout(ctx context.Context, session *model.Session) error {
	if err := as.store.Sessions().Delete(ctx, session.ID); err != nil {
		as.logger.Error("session deletion failed", err)
		return fault.Internal("error logging out")
	}
	return nil
}

func (as authService) Authenticate(ctx context.Context, session *model.Session) (*model.User, *model.Session, error) {
	if session.Expiration.Before(time.Now()) {
		return nil, session, fault.Unauthorized("session has expired")
	}

	user, err := as.store.Users().One(ctx, &model.UserF{ID: &session.UserID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, session, fault.NotFound("user not found")
		}
		as.logger.Error("user retrieval failed", err)
		return nil, session, fault.Internal("error authenticating user")
	}

	session, _ = as.refresh(ctx, session)
	return user, session, nil
}

func (as authService) Session(ctx context.Context, token uuid.UUID) (*model.Session, error) {
	session, err := as.store.Sessions().One(ctx, &model.SessionF{Token: &token})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("session not found")
		}
		as.logger.Error("session retrieval failed", err)
		return nil, fault.Internal("error retrieving session")
	}

	if session.Expiration.Before(time.Now()) {
		return nil, fault.Unauthorized("session has expired")
	}

	return session, nil
}

func (as authService) refresh(ctx context.Context, session *model.Session) (*model.Session, error) {
	if session.Expiration.Before(time.Now()) {
		return session, fault.Unauthorized("session has expired")
	}

	expiration := as.expiration(session.CreatedAt)

	var err error
	// if the expiration time isn't already set to the session's absolute deadline
	if expiration != session.CreatedAt.Add(model.SessionTokenAbsoluteDuration) {
		session, err = as.store.Sessions().Update(ctx, session.ID, &model.SessionU{Expiration: &expiration})
		if err != nil {
			as.logger.Error("session update failed", err)
			return nil, fault.Internal("error refreshing session")
		}
	}

	return session, nil
}

// expiration returns the correct expiration time for a session token when refreshed.
// Expirations last 7 days, but can be extended back to 7 days if the session is refreshed, unless
// extending the expiration would exceed the absolute lifetime of a session token, which is 1 month, in which
// case the expiration is set to the absolute deadline.
func (as authService) expiration(createdAt time.Time) time.Time {
	relative := time.Now().Add(model.SessionTokenDuration)        // now + duration (1 week)
	absolute := createdAt.Add(model.SessionTokenAbsoluteDuration) // createdAt + duration (1 month)
	return absolute.Add(relative.Sub(absolute))                   // min(relative, absolute)
}

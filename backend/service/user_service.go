package service

import (
	"cine/datastore"
	"cine/entity/model"
	"cine/pkg/fault"
	"cine/pkg/logger"
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetDetailedUser(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*model.DetailedUser, error)
	UpdateUser(ctx context.Context, id uuid.UUID, userU *model.UserU) (*model.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	FollowUser(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error
	UnfollowUser(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error
}

type userService struct {
	store  datastore.Store
	logger logger.Logger
}

func NewUserService(store datastore.Store, logger logger.Logger) UserService {
	return &userService{
		store:  store,
		logger: logger,
	}
}

func (us userService) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user, err := us.store.Users().One(ctx, &model.UserF{ID: &id})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("user not found")
		}
		us.logger.Error("user retrieval failed", err)
		return nil, fault.Internal("error retrieving user")
	}

	return user, nil
}

func (us userService) GetDetailedUser(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*model.DetailedUser, error) {
	user, err := us.store.Users().OneDetailed(ctx, id, userID)
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("user not found")
		}
		us.logger.Error("user retrieval failed", err)
		return nil, fault.Internal("error retrieving user")
	}

	return user, nil
}

func (us userService) UpdateUser(ctx context.Context, id uuid.UUID, userU *model.UserU) (*model.User, error) {
	if !us.hasFieldToUpdate(userU) {
		return nil, fault.BadRequest("no fields to update")
	}

	exists, err := us.store.Users().Exists(ctx, &model.UserF{ID: &id})
	if err != nil {
		us.logger.Error("user retrieval failed", err)
		return nil, fault.Internal("error updating user")
	} else if !exists {
		return nil, fault.NotFound("user not found")
	}

	if userU.Username != nil {
		exists, err = us.store.Users().Exists(ctx, &model.UserF{Username: userU.Username})
		if err != nil {
			us.logger.Error("exists check on username failed", err)
			return nil, fault.Internal("error updating user")
		} else if exists {
			return nil, fault.Conflict("username already exists")
		}
	}
	if userU.Email != nil {
		exists, err = us.store.Users().Exists(ctx, &model.UserF{Email: userU.Email})
		if err != nil {
			us.logger.Error("exists check on email failed", err)
			return nil, fault.Internal("error updating user")
		} else if exists {
			return nil, fault.Conflict("email already exists")
		}
	}
	if userU.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*userU.Password), bcrypt.DefaultCost)
		if err != nil {
			us.logger.Error("password hashing failed", err)
			return nil, fault.Internal("error updating user")
		}
		password := string(hashed)
		userU.Password = &password
	}

	user, err := us.store.Users().Update(ctx, id, userU)
	if err != nil {
		us.logger.Error("user update failed", err)
		return nil, fault.Internal("error updating user")
	}

	return user, nil
}

func (us userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := us.store.Users().Delete(ctx, id)
	if err != nil {
		us.logger.Error("user deletion failed", err)
		return fault.Internal("error deleting user")
	}

	return nil
}

func (us userService) FollowUser(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error {
	user, err := us.store.Users().One(ctx, &model.UserF{ID: &followerID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("user not found")
		}
		us.logger.Error("user retrieval failed", err)
		return fault.Internal("error following user")
	}

	if err = us.store.Users().FollowUser(ctx, user, followeeID); err != nil {
		if datastore.IsConstraint(err) {
			return fault.Conflict("already following user")
		}
		us.logger.Error("user follow failed", err)
		return fault.Internal("error following user")
	}

	return nil
}

func (us userService) UnfollowUser(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) error {
	user, err := us.store.Users().One(ctx, &model.UserF{ID: &followerID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("user not found")
		}
		us.logger.Error("user retrieval failed", err)
		return fault.Internal("error unfollowing user")
	}

	if err = us.store.Users().UnfollowUser(ctx, user, followeeID); err != nil {
		if datastore.IsConstraint(err) {
			return fault.Conflict("not following user")
		}
		us.logger.Error("user unfollow failed", err)
		return fault.Internal("error unfollowing user")
	}

	return nil
}

func (us userService) hasFieldToUpdate(userU *model.UserU) bool {
	return userU.DisplayName != nil ||
		userU.Username != nil ||
		userU.Email != nil ||
		userU.Password != nil ||
		userU.ProfilePicture != nil
}

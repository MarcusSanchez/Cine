package ent

import (
	"cine/datastore/ent/ent"
	"cine/datastore/ent/ent/predicate"
	User "cine/datastore/ent/ent/user"
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type userRepository struct {
	client *ent.Client
}

func newUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{client: client}
}

func (ur *userRepository) One(ctx context.Context, userFs ...*model.UserF) (*model.User, error) {
	q := ur.client.User.Query()
	q = q.Where(ur.filters(userFs)...)

	user, err := q.First(ctx)
	return c.user(user), c.error(err)
}

func (ur *userRepository) All(ctx context.Context, userFs ...*model.UserF) ([]*model.User, error) {
	q := ur.client.User.Query()
	q = q.Where(ur.filters(userFs)...)

	users, err := q.All(ctx)
	return c.users(users), c.error(err)
}

func (ur *userRepository) Exists(ctx context.Context, userFs ...*model.UserF) (bool, error) {
	q := ur.client.User.Query()
	q = q.Where(ur.filters(userFs)...)

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (ur *userRepository) Count(ctx context.Context, userFs ...*model.UserF) (int, error) {
	q := ur.client.User.Query()
	q = q.Where(ur.filters(userFs)...)

	count, err := q.Count(ctx)
	return count, c.error(err)
}

func (ur *userRepository) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	s := ur.create(user)

	iUser, err := s.Save(ctx)
	return c.user(iUser), c.error(err)
}

func (ur *userRepository) InsertBulk(ctx context.Context, users []*model.User) ([]*model.User, error) {
	q := ur.createBulk(users)

	iUsers, err := q.Save(ctx)
	return c.users(iUsers), c.error(err)
}

func (ur *userRepository) Update(ctx context.Context, id uuid.UUID, userU *model.UserU) (*model.User, error) {
	q := ur.client.User.UpdateOneID(id)

	q.SetUpdatedAt(time.Now())
	q.SetNillableDisplayName(userU.DisplayName)
	q.SetNillableUsername(userU.Username)
	q.SetNillableEmail(userU.Email)
	q.SetNillablePassword(userU.Password)
	q.SetNillableProfilePicture(userU.ProfilePicture)

	user, err := q.Save(ctx)
	return c.user(user), c.error(err)
}

func (ur *userRepository) UpdateExec(ctx context.Context, userU *model.UserU, userFs ...*model.UserF) (int, error) {
	q := ur.client.User.Update()
	q = q.Where(ur.filters(userFs)...)

	q.SetUpdatedAt(time.Now())
	q.SetNillableDisplayName(userU.DisplayName)
	q.SetNillableUsername(userU.Username)
	q.SetNillableEmail(userU.Email)
	q.SetNillablePassword(userU.Password)
	q.SetNillableProfilePicture(userU.ProfilePicture)

	affected, err := q.Save(ctx)
	return affected, c.error(err)
}

func (ur *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := ur.client.User.DeleteOneID(id)

	err := q.Exec(ctx)
	return c.error(err)
}

func (ur *userRepository) DeleteExec(ctx context.Context, userFs ...*model.UserF) (int, error) {
	q := ur.client.User.Delete()
	q = q.Where(ur.filters(userFs)...)

	affected, err := q.Exec(ctx)
	return affected, c.error(err)
}

func (ur *userRepository) OneFriend(ctx context.Context, user *model.User, friendID uuid.UUID) (*model.User, error) {
	q := ur.client.User.Query().
		Where(User.ID(user.ID)).
		QueryFriends().
		Where(User.ID(friendID))

	friend, err := q.First(ctx)
	return c.user(friend), c.error(err)
}

func (ur *userRepository) AllFriends(ctx context.Context, user *model.User) ([]*model.User, error) {
	q := ur.client.User.Query()
	q = q.Where(User.ID(user.ID)).
		QueryFriends()

	friends, err := q.All(ctx)
	return c.users(friends), c.error(err)
}

func (ur *userRepository) AddFriend(ctx context.Context, user *model.User, friendID uuid.UUID) error {
	q := ur.client.User.UpdateOneID(user.ID)
	q = q.AddFriendIDs(friendID)

	_, err := q.Save(ctx)
	return c.error(err)
}

func (ur *userRepository) RemoveFriend(ctx context.Context, user *model.User, friendID uuid.UUID) error {
	q := ur.client.User.UpdateOneID(user.ID)
	q = q.RemoveFriendIDs(friendID)

	_, err := q.Save(ctx)
	return c.error(err)
}

func (ur *userRepository) OneFollowed(ctx context.Context, user *model.User, followedID uuid.UUID) (*model.User, error) {
	q := ur.client.User.Query()
	q = q.Where(User.ID(user.ID)).
		QueryFollowing()
	q = q.Where(User.ID(followedID))

	followed, err := q.First(ctx)
	return c.user(followed), c.error(err)
}

func (ur *userRepository) AllFollowed(ctx context.Context, user *model.User) ([]*model.User, error) {
	q := ur.client.User.Query()
	q = q.Where(User.ID(user.ID)).
		QueryFollowing()

	followed, err := q.All(ctx)
	return c.users(followed), c.error(err)
}

func (ur *userRepository) FollowUser(ctx context.Context, user *model.User, userToFollowID uuid.UUID) error {
	q := ur.client.User.UpdateOneID(user.ID)
	q = q.AddFollowingIDs(userToFollowID)

	_, err := q.Save(ctx)
	return c.error(err)

}

func (ur *userRepository) UnfollowUser(ctx context.Context, user *model.User, followedID uuid.UUID) error {
	q := ur.client.User.UpdateOneID(user.ID)
	q = q.RemoveFollowingIDs(followedID)

	_, err := q.Save(ctx)
	return c.error(err)
}

func (ur *userRepository) OneFollower(ctx context.Context, user *model.User, followerID uuid.UUID) (*model.User, error) {
	q := ur.client.User.Query()
	q = q.Where(User.ID(user.ID)).
		QueryFollowers()
	q = q.Where(User.ID(followerID))

	follower, err := q.First(ctx)
	return c.user(follower), c.error(err)
}

func (ur *userRepository) AllFollowers(ctx context.Context, user *model.User) ([]*model.User, error) {
	q := ur.client.User.Query()
	q = q.Where(User.ID(user.ID)).
		QueryFollowers()

	followers, err := q.All(ctx)
	return c.users(followers), c.error(err)
}

func (ur *userRepository) filters(userFs []*model.UserF) []predicate.User {
	var userF *model.UserF
	if len(userFs) > 0 {
		userF = userFs[0]
	}
	var filters []predicate.User
	if userF != nil {
		if userF.ID != nil {
			filters = append(filters, User.ID(*userF.ID))
		}
		if userF.DisplayName != nil {
			filters = append(filters, User.DisplayName(*userF.DisplayName))
		}
		if userF.Username != nil {
			filters = append(filters, User.Username(*userF.Username))
		}
		if userF.Email != nil {
			filters = append(filters, User.Email(*userF.Email))
		}
		if userF.Password != nil {
			filters = append(filters, User.Password(*userF.Password))
		}
		if userF.ProfilePicture != nil {
			filters = append(filters, User.ProfilePicture(*userF.ProfilePicture))
		}
		if userF.CreatedAt != nil {
			filters = append(filters, User.CreatedAt(*userF.CreatedAt))
		}
		if userF.UpdatedAt != nil {
			filters = append(filters, User.UpdatedAt(*userF.UpdatedAt))
		}
	}
	return filters
}

func (ur *userRepository) create(user *model.User) *ent.UserCreate {
	return ur.client.User.Create().
		SetID(uuid.New()).
		SetDisplayName(user.DisplayName).
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetProfilePicture(user.ProfilePicture).
		SetCreatedAt(time.Now())
}

func (ur *userRepository) createBulk(users []*model.User) *ent.UserCreateBulk {
	builders := make([]*ent.UserCreate, 0, len(users))
	for _, user := range users {
		builders = append(builders, ur.create(user))
	}
	return ur.client.User.CreateBulk(builders...)
}

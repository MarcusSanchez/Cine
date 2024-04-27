package ent

import (
	"cine/datastore/ent/ent"
	Like "cine/datastore/ent/ent/like"
	"cine/datastore/ent/ent/predicate"
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type likeRepository struct {
	client *ent.Client
}

func newLikeRepository(client *ent.Client) repository.LikeRepository {
	return &likeRepository{client: client}
}

func (lr *likeRepository) One(ctx context.Context, likeFs ...*model.LikeF) (*model.Like, error) {
	q := lr.client.Like.Query()
	q = q.Where(lr.filters(likeFs)...)

	like, err := q.First(ctx)
	return c.like(like), c.error(err)
}

func (lr *likeRepository) All(ctx context.Context, likeFs ...*model.LikeF) ([]*model.Like, error) {
	q := lr.client.Like.Query()
	q = q.Where(lr.filters(likeFs)...)

	likes, err := q.All(ctx)
	return c.likes(likes), c.error(err)
}

func (lr *likeRepository) Exists(ctx context.Context, likeFs ...*model.LikeF) (bool, error) {
	q := lr.client.Like.Query()
	q = q.Where(lr.filters(likeFs)...)

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (lr *likeRepository) Count(ctx context.Context, likeFs ...*model.LikeF) (int, error) {
	q := lr.client.Like.Query()
	q = q.Where(lr.filters(likeFs)...)

	count, err := q.Count(ctx)
	return count, c.error(err)
}

func (lr *likeRepository) Insert(ctx context.Context, like *model.Like) (*model.Like, error) {
	i := lr.create(like)

	iLike, err := i.Save(ctx)
	return c.like(iLike), c.error(err)
}

func (lr *likeRepository) InsertBulk(ctx context.Context, likes []*model.Like) ([]*model.Like, error) {
	i := lr.createBulk(likes)

	iLikes, err := i.Save(ctx)
	return c.likes(iLikes), c.error(err)
}

func (lr *likeRepository) Update(ctx context.Context, id uuid.UUID, _ *model.LikeU) (*model.Like, error) {
	q := lr.client.Like.UpdateOneID(id)

	q.SetUpdatedAt(time.Now())

	like, err := q.Save(ctx)
	return c.like(like), c.error(err)
}

func (lr *likeRepository) UpdateExec(ctx context.Context, _ *model.LikeU, likeFs ...*model.LikeF) (int, error) {
	q := lr.client.Like.Update()
	q = q.Where(lr.filters(likeFs)...)

	q.SetUpdatedAt(time.Now())

	affected, err := q.Save(ctx)
	return affected, c.error(err)
}

func (lr *likeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := lr.client.Like.DeleteOneID(id)

	err := q.Exec(ctx)
	return c.error(err)
}

func (lr *likeRepository) DeleteExec(ctx context.Context, likeFs ...*model.LikeF) (int, error) {
	q := lr.client.Like.Delete()
	q = q.Where(lr.filters(likeFs)...)

	affected, err := q.Exec(ctx)
	return affected, c.error(err)
}

func (lr *likeRepository) filters(likeFs []*model.LikeF) []predicate.Like {
	var likeF *model.LikeF
	if len(likeFs) > 0 {
		likeF = likeFs[0]
	}
	var filters []predicate.Like
	if likeF != nil {
		if likeF.ID != nil {
			filters = append(filters, Like.ID(*likeF.ID))
		}
		if likeF.UserID != nil {
			filters = append(filters, Like.UserID(*likeF.UserID))
		}
		if likeF.CommentID != nil {
			filters = append(filters, Like.CommentID(*likeF.CommentID))
		}
		if likeF.CreatedAt != nil {
			filters = append(filters, Like.CreatedAt(*likeF.CreatedAt))
		}
		if likeF.UpdatedAt != nil {
			filters = append(filters, Like.UpdatedAt(*likeF.UpdatedAt))
		}
	}
	return filters
}

func (lr *likeRepository) create(like *model.Like) *ent.LikeCreate {
	return lr.client.Like.Create().
		SetID(uuid.New()).
		SetUserID(like.UserID).
		SetCommentID(like.CommentID).
		SetCreatedAt(time.Now())
}

func (lr *likeRepository) createBulk(likes []*model.Like) *ent.LikeCreateBulk {
	builders := make([]*ent.LikeCreate, 0, len(likes))
	for _, like := range likes {
		builders = append(builders, lr.create(like))
	}
	return lr.client.Like.CreateBulk(builders...)
}

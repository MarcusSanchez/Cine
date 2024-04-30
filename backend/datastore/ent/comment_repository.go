package ent

import (
	"cine/datastore/ent/ent"
	Comment "cine/datastore/ent/ent/comment"
	"cine/datastore/ent/ent/predicate"
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type commentRepository struct {
	client *ent.Client
}

func newCommentRepository(client *ent.Client) repository.CommentRepository {
	return &commentRepository{client: client}
}

func (cr *commentRepository) One(ctx context.Context, commentFs ...*model.CommentF) (*model.Comment, error) {
	q := cr.client.Comment.Query()
	q = q.Where(cr.filters(commentFs)...)

	comment, err := q.First(ctx)
	return c.comment(comment), c.error(err)
}

func (cr *commentRepository) All(ctx context.Context, commentFs ...*model.CommentF) ([]*model.Comment, error) {
	q := cr.client.Comment.Query()
	q = q.Where(cr.filters(commentFs)...)

	comments, err := q.All(ctx)
	return c.comments(comments), c.error(err)
}

func (cr *commentRepository) Exists(ctx context.Context, commentFs ...*model.CommentF) (bool, error) {
	q := cr.client.Comment.Query()
	q = q.Where(cr.filters(commentFs)...)

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (cr *commentRepository) Count(ctx context.Context, commentFs ...*model.CommentF) (int, error) {
	q := cr.client.Comment.Query()
	q = q.Where(cr.filters(commentFs)...)

	count, err := q.Count(ctx)
	return count, c.error(err)
}

func (cr *commentRepository) Insert(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	i := cr.create(comment)

	iComment, err := i.Save(ctx)
	return c.comment(iComment), c.error(err)
}

func (cr *commentRepository) InsertBulk(ctx context.Context, comments []*model.Comment) ([]*model.Comment, error) {
	i := cr.createBulk(comments)

	iComments, err := i.Save(ctx)
	return c.comments(iComments), c.error(err)
}

func (cr *commentRepository) Update(ctx context.Context, id uuid.UUID, commentU *model.CommentU) (*model.Comment, error) {
	q := cr.client.Comment.UpdateOneID(id)

	q.SetUpdatedAt(time.Now())
	q.SetNillableContent(commentU.Content)

	comment, err := q.Save(ctx)
	return c.comment(comment), c.error(err)
}

func (cr *commentRepository) UpdateExec(ctx context.Context, commentU *model.CommentU, commentFs ...*model.CommentF) (int, error) {
	q := cr.client.Comment.Update()
	q = q.Where(cr.filters(commentFs)...)

	q.SetUpdatedAt(time.Now())
	q.SetNillableContent(commentU.Content)

	affected, err := q.Save(ctx)
	return affected, c.error(err)
}

func (cr *commentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := cr.client.Comment.DeleteOneID(id)

	err := q.Exec(ctx)
	return c.error(err)
}

func (cr *commentRepository) DeleteExec(ctx context.Context, commentFs ...*model.CommentF) (int, error) {
	q := cr.client.Comment.Delete()
	q = q.Where(cr.filters(commentFs)...)

	affected, err := q.Exec(ctx)
	return affected, c.error(err)
}

func (cr *commentRepository) AllAsDetailed(ctx context.Context, mediaID, userID uuid.UUID) ([]*model.DetailedComment, error) {
	q := cr.client.Comment.Query().WithLikes().WithReplies().WithUser()
	q = q.Where(Comment.MediaID(mediaID), Comment.Not(Comment.HasReplyingTo())) // top-level comments only

	comments, err := q.All(ctx)
	if err != nil {
		return nil, c.error(err)
	}

	return cr.detailedComments(comments, userID), nil
}

func (cr *commentRepository) AllRepliesAsDetailed(ctx context.Context, comment *model.Comment, userID uuid.UUID) ([]*model.DetailedComment, error) {
	q := cr.client.Comment.Query()
	q = q.Where(Comment.ID(comment.ID)).
		QueryReplies().
		WithUser().WithLikes().WithReplies()

	replies, err := q.All(ctx)
	if err != nil {
		return nil, c.error(err)
	}

	return cr.detailedComments(replies, userID), nil
}

func (cr *commentRepository) filters(commentFs []*model.CommentF) []predicate.Comment {
	var commentF *model.CommentF
	if len(commentFs) > 0 {
		commentF = commentFs[0]
	}
	var filters []predicate.Comment
	if commentF != nil {
		if commentF.ID != nil {
			filters = append(filters, Comment.ID(*commentF.ID))
		}
		if commentF.UserID != nil {
			filters = append(filters, Comment.UserID(*commentF.UserID))
		}
		if commentF.MediaID != nil {
			filters = append(filters, Comment.MediaID(*commentF.MediaID))
		}
		if commentF.ReplyingToID != nil {
			filters = append(filters, Comment.ReplyingToID(*commentF.ReplyingToID))
		}
		if commentF.Content != nil {
			filters = append(filters, Comment.Content(*commentF.Content))
		}
		if commentF.CreatedAt != nil {
			filters = append(filters, Comment.CreatedAt(*commentF.CreatedAt))
		}
		if commentF.UpdatedAt != nil {
			filters = append(filters, Comment.UpdatedAt(*commentF.UpdatedAt))
		}
	}
	return filters
}

func (cr *commentRepository) create(comment *model.Comment) *ent.CommentCreate {
	return cr.client.Comment.Create().
		SetID(uuid.New()).
		SetUserID(*comment.UserID).
		SetNillableReplyingToID(comment.ReplyingToID).
		SetMediaID(comment.MediaID).
		SetContent(comment.Content).
		SetCreatedAt(time.Now())
}

func (cr *commentRepository) createBulk(comments []*model.Comment) *ent.CommentCreateBulk {
	builders := make([]*ent.CommentCreate, 0, len(comments))
	for _, comment := range comments {
		builders = append(builders, cr.create(comment))
	}
	return cr.client.Comment.CreateBulk(builders...)
}

func (cr *commentRepository) detailedComments(comments []*ent.Comment, userID uuid.UUID) []*model.DetailedComment {
	detailedComments := make([]*model.DetailedComment, 0, len(comments))
	for _, comment := range comments {
		detailedComments = append(detailedComments, &model.DetailedComment{
			Comment:      c.comment(comment),
			User:         c.user(comment.Edges.User),
			RepliesCount: len(comment.Edges.Replies),
			LikesCount:   len(comment.Edges.Likes),
			LikedByUser:  cr.likedByUser(comment, userID),
		})
	}
	return detailedComments
}

func (cr *commentRepository) likedByUser(comment *ent.Comment, userID uuid.UUID) bool {
	if userID != uuid.Nil {
		for _, like := range comment.Edges.Likes {
			if like.UserID == userID {
				return true
			}
		}
	}
	return false
}

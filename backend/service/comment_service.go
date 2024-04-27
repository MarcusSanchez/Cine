package service

import (
	"cine/datastore"
	"cine/entity/model"
	"cine/pkg/fault"
	"cine/pkg/logger"
	"context"
	"github.com/google/uuid"
)

type CommentService interface {
	CreateComment(ctx context.Context, input CreateCommentInput) (*model.Comment, error)
	UpdateComment(ctx context.Context, userID, commentID uuid.UUID, content string) (*model.Comment, error)
	DeleteComment(ctx context.Context, userID, commentID uuid.UUID) error
	GetComments(ctx context.Context, ref int, mediaType model.MediaType) ([]*model.CommentWithRelationsCount, error)
	GetCommentReplies(ctx context.Context, commentID uuid.UUID) ([]*model.CommentWithRelationsCount, error)
	LikeComment(ctx context.Context, like *model.Like) (*model.Like, error)
	UnlikeComment(ctx context.Context, userID, likeID uuid.UUID) error
}

type commentService struct {
	store  datastore.Store
	logger logger.Logger
	media  MediaService
}

func NewCommentService(store datastore.Store, logger logger.Logger, media MediaService) CommentService {
	return &commentService{
		store:  store,
		logger: logger,
		media:  media,
	}
}

type CreateCommentInput struct {
	UserID    uuid.UUID
	Comment   *model.Comment
	Ref       int
	MediaType model.MediaType
}

func (cs *commentService) CreateComment(ctx context.Context, input CreateCommentInput) (*model.Comment, error) {
	exists, err := cs.store.Users().Exists(ctx, &model.UserF{ID: &input.UserID})
	if err != nil {
		cs.logger.Error("failed checking user existence", err)
		return nil, fault.Internal("failed to create comment")
	} else if !exists {
		return nil, fault.NotFound("user not found")
	}

	media, err := cs.media.GetMedia(ctx, input.Ref, input.MediaType)
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("media not found")
		}
		cs.logger.Error("failed getting media", err)
		return nil, fault.Internal("failed to create comment")
	}
	input.Comment.MediaID = media.ID

	if input.Comment.ReplyingToID != nil {
		exists, err = cs.store.Comments().Exists(ctx, &model.CommentF{ID: input.Comment.ReplyingToID})
		if err != nil {
			cs.logger.Error("failed checking comment existence", err)
			return nil, fault.Internal("failed to create comment")
		} else if !exists {
			return nil, fault.NotFound("comment being replied to not found")
		}
	}

	comment, err := cs.store.Comments().Insert(ctx, input.Comment)
	if err != nil {
		cs.logger.Error("failed creating comment", err)
		return nil, fault.Internal("failed to create comment")
	}

	return comment, nil
}

func (cs *commentService) UpdateComment(ctx context.Context, userID, commentID uuid.UUID, content string) (*model.Comment, error) {
	comment, err := cs.store.Comments().One(ctx, &model.CommentF{ID: &commentID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("comment not found")
		}
		cs.logger.Error("failed getting comment", err)
		return nil, fault.Internal("failed to update comment")
	}

	if comment.UserID == nil || *comment.UserID != userID {
		return nil, fault.Forbidden("you are not allowed to update this comment")
	}

	comment, err = cs.store.Comments().Update(ctx, commentID, &model.CommentU{Content: &content})
	if err != nil {
		cs.logger.Error("failed updating comment", err)
		return nil, fault.Internal("failed to update comment")
	}

	return comment, nil
}

func (cs *commentService) DeleteComment(ctx context.Context, userID, commentID uuid.UUID) error {
	comment, err := cs.store.Comments().One(ctx, &model.CommentF{ID: &commentID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("comment not found")
		}
		cs.logger.Error("failed getting comment", err)
		return fault.Internal("failed to delete comment")
	}

	if comment.UserID == nil || *comment.UserID != userID {
		return fault.Forbidden("you are not allowed to delete this comment")
	}

	if err = cs.store.Comments().Delete(ctx, commentID); err != nil {
		cs.logger.Error("failed deleting comment", err)
		return fault.Internal("failed to delete comment")
	}

	return nil
}

func (cs *commentService) GetComments(ctx context.Context, ref int, mediaType model.MediaType) ([]*model.CommentWithRelationsCount, error) {
	media, err := cs.media.GetMedia(ctx, ref, mediaType)
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("media not found")
		}
		cs.logger.Error("failed getting media", err)
		return nil, fault.Internal("failed to get comments")
	}

	comments, err := cs.store.Comments().AllWithReplyAndLikeCount(ctx, media.ID)
	if err != nil {
		cs.logger.Error("failed getting comments", err)
		return nil, fault.Internal("failed to get comments")
	}

	return comments, nil
}

func (cs *commentService) GetCommentReplies(ctx context.Context, commentID uuid.UUID) ([]*model.CommentWithRelationsCount, error) {
	comment, err := cs.store.Comments().One(ctx, &model.CommentF{ID: &commentID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("comment not found")
		}
		cs.logger.Error("failed getting comment", err)
		return nil, fault.Internal("failed to get comment replies")
	}

	comments, err := cs.store.Comments().AllRepliesWithReplyAndLikeCount(ctx, comment)
	if err != nil {
		cs.logger.Error("failed getting comment replies", err)
		return nil, fault.Internal("failed to get comment replies")
	}

	return comments, nil
}

func (cs *commentService) LikeComment(ctx context.Context, like *model.Like) (*model.Like, error) {
	exists, err := cs.store.Users().Exists(ctx, &model.UserF{ID: &like.UserID})
	if err != nil {
		cs.logger.Error("failed checking user existence", err)
		return nil, fault.Internal("failed to like comment")
	} else if !exists {
		return nil, fault.NotFound("user not found")
	}

	exists, err = cs.store.Comments().Exists(ctx, &model.CommentF{ID: &like.CommentID})
	if err != nil {
		cs.logger.Error("failed checking comment existence", err)
		return nil, fault.Internal("failed to like comment")
	} else if !exists {
		return nil, fault.NotFound("comment not found")
	}

	like, err = cs.store.Likes().Insert(ctx, like)
	if err != nil {
		cs.logger.Error("failed liking comment", err)
		return nil, fault.Internal("failed to like comment")
	}

	return like, err
}

func (cs *commentService) UnlikeComment(ctx context.Context, userID, likeID uuid.UUID) error {
	like, err := cs.store.Likes().One(ctx, &model.LikeF{ID: &likeID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("like not found")
		}
		cs.logger.Error("failed getting like", err)
		return fault.Internal("failed to unlike comment")
	}

	if like.UserID != userID {
		return fault.Forbidden("you are not allowed to unlike this comment")
	}

	if err = cs.store.Likes().Delete(ctx, likeID); err != nil {
		cs.logger.Error("failed unliking comment", err)
		return fault.Internal("failed to unlike comment")
	}

	return nil
}

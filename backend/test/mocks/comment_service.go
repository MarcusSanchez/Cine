package mocks

import (
	"cine/entity/model"
	"cine/service"
	"context"
	"github.com/google/uuid"
)

var _ service.CommentService = (*CommentServiceMock)(nil)

type CommentServiceMock struct {
	CreateCommentFn     func(ctx context.Context, ref int, mediaType model.MediaType, comment *model.Comment) (*model.Comment, error)
	UpdateCommentFn     func(ctx context.Context, userID, commentID uuid.UUID, content string) (*model.Comment, error)
	DeleteCommentFn     func(ctx context.Context, userID, commentID uuid.UUID) error
	GetCommentsFn       func(ctx context.Context, ref int, mediaType model.MediaType, userID uuid.UUID) ([]*model.DetailedComment, error)
	GetCommentRepliesFn func(ctx context.Context, commentID, userID uuid.UUID) ([]*model.DetailedComment, error)
	LikeCommentFn       func(ctx context.Context, like *model.Like) (*model.Like, error)
	UnlikeCommentFn     func(ctx context.Context, userID, commentID uuid.UUID) error
}

func NewCommentService() *CommentServiceMock {
	return &CommentServiceMock{}
}

func (m *CommentServiceMock) CreateComment(ctx context.Context, ref int, mediaType model.MediaType, comment *model.Comment) (*model.Comment, error) {
	if m.CreateCommentFn != nil {
		return m.CreateCommentFn(ctx, ref, mediaType, comment)
	}
	return &model.Comment{}, nil
}

func (m *CommentServiceMock) UpdateComment(ctx context.Context, userID, commentID uuid.UUID, content string) (*model.Comment, error) {
	if m.UpdateCommentFn != nil {
		return m.UpdateCommentFn(ctx, userID, commentID, content)
	}
	return &model.Comment{}, nil
}

func (m *CommentServiceMock) DeleteComment(ctx context.Context, userID, commentID uuid.UUID) error {
	if m.DeleteCommentFn != nil {
		return m.DeleteCommentFn(ctx, userID, commentID)
	}
	return nil
}

func (m *CommentServiceMock) GetComments(ctx context.Context, ref int, mediaType model.MediaType, userID uuid.UUID) ([]*model.DetailedComment, error) {
	if m.GetCommentsFn != nil {
		return m.GetCommentsFn(ctx, ref, mediaType, userID)
	}
	return []*model.DetailedComment{}, nil
}

func (m *CommentServiceMock) GetCommentReplies(ctx context.Context, commentID, userID uuid.UUID) ([]*model.DetailedComment, error) {
	if m.GetCommentRepliesFn != nil {
		return m.GetCommentRepliesFn(ctx, commentID, userID)
	}
	return []*model.DetailedComment{}, nil
}

func (m *CommentServiceMock) LikeComment(ctx context.Context, like *model.Like) (*model.Like, error) {
	if m.LikeCommentFn != nil {
		return m.LikeCommentFn(ctx, like)
	}
	return &model.Like{}, nil
}

func (m *CommentServiceMock) UnlikeComment(ctx context.Context, userID, commentID uuid.UUID) error {
	if m.UnlikeCommentFn != nil {
		return m.UnlikeCommentFn(ctx, userID, commentID)
	}
	return nil
}

package mocks

import (
	"cine/entity/model"
	"cine/service"
	"context"
	"github.com/google/uuid"
)

var _ service.ReviewService = (*ReviewServiceMock)(nil)

type ReviewServiceMock struct {
	CreateReviewFn  func(ctx context.Context, input *service.CreateReviewInput) (*model.Review, error)
	UpdateReviewFn  func(ctx context.Context, userID, reviewID uuid.UUID, reviewU *model.ReviewU) (*model.Review, error)
	DeleteReviewFn  func(ctx context.Context, userID, reviewID uuid.UUID) error
	GetAllReviewsFn func(ctx context.Context, ref int, mediaType model.MediaType) ([]*model.DetailedReview, error)
}

func NewReviewService() *ReviewServiceMock {
	return &ReviewServiceMock{}
}

func (m *ReviewServiceMock) CreateReview(ctx context.Context, input *service.CreateReviewInput) (*model.Review, error) {
	if m.CreateReviewFn != nil {
		return m.CreateReviewFn(ctx, input)
	}
	return &model.Review{}, nil
}

func (m *ReviewServiceMock) UpdateReview(ctx context.Context, userID, reviewID uuid.UUID, reviewU *model.ReviewU) (*model.Review, error) {
	if m.UpdateReviewFn != nil {
		return m.UpdateReviewFn(ctx, userID, reviewID, reviewU)
	}
	return &model.Review{}, nil
}

func (m *ReviewServiceMock) DeleteReview(ctx context.Context, userID, reviewID uuid.UUID) error {
	if m.DeleteReviewFn != nil {
		return m.DeleteReviewFn(ctx, userID, reviewID)
	}
	return nil
}

func (m *ReviewServiceMock) GetAllReviews(ctx context.Context, ref int, mediaType model.MediaType) ([]*model.DetailedReview, error) {
	if m.GetAllReviewsFn != nil {
		return m.GetAllReviewsFn(ctx, ref, mediaType)
	}
	return []*model.DetailedReview{}, nil
}

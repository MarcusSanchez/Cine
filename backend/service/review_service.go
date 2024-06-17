package service

import (
	"cine/datastore"
	"cine/entity/model"
	"cine/pkg/fault"
	"cine/pkg/logger"
	"context"
	"github.com/google/uuid"
)

type ReviewService interface {
	CreateReview(ctx context.Context, input *CreateReviewInput) (*model.Review, error)
	UpdateReview(ctx context.Context, userID, reviewID uuid.UUID, reviewU *model.ReviewU) (*model.Review, error)
	DeleteReview(ctx context.Context, userID, reviewID uuid.UUID) error
	GetAllReviews(ctx context.Context, ref int, mediaType model.MediaType) ([]*model.DetailedReview, error)
}

type reviewService struct {
	store  datastore.Store
	logger logger.Logger
	media  MediaService
}

func NewReviewService(store datastore.Store, logger logger.Logger, media MediaService) ReviewService {
	return &reviewService{store: store, logger: logger, media: media}
}

type CreateReviewInput struct {
	UserID    uuid.UUID
	Ref       int
	MediaType model.MediaType
	Review    *model.Review
}

func (rs *reviewService) CreateReview(ctx context.Context, input *CreateReviewInput) (*model.Review, error) {
	exists, err := rs.store.Users().Exists(ctx, &model.UserF{ID: &input.UserID})
	if err != nil {
		rs.logger.Error("exists check on review failed", err)
		return nil, fault.Internal("error creating review")
	} else if !exists {
		return nil, fault.Conflict("review already exists")
	}

	media, err := rs.media.GetMedia(ctx, input.Ref, input.MediaType)
	if e, ok := fault.As(err); ok {
		if e.Code == fault.CodeNotFound {
			return nil, fault.NotFound("media not found")
		}
		rs.logger.Error("failed getting media", err)
		return nil, fault.Internal("failed to create review")
	}
	input.Review.MediaID = media.ID

	exists, err = rs.store.Reviews().Exists(ctx, &model.ReviewF{UserID: &input.UserID, MediaID: &media.ID})
	if err != nil {
		rs.logger.Error("exists check on review failed", err)
		return nil, fault.Internal("error creating review")
	} else if exists {
		return nil, fault.Conflict("a review already exists for this " + string(input.MediaType))
	}

	review, err := rs.store.Reviews().Insert(ctx, input.Review)
	if err != nil {
		rs.logger.Error("failed inserting review", err)
		return nil, fault.Internal("failed to create review")
	}

	return review, nil
}

func (rs *reviewService) UpdateReview(ctx context.Context, userID, reviewID uuid.UUID, reviewU *model.ReviewU) (*model.Review, error) {
	if !rs.hasFieldToUpdate(reviewU) {
		return nil, fault.BadRequest("no fields to update")
	}

	review, err := rs.store.Reviews().One(ctx, &model.ReviewF{ID: &reviewID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("review not found")
		}
		rs.logger.Error("failed to fetch review", err)
		return nil, fault.Internal("error updating review")
	}

	if review.UserID != userID {
		return nil, fault.Forbidden("you are not allowed to update this review")
	}

	review, err = rs.store.Reviews().Update(ctx, review.ID, reviewU)
	if err != nil {
		rs.logger.Error("failed updating review", err)
		return nil, fault.Internal("error updating review")
	}

	return review, nil
}

func (rs *reviewService) DeleteReview(ctx context.Context, userID, reviewID uuid.UUID) error {
	review, err := rs.store.Reviews().One(ctx, &model.ReviewF{ID: &reviewID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("review not found")
		}
		rs.logger.Error("failed getting review", err)
		return fault.Internal("error deleting review")
	}

	if review.UserID != userID {
		return fault.Forbidden("you are not allowed to delete this review")
	}

	if err = rs.store.Reviews().Delete(ctx, review.ID); err != nil {
		rs.logger.Error("failed deleting review", err)
		return fault.Internal("error deleting review")
	}

	return nil
}

func (rs *reviewService) GetAllReviews(ctx context.Context, ref int, mediaType model.MediaType) ([]*model.DetailedReview, error) {
	media, err := rs.media.GetMedia(ctx, ref, mediaType)
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("media not found")
		}
		rs.logger.Error("failed getting media", err)
		return nil, fault.Internal("error getting media")
	}

	reviews, err := rs.store.Reviews().AllWithUser(ctx, &model.ReviewF{MediaID: &media.ID})
	if err != nil {
		rs.logger.Error("failed getting reviews", err)
		return nil, fault.Internal("error getting reviews")
	}

	return reviews, nil
}

func (rs *reviewService) hasFieldToUpdate(reviewU *model.ReviewU) bool {
	return reviewU.Rating != nil || reviewU.Content != nil
}

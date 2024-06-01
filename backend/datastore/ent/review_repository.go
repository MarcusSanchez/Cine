package ent

import (
	"cine/datastore/ent/ent"
	"cine/datastore/ent/ent/predicate"
	Review "cine/datastore/ent/ent/review"
	User "cine/datastore/ent/ent/user"
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type reviewRepository struct {
	client *ent.Client
}

func newReviewRepository(client *ent.Client) repository.ReviewRepository {
	return &reviewRepository{client: client}
}

func (rr *reviewRepository) One(ctx context.Context, reviewFs ...*model.ReviewF) (*model.Review, error) {
	q := rr.client.Review.Query()
	q = q.Where(rr.filters(reviewFs)...)

	review, err := q.First(ctx)
	return c.review(review), c.error(err)
}

func (rr *reviewRepository) All(ctx context.Context, reviewFs ...*model.ReviewF) ([]*model.Review, error) {
	q := rr.client.Review.Query()
	q = q.Where(rr.filters(reviewFs)...)

	reviews, err := q.All(ctx)
	return c.reviews(reviews), c.error(err)
}

func (rr *reviewRepository) Exists(ctx context.Context, reviewFs ...*model.ReviewF) (bool, error) {
	q := rr.client.Review.Query()
	q = q.Where(rr.filters(reviewFs)...)

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (rr *reviewRepository) Count(ctx context.Context, reviewFs ...*model.ReviewF) (int, error) {
	q := rr.client.Review.Query()
	q = q.Where(rr.filters(reviewFs)...)

	count, err := q.Count(ctx)
	return count, c.error(err)
}

func (rr *reviewRepository) Insert(ctx context.Context, review *model.Review) (*model.Review, error) {
	i := rr.create(review)

	iReview, err := i.Save(ctx)
	return c.review(iReview), c.error(err)
}

func (rr *reviewRepository) InsertBulk(ctx context.Context, reviews []*model.Review) ([]*model.Review, error) {
	i := rr.createBulk(reviews)

	iReviews, err := i.Save(ctx)
	return c.reviews(iReviews), c.error(err)
}

func (rr *reviewRepository) Update(ctx context.Context, id uuid.UUID, reviewU *model.ReviewU) (*model.Review, error) {
	q := rr.client.Review.UpdateOneID(id)

	q.SetUpdatedAt(time.Now())
	q.SetNillableContent(reviewU.Content)
	q.SetNillableRating(reviewU.Rating)

	review, err := q.Save(ctx)
	return c.review(review), c.error(err)
}

func (rr *reviewRepository) UpdateExec(ctx context.Context, reviewU *model.ReviewU, reviewFs ...*model.ReviewF) (int, error) {
	q := rr.client.Review.Update()
	q = q.Where(rr.filters(reviewFs)...)

	q.SetUpdatedAt(time.Now())
	q.SetNillableContent(reviewU.Content)
	q.SetNillableRating(reviewU.Rating)

	affected, err := q.Save(ctx)
	return affected, c.error(err)
}

func (rr *reviewRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := rr.client.Review.DeleteOneID(id)

	err := q.Exec(ctx)
	return c.error(err)
}

func (rr *reviewRepository) DeleteExec(ctx context.Context, reviewFs ...*model.ReviewF) (int, error) {
	q := rr.client.Review.Delete()
	q = q.Where(rr.filters(reviewFs)...)

	affected, err := q.Exec(ctx)
	return affected, c.error(err)
}

func (rr *reviewRepository) AllWithUser(ctx context.Context, reviewFs ...*model.ReviewF) ([]*model.DetailedReview, error) {
	q := rr.client.Review.Query()
	q = q.Where(rr.filters(reviewFs)...).
		WithUser(func(q *ent.UserQuery) {
			q.Select(
				User.FieldID,
				User.FieldDisplayName,
				User.FieldUsername,
				User.FieldProfilePicture,
			)
		})

	reviews, err := q.All(ctx)
	return rr.detailedReviews(reviews), c.error(err)
}

func (rr *reviewRepository) filters(reviewFs []*model.ReviewF) []predicate.Review {
	var reviewF *model.ReviewF
	if len(reviewFs) > 0 {
		reviewF = reviewFs[0]
	}
	var filters []predicate.Review
	if reviewF != nil {
		if reviewF.ID != nil {
			filters = append(filters, Review.ID(*reviewF.ID))
		}
		if reviewF.UserID != nil {
			filters = append(filters, Review.UserID(*reviewF.UserID))
		}
		if reviewF.MediaID != nil {
			filters = append(filters, Review.MediaID(*reviewF.MediaID))
		}
		if reviewF.Content != nil {
			filters = append(filters, Review.Content(*reviewF.Content))
		}
		if reviewF.Rating != nil {
			filters = append(filters, Review.Rating(*reviewF.Rating))
		}
		if reviewF.CreatedAt != nil {
			filters = append(filters, Review.CreatedAt(*reviewF.CreatedAt))
		}
		if reviewF.UpdatedAt != nil {
			filters = append(filters, Review.UpdatedAt(*reviewF.UpdatedAt))
		}
	}
	return filters
}

func (rr *reviewRepository) create(review *model.Review) *ent.ReviewCreate {
	return rr.client.Review.Create().
		SetID(uuid.New()).
		SetUserID(review.UserID).
		SetMediaID(review.MediaID).
		SetContent(review.Content).
		SetRating(review.Rating).
		SetCreatedAt(time.Now())
}

func (rr *reviewRepository) createBulk(reviews []*model.Review) *ent.ReviewCreateBulk {
	builders := make([]*ent.ReviewCreate, 0, len(reviews))
	for _, review := range reviews {
		builders = append(builders, rr.create(review))
	}
	return rr.client.Review.CreateBulk(builders...)
}

func (rr *reviewRepository) detailedReviews(reviews []*ent.Review) []*model.DetailedReview {
	detailedReviews := make([]*model.DetailedReview, 0, len(reviews))
	for _, review := range reviews {
		detailedReviews = append(detailedReviews, &model.DetailedReview{
			Review: c.review(review),
			User:   c.user(review.Edges.User),
		})
	}
	return detailedReviews
}

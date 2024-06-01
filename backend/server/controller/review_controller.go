package controller

import (
	"cine/entity/model"
	"cine/entity/schemas"
	"cine/pkg/fault"
	"cine/server/middleware"
	"cine/service"
	"github.com/MarcusSanchez/go-parse"
	"github.com/MarcusSanchez/go-z"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type ReviewController struct {
	review service.ReviewService
}

func NewReviewController(review service.ReviewService) *ReviewController {
	return &ReviewController{review: review}
}

func (rc *ReviewController) Routes(router fiber.Router, mw *middleware.Middleware) {
	review := router.Group("/reviews")
	review.Post("/:mediaType/:ref", mw.SignedIn, mw.CSRF, mw.ParseMediaType("mediaType"), mw.ParseInt("ref"), rc.CreateReview)
	review.Put("/:reviewID", mw.SignedIn, mw.CSRF, mw.ParseUUID("reviewID"), rc.UpdateReview)
	review.Delete("/:reviewID", mw.SignedIn, mw.CSRF, mw.ParseUUID("reviewID"), rc.DeleteReview)
	review.Get("/:mediaType/:ref", mw.SignedIn, mw.ParseMediaType("mediaType"), mw.ParseInt("ref"), rc.GetAllReviews)
}

// CreateReview [POST] /api/reviews/:mediaType/:ref
func (rc *ReviewController) CreateReview(c *fiber.Ctx) error {

	type Payload struct {
		Content string `json:"content" z:"content"`
		Rating  int    `json:"rating"  z:"rating"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	schema := z.Struct{
		"content": schemas.ReviewContentSchema,
		"rating":  schemas.ReviewRatingSchema,
	}
	if errs := schema.Validate(p); errs != nil {
		return fault.BadRequest(errs.One())
	}

	session := c.Locals("session").(*model.Session)
	ref := c.Locals("ref").(int)
	mediaType := c.Locals("mediaType").(model.MediaType)

	review, err := rc.review.CreateReview(
		c.Context(), &service.CreateReviewInput{
			UserID:    session.UserID,
			Ref:       ref,
			MediaType: mediaType,
			Review: &model.Review{
				UserID:  session.UserID,
				Content: p.Content,
				Rating:  p.Rating,
			},
		},
	)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"review": review})
}

// UpdateReview [PUT] /api/reviews/:reviewID
func (rc *ReviewController) UpdateReview(c *fiber.Ctx) error {

	type Payload struct {
		Content *string `json:"content,optional" z:"content"`
		Rating  *int    `json:"rating,optional"  z:"rating"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	schema := z.Struct{
		"content": schemas.ReviewContentSchema.Optional(),
		"rating":  schemas.ReviewRatingSchema.Optional(),
	}
	if errs := schema.Validate(p); errs != nil {
		return fault.BadRequest(errs.One())
	}

	session := c.Locals("session").(*model.Session)
	reviewID := c.Locals("reviewID").(uuid.UUID)

	review, err := rc.review.UpdateReview(c.Context(),
		session.UserID, reviewID, &model.ReviewU{
			Content: p.Content,
			Rating:  p.Rating,
		},
	)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"review": review})
}

// DeleteReview [DELETE] /api/reviews/:reviewID
func (rc *ReviewController) DeleteReview(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	reviewID := c.Locals("reviewID").(uuid.UUID)

	err := rc.review.DeleteReview(c.Context(), session.UserID, reviewID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// GetAllReviews [GET] /api/reviews/:mediaType/:ref
func (rc *ReviewController) GetAllReviews(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)
	mediaType := c.Locals("mediaType").(model.MediaType)

	reviews, err := rc.review.GetAllReviews(c.Context(), ref, mediaType)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed_reviews": reviews})
}

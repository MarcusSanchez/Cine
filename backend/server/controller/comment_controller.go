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

type CommentController struct {
	comment service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{comment: commentService}
}

func (cc *CommentController) Routes(router fiber.Router, mw *middleware.Middleware) {
	comment := router.Group("/comments")

	comment.Get("/:mediaType/:ref", mw.SignedIn, mw.ParseInt("ref"), mw.ParseMediaType("mediaType"), cc.GetComments)
	comment.Get("/:commentID/replies", mw.SignedIn, mw.ParseUUID("commentID"), cc.GetCommentReplies)

	comment.Post("/", mw.SignedIn, mw.CSRF, cc.CreateComment)
	comment.Post("/:commentID", mw.SignedIn, mw.CSRF, mw.ParseUUID("commentID"), cc.UpdateContent)
	comment.Post("/like/:commentID", mw.SignedIn, mw.CSRF, mw.ParseUUID("commentID"), cc.LikeComment)

	comment.Delete("/:commentID", mw.SignedIn, mw.CSRF, mw.ParseUUID("commentID"), cc.DeleteComment)
	comment.Delete("/like/:likeID", mw.SignedIn, mw.CSRF, mw.ParseUUID("likeID"), cc.UnlikeComment)
}

// CreateComment [POST] /api/comments/:mediaType/:ref
func (cc *CommentController) CreateComment(c *fiber.Ctx) error {

	type Payload struct {
		Content      string  `json:"content"                 z:"content"`
		ReplyingToID *string `json:"replying_to_id,optional" z:"replying_to_id"`
		MediaType    string  `json:"media_type"              z:"media_type"`
		Ref          int     `json:"ref"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	schema := z.Struct{
		"content":        schemas.CommentContentSchema,
		"replying_to_id": schemas.CommentReplyingToIDSchema.Optional(),
		"media_type":     schemas.MediaTypeSchema,
	}
	if errs := schema.Validate(p); errs != nil {
		return fault.Validation(errs.One())
	}

	var replyingToID *uuid.UUID
	if p.ReplyingToID != nil {
		id, _ := uuid.Parse(*p.ReplyingToID)
		replyingToID = &id
	}

	session := c.Locals("session").(*model.Session)

	comment, err := cc.comment.CreateComment(c.Context(), service.CreateCommentInput{
		UserID: session.UserID,
		Comment: &model.Comment{
			Content:      p.Content,
			ReplyingToID: replyingToID,
		},
		Ref:       p.Ref,
		MediaType: model.MediaType(p.MediaType),
	})
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"comment": comment})
}

// UpdateContent [POST] /api/comments/:commenIDt
func (cc *CommentController) UpdateContent(c *fiber.Ctx) error {
	commentID := c.Locals("commentID").(uuid.UUID)

	type Payload struct {
		Content string `json:"content"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	if errs := schemas.CommentContentSchema.Validate(p.Content); errs != nil {
		return fault.Validation(errs.One())
	}

	session := c.Locals("session").(*model.Session)

	comment, err := cc.comment.UpdateComment(c.Context(), session.UserID, commentID, p.Content)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"comment": comment})
}

// DeleteComment [DELETE] /api/comments/:commenIDt
func (cc *CommentController) DeleteComment(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	commentID := c.Locals("commentID").(uuid.UUID)

	err := cc.comment.DeleteComment(c.Context(), session.UserID, commentID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// GetComments [GET] /api/comments/:mediaType/:ref
func (cc *CommentController) GetComments(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)
	mediaType := c.Locals("mediaType").(model.MediaType)

	comments, err := cc.comment.GetComments(c.Context(), ref, mediaType)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"comments": comments})
}

// GetCommentReplies [GET] /api/comments/:commentID/replies
func (cc *CommentController) GetCommentReplies(c *fiber.Ctx) error {
	commentID := c.Locals("commentID").(uuid.UUID)

	replies, err := cc.comment.GetCommentReplies(c.Context(), commentID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"replies": replies})
}

// LikeComment [POST] /api/comments/:commentID/like
func (cc *CommentController) LikeComment(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	commentID := c.Locals("commentID").(uuid.UUID)

	like, err := cc.comment.LikeComment(
		c.Context(), &model.Like{
			UserID:    session.UserID,
			CommentID: commentID,
		},
	)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"like": like})
}

// UnlikeComment [DELETE] /api/comments/:likIDe
func (cc *CommentController) UnlikeComment(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	likeID := c.Locals("likeID").(uuid.UUID)

	err := cc.comment.UnlikeComment(c.Context(), session.UserID, likeID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

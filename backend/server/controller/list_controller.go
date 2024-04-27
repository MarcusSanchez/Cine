package controller

import (
	"cine/entity/model"
	"cine/entity/schemas"
	"cine/pkg/fault"
	"cine/server/middleware"
	"cine/service"
	"github.com/MarcusSanchez/go-parse"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type ListController struct {
	list service.ListService
}

func NewListController(listService service.ListService) *ListController {
	return &ListController{list: listService}
}

func (lc *ListController) Routes(router fiber.Router, mw *middleware.Middleware) {
	list := router.Group("/list")
	list.Post("/", mw.SignedIn, mw.CSRF, lc.CreateList)
	list.Delete("/:list-id", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id"), lc.DeleteList)
	list.Put("/:list-id", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id"), lc.UpdateList)
	list.Post("/:list-id/user/:user-id", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id", "user-id"), lc.AddUserToList)
	list.Delete("/:list-id/user/:user-id", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id", "user-id"), lc.RemoveUserFromList)
	list.Get("/:list-id", mw.SignedIn, mw.ParseUUID("list-id"), lc.GetList)
	list.Get("/:list-id/detailed", mw.SignedIn, mw.ParseUUID("list-id"), lc.GetDetailedList)
	list.Post("/:list-id/movie/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id"), lc.AddMovieToList)
	list.Delete("/:list-id/movie/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id"), mw.ParseInt("ref"), lc.RemoveMovieFromList)
	list.Post("/:list-id/show/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id"), mw.ParseInt("ref"), lc.AddShowToList)
	list.Delete("/:list-id/show/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("list-id"), mw.ParseInt("ref"), lc.RemoveShowFromList)
}

// CreateList [POST] /api/list
func (lc *ListController) CreateList(c *fiber.Ctx) error {

	type Payload struct {
		Title string `json:"title"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	if errs := schemas.ListTitleSchema.Validate(p.Title); errs != nil {
		return fault.Validation(errs.One())
	}

	session := c.Locals("session").(*model.Session)

	list, err := lc.list.CreateList(c.Context(), session.UserID, p.Title)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"list": list})
}

// DeleteList [DELETE] /api/list/:list-id
func (lc *ListController) DeleteList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)

	err := lc.list.DeleteList(c.Context(), session.UserID, listID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// UpdateList [PUT] /api/list/:list-id
func (lc *ListController) UpdateList(c *fiber.Ctx) error {

	type Payload struct {
		Title  *string `json:"title,optional"`
		Public *bool   `json:"public,optional"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	if errs := schemas.ListTitleSchema.Validate(*p.Title); errs != nil {
		return fault.Validation(errs.One())
	}

	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)
	updater := &model.ListU{Title: p.Title, Public: p.Public}

	list, err := lc.list.UpdateList(c.Context(), session.UserID, listID, updater)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"list": list})
}

// AddUserToList [POST] /api/list/:list-id/user/:user-id
func (lc *ListController) AddUserToList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)
	userID := c.Locals("user-id").(uuid.UUID)

	err := lc.list.AddMemberToList(c.Context(), session.UserID, listID, userID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// RemoveUserFromList [DELETE] /api/list/:list-id/user/:user-id
func (lc *ListController) RemoveUserFromList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)
	userID := c.Locals("user-id").(uuid.UUID)

	err := lc.list.RemoveMemberFromList(c.Context(), session.UserID, listID, userID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// GetList [GET] /api/list/:list-id
func (lc *ListController) GetList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)

	list, err := lc.list.GetList(c.Context(), session.UserID, listID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"list": list})
}

// GetDetailedList [GET] /api/list/:list-id/detailed
func (lc *ListController) GetDetailedList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)

	detailedList, err := lc.list.GetDetailedList(c.Context(), session.UserID, listID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed-list": detailedList})
}

// AddMovieToList [POST] /api/list/:list-id/movie/:ref
func (lc *ListController) AddMovieToList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.AddMovieToList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// RemoveMovieFromList [DELETE] /api/list/:list-id/movie/:ref
func (lc *ListController) RemoveMovieFromList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.RemoveMovieFromList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// AddShowToList [POST] /api/list/:list-id/show/:ref
func (lc *ListController) AddShowToList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.AddShowToList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// RemoveShowFromList [DELETE] /api/list/:list-id/show/:ref
func (lc *ListController) RemoveShowFromList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("list-id").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.RemoveShowFromList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}
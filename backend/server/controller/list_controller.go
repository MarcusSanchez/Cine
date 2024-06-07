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
	list := router.Group("/lists")

	list.Get("/", mw.SignedIn, lc.GetYourLists)
	list.Get("/:userID", mw.SignedIn, mw.ParseUUID("userID"), lc.GetUsersPublicLists)
	list.Get("/:listID/detailed", mw.SignedIn, mw.ParseUUID("listID"), lc.GetDetailedList)

	list.Put("/:listID", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID"), lc.UpdateList)

	list.Post("/", mw.SignedIn, mw.CSRF, lc.CreateList)
	list.Post("/:listID/members/:userID", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID", "userID"), lc.AddMemberToList)
	list.Post("/:listID/movie/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID"), mw.ParseInt("ref"), lc.AddMovieToList)
	list.Post("/:listID/show/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID"), mw.ParseInt("ref"), lc.AddShowToList)

	list.Delete("/:listID", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID"), lc.DeleteList)
	list.Delete("/:listID/members/:userID", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID", "userID"), lc.RemoveMemberFromList)
	list.Delete("/:listID/movie/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID"), mw.ParseInt("ref"), lc.RemoveMovieFromList)
	list.Delete("/:listID/show/:ref", mw.SignedIn, mw.CSRF, mw.ParseUUID("listID"), mw.ParseInt("ref"), lc.RemoveShowFromList)
}

// CreateList [POST] /api/lists
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

// DeleteList [DELETE] /api/lists/:listID
func (lc *ListController) DeleteList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)

	err := lc.list.DeleteList(c.Context(), session.UserID, listID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// UpdateList [PUT] /api/lists/:listID
func (lc *ListController) UpdateList(c *fiber.Ctx) error {

	type Payload struct {
		Title  *string `json:"title,optional"`
		Public *bool   `json:"public,optional"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	if errs := schemas.ListTitleSchema.Optional().Validate(*p.Title); errs != nil {
		return fault.Validation(errs.One())
	}

	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)
	updater := &model.ListU{Title: p.Title, Public: p.Public}

	list, err := lc.list.UpdateList(c.Context(), session.UserID, listID, updater)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"list": list})
}

// AddMemberToList [POST] /api/lists/:listID/members/:userID
func (lc *ListController) AddMemberToList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)
	userID := c.Locals("userID").(uuid.UUID)

	err := lc.list.AddMemberToList(c.Context(), session.UserID, listID, userID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// RemoveMemberFromList [DELETE] /api/lists/:listID/members/:userID
func (lc *ListController) RemoveMemberFromList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)
	userID := c.Locals("userID").(uuid.UUID)

	err := lc.list.RemoveMemberFromList(c.Context(), session.UserID, listID, userID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// GetYourLists [GET] /api/lists/
func (lc *ListController) GetYourLists(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)

	lists, err := lc.list.GetAllLists(c.Context(), session.UserID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed_lists": lists})
}

// GetUsersPublicLists [GET] /api/lists/:userID
func (lc *ListController) GetUsersPublicLists(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	lists, err := lc.list.GetPublicLists(c.Context(), userID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed_lists": lists})
}

// GetDetailedList [GET] /api/lists/:listID/detailed
func (lc *ListController) GetDetailedList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)

	detailedList, err := lc.list.GetDetailedList(c.Context(), session.UserID, listID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed_list": detailedList})
}

// AddMovieToList [POST] /api/lists/:listID/movie/:ref
func (lc *ListController) AddMovieToList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.AddMovieToList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// RemoveMovieFromList [DELETE] /api/lists/:listID/movie/:ref
func (lc *ListController) RemoveMovieFromList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.RemoveMovieFromList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// AddShowToList [POST] /api/lists/:listID/show/:ref
func (lc *ListController) AddShowToList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.AddShowToList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

// RemoveShowFromList [DELETE] /api/lists/:listID/show/:ref
func (lc *ListController) RemoveShowFromList(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)
	listID := c.Locals("listID").(uuid.UUID)
	ref := c.Locals("ref").(int)

	err := lc.list.RemoveShowFromList(c.Context(), session.UserID, listID, ref)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

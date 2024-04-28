package controller

import (
	"cine/server/middleware"
	"cine/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type MediaController struct {
	media service.MediaService
}

func NewMediaController(mediaService service.MediaService) *MediaController {
	return &MediaController{media: mediaService}
}

func (mc *MediaController) Routes(router fiber.Router, mw *middleware.Middleware) {
	media := router.Group("/medias")
	media.Get("/movie/:ref", mw.SignedIn, mw.ParseInt("ref"), mc.GetMovie)
	media.Get("/show/:ref", mw.SignedIn, mw.ParseInt("ref"), mc.GetShow)
}

// GetMovie [Get] /medias/movie/:ref
func (mc *MediaController) GetMovie(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)

	movie, err := mc.media.GetDetailedMovie(c.Context(), ref)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed-movie": movie})
}

// GetShow [Get] /medias/show/:ref
func (mc *MediaController) GetShow(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)

	show, err := mc.media.GetDetailedShow(c.Context(), ref)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed-show": show})
}

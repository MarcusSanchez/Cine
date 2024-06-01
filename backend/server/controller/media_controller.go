package controller

import (
	"cine/pkg/tmdb"
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
	media.Get("/movie/:ref/credits", mw.SignedIn, mw.ParseInt("ref"), mc.GetMovieCredits)
	media.Get("/show/:ref/credits", mw.SignedIn, mw.ParseInt("ref"), mc.GetShowCredits)
	media.Get("/show/:ref/season/:season", mw.SignedIn, mw.ParseInt("ref"), mw.ParseInt("season"), mc.GetShowDetailedSeason)
	media.Get("/movie/:list", mw.SignedIn, mw.ParseInt("ref"), mw.ParseMovieList("list"), mc.GetMovieList)
	media.Get("/show/:list", mw.SignedIn, mw.ParseInt("ref"), mw.ParseShowList("list"), mc.GetShowList)
}

// GetMovie [Get] /api/medias/movie/:ref
func (mc *MediaController) GetMovie(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)

	movie, err := mc.media.GetDetailedMovie(c.Context(), ref)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed_movie": movie})
}

// GetShow [Get] /api/medias/show/:ref
func (mc *MediaController) GetShow(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)

	show, err := mc.media.GetDetailedShow(c.Context(), ref)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed_show": show})
}

// GetMovieCredits [Get] /api/medias/movie/:ref/credits
func (mc *MediaController) GetMovieCredits(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)

	credits, err := mc.media.GetMovieCredits(c.Context(), ref)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"movie_credits": credits})
}

// GetShowCredits [Get] /api/medias/show/:ref/credits
func (mc *MediaController) GetShowCredits(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)

	credits, err := mc.media.GetShowCredits(c.Context(), ref)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"show_credits": credits})
}

// GetShowDetailedSeason [Get] /api/medias/show/:ref/season/:season
func (mc *MediaController) GetShowDetailedSeason(c *fiber.Ctx) error {
	ref := c.Locals("ref").(int)
	season := c.Locals("season").(int)

	detailedSeason, err := mc.media.GetShowDetailedSeason(c.Context(), ref, season)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"detailed_season": detailedSeason})
}

// GetMovieList [Get] /api/medias/movie/:list
func (mc *MediaController) GetMovieList(c *fiber.Ctx) error {
	list := c.Locals("list").(tmdb.MovieList)

	movies, err := mc.media.GetMovieList(c.Context(), list)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"movies": movies})
}

// GetShowList [Get] /api/medias/show/:list
func (mc *MediaController) GetShowList(c *fiber.Ctx) error {
	list := c.Locals("list").(tmdb.ShowList)

	shows, err := mc.media.GetShowList(c.Context(), list)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"shows": shows})
}

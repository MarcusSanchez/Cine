package middleware

import (
	"cine/entity/model"
	"cine/entity/schemas"
	"cine/pkg/fault"
	"cine/pkg/tmdb"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (m *Middleware) ParseUUID(keys ...string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		for _, key := range keys {
			id, err := uuid.Parse(c.Params(key))
			if err != nil {
				return fault.BadRequest(key + " must be a valid UUID")
			}
			c.Locals(key, id)
		}
		return c.Next()
	}
}

func (m *Middleware) ParseInt(keys ...string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		for _, key := range keys {
			id, err := c.ParamsInt(key)
			if err != nil {
				return fault.BadRequest(key + " must be a valid integer")
			}
			c.Locals(key, id)
		}
		return c.Next()
	}
}

func (m *Middleware) ParseMediaType(key string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		mediaType := c.Params(key)
		if errs := schemas.MediaTypeSchema.Validate(mediaType); errs != nil {
			return fault.BadRequest(key + " must be a valid media type")
		}
		c.Locals(key, model.MediaType(mediaType))
		return c.Next()
	}
}

func (m *Middleware) ParseMovieList(key string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		list := c.Params(key)
		if errs := schemas.MovieListSchema.Validate(list); errs != nil {
			return fault.BadRequest(key + " must be a valid movie list")
		}
		c.Locals(key, tmdb.MovieList(list))
		return c.Next()
	}
}

func (m *Middleware) ParseShowList(key string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		list := c.Params(key)
		if errs := schemas.ShowListSchema.Validate(list); errs != nil {
			return fault.BadRequest(key + " must be a valid show list")
		}
		c.Locals(key, tmdb.ShowList(list))
		return c.Next()
	}
}

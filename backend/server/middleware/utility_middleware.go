package middleware

import (
	"cine/pkg/fault"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (m *Middleware) ParseUUID(key string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params(key))
		if err != nil {
			return fault.BadRequest(key + " must be a valid UUID")
		}
		c.Locals(key, id)
		return c.Next()
	}
}

func (m *Middleware) ParseInt(key string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt(key)
		if err != nil {
			return fault.BadRequest(key + " must be a valid integer")
		}
		c.Locals(key, id)
		return c.Next()
	}
}

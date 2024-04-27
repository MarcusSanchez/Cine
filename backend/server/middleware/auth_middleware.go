package middleware

import (
	"cine/entity/model"
	"cine/pkg/fault"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// SignedOut ensures the user is signed out
func (m *Middleware) SignedOut(c *fiber.Ctx) error {
	value := c.Get("X-Session-Token")
	if value == "" {
		return c.Next()
	}

	token, err := uuid.Parse(value)
	if err != nil {
		return c.Next()
	}

	_, err = m.auth.Session(c.Context(), token)
	if e, ok := fault.As(err); ok {
		if e.Code == fault.CodeNotFound || e.Code == fault.CodeUnauthorized {
			return c.Next()
		}
		return e
	}

	return fault.Forbidden("must be signed out")
}

// SignedIn ensures the user is signed in
func (m *Middleware) SignedIn(c *fiber.Ctx) error {
	value := c.Get("X-Session-Token")
	if value == "" {
		return fault.Unauthorized("missing access token")
	}

	access, err := uuid.Parse(value)
	if err != nil {
		return fault.Unauthorized("invalid access token")
	}

	session, err := m.auth.Session(c.Context(), access)
	if err != nil {
		return err
	}

	c.Locals("session", session)
	return c.Next()
}

// CSRF checks the csrf token against the session
func (m *Middleware) CSRF(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)

	value := c.Get("X-CSRF-Token")
	if value == "" {
		return fault.BadRequest("missing csrf token")
	}

	csrf, err := uuid.Parse(value)
	if err != nil {
		return fault.BadRequest("invalid csrf token")
	}

	if session.CSRF != csrf {
		m.logger.Warn("potential csrf attack", "token mismatch for session "+session.ID.String())
		return fault.Forbidden("csrf token mismatch")
	}

	return c.Next()
}

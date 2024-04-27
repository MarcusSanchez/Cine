package controller

import (
	"cine/server/middleware"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Routes(router fiber.Router, mw *middleware.Middleware)
}

type Controllers []Controller

func NewControllers(
	userController *UserController,
	authController *AuthController,
	listController *ListController,
) Controllers {
	return Controllers{
		userController,
		authController,
		listController,
	}
}

func (cs Controllers) Register(router fiber.Router, mw *middleware.Middleware) {
	for _, c := range cs {
		c.Routes(router, mw)
	}
}
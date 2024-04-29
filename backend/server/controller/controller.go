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
	reviewController *ReviewController,
	commentController *CommentController,
	mediaController *MediaController,
) Controllers {
	return Controllers{
		userController,
		authController,
		listController,
		reviewController,
		commentController,
		mediaController,
	}
}

func (cs Controllers) Register(router fiber.Router, mw *middleware.Middleware) {
	for _, c := range cs {
		c.Routes(router, mw)
	}
}

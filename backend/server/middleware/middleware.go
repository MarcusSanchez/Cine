package middleware

import (
	"cine/config"
	"cine/pkg/logger"
	"cine/service"
	"github.com/gofiber/fiber/v2"
	log "github.com/gofiber/fiber/v2/middleware/logger"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
)

type Middleware struct {
	config *config.Config
	logger logger.Logger
	auth   service.AuthService
}

func NewMiddleware(
	config *config.Config,
	logger logger.Logger,
	authService service.AuthService,
) *Middleware {
	return &Middleware{
		config: config,
		logger: logger,
		auth:   authService,
	}
}

func (m *Middleware) All(app *fiber.App) {
	app.Use(log.New())
	app.Use(recovery.New())
}

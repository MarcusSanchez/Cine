package server

import (
	"cine/config"
	"cine/pkg/fault"
	"cine/pkg/logger"
	"cine/server/controller"
	"cine/server/middleware"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func InvokeServer(
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	controllers controller.Controllers,
	middleware *middleware.Middleware,
	logger logger.Logger,
	config *config.Config,
) {
	fc := fiber.Config{ErrorHandler: errorHandler}
	server := fiber.New(fc)
	middleware.All(server)

	router := server.Group("/api")
	controllers.Register(router, middleware)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := server.Listen(":" + config.Port)
				if err != nil {
					logger.Error("failed to listen", err)
					_ = shutdowner.Shutdown()
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return server.Shutdown()
		},
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	if e, ok := fault.As(err); ok {
		return c.Status(e.Code.Status()).JSON(fiber.Map{
			"error":   e.Code.String(),
			"message": e.Message,
		})
	}
	return fiber.DefaultErrorHandler(c, err)
}

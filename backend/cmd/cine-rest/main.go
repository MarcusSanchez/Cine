package main

import (
	"cine/config"
	"cine/datastore/ent"
	"cine/pkg/logger"
	"cine/pkg/tmdb"
	"cine/server"
	"cine/server/controller"
	"cine/server/middleware"
	"cine/service"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			logger.NewLogger,
			config.NewConfig,

			ent.NewStore,
			tmdb.NewTheMovieDatabaseAPI,

			service.NewAuthService,
			service.NewUserService,
			service.NewMediaService,
			service.NewListService,
			service.NewCommentService,
			service.NewReviewService,

			middleware.NewMiddleware,

			controller.NewUserController,
			controller.NewAuthController,
			controller.NewListController,
			controller.NewCommentController,
			controller.NewReviewController,
			controller.NewMediaController,
			controller.NewControllers,
		),
		fx.Invoke(
			server.InvokeServer,
		),
	).Run()
}

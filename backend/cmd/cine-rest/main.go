package main

import (
	"cine/config"
	"cine/datastore/ent"
	"cine/pkg/logger"
	"cine/pkg/tmdb"
	"cine/server/controller"
	"cine/server/middleware"
	"cine/service"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"log"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(
			logger.NewLogger,
			config.NewConfig,

			ent.NewStore,
			tmdb.NewTheMovieDatabaseAPI,

			service.NewAuthService,
			service.NewUserService,

			middleware.NewMiddleware,

			controller.NewUserController,
			controller.NewAuthController,
			controller.NewControllers,
		),
		fx.Invoke(
			//server.InvokeServer,
			invokeAPI,
		),
	).Run()
}

func invokeAPI(api tmdb.API) {

	show, err := api.SearchShowByRef(550)
	if err != nil {
		if tmdb.IsNotFound(err) {
			log.Println("Show not found")
			return
		}
		log.Fatal(err)
	}

	fmt.Println(show.Name)
}

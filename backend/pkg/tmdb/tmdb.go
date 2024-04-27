// Package tmdb handles all the API calls to 'The Movie Database' API.
package tmdb

import (
	"cine/config"
	"cine/pkg/logger"
	"github.com/go-resty/resty/v2"
)

const url = "https://api.themoviedb.org/3"

type API interface {
	searchAPI
	listMovieAPI
	listShowAPI
}

type api struct {
	key       string
	readToken string
	logger    logger.Logger
	client    *resty.Client
}

func NewTheMovieDatabaseAPI(config *config.Config, logger logger.Logger) API {
	return &api{
		key:       config.TMDBApiKey,
		readToken: config.TMDBReadToken,
		logger:    logger,
		client:    resty.New(),
	}
}

package config

import (
	"cine/pkg/logger"
	"errors"
	"github.com/MarcusSanchez/go-z"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"os"
)

type Config struct {
	Port          string `z:"port"`
	Datasource    string `z:"datasource"`
	Environment   string `z:"environment"`
	TMDBApiKey    string `z:"tmdb_api"`
	TMDBReadToken string `z:"tmdb_read_token"`
}

func NewConfig(shutdowner fx.Shutdowner, logger logger.Logger) *Config {
	// load .env file if environment is not already set
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load("./config/.env"); err != nil {
			logger.Error("failed to load .env file", err)
			_ = shutdowner.Shutdown()
		}
	}

	cfg := &Config{
		Port:          os.Getenv("PORT"),
		Datasource:    os.Getenv("DATASOURCE"),
		Environment:   os.Getenv("ENVIRONMENT"),
		TMDBApiKey:    os.Getenv("THE_MOVIE_DATABASE_API_KEY"),
		TMDBReadToken: os.Getenv("THE_MOVIE_DATABASE_READ_TOKEN"),
	}

	if errs := cfg.validate(); errs != nil {
		for _, err := range errs.All() {
			logger.Error("failed to validate config", errors.New(err))
		}
		_ = shutdowner.Shutdown()
	}

	return cfg
}

func (c *Config) validate() z.Errors {
	schema := z.Struct{
		"port": z.String().
			NotEmpty("port must be set").
			Min(4, "port must be at least 4 digits long").
			Max(5, "port must be at most 5 digits long").
			Regex(`^\d+$`, "port must be a number"),
		"environment": z.String().
			NotEmpty("environment must be set").
			In([]string{"development", "production"}, "environment must be either development or production"),
		"datasource": z.String().
			NotEmpty("datasource must be set"),
		"tmdb_api": z.String().
			NotEmpty("tmdb_api must be set"),
		"tmdb_read_token": z.String().
			NotEmpty("tmdb_read_token must be set"),
	}
	return schema.Validate(c)
}

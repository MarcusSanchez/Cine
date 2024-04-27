package tmdb

import (
	"fmt"
	"github.com/MarcusSanchez/go-parse"
)

type listMovieAPI interface {
	ListNowPlayingMovies() ([]Movie, error)
	ListPopularMovies() ([]Movie, error)
	ListTopRatedMovies() ([]Movie, error)
	ListUpcomingMovies() ([]Movie, error)
}

func (a *api) ListNowPlayingMovies() ([]Movie, error) {
	endpoint := "/movie/now_playing"
	return a.movieListRequest(endpoint, "now-playing")
}

func (a *api) ListPopularMovies() ([]Movie, error) {
	endpoint := "/movie/popular"
	return a.movieListRequest(endpoint, "popular")
}

func (a *api) ListTopRatedMovies() ([]Movie, error) {
	endpoint := "/movie/top_rated"
	return a.movieListRequest(endpoint, "top-rated")
}

func (a *api) ListUpcomingMovies() ([]Movie, error) {
	endpoint := "/movie/upcoming"
	return a.movieListRequest(endpoint, "upcoming")
}

func (a *api) movieListRequest(endpoint, listName string) ([]Movie, error) {
	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch '"+listName+"' movie-list", err)
		return nil, ErrorInternal("failed to fetch '" + listName + "' movie-list")
	}

	if !resp.IsSuccess() {
		a.logger.Warn(
			"fetched '"+listName+"' movie-list response was not successful",
			fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
		)
		return nil, ErrorInternal("failed to fetch '" + listName + "' movie-list")
	}

	type Result struct {
		Movies []Movie `json:"results"`
	}

	r, err := parse.JSON[Result](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse '"+listName+"' movie-list", err)
		return nil, ErrorInternal("failed to fetch '" + listName + "' movie-list")
	}

	return r.Movies, nil
}

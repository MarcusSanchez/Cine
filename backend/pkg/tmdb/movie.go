package tmdb

import (
	"fmt"
	"github.com/MarcusSanchez/go-parse"
	"net/http"
	"strconv"
)

type movieAPI interface {
	GetMovie(ref int) (*DetailedMovie, error)
	GetMovieCredits(ref int) (*MovieCredits, error)

	ListNowPlayingMovies() ([]Movie, error)
	ListPopularMovies() ([]Movie, error)
	ListTopRatedMovies() ([]Movie, error)
	ListUpcomingMovies() ([]Movie, error)
}

func (a *api) GetMovie(ref int) (*DetailedMovie, error) {
	endpoint := "/movie/" + strconv.Itoa(ref)

	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch movie by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch movie by ref: " + strconv.Itoa(ref))
	}

	if !resp.IsSuccess() {
		switch resp.StatusCode() {
		case http.StatusNotFound:
			return nil, ErrorNotFound("movie by ref: " + strconv.Itoa(ref))
		default:
			a.logger.Warn(
				"movie by ref '"+strconv.Itoa(ref)+"' response was not successful",
				fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
			)
			return nil, ErrorInternal("failed to fetch movie by ref: " + strconv.Itoa(ref))
		}
	}

	movie, err := parse.JSON[DetailedMovie](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse movie by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch movie by ref: " + strconv.Itoa(ref))
	}

	return movie, nil
}

func (a *api) GetMovieCredits(ref int) (*MovieCredits, error) {
	endpoint := "/movie/" + strconv.Itoa(ref) + "/credits"

	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch movie credits by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch movie credits by ref: " + strconv.Itoa(ref))
	}

	if !resp.IsSuccess() {
		switch resp.StatusCode() {
		case http.StatusNotFound:
			return nil, ErrorNotFound("movie credits by ref: " + strconv.Itoa(ref))
		default:
			a.logger.Warn(
				"movie credits by ref '"+strconv.Itoa(ref)+"' response was not successful",
				fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
			)
			return nil, ErrorInternal("failed to fetch movie credits by ref: " + strconv.Itoa(ref))
		}
	}

	credits, err := parse.JSON[MovieCredits](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse movie credits by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch movie credits by ref: " + strconv.Itoa(ref))
	}

	return credits, nil
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

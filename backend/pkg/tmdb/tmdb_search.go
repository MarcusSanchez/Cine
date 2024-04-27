package tmdb

import (
	"fmt"
	"github.com/MarcusSanchez/go-parse"
	"net/http"
	"strconv"
)

type searchAPI interface {
	SearchMovies(query string, filter ...SearchMovieFilter) ([]Movie, error)
	SearchShows(query string, filter ...SearchShowFilter) ([]Show, error)
	SearchMovieByRef(ref int) (*DetailedMovie, error)
	SearchShowByRef(ref int) (*DetailedShow, error)
}

type SearchMovieFilter struct {
	Language           *string
	Region             *string
	Year               *int
	PrimaryReleaseYear *int
	Page               *int
}

func (a *api) SearchMovies(query string, filters ...SearchMovieFilter) ([]Movie, error) {
	endpoint := "/search/movie"

	request := a.client.R().
		SetHeader("Authorization", "Bearer "+a.readToken).
		SetHeader("Accept", "application/json").
		SetQueryParam("query", query)

	if len(filters) > 0 {
		filter := filters[0]
		if filter.Language != nil {
			request.SetQueryParam("language", *filter.Language)
		}
		if filter.Region != nil {
			request.SetQueryParam("region", *filter.Region)
		}
		if filter.Year != nil {
			request.SetQueryParam("year", strconv.Itoa(*filter.Year))
		}
		if filter.PrimaryReleaseYear != nil {
			request.SetQueryParam("primary_release_year", strconv.Itoa(*filter.PrimaryReleaseYear))
		}
		if filter.Page != nil {
			request.SetQueryParam("page", strconv.Itoa(*filter.Page))
		}
	}

	resp, err := request.Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch movies for query: "+query, err)
		return nil, ErrorInternal("failed to fetch movies for query: " + query)
	}

	if !resp.IsSuccess() {
		a.logger.Warn(
			"movie search query '"+query+"' response was not successful",
			fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
		)
		return nil, ErrorInternal("failed to fetch movies for query: " + query)
	}

	type Result struct {
		Movies []Movie `json:"results"`
	}

	r, err := parse.JSON[Result](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse movies for query: "+query, err)
		return nil, ErrorInternal("failed to fetch movies for query: " + query)
	}

	return r.Movies, nil
}

type SearchShowFilter struct {
	Language         *string
	Year             *int
	FirstAirDateYear *int
	Page             *int
}

func (a *api) SearchShows(query string, filters ...SearchShowFilter) ([]Show, error) {
	endpoint := "/search/tv"

	request := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		SetQueryParam("query", query)

	if len(filters) > 0 {
		filter := filters[0]
		if filter.Language != nil {
			request.SetQueryParam("language", *filter.Language)
		}
		if filter.Year != nil {
			request.SetQueryParam("year", strconv.Itoa(*filter.Year))
		}
		if filter.FirstAirDateYear != nil {
			request.SetQueryParam("first_air_date_year", strconv.Itoa(*filter.FirstAirDateYear))
		}
		if filter.Page != nil {
			request.SetQueryParam("page", strconv.Itoa(*filter.Page))
		}
	}

	resp, err := request.Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch shows for query: "+query, err)
		return nil, ErrorInternal("failed to fetch shows for query: " + query)
	}

	if resp.StatusCode() != 200 {
		a.logger.Warn(
			"show search query '"+query+"' response was not successful",
			fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
		)
		return nil, ErrorInternal("failed to fetch shows for query: " + query)
	}

	type Result struct {
		Shows []Show `json:"results"`
	}

	r, err := parse.JSON[Result](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse shows for query: "+query, err)
		return nil, ErrorInternal("failed to fetch shows for query: " + query)
	}

	return r.Shows, nil
}

func (a *api) SearchMovieByRef(ref int) (*DetailedMovie, error) {
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

func (a *api) SearchShowByRef(ref int) (*DetailedShow, error) {
	endpoint := "/tv/" + strconv.Itoa(ref)

	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch show by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch show by ref: " + strconv.Itoa(ref))
	}

	if !resp.IsSuccess() {
		switch resp.StatusCode() {
		case http.StatusNotFound:
			return nil, ErrorNotFound("show by ref: " + strconv.Itoa(ref))
		default:
			a.logger.Warn(
				"show by ref '"+strconv.Itoa(ref)+"' response was not successful",
				fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
			)
			return nil, ErrorInternal("failed to fetch show by ref: " + strconv.Itoa(ref))
		}
	}

	show, err := parse.JSON[DetailedShow](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse show by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch show by ref: " + strconv.Itoa(ref))
	}

	return show, nil
}

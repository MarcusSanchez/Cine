package tmdb

import (
	"fmt"
	"github.com/MarcusSanchez/go-parse"
	"strconv"
)

type searchAPI interface {
	SearchMovies(query string, filter ...SearchMovieFilter) ([]Movie, error)
	SearchShows(query string, filter ...SearchShowFilter) ([]Show, error)
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

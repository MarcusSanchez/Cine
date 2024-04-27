package tmdb

import (
	"fmt"
	"github.com/MarcusSanchez/go-parse"
)

type listShowAPI interface {
	ListAiringTodayShows() ([]Show, error)
	ListPopularShows() ([]Show, error)
	ListTopRatedShows() ([]Show, error)
	ListOnTheAirShows() ([]Show, error)
}

func (a *api) ListAiringTodayShows() ([]Show, error) {
	return a.showListRequest("/tv/airing_today", "airing-today")
}

func (a *api) ListPopularShows() ([]Show, error) {
	return a.showListRequest("/tv/popular", "popular")
}

func (a *api) ListTopRatedShows() ([]Show, error) {
	return a.showListRequest("/tv/top_rated", "top-rated")
}

func (a *api) ListOnTheAirShows() ([]Show, error) {
	return a.showListRequest("/tv/on_the_air", "on-the-air")
}

func (a *api) showListRequest(endpoint, listName string) ([]Show, error) {
	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch '"+listName+"' show-list", err)
		return nil, ErrorInternal("failed to fetch '" + listName + "' show-list")
	}

	if !resp.IsSuccess() {
		a.logger.Warn(
			"fetched '"+listName+"' show-list response was not successful",
			fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
		)
		return nil, ErrorInternal("failed to fetch '" + listName + "' show-list")
	}

	type Result struct {
		Shows []Show `json:"results"`
	}

	r, err := parse.JSON[Result](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse '"+listName+"' show-list", err)
		return nil, ErrorInternal("failed to parse '" + listName + "' show-list")
	}

	return r.Shows, nil
}

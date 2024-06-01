package tmdb

import (
	"fmt"
	"github.com/MarcusSanchez/go-parse"
	"net/http"
	"strconv"
)

type showAPI interface {
	GetShow(ref int) (*DetailedShow, error)
	GetShowCredits(ref int) (*ShowCredits, error)
	GetShowSeasonDetails(ref int, seasonNumber int) (*DetailedSeason, error)

	ListAiringTodayShows() ([]Show, error)
	ListPopularShows() ([]Show, error)
	ListTopRatedShows() ([]Show, error)
	ListOnTheAirShows() ([]Show, error)
}

func (a *api) GetShow(ref int) (*DetailedShow, error) {
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

func (a *api) GetShowCredits(ref int) (*ShowCredits, error) {
	endpoint := "/tv/" + strconv.Itoa(ref) + "/aggregate_credits"

	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch show credits by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch show credits by ref: " + strconv.Itoa(ref))
	}

	if !resp.IsSuccess() {
		switch resp.StatusCode() {
		case http.StatusNotFound:
			return nil, ErrorNotFound("show credits by ref: " + strconv.Itoa(ref))
		default:
			a.logger.Warn(
				"show credits by ref '"+strconv.Itoa(ref)+"' response was not successful",
				fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
			)
			return nil, ErrorInternal("failed to fetch show credits by ref: " + strconv.Itoa(ref))
		}
	}

	credits, err := parse.JSON[ShowCredits](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse show credits by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch movie credits by ref: " + strconv.Itoa(ref))
	}

	return credits, nil
}

func (a *api) GetShowSeasonDetails(ref int, seasonNumber int) (*DetailedSeason, error) {
	endpoint := "/tv/" + strconv.Itoa(ref) + "/season/" + strconv.Itoa(seasonNumber)

	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+a.readToken).
		Get(url + endpoint)
	if err != nil {
		a.logger.Error("failed to fetch season details", err)
		return nil, ErrorInternal("failed to fetch season details")
	}

	if !resp.IsSuccess() {
		switch resp.StatusCode() {
		case http.StatusNotFound:
			return nil, ErrorNotFound("show ref or season number not found")
		default:
			a.logger.Warn(
				"season details by ref '"+strconv.Itoa(ref)+"' response was not successful",
				fmt.Sprintf("status: %d | body: %s", resp.StatusCode(), prettyJSON(resp.Body())),
			)
			return nil, ErrorInternal("failed to fetch season details")
		}
	}

	detailedSeason, err := parse.JSON[DetailedSeason](resp.Body())
	if err != nil {
		a.logger.Error("failed to parse show credits by ref: "+strconv.Itoa(ref), err)
		return nil, ErrorInternal("failed to fetch movie credits by ref: " + strconv.Itoa(ref))
	}

	return detailedSeason, nil
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

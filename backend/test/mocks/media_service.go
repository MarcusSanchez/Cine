package mocks

import (
	"cine/entity/model"
	"cine/pkg/tmdb"
	"cine/service"
	"context"
	"github.com/google/uuid"
)

var _ service.MediaService = (*MediaServiceMock)(nil)

type MediaServiceMock struct {
	CreateMediaFn           func(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error)
	GetMediaFn              func(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error)
	GetMediaByIDFn          func(ctx context.Context, id uuid.UUID) (*model.Media, error)
	GetDetailedMovieFn      func(ctx context.Context, ref int) (*tmdb.DetailedMovie, error)
	GetDetailedShowFn       func(ctx context.Context, ref int) (*tmdb.DetailedShow, error)
	GetMovieCreditsFn       func(ctx context.Context, ref int) (*tmdb.MovieCredits, error)
	GetShowCreditsFn        func(ctx context.Context, ref int) (*tmdb.ShowCredits, error)
	GetShowDetailedSeasonFn func(ctx context.Context, ref int, seasonNumber int) (*tmdb.DetailedSeason, error)
	GetMovieListFn          func(ctx context.Context, list tmdb.MovieList) ([]tmdb.Movie, error)
	GetShowListFn           func(ctx context.Context, list tmdb.ShowList) ([]tmdb.Show, error)
	SearchMoviesFn          func(ctx context.Context, query string, page int) ([]tmdb.Movie, error)
	SearchShowsFn           func(ctx context.Context, query string, page int) ([]tmdb.Show, error)
}

func NewMediaService() *MediaServiceMock {
	return &MediaServiceMock{}
}

func (m *MediaServiceMock) CreateMedia(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error) {
	if m.CreateMediaFn != nil {
		return m.CreateMediaFn(ctx, ref, mediaType)
	}
	return &model.Media{}, nil
}

func (m *MediaServiceMock) GetMedia(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error) {
	if m.GetMediaFn != nil {
		return m.GetMediaFn(ctx, ref, mediaType)
	}
	return &model.Media{}, nil
}

func (m *MediaServiceMock) GetMediaByID(ctx context.Context, id uuid.UUID) (*model.Media, error) {
	if m.GetMediaByIDFn != nil {
		return m.GetMediaByIDFn(ctx, id)
	}
	return &model.Media{}, nil
}

func (m *MediaServiceMock) GetDetailedMovie(ctx context.Context, ref int) (*tmdb.DetailedMovie, error) {
	if m.GetDetailedMovieFn != nil {
		return m.GetDetailedMovieFn(ctx, ref)
	}
	return &tmdb.DetailedMovie{}, nil
}

func (m *MediaServiceMock) GetDetailedShow(ctx context.Context, ref int) (*tmdb.DetailedShow, error) {
	if m.GetDetailedShowFn != nil {
		return m.GetDetailedShowFn(ctx, ref)
	}
	return &tmdb.DetailedShow{}, nil
}

func (m *MediaServiceMock) GetMovieCredits(ctx context.Context, ref int) (*tmdb.MovieCredits, error) {
	if m.GetMovieCreditsFn != nil {
		return m.GetMovieCreditsFn(ctx, ref)
	}
	return &tmdb.MovieCredits{}, nil
}

func (m *MediaServiceMock) GetShowCredits(ctx context.Context, ref int) (*tmdb.ShowCredits, error) {
	if m.GetShowCreditsFn != nil {
		return m.GetShowCreditsFn(ctx, ref)
	}
	return &tmdb.ShowCredits{}, nil
}

func (m *MediaServiceMock) GetShowDetailedSeason(ctx context.Context, ref int, seasonNumber int) (*tmdb.DetailedSeason, error) {
	if m.GetShowDetailedSeasonFn != nil {
		return m.GetShowDetailedSeasonFn(ctx, ref, seasonNumber)
	}
	return &tmdb.DetailedSeason{}, nil
}

func (m *MediaServiceMock) GetMovieList(ctx context.Context, list tmdb.MovieList) ([]tmdb.Movie, error) {
	if m.GetMovieListFn != nil {
		return m.GetMovieListFn(ctx, list)
	}
	return []tmdb.Movie{}, nil
}

func (m *MediaServiceMock) GetShowList(ctx context.Context, list tmdb.ShowList) ([]tmdb.Show, error) {
	if m.GetShowListFn != nil {
		return m.GetShowListFn(ctx, list)
	}
	return []tmdb.Show{}, nil
}

func (m *MediaServiceMock) SearchMovies(ctx context.Context, query string, page int) ([]tmdb.Movie, error) {
	if m.SearchMoviesFn != nil {
		return m.SearchMoviesFn(ctx, query, page)
	}
	return []tmdb.Movie{}, nil
}

func (m *MediaServiceMock) SearchShows(ctx context.Context, query string, page int) ([]tmdb.Show, error) {
	if m.SearchShowsFn != nil {
		return m.SearchShowsFn(ctx, query, page)
	}
	return []tmdb.Show{}, nil
}

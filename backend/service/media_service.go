package service

import (
	"cine/datastore"
	"cine/entity/model"
	"cine/pkg/fault"
	"cine/pkg/logger"
	"cine/pkg/tmdb"
	"context"
	"github.com/google/uuid"
)

type MediaService interface {
	CreateMedia(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error)

	GetMedia(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error)
	GetMediaByID(ctx context.Context, id uuid.UUID) (*model.Media, error)

	GetDetailedMovie(ctx context.Context, ref int) (*tmdb.DetailedMovie, error)
	GetDetailedShow(ctx context.Context, ref int) (*tmdb.DetailedShow, error)

	GetMovieCredits(ctx context.Context, ref int) (*tmdb.MovieCredits, error)
	GetShowCredits(ctx context.Context, ref int) (*tmdb.ShowCredits, error)

	GetShowDetailedSeason(ctx context.Context, ref int, seasonNumber int) (*tmdb.DetailedSeason, error)

	GetMovieList(ctx context.Context, list tmdb.MovieList) ([]tmdb.Movie, error)
	GetShowList(ctx context.Context, list tmdb.ShowList) ([]tmdb.Show, error)
}

type mediaService struct {
	logger logger.Logger
	tmdb   tmdb.API
	store  datastore.Store
}

func NewMediaService(
	logger logger.Logger,
	tmdbAPI tmdb.API,
	store datastore.Store,
) MediaService {
	return &mediaService{
		logger: logger,
		tmdb:   tmdbAPI,
		store:  store,
	}
}

func (ms *mediaService) CreateMedia(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error) {
	exists, err := ms.store.Medias().Exists(ctx, &model.MediaF{Ref: &ref, MediaType: &mediaType})
	if err != nil {
		ms.logger.Error("exists check on media failed", err)
		return nil, fault.Internal("error creating media")
	} else if exists {
		return nil, fault.Conflict("media already exists")
	}

	media, err := ms.mediaFromRef(ref, mediaType)
	if err != nil {
		return nil, err
	}

	media, err = ms.store.Medias().Insert(ctx, media)
	if err != nil {
		ms.logger.Error("media insert failed", err)
		return nil, fault.Internal("error creating media")
	}

	return media, nil
}

func (ms *mediaService) GetMedia(ctx context.Context, ref int, mediaType model.MediaType) (*model.Media, error) {
	media, err := ms.store.Medias().One(ctx, &model.MediaF{Ref: &ref, MediaType: &mediaType})
	if err != nil {
		if datastore.IsNotFound(err) {
			// if we don't find the media in our database, it doesn't mean that it doesn't exist in TMDB,
			// so we need to check if it exists there, and if so, create a record for it
			media, err = ms.CreateMedia(ctx, ref, mediaType)
			if e, ok := fault.As(err); ok {
				if e.Code == fault.CodeNotFound {
					return nil, fault.NotFound("media not found")
				}
				return nil, fault.Internal("error getting media")
			}
		} else {
			ms.logger.Error("failed to retrieve media", err)
			return nil, fault.Internal("error getting media")
		}
	}

	return media, nil
}

func (ms *mediaService) GetMediaByID(ctx context.Context, id uuid.UUID) (*model.Media, error) {
	media, err := ms.store.Medias().One(ctx, &model.MediaF{ID: &id})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("media not found")
		}
		ms.logger.Error("media get failed", err)
		return nil, fault.Internal("error getting media")
	}

	return media, nil
}

func (ms *mediaService) GetDetailedMovie(ctx context.Context, ref int) (*tmdb.DetailedMovie, error) {
	movie, err := ms.tmdb.GetMovie(ref)
	if err != nil {
		if tmdb.IsNotFound(err) {
			return nil, fault.NotFound("movie not found")
		}
		ms.logger.Error("failed to search movie by ref", err)
		return nil, fault.Internal("error getting movie")
	}

	return movie, err
}

func (ms *mediaService) GetDetailedShow(ctx context.Context, ref int) (*tmdb.DetailedShow, error) {
	show, err := ms.tmdb.GetShow(ref)
	if err != nil {
		if tmdb.IsNotFound(err) {
			return nil, fault.NotFound("show not found")
		}
		ms.logger.Error("failed to search show by ref", err)
		return nil, fault.Internal("error getting show")
	}

	return show, err
}

func (ms *mediaService) GetMovieCredits(_ context.Context, ref int) (*tmdb.MovieCredits, error) {
	credits, err := ms.tmdb.GetMovieCredits(ref)
	if err != nil {
		if tmdb.IsNotFound(err) {
			return nil, fault.NotFound("movie not found")
		}
		ms.logger.Error("failed to search movie credits by ref", err)
		return nil, fault.Internal("error getting movie credits")
	}

	return credits, nil
}

func (ms *mediaService) GetShowCredits(_ context.Context, ref int) (*tmdb.ShowCredits, error) {
	credits, err := ms.tmdb.GetShowCredits(ref)
	if err != nil {
		if tmdb.IsNotFound(err) {
			return nil, fault.NotFound("show not found")
		}
		ms.logger.Error("failed to search show credits by ref", err)
		return nil, fault.Internal("error getting show credits")
	}

	return credits, nil
}

func (ms *mediaService) GetShowDetailedSeason(_ context.Context, ref int, seasonNumber int) (*tmdb.DetailedSeason, error) {
	season, err := ms.tmdb.GetShowSeasonDetails(ref, seasonNumber)
	if err != nil {
		if tmdb.IsNotFound(err) {
			return nil, fault.NotFound("show not found")
		}
		ms.logger.Error("failed to search show season by ref", err)
		return nil, fault.Internal("error getting show season")
	}

	return season, nil
}

func (ms *mediaService) GetMovieList(_ context.Context, list tmdb.MovieList) (movies []tmdb.Movie, err error) {
	switch list {
	case tmdb.MovieListNowPlaying:
		movies, err = ms.tmdb.ListNowPlayingMovies()
	case tmdb.MovieListPopular:
		movies, err = ms.tmdb.ListPopularMovies()
	case tmdb.MovieListTopRated:
		movies, err = ms.tmdb.ListTopRatedMovies()
	case tmdb.MovieListUpcoming:
		movies, err = ms.tmdb.ListUpcomingMovies()
	default:
		return nil, fault.BadRequest("invalid movie list")
	}
	if err != nil {
		ms.logger.Error("failed to search movie list", err)
		return nil, fault.Internal("error getting movie list")
	}
	return movies, nil
}

func (ms *mediaService) GetShowList(_ context.Context, list tmdb.ShowList) (shows []tmdb.Show, err error) {
	switch list {
	case tmdb.ShowListAiringToday:
		shows, err = ms.tmdb.ListAiringTodayShows()
	case tmdb.ShowListOnTheAir:
		shows, err = ms.tmdb.ListOnTheAirShows()
	case tmdb.ShowListPopular:
		shows, err = ms.tmdb.ListPopularShows()
	case tmdb.ShowListTopRated:
		shows, err = ms.tmdb.ListTopRatedShows()
	default:
		shows, err = nil, fault.BadRequest("invalid show list")
	}
	if err != nil {
		ms.logger.Error("failed to search show list", err)
		return nil, fault.Internal("error getting show list")
	}
	return shows, nil
}

func (ms *mediaService) mediaFromRef(ref int, mediaType model.MediaType) (*model.Media, error) {
	switch mediaType {
	case model.MediaTypeMovie:
		movie, err := ms.tmdb.GetMovie(ref)
		if err != nil {
			if tmdb.IsNotFound(err) {
				return nil, fault.NotFound("movie not found")
			}
			ms.logger.Error("failed to search movie by ref", err)
			return nil, fault.Internal("error getting movie")
		}
		return &model.Media{
			Ref:          movie.ID,
			MediaType:    model.MediaTypeMovie,
			Overview:     movie.Overview,
			BackdropPath: movie.BackdropPath,
			Language:     movie.OriginalLanguage,
			PosterPath:   movie.PosterPath,
			ReleaseDate:  movie.ReleaseDate,
			Title:        movie.Title,
		}, nil
	case model.MediaTypeShow:
		show, err := ms.tmdb.GetShow(ref)
		if err != nil {
			if tmdb.IsNotFound(err) {
				return nil, fault.NotFound("show not found")
			}
			ms.logger.Error("failed to search movie by ref", err)
			return nil, fault.Internal("error getting movie")
		}
		return &model.Media{
			Ref:          show.ID,
			MediaType:    model.MediaTypeShow,
			Overview:     show.Overview,
			BackdropPath: show.BackdropPath,
			Language:     show.OriginalLanguage,
			PosterPath:   show.PosterPath,
			ReleaseDate:  show.FirstAirDate,
			Title:        show.Name,
		}, nil
	}
	return nil, fault.BadRequest("invalid media type") // unreachable
}

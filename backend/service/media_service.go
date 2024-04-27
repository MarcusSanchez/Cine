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

func (ms *mediaService) mediaFromRef(ref int, mediaType model.MediaType) (*model.Media, error) {
	switch mediaType {
	case model.MediaTypeMovie:
		movie, err := ms.tmdb.SearchMovieByRef(ref)
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
		show, err := ms.tmdb.SearchShowByRef(ref)
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

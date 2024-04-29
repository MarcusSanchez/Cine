package ent

import (
	"cine/datastore/ent/ent"
	Media "cine/datastore/ent/ent/media"
	"cine/datastore/ent/ent/predicate"
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type mediaRepository struct {
	client *ent.Client
}

func newMediaRepository(client *ent.Client) repository.MediaRepository {
	return &mediaRepository{client: client}
}

func (mr *mediaRepository) One(ctx context.Context, mediaFs ...*model.MediaF) (*model.Media, error) {
	q := mr.client.Media.Query()
	q = q.Where(mr.filters(mediaFs)...)

	media, err := q.First(ctx)
	return c.media(media), c.error(err)
}

func (mr *mediaRepository) All(ctx context.Context, mediaFs ...*model.MediaF) ([]*model.Media, error) {
	q := mr.client.Media.Query()
	q = q.Where(mr.filters(mediaFs)...)

	medias, err := q.All(ctx)
	return c.medias(medias), c.error(err)
}

func (mr *mediaRepository) Exists(ctx context.Context, mediaFs ...*model.MediaF) (bool, error) {
	q := mr.client.Media.Query()
	q = q.Where(mr.filters(mediaFs)...)

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (mr *mediaRepository) Count(ctx context.Context, mediaFs ...*model.MediaF) (int, error) {
	q := mr.client.Media.Query()
	q = q.Where(mr.filters(mediaFs)...)

	count, err := q.Count(ctx)
	return count, c.error(err)
}

func (mr *mediaRepository) Insert(ctx context.Context, media *model.Media) (*model.Media, error) {
	i := mr.create(media)

	iMedia, err := i.Save(ctx)
	return c.media(iMedia), c.error(err)
}

func (mr *mediaRepository) InsertBulk(ctx context.Context, medias []*model.Media) ([]*model.Media, error) {
	i := mr.createBulk(medias)

	iMedias, err := i.Save(ctx)
	return c.medias(iMedias), c.error(err)
}

func (mr *mediaRepository) Update(ctx context.Context, id uuid.UUID, _ *model.MediaU) (*model.Media, error) {
	q := mr.client.Media.UpdateOneID(id)

	q.SetUpdatedAt(time.Now())

	media, err := q.Save(ctx)
	return c.media(media), c.error(err)
}

func (mr *mediaRepository) UpdateExec(ctx context.Context, _ *model.MediaU, mediaFs ...*model.MediaF) (int, error) {
	q := mr.client.Media.Update()
	q = q.Where(mr.filters(mediaFs)...)

	q.SetUpdatedAt(time.Now())

	affected, err := q.Save(ctx)
	return affected, c.error(err)
}

func (mr *mediaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := mr.client.Media.DeleteOneID(id)

	err := q.Exec(ctx)
	return c.error(err)
}

func (mr *mediaRepository) DeleteExec(ctx context.Context, mediaFs ...*model.MediaF) (int, error) {
	q := mr.client.Media.Delete()
	q = q.Where(mr.filters(mediaFs)...)

	affected, err := q.Exec(ctx)
	return affected, c.error(err)
}

func (mr *mediaRepository) filters(mediaFs []*model.MediaF) []predicate.Media {
	var mediaF *model.MediaF
	if len(mediaFs) > 0 {
		mediaF = mediaFs[0]
	}
	var filters []predicate.Media
	if mediaF != nil {
		if mediaF.ID != nil {
			filters = append(filters, Media.ID(*mediaF.ID))
		}
		if mediaF.Ref != nil {
			filters = append(filters, Media.Ref(*mediaF.Ref))
		}
		if mediaF.MediaType != nil {
			filters = append(filters, Media.MediaTypeEQ(Media.MediaType(*mediaF.MediaType)))
		}
		if mediaF.Overview != nil {
			filters = append(filters, Media.OverviewEQ(*mediaF.Overview))
		}
		if mediaF.BackdropPath != nil {
			filters = append(filters, Media.BackdropPathEQ(*mediaF.BackdropPath))
		}
		if mediaF.Language != nil {
			filters = append(filters, Media.LanguageEQ(*mediaF.Language))
		}
		if mediaF.PosterPath != nil {
			filters = append(filters, Media.PosterPathEQ(*mediaF.PosterPath))
		}
		if mediaF.ReleaseDate != nil {
			filters = append(filters, Media.ReleaseDateEQ(*mediaF.ReleaseDate))
		}
		if mediaF.Title != nil {
			filters = append(filters, Media.TitleEQ(*mediaF.Title))
		}
		if mediaF.CreatedAt != nil {
			filters = append(filters, Media.CreatedAt(*mediaF.CreatedAt))
		}
		if mediaF.UpdatedAt != nil {
			filters = append(filters, Media.UpdatedAt(*mediaF.UpdatedAt))
		}
	}
	return filters
}

func (mr *mediaRepository) create(media *model.Media) *ent.MediaCreate {
	return mr.client.Media.Create().
		SetID(uuid.New()).
		SetRef(media.Ref).
		SetMediaType(Media.MediaType(media.MediaType)).
		SetOverview(media.Overview).
		SetBackdropPath(media.BackdropPath).
		SetLanguage(media.Language).
		SetPosterPath(media.PosterPath).
		SetReleaseDate(media.ReleaseDate).
		SetTitle(media.Title).
		SetCreatedAt(time.Now())
}

func (mr *mediaRepository) createBulk(medias []*model.Media) *ent.MediaCreateBulk {
	builders := make([]*ent.MediaCreate, 0, len(medias))
	for _, media := range medias {
		builders = append(builders, mr.create(media))
	}
	return mr.client.Media.CreateBulk(builders...)
}

package service

import (
	"cine/datastore"
	"cine/entity/model"
	"cine/pkg/fault"
	"cine/pkg/logger"
	"context"
	"github.com/google/uuid"
	"time"
)

type ListService interface {
	CreateList(ctx context.Context, ownerID uuid.UUID, title string) (*model.List, error)
	DeleteList(ctx context.Context, ownerID uuid.UUID, id uuid.UUID) error
	UpdateList(ctx context.Context, ownerID uuid.UUID, id uuid.UUID, listU *model.ListU) (*model.List, error)
	AddMemberToList(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error
	RemoveMemberFromList(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error
	GetAllLists(ctx context.Context, memberID uuid.UUID) ([]*model.List, error)
	GetPublicLists(ctx context.Context, userID uuid.UUID) ([]*model.List, error)
	GetPrivateDetailedList(ctx context.Context, memberID uuid.UUID, id uuid.UUID) (*DetailedList, error)
	GetPublicDetailedList(ctx context.Context, id uuid.UUID) (*DetailedList, error)
	AddMovieToList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
	RemoveMovieFromList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
	AddShowToList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
	RemoveShowFromList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
}

type listService struct {
	store  datastore.Store
	logger logger.Logger
	media  MediaService
}

func NewListService(
	store datastore.Store,
	logger logger.Logger,
	mediaService MediaService,
) ListService {
	return &listService{
		store:  store,
		logger: logger,
		media:  mediaService,
	}
}

func (ls *listService) CreateList(ctx context.Context, ownerID uuid.UUID, title string) (*model.List, error) {
	exists, err := ls.store.Users().Exists(ctx, &model.UserF{ID: &ownerID})
	if err != nil {
		ls.logger.Error("error fetching user", err)
		return nil, fault.Internal("error creating list")
	} else if !exists {
		return nil, fault.NotFound("user not found")
	}

	tx, err := ls.store.Transaction(ctx)
	if err != nil {
		ls.logger.Error("error starting transaction", err)
		return nil, fault.Internal("error creating list")
	}
	defer tx.Rollback()

	list, err := tx.Lists().Insert(
		ctx, &model.List{
			OwnerID:   ownerID,
			Title:     title,
			Public:    false,
			CreatedAt: time.Now(),
		},
	)
	if err != nil {
		ls.logger.Error("error creating list", err)
		return nil, fault.Internal("error creating list")
	}

	if err = tx.Lists().AddMember(ctx, list, ownerID); err != nil {
		ls.logger.Error("error adding user to list", err)
		return nil, fault.Internal("error creating list")
	}

	if err = tx.Commit(); err != nil {
		ls.logger.Error("error committing transaction", err)
		return nil, fault.Internal("error creating list")
	}

	return list, nil
}

func (ls *listService) DeleteList(ctx context.Context, ownerID uuid.UUID, id uuid.UUID) error {
	exists, err := ls.store.Lists().Exists(ctx, &model.ListF{ID: &id, OwnerID: &ownerID})
	if err != nil {
		ls.logger.Error("error checking list existence", err)
		return fault.Internal("error deleting list")
	} else if !exists {
		return fault.NotFound("list not found")
	}

	if err = ls.store.Lists().Delete(ctx, id); err != nil {
		ls.logger.Error("error deleting list", err)
		return fault.Internal("error deleting list")
	}

	return nil
}

func (ls *listService) UpdateList(ctx context.Context, ownerID uuid.UUID, id uuid.UUID, listU *model.ListU) (*model.List, error) {
	if listU.Title == nil && listU.Public == nil {
		return nil, fault.BadRequest("no fields to update")
	}

	exists, err := ls.store.Lists().Exists(ctx, &model.ListF{ID: &id, OwnerID: &ownerID})
	if err != nil {
		ls.logger.Error("error checking list existence", err)
		return nil, fault.Internal("error updating list")
	} else if !exists {
		return nil, fault.NotFound("list not found")
	}

	list, err := ls.store.Lists().Update(ctx, id, listU)
	if err != nil {
		ls.logger.Error("error updating list", err)
		return nil, fault.Internal("error updating list")
	}

	return list, nil
}

func (ls *listService) AddMemberToList(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error {
	exists, err := ls.store.Users().Exists(ctx, &model.UserF{ID: &userID})
	if err != nil {
		ls.logger.Error("error checking user existence", err)
		return fault.Internal("error adding user to list")
	} else if !exists {
		return fault.NotFound("user not found")
	}

	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &listID, OwnerID: &ownerID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return fault.Internal("error adding user to list")
	}

	exists, err = ls.store.Lists().Exists(ctx, &model.ListF{ID: &listID, HasMemberID: &userID})
	if err != nil {
		ls.logger.Error("error checking user existence", err)
		return fault.Internal("error adding user to list")
	} else if exists {
		return fault.Conflict("user already exists in list")
	}

	if err = ls.store.Lists().AddMember(ctx, list, userID); err != nil {
		ls.logger.Error("error adding user to list", err)
		return fault.Internal("error adding user to list")
	}

	return nil
}

func (ls *listService) RemoveMemberFromList(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error {
	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &listID, OwnerID: &ownerID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return fault.Internal("error removing user from list")
	}

	exists, err := ls.store.Lists().Exists(ctx, &model.ListF{ID: &listID, HasMemberID: &userID})
	if err != nil {
		ls.logger.Error("error checking user existence", err)
		return fault.Internal("error removing user from list")
	} else if !exists {
		return fault.NotFound("user not found")
	}

	if err = ls.store.Lists().RemoveMember(ctx, list, userID); err != nil {
		ls.logger.Error("error removing user from list", err)
		return fault.Internal("error removing user from list")
	}

	return nil
}

func (ls *listService) GetAllLists(ctx context.Context, memberID uuid.UUID) ([]*model.List, error) {
	exists, err := ls.store.Users().Exists(ctx, &model.UserF{ID: &memberID})
	if err != nil {
		ls.logger.Error("error checking user existence", err)
		return nil, fault.Internal("error fetching list")
	} else if !exists {
		return nil, fault.NotFound("user not found")
	}

	lists, err := ls.store.Lists().All(ctx, &model.ListF{HasMemberID: &memberID})
	if err != nil {
		ls.logger.Error("error fetching list", err)
		return nil, fault.Internal("error fetching list")
	}

	return lists, nil
}

func (ls *listService) GetPublicLists(ctx context.Context, userID uuid.UUID) ([]*model.List, error) {
	exists, err := ls.store.Users().Exists(ctx, &model.UserF{ID: &userID})
	if err != nil {
		ls.logger.Error("error checking user existence", err)
		return nil, fault.Internal("error fetching list")
	} else if !exists {
		return nil, fault.NotFound("user not found")
	}

	public := true

	lists, err := ls.store.Lists().All(ctx, &model.ListF{HasMemberID: &userID, Public: &public})
	if err != nil {
		ls.logger.Error("error fetching list", err)
		return nil, fault.Internal("error fetching list")
	}

	return lists, nil
}

type DetailedList struct {
	List    *model.List    `json:"list"`
	Members []*model.User  `json:"memberIDs"`
	Movies  []*model.Media `json:"movies"`
	Shows   []*model.Media `json:"shows"`
}

func (ls *listService) GetPrivateDetailedList(ctx context.Context, memberID uuid.UUID, id uuid.UUID) (*DetailedList, error) {
	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &id, HasMemberID: &memberID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return nil, fault.Internal("error fetching list")
	}

	users, err := ls.store.Lists().AllMembers(ctx, list)
	if err != nil {
		ls.logger.Error("error fetching users", err)
		return nil, fault.Internal("error fetching list")
	}

	media, err := ls.store.Lists().AllMedia(ctx, list)
	if err != nil {
		ls.logger.Error("error fetching media", err)
		return nil, fault.Internal("error fetching list")
	}

	return &DetailedList{
		List:    list,
		Members: users,
		Movies:  ls.filterMedia(media, model.MediaTypeMovie),
		Shows:   ls.filterMedia(media, model.MediaTypeShow),
	}, nil
}

func (ls *listService) GetPublicDetailedList(ctx context.Context, id uuid.UUID) (*DetailedList, error) {
	public := true

	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &id, Public: &public})
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return nil, fault.Internal("error fetching list")
	}

	users, err := ls.store.Lists().AllMembers(ctx, list)
	if err != nil {
		ls.logger.Error("error fetching users", err)
		return nil, fault.Internal("error fetching list")
	}

	media, err := ls.store.Lists().AllMedia(ctx, list)
	if err != nil {
		ls.logger.Error("error fetching media", err)
		return nil, fault.Internal("error fetching list")
	}

	return &DetailedList{
		List:    list,
		Members: users,
		Movies:  ls.filterMedia(media, model.MediaTypeMovie),
		Shows:   ls.filterMedia(media, model.MediaTypeShow),
	}, nil
}

func (ls *listService) AddMovieToList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &listID, HasMemberID: &memberID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return fault.Internal("error adding movie to list")
	}

	media, err := ls.media.GetMedia(ctx, ref, model.MediaTypeMovie)
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("movie not found")
		}
		ls.logger.Error("error fetching movie", err)
		return fault.Internal("error adding movie to list")
	}

	if err = ls.store.Lists().AddMedia(ctx, list, media.ID); err != nil {
		ls.logger.Error("error adding movie to list", err)
		return fault.Internal("error adding movie to list")
	}

	return nil
}

func (ls *listService) RemoveMovieFromList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &listID, HasMemberID: &memberID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return fault.Internal("error removing movie from list")
	}

	media, err := ls.media.GetMedia(ctx, ref, model.MediaTypeMovie)
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("movie not found")
		}
		ls.logger.Error("error fetching movie", err)
		return fault.Internal("error removing movie from list")
	}

	if err = ls.store.Lists().RemoveMedia(ctx, list, media.ID); err != nil {
		ls.logger.Error("error removing movie from list", err)
		return fault.Internal("error removing movie from list")
	}

	return nil
}

func (ls *listService) AddShowToList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &listID, HasMemberID: &memberID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return fault.Internal("error adding show to list")
	}

	media, err := ls.media.GetMedia(ctx, ref, model.MediaTypeShow)
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("show not found")
		}
		ls.logger.Error("error fetching show", err)
		return fault.Internal("error adding show to list")
	}

	if err = ls.store.Lists().AddMedia(ctx, list, media.ID); err != nil {
		ls.logger.Error("error adding show to list", err)
		return fault.Internal("error adding show to list")
	}

	return nil
}

func (ls *listService) RemoveShowFromList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	list, err := ls.store.Lists().One(ctx, &model.ListF{ID: &listID, HasMemberID: &memberID})
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("list not found")
		}
		ls.logger.Error("error fetching list", err)
		return fault.Internal("error removing show from list")
	}

	media, err := ls.media.GetMedia(ctx, ref, model.MediaTypeShow)
	if err != nil {
		if datastore.IsNotFound(err) {
			return fault.NotFound("show not found")
		}
		ls.logger.Error("error fetching show", err)
		return fault.Internal("error removing show from list")
	}

	if err = ls.store.Lists().RemoveMedia(ctx, list, media.ID); err != nil {
		ls.logger.Error("error removing show from list", err)
		return fault.Internal("error removing show from list")
	}

	return nil
}

func (ls *listService) filterMedia(medias []*model.Media, mediaType model.MediaType) []*model.Media {
	var movies []*model.Media
	for _, media := range medias {
		if media.MediaType == mediaType {
			movies = append(movies, media)
		}
	}
	return movies
}

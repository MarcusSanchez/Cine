package ent

import (
	"cine/datastore/ent/ent"
	List "cine/datastore/ent/ent/list"
	"cine/datastore/ent/ent/predicate"
	User "cine/datastore/ent/ent/user"
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type listRepository struct {
	client *ent.Client
}

func newListRepository(client *ent.Client) repository.ListRepository {
	return &listRepository{client: client}
}

func (lr *listRepository) One(ctx context.Context, listFs ...*model.ListF) (*model.List, error) {
	q := lr.client.List.Query()
	q = q.Where(lr.filters(listFs)...)

	list, err := q.First(ctx)
	return c.list(list), c.error(err)
}

func (lr *listRepository) All(ctx context.Context, listFs ...*model.ListF) ([]*model.List, error) {
	q := lr.client.List.Query()
	q = q.Where(lr.filters(listFs)...)

	lists, err := q.All(ctx)
	return c.lists(lists), c.error(err)
}

func (lr *listRepository) Exists(ctx context.Context, listFs ...*model.ListF) (bool, error) {
	q := lr.client.List.Query()
	q = q.Where(lr.filters(listFs)...)

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (lr *listRepository) Count(ctx context.Context, listFs ...*model.ListF) (int, error) {
	q := lr.client.List.Query()
	q = q.Where(lr.filters(listFs)...)

	count, err := q.Count(ctx)
	return count, c.error(err)
}

func (lr *listRepository) Insert(ctx context.Context, list *model.List) (*model.List, error) {
	i := lr.create(list)

	iList, err := i.Save(ctx)
	return c.list(iList), c.error(err)
}

func (lr *listRepository) InsertBulk(ctx context.Context, lists []*model.List) ([]*model.List, error) {
	i := lr.createBulk(lists)

	iLists, err := i.Save(ctx)
	return c.lists(iLists), c.error(err)
}

func (lr *listRepository) Update(ctx context.Context, id uuid.UUID, listU *model.ListU) (*model.List, error) {
	q := lr.client.List.UpdateOneID(id)

	q.SetUpdatedAt(time.Now())
	q.SetNillableTitle(listU.Title)
	q.SetNillablePublic(listU.Public)

	list, err := q.Save(ctx)
	return c.list(list), c.error(err)
}

func (lr *listRepository) UpdateExec(ctx context.Context, listU *model.ListU, listFs ...*model.ListF) (int, error) {
	q := lr.client.List.Update()
	q = q.Where(lr.filters(listFs)...)

	q.SetUpdatedAt(time.Now())
	q.SetNillableTitle(listU.Title)
	q.SetNillablePublic(listU.Public)

	affected, err := q.Save(ctx)
	return affected, c.error(err)
}

func (lr *listRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := lr.client.List.DeleteOneID(id)

	err := q.Exec(ctx)
	return c.error(err)
}

func (lr *listRepository) DeleteExec(ctx context.Context, listFs ...*model.ListF) (int, error) {
	q := lr.client.List.Delete()
	q = q.Where(lr.filters(listFs)...)

	affected, err := q.Exec(ctx)
	return affected, c.error(err)
}

func (lr *listRepository) AllUsers(ctx context.Context, list *model.List) ([]*model.User, error) {
	q := c.entList(list).QueryUsers()

	users, err := q.All(ctx)
	return c.users(users), c.error(err)
}

func (lr *listRepository) ExistsUser(ctx context.Context, list *model.List, userID uuid.UUID) (bool, error) {
	q := lr.client.List.Query()
	q = q.Where(List.ID(list.ID), List.HasUsersWith(User.ID(userID)))

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (lr *listRepository) AddUser(ctx context.Context, list *model.List, userID uuid.UUID) error {
	q := lr.client.List.UpdateOne(c.entList(list))
	q = q.AddUserIDs(userID)

	_, err := q.Save(ctx)
	return c.error(err)
}

func (lr *listRepository) RemoveUser(ctx context.Context, list *model.List, userID uuid.UUID) error {
	q := lr.client.List.UpdateOne(c.entList(list))
	q = q.RemoveUserIDs(userID)

	_, err := q.Save(ctx)
	return c.error(err)
}

func (lr *listRepository) AddMedia(ctx context.Context, list *model.List, mediaID uuid.UUID) error {
	q := lr.client.List.UpdateOne(c.entList(list))
	q = q.AddMediaIDs(mediaID)

	_, err := q.Save(ctx)
	return c.error(err)
}

func (lr *listRepository) RemoveMedia(ctx context.Context, list *model.List, mediaID uuid.UUID) error {
	q := lr.client.List.UpdateOne(c.entList(list))
	q = q.RemoveMediaIDs(mediaID)

	_, err := q.Save(ctx)
	return c.error(err)
}

func (lr *listRepository) AllMedia(ctx context.Context, list *model.List) ([]*model.Media, error) {
	q := c.entList(list).QueryMedias()

	medias, err := q.All(ctx)
	return c.medias(medias), c.error(err)
}

func (lr *listRepository) filters(listFs []*model.ListF) []predicate.List {
	var listF *model.ListF
	if len(listFs) > 0 {
		listF = listFs[0]
	}
	var filters []predicate.List
	if listF != nil {
		if listF.ID != nil {
			filters = append(filters, List.ID(*listF.ID))
		}
		if listF.OwnerID != nil {
			filters = append(filters, List.OwnerID(*listF.OwnerID))
		}
		if listF.Title != nil {
			filters = append(filters, List.Title(*listF.Title))
		}
		if listF.Public != nil {
			filters = append(filters, List.Public(*listF.Public))
		}
		if listF.CreatedAt != nil {
			filters = append(filters, List.CreatedAt(*listF.CreatedAt))
		}
		if listF.UpdatedAt != nil {
			filters = append(filters, List.UpdatedAt(*listF.UpdatedAt))
		}
		if listF.HasMemberID != nil {
			filters = append(filters, List.HasUsersWith(User.ID(*listF.HasMemberID)))
		}
	}
	return filters
}

func (lr *listRepository) create(list *model.List) *ent.ListCreate {
	return lr.client.List.Create().
		SetID(uuid.New()).
		SetOwnerID(list.OwnerID).
		SetTitle(list.Title).
		SetPublic(list.Public).
		SetCreatedAt(time.Now())
}

func (lr *listRepository) createBulk(lists []*model.List) *ent.ListCreateBulk {
	builders := make([]*ent.ListCreate, 0, len(lists))
	for _, list := range lists {
		builders = append(builders, lr.create(list))
	}
	return lr.client.List.CreateBulk(builders...)
}

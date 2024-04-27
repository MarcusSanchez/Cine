package ent

import (
	"cine/datastore/ent/ent"
	"cine/datastore/ent/ent/predicate"
	Session "cine/datastore/ent/ent/session"
	"cine/entity/model"
	"cine/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type sessionRepository struct {
	client *ent.Client
}

func newSessionRepository(client *ent.Client) repository.SessionRepository {
	return &sessionRepository{client: client}
}

func (sr *sessionRepository) One(ctx context.Context, sessionFs ...*model.SessionF) (*model.Session, error) {
	q := sr.client.Session.Query()
	q = q.Where(sr.filters(sessionFs)...)

	session, err := q.First(ctx)
	return c.session(session), c.error(err)
}

func (sr *sessionRepository) All(ctx context.Context, sessionFs ...*model.SessionF) ([]*model.Session, error) {
	q := sr.client.Session.Query()
	q = q.Where(sr.filters(sessionFs)...)

	sessions, err := q.All(ctx)
	return c.sessions(sessions), c.error(err)
}

func (sr *sessionRepository) Exists(ctx context.Context, sessionFs ...*model.SessionF) (bool, error) {
	q := sr.client.Session.Query()
	q = q.Where(sr.filters(sessionFs)...)

	exists, err := q.Exist(ctx)
	return exists, c.error(err)
}

func (sr *sessionRepository) Count(ctx context.Context, sessionFs ...*model.SessionF) (int, error) {
	q := sr.client.Session.Query()
	q = q.Where(sr.filters(sessionFs)...)

	count, err := q.Count(ctx)
	return count, c.error(err)
}

func (sr *sessionRepository) Insert(ctx context.Context, session *model.Session) (*model.Session, error) {
	i := sr.create(session)

	iSession, err := i.Save(ctx)
	return c.session(iSession), c.error(err)
}

func (sr *sessionRepository) InsertBulk(ctx context.Context, sessions []*model.Session) ([]*model.Session, error) {
	i := sr.createBulk(sessions)

	iSessions, err := i.Save(ctx)
	return c.sessions(iSessions), c.error(err)
}

func (sr *sessionRepository) Update(ctx context.Context, id uuid.UUID, sessionU *model.SessionU) (*model.Session, error) {
	q := sr.client.Session.UpdateOneID(id)

	q.SetUpdatedAt(time.Now())
	q.SetNillableExpiration(sessionU.Expiration)

	session, err := q.Save(ctx)
	return c.session(session), c.error(err)
}

func (sr *sessionRepository) UpdateExec(ctx context.Context, sessionU *model.SessionU, sessionFs ...*model.SessionF) (int, error) {
	q := sr.client.Session.Update()
	q = q.Where(sr.filters(sessionFs)...)

	q.SetUpdatedAt(time.Now())
	q.SetNillableExpiration(sessionU.Expiration)

	affected, err := q.Save(ctx)
	return affected, c.error(err)
}

func (sr *sessionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	q := sr.client.Session.DeleteOneID(id)

	err := q.Exec(ctx)
	return c.error(err)
}

func (sr *sessionRepository) DeleteExec(ctx context.Context, sessionFs ...*model.SessionF) (int, error) {
	q := sr.client.Session.Delete()
	q = q.Where(sr.filters(sessionFs)...)

	affected, err := q.Exec(ctx)
	return affected, c.error(err)
}

func (sr *sessionRepository) filters(sessionFs []*model.SessionF) []predicate.Session {
	var sessionF *model.SessionF
	if len(sessionFs) > 0 {
		sessionF = sessionFs[0]
	}
	var filters []predicate.Session
	if sessionF != nil {
		if sessionF.ID != nil {
			filters = append(filters, Session.ID(*sessionF.ID))
		}
		if sessionF.UserID != nil {
			filters = append(filters, Session.UserID(*sessionF.UserID))
		}
		if sessionF.CSRF != nil {
			filters = append(filters, Session.Csrf(*sessionF.CSRF))
		}
		if sessionF.Token != nil {
			filters = append(filters, Session.Token(*sessionF.Token))
		}
		if sessionF.Expiration != nil {
			filters = append(filters, Session.Expiration(*sessionF.Expiration))
		}
		if sessionF.CreatedAt != nil {
			filters = append(filters, Session.CreatedAt(*sessionF.CreatedAt))
		}
		if sessionF.UpdatedAt != nil {
			filters = append(filters, Session.UpdatedAt(*sessionF.UpdatedAt))
		}
	}
	return filters
}

func (sr *sessionRepository) create(session *model.Session) *ent.SessionCreate {
	return sr.client.Session.Create().
		SetID(uuid.New()).
		SetUserID(session.UserID).
		SetCsrf(session.CSRF).
		SetToken(session.Token).
		SetExpiration(session.Expiration).
		SetCreatedAt(time.Now())
}

func (sr *sessionRepository) createBulk(sessions []*model.Session) *ent.SessionCreateBulk {
	builders := make([]*ent.SessionCreate, 0, len(sessions))
	for _, session := range sessions {
		builders = append(builders, sr.create(session))
	}
	return sr.client.Session.CreateBulk(builders...)
}

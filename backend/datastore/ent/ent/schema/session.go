package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.UUID("user_id", uuid.UUID{}).Immutable(),
		field.UUID("csrf", uuid.UUID{}).Unique().Immutable(),
		field.UUID("token", uuid.UUID{}).Unique().Immutable(),
		field.Time("expiration"),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M User <-- Session
		edge.From("user", User.Type).Ref("sessions").Field("user_id").Unique().Required().Immutable(),
	}
}

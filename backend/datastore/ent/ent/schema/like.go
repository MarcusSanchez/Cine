package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.UUID("user_id", uuid.UUID{}).Immutable(),
		field.UUID("comment_id", uuid.UUID{}).Immutable(),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M User <-- Like
		edge.From("user", User.Type).Ref("likes").Field("user_id").Unique().Required().Immutable(),
		// O2M Comment <-- Like
		edge.From("comment", Comment.Type).Ref("likes").Field("comment_id").Unique().Required().Immutable(),
	}
}

func (Like) Indexes() []ent.Index {
	return []ent.Index{
		// a user can only like a comment once
		index.Fields("user_id", "comment_id").Unique(),
	}
}

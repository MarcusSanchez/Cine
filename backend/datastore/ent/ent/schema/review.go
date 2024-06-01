package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.UUID("user_id", uuid.UUID{}).Immutable(),
		field.UUID("media_id", uuid.UUID{}).Immutable(),
		field.String("content"),
		field.Int("rating"),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M User <-- Review
		edge.From("user", User.Type).Ref("reviews").Field("user_id").Unique().Required().Immutable(),
		// O2M Media <-- Review
		edge.From("media", Media.Type).Ref("reviews").Field("media_id").Unique().Required().Immutable(),
	}
}

func (Review) Indexes() []ent.Index {
	return []ent.Index{
		// a user can only review a movie/tv show once
		index.Fields("user_id", "media_id").Unique(),
	}
}

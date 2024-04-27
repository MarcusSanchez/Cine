package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.UUID("user_id", uuid.UUID{}).Nillable().Immutable(),
		field.UUID("media_id", uuid.UUID{}).Immutable(),
		field.UUID("replying_to", uuid.UUID{}).Nillable().Optional().Immutable(),
		field.String("content"),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M User <-- Comment
		edge.From("user", User.Type).Ref("comments").Field("user_id").Unique().Required().Immutable(),
		// O2M Comment <-- Like
		edge.To("likes", Like.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// O2M Media <-- Comment
		edge.From("media", Media.Type).Ref("comments").Field("media_id").Unique().Required().Immutable(),
	}
}

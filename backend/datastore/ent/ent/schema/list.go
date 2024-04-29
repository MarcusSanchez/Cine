package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// List holds the schema definition for the List entity.
type List struct {
	ent.Schema
}

// Fields of the List.
func (List) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.UUID("owner_id", uuid.UUID{}).Immutable(),
		field.String("title"),
		field.Bool("public"),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the List.
func (List) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M User (owner) <--> List
		edge.From("owner", User.Type).Ref("owned_lists").Field("owner_id").Unique().Immutable().Required(),
		// M2M User (members) <--> List
		edge.From("members", User.Type).Ref("lists"),
		// M2M Media <--> List
		edge.From("medias", Media.Type).Ref("lists"),
	}
}

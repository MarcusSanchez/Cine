package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.Int("ref").Immutable(),
		field.Enum("media_type").Values("movie", "show").Immutable(),
		field.String("overview"),
		field.String("backdrop_path"),
		field.String("language"),
		field.String("poster_path"),
		field.String("release_date"),
		field.String("title"),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M Media <-- Comment
		edge.To("comments", Comment.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// M2M Media <--> List
		edge.To("lists", List.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// O2M Media <-- Review
		edge.To("reviews", Review.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Media) Indexes() []ent.Index {
	return []ent.Index{
		// the same ref can be used for a show and/or movie media type, but not for the same media type twice
		index.Fields("ref", "media_type").Unique(),
	}
}

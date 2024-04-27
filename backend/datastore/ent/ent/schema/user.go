package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.String("display_name"),
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("password").Sensitive(),
		field.String("profile_picture"),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M User <-- Session
		edge.To("sessions", Session.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// O2M User <-- Comment
		edge.To("comments", Comment.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
		// O2M User <-- Like
		edge.To("likes", Like.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// M2M User <--> User (Followers)
		edge.To("following", User.Type).From("followers").Annotations(entsql.OnDelete(entsql.Cascade)),
		// M2M User <--> User (friends)
		edge.To("friends", User.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// O2M User <-- Review
		edge.To("reviews", Review.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// M2M User (members) <--> List
		edge.To("lists", List.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		// O2M User (owner) <-- List
		edge.To("owned_lists", List.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

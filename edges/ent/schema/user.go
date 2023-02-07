package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.Int("age"),
		// a foreign key
		field.Int("spouse_id").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	// return []ent.Edge{
	// 	edge.From("groups", Group.Type).
	// 		Ref("users"),
	// 	edge.To("pets", Pet.Type),
	// }

	return []ent.Edge{
		edge.To("spouse", User.Type).
			Unique().
			Field("spouse_id"),
	}
}

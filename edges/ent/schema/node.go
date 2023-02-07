package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Node is a kind of recursive relation sttucture.
// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value").
			NonNegative().
			Unique(),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("next", Node.Type).
		// Unique().
		// From("prev").
		// Unique(),

		edge.To("next", Node.Type).
			Unique(),

		edge.From("prev", Node.Type).
			Ref("next").
			Unique(),
	}
}

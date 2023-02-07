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
		// Test O2M same type
		// One parent could have many children. Otherwise, 
		// One children could only have one parent.
		edge.To("children", Node.Type).
			From("parent").
			Unique(),
		// Or use separated form
		// edge.To("children", Node.Type),
		// edge.From("parent", Node.Type).
		// 	Ref("children").
		// 	Unique(),

		// edge.To("next", Node.Type).
		// Unique().
		// From("prev").
		// Unique(),

		// edge.To("next", Node.Type).
		// 	Unique(),

		// edge.From("prev", Node.Type).
		// 	Ref("next").
		// 	Unique(),
	}
}

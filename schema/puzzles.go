package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Puzzle holds the schema definition for the Puzzle entity.
type Puzzle struct {
	ent.Schema
}

// Fields of the Puzzle.
func (Puzzle) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("answer"),
	}
}

// Edges of the Puzzle.
func (Puzzle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("adventure", Adventure.Type).Ref("puzzles").Unique(),
	}
}

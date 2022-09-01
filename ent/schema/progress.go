package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Progress holds the schema definition for the Progress entity.
type Progress struct {
	ent.Schema
}

// Fields of the Progress.
func (Progress) Fields() []ent.Field {
	return nil
}

// Edges of the Progress.
func (Progress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("adventure", Adventure.Type).Unique(),
		edge.To("question", Question.Type).Unique(),
	}
}

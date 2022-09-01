package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Adventure holds the schema definition for the Adventure entity.
type Adventure struct {
	ent.Schema
}

// Fields of the Adventure.
func (Adventure) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Adventure.
func (Adventure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
	}
}

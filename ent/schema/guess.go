package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Guess holds the schema definition for the Guess entity.
type Guess struct {
	ent.Schema
}

// Fields of the Guess.
func (Guess) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Time("submitted").Default(time.Now()),
	}
}

// Edges of the Guess.
func (Guess) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("question", Question.Type),
	}
}

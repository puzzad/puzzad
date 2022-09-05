package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Guess holds the schema definition for the Guess entity.
type Guess struct {
	ent.Schema
}

// Fields of the Guess.
func (Guess) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
	}
}

func (Guess) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

// Edges of the Guess.
func (Guess) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("question", Question.Type),
		edge.To("team", User.Type),
	}
}

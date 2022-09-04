package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Access holds the schema definition for the Access entity.
type Access struct {
	ent.Schema
}

func (Access) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//This crashes the ent schema description command, but its written as per the docs
		field.ID("team_id", "adventure_id"),
	}
}

// Fields of the Access.
func (Access) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("Paid", "Unpaid", "Expired").Default("Unpaid"),
		field.Int("team_id"),
		field.Int("adventure_id"),
	}
}

// Edges of the Access.
func (Access) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("team", Team.Type).Unique().Required().Field("team_id"),
		edge.To("adventures", Adventure.Type).Unique().Required().Field("adventure_id"),
	}
}

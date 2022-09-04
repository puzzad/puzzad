package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			MaxLen(20).
			MinLen(1),
		field.String("code").
			DefaultFunc(func() string {
				return "t" + uuid.New().String()
			}).
			Unique(),
		field.String("verifyCode").
			DefaultFunc(func() string {
				return uuid.New().String()
			}).
			Unique(),
		field.String("email").Unique(),
		field.Enum("status").Values("Unverified", "Verified", "Disabled").Default("Unverified"),
	}
}

func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("adventures", Adventure.Type).Through("access", Access.Type),
		edge.To("progress", Progress.Type),
	}
}

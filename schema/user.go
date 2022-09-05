package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("verifyCode").
			DefaultFunc(func() string {
				return uuid.New().String()
			}).
			Unique(),
		field.Time("verifyExpiry").
			Default(func() time.Time {
				return time.Now().Add(time.Hour)
			}),
		field.String("email").Unique(),
		field.Enum("status").Values("Unverified", "Verified", "Disabled").Default("Unverified"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("adventures", Adventure.Type).Through("access", Access.Type),
		edge.To("progress", Progress.Type),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/csmith/aca"
	"math/rand"
	"time"
)

var acaGenerator *aca.Generator

func init() {
	var err error
	acaGenerator, err = aca.NewGenerator("-", rand.NewSource(time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}
}

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("Paid", "Unpaid", "Expired").Default("Unpaid"),
		field.String("code").
			DefaultFunc(func() string {
				return acaGenerator.Generate()
			}).
			MaxLen(40).
			Unique(),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("game").Unique().Required(),
		edge.To("adventure", Adventure.Type).Unique().Required(),
		edge.To("current_puzzle", Puzzle.Type).Unique().Required(),
	}
}

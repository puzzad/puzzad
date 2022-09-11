// Code generated by ent, DO NOT EDIT.

package adventure

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/greboid/puzzad/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Public applies equality check predicate on the "public" field. It's identical to PublicEQ.
func Public(v bool) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublic), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Adventure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Adventure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Adventure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Adventure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Adventure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Adventure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// PublicEQ applies the EQ predicate on the "public" field.
func PublicEQ(v bool) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublic), v))
	})
}

// PublicNEQ applies the NEQ predicate on the "public" field.
func PublicNEQ(v bool) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublic), v))
	})
}

// HasGame applies the HasEdge predicate on the "game" edge.
func HasGame() predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GameTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GameTable, GameColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGameWith applies the HasEdge predicate on the "game" edge with a given conditions (other predicates).
func HasGameWith(preds ...predicate.Game) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GameInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GameTable, GameColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPuzzles applies the HasEdge predicate on the "puzzles" edge.
func HasPuzzles() predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PuzzlesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PuzzlesTable, PuzzlesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPuzzlesWith applies the HasEdge predicate on the "puzzles" edge with a given conditions (other predicates).
func HasPuzzlesWith(preds ...predicate.Puzzle) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PuzzlesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PuzzlesTable, PuzzlesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Adventure) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Adventure) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Adventure) predicate.Adventure {
	return predicate.Adventure(func(s *sql.Selector) {
		p(s.Not())
	})
}

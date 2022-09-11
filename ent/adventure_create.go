// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/game"
	"github.com/greboid/puzzad/ent/puzzle"
)

// AdventureCreate is the builder for creating a Adventure entity.
type AdventureCreate struct {
	config
	mutation *AdventureMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AdventureCreate) SetName(s string) *AdventureCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AdventureCreate) SetDescription(s string) *AdventureCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *AdventureCreate) SetNillableDescription(s *string) *AdventureCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetPrice sets the "price" field.
func (ac *AdventureCreate) SetPrice(f float64) *AdventureCreate {
	ac.mutation.SetPrice(f)
	return ac
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (ac *AdventureCreate) SetNillablePrice(f *float64) *AdventureCreate {
	if f != nil {
		ac.SetPrice(*f)
	}
	return ac
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (ac *AdventureCreate) AddGameIDs(ids ...int) *AdventureCreate {
	ac.mutation.AddGameIDs(ids...)
	return ac
}

// AddGame adds the "game" edges to the Game entity.
func (ac *AdventureCreate) AddGame(g ...*Game) *AdventureCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ac.AddGameIDs(ids...)
}

// AddPuzzleIDs adds the "puzzles" edge to the Puzzle entity by IDs.
func (ac *AdventureCreate) AddPuzzleIDs(ids ...int) *AdventureCreate {
	ac.mutation.AddPuzzleIDs(ids...)
	return ac
}

// AddPuzzles adds the "puzzles" edges to the Puzzle entity.
func (ac *AdventureCreate) AddPuzzles(p ...*Puzzle) *AdventureCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPuzzleIDs(ids...)
}

// Mutation returns the AdventureMutation object of the builder.
func (ac *AdventureCreate) Mutation() *AdventureMutation {
	return ac.mutation
}

// Save creates the Adventure in the database.
func (ac *AdventureCreate) Save(ctx context.Context) (*Adventure, error) {
	var (
		err  error
		node *Adventure
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdventureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Adventure)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AdventureMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdventureCreate) SaveX(ctx context.Context) *Adventure {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdventureCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdventureCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdventureCreate) defaults() {
	if _, ok := ac.mutation.Description(); !ok {
		v := adventure.DefaultDescription
		ac.mutation.SetDescription(v)
	}
	if _, ok := ac.mutation.Price(); !ok {
		v := adventure.DefaultPrice
		ac.mutation.SetPrice(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdventureCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Adventure.name"`)}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Adventure.description"`)}
	}
	if _, ok := ac.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Adventure.price"`)}
	}
	return nil
}

func (ac *AdventureCreate) sqlSave(ctx context.Context) (*Adventure, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AdventureCreate) createSpec() (*Adventure, *sqlgraph.CreateSpec) {
	var (
		_node = &Adventure{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: adventure.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adventure.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adventure.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adventure.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := ac.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: adventure.FieldPrice,
		})
		_node.Price = value
	}
	if nodes := ac.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adventure.GameTable,
			Columns: []string{adventure.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PuzzlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adventure.PuzzlesTable,
			Columns: []string{adventure.PuzzlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: puzzle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdventureCreateBulk is the builder for creating many Adventure entities in bulk.
type AdventureCreateBulk struct {
	config
	builders []*AdventureCreate
}

// Save creates the Adventure entities in the database.
func (acb *AdventureCreateBulk) Save(ctx context.Context) ([]*Adventure, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Adventure, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdventureMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdventureCreateBulk) SaveX(ctx context.Context) []*Adventure {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdventureCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdventureCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

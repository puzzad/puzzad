// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/greboid/puzzad/ent/adventure"
	"github.com/greboid/puzzad/ent/predicate"
	"github.com/greboid/puzzad/ent/puzzle"
)

// PuzzleQuery is the builder for querying Puzzle entities.
type PuzzleQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.Puzzle
	withAdventure *AdventureQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PuzzleQuery builder.
func (pq *PuzzleQuery) Where(ps ...predicate.Puzzle) *PuzzleQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PuzzleQuery) Limit(limit int) *PuzzleQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PuzzleQuery) Offset(offset int) *PuzzleQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PuzzleQuery) Unique(unique bool) *PuzzleQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PuzzleQuery) Order(o ...OrderFunc) *PuzzleQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryAdventure chains the current query on the "adventure" edge.
func (pq *PuzzleQuery) QueryAdventure() *AdventureQuery {
	query := &AdventureQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(puzzle.Table, puzzle.FieldID, selector),
			sqlgraph.To(adventure.Table, adventure.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, puzzle.AdventureTable, puzzle.AdventureColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Puzzle entity from the query.
// Returns a *NotFoundError when no Puzzle was found.
func (pq *PuzzleQuery) First(ctx context.Context) (*Puzzle, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{puzzle.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PuzzleQuery) FirstX(ctx context.Context) *Puzzle {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Puzzle ID from the query.
// Returns a *NotFoundError when no Puzzle ID was found.
func (pq *PuzzleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{puzzle.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PuzzleQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Puzzle entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Puzzle entity is found.
// Returns a *NotFoundError when no Puzzle entities are found.
func (pq *PuzzleQuery) Only(ctx context.Context) (*Puzzle, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{puzzle.Label}
	default:
		return nil, &NotSingularError{puzzle.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PuzzleQuery) OnlyX(ctx context.Context) *Puzzle {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Puzzle ID in the query.
// Returns a *NotSingularError when more than one Puzzle ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PuzzleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{puzzle.Label}
	default:
		err = &NotSingularError{puzzle.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PuzzleQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Puzzles.
func (pq *PuzzleQuery) All(ctx context.Context) ([]*Puzzle, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PuzzleQuery) AllX(ctx context.Context) []*Puzzle {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Puzzle IDs.
func (pq *PuzzleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(puzzle.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PuzzleQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PuzzleQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PuzzleQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PuzzleQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PuzzleQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PuzzleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PuzzleQuery) Clone() *PuzzleQuery {
	if pq == nil {
		return nil
	}
	return &PuzzleQuery{
		config:        pq.config,
		limit:         pq.limit,
		offset:        pq.offset,
		order:         append([]OrderFunc{}, pq.order...),
		predicates:    append([]predicate.Puzzle{}, pq.predicates...),
		withAdventure: pq.withAdventure.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithAdventure tells the query-builder to eager-load the nodes that are connected to
// the "adventure" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PuzzleQuery) WithAdventure(opts ...func(*AdventureQuery)) *PuzzleQuery {
	query := &AdventureQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withAdventure = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Puzzle.Query().
//		GroupBy(puzzle.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PuzzleQuery) GroupBy(field string, fields ...string) *PuzzleGroupBy {
	grbuild := &PuzzleGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = puzzle.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Puzzle.Query().
//		Select(puzzle.FieldTitle).
//		Scan(ctx, &v)
func (pq *PuzzleQuery) Select(fields ...string) *PuzzleSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PuzzleSelect{PuzzleQuery: pq}
	selbuild.label = puzzle.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

func (pq *PuzzleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !puzzle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PuzzleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Puzzle, error) {
	var (
		nodes       = []*Puzzle{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withAdventure != nil,
		}
	)
	if pq.withAdventure != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, puzzle.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Puzzle).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Puzzle{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withAdventure; query != nil {
		if err := pq.loadAdventure(ctx, query, nodes, nil,
			func(n *Puzzle, e *Adventure) { n.Edges.Adventure = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PuzzleQuery) loadAdventure(ctx context.Context, query *AdventureQuery, nodes []*Puzzle, init func(*Puzzle), assign func(*Puzzle, *Adventure)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Puzzle)
	for i := range nodes {
		if nodes[i].adventure_puzzles == nil {
			continue
		}
		fk := *nodes[i].adventure_puzzles
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(adventure.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "adventure_puzzles" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PuzzleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PuzzleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *PuzzleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   puzzle.Table,
			Columns: puzzle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: puzzle.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, puzzle.FieldID)
		for i := range fields {
			if fields[i] != puzzle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PuzzleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(puzzle.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = puzzle.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PuzzleGroupBy is the group-by builder for Puzzle entities.
type PuzzleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PuzzleGroupBy) Aggregate(fns ...AggregateFunc) *PuzzleGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PuzzleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *PuzzleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !puzzle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PuzzleGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PuzzleSelect is the builder for selecting fields of Puzzle entities.
type PuzzleSelect struct {
	*PuzzleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PuzzleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PuzzleQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *PuzzleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"math"

	"github.com/go-ent/ent"
	"github.com/go-ent/ent/dialect/sql"
	"github.com/go-ent/ent/dialect/sql/sqlgraph"
	"github.com/go-ent/ent/ent/gen/airdropuser"
	"github.com/go-ent/ent/ent/gen/predicate"
	"github.com/go-ent/ent/schema/field"
)

// AirdropUserQuery is the builder for querying AirdropUser entities.
type AirdropUserQuery struct {
	config
	ctx        *QueryContext
	order      []airdropuser.OrderOption
	inters     []Interceptor
	predicates []predicate.AirdropUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AirdropUserQuery builder.
func (auq *AirdropUserQuery) Where(ps ...predicate.AirdropUser) *AirdropUserQuery {
	auq.predicates = append(auq.predicates, ps...)
	return auq
}

// Limit the number of records to be returned by this query.
func (auq *AirdropUserQuery) Limit(limit int) *AirdropUserQuery {
	auq.ctx.Limit = &limit
	return auq
}

// Offset to start from.
func (auq *AirdropUserQuery) Offset(offset int) *AirdropUserQuery {
	auq.ctx.Offset = &offset
	return auq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auq *AirdropUserQuery) Unique(unique bool) *AirdropUserQuery {
	auq.ctx.Unique = &unique
	return auq
}

// Order specifies how the records should be ordered.
func (auq *AirdropUserQuery) Order(o ...airdropuser.OrderOption) *AirdropUserQuery {
	auq.order = append(auq.order, o...)
	return auq
}

// First returns the first AirdropUser entity from the query.
// Returns a *NotFoundError when no AirdropUser was found.
func (auq *AirdropUserQuery) First(ctx context.Context) (*AirdropUser, error) {
	nodes, err := auq.Limit(1).All(setContextOp(ctx, auq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{airdropuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auq *AirdropUserQuery) FirstX(ctx context.Context) *AirdropUser {
	node, err := auq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AirdropUser ID from the query.
// Returns a *NotFoundError when no AirdropUser ID was found.
func (auq *AirdropUserQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = auq.Limit(1).IDs(setContextOp(ctx, auq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{airdropuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auq *AirdropUserQuery) FirstIDX(ctx context.Context) uint {
	id, err := auq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AirdropUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AirdropUser entity is found.
// Returns a *NotFoundError when no AirdropUser entities are found.
func (auq *AirdropUserQuery) Only(ctx context.Context) (*AirdropUser, error) {
	nodes, err := auq.Limit(2).All(setContextOp(ctx, auq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{airdropuser.Label}
	default:
		return nil, &NotSingularError{airdropuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auq *AirdropUserQuery) OnlyX(ctx context.Context) *AirdropUser {
	node, err := auq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AirdropUser ID in the query.
// Returns a *NotSingularError when more than one AirdropUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (auq *AirdropUserQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = auq.Limit(2).IDs(setContextOp(ctx, auq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{airdropuser.Label}
	default:
		err = &NotSingularError{airdropuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auq *AirdropUserQuery) OnlyIDX(ctx context.Context) uint {
	id, err := auq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AirdropUsers.
func (auq *AirdropUserQuery) All(ctx context.Context) ([]*AirdropUser, error) {
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryAll)
	if err := auq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AirdropUser, *AirdropUserQuery]()
	return withInterceptors[[]*AirdropUser](ctx, auq, qr, auq.inters)
}

// AllX is like All, but panics if an error occurs.
func (auq *AirdropUserQuery) AllX(ctx context.Context) []*AirdropUser {
	nodes, err := auq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AirdropUser IDs.
func (auq *AirdropUserQuery) IDs(ctx context.Context) (ids []uint, err error) {
	if auq.ctx.Unique == nil && auq.path != nil {
		auq.Unique(true)
	}
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryIDs)
	if err = auq.Select(airdropuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auq *AirdropUserQuery) IDsX(ctx context.Context) []uint {
	ids, err := auq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auq *AirdropUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryCount)
	if err := auq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, auq, querierCount[*AirdropUserQuery](), auq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (auq *AirdropUserQuery) CountX(ctx context.Context) int {
	count, err := auq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auq *AirdropUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, auq.ctx, ent.OpQueryExist)
	switch _, err := auq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (auq *AirdropUserQuery) ExistX(ctx context.Context) bool {
	exist, err := auq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AirdropUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auq *AirdropUserQuery) Clone() *AirdropUserQuery {
	if auq == nil {
		return nil
	}
	return &AirdropUserQuery{
		config:     auq.config,
		ctx:        auq.ctx.Clone(),
		order:      append([]airdropuser.OrderOption{}, auq.order...),
		inters:     append([]Interceptor{}, auq.inters...),
		predicates: append([]predicate.AirdropUser{}, auq.predicates...),
		// clone intermediate query.
		sql:  auq.sql.Clone(),
		path: auq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Address string `json:"address,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AirdropUser.Query().
//		GroupBy(airdropuser.FieldAddress).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (auq *AirdropUserQuery) GroupBy(field string, fields ...string) *AirdropUserGroupBy {
	auq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AirdropUserGroupBy{build: auq}
	grbuild.flds = &auq.ctx.Fields
	grbuild.label = airdropuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Address string `json:"address,omitempty"`
//	}
//
//	client.AirdropUser.Query().
//		Select(airdropuser.FieldAddress).
//		Scan(ctx, &v)
func (auq *AirdropUserQuery) Select(fields ...string) *AirdropUserSelect {
	auq.ctx.Fields = append(auq.ctx.Fields, fields...)
	sbuild := &AirdropUserSelect{AirdropUserQuery: auq}
	sbuild.label = airdropuser.Label
	sbuild.flds, sbuild.scan = &auq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AirdropUserSelect configured with the given aggregations.
func (auq *AirdropUserQuery) Aggregate(fns ...AggregateFunc) *AirdropUserSelect {
	return auq.Select().Aggregate(fns...)
}

func (auq *AirdropUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range auq.inters {
		if inter == nil {
			return fmt.Errorf("gen: uninitialized interceptor (forgotten import gen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, auq); err != nil {
				return err
			}
		}
	}
	for _, f := range auq.ctx.Fields {
		if !airdropuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if auq.path != nil {
		prev, err := auq.path(ctx)
		if err != nil {
			return err
		}
		auq.sql = prev
	}
	return nil
}

func (auq *AirdropUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AirdropUser, error) {
	var (
		nodes = []*AirdropUser{}
		_spec = auq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AirdropUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AirdropUser{config: auq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, auq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auq *AirdropUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auq.querySpec()
	_spec.Node.Columns = auq.ctx.Fields
	if len(auq.ctx.Fields) > 0 {
		_spec.Unique = auq.ctx.Unique != nil && *auq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, auq.driver, _spec)
}

func (auq *AirdropUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(airdropuser.Table, airdropuser.Columns, sqlgraph.NewFieldSpec(airdropuser.FieldID, field.TypeUint))
	_spec.From = auq.sql
	if unique := auq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if auq.path != nil {
		_spec.Unique = true
	}
	if fields := auq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airdropuser.FieldID)
		for i := range fields {
			if fields[i] != airdropuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auq *AirdropUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auq.driver.Dialect())
	t1 := builder.Table(airdropuser.Table)
	columns := auq.ctx.Fields
	if len(columns) == 0 {
		columns = airdropuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auq.sql != nil {
		selector = auq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auq.ctx.Unique != nil && *auq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range auq.predicates {
		p(selector)
	}
	for _, p := range auq.order {
		p(selector)
	}
	if offset := auq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AirdropUserGroupBy is the group-by builder for AirdropUser entities.
type AirdropUserGroupBy struct {
	selector
	build *AirdropUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (augb *AirdropUserGroupBy) Aggregate(fns ...AggregateFunc) *AirdropUserGroupBy {
	augb.fns = append(augb.fns, fns...)
	return augb
}

// Scan applies the selector query and scans the result into the given value.
func (augb *AirdropUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, augb.build.ctx, ent.OpQueryGroupBy)
	if err := augb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AirdropUserQuery, *AirdropUserGroupBy](ctx, augb.build, augb, augb.build.inters, v)
}

func (augb *AirdropUserGroupBy) sqlScan(ctx context.Context, root *AirdropUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(augb.fns))
	for _, fn := range augb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*augb.flds)+len(augb.fns))
		for _, f := range *augb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*augb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := augb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AirdropUserSelect is the builder for selecting fields of AirdropUser entities.
type AirdropUserSelect struct {
	*AirdropUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aus *AirdropUserSelect) Aggregate(fns ...AggregateFunc) *AirdropUserSelect {
	aus.fns = append(aus.fns, fns...)
	return aus
}

// Scan applies the selector query and scans the result into the given value.
func (aus *AirdropUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aus.ctx, ent.OpQuerySelect)
	if err := aus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AirdropUserQuery, *AirdropUserSelect](ctx, aus.AirdropUserQuery, aus, aus.inters, v)
}

func (aus *AirdropUserSelect) sqlScan(ctx context.Context, root *AirdropUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aus.fns))
	for _, fn := range aus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"myapp/ent/carutility"
	"myapp/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CarUtilityDelete is the builder for deleting a CarUtility entity.
type CarUtilityDelete struct {
	config
	hooks    []Hook
	mutation *CarUtilityMutation
}

// Where appends a list predicates to the CarUtilityDelete builder.
func (cud *CarUtilityDelete) Where(ps ...predicate.CarUtility) *CarUtilityDelete {
	cud.mutation.Where(ps...)
	return cud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cud *CarUtilityDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cud.hooks) == 0 {
		affected, err = cud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarUtilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cud.mutation = mutation
			affected, err = cud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cud.hooks) - 1; i >= 0; i-- {
			if cud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cud *CarUtilityDelete) ExecX(ctx context.Context) int {
	n, err := cud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cud *CarUtilityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: carutility.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carutility.FieldID,
			},
		},
	}
	if ps := cud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cud.driver, _spec)
}

// CarUtilityDeleteOne is the builder for deleting a single CarUtility entity.
type CarUtilityDeleteOne struct {
	cud *CarUtilityDelete
}

// Exec executes the deletion query.
func (cudo *CarUtilityDeleteOne) Exec(ctx context.Context) error {
	n, err := cudo.cud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{carutility.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cudo *CarUtilityDeleteOne) ExecX(ctx context.Context) {
	cudo.cud.ExecX(ctx)
}

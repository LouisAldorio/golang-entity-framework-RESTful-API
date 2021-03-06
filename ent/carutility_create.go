// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myapp/ent/car"
	"myapp/ent/carutility"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CarUtilityCreate is the builder for creating a CarUtility entity.
type CarUtilityCreate struct {
	config
	mutation *CarUtilityMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cuc *CarUtilityCreate) SetName(s string) *CarUtilityCreate {
	cuc.mutation.SetName(s)
	return cuc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuc *CarUtilityCreate) SetNillableName(s *string) *CarUtilityCreate {
	if s != nil {
		cuc.SetName(*s)
	}
	return cuc
}

// AddCarIDs adds the "cars" edge to the Car entity by IDs.
func (cuc *CarUtilityCreate) AddCarIDs(ids ...int) *CarUtilityCreate {
	cuc.mutation.AddCarIDs(ids...)
	return cuc
}

// AddCars adds the "cars" edges to the Car entity.
func (cuc *CarUtilityCreate) AddCars(c ...*Car) *CarUtilityCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuc.AddCarIDs(ids...)
}

// Mutation returns the CarUtilityMutation object of the builder.
func (cuc *CarUtilityCreate) Mutation() *CarUtilityMutation {
	return cuc.mutation
}

// Save creates the CarUtility in the database.
func (cuc *CarUtilityCreate) Save(ctx context.Context) (*CarUtility, error) {
	var (
		err  error
		node *CarUtility
	)
	cuc.defaults()
	if len(cuc.hooks) == 0 {
		if err = cuc.check(); err != nil {
			return nil, err
		}
		node, err = cuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarUtilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuc.check(); err != nil {
				return nil, err
			}
			cuc.mutation = mutation
			if node, err = cuc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cuc.hooks) - 1; i >= 0; i-- {
			if cuc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cuc *CarUtilityCreate) SaveX(ctx context.Context) *CarUtility {
	v, err := cuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cuc *CarUtilityCreate) Exec(ctx context.Context) error {
	_, err := cuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuc *CarUtilityCreate) ExecX(ctx context.Context) {
	if err := cuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuc *CarUtilityCreate) defaults() {
	if _, ok := cuc.mutation.Name(); !ok {
		v := carutility.DefaultName
		cuc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuc *CarUtilityCreate) check() error {
	if _, ok := cuc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (cuc *CarUtilityCreate) sqlSave(ctx context.Context) (*CarUtility, error) {
	_node, _spec := cuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cuc *CarUtilityCreate) createSpec() (*CarUtility, *sqlgraph.CreateSpec) {
	var (
		_node = &CarUtility{config: cuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: carutility.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carutility.FieldID,
			},
		}
	)
	if value, ok := cuc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: carutility.FieldName,
		})
		_node.Name = value
	}
	if nodes := cuc.mutation.CarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   carutility.CarsTable,
			Columns: carutility.CarsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
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

// CarUtilityCreateBulk is the builder for creating many CarUtility entities in bulk.
type CarUtilityCreateBulk struct {
	config
	builders []*CarUtilityCreate
}

// Save creates the CarUtility entities in the database.
func (cucb *CarUtilityCreateBulk) Save(ctx context.Context) ([]*CarUtility, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cucb.builders))
	nodes := make([]*CarUtility, len(cucb.builders))
	mutators := make([]Mutator, len(cucb.builders))
	for i := range cucb.builders {
		func(i int, root context.Context) {
			builder := cucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarUtilityMutation)
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
					_, err = mutators[i+1].Mutate(root, cucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cucb *CarUtilityCreateBulk) SaveX(ctx context.Context) []*CarUtility {
	v, err := cucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cucb *CarUtilityCreateBulk) Exec(ctx context.Context) error {
	_, err := cucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cucb *CarUtilityCreateBulk) ExecX(ctx context.Context) {
	if err := cucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// DocumentUpdate is the builder for updating Document entities.
type DocumentUpdate struct {
	config
	hooks    []Hook
	mutation *DocumentMutation
}

// Where appends a list predicates to the DocumentUpdate builder.
func (du *DocumentUpdate) Where(ps ...predicate.Document) *DocumentUpdate {
	du.mutation.Where(ps...)
	return du
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (du *DocumentUpdate) AddAnnotationIDs(ids ...int) *DocumentUpdate {
	du.mutation.AddAnnotationIDs(ids...)
	return du
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (du *DocumentUpdate) AddAnnotations(a ...*Annotation) *DocumentUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.AddAnnotationIDs(ids...)
}

// Mutation returns the DocumentMutation object of the builder.
func (du *DocumentUpdate) Mutation() *DocumentMutation {
	return du.mutation
}

// ClearAnnotations clears all "annotations" edges to the Annotation entity.
func (du *DocumentUpdate) ClearAnnotations() *DocumentUpdate {
	du.mutation.ClearAnnotations()
	return du
}

// RemoveAnnotationIDs removes the "annotations" edge to Annotation entities by IDs.
func (du *DocumentUpdate) RemoveAnnotationIDs(ids ...int) *DocumentUpdate {
	du.mutation.RemoveAnnotationIDs(ids...)
	return du
}

// RemoveAnnotations removes "annotations" edges to Annotation entities.
func (du *DocumentUpdate) RemoveAnnotations(a ...*Annotation) *DocumentUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.RemoveAnnotationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DocumentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DocumentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DocumentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DocumentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DocumentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(document.Table, document.Columns, sqlgraph.NewFieldSpec(document.FieldID, field.TypeString))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if du.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   document.AnnotationsTable,
			Columns: []string{document.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedAnnotationsIDs(); len(nodes) > 0 && !du.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   document.AnnotationsTable,
			Columns: []string{document.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   document.AnnotationsTable,
			Columns: []string{document.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DocumentUpdateOne is the builder for updating a single Document entity.
type DocumentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DocumentMutation
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (duo *DocumentUpdateOne) AddAnnotationIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.AddAnnotationIDs(ids...)
	return duo
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (duo *DocumentUpdateOne) AddAnnotations(a ...*Annotation) *DocumentUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.AddAnnotationIDs(ids...)
}

// Mutation returns the DocumentMutation object of the builder.
func (duo *DocumentUpdateOne) Mutation() *DocumentMutation {
	return duo.mutation
}

// ClearAnnotations clears all "annotations" edges to the Annotation entity.
func (duo *DocumentUpdateOne) ClearAnnotations() *DocumentUpdateOne {
	duo.mutation.ClearAnnotations()
	return duo
}

// RemoveAnnotationIDs removes the "annotations" edge to Annotation entities by IDs.
func (duo *DocumentUpdateOne) RemoveAnnotationIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.RemoveAnnotationIDs(ids...)
	return duo
}

// RemoveAnnotations removes "annotations" edges to Annotation entities.
func (duo *DocumentUpdateOne) RemoveAnnotations(a ...*Annotation) *DocumentUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.RemoveAnnotationIDs(ids...)
}

// Where appends a list predicates to the DocumentUpdate builder.
func (duo *DocumentUpdateOne) Where(ps ...predicate.Document) *DocumentUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DocumentUpdateOne) Select(field string, fields ...string) *DocumentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Document entity.
func (duo *DocumentUpdateOne) Save(ctx context.Context) (*Document, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DocumentUpdateOne) SaveX(ctx context.Context) *Document {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DocumentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DocumentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DocumentUpdateOne) sqlSave(ctx context.Context) (_node *Document, err error) {
	_spec := sqlgraph.NewUpdateSpec(document.Table, document.Columns, sqlgraph.NewFieldSpec(document.FieldID, field.TypeString))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Document.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, document.FieldID)
		for _, f := range fields {
			if !document.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != document.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if duo.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   document.AnnotationsTable,
			Columns: []string{document.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedAnnotationsIDs(); len(nodes) > 0 && !duo.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   document.AnnotationsTable,
			Columns: []string{document.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   document.AnnotationsTable,
			Columns: []string{document.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Document{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}

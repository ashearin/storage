// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/node"
)

// Annotation is the model entity for the Annotation schema.
type Annotation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DocumentID holds the value of the "document_id" field.
	DocumentID *uuid.UUID `json:"-"`
	// NodeID holds the value of the "node_id" field.
	NodeID *uuid.UUID `json:"-"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// IsUnique holds the value of the "is_unique" field.
	IsUnique bool `json:"is_unique,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnnotationQuery when eager-loading is set.
	Edges        AnnotationEdges `json:"-"`
	selectValues sql.SelectValues
}

// AnnotationEdges holds the relations/edges for other nodes in the graph.
type AnnotationEdges struct {
	// Document holds the value of the document edge.
	Document *Document `json:"-"`
	// Node holds the value of the node edge.
	Node *Node `json:"-"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DocumentOrErr returns the Document value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnnotationEdges) DocumentOrErr() (*Document, error) {
	if e.Document != nil {
		return e.Document, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: document.Label}
	}
	return nil, &NotLoadedError{edge: "document"}
}

// NodeOrErr returns the Node value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnnotationEdges) NodeOrErr() (*Node, error) {
	if e.Node != nil {
		return e.Node, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: node.Label}
	}
	return nil, &NotLoadedError{edge: "node"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Annotation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case annotation.FieldDocumentID, annotation.FieldNodeID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case annotation.FieldIsUnique:
			values[i] = new(sql.NullBool)
		case annotation.FieldID:
			values[i] = new(sql.NullInt64)
		case annotation.FieldName, annotation.FieldValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Annotation fields.
func (a *Annotation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case annotation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case annotation.FieldDocumentID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field document_id", values[i])
			} else if value.Valid {
				a.DocumentID = new(uuid.UUID)
				*a.DocumentID = *value.S.(*uuid.UUID)
			}
		case annotation.FieldNodeID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				a.NodeID = new(uuid.UUID)
				*a.NodeID = *value.S.(*uuid.UUID)
			}
		case annotation.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case annotation.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				a.Value = value.String
			}
		case annotation.FieldIsUnique:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_unique", values[i])
			} else if value.Valid {
				a.IsUnique = value.Bool
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Annotation.
// This includes values selected through modifiers, order, etc.
func (a *Annotation) GetValue(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryDocument queries the "document" edge of the Annotation entity.
func (a *Annotation) QueryDocument() *DocumentQuery {
	return NewAnnotationClient(a.config).QueryDocument(a)
}

// QueryNode queries the "node" edge of the Annotation entity.
func (a *Annotation) QueryNode() *NodeQuery {
	return NewAnnotationClient(a.config).QueryNode(a)
}

// Update returns a builder for updating this Annotation.
// Note that you need to call Annotation.Unwrap() before calling this method if this Annotation
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Annotation) Update() *AnnotationUpdateOne {
	return NewAnnotationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Annotation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Annotation) Unwrap() *Annotation {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Annotation is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Annotation) String() string {
	var builder strings.Builder
	builder.WriteString("Annotation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	if v := a.DocumentID; v != nil {
		builder.WriteString("document_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := a.NodeID; v != nil {
		builder.WriteString("node_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(a.Value)
	builder.WriteString(", ")
	builder.WriteString("is_unique=")
	builder.WriteString(fmt.Sprintf("%v", a.IsUnique))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (a *Annotation) MarshalJSON() ([]byte, error) {
	type Alias Annotation
	return json.Marshal(&struct {
		*Alias
		AnnotationEdges
	}{
		Alias:           (*Alias)(a),
		AnnotationEdges: a.Edges,
	})
}

// Annotations is a parsable slice of Annotation.
type Annotations []*Annotation

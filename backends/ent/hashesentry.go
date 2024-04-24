// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/protobom/storage/backends/ent/externalreference"
	"github.com/protobom/storage/backends/ent/hashesentry"
	"github.com/protobom/storage/backends/ent/node"
)

// HashesEntry is the model entity for the HashesEntry schema.
type HashesEntry struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HashAlgorithmType holds the value of the "hash_algorithm_type" field.
	HashAlgorithmType hashesentry.HashAlgorithmType `json:"hash_algorithm_type,omitempty"`
	// HashData holds the value of the "hash_data" field.
	HashData string `json:"hash_data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HashesEntryQuery when eager-loading is set.
	Edges                     HashesEntryEdges `json:"edges"`
	external_reference_hashes *int
	node_hashes               *string
	selectValues              sql.SelectValues
}

// HashesEntryEdges holds the relations/edges for other nodes in the graph.
type HashesEntryEdges struct {
	// ExternalReferences holds the value of the external_references edge.
	ExternalReferences *ExternalReference `json:"external_references,omitempty"`
	// Nodes holds the value of the nodes edge.
	Nodes *Node `json:"nodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ExternalReferencesOrErr returns the ExternalReferences value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HashesEntryEdges) ExternalReferencesOrErr() (*ExternalReference, error) {
	if e.ExternalReferences != nil {
		return e.ExternalReferences, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: externalreference.Label}
	}
	return nil, &NotLoadedError{edge: "external_references"}
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HashesEntryEdges) NodesOrErr() (*Node, error) {
	if e.Nodes != nil {
		return e.Nodes, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: node.Label}
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HashesEntry) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hashesentry.FieldID:
			values[i] = new(sql.NullInt64)
		case hashesentry.FieldHashAlgorithmType, hashesentry.FieldHashData:
			values[i] = new(sql.NullString)
		case hashesentry.ForeignKeys[0]: // external_reference_hashes
			values[i] = new(sql.NullInt64)
		case hashesentry.ForeignKeys[1]: // node_hashes
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HashesEntry fields.
func (he *HashesEntry) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hashesentry.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			he.ID = int(value.Int64)
		case hashesentry.FieldHashAlgorithmType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash_algorithm_type", values[i])
			} else if value.Valid {
				he.HashAlgorithmType = hashesentry.HashAlgorithmType(value.String)
			}
		case hashesentry.FieldHashData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash_data", values[i])
			} else if value.Valid {
				he.HashData = value.String
			}
		case hashesentry.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field external_reference_hashes", value)
			} else if value.Valid {
				he.external_reference_hashes = new(int)
				*he.external_reference_hashes = int(value.Int64)
			}
		case hashesentry.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_hashes", values[i])
			} else if value.Valid {
				he.node_hashes = new(string)
				*he.node_hashes = value.String
			}
		default:
			he.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HashesEntry.
// This includes values selected through modifiers, order, etc.
func (he *HashesEntry) Value(name string) (ent.Value, error) {
	return he.selectValues.Get(name)
}

// QueryExternalReferences queries the "external_references" edge of the HashesEntry entity.
func (he *HashesEntry) QueryExternalReferences() *ExternalReferenceQuery {
	return NewHashesEntryClient(he.config).QueryExternalReferences(he)
}

// QueryNodes queries the "nodes" edge of the HashesEntry entity.
func (he *HashesEntry) QueryNodes() *NodeQuery {
	return NewHashesEntryClient(he.config).QueryNodes(he)
}

// Update returns a builder for updating this HashesEntry.
// Note that you need to call HashesEntry.Unwrap() before calling this method if this HashesEntry
// was returned from a transaction, and the transaction was committed or rolled back.
func (he *HashesEntry) Update() *HashesEntryUpdateOne {
	return NewHashesEntryClient(he.config).UpdateOne(he)
}

// Unwrap unwraps the HashesEntry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (he *HashesEntry) Unwrap() *HashesEntry {
	_tx, ok := he.config.driver.(*txDriver)
	if !ok {
		panic("ent: HashesEntry is not a transactional entity")
	}
	he.config.driver = _tx.drv
	return he
}

// String implements the fmt.Stringer.
func (he *HashesEntry) String() string {
	var builder strings.Builder
	builder.WriteString("HashesEntry(")
	builder.WriteString(fmt.Sprintf("id=%v, ", he.ID))
	builder.WriteString("hash_algorithm_type=")
	builder.WriteString(fmt.Sprintf("%v", he.HashAlgorithmType))
	builder.WriteString(", ")
	builder.WriteString("hash_data=")
	builder.WriteString(he.HashData)
	builder.WriteByte(')')
	return builder.String()
}

// HashesEntries is a parsable slice of HashesEntry.
type HashesEntries []*HashesEntry

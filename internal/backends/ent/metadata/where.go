// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package metadata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContainsFold(FieldID, id))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldVersion, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldName, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldDate, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldComment, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContainsFold(FieldVersion, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContainsFold(FieldName, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldDate, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContainsFold(FieldComment, v))
}

// HasTools applies the HasEdge predicate on the "tools" edge.
func HasTools() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ToolsTable, ToolsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToolsWith applies the HasEdge predicate on the "tools" edge with a given conditions (other predicates).
func HasToolsWith(preds ...predicate.Tool) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newToolsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthors applies the HasEdge predicate on the "authors" edge.
func HasAuthors() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuthorsTable, AuthorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorsWith applies the HasEdge predicate on the "authors" edge with a given conditions (other predicates).
func HasAuthorsWith(preds ...predicate.Person) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newAuthorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDocumentTypes applies the HasEdge predicate on the "document_types" edge.
func HasDocumentTypes() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DocumentTypesTable, DocumentTypesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentTypesWith applies the HasEdge predicate on the "document_types" edge with a given conditions (other predicates).
func HasDocumentTypesWith(preds ...predicate.DocumentType) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newDocumentTypesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDocument applies the HasEdge predicate on the "document" edge.
func HasDocument() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, DocumentTable, DocumentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentWith applies the HasEdge predicate on the "document" edge with a given conditions (other predicates).
func HasDocumentWith(preds ...predicate.Document) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newDocumentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(sql.NotPredicates(p))
}

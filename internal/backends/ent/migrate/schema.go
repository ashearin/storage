// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DocumentsColumns holds the columns for the "documents" table.
	DocumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// DocumentsTable holds the schema information for the "documents" table.
	DocumentsTable = &schema.Table{
		Name:       "documents",
		Columns:    DocumentsColumns,
		PrimaryKey: []*schema.Column{DocumentsColumns[0]},
	}
	// DocumentTypesColumns holds the columns for the "document_types" table.
	DocumentTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"OTHER", "DESIGN", "SOURCE", "BUILD", "ANALYZED", "DEPLOYED", "RUNTIME", "DISCOVERY", "DECOMISSION"}},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata_id", Type: field.TypeString, Nullable: true},
	}
	// DocumentTypesTable holds the schema information for the "document_types" table.
	DocumentTypesTable = &schema.Table{
		Name:       "document_types",
		Columns:    DocumentTypesColumns,
		PrimaryKey: []*schema.Column{DocumentTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "document_types_metadata_document_types",
				Columns:    []*schema.Column{DocumentTypesColumns[4]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_document_types",
				Unique:  true,
				Columns: []*schema.Column{DocumentTypesColumns[4], DocumentTypesColumns[1], DocumentTypesColumns[2], DocumentTypesColumns[3]},
			},
		},
	}
	// EdgeTypesColumns holds the columns for the "edge_types" table.
	EdgeTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "amends", "ancestor", "buildDependency", "buildTool", "contains", "contained_by", "copy", "dataFile", "dependencyManifest", "dependsOn", "dependencyOf", "descendant", "describes", "describedBy", "devDependency", "devTool", "distributionArtifact", "documentation", "dynamicLink", "example", "expandedFromArchive", "fileAdded", "fileDeleted", "fileModified", "generates", "generatedFrom", "metafile", "optionalComponent", "optionalDependency", "other", "packages", "patch", "prerequisite", "prerequisiteFor", "providedDependency", "requirementFor", "runtimeDependency", "specificationFor", "staticLink", "test", "testCase", "testDependency", "testTool", "variant"}},
		{Name: "node_id", Type: field.TypeString},
		{Name: "to_node_id", Type: field.TypeString},
	}
	// EdgeTypesTable holds the schema information for the "edge_types" table.
	EdgeTypesTable = &schema.Table{
		Name:       "edge_types",
		Columns:    EdgeTypesColumns,
		PrimaryKey: []*schema.Column{EdgeTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "edge_types_nodes_from",
				Columns:    []*schema.Column{EdgeTypesColumns[2]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "edge_types_nodes_to",
				Columns:    []*schema.Column{EdgeTypesColumns[3]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_edge_types",
				Unique:  true,
				Columns: []*schema.Column{EdgeTypesColumns[1], EdgeTypesColumns[2], EdgeTypesColumns[3]},
			},
			{
				Name:    "edgetype_node_id_to_node_id",
				Unique:  true,
				Columns: []*schema.Column{EdgeTypesColumns[2], EdgeTypesColumns[3]},
			},
		},
	}
	// ExternalReferencesColumns holds the columns for the "external_references" table.
	ExternalReferencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "authority", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "ATTESTATION", "BINARY", "BOM", "BOWER", "BUILD_META", "BUILD_SYSTEM", "CERTIFICATION_REPORT", "CHAT", "CODIFIED_INFRASTRUCTURE", "COMPONENT_ANALYSIS_REPORT", "CONFIGURATION", "DISTRIBUTION_INTAKE", "DOCUMENTATION", "DOWNLOAD", "DYNAMIC_ANALYSIS_REPORT", "EOL_NOTICE", "EVIDENCE", "EXPORT_CONTROL_ASSESSMENT", "FORMULATION", "FUNDING", "ISSUE_TRACKER", "LICENSE", "LOG", "MAILING_LIST", "MATURITY_REPORT", "MAVEN_CENTRAL", "METRICS", "MODEL_CARD", "NPM", "NUGET", "OTHER", "POAM", "PRIVACY_ASSESSMENT", "PRODUCT_METADATA", "PURCHASE_ORDER", "QUALITY_ASSESSMENT_REPORT", "QUALITY_METRICS", "RELEASE_HISTORY", "RELEASE_NOTES", "RISK_ASSESSMENT", "RUNTIME_ANALYSIS_REPORT", "SECURE_SOFTWARE_ATTESTATION", "SECURITY_ADVERSARY_MODEL", "SECURITY_ADVISORY", "SECURITY_CONTACT", "SECURITY_FIX", "SECURITY_OTHER", "SECURITY_PENTEST_REPORT", "SECURITY_POLICY", "SECURITY_SWID", "SECURITY_THREAT_MODEL", "SOCIAL", "SOURCE_ARTIFACT", "STATIC_ANALYSIS_REPORT", "SUPPORT", "VCS", "VULNERABILITY_ASSERTION", "VULNERABILITY_DISCLOSURE_REPORT", "VULNERABILITY_EXPLOITABILITY_ASSESSMENT", "WEBSITE"}},
		{Name: "node_id", Type: field.TypeString, Nullable: true},
	}
	// ExternalReferencesTable holds the schema information for the "external_references" table.
	ExternalReferencesTable = &schema.Table{
		Name:       "external_references",
		Columns:    ExternalReferencesColumns,
		PrimaryKey: []*schema.Column{ExternalReferencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "external_references_nodes_external_references",
				Columns:    []*schema.Column{ExternalReferencesColumns[5]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_external_references",
				Unique:  true,
				Columns: []*schema.Column{ExternalReferencesColumns[5], ExternalReferencesColumns[1], ExternalReferencesColumns[4]},
			},
		},
	}
	// HashesEntriesColumns holds the columns for the "hashes_entries" table.
	HashesEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash_algorithm_type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "MD5", "SHA1", "SHA256", "SHA384", "SHA512", "SHA3_256", "SHA3_384", "SHA3_512", "BLAKE2B_256", "BLAKE2B_384", "BLAKE2B_512", "BLAKE3", "MD2", "ADLER32", "MD4", "MD6", "SHA224"}},
		{Name: "hash_data", Type: field.TypeString},
		{Name: "external_reference_id", Type: field.TypeInt, Nullable: true},
		{Name: "node_id", Type: field.TypeString, Nullable: true},
	}
	// HashesEntriesTable holds the schema information for the "hashes_entries" table.
	HashesEntriesTable = &schema.Table{
		Name:       "hashes_entries",
		Columns:    HashesEntriesColumns,
		PrimaryKey: []*schema.Column{HashesEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hashes_entries_external_references_hashes",
				Columns:    []*schema.Column{HashesEntriesColumns[3]},
				RefColumns: []*schema.Column{ExternalReferencesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "hashes_entries_nodes_hashes",
				Columns:    []*schema.Column{HashesEntriesColumns[4]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_hashes_entries",
				Unique:  true,
				Columns: []*schema.Column{HashesEntriesColumns[3], HashesEntriesColumns[4], HashesEntriesColumns[1], HashesEntriesColumns[2]},
			},
		},
	}
	// IdentifiersEntriesColumns holds the columns for the "identifiers_entries" table.
	IdentifiersEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "software_identifier_type", Type: field.TypeEnum, Enums: []string{"UNKNOWN_IDENTIFIER_TYPE", "PURL", "CPE22", "CPE23", "GITOID"}},
		{Name: "software_identifier_value", Type: field.TypeString},
		{Name: "node_id", Type: field.TypeString, Nullable: true},
	}
	// IdentifiersEntriesTable holds the schema information for the "identifiers_entries" table.
	IdentifiersEntriesTable = &schema.Table{
		Name:       "identifiers_entries",
		Columns:    IdentifiersEntriesColumns,
		PrimaryKey: []*schema.Column{IdentifiersEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "identifiers_entries_nodes_identifiers",
				Columns:    []*schema.Column{IdentifiersEntriesColumns[3]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_identifiers_entries",
				Unique:  true,
				Columns: []*schema.Column{IdentifiersEntriesColumns[3], IdentifiersEntriesColumns[1], IdentifiersEntriesColumns[2]},
			},
		},
	}
	// MetadataColumns holds the columns for the "metadata" table.
	MetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "comment", Type: field.TypeString},
	}
	// MetadataTable holds the schema information for the "metadata" table.
	MetadataTable = &schema.Table{
		Name:       "metadata",
		Columns:    MetadataColumns,
		PrimaryKey: []*schema.Column{MetadataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "metadata_documents_metadata",
				Columns:    []*schema.Column{MetadataColumns[0]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_metadata",
				Unique:  true,
				Columns: []*schema.Column{MetadataColumns[0], MetadataColumns[1], MetadataColumns[2]},
			},
		},
	}
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"PACKAGE", "FILE"}},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "file_name", Type: field.TypeString},
		{Name: "url_home", Type: field.TypeString},
		{Name: "url_download", Type: field.TypeString},
		{Name: "licenses", Type: field.TypeJSON},
		{Name: "license_concluded", Type: field.TypeString},
		{Name: "license_comments", Type: field.TypeString},
		{Name: "copyright", Type: field.TypeString},
		{Name: "source_info", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "release_date", Type: field.TypeTime},
		{Name: "build_date", Type: field.TypeTime},
		{Name: "valid_until_date", Type: field.TypeTime},
		{Name: "attribution", Type: field.TypeJSON},
		{Name: "file_types", Type: field.TypeJSON},
		{Name: "node_list_id", Type: field.TypeInt, Nullable: true},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "nodes_node_lists_nodes",
				Columns:    []*schema.Column{NodesColumns[20]},
				RefColumns: []*schema.Column{NodeListsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_nodes",
				Unique:  true,
				Columns: []*schema.Column{NodesColumns[0], NodesColumns[20]},
			},
		},
	}
	// NodeListsColumns holds the columns for the "node_lists" table.
	NodeListsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "root_elements", Type: field.TypeJSON},
		{Name: "document_id", Type: field.TypeString, Unique: true},
	}
	// NodeListsTable holds the schema information for the "node_lists" table.
	NodeListsTable = &schema.Table{
		Name:       "node_lists",
		Columns:    NodeListsColumns,
		PrimaryKey: []*schema.Column{NodeListsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_lists_documents_node_list",
				Columns:    []*schema.Column{NodeListsColumns[2]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_org", Type: field.TypeBool},
		{Name: "email", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "metadata_id", Type: field.TypeString, Nullable: true},
		{Name: "node_suppliers", Type: field.TypeString, Nullable: true},
		{Name: "node_id", Type: field.TypeString, Nullable: true},
		{Name: "person_contacts", Type: field.TypeInt, Nullable: true},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "persons_metadata_authors",
				Columns:    []*schema.Column{PersonsColumns[6]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "persons_nodes_suppliers",
				Columns:    []*schema.Column{PersonsColumns[7]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "persons_nodes_originators",
				Columns:    []*schema.Column{PersonsColumns[8]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "persons_persons_contacts",
				Columns:    []*schema.Column{PersonsColumns[9]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_person_metadata_id",
				Unique:  true,
				Columns: []*schema.Column{PersonsColumns[6], PersonsColumns[1], PersonsColumns[2], PersonsColumns[3], PersonsColumns[4], PersonsColumns[5]},
				Annotation: &entsql.IndexAnnotation{
					Where: "metadata_id IS NOT NULL AND node_id IS NULL",
				},
			},
			{
				Name:    "idx_person_node_id",
				Unique:  true,
				Columns: []*schema.Column{PersonsColumns[8], PersonsColumns[1], PersonsColumns[2], PersonsColumns[3], PersonsColumns[4], PersonsColumns[5]},
				Annotation: &entsql.IndexAnnotation{
					Where: "metadata_id IS NULL AND node_id IS NOT NULL",
				},
			},
		},
	}
	// PurposesColumns holds the columns for the "purposes" table.
	PurposesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "primary_purpose", Type: field.TypeEnum, Enums: []string{"UNKNOWN_PURPOSE", "APPLICATION", "ARCHIVE", "BOM", "CONFIGURATION", "CONTAINER", "DATA", "DEVICE", "DEVICE_DRIVER", "DOCUMENTATION", "EVIDENCE", "EXECUTABLE", "FILE", "FIRMWARE", "FRAMEWORK", "INSTALL", "LIBRARY", "MACHINE_LEARNING_MODEL", "MANIFEST", "MODEL", "MODULE", "OPERATING_SYSTEM", "OTHER", "PATCH", "PLATFORM", "REQUIREMENT", "SOURCE", "SPECIFICATION", "TEST"}},
		{Name: "node_id", Type: field.TypeString, Nullable: true},
	}
	// PurposesTable holds the schema information for the "purposes" table.
	PurposesTable = &schema.Table{
		Name:       "purposes",
		Columns:    PurposesColumns,
		PrimaryKey: []*schema.Column{PurposesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "purposes_nodes_primary_purpose",
				Columns:    []*schema.Column{PurposesColumns[2]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_purposes",
				Unique:  true,
				Columns: []*schema.Column{PurposesColumns[2], PurposesColumns[1]},
			},
		},
	}
	// ToolsColumns holds the columns for the "tools" table.
	ToolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "vendor", Type: field.TypeString},
		{Name: "metadata_id", Type: field.TypeString, Nullable: true},
	}
	// ToolsTable holds the schema information for the "tools" table.
	ToolsTable = &schema.Table{
		Name:       "tools",
		Columns:    ToolsColumns,
		PrimaryKey: []*schema.Column{ToolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tools_metadata_tools",
				Columns:    []*schema.Column{ToolsColumns[4]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_tools",
				Unique:  true,
				Columns: []*schema.Column{ToolsColumns[4], ToolsColumns[1], ToolsColumns[2], ToolsColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DocumentsTable,
		DocumentTypesTable,
		EdgeTypesTable,
		ExternalReferencesTable,
		HashesEntriesTable,
		IdentifiersEntriesTable,
		MetadataTable,
		NodesTable,
		NodeListsTable,
		PersonsTable,
		PurposesTable,
		ToolsTable,
	}
)

func init() {
	DocumentTypesTable.ForeignKeys[0].RefTable = MetadataTable
	EdgeTypesTable.ForeignKeys[0].RefTable = NodesTable
	EdgeTypesTable.ForeignKeys[1].RefTable = NodesTable
	ExternalReferencesTable.ForeignKeys[0].RefTable = NodesTable
	HashesEntriesTable.ForeignKeys[0].RefTable = ExternalReferencesTable
	HashesEntriesTable.ForeignKeys[1].RefTable = NodesTable
	IdentifiersEntriesTable.ForeignKeys[0].RefTable = NodesTable
	MetadataTable.ForeignKeys[0].RefTable = DocumentsTable
	NodesTable.ForeignKeys[0].RefTable = NodeListsTable
	NodeListsTable.ForeignKeys[0].RefTable = DocumentsTable
	PersonsTable.ForeignKeys[0].RefTable = MetadataTable
	PersonsTable.ForeignKeys[1].RefTable = NodesTable
	PersonsTable.ForeignKeys[2].RefTable = NodesTable
	PersonsTable.ForeignKeys[3].RefTable = PersonsTable
	PurposesTable.ForeignKeys[0].RefTable = NodesTable
	ToolsTable.ForeignKeys[0].RefTable = MetadataTable
}

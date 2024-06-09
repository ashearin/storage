// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

	sqlite "github.com/glebarez/go-sqlite"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
)

// Backend implements the protobom.pkg.storage.Backend interface.
type Backend struct {
	client *ent.Client
	ctx    context.Context

	// Options is the set of options common to all ent Backends.
	Options *BackendOptions
}

var _ storage.Backend = (*Backend)(nil)

func NewBackend(opts ...Option) *Backend {
	backend := &Backend{
		Options: NewBackendOptions(),
	}

	for _, opt := range opts {
		opt(backend)
	}

	return backend
}

func (backend *Backend) InitClient() error {
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	// Register the SQLite driver as "sqlite3".
	if !slices.Contains(sql.Drivers(), "sqlite3") {
		sqlite.RegisterAsSQLITE3()
	}

	clientOpts := []ent.Option{}
	if backend.Options.Debug {
		clientOpts = append(clientOpts, ent.Debug())
	}

	client, err := ent.Open("sqlite3", backend.Options.DatabaseFile+dsnParams, clientOpts...)
	if err != nil {
		return fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	backend.client = client
	backend.ctx = ent.NewContext(context.Background(), client)

	// Run the auto migration tool.
	if err := backend.client.Schema.Create(backend.ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %w", err)
	}

	return nil
}

func (backend *Backend) CloseClient() {
	backend.client.Close()
}

func (backend *Backend) Debug() *Backend {
	backend.Options.Debug = true

	return backend
}

func (backend *Backend) WithBackendOptions(opts *BackendOptions) *Backend {
	backend.Options = opts

	return backend
}

func (backend *Backend) WithDatabaseFile(file string) *Backend {
	backend.Options.DatabaseFile = file

	return backend
}

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package asset

import (
	"database/sql"
	"time"
)

// SQLRepository provides a SQL backed storage facility for financial market data.
type SQLRepository struct {
	// db is the database connection.
	db *sql.DB

	// dialect is the database dialect to use.
	dialect SQLRepositoryDialect

	// assetsQuery is the prepared assets query.
	assetsQuery *sql.Stmt

	// getSinceQuery is the prepared get since query.
	getSinceQuery *sql.Stmt

	// lastDateQuery is the prepared last date query.
	lastDateQuery *sql.Stmt

	// appendQuery is the prepared append query.
	appendQuery *sql.Stmt
}

// NewSQLRepository takes a database driver, URL, and dialect for the asset repository and connects to it.
func NewSQLRepository(dbDriver, dbURL string, dialect SQLRepositoryDialect) (*SQLRepository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the database connection.
func (s *SQLRepository) Close() error { _ = "STUB: not implemented"; return nil }

// Assets returns the names of all assets in the respository.
func (s *SQLRepository) Assets() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Get attempts to return a channel of snapshots for the asset with the given name.
func (s *SQLRepository) Get(name string) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSince attempts to return a channel of snapshots for the asset with the given name since the given date.
func (s *SQLRepository) GetSince(name string, date time.Time) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastDate returns the date of the last snapshot for the asset with the given name.
func (s *SQLRepository) LastDate(name string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Append adds the given snapshots to the asset with the given name.
func (s *SQLRepository) Append(name string, snapshots <-chan *Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// Drop drops the snapshots table.
func (s *SQLRepository) Drop() error { _ = "STUB: not implemented"; return nil }

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package asset

import (
	"time"
)

// InMemoryRepository stores and retrieves asset snapshots using
// an in memory storage.
type InMemoryRepository struct {
	// storage is the in memory storage for assets.
	storage map[string][]*Snapshot
}

// NewInMemoryRepository initializes an in memory repository.
func NewInMemoryRepository() *InMemoryRepository { _ = "STUB: not implemented"; return nil }

// Assets returns the names of all assets in the repository.
func (r *InMemoryRepository) Assets() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Get attempts to return a channel of snapshots for the asset with the given name.
func (r *InMemoryRepository) Get(name string) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSince attempts to return a channel of snapshots for the asset with the given name since the given date.
func (r *InMemoryRepository) GetSince(name string, date time.Time) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastDate returns the date of the last snapshot for the asset with the given name.
func (r *InMemoryRepository) LastDate(name string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Append adds the given snapshows to the asset with the given name.
func (r *InMemoryRepository) Append(name string, snapshots <-chan *Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

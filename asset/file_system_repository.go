// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package asset

import (
	"time"

	"github.com/cinar/indicator/v2/helper"
)

// FileSystemRepository stores and retrieves asset snapshots using
// the local file system.
type FileSystemRepository struct {
	// base is the root directory where asset snapshots are stored.
	base string

	// csvOptions are the CSV options used for reading and writing snapshots.
	csvOptions []helper.CsvOption[Snapshot]
}

// NewFileSystemRepository initializes a file system repository with
// the given base directory and the CSV options.
func NewFileSystemRepository(base string, csvOptions ...helper.CsvOption[Snapshot]) *FileSystemRepository {
	_ = "STUB: not implemented"
	return nil
}

// Assets returns the names of all assets in the repository.
func (r *FileSystemRepository) Assets() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get attempts to return a channel of snapshots for the asset with the given name.
func (r *FileSystemRepository) Get(name string) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSince attempts to return a channel of snapshots for the asset with the given name since the given date.
func (r *FileSystemRepository) GetSince(name string, date time.Time) (<-chan *Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastDate returns the date of the last snapshot for the asset with the given name.
func (r *FileSystemRepository) LastDate(name string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Append adds the given snapshows to the asset with the given name.
func (r *FileSystemRepository) Append(name string, snapshots <-chan *Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// getCsvFileName gets the CSV file name for the given asset name.
func (r *FileSystemRepository) getCsvFileName(name string) string {
	_ = "STUB: not implemented"
	return ""
}

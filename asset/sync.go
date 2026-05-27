// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package asset

import (
	"log/slog"
	"time"
)

const (
	// DefaultSyncWorkers is the default number of workers to use to synchronize.
	DefaultSyncWorkers = 1

	// DefaultSyncDelay is the default delay in seconds between each get request.
	DefaultSyncDelay = 5
)

// Sync represents the configuration parameters for synchronizing assets between repositories.
type Sync struct {
	// Number of workers to use.
	Workers int

	// Delay between repository get requests to minimize the load to the remote server.
	Delay int

	// Assets is the name of the assets to be synced. If it is empty, all assets in the target repository
	// will be synced instead.
	Assets []string

	// Logger is the slog logger instance.
	Logger *slog.Logger
}

// NewSync function initializes a new sync instance with the default parameters.
func NewSync() *Sync { _ = "STUB: not implemented"; return nil }

// Run synchronizes assets between the source and target repositories using multi-worker concurrency.
func (s *Sync) Run(source, target Repository, defaultStartDate time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

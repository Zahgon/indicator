// Package backtest contains the backtest functions.
//
// This package belongs to the Indicator project. Indicator is
// a Golang module that supplies a variety of technical
// indicators, strategies, and a backtesting framework
// for analysis.
//
// # License
//
//	Copyright (c) 2021-2026 Onur Cinar.
//	The source code is provided under GNU AGPLv3 License.
//	https://github.com/cinar/indicator
//
// # Disclaimer
//
// The information provided on this project is strictly for
// informational purposes and is not to be construed as
// advice or solicitation to buy or sell any security.
package backtest

import (
	"log/slog"
	"sync"

	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/strategy"
)

const (
	// DefaultBacktestWorkers is the default number of backtest workers.
	DefaultBacktestWorkers = 1

	// DefaultLastDays is the default number of days backtest should go back.
	DefaultLastDays = 365
)

// Backtest function rigorously evaluates the potential performance of the
// specified strategies applied to a defined set of assets. It generates
// comprehensive visual representations for each strategy-asset pairing.
type Backtest struct {
	// repository is the repository to retrieve the assets from.
	repository asset.Repository

	// report is the report writer for the backtest.
	report Report

	// Names is the names of the assets to backtest.
	Names []string

	// Strategies is the list of strategies to apply.
	Strategies []strategy.Strategy

	// Workers is the number of concurrent workers.
	Workers int

	// LastDays is the number of days backtest should go back.
	LastDays int

	// Logger is the slog logger instance.
	Logger *slog.Logger
}

// NewBacktest function initializes a new backtest instance.
func NewBacktest(repository asset.Repository, report Report) *Backtest {
	_ = "STUB: not implemented"
	return nil
}

// Run executes a comprehensive performance evaluation of the designated strategies,
// applied to a specified collection of assets. In the absence of explicitly defined
// assets, encompasses all assets within the repository. Likewise, in the absence of
// explicitly defined strategies, encompasses all the registered strategies.
func (b *Backtest) Run() error {
	_ = "STUB: not implemented"
	// When asset names are absent, considers all assets within the provided repository for evaluation.
	return nil
}

// When strategies are absent, considers all strategies.

// Begin report.

// Run the backtest workers.

// Wait for all workers to finish.

// End report.

// worker is a backtesting worker that concurrently executes backtests for individual
// assets. It receives asset names from the provided channel, and performs backtests
// using the given strategies.
func (b *Backtest) worker(names <-chan string, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// We don't expect the snapshots to be a stream during backtesting.

// Backtesting asset has begun.

// Backtest strategies on the given asset.

// Backtesting asset had ended

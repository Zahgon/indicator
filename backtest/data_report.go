// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package backtest

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/strategy"
)

// DataStrategyResult is the strategy result.
type DataStrategyResult struct {
	// Asset is the asset name.
	Asset string

	// Strategy is the strategy instnace.
	Strategy strategy.Strategy

	// Outcome is the strategy outcome.
	Outcome float64

	// Action is the final action recommended by the strategy.
	Action strategy.Action

	// Transactions are the action recommendations.
	Transactions []strategy.Action
}

// DataReport is the bactest data report enablign programmatic access to the backtest results.
type DataReport struct {
	// Results are the backtest results for the assets.
	Results map[string][]*DataStrategyResult
}

// NewDataReport initializes a new data report instance.
func NewDataReport() *DataReport { _ = "STUB: not implemented"; return nil }

// Begin is called when the backtest begins.
func (*DataReport) Begin(_ []string, _ []strategy.Strategy) error {
	_ = "STUB: not implemented"

	// AssetBegin is called when backtesting for the given asset begins.
	return nil
}

func (d *DataReport) AssetBegin(name string, strategies []strategy.Strategy) error {
	_ = "STUB: not implemented"
	return nil
}

// Write writes the given strategy actions and outomes to the report.
func (d *DataReport) Write(assetName string, currentStrategy strategy.Strategy, snapshots <-chan *asset.Snapshot, actions <-chan strategy.Action, outcomes <-chan float64) error {
	_ = "STUB: not implemented"
	return nil
}

// AssetEnd is called when backtesting for the given asset ends.
func (*DataReport) AssetEnd(_ string) error {
	_ = "STUB: not implemented"

	// End is called when the backtest ends.
	return nil
}

func (*DataReport) End() error { _ = "STUB: not implemented"; return nil }

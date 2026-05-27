// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package decorator

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
)

// StopLossStrategy prevents a loss by recommending a sell action when the assets drop below the given threshold.
type StopLossStrategy struct {
	// InnertStrategy is the inner strategy.
	InnertStrategy strategy.Strategy

	// Percentage is the loss threshold in percentage.
	Percentage float64
}

// NewStopLossStrategy function initializes a new stop loss strategy instance.
func NewStopLossStrategy(innerStrategy strategy.Strategy, percentage float64) *StopLossStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (s *StopLossStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (s *StopLossStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// If action is Buy and the asset is not yet bought, buy it as recommended.

// If asset is bought and action is sell or closing is less than or equal to stop loss at, recommend sell.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (s *StopLossStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	return nil
}

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultTrimaStrategyShortPeriod is the first TRIMA period.
	DefaultTrimaStrategyShortPeriod = 20

	// DefaultTrimaStrategyLongPeriod is the second TRIMA period.
	DefaultTrimaStrategyLongPeriod = 50
)

// TrimaStrategy represents the configuration parameters for calculating the TRIMA strategy.
// A bullish cross occurs when the short TRIMA moves above the long TRIMA.
// A bearish cross occurs when the short TRIMA moves below the long TRIME.
type TrimaStrategy struct {
	// Trima1 represents the configuration parameters for calculating the short TRIMA.
	Short *trend.Trima[float64]

	// Trima2 represents the configuration parameters for calculating the long TRIMA.
	Long *trend.Trima[float64]
}

// NewTrimaStrategy function initializes a new TRIMA strategy instance
// with the default parameters.
func NewTrimaStrategy() *TrimaStrategy { _ = "STUB: not implemented"; return nil }

// Name returns the name of the strategy.
func (*TrimaStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a
// stream of actionable recommendations.
func (t *TrimaStrategy) Compute(c <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// TRIMA starts only after a full periods for each EMA used.

// Report processes the provided asset snapshots and generates a
// report annotated with the recommended actions.
func (t *TrimaStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> closings[0] -> shorts
	//                 closings[1] -> longs
	//                 closings[2] -> closings
	// snapshots[2] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

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
	// DefaultDemaStrategyPeriod1 is the first DEMA period.
	DefaultDemaStrategyPeriod1 = 5

	// DefaultDemaStrategyPeriod2 is the second DEMA period.
	DefaultDemaStrategyPeriod2 = 35
)

// DemaStrategy represents the configuration parameters for calculating the DEMA strategy.
// A bullish cross occurs when DEMA with 5 days period moves above DEMA with 35 days period.
// A bearish cross occurs when DEMA with 35 days period moves above DEMA With 5 days period.
type DemaStrategy struct {
	// Dema1 represents the configuration parameters for
	// calculating the first DEMA.
	Dema1 *trend.Dema[float64]

	// Dema2 represents the configuration parameters for
	// calculating the second DEMA.
	Dema2 *trend.Dema[float64]
}

// NewDemaStrategy function initializes a new DEMA strategy instance
// with the default parameters.
func NewDemaStrategy() *DemaStrategy { _ = "STUB: not implemented"; return nil }

// Name returns the name of the strategy.
func (*DemaStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a
// stream of actionable recommendations.
func (d *DemaStrategy) Compute(c <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// DEMA starts only after a full periods for each EMA used.

// Report processes the provided asset snapshots and generates a
// report annotated with the recommended actions.
func (d *DemaStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> closings[0] -> demas1
	//                 closings[1] -> demas2
	//                 closings[2] -> closings
	// snapshots[2] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

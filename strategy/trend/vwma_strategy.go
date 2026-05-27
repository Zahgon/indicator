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
	// DefaultVwmaStrategyPeriod is the default VWMA period.
	DefaultVwmaStrategyPeriod = 20
)

// VwmaStrategy represents the configuration parameters for calculating the VWMA strategy.
// The VwmaStrategy function uses SMA and VWMA indicators to provide a BUY action when
// VWMA is above SMA, and a SELL signal when VWMA is below SMA, a HOLD otherwse.
type VwmaStrategy struct {
	// VWMA indicator.
	Vwma *trend.Vwma[float64]

	// SMA indicator.
	Sma *trend.Sma[float64]
}

// NewVwmaStrategy function initializes a new VWMA strategy instance with the default parameters.
func NewVwmaStrategy() *VwmaStrategy { _ = "STUB: not implemented"; return nil }

// Name returns the name of the strategy.
func (*VwmaStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (v *VwmaStrategy) Compute(c <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// VWMA starts only after a full period.

// Report processes the provided asset snapshots and generates a
// report annotated with the recommended actions.
func (v *VwmaStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> closings
	// snapshots[2] -> sma
	//                 vwma
	// snapshots[3] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

// calculateSmaAndVwma calculates the SMA and VWMA using the given channel of snapshots.
func (v *VwmaStrategy) calculateSmaAndVwma(c <-chan *asset.Snapshot) (<-chan float64, <-chan float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

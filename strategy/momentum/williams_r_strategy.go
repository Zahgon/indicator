// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/momentum"
	"github.com/cinar/indicator/v2/strategy"
)

const (
	// DefaultWilliamsRStrategyBuyAt defines the default Williams R level at which a Buy action is generated.
	DefaultWilliamsRStrategyBuyAt = -80.0

	// DefaultWilliamsRStrategySellAt defines the default Williams R level at which a Sell action is generated.
	DefaultWilliamsRStrategySellAt = -20.0
)

// WilliamsRStrategy represents the configuration parameters for calculating the Williams R strategy.
type WilliamsRStrategy struct {
	// WilliamsR represents the configuration parameters for calculating the Williams %R.
	WilliamsR *momentum.WilliamsR[float64]

	// BuyAt defines the Williams R level at which a Buy action is generated.
	BuyAt float64

	// SellAt defines the Williams R level at which a Sell action is generated.
	SellAt float64
}

// NewWilliamsRStrategy function initializes a new Williams R strategy instance with the default parameters.
func NewWilliamsRStrategy() *WilliamsRStrategy { _ = "STUB: not implemented"; return nil }

// NewWilliamsRStrategyWith function initializes a new Williams R strategy instance with the given parameters.
func NewWilliamsRStrategyWith(buyAt, sellAt float64) *WilliamsRStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (r *WilliamsRStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (r *WilliamsRStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Williams R starts only after the idle period.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (r *WilliamsRStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> Compute          -> actions -> annotations
	// snapshots[2] -> closings         -> close
	// snapshots[3] -> highs   -|
	// snapshots[4] -> lows    -+-> WilliamsR.Compute -> wr
	// snapshots[5] -> closings-|
	//
	return nil
}

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
	// DefaultRsiStrategyBuyAt defines the default RSI level at which a Buy action is generated.
	DefaultRsiStrategyBuyAt = 30

	// DefaultRsiStrategySellAt defines the default RSI level at which a Sell action is generated.
	DefaultRsiStrategySellAt = 70
)

// RsiStrategy represents the configuration parameters for calculating the RSI strategy.
type RsiStrategy struct {
	// Rsi represents the configuration parameters for calculating the Relative Strength Index (RSI).
	Rsi *momentum.Rsi[float64]

	// BuyAt defines the RSI level at which a Buy action is generated.
	BuyAt float64

	// SellAt defines the RSI level at which a Sell action is generated.
	SellAt float64
}

// NewRsiStrategy function initializes a new RSI strategy instance with the default parameters.
func NewRsiStrategy() *RsiStrategy { _ = "STUB: not implemented"; return nil }

// NewRsiStrategyWith function initializes a new RSI strategy instance with the given parameters.
func NewRsiStrategyWith(buyAt, sellAt float64) *RsiStrategy { _ = "STUB: not implemented"; return nil }

// Name returns the name of the strategy.
func (r *RsiStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (r *RsiStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// RSI starts only after the idle period.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (r *RsiStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> Compute     -> actions -> annotations
	// snapshots[2] -> closings[0] -> close
	//              -> closings[1] -> Rsi.Compute -> rsi
	//
	return nil
}

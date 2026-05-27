// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/volume"
)

const (
	// DefaultMoneyFlowIndexStrategySellAt is the default sell at of 80.
	DefaultMoneyFlowIndexStrategySellAt = 80

	// DefaultMoneyFlowIndexStrategyBuyAt is the default buy at of 20.
	DefaultMoneyFlowIndexStrategyBuyAt = 20
)

// MoneyFlowIndexStrategy represents the configuration parameters for calculating the Money Flow Index strategy.
// Recommends a Sell action when it crosses over 80, and recommends a Buy action when it crosses below 20.
type MoneyFlowIndexStrategy struct {
	// MoneyFlowIndex is the Money Flow Index indicator instance.
	MoneyFlowIndex *volume.Mfi[float64]

	// SellAt is the sell at value.
	SellAt float64

	// BuyAt is the buy at value.
	BuyAt float64
}

// NewMoneyFlowIndexStrategy function initializes a new Money Flow Index strategy instance with the default parameters.
func NewMoneyFlowIndexStrategy() *MoneyFlowIndexStrategy { _ = "STUB: not implemented"; return nil }

// NewMoneyFlowIndexStrategyWith function initializes a new Money Flow Index strategy instance with the
// given parameters.
func NewMoneyFlowIndexStrategyWith(sellAt, buyAt float64) *MoneyFlowIndexStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (m *MoneyFlowIndexStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (m *MoneyFlowIndexStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Money Flow Index starts only after a full period.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (m *MoneyFlowIndexStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> highs       |
	// snapshots[2] -> lows        |
	// snapshots[3] -> closings[0] -> closings
	//                 closings[1] -> money flow index
	// snapshots[4] -> volumes
	// snapshots[5] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

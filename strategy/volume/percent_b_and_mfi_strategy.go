// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/volatility"
	"github.com/cinar/indicator/v2/volume"
)

const (
	// DefaultPercentBandMFIStrategyPercentBBuyAt is the default buy for %B at of 0.8.
	DefaultPercentBandMFIStrategyPercentBBuyAt = 0.8

	// DefaultPercentBandMFIStrategyPercentBSellAt is the default sell for %B at of 0.2.
	DefaultPercentBandMFIStrategyPercentBSellAt = 0.2

	// DefaultPercentBandMFIStrategyMfiBuyAt is the default buy for MFI at of 80.
	DefaultPercentBandMFIStrategyMfiBuyAt = 80

	// DefaultPercentBandMFIStrategyMfiSellAt is the default sell for MFI at of 20.
	DefaultPercentBandMFIStrategyMfiSellAt = 20
)

// PercentBandMFIStrategy represents the configuration parameters for calculating the %B combined with MFI strategy.
// Recommends a Buy action when %B is above 0.8 and MFI is above 80, and recommends a Sell action when %B is below 0.2
// and MFI is below 20.
type PercentBandMFIStrategy struct {
	// MoneyFlowIndex is the Money Flow Index indicator instance.
	MoneyFlowIndex *volume.Mfi[float64]

	// PercentB is the %B indicator instance.
	PercentB *volatility.PercentB[float64]

	// SellPercentBAt is the sell at value of %B.
	SellPercentBAt float64

	// BuyPercentBAt is the buy at value of %B.
	BuyPercentBAt float64

	// SellMfiAt is the sell at value of MFI.
	SellMfiAt float64

	// BuyMfiAt is the buy at value of MFI.
	BuyMfiAt float64
}

// NewPercentBandMFIStrategy function initializes a new PercentBandMFI strategy instance with the default parameters.
func NewPercentBandMFIStrategy() *PercentBandMFIStrategy { _ = "STUB: not implemented"; return nil }

// NewPercentBandMFIStrategyWith function initializes a new PercentBandMFI strategy instance with the
// given parameters.
func NewPercentBandMFIStrategyWith(sellPercentBAt, buyPercentBAt, sellMfiAt, buyMfiAt float64) *PercentBandMFIStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (m *PercentBandMFIStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (m *PercentBandMFIStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// strategy starts only after a full period.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (m *PercentBandMFIStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> highs       |
	// snapshots[2] -> lows        |
	// snapshots[3] -> closings[0] -> closings
	//                 closings[1] -> money flow index
	//                 closings[2] -> percent b
	// snapshots[4] -> volumes
	// snapshots[5] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

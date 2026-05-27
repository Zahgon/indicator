// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/momentum"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultTripleRsiStrategyPeriod defines the default period for the RSI.
	DefaultTripleRsiStrategyPeriod = 5

	// DefaultTripleRsiStrategyMovingAveragePeriod defines the default period for the SMA.
	DefaultTripleRsiStrategyMovingAveragePeriod = 200

	// DefaultTripleRsiStrategyDownDays defines the default number of down days for the RSI.
	DefaultTripleRsiStrategyDownDays = 3

	// DefaultTripleRsiStrategyBuySignalAt defines the default RSI level at which a Buy signal is confirmed.
	DefaultTripleRsiStrategyBuySignalAt = 60

	// DefaultTripleRsiStrategyBuyAt defines the default RSI level at which a Buy action is generated.
	DefaultTripleRsiStrategyBuyAt = 30

	// DefaultTripleRsiStrategySellAt defines the default RSI level at which a Sell action is generated.
	DefaultTripleRsiStrategySellAt = 50
)

// TripleRsiStrategy represents the configuration parameters for calculating the Triple RSI strategy.
// It assumes that the moving average period is longer than the RSI period.
//
// Recommend Buy:
// - The 5-period RSI is below 30.
// - The 5-period RSI reading is down for the 3rd period in a row.
// - The 5-period RSI reading was below 60 three trading periods ago.
// - The close is higher than the 200-period moving average.
//
// Recommend Sell:
// - Sell at the close when the 5-period RSI crosses above 50.
//
// Based on [Triple RSI Trading Strategy: Enhance Your Win Rate to 90% — Advanced Insights](https://tradingstrategy.medium.com/triple-rsi-trading-strategy-enhance-your-win-rate-to-90-advanced-insights-6143059ce41d).
type TripleRsiStrategy struct {
	// Rsi represents the configuration parameters for calculating the Relative Strength Index (RSI).
	Rsi *momentum.Rsi[float64]

	// Sma represents the configuration parameters for calculating the Simple Moving Average (SMA).
	Sma *trend.Sma[float64]

	// DownDays is the number of down days for RSI.
	DownDays int

	// BuySignalAt defines the RSI level at which a Buy signal is confirmed.
	BuySignalAt float64

	// BuyAt defines the RSI level at which a Buy action is generated.
	BuyAt float64

	// SellAt defines the RSI level at which a Sell action is generated.
	SellAt float64
}

// NewTripleRsiStrategy function initializes a new Triple RSI strategy instance with the default parameters.
func NewTripleRsiStrategy() *TripleRsiStrategy { _ = "STUB: not implemented"; return nil }

// NewTripleRsiStrategyWith function initializes a new RSI strategy instance with the given parameters.
func NewTripleRsiStrategyWith(period, smaPeriod, downDays int, buySignalAt, buyAt, sellAt float64) *TripleRsiStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (t *TripleRsiStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// IdlePeriod is the initial period that the Triple RSI strategy won't yield any results.
func (t *TripleRsiStrategy) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (t *TripleRsiStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Skip RSI results until SMA is ready.

// Skip closing values until SMA is ready.

// Recommend Sell:
// - Sell at the close when the 5-period RSI crosses above 50.

// Recommend Buy:
// - The 5-period RSI is below 30.

// - The 5-period RSI reading is down for the 3rd period in a row.

// - The 5-period RSI reading was below 60 three trading periods ago.

// - The close is higher than the 200-period moving average.

// Shift actions until strategy is ready.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (t *TripleRsiStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> Compute     -> actions -> annotations
	// snapshots[2] -> closings[0] -> close
	//              -> closings[1] -> Rsi.Compute -> rsi
	//              -> closings[2] -> Sma.Compute -> sma
	//
	return nil
}

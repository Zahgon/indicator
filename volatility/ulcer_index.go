// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultUlcerIndexPeriod is the default period for the Ulcer Index.
	DefaultUlcerIndexPeriod = 14
)

// UlcerIndex represents the configuration parameters for calculating the Ulcer Index (UI).
// It measures downside risk. The index increases in value as the price moves farther away
// from a recent high and falls as the price rises to new highs.
//
//	High Closings = Max(period, Closings)
//	Percentage Drawdown = 100 * ((Closings - High Closings) / High Closings)
//	Squared Average = Sma(period, Percent Drawdown * Percent Drawdown)
//	Ulcer Index = Sqrt(Squared Average)
//
// Example:
//
//	ui := volatility.NewUlcerIndex[float64]()
//	ui.Compute(closings)
type UlcerIndex[T helper.Number] struct {
	// Time period.
	Period int
}

// NewUlcerIndex function initializes a new Ulcer Index instance with the default parameters.
func NewUlcerIndex[T helper.Number]() *UlcerIndex[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the Ulcer Index over the specified period.
func (u *UlcerIndex[T]) Compute(closings <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

//	High Closings = Max(period, Closings)

//	Percentage Drawdown = 100 * ((Closings - High Closings) / High Closings)

//	Squared Average = Sma(period, Percent Drawdown * Percent Drawdown)

// Ulcer Index = Sqrt(Squared Average)

// IdlePeriod is the initial period that Ulcer Index won't yield any results.
func (u *UlcerIndex[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

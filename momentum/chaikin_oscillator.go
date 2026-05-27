// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
	"github.com/cinar/indicator/v2/volume"
)

const (
	// DefaultChaikinOscillatorShortPeriod is the default short period for the Chaikin Oscillator.
	DefaultChaikinOscillatorShortPeriod = 3

	// DefaultChaikinOscillatorLongPeriod is the default long period for the Chaikin Oscillator.
	DefaultChaikinOscillatorLongPeriod = 10
)

// ChaikinOscillator represents the configuration parameter for calculating the Chaikin Oscillator. It measures
// the momentum of the Accumulation/Distribution (A/D) using the Moving Average Convergence Divergence (MACD)
// formula. It takes the difference between fast and slow periods EMA of the A/D. Cross above the A/D line
// indicates bullish.
//
//	CO = Ema(fastPeriod, AD) - Ema(slowPeriod, AD)
//
// Example:
//
//	co := momentum.ChaikinOscillator[float64]()
//	values := co.Compute(lows, highs)
type ChaikinOscillator[T helper.Number] struct {
	// Ad is the Accumulation/Distribution (A/D) instance.
	Ad *volume.Ad[T]

	// ShortEma is the SMA for the short period.
	ShortEma *trend.Ema[T]

	// LongEma is the SMA for the long period.
	LongEma *trend.Ema[T]
}

// NewChaikinOscillator function initializes a new Chaikin Oscillator instance.
func NewChaikinOscillator[T helper.Number]() *ChaikinOscillator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the Chaikin Oscillator.
func (c *ChaikinOscillator[T]) Compute(highs, lows, closings, volumes <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IdlePeriod is the initial period that Chaikin Oscillator won't yield any results.
func (c *ChaikinOscillator[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

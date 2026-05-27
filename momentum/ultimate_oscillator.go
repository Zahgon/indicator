// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultUltimateOscillatorShortPeriod is the default short period for the Ultimate Oscillator (UO).
	DefaultUltimateOscillatorShortPeriod = 7

	// DefaultUltimateOscillatorMediumPeriod is the default medium period for the Ultimate Oscillator (UO).
	DefaultUltimateOscillatorMediumPeriod = 14

	// DefaultUltimateOscillatorLongPeriod is the default long period for the Ultimate Oscillator (UO).
	DefaultUltimateOscillatorLongPeriod = 28
)

// UltimateOscillator represents the configuration parameter for calculating the Ultimate Oscillator (UO).
// It was developed by Larry Williams in 1976 to measure the price momentum of an asset across multiple
// timeframes. By using the weighted average of three different timeframes the indicator has less
// volatility and fewer trade signals compared to other oscillators that rely on a single timeframe.
//
//	BP = Close - Minimum(Low, Prior Close)
//	TR = Maximum(High, Prior Close) - Minimum(Low, Prior Close)
//	Average7 = Sum(BP for 7 periods) / Sum(TR for 7 periods)
//	Average14 = Sum(BP for 14 periods) / Sum(TR for 14 periods)
//	Average28 = Sum(BP for 28 periods) / Sum(TR for 28 periods)
//	UO = 100 * [(4 * Average7) + (2 * Average14) + Average28] / (4 + 2 + 1)
//
// Example:
//
//	uo := momentum.NewUltimateOscillator[float64]()
//	values := uo.Compute(highs, lows, closings)
type UltimateOscillator[T helper.Number] struct {
	// ShortPeriod is the short period for the UO.
	ShortPeriod int

	// MediumPeriod is the medium period for the UO.
	MediumPeriod int

	// LongPeriod is the long period for the UO.
	LongPeriod int
}

// NewUltimateOscillator function initializes a new Ultimate Oscillator instance.
func NewUltimateOscillator[T helper.Number]() *UltimateOscillator[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewUltimateOscillatorWithPeriods function initializes a new Ultimate Oscillator instance with the given periods.
func NewUltimateOscillatorWithPeriods[T helper.Number](shortPeriod, mediumPeriod, longPeriod int) *UltimateOscillator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the Ultimate Oscillator.
func (u *UltimateOscillator[T]) Compute(highs, lows, closings <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// Align sums to the long period

// UO = 100 * [(4 * Average7) + (2 * Average14) + Average28] / (4 + 2 + 1)
// (4 + 2 + 1) = 7

// IdlePeriod is the initial period that Ultimate Oscillator won't yield any results.
func (u *UltimateOscillator[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

func (u *UltimateOscillator[T]) String() string { _ = "STUB: not implemented"; return "" }

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultMacdPeriod1 is the period for the first EMA.
	DefaultMacdPeriod1 = 12

	// DefaultMacdPeriod2 is the period for the second EMA.
	DefaultMacdPeriod2 = 26

	// DefaultMacdPeriod3 is the period for the third EMA.
	DefaultMacdPeriod3 = 9
)

// Macd represents the configuration parameters for calculating the
// Moving Average Convergence Divergence (MACD).
//
//	MACD = 12-Period EMA - 26-Period EMA.
//	Signal = 9-Period EMA of MACD.
//
// Example:
type Macd[T helper.Number] struct {
	Ema1 *Ema[T]
	Ema2 *Ema[T]
	Ema3 *Ema[T]
}

// NewMacd function initializes a new MACD instance with the default parameters.
func NewMacd[T helper.Number]() *Macd[T] { _ = "STUB: not implemented"; return nil }

// NewMacdWithPeriod function initializes a new MACD instance with the given parameters.
func NewMacdWithPeriod[T helper.Number](period1, period2, period3 int) *Macd[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the MACD
// and the signal line.
func (m *Macd[T]) Compute(c <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IdlePeriod is the initial period that MACD won't yield any results.
func (m *Macd[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

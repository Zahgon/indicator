// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultChopPeriod is the default period for the Choppiness Index (CHOP).
	DefaultChopPeriod = 14
)

// Chop represents the configuration parameters for calculating the Choppiness Index (CHOP).
// It is a technical analysis indicator that measures the market's trendiness or choppiness.
//
//	CHOP = 100 * LOG10( SUM(ATR(1), n) / (MAX(High, n) - MIN(Low, n)) ) / LOG10(n)
type Chop[T helper.Number] struct {
	// Period is the period for the CHOP.
	Period int
}

// NewChop function initializes a new CHOP instance with the default parameters.
func NewChop[T helper.Number]() *Chop[T] { _ = "STUB: not implemented"; return nil }

// NewChopWithPeriod function initializes a new CHOP instance with the given period.
func NewChopWithPeriod[T helper.Number](period int) *Chop[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes channels of highs, lows, and closings, and computes the CHOP over the specified period.
func (c *Chop[T]) Compute(highs, lows, closings <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// TR calculation
// Use previous closing by skipping highs and lows by one.

// MAX(High, n) and MIN(Low, n)
// They should be aligned with the TR bars (starting from bar 1).

// IdlePeriod is the initial period that CHOP won't yield any results.
func (c *Chop[T]) IdlePeriod() int {
	_ = "STUB: not implemented"

	// String function returns a string representation of the CHOP.
	return 0
}

func (c *Chop[T]) String() string { _ = "STUB: not implemented"; return "" }

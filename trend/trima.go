// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultTrimaPeriod is the default period for TRIMA.
	DefaultTrimaPeriod = 15
)

// Trima represents the configuration parameters for calculating the
// Triangular Moving Average (TRIMA).
//
// If period is even:
//
//	TRIMA = SMA(period / 2, SMA((period / 2) + 1, values))
//
// If period is odd:
//
//	TRIMA = SMA((period + 1) / 2, SMA((period + 1) / 2, values))
type Trima[T helper.Number] struct {
	// Time period.
	Period int
}

// NewTrima function initializes a new TRIMA instance
// with the default parameters.
func NewTrima[T helper.Number]() *Trima[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the TRIMA
// and the signal line.
func (t *Trima[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that TRIMA won't yield any results.
func (t *Trima[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// calculatePeriods calculates the individual periods to use based on the
// TRIMA period.
func (t *Trima[T]) calculatePeriods() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

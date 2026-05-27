// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultTrixPeriod is the default time period for TRIX.
	DefaultTrixPeriod = 15
)

// Trix represents the configuration parameters for calculating the Triple Exponential Average (TRIX).
// TRIX indicator is an oscillator used to identify oversold and overbought markets, and it can also
// be used as a momentum indicator. Like many oscillators, TRIX oscillates around a zero line.
//
//	EMA1 = EMA(period, values)
//	EMA2 = EMA(period, EMA1)
//	EMA3 = EMA(period, EMA2)
//	TRIX = (EMA3 - Previous EMA3) / Previous EMA3
//
// Example:
//
//	trix := trend.NewTrix[float64]()
//	result := trix.Compute(values)
type Trix[T helper.Number] struct {
	// Time period.
	Period int
}

// NewTrix function initializes a new TRIX instance with the default parameters.
func NewTrix[T helper.Number]() *Trix[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the TRIX and the signal line.
func (t *Trix[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that TRIX won't yield any results.
func (t *Trix[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

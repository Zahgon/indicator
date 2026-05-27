// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultEmaPeriod is the default EMA period of 20.
	DefaultEmaPeriod = 20

	// DefaultEmaSmoothing is the default EMA smooting of 2.
	DefaultEmaSmoothing = 2
)

// Ema represents the parameters for calculating the Exponential Moving Average.
//
// Example:
//
//	ema := trend.NewEma[float64]()
//	ema.Period = 10
//
//	result := ema.Compute(c)
type Ema[T helper.Number] struct {
	// Time period.
	Period int

	// Smoothing constant.
	Smoothing T
}

// NewEma function initializes a new EMA instance with the default parameters.
func NewEma[T helper.Number]() *Ema[T] { _ = "STUB: not implemented"; return nil }

// NewEmaWithPeriod function initializes a new EMA instance with the given period.
func NewEmaWithPeriod[T helper.Number](period int) *Ema[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the EMA over the specified period.
func (e *Ema[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// Initial EMA value is the SMA.

// IdlePeriod is the initial period that EMA yield any results.
func (e *Ema[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the EMA.
func (e *Ema[T]) String() string { _ = "STUB: not implemented"; return "" }

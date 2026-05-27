// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

// Wma represents the configuration parameters for calculating the Weighted Moving Average (WMA).
// It calculates a moving average by putting more weight on recent data and less on past data.
//
//	WMA = ((Value1 * 1/N) + (Value2 * 2/N) + ...) / 2
type Wma[T helper.Number] struct {
	// Time period.
	Period int
}

// NewWmaWith function initializes a new WMA instance with the given parameters.
func NewWmaWith[T helper.Number](period int) *Wma[T] { _ = "STUB: not implemented"; return nil }

// Compute computes the WMA over the input stream.
func (w *Wma[T]) Compute(values <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that WMA won't yield any results.
func (w *Wma[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the WMA.
func (w *Wma[T]) String() string { _ = "STUB: not implemented"; return "" }

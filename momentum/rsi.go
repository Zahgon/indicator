// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultRsiPeriod is the default period for the Relative Strength Index (RSI).
	DefaultRsiPeriod = 14
)

// Rsi represents the configuration parameter for calculating the Relative Strength Index (RSI).  It is a momentum
// indicator that measures the magnitude of recent price changes to evaluate overbought and oversold conditions.
//
//	RS = Average Gain / Average Loss
//	RSI = 100 - (100 / (1 + RS))
//
// Example:
//
//	rsi := momentum.NewRsi[float64]()
//	result := rsi.Compute(closings)
type Rsi[T helper.Number] struct {
	// Rma is the RMA instance.
	Rma *trend.Rma[T]
}

// NewRsi function initializes a new Relative Strength Index instance with the default parameters.
func NewRsi[T helper.Number]() *Rsi[T] { _ = "STUB: not implemented"; return nil }

// NewRsiWithPeriod function initializes a new Relative Strength Index instance with the given period.
func NewRsiWithPeriod[T helper.Number](period int) *Rsi[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of closings numbers and computes the Relative Strength Index.
func (r *Rsi[T]) Compute(closings <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// RSI = 100 - (100 / (1 + RS))

// - (100 / (1 + RS))

// 100 / (1 + RS)

// 1 / (1 + RS)

// 1 + RS

// IdlePeriod is the initial period that Relative Strength Index won't yield any results.
func (r *Rsi[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

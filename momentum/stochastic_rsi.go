// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultStochasticRsiPeriod is the default period for the Stochastic Relative Strength Index (RSI).
	DefaultStochasticRsiPeriod = 14
)

// StochasticRsi represents the configuration parameter for calculating the Stochastic Relative Strength Index (RSI).
// It is a momentum indicator that focuses on the historical performance to evaluate overbought and
// oversold conditions.
//
//	                    RSI - Min(RSI)
//	Stochastic RSI = -------------------------
//	                   Max(RSI) - Min(RSI)
//
// Example:
//
//	stochasticRsi := momentum.NewStochasticRsi[float64]()
//	result := stochasticRsi.Compute(closings)
type StochasticRsi[T helper.Number] struct {
	// Rsi is that RSI instance.
	Rsi *Rsi[T]

	// Min is the Moving Min instance.
	Min *trend.MovingMin[T]

	// Max is the Moving Max instance.
	Max *trend.MovingMax[T]
}

// NewStochasticRsi function initializes a new Storchastic RSI instance with the default parameters.
func NewStochasticRsi[T helper.Number]() *StochasticRsi[T] { _ = "STUB: not implemented"; return nil }

// NewStochasticRsiWithPeriod function initializes a new Stochastic RSI instance with the given period.
func NewStochasticRsiWithPeriod[T helper.Number](period int) *StochasticRsi[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of closings numbers and computes the Stochastic RSI.
func (s *StochasticRsi[T]) Compute(closings <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// IdlePeriod is the initial period that Stochasic RSI won't yield any results.
func (s *StochasticRsi[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

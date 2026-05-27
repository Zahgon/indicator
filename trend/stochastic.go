// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import "github.com/cinar/indicator/v2/helper"

const (
	// DefaultStochasticPeriod is the default period for the Stochastic indicator.
	DefaultStochasticPeriod = 10

	// DefaultStochasticSmaPeriod is the default period for the SMA of %K.
	DefaultStochasticSmaPeriod = 3
)

// Stochastic represents the configuration parameters for calculating
// the Stochastic indicator on a single input series. This is different
// from the Stochastic Oscillator which operates on high, low, and close.
// This generic version is useful for applying stochastic calculation to
// any series, such as MACD values in the Schaff Trend Cycle (STC).
//
//	K = (Value - Min(Value, period)) / (Max(Value, period) - Min(Value, period)) * 100
//	D = SMA(K, dPeriod)
//
// Example:
//
//	s := trend.NewStochastic[float64]()
//	k, d := s.Compute(values)
type Stochastic[T helper.Number] struct {
	// Period is the period for the min/max calculation.
	Period int

	// Sma is the SMA instance for %D calculation.
	Sma *Sma[T]
}

// NewStochastic function initializes a new Stochastic instance with the default parameters.
func NewStochastic[T helper.Number]() *Stochastic[T] { _ = "STUB: not implemented"; return nil }

// NewStochasticWithPeriod function initializes a new Stochastic instance with the given period.
func NewStochasticWithPeriod[T helper.Number](period int) *Stochastic[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the Stochastic indicator.
// Returns %K and %D.
func (s *Stochastic[T]) Compute(values <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IdlePeriod is the initial period that Stochastic won't yield any results.
func (s *Stochastic[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

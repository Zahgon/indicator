// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import "github.com/cinar/indicator/v2/helper"

const (
	// DefaultSlowStochasticPeriod is the default period for the Slow Stochastic indicator.
	DefaultSlowStochasticPeriod = 10

	// DefaultSlowStochasticKPeriod is the default period for the Fast %K SMA smoothing.
	DefaultSlowStochasticKPeriod = 3

	// DefaultSlowStochasticDPeriod is the default period for the Slow %D SMA.
	DefaultSlowStochasticDPeriod = 3
)

// SlowStochastic represents the configuration parameters for calculating
// the Slow Stochastic indicator. This applies additional smoothing to the
// Fast Stochastic values.
//
//	Fast %K = Stochastic(values, period)
//	Slow %K = SMA(Fast %K, kPeriod)
//	Slow %D = SMA(Slow %K, dPeriod)
//
// Example:
//
//	s := trend.NewSlowStochastic[float64]()
//	k, d := s.Compute(values)
type SlowStochastic[T helper.Number] struct {
	// Period is the period for the min/max calculation.
	Period int

	// KPeriod is the period for smoothing Fast %K to get Slow %K.
	KPeriod int

	// DPeriod is the period for smoothing Slow %K to get Slow %D.
	DPeriod int
}

// NewSlowStochastic function initializes a new SlowStochastic instance with the default parameters.
func NewSlowStochastic[T helper.Number]() *SlowStochastic[T] { _ = "STUB: not implemented"; return nil }

// NewSlowStochasticWithPeriod function initializes a new SlowStochastic instance with the given periods.
func NewSlowStochasticWithPeriod[T helper.Number](period, kPeriod, dPeriod int) *SlowStochastic[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the Slow Stochastic indicator.
// Returns Slow %K and Slow %D.
func (s *SlowStochastic[T]) Compute(values <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IdlePeriod is the initial period that Slow Stochastic won't yield any results.
func (s *SlowStochastic[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

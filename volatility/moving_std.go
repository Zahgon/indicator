// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultMovingStdPeriod is the default time period for Moving Standard Deviation.
	DefaultMovingStdPeriod = 1
)

// MovingStd represents the configuration parameters for calculating the Moving Standard Deviation
// over the specified period.
//
//	Std = Sqrt(1/Period * Sum(Pow(value - sma), 2))
type MovingStd[T helper.Number] struct {
	// Time period.
	Period int
}

// NewMovingStd function initializes a new Moving Standard Deviation instance with the default parameters.
func NewMovingStd[T helper.Number]() *MovingStd[T] { _ = "STUB: not implemented"; return nil }

// NewMovingStdWithPeriod function initializes a new Moving Standard Deviation instance with the given period.
func NewMovingStdWithPeriod[T helper.Number](period int) *MovingStd[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the Moving Standard Deviation over the specified period.
func (m *MovingStd[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

//	Std = Sqrt(1/Period * Sum(Pow(value - sma), 2))

// IdlePeriod is the initial period that Moving Standard Deviation won't yield any results.
func (m *MovingStd[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

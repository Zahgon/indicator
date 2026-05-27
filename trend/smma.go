// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultSmmaPeriod is the default SMMA period of 7.
	DefaultSmmaPeriod = 7
)

// Smma represents the parameters for calculating the Smoothed Moving Average (SMMA).
//
//	SMMA[0] = SMA(N)
//	SMMA[i] = ((SMMA[i-1] * (N - 1)) + Close[i]) / N
//
// Example:
//
//	smma := trend.NewSmma[float64]()
//	smma.Period = 10
//
//	result := smma.Compute(c)
type Smma[T helper.Number] struct {
	// Time period.
	Period int
}

// NewSmma function initializes a new SMMA instance with the default parameters.
func NewSmma[T helper.Number]() *Smma[T] { _ = "STUB: not implemented"; return nil }

// NewSmmaWithPeriod function initializes a new SMMA instance with the given period.
func NewSmmaWithPeriod[T helper.Number](period int) *Smma[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the SMMA over the specified period.
func (s *Smma[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// Initial SMMA value is the SMA.

// IdlePeriod is the initial period that SMMA yield any results.
func (s *Smma[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the SMMA.
func (s *Smma[T]) String() string { _ = "STUB: not implemented"; return "" }

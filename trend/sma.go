// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultSmaPeriod is the default SMA period.
	DefaultSmaPeriod = 50
)

// Sma represents the parameters for calculating the Simple Moving Average.
//
// Example:
//
//	sma := trend.NewSma[float64]()
//	sma.Period = 10
//
//	result := sma.Compute(c)
type Sma[T helper.Number] struct {
	// Period is the time period for the SMA.
	Period int
}

// NewSma function initializes a new SMA instance with the default parameters.
func NewSma[T helper.Number]() *Sma[T] { _ = "STUB: not implemented"; return nil }

// NewSmaWithPeriod function initializes a new SMA instance with the default parameters.
func NewSmaWithPeriod[T helper.Number](period int) *Sma[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the SMA over the specified period.
func (s *Sma[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that SMA won't yield any results.
func (s *Sma[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the SMA.
func (s *Sma[T]) String() string { _ = "STUB: not implemented"; return "" }

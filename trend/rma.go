// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import "github.com/cinar/indicator/v2/helper"

const (
	// DefaultRmaPeriod is the default RMA period.
	DefaultRmaPeriod = 20
)

// Rma represents the parameters for calculating Rolling Moving Average (RMA).
//
//	R[0] to R[p-1] is SMA(values)
//	R[p] and after is R[i] = ((R[i-1]*(p-1)) + v[i]) / p
//
// Example:
//
//	rma := trend.NewRma[float64]()
//	rma.Period = 10
//
//	result := rma.Compute(c)
type Rma[T helper.Number] struct {
	// Time period.
	Period int
}

// NewRma function initializes a new RMA instance with the default parameters.
func NewRma[T helper.Number]() *Rma[T] { _ = "STUB: not implemented"; return nil }

// NewRmaWithPeriod function initializes a new RMA instance with the given period.
func NewRmaWithPeriod[T helper.Number](period int) *Rma[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the RMA over the specified period.
func (r *Rma[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// Initial RMA value is the SMA.

// IdlePeriod is the initial period that RMA won't yield any results.
func (r *Rma[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

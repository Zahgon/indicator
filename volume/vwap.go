// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultVwapPeriod is the default period for the VWAP.
	DefaultVwapPeriod = 14
)

// Vwap holds configuration parameters for calculating the Volume Weighted Average Price (VWAP). It provides the
// average price the asset has traded.
//
//	VWAP = Sum(Closing * Volume) / Sum(Volume)
//
// Example:
//
//	vwap := volume.NewVwap[float64]()
//	result := vwap.Compute(closings, volumes)
type Vwap[T helper.Number] struct {
	// Sum is the Moving Sum instance.
	Sum *trend.MovingSum[T]
}

// NewVwap function initializes a new VWAP instance with the default parameters.
func NewVwap[T helper.Number]() *Vwap[T] { _ = "STUB: not implemented"; return nil }

// NewVwapWithPeriod function initializes a new VWAP instance with the given period.
func NewVwapWithPeriod[T helper.Number](period int) *Vwap[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the VWAP.
func (v *Vwap[T]) Compute(closings, volumes <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// IdlePeriod is the initial period that VWAP won't yield any results.
func (v *Vwap[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

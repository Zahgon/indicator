// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultFiPeriod is the default period for the FI.
	DefaultFiPeriod = 13
)

// Fi holds configuration parameters for calculating the Force Index (FI). It uses the closing price and the volume to
// assess the power behind a move and identify turning points.
//
//	FI = EMA(period, (Current - Previous) * Volume)
//
// Example:
//
//	fi := volume.NewFi[float64]()
//	result := fi.Compute(closings, volumes)
type Fi[T helper.Number] struct {
	// Ema is the EMA instance.
	Ema *trend.Ema[T]
}

// NewFi function initializes a new FI instance with the default parameters.
func NewFi[T helper.Number]() *Fi[T] { _ = "STUB: not implemented"; return nil }

// NewFiWithPeriod function initializes a new FI instance with the given period.
func NewFiWithPeriod[T helper.Number](period int) *Fi[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the FI.
func (f *Fi[T]) Compute(closings, volumes <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that FI won't yield any results.
func (f *Fi[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

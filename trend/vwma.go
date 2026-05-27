// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultVwmaPeriod is the default period for the VWMA.
	DefaultVwmaPeriod = 20
)

// Vwma represents the configuration parameters for calculating the Volume Weighted Moving Average (VWMA)
// It averages the price data with an emphasis on volume, meaning areas with higher volume will have a
// greater weight.
//
//	VWMA = Sum(Price * Volume) / Sum(Volume)
type Vwma[T helper.Number] struct {
	// Time period.
	Period int
}

// NewVwma function initializes a new VWMA instance with the default parameters.
func NewVwma[T helper.Number]() *Vwma[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the VWMA and the signal line.
func (v *Vwma[T]) Compute(closing, volume <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that VWMA won't yield any results.
func (v *Vwma[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/helper"
)

// Vpt holds configuration parameters for calculating the Volume Price Trend (VPT). It provides a correlation
// between the volume and the price.
//
//	VPT = Previous VPT + (Volume * (Current Closing - Previous Closing) / Previous Closing)
//
// Example:
//
//	vpt := volume.NewVpt[float64]()
//	result := vpt.Compute(closings, volumes)
type Vpt[T helper.Number] struct{}

// NewVpt function initializes a new VPT instance with the default parameters.
func NewVpt[T helper.Number]() *Vpt[T] {
	_ = "STUB: not implemented"

	// Compute function takes a channel of numbers and computes the VPT.
	return nil
}

func (*Vpt[T]) Compute(closings, volumes <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that VPT won't yield any results.
func (*Vpt[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

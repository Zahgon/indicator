// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

// DefaultDpoPeriod is the default period for DPO calculation.
const DefaultDpoPeriod = 20

// Dpo computes the Detrended Price Oscillator.
// Formula (common approximation):
// Let k = floor(period/2) + 1.
// For time index t >= period-1+k:
//
//	DPO[t] = Price[t] - SMA[t - k]
//
// Example:
//
//	dpo := trend.NewDpoWithPeriod[float64](20)
//	out := dpo.Compute(c)
type Dpo[T helper.Float] struct {
	// Period is the SMA window length. Typical default is 20.
	// Note: values <= 1 are considered invalid and will be replaced with DefaultDpoPeriod by constructors.
	period int
}

// NewDpo creates a new DPO instance with default parameters.
func NewDpo[T helper.Float]() *Dpo[T] { _ = "STUB: not implemented"; return nil }

// NewDpoWithPeriod initializes a new DPO instance with the given period.
// Periods <= 1 are clamped to DefaultDpoPeriod.
func NewDpoWithPeriod[T helper.Float](period int) *Dpo[T] { _ = "STUB: not implemented"; return nil }

// Compute calculates the DPO indicator over the input price channel.
func (d *Dpo[T]) Compute(closing <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// compute SMA on the first duplicated stream

// align the original price stream and the SMA stream according to DPO formula

// DPO = Price - shifted SMA

// IdlePeriod returns the number of leading samples to discard before the first DPO value is available.
func (d *Dpo[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the DPO.
func (d *Dpo[T]) String() string { _ = "STUB: not implemented"; return "" }

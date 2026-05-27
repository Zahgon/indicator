// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultPoPeriod is the default period for the Projection Oscillator (PO).
	DefaultPoPeriod = 14
)

// Po represents the configuration parameters for calculating the Projection Oscillator (PO). It uses the linear
// regression slope, along with highs and lows. Period defines the moving window to calculates the PO.
//
//	PL = Min(period, (high + MLS(period, x, high)))
//	PH = Max(period, (low + MLS(period, x, low)))
//	PO = 100 * (Closing - PL) / (PH - PL)
//
// Example:
//
//	po := volatility.NewPo()
//	ps := po.Compute(highs, lows, closings)
type Po[T helper.Number] struct {
	// Mls is the Moving Least Square instance.
	mls *trend.Mls[T]

	// Min is the Moving Min instance.
	min *trend.MovingMin[T]

	// Max is the Moving Max instance.
	max *trend.MovingMax[T]
}

// NewPo function initializes a new PO instance with the default parameters.
func NewPo[T helper.Number]() *Po[T] { _ = "STUB: not implemented"; return nil }

// NewPoWithPeriod function initializes a new PO instance with the given period.
func NewPoWithPeriod[T helper.Number](period int) *Po[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the PO over the specified period.
func (p *Po[T]) Compute(highs, lows, closings <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// PL = Min(period, (high + MLS(period, x, high)))

// PH = Max(period, (low + MLS(period, x, low)))

// PO = 100 * (Closing - PL) / (PH - PL)

// IdlePeriod is the initial period that PO won't yield any results.
func (p *Po[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

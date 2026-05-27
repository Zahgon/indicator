// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
)

// PercentB represents the parameters for calculating the %B indicator.
//
//	%B = (Close - Lower Band) / (Upper Band - Lower Band)
type PercentB[T helper.Number] struct {
	// BollingerBands is the underlying Bollinger Bands indicator used for calculations.
	BollingerBands *BollingerBands[T]
}

// NewPercentB function initializes a new %B instance with the default parameters.
func NewPercentB[T helper.Number]() *PercentB[T] { _ = "STUB: not implemented"; return nil }

// NewPercentBWithPeriod function initializes a new %B instance with the given period.
func NewPercentBWithPeriod[T helper.Number](period int) *PercentB[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the %B over the specified period.
func (p *PercentB[T]) Compute(closings <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// Compute the Bollinger Bands

// Skip closings until the Bollinger Bands are available

// Drain the middle bands

// %B = (Close - Lower Band) / (Upper Band - Lower Band)

// IdlePeriod is the initial period that %B yield any results.
func (p *PercentB[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the %B.
func (p *PercentB[T]) String() string { _ = "STUB: not implemented"; return "" }

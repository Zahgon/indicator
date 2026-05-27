// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
)

// TrueRange represents the parameters for calculating the True Range (TR).
// It evaluates the greatest distance covered by price in a single period,
// accounting for gaps.
//
//	TR = Max((High - Low), (High - Previous Closing), (Previous Closing - Low))
//
// Example:
//
//	tr := volatility.NewTrueRange[float64]()
//	result := tr.Compute(highs, lows, closings)
type TrueRange[T helper.Number] struct{}

// NewTrueRange function initializes a new TrueRange instance.
func NewTrueRange[T helper.Number]() *TrueRange[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes channels of highs, lows, and closings and computes the True Range.
func (tr *TrueRange[T]) Compute(highs, lows, closings <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// IdlePeriod is the initial period that TrueRange won't yield any results.
func (tr *TrueRange[T]) IdlePeriod() int {
	_ = "STUB: not implemented"

	// String is the string representation of the TrueRange.
	return 0
}

func (tr *TrueRange[T]) String() string { _ = "STUB: not implemented"; return "" }

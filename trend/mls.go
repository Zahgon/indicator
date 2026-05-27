// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

// Mls represents the configuration parameters for calculating the Moving Least Square (MLS). It is a regression
// analysis to determine the line of best fit for the given set of data.
//
//	y = mx + b
//	b = y-intercept
//	y = slope
//
//	m = (period * sumXY - sumX * sumY) / (period * sumX2 - sumX * sumX)
//	b = (sumY - m * sumX) / period
//
// Example:
//
//	mls := trend.NewMlsWithPeriod[float64](14)
//	ms, bs := mls.Compute(x , y)
type Mls[T helper.Number] struct {
	// Sum is the moving sum instance.
	Sum *MovingSum[T]
}

// NewMlsWithPeriod function initializes a new MLS instance with the given period.
func NewMlsWithPeriod[T helper.Number](period int) *Mls[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the MLS m and b.
func (m *Mls[T]) Compute(x, y <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// m = (period * sumXY - sumX * sumY) / (period * sumX2 - sumX * sumX)

// b = (sumY - m * sumX) / period

// IdlePeriod is the initial period that MLS won't yield any results.
func (m *Mls[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

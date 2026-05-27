// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import "github.com/cinar/indicator/v2/helper"

// MovingMin represents the configuration parameters for calculating the
// Moving Min over the specified period.
//
// Example:
type MovingMin[T helper.Number] struct {
	// Time period.
	Period int
}

// NewMovingMin function initializes a new Moving Min instance with the default parameters.
func NewMovingMin[T helper.Number]() *MovingMin[T] { _ = "STUB: not implemented"; return nil }

// NewMovingMinWithPeriod function initializes a new Moving Min instance with the given period.
func NewMovingMinWithPeriod[T helper.Number](period int) *MovingMin[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the
// Moving Min over the specified period.
func (m *MovingMin[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that Mocing Min won't yield any results.
func (m *MovingMin[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

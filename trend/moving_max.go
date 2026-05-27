// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import "github.com/cinar/indicator/v2/helper"

// MovingMax represents the configuration parameters for calculating the
// Moving Max over the specified period.
//
// Example:
type MovingMax[T helper.Number] struct {
	// Time period.
	Period int
}

// NewMovingMax function initializes a new Moving Max instance with the default parameters.
func NewMovingMax[T helper.Number]() *MovingMax[T] { _ = "STUB: not implemented"; return nil }

// NewMovingMaxWithPeriod function initializes a new Moving Max instance with the given period.
func NewMovingMaxWithPeriod[T helper.Number](period int) *MovingMax[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of numbers and computes the
// Moving Max over the specified period.
func (m *MovingMax[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that Mocing Max won't yield any results.
func (m *MovingMax[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

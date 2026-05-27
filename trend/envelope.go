// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultEnvelopePercentage is the default envelope percentage of 20%.
	DefaultEnvelopePercentage = 20

	// DefaultEnvelopePeriod is the default envelope period of 20.
	DefaultEnvelopePeriod = 20
)

// Envelope represents the parameters neededd to calcualte the Envelope.
type Envelope[T helper.Number] struct {
	// Ma is the moving average used.
	Ma Ma[T]

	// Percentage is the envelope percentage.
	Percentage T
}

// NewEnvelope function initializes a new Envelope instance with the default parameters.
func NewEnvelope[T helper.Number](ma Ma[T], percentage T) *Envelope[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewEnvelopeWithSma function initalizes a new Envelope instance using SMA.
func NewEnvelopeWithSma[T helper.Number]() *Envelope[T] { _ = "STUB: not implemented"; return nil }

// NewEnvelopeWithEma function initializes a new Envelope instance using EMA.
func NewEnvelopeWithEma[T helper.Number]() *Envelope[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the Envelope over the specified period.
func (e *Envelope[T]) Compute(closings <-chan T) (<-chan T, <-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// IdlePeriod is the initial period that Envelope yield any results.
func (e *Envelope[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the Envelope.
func (e *Envelope[T]) String() string { _ = "STUB: not implemented"; return "" }

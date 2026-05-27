// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultRocPeriod is the default ROC period.
	DefaultRocPeriod = 9
)

// Roc represents the configuration parameters for calculating the Rate Of Change (ROC) indicator.
//
//	ROC = (Current Price - Price n periods ago) / Price n periods ago
type Roc[T helper.Float] struct {
	// Time period.
	Period int
}

// NewRoc function initializes a new Roc instance with the default parameters.
func NewRoc[T helper.Float]() *Roc[T] { _ = "STUB: not implemented"; return nil }

// NewRocWithPeriod function initializes a new Roc instance with the given parameters.
func NewRocWithPeriod[T helper.Float](period int) *Roc[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the ROC and the signal line.
func (r *Roc[T]) Compute(values <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that ROC won't yield any results.
func (r *Roc[T]) IdlePeriod() int {
	_ = "STUB: not implemented"

	// String is the string representation of the ROC.
	return 0
}

func (r *Roc[T]) String() string { _ = "STUB: not implemented"; return "" }

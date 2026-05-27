// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultEmvPeriod is the default period for the EMV.
	DefaultEmvPeriod = 14
)

// Emv holds configuration parameters for calculating the Ease of Movement (EMV). It is a volume based oscillator
// measuring the ease of price movement.
//
//	Distance Moved = ((High + Low) / 2) - ((Priod High + Prior Low) /2)
//	Box Ratio = ((Volume / 100000000) / (High - Low))
//	EMV(1) = Distance Moved / Box Ratio
//	EMV(14) = SMA(14, EMV(1))
//
// Example:
//
//	emv := volume.NewEmv[float64]()
//	result := emv.Compute(highs, lows, volumes)
type Emv[T helper.Number] struct {
	// Sma is the SMA instance.
	Sma *trend.Sma[T]
}

// NewEmv function initializes a new EMV instance with the default parameters.
func NewEmv[T helper.Number]() *Emv[T] { _ = "STUB: not implemented"; return nil }

// NewEmvWithPeriod function initializes a new EMV instance with the given period.
func NewEmvWithPeriod[T helper.Number](period int) *Emv[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the EMV.
func (e *Emv[T]) Compute(highs, lows, volumes <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

//	Distance Moved = ((High + Low) / 2) - ((Priod High + Prior Low) /2)

// Box Ratio = ((Volume / 100000000) / (High - Low))

// EMV(1) = Distance Moved / Box Ratio
// EMV(14) = SMA(14, EMV(1))

// IdlePeriod is the initial period that EMV won't yield any results.
func (e *Emv[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

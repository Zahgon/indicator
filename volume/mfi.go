// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultMfiPeriod is the default period of the MFI.
	DefaultMfiPeriod = 14
)

// Mfi holds configuration parameters for calculating the Money Flow Index (MFI). It analyzes both the closing price
// and the volume to measure to identify overbought and oversold states. It is similar to the Relative Strength
// Index (RSI), but it also uses the volume.
//
//	Raw Money Flow = Typical Price * Volume
//	Money Ratio = Positive Money Flow / Negative Money Flow
//	Money Flow Index = 100 - (100 / (1 + Money Ratio))
//
// Example:
//
//	mfi := volume.NewMfi[float64]()
//	result := mfi.Compute(highs, lows, closings, volumes)
type Mfi[T helper.Number] struct {
	// TypicalPrice is the Typical Price instance.
	TypicalPrice *trend.TypicalPrice[T]

	// Sum is the Moving Sum instance.
	Sum *trend.MovingSum[T]
}

// NewMfi function initializes a new MFI instance with the default parameters.
func NewMfi[T helper.Number]() *Mfi[T] { _ = "STUB: not implemented"; return nil }

// NewMfiWithPeriod function initializes a new MFI instance with the given period.
func NewMfiWithPeriod[T helper.Number](period int) *Mfi[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the MFI.
func (m *Mfi[T]) Compute(highs, lows, closings, volumes <-chan T) <-chan T {
	_ = "STUB: not implemented"
	//	Raw Money Flow = Typical Price * Volume
	return nil
}

// Money Ratio = Positive Money Flow / Negative Money Flow

// Money Flow Index = 100 - (100 / (1 + Money Ratio))

// IdlePeriod is the initial period that MFI won't yield any results.
func (m *Mfi[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

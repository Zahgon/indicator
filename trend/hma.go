// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

// Hma represents the configuration parameters for calculating the Hull Moving Average (HMA). Developed by
// Alan Hull in 2005, HMA attempts to minimize the lag of a traditional moving average.
//
//	WMA1 = WMA(period/2 , values)
//	WMA2 = WMA(period, values)
//	WMA3 = WMA(sqrt(period), (2 * WMA1) - WMA2)
//	HMA = WMA3
type Hma[T helper.Number] struct {
	// First WMA.
	wma1 *Wma[T]

	// Second WMA.
	wma2 *Wma[T]

	// Third WMA.
	wma3 *Wma[T]
}

// NewHmaWithPeriod function initializes a new HMA instance with the given parameters.
func NewHmaWithPeriod[T helper.Number](period int) *Hma[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the HMA and the signal line.
func (h *Hma[T]) Compute(values <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

//	WMA1 = WMA(period/2 , values)

//	WMA2 = WMA(period, values)

// WMA3 = WMA(sqrt(period), (2 * WMA1) - WMA2)

// HMA = WMA3

// IdlePeriod is the initial period that HMA won't yield any results.
func (h *Hma[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// String is the string representation of the HMA.
func (h *Hma[T]) String() string { _ = "STUB: not implemented"; return "" }

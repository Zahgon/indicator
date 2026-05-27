// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultChandelierExitPeriod is the default period for the Chandelier Exit.
	DefaultChandelierExitPeriod = 22

	// DefaultChandelierExitMultiplier is the default multiplier for the Chandelier Exit.
	DefaultChandelierExitMultiplier = 3
)

// ChandelierExit represents the configuration parameters for calculating the Chandelier Exit.
// It sets a trailing stop-loss based on the Average True Value (ATR).
//
//	Chandelier Exit Long = 22-Period SMA High - ATR(22) * 3
//	Chandelier Exit Short = 22-Period SMA Low + ATR(22) * 3
//
// Example:
//
//	ce := volatility.NewChandelierExit[float64]()
//	ceLong, ceShort := ce.Compute(highs, lows, closings)
type ChandelierExit[T helper.Number] struct {
	// Period is time period.
	Period int

	// Multiplier is for sensitivity.
	Multiplier T
}

// NewChandelierExit function initializes a new Chandelier Exit instance with the default parameters.
func NewChandelierExit[T helper.Number]() *ChandelierExit[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the Chandelier Exit over the specified period.
func (c *ChandelierExit[T]) Compute(highs, lows, closings <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IdlePeriod is the initial period that Chandelier Exit won't yield any results.
func (c *ChandelierExit[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

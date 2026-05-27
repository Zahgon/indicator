// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultPpoShortPeriod is the default short period for the Percentage Price Oscillator.
	DefaultPpoShortPeriod = 12

	// DefaultPpoLongPeriod is the default long period for the Percentage Price Oscillator.
	DefaultPpoLongPeriod = 26

	// DefaultPpoSignalPeriod is the default signal period for the Percentage Price Oscillator.
	DefaultPpoSignalPeriod = 9
)

// Ppo represents the configuration parameter for calculating the Percentage Price Oscillator (PPO). It is a momentum
// oscillator for the price. It is used to indicate the ups and downs based on the price. A breakout is confirmed
// when PPO is positive.
//
//	PPO = ((EMA(shortPeriod, prices) - EMA(longPeriod, prices)) / EMA(longPeriod, prices)) * 100
//	Signal = EMA(9, PPO)
//	Histogram = PPO - Signal
//
// Example:
//
//	ppo := momentum.Ppo[float64]()
//	p, s, h := ppo.Compute(closings)
type Ppo[T helper.Number] struct {
	// ShortEma is the short EMA instance.
	ShortEma *trend.Ema[T]

	// LongEma is the long EMA instance.
	LongEma *trend.Ema[T]

	// SignalEma is the signal EMA instance.
	SignalEma *trend.Ema[T]
}

// NewPpo function initializes a new Percentage Price Oscillator instance.
func NewPpo[T helper.Number]() *Ppo[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the Percentage Price Oscillator.
// Returns ppo, signal, histogram.
func (p *Ppo[T]) Compute(closings <-chan T) (<-chan T, <-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//	PPO = ((EMA(shortPeriod, prices) - EMA(longPeriod, prices)) / EMA(longPeriod, prices)) * 100

//	Signal = EMA(9, PPO)

//	Histogram = PPO - Signal

// IdlePeriod is the initial period that Percentage Price Oscillator won't yield any results.
func (p *Ppo[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

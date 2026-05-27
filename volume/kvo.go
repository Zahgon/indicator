// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultKvoShortPeriod is the default short period for the Klinger Volume Oscillator.
	DefaultKvoShortPeriod = 34

	// DefaultKvoLongPeriod is the default long period for the Klinger Volume Oscillator.
	DefaultKvoLongPeriod = 55

	// DefaultKvoSignalPeriod is the default signal period for the Klinger Volume Oscillator.
	DefaultKvoSignalPeriod = 13
)

// Kvo represents the configuration parameters for calculating the Klinger Volume Oscillator (KVO).
// It is a volume-based oscillator that identifies long-term money flow trends using EMA differences.
//
//	Trend = +1 if High > High[1] and Low >= Low[1]
//	Trend = -1 if High <= High[1] and Low < Low[1]
//	Trend = 0 otherwise
//	VF = Volume * Trend
//	KVO = EMA(VF, shortPeriod) - EMA(VF, longPeriod)
//	Signal = EMA(KVO, signalPeriod)
//
// Example:
//
//	kvo := volume.NewKvo[float64]()
//	kvoResult, signalResult := kvo.Compute(highs, lows, volumes)
type Kvo[T helper.Number] struct {
	// ShortEma is the short EMA instance.
	ShortEma *trend.Ema[T]

	// LongEma is the long EMA instance.
	LongEma *trend.Ema[T]

	// SignalEma is the signal EMA instance.
	SignalEma *trend.Ema[T]
}

// NewKvo function initializes a new KVO instance.
func NewKvo[T helper.Number]() *Kvo[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes channels of numbers and computes the Klinger Volume Oscillator.
// Returns kvo and signal.
func (k *Kvo[T]) Compute(highs, lows, volumes <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IdlePeriod is the initial period that KVO won't yield any results.
func (k *Kvo[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

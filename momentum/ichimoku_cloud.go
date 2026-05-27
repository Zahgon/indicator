// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultIchimokuCloudConversionPeriod is the default conversion period for the Ichimoku Cloud.
	DefaultIchimokuCloudConversionPeriod = 9

	// DefaultIchimokuCloudBasePeriod is the default base period for the Ichimoku Cloud.
	DefaultIchimokuCloudBasePeriod = 26

	// DefaultIchimokuCloudLeadingPeriod is the default leading period for the Ichimoku Cloud.
	DefaultIchimokuCloudLeadingPeriod = 52

	// DefaultIchimokuCloudLaggingPeriod is the default lagging period for the Ichimoku Cloud.
	DefaultIchimokuCloudLaggingPeriod = 26
)

// IchimokuCloud represents the configuration parameter for calculating the Ichimoku Cloud. It is also known as the
// Ichimoku Kinko Hyo, is a versatile indicator that defines support and resistance, identifies trend direction,
// gauges momentum, and provides trading signals.
//
//	Tenkan-sen (Conversion Line) = (9-Period High + 9-Period Low) / 2
//	Kijun-sen (Base Line) = (26-Period High + 26-Period Low) / 2
//	Senkou Span A (Leading Span A) = (Conversion Line + Base Line) / 2
//	Senkou Span B (Leading Span B) = (52-Period High + 52-Period Low) / 2
//	Chikou Span (Lagging Span) = Closing plotted 26 days in the past.
//
// Example:
//
//	ic := momentum.IchimokuCloud[float64]()
//	conversionLine, baseLine, leadingSpanA, leasingSpanB, laggingSpan := ic.Compute(highs, lows, closings)
type IchimokuCloud[T helper.Number] struct {
	// ConversionMax is the conversion Moving Max instance.
	ConversionMax *trend.MovingMax[T]

	// ConversionMin is the conversion Moving Min instance.
	ConversionMin *trend.MovingMin[T]

	// BaseMax is the base Moving Max instance.
	BaseMax *trend.MovingMax[T]

	// BaseMin is the base Moving Min instance.
	BaseMin *trend.MovingMin[T]

	// LeadingMax is the leading Moving Max instance.
	LeadingMax *trend.MovingMax[T]

	// LeadingMin is the leading Moving Min instance.
	LeadingMin *trend.MovingMin[T]

	// LaggingPeriod is the lagging period.
	LaggingPeriod int
}

// NewIchimokuCloud function initializes a new Ichimoku Cloud instance.
func NewIchimokuCloud[T helper.Number]() *IchimokuCloud[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the Ichimoku Cloud.
// Returns conversionLine, baseLine, leadingSpanA, leadingSpanB, laggingSpan
func (i *IchimokuCloud[T]) Compute(highs, lows, closings <-chan T) (<-chan T, <-chan T, <-chan T, <-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

//	Tenkan-sen (Conversion Line) = (9-Period High + 9-Period Low) / 2

//	Kijun-sen (Base Line) = (26-Period High + 26-Period Low) / 2

//	Senkou Span A (Leading Span A) = (Conversion Line + Base Line) / 2

//	Senkou Span B (Leading Span B) = (52-Period High + 52-Period Low) / 2

//	Chikou Span (Lagging Span) = Closing plotted 26 days in the past.

// IdlePeriod is the initial period that Ichimoku Cloud won't yield any results.
func (i *IchimokuCloud[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

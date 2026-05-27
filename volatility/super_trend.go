// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volatility

import (
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultSuperTrendPeriod is the default period value.
	DefaultSuperTrendPeriod = 14

	// DefaultSuperTrendMultiplier is the default multiplier value.
	DefaultSuperTrendMultiplier = 2.5
)

// SuperTrend represents the configuration parameters for calculating the Super Trend.
//
//	BasicUpperBands = (High + Low) / 2 + Multiplier * ATR
//	BasicLowerBands = (High + Low) / 2 - Multiplier * ATR
//	FinalUpperBands = If (BasicUpperBand < PreviousFinalUpperBand)
//	                  Or (PreviousClose > PreviousFinalUpperBand)
//	                  Then BasicUpperBand Else PreviousFinalUpperBand
//	FinalLowerBands = If (BasicLowerBand > PreviousFinalLowerBand)
//	                  Or (PreviousClose < PreviousFinalLowerBand)
//	                  Then BasicLowerBand Else PreviousFinalLowerBand
//	SuperTrend = If upTrend
//				 Then
//	               If (Close <= FinalUpperBand) Then FinalUpperBand Else FinalLowerBand
//	             Else
//	               If (Close >= FinalLowerBand) Then FinalLowerBand Else FinalUpperBand
//
//	UpTrend = If (SuperTrend == FinalUpperBand) Then True Else False
//
// Example:
type SuperTrend[T helper.Number] struct {
	Atr        *Atr[T]
	Multiplier T
}

// NewSuperTrend function initializes a new Super Trend instance with the default parameters.
func NewSuperTrend[T helper.Number]() *SuperTrend[T] { _ = "STUB: not implemented"; return nil }

// NewSuperTrendWithPeriod initializes a new Super Trend instance with the given period and multiplier.
func NewSuperTrendWithPeriod[T helper.Number](period int, multiplier T) *SuperTrend[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewSuperTrendWithMa function initializes a new Super Trend instance with the given moving average instance
// and multiplier.
func NewSuperTrendWithMa[T helper.Number](ma trend.Ma[T], multiplier T) *SuperTrend[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function calculates the Super Trend, using separate channels for highs, lows, and closings.
func (s *SuperTrend[T]) Compute(highs, lows, closings <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

//	BasicUpperBands = (High + Low) / 2 + Multiplier * ATR

//	BasicLowerBands = (High + Low) / 2 - Multiplier * ATR

//	FinalUpperBands = If (BasicUpperBand < PreviousFinalUpperBand)
//	                  Or (PreviousClose > PreviousFinalUpperBand)
//	                  Then BasicUpperBand Else PreviousFinalUpperBand

//	FinalLowerBands = If (BasicLowerBand > PreviousFinalLowerBand)
//	                  Or (PreviousClose < PreviousFinalLowerBand)
//	                  Then BasicLowerBand Else PreviousFinalLowerBand

//	SuperTrend = If upTrend
//				 Then
//	               If (Close <= FinalUpperBand) Then FinalUpperBand Else FinalLowerBand
//	             Else
//	               If (Close >= FinalLowerBand) Then FinalLowerBand Else FinalUpperBand
//
//	UpTrend = If (SuperTrend == FinalUpperBand) Then True Else False

// IdlePeriod is the initial period that Super Trend won't yield any results.
func (s *SuperTrend[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package momentum

import (
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultCoppockCurveRocPeriod1 is the default first ROC period.
	DefaultCoppockCurveRocPeriod1 = 14

	// DefaultCoppockCurveRocPeriod2 is the default second ROC period.
	DefaultCoppockCurveRocPeriod2 = 11

	// DefaultCoppockCurveWmaPeriod is the default WMA period.
	DefaultCoppockCurveWmaPeriod = 10
)

// CoppockCurve represents the configuration parameters for calculating the
// Coppock Curve oscillator. The Coppock Curve is a momentum indicator
// used to identify long-term buying opportunities in equity indices.
//
//	Coppock Curve = WMA(ROC(14) + ROC(11), 10)
//
// Example:
//
//	cc := momentum.NewCoppockCurve[float64]()
//	result := cc.Compute(closings)
type CoppockCurve[T helper.Float] struct {
	// RocPeriod1 is the first ROC period.
	RocPeriod1 int

	// RocPeriod2 is the second ROC period.
	RocPeriod2 int

	// WmaPeriod is the WMA period for smoothing.
	WmaPeriod int
}

// NewCoppockCurve function initializes a new CoppockCurve instance with default parameters.
func NewCoppockCurve[T helper.Float]() *CoppockCurve[T] { _ = "STUB: not implemented"; return nil }

// NewCoppockCurveWithPeriods function initializes a new CoppockCurve instance with the given periods.
func NewCoppockCurveWithPeriods[T helper.Float](rocPeriod1, rocPeriod2, wmaPeriod int) *CoppockCurve[T] {
	_ = "STUB: not implemented"
	return nil
}

// Compute function takes a channel of closings and computes the Coppock Curve.
func (c *CoppockCurve[T]) Compute(values <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// Align ROC streams. Both Compute calls return streams that have already skipped their own IdlePeriod.
// To align them to maxRocPeriod, we skip the remaining difference.

// IdlePeriod is the initial period that Coppock Curve won't yield any results.
func (c *CoppockCurve[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }

// maxRocPeriod (from ROC) + (WmaPeriod - 1) (from WMA)

// String is the string representation of the Coppock Curve.
func (c *CoppockCurve[T]) String() string { _ = "STUB: not implemented"; return "" }

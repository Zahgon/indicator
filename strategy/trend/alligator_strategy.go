// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/trend"
)

const (
	// DefaultAlligatorStrategyJawPeriod is the default jaw period of 13.
	DefaultAlligatorStrategyJawPeriod = 13

	// DefaultAlligatorStrategyTeethPeriod is the default teeth period of 8.
	DefaultAlligatorStrategyTeethPeriod = 8

	// DefaultAlligatorStrategyLipPeriod is the default lip period of 5.
	DefaultAlligatorStrategyLipPeriod = 5
)

// AlligatorStrategy represents the configuration parameters for calculating the
// Alligator strategy. It is a technical indicator to help identify the presence
// and the direction of the trend. It uses three Smooted Moving Averges (SMMAs).
type AlligatorStrategy struct {
	// Jaw represents the slowest moving aveage.
	Jaw *trend.Smma[float64]

	// Teeth represents the medium moving average.
	Teeth *trend.Smma[float64]

	// Lip represents the fastest moving average.
	Lip *trend.Smma[float64]
}

// NewAlligatorStrategy function initializes a new Alligator strategy instance.
func NewAlligatorStrategy() *AlligatorStrategy { _ = "STUB: not implemented"; return nil }

// NewAlligatorStrategyWith function initializes a new Alligator strategy instance with the given parameters.
func NewAlligatorStrategyWith(jawPeriod, teethPeriod, lipPeriod int) *AlligatorStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (a *AlligatorStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (a *AlligatorStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Alligator strategy starts only after a full period.

// Report processes the provided asset snapshots and generates a
// report annotated with the recommended actions.
func (a *AlligatorStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> closings[0] -> closings
	//                 closings[1] -> jaw
	//                 closings[2] -> teeth
	//                 closings[3] -> lip
	// snapshots[2] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

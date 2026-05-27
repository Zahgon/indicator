// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/volume"
)

// ForceIndexStrategy represents the configuration parameters for calculating the Force Index strategy.
// It recommends a Buy action when it crosses above zero, and a Sell action when it crosses below zero.
type ForceIndexStrategy struct {
	// ForceIndex is the Force Index instance.
	ForceIndex *volume.Fi[float64]
}

// NewForceIndexStrategy function initializes a new Force Index strategy instance with the default parameters.
func NewForceIndexStrategy() *ForceIndexStrategy { _ = "STUB: not implemented"; return nil }

// NewForceIndexStrategyWith function initializes a new Force Index strategy instance with the given parameters.
func NewForceIndexStrategyWith(period int) *ForceIndexStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (f *ForceIndexStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (f *ForceIndexStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Force Index starts only after a full period.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (f *ForceIndexStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> closings[0] -> closings
	//                 closings[1] -> force index
	// snapshots[2] -> volumes
	// snapshots[3] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

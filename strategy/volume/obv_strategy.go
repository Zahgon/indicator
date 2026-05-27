// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package volume

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/trend"
	"github.com/cinar/indicator/v2/volume"
)

const (
	// DefaultObvStrategyPeriod is the default OBV strategy period.
	DefaultObvStrategyPeriod = 10
)

// ObvStrategy represents the configuration parameters for calculating the On-Balance Volume (OBV) strategy.
// Recommends a Buy action when OBV crosses above its SMA, and recommends a Sell action when OBV crosses below its SMA.
type ObvStrategy struct {
	// Obv is the OBV indicator instance.
	Obv *volume.Obv[float64]

	// Sma is the SMA indicator instance.
	Sma *trend.Sma[float64]
}

// NewObvStrategy function initializes a new OBV strategy instance with the default parameters.
func NewObvStrategy() *ObvStrategy { _ = "STUB: not implemented"; return nil }

// NewObvStrategyWith function initializes a new OBV strategy instance with the given period.
func NewObvStrategyWith(period int) *ObvStrategy { _ = "STUB: not implemented"; return nil }

// Name function returns the name of the strategy.
func (s *ObvStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute function processes the provided asset snapshots and generates a stream of actionable recommendations.
func (s *ObvStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Align OBV with SMA

// OBV starts after its idle period (0), but SMA starts after its idle period.

// Report function processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (s *ObvStrategy) Report(snapshots <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> closings (for report)
	// snapshots[2] -> closings (for obv)
	// snapshots[3] -> volumes  (for obv)
	// snapshots[4] -> actions / outcomes
	//
	return nil
}

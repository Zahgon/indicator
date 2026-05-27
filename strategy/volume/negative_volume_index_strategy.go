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
	// DefaultNegativeVolumeIndexStrategyEmaPeriod is the default EMA period of 255.
	DefaultNegativeVolumeIndexStrategyEmaPeriod = 255
)

// NegativeVolumeIndexStrategy represents the configuration parameters for calculating the Negative Volume Index
// strategy. Recommends a Buy action when it crosses below its EMA, recommends a Sell action when it crosses
// above its EMA, and recommends a Hold action otherwise.
type NegativeVolumeIndexStrategy struct {
	// NegativeVolumeIndex is the Negative Volume Index indicator instance.
	NegativeVolumeIndex *volume.Nvi[float64]

	// NegativeVolumeIndexEma is the Negative Volume Index EMA instance.
	NegativeVolumeIndexEma *trend.Ema[float64]
}

// NewNegativeVolumeIndexStrategy function initializes a new Negative Volume Index strategy instance with the
// default parameters.
func NewNegativeVolumeIndexStrategy() *NegativeVolumeIndexStrategy {
	_ = "STUB: not implemented"
	return nil
}

// NewNegativeVolumeIndexStrategyWith function initializes a new Negative Volume Index strategy instance with the
// given parameters.
func NewNegativeVolumeIndexStrategyWith(emaPeriod int) *NegativeVolumeIndexStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (n *NegativeVolumeIndexStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (n *NegativeVolumeIndexStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Negative Volume Index starts only after a full period.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (n *NegativeVolumeIndexStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> closings[0] -> closings
	//                 closings[1] -> negative volume index[0] -> negative volume index
	//                                negative volume index[1] -> negative volume index ema
	// snapshots[2] -> volumes
	// snapshots[3] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

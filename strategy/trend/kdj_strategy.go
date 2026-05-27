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

// KdjStrategy represents the configuration parameters for calculating the KDJ strategy.
// Generates BUY action when j value crosses above both k and d values.
// Generates SELL action when j value crosses below both k and d values.
type KdjStrategy struct {
	// Kdj represents the configuration parameters for calculating the KDJ.
	Kdj *trend.Kdj[float64]
}

// NewKdjStrategy function initializes a new KDJ strategy instance.
func NewKdjStrategy() *KdjStrategy { _ = "STUB: not implemented"; return nil }

// Name returns the name of the strategy.
func (*KdjStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a
// stream of actionable recommendations.
func (kdj *KdjStrategy) Compute(c <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Generates BUY action when j value crosses above both k and d values.

// Generates SELL action when j value crosses below both k and d values.

// KDJ starts only after a full period.

// Report processes the provided asset snapshots and generates a
// report annotated with the recommended actions.
func (kdj *KdjStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	//
	// snapshots[0] -> dates
	// snapshots[1] -> highs       |
	// snapshots[2] -> lows        |
	// snapshots[3] -> closings[1] |> kdj
	//                 closings[0] -> closings
	// snapshots[4] -> actions     -> annotations
	//              -> outcomes
	//
	return nil
}

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package decorator

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
)

// NoLossStrategy prevents selling an asset at a loss. It modifies the recommendations of another strategy to ensure
// that the asset is only sold if its value is at or above the original purchase price.
type NoLossStrategy struct {
	// InnertStrategy is the inner strategy.
	InnertStrategy strategy.Strategy
}

// NewNoLossStrategy function initializes a new no loss strategy instance.
func NewNoLossStrategy(innerStrategy strategy.Strategy) *NoLossStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (n *NoLossStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (n *NoLossStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// If action is Buy and the asset is not yet bought, buy it as recommended.

// If the action is sell and the asset was bought at a lower amount, sell it as recommended.

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (n *NoLossStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	return nil
}

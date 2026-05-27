// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package decorator

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
)

// InverseStrategy reverses the advice of another strategy. For example, if the original strategy suggests buying an
// asset, InverseStrategy would recommend selling it.
type InverseStrategy struct {
	// InnerStrategy is the inner strategy.
	InnerStrategy strategy.Strategy
}

// NewInverseStrategy function initializes a new inverse strategy instance.
func NewInverseStrategy(innerStrategy strategy.Strategy) *InverseStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (i *InverseStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (i *InverseStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	_ = "STUB: not implemented"
	return nil
}

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (i *InverseStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	return nil
}

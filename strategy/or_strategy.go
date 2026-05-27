// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package strategy

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
)

// OrStrategy emits actionable recommendations when **at least one** strategy in the group recommends an
// action **without any conflicting recommendations** from other strategies.
type OrStrategy struct {
	// Strategies are the group of strategies that will be consulted to make an actionable recommendation.
	Strategies []Strategy

	// name is the name of this group of strategies.
	name string
}

// NewOrStrategy function initializes an empty or strategies group with the given name.
func NewOrStrategy(name string, strategies ...Strategy) *OrStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (a *OrStrategy) Name() string {
	_ = "STUB: not implemented"

	// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
	return ""
}

func (a *OrStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan Action {
	_ = "STUB: not implemented"
	return nil
}

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (a *OrStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	return nil
}

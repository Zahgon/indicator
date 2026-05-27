// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package strategy

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
)

// MajorityStrategy emits actionable recommendations aligned with what the strategies in the group recommends.
type MajorityStrategy struct {
	// Strategies are the group of strategies that will be consulted to make an actionable recommendation.
	Strategies []Strategy

	// name is the name of this group of strategies.
	name string
}

// NewMajorityStrategy function initializes an empty majority strategies group with the given name.
func NewMajorityStrategy(name string) *MajorityStrategy { _ = "STUB: not implemented"; return nil }

// NewMajorityStrategyWith function initializes a majority strategies group with the given name and strategies.
func NewMajorityStrategyWith(name string, strategies []Strategy) *MajorityStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (a *MajorityStrategy) Name() string {
	_ = "STUB: not implemented"

	// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
	return ""
}

func (a *MajorityStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan Action {
	_ = "STUB: not implemented"
	return nil
}

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (a *MajorityStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	return nil
}

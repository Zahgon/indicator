// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package strategy

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
)

// SplitStrategy leverages two separate strategies. It utilizes the first strategy to identify potential Buy
// opportunities, and the second strategy to identify potential Sell opportunities. When there is a
// conflicting recommendation, returns Hold.
type SplitStrategy struct {
	// BuyStrategy is used to identify potential Buy opportunities.
	BuyStrategy Strategy

	// SellStrategy is used to identify potential Sell opportunities.
	SellStrategy Strategy
}

// NewSplitStrategy function initializes a new split strategy with the given parameters.
func NewSplitStrategy(buyStrategy, sellStrategy Strategy) *SplitStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the strategy.
func (s *SplitStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (s *SplitStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan Action {
	_ = "STUB: not implemented"
	return nil
}

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (s *SplitStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	return nil
}

// AllSplitStrategies performs a cartesian product operation on the given strategies, resulting in a collection
// containing all split strategies formed by combining individual buy and sell strategies.
func AllSplitStrategies(strategies []Strategy) []Strategy { _ = "STUB: not implemented"; return nil }

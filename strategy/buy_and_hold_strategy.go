// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package strategy

import (
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
)

// BuyAndHoldStrategy defines an investment approach of acquiring and
// indefinitely retaining an asset. This strategy primarily serves as
// a benchmark for evaluating the performance of alternative
// strategies against a baseline of passive asset ownership.
type BuyAndHoldStrategy struct {
}

// NewBuyAndHoldStrategy function initializes a new buy and hold strategy instance.
func NewBuyAndHoldStrategy() *BuyAndHoldStrategy { _ = "STUB: not implemented"; return nil }

// Name returns the name of the strategy.
func (*BuyAndHoldStrategy) Name() string { _ = "STUB: not implemented"; return "" }

// Compute processes the provided asset snapshots and generates a
// stream of actionable recommendations.
func (*BuyAndHoldStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan Action {
	_ = "STUB: not implemented"
	return nil
}

// Report processes the provided asset snapshots and generates a
// report annotated with the recommended actions.
func (b *BuyAndHoldStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	_ = "STUB: not implemented"
	return nil
}

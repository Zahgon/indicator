// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package backtest

import (
	// Go embed report template.
	_ "embed"
	"log/slog"

	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/strategy"
)

const (
	// DefaultWriteStrategyReports is the default state of writing individual strategy reports.
	DefaultWriteStrategyReports = true
)

//go:embed "html_report.tmpl"
var htmlReportTmpl string

//go:embed "html_asset_report.tmpl"
var htmlAssetReportTmpl string

// HTMLReport is the backtest HTML report.
type HTMLReport struct {
	Report

	// outputDir is the output directory for the generated reports.
	outputDir string

	// assetResults is the mapping from the asset name to strategy results.
	assetResults map[string][]*htmlReportResult

	// bestResults is the best results for each asset.
	bestResults []*htmlReportResult

	// WriteStrategyReports indicates whether the individual strategy reports should be generated.
	WriteStrategyReports bool

	// DateFormat is the date format that is used in the reports.
	DateFormat string

	// Logger is the slog logger instance.
	Logger *slog.Logger
}

// htmlReportResult encapsulates the outcome of running a strategy.
type htmlReportResult struct {
	// AssetName is the name of the asset.
	AssetName string

	// StrategyName is the name of the strategy.
	StrategyName string

	// Action is the last recommended action by the strategy.
	Action strategy.Action

	// Since indicates how long the current action recommendation has been in effect.
	Since int

	// Outcome is the effectiveness of applying the recommended actions.
	Outcome float64

	// Transactions is the number of transactions made by the strategy.
	Transactions int
}

// NewHTMLReport initializes a new HTML report instance.
func NewHTMLReport(outputDir string) *HTMLReport { _ = "STUB: not implemented"; return nil }

// Begin is called when the backtest starts.
func (h *HTMLReport) Begin(assetNames []string, _ []strategy.Strategy) error {
	_ = "STUB: not implemented"
	// Make sure that output directory exists.
	return nil
}

// AssetBegin is called when backtesting for the given asset begins.
func (h *HTMLReport) AssetBegin(name string, strategies []strategy.Strategy) error {
	_ = "STUB: not implemented"
	return nil
}

// Write writes the given strategy actions and outomes to the report.
func (h *HTMLReport) Write(assetName string, currentStrategy strategy.Strategy, snapshots <-chan *asset.Snapshot, actions <-chan strategy.Action, outcomes <-chan float64) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate inidividual strategy report.

// Get asset strategy results.

// Append current strategy result for the asset.

// AssetEnd is called when backtesting for the given asset ends.
func (h *HTMLReport) AssetEnd(name string) error { _ = "STUB: not implemented"; return nil }

// Sort the backtest results by the outcomes.

// Report the best result for the current asset.

// Write the asset report.

// End is called when the backtest ends.
func (h *HTMLReport) End() error {
	_ = "STUB: not implemented"
	// Sort the best results by the outcomes.
	return nil
}

// strategyReportFileName defines the HTML report file name for the given asset and strategy.
func (*HTMLReport) strategyReportFileName(assetName, strategyName string) string {
	_ = "STUB: not implemented"
	return ""
}

// writeAssetReport generates a detailed report for the asset, summarizing the backtest results.
func (h *HTMLReport) writeAssetReport(name string, results []*htmlReportResult) error {
	_ = "STUB: not implemented"
	return nil
}

// writeReport generates a detailed report for the best results for all the assets.
func (h *HTMLReport) writeReport() error { _ = "STUB: not implemented"; return nil }

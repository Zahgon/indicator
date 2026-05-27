// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package backtest

const (
	// HTMLReportBuilderName is the name for the HTML report builder.
	HTMLReportBuilderName = "html"
)

// ReportBuilderFunc defines a function to build a new report using the given configuration parameter.
type ReportBuilderFunc func(config string) (Report, error)

// reportBuilders provides mapping for the report builders.
var reportBuilders = map[string]ReportBuilderFunc{
	HTMLReportBuilderName: htmlReportBuilder,
}

// RegisterReportBuilder registers the given builder.
func RegisterReportBuilder(name string, builder ReportBuilderFunc) {
	_ = "STUB: not implemented"
	return
}

// NewReport builds a new report by the given name type and the configuration.
func NewReport(name, config string) (Report, error) {
	_ = "STUB: not implemented"
	return *new(Report), nil
}

// htmlReportBuilder builds a new HTML report instance.
func htmlReportBuilder(config string) (Report, error) {
	_ = "STUB: not implemented"
	return *new(Report), nil
}

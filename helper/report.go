// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import (
	// Go embed report template.
	_ "embed"
	"io"
	"time"
)

//go:embed "report.tmpl"
var reportTmpl string

const (
	// DefaultReportDateFormat is the default date format used in the report.
	DefaultReportDateFormat = "2006-01-02"
)

// ReportColumn defines the interface that all report data columns must implement.
// This interface ensures that different types of data columns can be used
// consistently within the report generation process.
type ReportColumn interface {
	// Name returns the name of the report column.
	Name() string

	// Type returns the data type of the report column.
	Type() string

	// Role returns the role of the report column.
	Role() string

	// Value returns the next data value for the report column.
	Value() string
}

// Report generates an HTML file containing an interactive chart that
// visually represents the provided data and annotations.
//
// The generated HTML file can be opened in a web browser to explore
// the data visually, interact with the chart elements, and view
// the associated annotations.
type Report struct {
	Title       string
	Date        <-chan time.Time
	Columns     []ReportColumn
	Views       [][]int
	DateFormat  string
	GeneratedOn string
}

// NewReport takes a channel of time as the time axis and returns a new
// instance of the Report struct. This instance can later be used to
// add data and annotations and subsequently generate a report.
func NewReport(title string, date <-chan time.Time) *Report { _ = "STUB: not implemented"; return nil }

// AddChart adds a new chart to the report and returns its unique
// identifier. This identifier can be used later to refer to the
// chart and add columns to it.
func (r *Report) AddChart() int { _ = "STUB: not implemented"; return 0 }

// AddColumn adds a new data column to the specified charts. If no
// chart is specified, it will be added to the main chart.
func (r *Report) AddColumn(column ReportColumn, charts ...int) { _ = "STUB: not implemented"; return }

// WriteToWriter writes the report content to the provided io.Writer.
// This allows the report to be sent to various destinations, such
// as a file, a network socket, or even the standard output.
func (r *Report) WriteToWriter(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// WriteToFile writes the generated report content to a file with
// the specified name. This allows users to conveniently save the
// report for later viewing or analysis.
func (r *Report) WriteToFile(fileName string) error { _ = "STUB: not implemented"; return nil }

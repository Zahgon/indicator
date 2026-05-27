// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// annotationReportColumn is the annotation report column struct.
type annotationReportColumn struct {
	ReportColumn
	values <-chan string
}

// NewAnnotationReportColumn returns a new instance of an annotation column for a report.
func NewAnnotationReportColumn(values <-chan string) ReportColumn {
	_ = "STUB: not implemented"
	return *new(ReportColumn)
}

// Name returns the name of the report column.
func (*annotationReportColumn) Name() string {
	_ = "STUB: not implemented"

	// Type returns number as the data type.
	return ""
}

func (*annotationReportColumn) Type() string {
	_ = "STUB: not implemented"

	// Role returns the role of the report column.
	return ""
}

func (*annotationReportColumn) Role() string { _ = "STUB: not implemented"; return "" }

// Value returns the next data value for the report column.
func (c *annotationReportColumn) Value() string { _ = "STUB: not implemented"; return "" }

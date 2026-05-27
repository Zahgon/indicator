// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// numericReportColumn is the number report column struct.
type numericReportColumn[T Number] struct {
	ReportColumn
	name   string
	values <-chan T
}

// NewNumericReportColumn returns a new instance of a numeric data column for a report.
func NewNumericReportColumn[T Number](name string, values <-chan T) ReportColumn {
	_ = "STUB: not implemented"
	return *new(ReportColumn)
}

// Name returns the name of the report column.
func (c *numericReportColumn[T]) Name() string {
	_ = "STUB: not implemented"

	// Type returns number as the data type.
	return ""
}

func (*numericReportColumn[T]) Type() string {
	_ = "STUB: not implemented"

	// Role returns the role of the report column.
	return ""
}

func (*numericReportColumn[T]) Role() string {
	_ = "STUB: not implemented"

	// Value returns the next data value for the report column.
	return ""
}

func (c *numericReportColumn[T]) Value() string { _ = "STUB: not implemented"; return "" }

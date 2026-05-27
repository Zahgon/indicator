// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import (
	"io"
)

// ChanToJSON converts a channel of values into JSON format and writes it to the specified writer.
//
// Example:
//
//	input := helper.SliceToChan([]int{2, 4, 6, 8})
//
//	var buffer bytes.Buffer
//	err := helper.ChanToJSON(input, &buffer)
//
//	fmt.Println(buffer.String())
//	// Output: [2,4,6,8,9]
func ChanToJSON[T any](c <-chan T, w io.Writer) error { _ = "STUB: not implemented"; return nil }

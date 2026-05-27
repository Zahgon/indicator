// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Skip skips the specified number of elements from the
// given channel of type T.
//
// Example:
//
//	c := helper.SliceToChan([]int{2, 4, 6, 8})
//	actual := helper.Skip(c, 2)
//	fmt.Println(helper.ChanToSlice(actual)) // [6, 8]
func Skip[T any](c <-chan T, count int) <-chan T { _ = "STUB: not implemented"; return nil }

// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Pow takes a channel of type T values and returns the element-wise
// base-value exponential of y.
//
// Example:
//
//	c := helper.SliceToChan([]int{2, 3, 5, 10})
//	squared := helper.Pow(c, 2)
//	fmt.Println(helper.ChanToSlice(squared)) // [4, 9, 25, 100]
func Pow[T Number](c <-chan T, y T) <-chan T { _ = "STUB: not implemented"; return nil }

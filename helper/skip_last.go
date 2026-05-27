package helper

// SkipLast skips the specified number of elements
// from the end of the given channel.
//
// Example:
//
//	c := helper.SliceToChan([]int{2, 4, 6, 8})
//	actual := helper.SkipLast(c, 2)
//	fmt.Println(helper.ChanToSlice(actual)) // [2, 4]
func SkipLast[T any](c <-chan T, count int) <-chan T { _ = "STUB: not implemented"; return nil }

// Buffer to hold the last "count" elements

// send the oldest value

// drop the last `count` elements automatically

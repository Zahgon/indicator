// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Window returns a channel that emits the passed function result
// within a sliding window of size w from the input channel c.
// Note: the slice is in the same order than in source channel
// but the 1st element may not be 0, use modulo window size if
// order is important.
func Window[T any](c <-chan T, f func([]T, int) T, w int) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

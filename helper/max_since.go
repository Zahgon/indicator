// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// MaxSince returns a channel of T indicating since when
// (number of previous values) the respective value was the maximum
// within the window of size w.
func MaxSince[T Number](c <-chan T, w int) <-chan T { _ = "STUB: not implemented"; return nil }

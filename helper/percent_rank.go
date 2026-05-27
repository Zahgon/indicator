// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// PercentRank returns a channel that emits the percentile rank
// of each value compared to the previous period-1 values.
// The rank is between 0 and 100.
func PercentRank[T Number](c <-chan T, period int) <-chan T { _ = "STUB: not implemented"; return nil }

// Shift: remove oldest, add new

// Count how many values are less than current

// SortedPercentRank returns a channel that emits the percentile rank
// by sorting the window values. This is more accurate but slower.
func SortedPercentRank[T Number](c <-chan T, period int) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

// Shift: remove oldest, add new

// Sort copy for ranking

// Binary search for rank

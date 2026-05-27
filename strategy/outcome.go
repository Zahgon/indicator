// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package strategy

import "github.com/cinar/indicator/v2/helper"

// Outcome simulates the potential result of executing the given actions based on the provided values.
func Outcome[T helper.Number](values <-chan T, actions <-chan Action) <-chan float64 {
	_ = "STUB: not implemented"
	return nil
}

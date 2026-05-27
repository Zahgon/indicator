// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Duplicate duplicates a given receive-only channel by reading each value coming out of
// that channel and sending them on requested number of new output channels.
//
// Example:
//
//	expected := helper.SliceToChan([]float64{-10, 20, -4, -5})
//	outputs := helper.Duplicates[float64](helper.SliceToChan(expected), 2)
//
//	fmt.Println(helper.ChanToSlice(outputs[0])) // [-10, 20, -4, -5]
//	fmt.Println(helper.ChanToSlice(outputs[1])) // [-10, 20, -4, -5]
func Duplicate[T any](input <-chan T, count int) []<-chan T {
	_ = "STUB: not implemented"
	// TODO(cinar): Find a way to cast as a directional channel array.
	return nil
}

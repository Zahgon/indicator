// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import "sync"

// Waitable increments the wait group before reading from the channel
// and signals completion when the channel is closed.
func Waitable[T any](wg *sync.WaitGroup, c <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

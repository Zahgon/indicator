// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

// Field extracts a specific field from a channel of struct pointers and
// delivers it through a new channel.
func Field[T, S any](c <-chan *S, name string) (<-chan T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

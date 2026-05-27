// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import (
	"io"
	"log/slog"
)

// JSONToChan reads values from the specified reader in JSON format into a channel of values.
func JSONToChan[T any](r io.Reader) <-chan T { _ = "STUB: not implemented"; return nil }

// JSONToChanWithLogger reads values from the specified reader in JSON format into a channel of values.
func JSONToChanWithLogger[T any](r io.Reader, logger *slog.Logger) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

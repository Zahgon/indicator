// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import (
	"math/bits"
	"reflect"
	"strconv"
)

// kindToBits maps numeric kinds (e.g., int, float) to their corresponding sizes in bits.
var kindToBits = map[reflect.Kind]int{
	reflect.Int:     strconv.IntSize,
	reflect.Int8:    8,
	reflect.Int16:   16,
	reflect.Int32:   32,
	reflect.Int64:   64,
	reflect.Uint:    bits.UintSize,
	reflect.Uint16:  16,
	reflect.Uint32:  32,
	reflect.Uint64:  64,
	reflect.Float32: 32,
	reflect.Float64: 64,
}

// setReflectValueFromBool assigns the parsed boolean value to the specified variable.
func setReflectValueFromBool(value reflect.Value, stringValue string) error {
	_ = "STUB: not implemented"
	return nil
}

// setReflectValueFromInt assigns the parsed integer value to the specified variable.
func setReflectValueFromInt(value reflect.Value, stringValue string, bitSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// setReflectValueFromUint assigns the parsed unsigned integer value to the specified variable.
func setReflectValueFromUint(value reflect.Value, stringValue string, bitSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// setReflectValueFromFloat assigns the parsed float value to the specified variable.
func setReflectValueFromFloat(value reflect.Value, stringValue string, bitSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// setReflectValueFromTime assigns the parsed unsigned float value to the specified variable.
func setReflectValueFromTime(value reflect.Value, stringValue, format string) error {
	_ = "STUB: not implemented"
	return nil
}

// setReflectValue assigns the parsed value to the specified variable.
func setReflectValue(value reflect.Value, stringValue, format string) error {
	_ = "STUB: not implemented"
	return nil
}

// getReflectValue returns the string representation of the given value.
func getReflectValue(value reflect.Value, format string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

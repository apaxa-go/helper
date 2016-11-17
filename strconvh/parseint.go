package strconvh

import (
	"github.com/apaxa-go/helper/mathh"
	"strconv"
)

const defaultIntegerBase = 10

// Signed integers

// ParseInt interprets a string s in 10-base and returns the corresponding value i (int) and error.
func ParseInt(s string) (int, error) {
	if valueInt64, err := strconv.ParseInt(s, defaultIntegerBase, mathh.IntBits); err == nil {
		return int(valueInt64), nil
	} else {
		return 0, err
	}
}

// ParseInt8 interprets a string s in 10-base and returns the corresponding value i (int8) and error.
func ParseInt8(stringValue string) (int8, error) {
	if valueInt64, err := strconv.ParseInt(stringValue, defaultIntegerBase, 8); err == nil {
		return int8(valueInt64), nil
	} else {
		return 0, err
	}
}

// ParseInt16 interprets a string s in 10-base and returns the corresponding value i (int16) and error.
func ParseInt16(stringValue string) (int16, error) {
	if valueInt64, err := strconv.ParseInt(stringValue, defaultIntegerBase, 16); err == nil {
		return int16(valueInt64), nil
	} else {
		return 0, err
	}
}

// ParseInt32 interprets a string s in 10-base and returns the corresponding value i (int32) and error.
func ParseInt32(stringValue string) (int32, error) {
	if valueInt64, err := strconv.ParseInt(stringValue, defaultIntegerBase, 32); err == nil {
		return int32(valueInt64), nil
	} else {
		return 0, err
	}
}

// ParseInt64 interprets a string s in 10-base and returns the corresponding value i (int64) and error.
func ParseInt64(stringValue string) (int64, error) {
	return strconv.ParseInt(stringValue, defaultIntegerBase, 64)
}

// Unsigned integers

// ParseUint interprets a string s in 10-base and returns the corresponding value i (uint) and error.
func ParseUint(stringValue string) (uint, error) {
	if valueUint64, err := strconv.ParseUint(stringValue, defaultIntegerBase, mathh.UintBits); err == nil {
		return uint(valueUint64), nil
	} else {
		return 0, err
	}
}

// ParseUint8 interprets a string s in 10-base and returns the corresponding value i (uint8) and error.
func ParseUint8(stringValue string) (uint8, error) {
	if valueUint64, err := strconv.ParseUint(stringValue, defaultIntegerBase, 8); err == nil {
		return uint8(valueUint64), nil
	} else {
		return 0, err
	}
}

// ParseUint16 interprets a string s in 10-base and returns the corresponding value i (uint16) and error.
func ParseUint16(stringValue string) (uint16, error) {
	if valueUint64, err := strconv.ParseUint(stringValue, defaultIntegerBase, 16); err == nil {
		return uint16(valueUint64), nil
	} else {
		return 0, err
	}
}

// ParseUint32 interprets a string s in 10-base and returns the corresponding value i (uint32) and error.
func ParseUint32(stringValue string) (uint32, error) {
	if valueUint64, err := strconv.ParseUint(stringValue, defaultIntegerBase, 32); err == nil {
		return uint32(valueUint64), nil
	} else {
		return 0, err
	}
}

// ParseUint64 interprets a string s in 10-base and returns the corresponding value i (uint64) and error.
func ParseUint64(stringValue string) (uint64, error) {
	return strconv.ParseUint(stringValue, defaultIntegerBase, 64)
}

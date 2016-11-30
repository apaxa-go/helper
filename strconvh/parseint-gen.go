//replacer:generated-file
package strconvh

import (
	"github.com/apaxa-go/helper/mathh"
	"strconv"
)


// ParseUint64 interprets a string s in 10-base and returns the corresponding value i (uint64) and error.
func ParseUint64(stringValue string) (uint64, error) {
	return strconv.ParseUint(stringValue, defaultIntegerBase, mathh.Uint64Bits)
}


// ParseInt interprets a string s in 10-base and returns the corresponding value i (int) and error.
func ParseInt(stringValue string) (int, error) {
	if value64, err := strconv.ParseInt(stringValue, defaultIntegerBase, mathh.IntBits); err == nil {
		return int(value64), nil
	} else {
		return 0, err
	}
}

// ParseInt8 interprets a string s in 10-base and returns the corresponding value i (int8) and error.
func ParseInt8(stringValue string) (int8, error) {
	if value64, err := strconv.ParseInt(stringValue, defaultIntegerBase, mathh.Int8Bits); err == nil {
		return int8(value64), nil
	} else {
		return 0, err
	}
}

// ParseInt16 interprets a string s in 10-base and returns the corresponding value i (int16) and error.
func ParseInt16(stringValue string) (int16, error) {
	if value64, err := strconv.ParseInt(stringValue, defaultIntegerBase, mathh.Int16Bits); err == nil {
		return int16(value64), nil
	} else {
		return 0, err
	}
}

// ParseUint interprets a string s in 10-base and returns the corresponding value i (uint) and error.
func ParseUint(stringValue string) (uint, error) {
	if value64, err := strconv.ParseUint(stringValue, defaultIntegerBase, mathh.UintBits); err == nil {
		return uint(value64), nil
	} else {
		return 0, err
	}
}

// ParseUint8 interprets a string s in 10-base and returns the corresponding value i (uint8) and error.
func ParseUint8(stringValue string) (uint8, error) {
	if value64, err := strconv.ParseUint(stringValue, defaultIntegerBase, mathh.Uint8Bits); err == nil {
		return uint8(value64), nil
	} else {
		return 0, err
	}
}

// ParseUint16 interprets a string s in 10-base and returns the corresponding value i (uint16) and error.
func ParseUint16(stringValue string) (uint16, error) {
	if value64, err := strconv.ParseUint(stringValue, defaultIntegerBase, mathh.Uint16Bits); err == nil {
		return uint16(value64), nil
	} else {
		return 0, err
	}
}

// ParseUint32 interprets a string s in 10-base and returns the corresponding value i (uint32) and error.
func ParseUint32(stringValue string) (uint32, error) {
	if value64, err := strconv.ParseUint(stringValue, defaultIntegerBase, mathh.Uint32Bits); err == nil {
		return uint32(value64), nil
	} else {
		return 0, err
	}
}

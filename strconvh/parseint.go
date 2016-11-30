package strconvh

import (
	"github.com/apaxa-go/helper/mathh"
	"strconv"
)

const defaultIntegerBase = 10

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old int64	Int64
//replacer:new uint64	Uint64

// ParseInt64 interprets a string s in 10-base and returns the corresponding value i (int64) and error.
func ParseInt64(stringValue string) (int64, error) {
	return strconv.ParseInt(stringValue, defaultIntegerBase, mathh.Int64Bits)
}

//replacer:replace
//replacer:old int32	Int32
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32

// ParseInt32 interprets a string s in 10-base and returns the corresponding value i (int32) and error.
func ParseInt32(stringValue string) (int32, error) {
	if value64, err := strconv.ParseInt(stringValue, defaultIntegerBase, mathh.Int32Bits); err == nil {
		return int32(value64), nil
	} else {
		return 0, err
	}
}

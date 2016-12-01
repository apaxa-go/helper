package strconvh

import (
	"github.com/apaxa-go/helper/mathh"
	"testing"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
import "runtime"

const (
	maxUint8Str = "255"
	minUint8Str = "0"
	maxInt8Str  = "127"
	minInt8Str  = "-128"

	maxUint16Str = "65535"
	minUint16Str = "0"
	maxInt16Str  = "32767"
	minInt16Str  = "-32768"

	maxUint32Str = "4294967295"
	minUint32Str = "0"
	maxInt32Str  = "2147483647"
	minInt32Str  = "-2147483648"

	maxUint64Str = "18446744073709551615"
	minUint64Str = "0"
	maxInt64Str  = "9223372036854775807"
	minInt64Str  = "-9223372036854775808"
)

var (
	maxUintStr string
	minUintStr = "0"
	maxIntStr  string
	minIntStr  string
)

func init() {
	switch runtime.GOARCH {
	case "386":
		maxUintStr = maxUint32Str
		maxIntStr = maxInt32Str
		minIntStr = minInt32Str
	case "amd64":
		maxUintStr = maxUint64Str
		maxIntStr = maxInt64Str
		minIntStr = minInt64Str
	default:
		panic("Test does not support this platform: " + runtime.GOARCH)
	}
}

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32
//replacer:new int64	Int64

func TestFormatUint64(t *testing.T) {
	type testElement struct {
		i uint64
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxUint64, maxUint64Str},
		{mathh.MinUint64, minUint64Str},
	}
	for _, v := range test {
		s := FormatUint64(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

package strconvh

import (
	"github.com/apaxa-go/helper/mathh"
	"testing"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32
//replacer:new uint64	Uint64

func TestParseInt64(t *testing.T) {

	type testElement struct {
		s   string
		i   int64
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minInt64Str, mathh.MinInt64, false},
		{maxInt64Str, mathh.MaxInt64, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseInt64(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

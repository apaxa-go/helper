package mathh

import (
	"testing"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32

var testsMinMaxUint64 = []struct {
	a, b     uint64
	min, max uint64
}{
	{a: 0, b: 0, min: 0, max: 0},
	{a: 0, b: 1, min: 0, max: 1},
	{a: 2, b: 1, min: 1, max: 2},
	{a: 2, b: 2, min: 2, max: 2},
	{a: 4, b: 1, min: 1, max: 4},
}

//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32

var testsMinMaxInt64 = []struct {
	a, b     int64
	min, max int64
}{
	{a: 0, b: 0, min: 0, max: 0},
	{a: 0, b: 1, min: 0, max: 1},
	{a: 2, b: 1, min: 1, max: 2},
	{a: 2, b: 2, min: 2, max: 2},
	{a: 4, b: 1, min: 1, max: 4},
	{a: 0, b: -1, min: -1, max: 0},
	{a: -2, b: -1, min: -2, max: -1},
	{a: -2, b: -2, min: -2, max: -2},
	{a: -4, b: -1, min: -4, max: -1},
	{a: -4, b: 1, min: -4, max: 1},
	{a: 4, b: -1, min: -1, max: 4},
}

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

func TestMin2Uint64(t *testing.T) {
	for _, test := range testsMinMaxUint64 {
		if r := Min2Uint64(test.a, test.b); r != test.min {
			t.Errorf("Error Min2Uint64(%v, %v) - expected %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Uint64(t *testing.T) {
	for _, test := range testsMinMaxUint64 {
		if r := Max2Uint64(test.a, test.b); r != test.max {
			t.Errorf("Error Max2Uint64(%v, %v) - expected %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

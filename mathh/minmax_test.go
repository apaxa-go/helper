package mathh

import (
	"math"
	"testing"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint		Uint
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
//replacer:new int		Int
//replacer:new int8		Int8
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
//replacer:old Uint64
//replacer:new Uint
//replacer:new Uint8
//replacer:new Uint16
//replacer:new Uint32
//replacer:new Int
//replacer:new Int8
//replacer:new Int16
//replacer:new Int32
//replacer:new Int64

func TestMin2Uint64(t *testing.T) {
	for _, test := range testsMinMaxUint64 {
		if r := Min2Uint64(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Uint64(t *testing.T) {
	for _, test := range testsMinMaxUint64 {
		if r := Max2Uint64(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

//replacer:replace
//replacer:old float32	Float32
//replacer:new float64	Float64

var piFloat32 = PositiveInfFloat32()
var niFloat32 = NegativeInfFloat32()
var nzFloat32 = NegativeZeroFloat32()
var nanFloat32 = NaNFloat32()

var testsMinMaxFloat32 = []struct {
	a, b     float32
	min, max float32
}{
	{a: 0, b: 0, min: 0, max: 0},
	{a: 1, b: 1, min: 1, max: 1},
	{a: -1, b: -1, min: -1, max: -1},
	{a: 0, b: 1, min: 0, max: 1},
	{a: 1, b: 0, min: 0, max: 1},
	{a: 0, b: -1, min: -1, max: 0},
	{a: -1, b: 0, min: -1, max: 0},
	{a: -1, b: 1, min: -1, max: 1},
	{a: 1, b: -1, min: -1, max: 1},
	{a: 1, b: 2, min: 1, max: 2},
	{a: 2, b: 1, min: 1, max: 2},
	{a: -1, b: -2, min: -2, max: -1},
	{a: -2, b: -1, min: -2, max: -1},

	{a: niFloat32, b: niFloat32, min: niFloat32, max: niFloat32},
	{a: niFloat32, b: 1, min: niFloat32, max: 1},
	{a: 1, b: niFloat32, min: niFloat32, max: 1},
	{a: niFloat32, b: nzFloat32, min: niFloat32, max: nzFloat32},
	{a: nzFloat32, b: niFloat32, min: niFloat32, max: nzFloat32},
	{a: niFloat32, b: piFloat32, min: niFloat32, max: piFloat32},
	{a: piFloat32, b: niFloat32, min: niFloat32, max: piFloat32},

	{a: nzFloat32, b: nzFloat32, min: nzFloat32, max: nzFloat32},
	{a: nzFloat32, b: 0, min: nzFloat32, max: 0},
	{a: 0, b: nzFloat32, min: nzFloat32, max: 0},
	{a: nzFloat32, b: 1, min: nzFloat32, max: 1},
	{a: 1, b: nzFloat32, min: nzFloat32, max: 1},
	{a: nzFloat32, b: -1, min: -1, max: nzFloat32},
	{a: -1, b: nzFloat32, min: -1, max: nzFloat32},
	{a: nzFloat32, b: piFloat32, min: nzFloat32, max: piFloat32},
	{a: piFloat32, b: nzFloat32, min: nzFloat32, max: piFloat32},

	{a: piFloat32, b: piFloat32, min: piFloat32, max: piFloat32},
	{a: piFloat32, b: 1, min: 1, max: piFloat32},
	{a: 1, b: piFloat32, min: 1, max: piFloat32},

	{a: nanFloat32, b: nanFloat32, min: nanFloat32, max: nanFloat32},
	{a: nanFloat32, b: niFloat32, min: nanFloat32, max: nanFloat32},
	{a: niFloat32, b: nanFloat32, min: nanFloat32, max: nanFloat32},
	{a: nanFloat32, b: nzFloat32, min: nanFloat32, max: nanFloat32},
	{a: nzFloat32, b: nanFloat32, min: nanFloat32, max: nanFloat32},
	{a: nanFloat32, b: 0, min: nanFloat32, max: nanFloat32},
	{a: 0, b: nanFloat32, min: nanFloat32, max: nanFloat32},
	{a: nanFloat32, b: 1, min: nanFloat32, max: nanFloat32},
	{a: 1, b: nanFloat32, min: nanFloat32, max: nanFloat32},
	{a: nanFloat32, b: piFloat32, min: nanFloat32, max: nanFloat32},
	{a: piFloat32, b: nanFloat32, min: nanFloat32, max: nanFloat32},
}

// this function is required because NaN!=NaN.
func isEqualFloat32(a, b float32) bool { return math.Float32bits(a) == math.Float32bits(b) }

func TestMin2Float32(t *testing.T) {
	for _, test := range testsMinMaxFloat32 {
		if r := Min2Float32(test.a, test.b); !isEqualFloat32(r, test.min) {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Float32(t *testing.T) {
	for _, test := range testsMinMaxFloat32 {
		if r := Max2Float32(test.a, test.b); !isEqualFloat32(r, test.max) {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

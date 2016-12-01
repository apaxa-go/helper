package mathh

import "testing"

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

func TestPowUint64(t *testing.T) {
	type testElement struct {
		a uint64
		b uint64
		r uint64
	}
	tests := []testElement{
		{a: 0, b: 0, r: 1}, // Undefined
		{a: 0, b: 1, r: 0},
		{a: 0, b: 10, r: 0},
		{a: 1, b: 0, r: 1},
		{a: 1, b: 1, r: 1},
		{a: 1, b: 10, r: 1},
		{a: 2, b: 2, r: 4},
		{a: 2, b: 3, r: 8},
		{a: 3, b: 3, r: 27},
		{a: 4, b: 3, r: 64},
	}
	for _, test := range tests {
		if r := PowUint64(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModUint64(t *testing.T) {
	type testElement struct {
		a uint64
		b uint64
		m uint64
		r uint64
	}
	tests := []testElement{
		{a: 0, b: 0, m: 2, r: 1},
		{a: 0, b: 1, m: 3, r: 0},
		{a: 0, b: 10, m: 4, r: 0},
		{a: 1, b: 0, m: 5, r: 1},
		{a: 1, b: 1, m: 6, r: 1},
		{a: 1, b: 10, m: 7, r: 1},
		{a: 2, b: 2, m: 5, r: 4},
		{a: 2, b: 3, m: 3, r: 2},
		{a: 3, b: 3, m: 7, r: 6},
		{a: 4, b: 3, m: 10, r: 4},
		{a: 8, b: 8, m: 4, r: 0},
	}
	for _, test := range tests {
		if r := PowModUint64(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

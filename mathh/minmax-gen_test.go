//replacer:generated-file
package mathh

import (
	"testing"
)


var testsMinMaxUint = []struct {
	a, b     uint
	min, max uint
}{
	{a: 0, b: 0, min: 0, max: 0},
	{a: 0, b: 1, min: 0, max: 1},
	{a: 2, b: 1, min: 1, max: 2},
	{a: 2, b: 2, min: 2, max: 2},
	{a: 4, b: 1, min: 1, max: 4},
}


var testsMinMaxUint8 = []struct {
	a, b     uint8
	min, max uint8
}{
	{a: 0, b: 0, min: 0, max: 0},
	{a: 0, b: 1, min: 0, max: 1},
	{a: 2, b: 1, min: 1, max: 2},
	{a: 2, b: 2, min: 2, max: 2},
	{a: 4, b: 1, min: 1, max: 4},
}


var testsMinMaxUint16 = []struct {
	a, b     uint16
	min, max uint16
}{
	{a: 0, b: 0, min: 0, max: 0},
	{a: 0, b: 1, min: 0, max: 1},
	{a: 2, b: 1, min: 1, max: 2},
	{a: 2, b: 2, min: 2, max: 2},
	{a: 4, b: 1, min: 1, max: 4},
}


var testsMinMaxUint32 = []struct {
	a, b     uint32
	min, max uint32
}{
	{a: 0, b: 0, min: 0, max: 0},
	{a: 0, b: 1, min: 0, max: 1},
	{a: 2, b: 1, min: 1, max: 2},
	{a: 2, b: 2, min: 2, max: 2},
	{a: 4, b: 1, min: 1, max: 4},
}


var testsMinMaxInt = []struct {
	a, b     int
	min, max int
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


var testsMinMaxInt8 = []struct {
	a, b     int8
	min, max int8
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


var testsMinMaxInt16 = []struct {
	a, b     int16
	min, max int16
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


var testsMinMaxInt32 = []struct {
	a, b     int32
	min, max int32
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


func TestMin2Uint(t *testing.T) {
	for _, test := range testsMinMaxUint {
		if r := Min2Uint(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Uint(t *testing.T) {
	for _, test := range testsMinMaxUint {
		if r := Max2Uint(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Uint8(t *testing.T) {
	for _, test := range testsMinMaxUint8 {
		if r := Min2Uint8(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Uint8(t *testing.T) {
	for _, test := range testsMinMaxUint8 {
		if r := Max2Uint8(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Uint16(t *testing.T) {
	for _, test := range testsMinMaxUint16 {
		if r := Min2Uint16(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Uint16(t *testing.T) {
	for _, test := range testsMinMaxUint16 {
		if r := Max2Uint16(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Uint32(t *testing.T) {
	for _, test := range testsMinMaxUint32 {
		if r := Min2Uint32(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Uint32(t *testing.T) {
	for _, test := range testsMinMaxUint32 {
		if r := Max2Uint32(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Int(t *testing.T) {
	for _, test := range testsMinMaxInt {
		if r := Min2Int(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Int(t *testing.T) {
	for _, test := range testsMinMaxInt {
		if r := Max2Int(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Int8(t *testing.T) {
	for _, test := range testsMinMaxInt8 {
		if r := Min2Int8(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Int8(t *testing.T) {
	for _, test := range testsMinMaxInt8 {
		if r := Max2Int8(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Int16(t *testing.T) {
	for _, test := range testsMinMaxInt16 {
		if r := Min2Int16(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Int16(t *testing.T) {
	for _, test := range testsMinMaxInt16 {
		if r := Max2Int16(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Int32(t *testing.T) {
	for _, test := range testsMinMaxInt32 {
		if r := Min2Int32(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Int32(t *testing.T) {
	for _, test := range testsMinMaxInt32 {
		if r := Max2Int32(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

func TestMin2Int64(t *testing.T) {
	for _, test := range testsMinMaxInt64 {
		if r := Min2Int64(test.a, test.b); r != test.min {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.min, r)
		}
	}
}

func TestMax2Int64(t *testing.T) {
	for _, test := range testsMinMaxInt64 {
		if r := Max2Int64(test.a, test.b); r != test.max {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.max, r)
		}
	}
}

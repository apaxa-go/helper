//replacer:generated-file
package mathh

import "testing"

func TestPowUint(t *testing.T) {
	type testElement struct {
		a uint
		b uint
		r uint
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
		if r := PowUint(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModUint(t *testing.T) {
	type testElement struct {
		a uint
		b uint
		m uint
		r uint
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
		if r := PowModUint(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowUint8(t *testing.T) {
	type testElement struct {
		a uint8
		b uint8
		r uint8
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
		if r := PowUint8(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModUint8(t *testing.T) {
	type testElement struct {
		a uint8
		b uint8
		m uint8
		r uint8
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
		if r := PowModUint8(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowUint16(t *testing.T) {
	type testElement struct {
		a uint16
		b uint16
		r uint16
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
		if r := PowUint16(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModUint16(t *testing.T) {
	type testElement struct {
		a uint16
		b uint16
		m uint16
		r uint16
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
		if r := PowModUint16(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowUint32(t *testing.T) {
	type testElement struct {
		a uint32
		b uint32
		r uint32
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
		if r := PowUint32(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModUint32(t *testing.T) {
	type testElement struct {
		a uint32
		b uint32
		m uint32
		r uint32
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
		if r := PowModUint32(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowInt(t *testing.T) {
	type testElement struct {
		a int
		b int
		r int
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
		if r := PowInt(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModInt(t *testing.T) {
	type testElement struct {
		a int
		b int
		m int
		r int
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
		if r := PowModInt(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowInt8(t *testing.T) {
	type testElement struct {
		a int8
		b int8
		r int8
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
		if r := PowInt8(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModInt8(t *testing.T) {
	type testElement struct {
		a int8
		b int8
		m int8
		r int8
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
		if r := PowModInt8(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowInt16(t *testing.T) {
	type testElement struct {
		a int16
		b int16
		r int16
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
		if r := PowInt16(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModInt16(t *testing.T) {
	type testElement struct {
		a int16
		b int16
		m int16
		r int16
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
		if r := PowModInt16(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowInt32(t *testing.T) {
	type testElement struct {
		a int32
		b int32
		r int32
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
		if r := PowInt32(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModInt32(t *testing.T) {
	type testElement struct {
		a int32
		b int32
		m int32
		r int32
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
		if r := PowModInt32(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

func TestPowInt64(t *testing.T) {
	type testElement struct {
		a int64
		b int64
		r int64
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
		if r := PowInt64(test.a, test.b); r != test.r {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestPowModInt64(t *testing.T) {
	type testElement struct {
		a int64
		b int64
		m int64
		r int64
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
		if r := PowModInt64(test.a, test.b, test.m); r != test.r {
			t.Errorf("%v,%v,%v: expect %v, got %v", test.a, test.b, test.m, test.r, r)
		}
	}
}

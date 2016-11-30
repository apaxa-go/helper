//replacer:generated-file
package mathh

import "testing"


type testDivideUint struct {
	a uint
	b uint
	r uint
}

var testsDivideUint = []testDivideUint{
	{a: 3, b: 7, r: 0},
	{a: 3, b: 6, r: 1},
	{a: 3, b: 5, r: 1},
	{a: 3, b: 4, r: 1},
	{a: 3, b: 3, r: 1},
	{a: 3, b: 2, r: 2},
	{a: 3, b: 1, r: 3},
}

func TestDivideUint(t *testing.T) {
	for _, test := range testsDivideUint {
		if r := DivideRoundUint(test.a, test.b); r != test.r {
			t.Errorf("Error DivideUint(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideUintOverflow(t *testing.T) {
	for _, a := range testsUint {
		for _, b := range testsUint {
			if b == 0 {
				continue
			}
			validR := uint(divideAsUBig(customU(a), customU(b)))
			r := DivideRoundUint(a, b)
			if r != validR {
				t.Errorf("Error Divide(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func BenchmarkDivideUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%(testsLenUint-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideUint8 struct {
	a uint8
	b uint8
	r uint8
}

var testsDivideUint8 = []testDivideUint8{
	{a: 3, b: 7, r: 0},
	{a: 3, b: 6, r: 1},
	{a: 3, b: 5, r: 1},
	{a: 3, b: 4, r: 1},
	{a: 3, b: 3, r: 1},
	{a: 3, b: 2, r: 2},
	{a: 3, b: 1, r: 3},
}

func TestDivideUint8(t *testing.T) {
	for _, test := range testsDivideUint8 {
		if r := DivideRoundUint8(test.a, test.b); r != test.r {
			t.Errorf("Error DivideUint8(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideUint8Overflow(t *testing.T) {
	for _, a := range testsUint8 {
		for _, b := range testsUint8 {
			if b == 0 {
				continue
			}
			validR := uint8(divideAsUBig(customU(a), customU(b)))
			r := DivideRoundUint8(a, b)
			if r != validR {
				t.Errorf("Error Divide(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func BenchmarkDivideUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%(testsLenUint8-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideUint16 struct {
	a uint16
	b uint16
	r uint16
}

var testsDivideUint16 = []testDivideUint16{
	{a: 3, b: 7, r: 0},
	{a: 3, b: 6, r: 1},
	{a: 3, b: 5, r: 1},
	{a: 3, b: 4, r: 1},
	{a: 3, b: 3, r: 1},
	{a: 3, b: 2, r: 2},
	{a: 3, b: 1, r: 3},
}

func TestDivideUint16(t *testing.T) {
	for _, test := range testsDivideUint16 {
		if r := DivideRoundUint16(test.a, test.b); r != test.r {
			t.Errorf("Error DivideUint16(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideUint16Overflow(t *testing.T) {
	for _, a := range testsUint16 {
		for _, b := range testsUint16 {
			if b == 0 {
				continue
			}
			validR := uint16(divideAsUBig(customU(a), customU(b)))
			r := DivideRoundUint16(a, b)
			if r != validR {
				t.Errorf("Error Divide(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func BenchmarkDivideUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%(testsLenUint16-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideUint32 struct {
	a uint32
	b uint32
	r uint32
}

var testsDivideUint32 = []testDivideUint32{
	{a: 3, b: 7, r: 0},
	{a: 3, b: 6, r: 1},
	{a: 3, b: 5, r: 1},
	{a: 3, b: 4, r: 1},
	{a: 3, b: 3, r: 1},
	{a: 3, b: 2, r: 2},
	{a: 3, b: 1, r: 3},
}

func TestDivideUint32(t *testing.T) {
	for _, test := range testsDivideUint32 {
		if r := DivideRoundUint32(test.a, test.b); r != test.r {
			t.Errorf("Error DivideUint32(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideUint32Overflow(t *testing.T) {
	for _, a := range testsUint32 {
		for _, b := range testsUint32 {
			if b == 0 {
				continue
			}
			validR := uint32(divideAsUBig(customU(a), customU(b)))
			r := DivideRoundUint32(a, b)
			if r != validR {
				t.Errorf("Error Divide(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func BenchmarkDivideUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%(testsLenUint32-1)+1]) // -+1 is to avoid division by zero
	}
}

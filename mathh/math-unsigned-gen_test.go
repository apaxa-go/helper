//replacer:generated-file

package mathh

import "testing"

type testDivideUint struct {
	a     uint
	b     uint
	round uint
	up    uint
	down  uint
}

var testsDivideUint = []testDivideUint{
	{a: 3, b: 1, round: 3, up: 3, down: 3}, // 3
	{a: 3, b: 2, round: 2, up: 2, down: 1}, // 1.5
	{a: 3, b: 3, round: 1, up: 1, down: 1}, // 1
	{a: 3, b: 4, round: 1, up: 1, down: 0}, // 0.75
	{a: 3, b: 5, round: 1, up: 1, down: 0}, // 0.6
	{a: 3, b: 6, round: 1, up: 1, down: 0}, // 0.5
	{a: 3, b: 7, round: 0, up: 1, down: 0}, // 0.43...
	{a: 0, b: 7, round: 0, up: 0, down: 0}, // 0
}

func TestDivideRoundUint(t *testing.T) {
	for _, test := range testsDivideUint {
		if r := DivideRoundUint(test.a, test.b); r != test.round {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.round, r)
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
				t.Errorf("%v, %v: expect %v, got %v, ", a, b, validR, r)
			}
		}
	}
}

func TestDivideCeilUint(t *testing.T) {
	for _, test := range testsDivideUint {
		if r := DivideCeilUint(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideFloorUint(t *testing.T) {
	for _, test := range testsDivideUint {
		if r := DivideFloorUint(test.a, test.b); r != test.down {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.down, r)
		}
	}
}

func TestDivideRafzUint(t *testing.T) {
	for _, test := range testsDivideUint {
		if r := DivideRafzUint(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideTruncUint(t *testing.T) {
	for _, test := range testsDivideUint {
		if r := DivideTruncUint(test.a, test.b); r != test.a/test.b {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkDivideUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%(testsLenUint-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideUint8 struct {
	a     uint8
	b     uint8
	round uint8
	up    uint8
	down  uint8
}

var testsDivideUint8 = []testDivideUint8{
	{a: 3, b: 1, round: 3, up: 3, down: 3}, // 3
	{a: 3, b: 2, round: 2, up: 2, down: 1}, // 1.5
	{a: 3, b: 3, round: 1, up: 1, down: 1}, // 1
	{a: 3, b: 4, round: 1, up: 1, down: 0}, // 0.75
	{a: 3, b: 5, round: 1, up: 1, down: 0}, // 0.6
	{a: 3, b: 6, round: 1, up: 1, down: 0}, // 0.5
	{a: 3, b: 7, round: 0, up: 1, down: 0}, // 0.43...
	{a: 0, b: 7, round: 0, up: 0, down: 0}, // 0
}

func TestDivideRoundUint8(t *testing.T) {
	for _, test := range testsDivideUint8 {
		if r := DivideRoundUint8(test.a, test.b); r != test.round {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.round, r)
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
				t.Errorf("%v, %v: expect %v, got %v, ", a, b, validR, r)
			}
		}
	}
}

func TestDivideCeilUint8(t *testing.T) {
	for _, test := range testsDivideUint8 {
		if r := DivideCeilUint8(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideFloorUint8(t *testing.T) {
	for _, test := range testsDivideUint8 {
		if r := DivideFloorUint8(test.a, test.b); r != test.down {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.down, r)
		}
	}
}

func TestDivideRafzUint8(t *testing.T) {
	for _, test := range testsDivideUint8 {
		if r := DivideRafzUint8(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideTruncUint8(t *testing.T) {
	for _, test := range testsDivideUint8 {
		if r := DivideTruncUint8(test.a, test.b); r != test.a/test.b {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkDivideUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%(testsLenUint8-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideUint16 struct {
	a     uint16
	b     uint16
	round uint16
	up    uint16
	down  uint16
}

var testsDivideUint16 = []testDivideUint16{
	{a: 3, b: 1, round: 3, up: 3, down: 3}, // 3
	{a: 3, b: 2, round: 2, up: 2, down: 1}, // 1.5
	{a: 3, b: 3, round: 1, up: 1, down: 1}, // 1
	{a: 3, b: 4, round: 1, up: 1, down: 0}, // 0.75
	{a: 3, b: 5, round: 1, up: 1, down: 0}, // 0.6
	{a: 3, b: 6, round: 1, up: 1, down: 0}, // 0.5
	{a: 3, b: 7, round: 0, up: 1, down: 0}, // 0.43...
	{a: 0, b: 7, round: 0, up: 0, down: 0}, // 0
}

func TestDivideRoundUint16(t *testing.T) {
	for _, test := range testsDivideUint16 {
		if r := DivideRoundUint16(test.a, test.b); r != test.round {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.round, r)
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
				t.Errorf("%v, %v: expect %v, got %v, ", a, b, validR, r)
			}
		}
	}
}

func TestDivideCeilUint16(t *testing.T) {
	for _, test := range testsDivideUint16 {
		if r := DivideCeilUint16(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideFloorUint16(t *testing.T) {
	for _, test := range testsDivideUint16 {
		if r := DivideFloorUint16(test.a, test.b); r != test.down {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.down, r)
		}
	}
}

func TestDivideRafzUint16(t *testing.T) {
	for _, test := range testsDivideUint16 {
		if r := DivideRafzUint16(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideTruncUint16(t *testing.T) {
	for _, test := range testsDivideUint16 {
		if r := DivideTruncUint16(test.a, test.b); r != test.a/test.b {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkDivideUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%(testsLenUint16-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideUint32 struct {
	a     uint32
	b     uint32
	round uint32
	up    uint32
	down  uint32
}

var testsDivideUint32 = []testDivideUint32{
	{a: 3, b: 1, round: 3, up: 3, down: 3}, // 3
	{a: 3, b: 2, round: 2, up: 2, down: 1}, // 1.5
	{a: 3, b: 3, round: 1, up: 1, down: 1}, // 1
	{a: 3, b: 4, round: 1, up: 1, down: 0}, // 0.75
	{a: 3, b: 5, round: 1, up: 1, down: 0}, // 0.6
	{a: 3, b: 6, round: 1, up: 1, down: 0}, // 0.5
	{a: 3, b: 7, round: 0, up: 1, down: 0}, // 0.43...
	{a: 0, b: 7, round: 0, up: 0, down: 0}, // 0
}

func TestDivideRoundUint32(t *testing.T) {
	for _, test := range testsDivideUint32 {
		if r := DivideRoundUint32(test.a, test.b); r != test.round {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.round, r)
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
				t.Errorf("%v, %v: expect %v, got %v, ", a, b, validR, r)
			}
		}
	}
}

func TestDivideCeilUint32(t *testing.T) {
	for _, test := range testsDivideUint32 {
		if r := DivideCeilUint32(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideFloorUint32(t *testing.T) {
	for _, test := range testsDivideUint32 {
		if r := DivideFloorUint32(test.a, test.b); r != test.down {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.down, r)
		}
	}
}

func TestDivideRafzUint32(t *testing.T) {
	for _, test := range testsDivideUint32 {
		if r := DivideRafzUint32(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideTruncUint32(t *testing.T) {
	for _, test := range testsDivideUint32 {
		if r := DivideTruncUint32(test.a, test.b); r != test.a/test.b {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkDivideUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%(testsLenUint32-1)+1]) // -+1 is to avoid division by zero
	}
}

//replacer:generated-file
package mathh

import "testing"


type testDivideInt struct {
	a     int
	b     int
	round int
	ceil  int
	floor int
	rafz  int
}

var testsDivideInt = []testDivideInt{
	{a: 3, b: 1, round: 3, ceil: 3, floor: 3, rafz: 3},      // 3
	{a: 3, b: 2, round: 2, ceil: 2, floor: 1, rafz: 2},      // 1.5
	{a: 3, b: 3, round: 1, ceil: 1, floor: 1, rafz: 1},      // 1
	{a: 3, b: 4, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.75
	{a: 3, b: 5, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.6
	{a: 3, b: 6, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.5
	{a: 3, b: 7, round: 0, ceil: 1, floor: 0, rafz: 1},      // 0.43...
	{a: 0, b: 7, round: 0, ceil: 0, floor: 0, rafz: 0},      // 0
	{a: -3, b: 7, round: 0, ceil: 0, floor: -1, rafz: -1},   // -0.43...
	{a: -3, b: 6, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.5
	{a: -3, b: 5, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.6
	{a: -3, b: 4, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.75
	{a: -3, b: 3, round: -1, ceil: -1, floor: -1, rafz: -1}, // -1
	{a: -3, b: 2, round: -2, ceil: -1, floor: -2, rafz: -2}, // -1.5
	{a: -3, b: 1, round: -3, ceil: -3, floor: -3, rafz: -3}, // -3
}

func init() {
	// Extend tests based on the following rules:
	// a/b=c => -a/-b=c
	var ts []testDivideInt
	for _, v := range testsDivideInt {
		if v.a != MinInt && v.b != MinInt {
			v.a, v.b = -v.a, -v.b
			ts = append(ts, v)
		}
	}
	testsDivideInt = append(testsDivideInt, ts...)
}

func TestAbsInt(t *testing.T) {
	for _, test := range testsInt {
		if test == MinInt {
			continue
		}
		if r := AbsInt(test); r < 0 || ((test < 0 && r != -test) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsInt(%v) - got %v", test, r)
		}
	}
}

func TestAbsFixInt(t *testing.T) {
	for _, test := range testsInt {
		if r := AbsFixInt(test); r < 0 || ((test < 0 && r != -test && (test != MinInt || r != MaxInt)) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsFixInt(%v) - got %v", test, r)
		}
	}
}

func TestAntiAbsInt(t *testing.T) {
	for _, test := range testsInt {
		if r := AntiAbsInt(test); r > 0 || ((test > 0 && r != -test) || (test <= 0 && r != test)) {
			t.Errorf("Error AntiAbsInt(%v) - got %v", test, r)
		}
	}
}

func TestDivideRoundInt(t *testing.T) {
	for _, test := range testsDivideInt {
		if r := DivideRoundInt(test.a, test.b); r != test.round {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideRoundFixInt(t *testing.T) {
	for _, test := range testsDivideInt {
		if r := DivideRoundFixInt(test.a, test.b); r != test.round {
			t.Errorf("Error DivideFixInt(%v, %v) - expected %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideIntOverflow(t *testing.T) {
	for _, a := range testsInt {
		for _, b := range testsInt {
			if b == 0 || (a == MinInt && b == -1) {
				continue
			}
			validR := int(divideAsBig(customI(a), customI(b)))
			r := DivideRoundInt(a, b)
			if r != validR {
				t.Errorf("Error DivideInt(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideFixIntOverflow(t *testing.T) {
	for _, a := range testsInt {
		for _, b := range testsInt {
			if b == 0 {
				continue
			}
			var validR int
			if a == MinInt && b == -1 {
				validR = MaxInt
			} else {
				validR = int(divideAsBig(customI(a), customI(b)))
			}
			r := DivideRoundFixInt(a, b)
			if r != validR {
				t.Errorf("Error DivideFixInt(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideCeilInt(t *testing.T) {
	for _, test := range testsDivideInt {
		if r := DivideCeilInt(test.a, test.b); r != test.ceil {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.ceil, r)
		}
	}
}

func TestDivideFloorInt(t *testing.T) {
	for _, test := range testsDivideInt {
		if r := DivideFloorInt(test.a, test.b); r != test.floor {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.floor, r)
		}
	}
}

func TestDivideRafzInt(t *testing.T) {
	for _, test := range testsDivideInt {
		if r := DivideRafzInt(test.a, test.b); r != test.rafz {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.rafz, r)
		}
	}
}

func TestDivideTruncInt(t *testing.T) {
	for _, test := range testsDivideInt {
		if r := DivideTruncInt(test.a, test.b); r != test.a/test.b {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkAbsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkAbsFixInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsFixInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkAntiAbsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AntiAbsInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkDivideInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%(testsLenInt-1)+1]) // -+1 is to avoid division by zero
	}
}

func BenchmarkDivideFixInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundFixInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%(testsLenInt-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideInt8 struct {
	a     int8
	b     int8
	round int8
	ceil  int8
	floor int8
	rafz  int8
}

var testsDivideInt8 = []testDivideInt8{
	{a: 3, b: 1, round: 3, ceil: 3, floor: 3, rafz: 3},      // 3
	{a: 3, b: 2, round: 2, ceil: 2, floor: 1, rafz: 2},      // 1.5
	{a: 3, b: 3, round: 1, ceil: 1, floor: 1, rafz: 1},      // 1
	{a: 3, b: 4, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.75
	{a: 3, b: 5, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.6
	{a: 3, b: 6, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.5
	{a: 3, b: 7, round: 0, ceil: 1, floor: 0, rafz: 1},      // 0.43...
	{a: 0, b: 7, round: 0, ceil: 0, floor: 0, rafz: 0},      // 0
	{a: -3, b: 7, round: 0, ceil: 0, floor: -1, rafz: -1},   // -0.43...
	{a: -3, b: 6, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.5
	{a: -3, b: 5, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.6
	{a: -3, b: 4, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.75
	{a: -3, b: 3, round: -1, ceil: -1, floor: -1, rafz: -1}, // -1
	{a: -3, b: 2, round: -2, ceil: -1, floor: -2, rafz: -2}, // -1.5
	{a: -3, b: 1, round: -3, ceil: -3, floor: -3, rafz: -3}, // -3
}

func init() {
	// Extend tests based on the following rules:
	// a/b=c => -a/-b=c
	var ts []testDivideInt8
	for _, v := range testsDivideInt8 {
		if v.a != MinInt8 && v.b != MinInt8 {
			v.a, v.b = -v.a, -v.b
			ts = append(ts, v)
		}
	}
	testsDivideInt8 = append(testsDivideInt8, ts...)
}

func TestAbsInt8(t *testing.T) {
	for _, test := range testsInt8 {
		if test == MinInt8 {
			continue
		}
		if r := AbsInt8(test); r < 0 || ((test < 0 && r != -test) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsInt8(%v) - got %v", test, r)
		}
	}
}

func TestAbsFixInt8(t *testing.T) {
	for _, test := range testsInt8 {
		if r := AbsFixInt8(test); r < 0 || ((test < 0 && r != -test && (test != MinInt8 || r != MaxInt8)) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsFixInt8(%v) - got %v", test, r)
		}
	}
}

func TestAntiAbsInt8(t *testing.T) {
	for _, test := range testsInt8 {
		if r := AntiAbsInt8(test); r > 0 || ((test > 0 && r != -test) || (test <= 0 && r != test)) {
			t.Errorf("Error AntiAbsInt8(%v) - got %v", test, r)
		}
	}
}

func TestDivideRoundInt8(t *testing.T) {
	for _, test := range testsDivideInt8 {
		if r := DivideRoundInt8(test.a, test.b); r != test.round {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideRoundFixInt8(t *testing.T) {
	for _, test := range testsDivideInt8 {
		if r := DivideRoundFixInt8(test.a, test.b); r != test.round {
			t.Errorf("Error DivideFixInt8(%v, %v) - expected %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideInt8Overflow(t *testing.T) {
	for _, a := range testsInt8 {
		for _, b := range testsInt8 {
			if b == 0 || (a == MinInt8 && b == -1) {
				continue
			}
			validR := int8(divideAsBig(customI(a), customI(b)))
			r := DivideRoundInt8(a, b)
			if r != validR {
				t.Errorf("Error DivideInt8(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideFixInt8Overflow(t *testing.T) {
	for _, a := range testsInt8 {
		for _, b := range testsInt8 {
			if b == 0 {
				continue
			}
			var validR int8
			if a == MinInt8 && b == -1 {
				validR = MaxInt8
			} else {
				validR = int8(divideAsBig(customI(a), customI(b)))
			}
			r := DivideRoundFixInt8(a, b)
			if r != validR {
				t.Errorf("Error DivideFixInt8(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideCeilInt8(t *testing.T) {
	for _, test := range testsDivideInt8 {
		if r := DivideCeilInt8(test.a, test.b); r != test.ceil {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.ceil, r)
		}
	}
}

func TestDivideFloorInt8(t *testing.T) {
	for _, test := range testsDivideInt8 {
		if r := DivideFloorInt8(test.a, test.b); r != test.floor {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.floor, r)
		}
	}
}

func TestDivideRafzInt8(t *testing.T) {
	for _, test := range testsDivideInt8 {
		if r := DivideRafzInt8(test.a, test.b); r != test.rafz {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.rafz, r)
		}
	}
}

func TestDivideTruncInt8(t *testing.T) {
	for _, test := range testsDivideInt8 {
		if r := DivideTruncInt8(test.a, test.b); r != test.a/test.b {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkAbsInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkAbsFixInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsFixInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkAntiAbsInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AntiAbsInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkDivideInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%(testsLenInt8-1)+1]) // -+1 is to avoid division by zero
	}
}

func BenchmarkDivideFixInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundFixInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%(testsLenInt8-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideInt16 struct {
	a     int16
	b     int16
	round int16
	ceil  int16
	floor int16
	rafz  int16
}

var testsDivideInt16 = []testDivideInt16{
	{a: 3, b: 1, round: 3, ceil: 3, floor: 3, rafz: 3},      // 3
	{a: 3, b: 2, round: 2, ceil: 2, floor: 1, rafz: 2},      // 1.5
	{a: 3, b: 3, round: 1, ceil: 1, floor: 1, rafz: 1},      // 1
	{a: 3, b: 4, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.75
	{a: 3, b: 5, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.6
	{a: 3, b: 6, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.5
	{a: 3, b: 7, round: 0, ceil: 1, floor: 0, rafz: 1},      // 0.43...
	{a: 0, b: 7, round: 0, ceil: 0, floor: 0, rafz: 0},      // 0
	{a: -3, b: 7, round: 0, ceil: 0, floor: -1, rafz: -1},   // -0.43...
	{a: -3, b: 6, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.5
	{a: -3, b: 5, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.6
	{a: -3, b: 4, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.75
	{a: -3, b: 3, round: -1, ceil: -1, floor: -1, rafz: -1}, // -1
	{a: -3, b: 2, round: -2, ceil: -1, floor: -2, rafz: -2}, // -1.5
	{a: -3, b: 1, round: -3, ceil: -3, floor: -3, rafz: -3}, // -3
}

func init() {
	// Extend tests based on the following rules:
	// a/b=c => -a/-b=c
	var ts []testDivideInt16
	for _, v := range testsDivideInt16 {
		if v.a != MinInt16 && v.b != MinInt16 {
			v.a, v.b = -v.a, -v.b
			ts = append(ts, v)
		}
	}
	testsDivideInt16 = append(testsDivideInt16, ts...)
}

func TestAbsInt16(t *testing.T) {
	for _, test := range testsInt16 {
		if test == MinInt16 {
			continue
		}
		if r := AbsInt16(test); r < 0 || ((test < 0 && r != -test) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsInt16(%v) - got %v", test, r)
		}
	}
}

func TestAbsFixInt16(t *testing.T) {
	for _, test := range testsInt16 {
		if r := AbsFixInt16(test); r < 0 || ((test < 0 && r != -test && (test != MinInt16 || r != MaxInt16)) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsFixInt16(%v) - got %v", test, r)
		}
	}
}

func TestAntiAbsInt16(t *testing.T) {
	for _, test := range testsInt16 {
		if r := AntiAbsInt16(test); r > 0 || ((test > 0 && r != -test) || (test <= 0 && r != test)) {
			t.Errorf("Error AntiAbsInt16(%v) - got %v", test, r)
		}
	}
}

func TestDivideRoundInt16(t *testing.T) {
	for _, test := range testsDivideInt16 {
		if r := DivideRoundInt16(test.a, test.b); r != test.round {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideRoundFixInt16(t *testing.T) {
	for _, test := range testsDivideInt16 {
		if r := DivideRoundFixInt16(test.a, test.b); r != test.round {
			t.Errorf("Error DivideFixInt16(%v, %v) - expected %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideInt16Overflow(t *testing.T) {
	for _, a := range testsInt16 {
		for _, b := range testsInt16 {
			if b == 0 || (a == MinInt16 && b == -1) {
				continue
			}
			validR := int16(divideAsBig(customI(a), customI(b)))
			r := DivideRoundInt16(a, b)
			if r != validR {
				t.Errorf("Error DivideInt16(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideFixInt16Overflow(t *testing.T) {
	for _, a := range testsInt16 {
		for _, b := range testsInt16 {
			if b == 0 {
				continue
			}
			var validR int16
			if a == MinInt16 && b == -1 {
				validR = MaxInt16
			} else {
				validR = int16(divideAsBig(customI(a), customI(b)))
			}
			r := DivideRoundFixInt16(a, b)
			if r != validR {
				t.Errorf("Error DivideFixInt16(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideCeilInt16(t *testing.T) {
	for _, test := range testsDivideInt16 {
		if r := DivideCeilInt16(test.a, test.b); r != test.ceil {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.ceil, r)
		}
	}
}

func TestDivideFloorInt16(t *testing.T) {
	for _, test := range testsDivideInt16 {
		if r := DivideFloorInt16(test.a, test.b); r != test.floor {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.floor, r)
		}
	}
}

func TestDivideRafzInt16(t *testing.T) {
	for _, test := range testsDivideInt16 {
		if r := DivideRafzInt16(test.a, test.b); r != test.rafz {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.rafz, r)
		}
	}
}

func TestDivideTruncInt16(t *testing.T) {
	for _, test := range testsDivideInt16 {
		if r := DivideTruncInt16(test.a, test.b); r != test.a/test.b {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkAbsInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkAbsFixInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsFixInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkAntiAbsInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AntiAbsInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkDivideInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%(testsLenInt16-1)+1]) // -+1 is to avoid division by zero
	}
}

func BenchmarkDivideFixInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundFixInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%(testsLenInt16-1)+1]) // -+1 is to avoid division by zero
	}
}

type testDivideInt32 struct {
	a     int32
	b     int32
	round int32
	ceil  int32
	floor int32
	rafz  int32
}

var testsDivideInt32 = []testDivideInt32{
	{a: 3, b: 1, round: 3, ceil: 3, floor: 3, rafz: 3},      // 3
	{a: 3, b: 2, round: 2, ceil: 2, floor: 1, rafz: 2},      // 1.5
	{a: 3, b: 3, round: 1, ceil: 1, floor: 1, rafz: 1},      // 1
	{a: 3, b: 4, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.75
	{a: 3, b: 5, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.6
	{a: 3, b: 6, round: 1, ceil: 1, floor: 0, rafz: 1},      // 0.5
	{a: 3, b: 7, round: 0, ceil: 1, floor: 0, rafz: 1},      // 0.43...
	{a: 0, b: 7, round: 0, ceil: 0, floor: 0, rafz: 0},      // 0
	{a: -3, b: 7, round: 0, ceil: 0, floor: -1, rafz: -1},   // -0.43...
	{a: -3, b: 6, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.5
	{a: -3, b: 5, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.6
	{a: -3, b: 4, round: -1, ceil: 0, floor: -1, rafz: -1},  // -0.75
	{a: -3, b: 3, round: -1, ceil: -1, floor: -1, rafz: -1}, // -1
	{a: -3, b: 2, round: -2, ceil: -1, floor: -2, rafz: -2}, // -1.5
	{a: -3, b: 1, round: -3, ceil: -3, floor: -3, rafz: -3}, // -3
}

func init() {
	// Extend tests based on the following rules:
	// a/b=c => -a/-b=c
	var ts []testDivideInt32
	for _, v := range testsDivideInt32 {
		if v.a != MinInt32 && v.b != MinInt32 {
			v.a, v.b = -v.a, -v.b
			ts = append(ts, v)
		}
	}
	testsDivideInt32 = append(testsDivideInt32, ts...)
}

func TestAbsInt32(t *testing.T) {
	for _, test := range testsInt32 {
		if test == MinInt32 {
			continue
		}
		if r := AbsInt32(test); r < 0 || ((test < 0 && r != -test) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsInt32(%v) - got %v", test, r)
		}
	}
}

func TestAbsFixInt32(t *testing.T) {
	for _, test := range testsInt32 {
		if r := AbsFixInt32(test); r < 0 || ((test < 0 && r != -test && (test != MinInt32 || r != MaxInt32)) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsFixInt32(%v) - got %v", test, r)
		}
	}
}

func TestAntiAbsInt32(t *testing.T) {
	for _, test := range testsInt32 {
		if r := AntiAbsInt32(test); r > 0 || ((test > 0 && r != -test) || (test <= 0 && r != test)) {
			t.Errorf("Error AntiAbsInt32(%v) - got %v", test, r)
		}
	}
}

func TestDivideRoundInt32(t *testing.T) {
	for _, test := range testsDivideInt32 {
		if r := DivideRoundInt32(test.a, test.b); r != test.round {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideRoundFixInt32(t *testing.T) {
	for _, test := range testsDivideInt32 {
		if r := DivideRoundFixInt32(test.a, test.b); r != test.round {
			t.Errorf("Error DivideFixInt32(%v, %v) - expected %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideInt32Overflow(t *testing.T) {
	for _, a := range testsInt32 {
		for _, b := range testsInt32 {
			if b == 0 || (a == MinInt32 && b == -1) {
				continue
			}
			validR := int32(divideAsBig(customI(a), customI(b)))
			r := DivideRoundInt32(a, b)
			if r != validR {
				t.Errorf("Error DivideInt32(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideFixInt32Overflow(t *testing.T) {
	for _, a := range testsInt32 {
		for _, b := range testsInt32 {
			if b == 0 {
				continue
			}
			var validR int32
			if a == MinInt32 && b == -1 {
				validR = MaxInt32
			} else {
				validR = int32(divideAsBig(customI(a), customI(b)))
			}
			r := DivideRoundFixInt32(a, b)
			if r != validR {
				t.Errorf("Error DivideFixInt32(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideCeilInt32(t *testing.T) {
	for _, test := range testsDivideInt32 {
		if r := DivideCeilInt32(test.a, test.b); r != test.ceil {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.ceil, r)
		}
	}
}

func TestDivideFloorInt32(t *testing.T) {
	for _, test := range testsDivideInt32 {
		if r := DivideFloorInt32(test.a, test.b); r != test.floor {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.floor, r)
		}
	}
}

func TestDivideRafzInt32(t *testing.T) {
	for _, test := range testsDivideInt32 {
		if r := DivideRafzInt32(test.a, test.b); r != test.rafz {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.rafz, r)
		}
	}
}

func TestDivideTruncInt32(t *testing.T) {
	for _, test := range testsDivideInt32 {
		if r := DivideTruncInt32(test.a, test.b); r != test.a/test.b {
			t.Errorf("Expected f(%v, %v) = %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkAbsInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkAbsFixInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsFixInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkAntiAbsInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AntiAbsInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkDivideInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%(testsLenInt32-1)+1]) // -+1 is to avoid division by zero
	}
}

func BenchmarkDivideFixInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundFixInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%(testsLenInt32-1)+1]) // -+1 is to avoid division by zero
	}
}

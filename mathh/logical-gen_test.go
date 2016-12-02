//replacer:generated-file
package mathh

import (
	"testing"
)

const (
	testSegmentsAmountInt = 10
	testsLenInt           = testSegmentsAmountInt*6 + 1
)

var testsInt = []int{0}

func initInt() {
	for i := int(0); i < testSegmentsAmountInt; i++ {
		testsInt = append(testsInt, MinInt+i, MinInt/2+i, -1-i, 1+i, MaxInt/2-i, MaxInt-i)
	}
	if testsLenInt != len(testsInt) {
		panic("Wrong tests len")
	}
}

const (
	testSegmentsAmountInt8 = 10
	testsLenInt8           = testSegmentsAmountInt8*6 + 1
)

var testsInt8 = []int8{0}

func initInt8() {
	for i := int8(0); i < testSegmentsAmountInt8; i++ {
		testsInt8 = append(testsInt8, MinInt8+i, MinInt8/2+i, -1-i, 1+i, MaxInt8/2-i, MaxInt8-i)
	}
	if testsLenInt8 != len(testsInt8) {
		panic("Wrong tests len")
	}
}

const (
	testSegmentsAmountInt16 = 10
	testsLenInt16           = testSegmentsAmountInt16*6 + 1
)

var testsInt16 = []int16{0}

func initInt16() {
	for i := int16(0); i < testSegmentsAmountInt16; i++ {
		testsInt16 = append(testsInt16, MinInt16+i, MinInt16/2+i, -1-i, 1+i, MaxInt16/2-i, MaxInt16-i)
	}
	if testsLenInt16 != len(testsInt16) {
		panic("Wrong tests len")
	}
}

const (
	testSegmentsAmountInt32 = 10
	testsLenInt32           = testSegmentsAmountInt32*6 + 1
)

var testsInt32 = []int32{0}

func initInt32() {
	for i := int32(0); i < testSegmentsAmountInt32; i++ {
		testsInt32 = append(testsInt32, MinInt32+i, MinInt32/2+i, -1-i, 1+i, MaxInt32/2-i, MaxInt32-i)
	}
	if testsLenInt32 != len(testsInt32) {
		panic("Wrong tests len")
	}
}

const (
	testSegmentsAmountUint = 10
	testsLenUint           = testSegmentsAmountUint * 3
)

var testsUint = []uint{}

func initUint() {
	for i := uint(0); i < testSegmentsAmountUint; i++ {
		testsUint = append(testsUint, MinUint+i, MaxUint/2-i, MaxUint-i)
	}
	if testsLenUint != len(testsUint) {
		panic("Wrong tests len")
	}
}

const (
	testSegmentsAmountUint8 = 10
	testsLenUint8           = testSegmentsAmountUint8 * 3
)

var testsUint8 = []uint8{}

func initUint8() {
	for i := uint8(0); i < testSegmentsAmountUint8; i++ {
		testsUint8 = append(testsUint8, MinUint8+i, MaxUint8/2-i, MaxUint8-i)
	}
	if testsLenUint8 != len(testsUint8) {
		panic("Wrong tests len")
	}
}

const (
	testSegmentsAmountUint16 = 10
	testsLenUint16           = testSegmentsAmountUint16 * 3
)

var testsUint16 = []uint16{}

func initUint16() {
	for i := uint16(0); i < testSegmentsAmountUint16; i++ {
		testsUint16 = append(testsUint16, MinUint16+i, MaxUint16/2-i, MaxUint16-i)
	}
	if testsLenUint16 != len(testsUint16) {
		panic("Wrong tests len")
	}
}

const (
	testSegmentsAmountUint32 = 10
	testsLenUint32           = testSegmentsAmountUint32 * 3
)

var testsUint32 = []uint32{}

func initUint32() {
	for i := uint32(0); i < testSegmentsAmountUint32; i++ {
		testsUint32 = append(testsUint32, MinUint32+i, MaxUint32/2-i, MaxUint32-i)
	}
	if testsLenUint32 != len(testsUint32) {
		panic("Wrong tests len")
	}
}

func TestBtoInt(t *testing.T) {
	tests := []struct {
		b bool
		r int
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoInt(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotInt(t *testing.T) {
	tests := []struct {
		i int
		r int
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotInt(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignInt(t *testing.T) {
	for _, test := range testsInt {
		sign := SignInt(test)
		negative := NegativeInt(test)
		notNegative := NotNegativeInt(test)
		positive := PositiveInt(test)
		notPositive := NotPositiveInt(test)
		zero := ZeroInt(test)
		notZero := NotZeroInt(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareInt(t *testing.T) {
	for _, a := range testsInt {
		for _, b := range testsInt {
			sameSign := SameSignInt(a, b)
			notSameSign := NotSameSignInt(a, b)
			equal := EqualInt(a, b)
			notEqual := NotEqualInt(a, b)
			greater := GreaterInt(a, b)
			notGreater := NotGreaterInt(a, b)
			less := LessInt(a, b)
			notLess := NotLessInt(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotInt(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoInt(false)
		BtoInt(true)
	}
}

func BenchmarkNotInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotInt(0)
		NotInt(1)
	}
}

func BenchmarkNegativeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkNotNegativeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkPositiveInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkNotPositiveInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkZeroInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkNotZeroInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkSignInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignInt(testsInt[i%testsLenInt])
	}
}

func BenchmarkSameSignInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%testsLenInt])
	}
}

func BenchmarkEqualInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%testsLenInt])
	}
}

func BenchmarkNotEqualInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%testsLenInt])
	}
}

func BenchmarkGreaterInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%testsLenInt])
	}
}

func BenchmarkNotGreaterInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%testsLenInt])
	}
}

func BenchmarkLessInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%testsLenInt])
	}
}

func BenchmarkNotLessInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessInt(testsInt[i%testsLenInt], testsInt[(i+testsLenInt/4)%testsLenInt])
	}
}
func TestBtoInt8(t *testing.T) {
	tests := []struct {
		b bool
		r int8
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoInt8(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotInt8(t *testing.T) {
	tests := []struct {
		i int8
		r int8
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotInt8(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignInt8(t *testing.T) {
	for _, test := range testsInt8 {
		sign := SignInt8(test)
		negative := NegativeInt8(test)
		notNegative := NotNegativeInt8(test)
		positive := PositiveInt8(test)
		notPositive := NotPositiveInt8(test)
		zero := ZeroInt8(test)
		notZero := NotZeroInt8(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareInt8(t *testing.T) {
	for _, a := range testsInt8 {
		for _, b := range testsInt8 {
			sameSign := SameSignInt8(a, b)
			notSameSign := NotSameSignInt8(a, b)
			equal := EqualInt8(a, b)
			notEqual := NotEqualInt8(a, b)
			greater := GreaterInt8(a, b)
			notGreater := NotGreaterInt8(a, b)
			less := LessInt8(a, b)
			notLess := NotLessInt8(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotInt8(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoInt8(false)
		BtoInt8(true)
	}
}

func BenchmarkNotInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotInt8(0)
		NotInt8(1)
	}
}

func BenchmarkNegativeInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkNotNegativeInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkPositiveInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkNotPositiveInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkZeroInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkNotZeroInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkSignInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignInt8(testsInt8[i%testsLenInt8])
	}
}

func BenchmarkSameSignInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%testsLenInt8])
	}
}

func BenchmarkEqualInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%testsLenInt8])
	}
}

func BenchmarkNotEqualInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%testsLenInt8])
	}
}

func BenchmarkGreaterInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%testsLenInt8])
	}
}

func BenchmarkNotGreaterInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%testsLenInt8])
	}
}

func BenchmarkLessInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%testsLenInt8])
	}
}

func BenchmarkNotLessInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessInt8(testsInt8[i%testsLenInt8], testsInt8[(i+testsLenInt8/4)%testsLenInt8])
	}
}
func TestBtoInt16(t *testing.T) {
	tests := []struct {
		b bool
		r int16
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoInt16(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotInt16(t *testing.T) {
	tests := []struct {
		i int16
		r int16
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotInt16(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignInt16(t *testing.T) {
	for _, test := range testsInt16 {
		sign := SignInt16(test)
		negative := NegativeInt16(test)
		notNegative := NotNegativeInt16(test)
		positive := PositiveInt16(test)
		notPositive := NotPositiveInt16(test)
		zero := ZeroInt16(test)
		notZero := NotZeroInt16(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareInt16(t *testing.T) {
	for _, a := range testsInt16 {
		for _, b := range testsInt16 {
			sameSign := SameSignInt16(a, b)
			notSameSign := NotSameSignInt16(a, b)
			equal := EqualInt16(a, b)
			notEqual := NotEqualInt16(a, b)
			greater := GreaterInt16(a, b)
			notGreater := NotGreaterInt16(a, b)
			less := LessInt16(a, b)
			notLess := NotLessInt16(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotInt16(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoInt16(false)
		BtoInt16(true)
	}
}

func BenchmarkNotInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotInt16(0)
		NotInt16(1)
	}
}

func BenchmarkNegativeInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkNotNegativeInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkPositiveInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkNotPositiveInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkZeroInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkNotZeroInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkSignInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignInt16(testsInt16[i%testsLenInt16])
	}
}

func BenchmarkSameSignInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%testsLenInt16])
	}
}

func BenchmarkEqualInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%testsLenInt16])
	}
}

func BenchmarkNotEqualInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%testsLenInt16])
	}
}

func BenchmarkGreaterInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%testsLenInt16])
	}
}

func BenchmarkNotGreaterInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%testsLenInt16])
	}
}

func BenchmarkLessInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%testsLenInt16])
	}
}

func BenchmarkNotLessInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessInt16(testsInt16[i%testsLenInt16], testsInt16[(i+testsLenInt16/4)%testsLenInt16])
	}
}
func TestBtoInt32(t *testing.T) {
	tests := []struct {
		b bool
		r int32
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoInt32(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotInt32(t *testing.T) {
	tests := []struct {
		i int32
		r int32
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotInt32(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignInt32(t *testing.T) {
	for _, test := range testsInt32 {
		sign := SignInt32(test)
		negative := NegativeInt32(test)
		notNegative := NotNegativeInt32(test)
		positive := PositiveInt32(test)
		notPositive := NotPositiveInt32(test)
		zero := ZeroInt32(test)
		notZero := NotZeroInt32(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareInt32(t *testing.T) {
	for _, a := range testsInt32 {
		for _, b := range testsInt32 {
			sameSign := SameSignInt32(a, b)
			notSameSign := NotSameSignInt32(a, b)
			equal := EqualInt32(a, b)
			notEqual := NotEqualInt32(a, b)
			greater := GreaterInt32(a, b)
			notGreater := NotGreaterInt32(a, b)
			less := LessInt32(a, b)
			notLess := NotLessInt32(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotInt32(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoInt32(false)
		BtoInt32(true)
	}
}

func BenchmarkNotInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotInt32(0)
		NotInt32(1)
	}
}

func BenchmarkNegativeInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkNotNegativeInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkPositiveInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkNotPositiveInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkZeroInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkNotZeroInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkSignInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignInt32(testsInt32[i%testsLenInt32])
	}
}

func BenchmarkSameSignInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%testsLenInt32])
	}
}

func BenchmarkEqualInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%testsLenInt32])
	}
}

func BenchmarkNotEqualInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%testsLenInt32])
	}
}

func BenchmarkGreaterInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%testsLenInt32])
	}
}

func BenchmarkNotGreaterInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%testsLenInt32])
	}
}

func BenchmarkLessInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%testsLenInt32])
	}
}

func BenchmarkNotLessInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessInt32(testsInt32[i%testsLenInt32], testsInt32[(i+testsLenInt32/4)%testsLenInt32])
	}
}
func TestBtoUint(t *testing.T) {
	tests := []struct {
		b bool
		r uint
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoUint(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotUint(t *testing.T) {
	tests := []struct {
		i uint
		r uint
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotUint(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignUint(t *testing.T) {
	for _, test := range testsUint {
		sign := SignUint(test)
		negative := NegativeUint(test)
		notNegative := NotNegativeUint(test)
		positive := PositiveUint(test)
		notPositive := NotPositiveUint(test)
		zero := ZeroUint(test)
		notZero := NotZeroUint(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareUint(t *testing.T) {
	for _, a := range testsUint {
		for _, b := range testsUint {
			sameSign := SameSignUint(a, b)
			notSameSign := NotSameSignUint(a, b)
			equal := EqualUint(a, b)
			notEqual := NotEqualUint(a, b)
			greater := GreaterUint(a, b)
			notGreater := NotGreaterUint(a, b)
			less := LessUint(a, b)
			notLess := NotLessUint(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotUint(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoUint(false)
		BtoUint(true)
	}
}

func BenchmarkNotUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotUint(0)
		NotUint(1)
	}
}

func BenchmarkNegativeUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeUint(testsUint[i%testsLenUint])
	}
}

func BenchmarkNotNegativeUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeUint(testsUint[i%testsLenUint])
	}
}

func BenchmarkPositiveUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveUint(testsUint[i%testsLenUint])
	}
}

func BenchmarkNotPositiveUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveUint(testsUint[i%testsLenUint])
	}
}

func BenchmarkZeroUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroUint(testsUint[i%testsLenUint])
	}
}

func BenchmarkNotZeroUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroUint(testsUint[i%testsLenUint])
	}
}

func BenchmarkSignUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignUint(testsUint[i%testsLenUint])
	}
}

func BenchmarkSameSignUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%testsLenUint])
	}
}

func BenchmarkEqualUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%testsLenUint])
	}
}

func BenchmarkNotEqualUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%testsLenUint])
	}
}

func BenchmarkGreaterUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%testsLenUint])
	}
}

func BenchmarkNotGreaterUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%testsLenUint])
	}
}

func BenchmarkLessUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%testsLenUint])
	}
}

func BenchmarkNotLessUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessUint(testsUint[i%testsLenUint], testsUint[(i+testsLenUint/4)%testsLenUint])
	}
}
func TestBtoUint8(t *testing.T) {
	tests := []struct {
		b bool
		r uint8
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoUint8(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotUint8(t *testing.T) {
	tests := []struct {
		i uint8
		r uint8
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotUint8(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignUint8(t *testing.T) {
	for _, test := range testsUint8 {
		sign := SignUint8(test)
		negative := NegativeUint8(test)
		notNegative := NotNegativeUint8(test)
		positive := PositiveUint8(test)
		notPositive := NotPositiveUint8(test)
		zero := ZeroUint8(test)
		notZero := NotZeroUint8(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareUint8(t *testing.T) {
	for _, a := range testsUint8 {
		for _, b := range testsUint8 {
			sameSign := SameSignUint8(a, b)
			notSameSign := NotSameSignUint8(a, b)
			equal := EqualUint8(a, b)
			notEqual := NotEqualUint8(a, b)
			greater := GreaterUint8(a, b)
			notGreater := NotGreaterUint8(a, b)
			less := LessUint8(a, b)
			notLess := NotLessUint8(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotUint8(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoUint8(false)
		BtoUint8(true)
	}
}

func BenchmarkNotUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotUint8(0)
		NotUint8(1)
	}
}

func BenchmarkNegativeUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeUint8(testsUint8[i%testsLenUint8])
	}
}

func BenchmarkNotNegativeUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeUint8(testsUint8[i%testsLenUint8])
	}
}

func BenchmarkPositiveUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveUint8(testsUint8[i%testsLenUint8])
	}
}

func BenchmarkNotPositiveUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveUint8(testsUint8[i%testsLenUint8])
	}
}

func BenchmarkZeroUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroUint8(testsUint8[i%testsLenUint8])
	}
}

func BenchmarkNotZeroUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroUint8(testsUint8[i%testsLenUint8])
	}
}

func BenchmarkSignUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignUint8(testsUint8[i%testsLenUint8])
	}
}

func BenchmarkSameSignUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%testsLenUint8])
	}
}

func BenchmarkEqualUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%testsLenUint8])
	}
}

func BenchmarkNotEqualUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%testsLenUint8])
	}
}

func BenchmarkGreaterUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%testsLenUint8])
	}
}

func BenchmarkNotGreaterUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%testsLenUint8])
	}
}

func BenchmarkLessUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%testsLenUint8])
	}
}

func BenchmarkNotLessUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessUint8(testsUint8[i%testsLenUint8], testsUint8[(i+testsLenUint8/4)%testsLenUint8])
	}
}
func TestBtoUint16(t *testing.T) {
	tests := []struct {
		b bool
		r uint16
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoUint16(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotUint16(t *testing.T) {
	tests := []struct {
		i uint16
		r uint16
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotUint16(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignUint16(t *testing.T) {
	for _, test := range testsUint16 {
		sign := SignUint16(test)
		negative := NegativeUint16(test)
		notNegative := NotNegativeUint16(test)
		positive := PositiveUint16(test)
		notPositive := NotPositiveUint16(test)
		zero := ZeroUint16(test)
		notZero := NotZeroUint16(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareUint16(t *testing.T) {
	for _, a := range testsUint16 {
		for _, b := range testsUint16 {
			sameSign := SameSignUint16(a, b)
			notSameSign := NotSameSignUint16(a, b)
			equal := EqualUint16(a, b)
			notEqual := NotEqualUint16(a, b)
			greater := GreaterUint16(a, b)
			notGreater := NotGreaterUint16(a, b)
			less := LessUint16(a, b)
			notLess := NotLessUint16(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotUint16(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoUint16(false)
		BtoUint16(true)
	}
}

func BenchmarkNotUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotUint16(0)
		NotUint16(1)
	}
}

func BenchmarkNegativeUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeUint16(testsUint16[i%testsLenUint16])
	}
}

func BenchmarkNotNegativeUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeUint16(testsUint16[i%testsLenUint16])
	}
}

func BenchmarkPositiveUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveUint16(testsUint16[i%testsLenUint16])
	}
}

func BenchmarkNotPositiveUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveUint16(testsUint16[i%testsLenUint16])
	}
}

func BenchmarkZeroUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroUint16(testsUint16[i%testsLenUint16])
	}
}

func BenchmarkNotZeroUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroUint16(testsUint16[i%testsLenUint16])
	}
}

func BenchmarkSignUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignUint16(testsUint16[i%testsLenUint16])
	}
}

func BenchmarkSameSignUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%testsLenUint16])
	}
}

func BenchmarkEqualUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%testsLenUint16])
	}
}

func BenchmarkNotEqualUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%testsLenUint16])
	}
}

func BenchmarkGreaterUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%testsLenUint16])
	}
}

func BenchmarkNotGreaterUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%testsLenUint16])
	}
}

func BenchmarkLessUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%testsLenUint16])
	}
}

func BenchmarkNotLessUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessUint16(testsUint16[i%testsLenUint16], testsUint16[(i+testsLenUint16/4)%testsLenUint16])
	}
}
func TestBtoUint32(t *testing.T) {
	tests := []struct {
		b bool
		r uint32
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoUint32(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotUint32(t *testing.T) {
	tests := []struct {
		i uint32
		r uint32
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotUint32(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignUint32(t *testing.T) {
	for _, test := range testsUint32 {
		sign := SignUint32(test)
		negative := NegativeUint32(test)
		notNegative := NotNegativeUint32(test)
		positive := PositiveUint32(test)
		notPositive := NotPositiveUint32(test)
		zero := ZeroUint32(test)
		notZero := NotZeroUint32(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareUint32(t *testing.T) {
	for _, a := range testsUint32 {
		for _, b := range testsUint32 {
			sameSign := SameSignUint32(a, b)
			notSameSign := NotSameSignUint32(a, b)
			equal := EqualUint32(a, b)
			notEqual := NotEqualUint32(a, b)
			greater := GreaterUint32(a, b)
			notGreater := NotGreaterUint32(a, b)
			less := LessUint32(a, b)
			notLess := NotLessUint32(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotUint32(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoUint32(false)
		BtoUint32(true)
	}
}

func BenchmarkNotUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotUint32(0)
		NotUint32(1)
	}
}

func BenchmarkNegativeUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeUint32(testsUint32[i%testsLenUint32])
	}
}

func BenchmarkNotNegativeUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeUint32(testsUint32[i%testsLenUint32])
	}
}

func BenchmarkPositiveUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveUint32(testsUint32[i%testsLenUint32])
	}
}

func BenchmarkNotPositiveUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveUint32(testsUint32[i%testsLenUint32])
	}
}

func BenchmarkZeroUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroUint32(testsUint32[i%testsLenUint32])
	}
}

func BenchmarkNotZeroUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroUint32(testsUint32[i%testsLenUint32])
	}
}

func BenchmarkSignUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignUint32(testsUint32[i%testsLenUint32])
	}
}

func BenchmarkSameSignUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%testsLenUint32])
	}
}

func BenchmarkEqualUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%testsLenUint32])
	}
}

func BenchmarkNotEqualUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%testsLenUint32])
	}
}

func BenchmarkGreaterUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%testsLenUint32])
	}
}

func BenchmarkNotGreaterUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%testsLenUint32])
	}
}

func BenchmarkLessUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%testsLenUint32])
	}
}

func BenchmarkNotLessUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessUint32(testsUint32[i%testsLenUint32], testsUint32[(i+testsLenUint32/4)%testsLenUint32])
	}
}
func TestBtoUint64(t *testing.T) {
	tests := []struct {
		b bool
		r uint64
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoUint64(test.b); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotUint64(t *testing.T) {
	tests := []struct {
		i uint64
		r uint64
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotUint64(test.i); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignUint64(t *testing.T) {
	for _, test := range testsUint64 {
		sign := SignUint64(test)
		negative := NegativeUint64(test)
		notNegative := NotNegativeUint64(test)
		positive := PositiveUint64(test)
		notPositive := NotPositiveUint64(test)
		zero := ZeroUint64(test)
		notZero := NotZeroUint64(test)

		if (test < 0 && sign >= 0) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("%v: got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("%v: got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("%v: got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("%v: got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("%v: got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("%v: got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("%v: got %v", test, notZero)
		}
	}
}

func TestAllCompareUint64(t *testing.T) {
	for _, a := range testsUint64 {
		for _, b := range testsUint64 {
			sameSign := SameSignUint64(a, b)
			notSameSign := NotSameSignUint64(a, b)
			equal := EqualUint64(a, b)
			notEqual := NotEqualUint64(a, b)
			greater := GreaterUint64(a, b)
			notGreater := NotGreaterUint64(a, b)
			less := LessUint64(a, b)
			notLess := NotLessUint64(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("%v,%v: got %v", a, b, sameSign)
			}
			if notSameSign != NotUint64(sameSign) {
				t.Errorf("%v,%v: got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("%v,%v: got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("%v,%v: got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("%v,%v: got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("%v,%v: got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("%v,%v: got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("%v,%v: got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoUint64(false)
		BtoUint64(true)
	}
}

func BenchmarkNotUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotUint64(0)
		NotUint64(1)
	}
}

func BenchmarkNegativeUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeUint64(testsUint64[i%testsLenUint64])
	}
}

func BenchmarkNotNegativeUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeUint64(testsUint64[i%testsLenUint64])
	}
}

func BenchmarkPositiveUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveUint64(testsUint64[i%testsLenUint64])
	}
}

func BenchmarkNotPositiveUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveUint64(testsUint64[i%testsLenUint64])
	}
}

func BenchmarkZeroUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroUint64(testsUint64[i%testsLenUint64])
	}
}

func BenchmarkNotZeroUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroUint64(testsUint64[i%testsLenUint64])
	}
}

func BenchmarkSignUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignUint64(testsUint64[i%testsLenUint64])
	}
}

func BenchmarkSameSignUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%testsLenUint64])
	}
}

func BenchmarkEqualUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%testsLenUint64])
	}
}

func BenchmarkNotEqualUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%testsLenUint64])
	}
}

func BenchmarkGreaterUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%testsLenUint64])
	}
}

func BenchmarkNotGreaterUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%testsLenUint64])
	}
}

func BenchmarkLessUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%testsLenUint64])
	}
}

func BenchmarkNotLessUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%testsLenUint64])
	}
}

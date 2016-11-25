package mathh

import (
	"testing"
)

func init() {
	initInt64()
	// TODO initInt32()
	// TODO initInt16()
	// TODO initInt8()
	// TODO initInt()
}

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32

const (
	testSegmentsAmountInt64 = 10
	testsLenInt64           = testSegmentsAmountInt64*6 + 1
)

var testsInt64 = []int64{0}

func initInt64() {
	for i := int64(0); i < testSegmentsAmountInt64; i++ {
		testsInt64 = append(testsInt64, MinInt64+i, MinInt64/2+i, -1-i, 1+i, MaxInt64/2-i, MaxInt64-i)
	}
	if testsLenInt64 != len(testsInt64) {
		panic("Wrong tests len")
	}
}

//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32

const (
	testSegmentsAmountUint64 = 10
	testsLenUint64           = testSegmentsAmountUint64 * 3
)

var testsUint64 = []uint64{}

func initUint64() {
	for i := int64(0); i < testSegmentsAmountUint64; i++ {
		testsInt64 = append(testsInt64, MinUint64+i, MaxUint64/2-i, MaxUint64-i)
	}
	if testsLenUint64 != len(testsUint64) {
		panic("Wrong tests len")
	}
}

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
func TestBtoInt64(t *testing.T) {
	tests := []struct {
		b bool
		r int64
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoInt64(test.b); r != test.r {
			t.Errorf("Error %v to int64 - expected %v, got %v", test.b, test.r, r)
		}
	}
}

func TestNotInt64(t *testing.T) {
	tests := []struct {
		i int64
		r int64
	}{{i: 0, r: 1}, {i: 1, r: 0}}
	for _, test := range tests {
		if r := NotInt64(test.i); r != test.r {
			t.Errorf("Error NotInt64(%v) - expected %v, got %v", test.i, test.r, r)
		}
	}
}

func TestAllSignInt64(t *testing.T) {
	for _, test := range testsInt64 {
		sign := SignInt64(test)
		negative := NegativeInt64(test)
		notNegative := NotNegativeInt64(test)
		positive := PositiveInt64(test)
		notPositive := NotPositiveInt64(test)
		zero := ZeroInt64(test)
		notZero := NotZeroInt64(test)

		if (test < 0 && sign != -1) || (test == 0 && sign != 0) || (test > 0 && sign != 1) {
			t.Errorf("Error SignInt64(%v) - got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("Error NegativeInt64(%v) - got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("Error NotNegativeInt64(%v) - got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("Error PositiveInt64(%v) - got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("Error NotPositiveInt64(%v) - got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("Error ZeroInt64(%v) - got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("Error NotZeroInt64(%v) - got %v", test, notZero)
		}
	}
}

func TestAllCompareInt64(t *testing.T) {
	for _, a := range testsInt64 {
		for _, b := range testsInt64 {
			sameSign := SameSignInt64(a, b)
			notSameSign := NotSameSignInt64(a, b)
			equal := EqualInt64(a, b)
			notEqual := NotEqualInt64(a, b)
			greater := GreaterInt64(a, b)
			notGreater := NotGreaterInt64(a, b)
			less := LessInt64(a, b)
			notLess := NotLessInt64(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("Error SameSignInt64(%v,%v) - got %v", a, b, sameSign)
			}
			if notSameSign != NotInt64(sameSign) {
				t.Errorf("Error NotSameSignInt64(%v,%v) - got %v", a, b, notSameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("Error EqualInt64(%v,%v) - got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("Error NotEqualInt64(%v,%v) - got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("Error GreaterInt64(%v,%v) - got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("Error NotGreateInt64(%v,%v) - got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("Error LessInt64(%v,%v) - got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("Error NotLessInt64(%v,%v) - got %v", a, b, notLess)
			}
		}
	}
}

func BenchmarkBtoInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BtoInt64(false)
		BtoInt64(true)
	}
}

func BenchmarkNotInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotInt64(0)
		NotInt64(1)
	}
}

func BenchmarkNegativeInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NegativeInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkNotNegativeInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkPositiveInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkNotPositiveInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkZeroInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkNotZeroInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkSignInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkSameSignInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%testsLenInt64])
	}
}

func BenchmarkEqualInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%testsLenInt64])
	}
}

func BenchmarkNotEqualInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%testsLenInt64])
	}
}

func BenchmarkGreaterInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%testsLenInt64])
	}
}

func BenchmarkNotGreaterInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%testsLenInt64])
	}
}

func BenchmarkLessInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%testsLenInt64])
	}
}

func BenchmarkNotLessInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%testsLenInt64])
	}
}

package mathh

import (
	"math"
	"testing"
)

const testSegmentsLen = 5
const testsLen = testSegmentsLen*6 + 1

var testsInt64 = []int64{0}

func init() {
	for i := int64(0); i < testSegmentsLen; i++ {
		testsInt64 = append(testsInt64, math.MinInt64+i, math.MinInt64/2+i, -1-i, 1+i, math.MaxInt64-i, math.MaxInt64/2-i)
	}
	if testsLen != len(testsInt64) {
		panic("Wrong tests len")
	}
}

func TestBtoInt64(t *testing.T) {
	tests := []struct {
		b bool
		r int64
	}{{b: false, r: 0}, {b: true, r: 1}}
	for _, test := range tests {
		if r := BtoInt64(test.b); r != test.r {
			t.Errorf("Error %v to int - expected %v, got %v", test.b, test.r, r)
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
			t.Errorf("Error Not(%v) - expected %v, got %v", test.i, test.r, r)
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
			t.Errorf("Error Sign(%v) - got %v", test, sign)
		}
		if (test < 0 && negative != 1) || (test >= 0 && negative != 0) {
			t.Errorf("Error Negative(%v) - got %v", test, negative)
		}
		if (test >= 0 && notNegative != 1) || (test < 0 && notNegative != 0) {
			t.Errorf("Error NotNegative(%v) - got %v", test, notNegative)
		}
		if (test > 0 && positive != 1) || (test <= 0 && positive != 0) {
			t.Errorf("Error Positive(%v) - got %v", test, positive)
		}
		if (test <= 0 && notPositive != 1) || (test > 0 && notPositive != 0) {
			t.Errorf("Error NotPositive(%v) - got %v", test, notPositive)
		}
		if (test == 0 && zero != 1) || (test != 0 && zero != 0) {
			t.Errorf("Error Zero(%v) - got %v", test, zero)
		}
		if (test != 0 && notZero != 1) || (test == 0 && notZero != 0) {
			t.Errorf("Error NotZero(%v) - got %v", test, notZero)
		}
	}
}

func TestAllCompareInt64(t *testing.T) {
	for _, a := range testsInt64 {
		for _, b := range testsInt64 {
			sameSign := SameSignInt64(a, b)
			equal := EqualInt64(a, b)
			notEqual := NotEqualInt64(a, b)
			greater := GreaterInt64(a, b)
			notGreater := NotGreaterInt64(a, b)
			less := LessInt64(a, b)
			notLess := NotLessInt64(a, b)

			if (((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 0) || (!((a < 0 && b > 0) || (a > 0 && b < 0)) && sameSign != 1) {
				t.Errorf("Error SameSign(%v,%v) - got %v", a, b, sameSign)
			}
			if (a == b && equal != 1) || (a != b && equal != 0) {
				t.Errorf("Error Equal(%v,%v) - got %v", a, b, equal)
			}
			if (a != b && notEqual != 1) || (a == b && notEqual != 0) {
				t.Errorf("Error NotEqual(%v,%v) - got %v", a, b, notEqual)
			}
			if (a > b && greater != 1) || (a <= b && greater != 0) {
				t.Errorf("Error Greater(%v,%v) - got %v", a, b, greater)
			}
			if (a <= b && notGreater != 1) || (a > b && notGreater != 0) {
				t.Errorf("Error NotGreate(%v,%v) - got %v", a, b, notGreater)
			}
			if (a < b && less != 1) || (a >= b && less != 0) {
				t.Errorf("Error Less(%v,%v) - got %v", a, b, less)
			}
			if (a >= b && notLess != 1) || (a < b && notLess != 0) {
				t.Errorf("Error NotLess(%v,%v) - got %v", a, b, notLess)
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
		NegativeInt64(testsInt64[i%testsLen])
	}
}

func BenchmarkNotNegativeInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotNegativeInt64(testsInt64[i%testsLen])
	}
}

func BenchmarkPositiveInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PositiveInt64(testsInt64[i%testsLen])
	}
}

func BenchmarkNotPositiveInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotPositiveInt64(testsInt64[i%testsLen])
	}
}

func BenchmarkZeroInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroInt64(testsInt64[i%testsLen])
	}
}

func BenchmarkNotZeroInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotZeroInt64(testsInt64[i%testsLen])
	}
}

func BenchmarkSignInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SignInt64(testsInt64[i%testsLen])
	}
}

func BenchmarkSameSignInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameSignInt64(testsInt64[i%testsLen], testsInt64[(i+testsLen/4)%testsLen])
	}
}

func BenchmarkEqualInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualInt64(testsInt64[i%testsLen], testsInt64[(i+testsLen/4)%testsLen])
	}
}

func BenchmarkNotEqualInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEqualInt64(testsInt64[i%testsLen], testsInt64[(i+testsLen/4)%testsLen])
	}
}

func BenchmarkGreaterInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreaterInt64(testsInt64[i%testsLen], testsInt64[(i+testsLen/4)%testsLen])
	}
}

func BenchmarkNotGreaterInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotGreaterInt64(testsInt64[i%testsLen], testsInt64[(i+testsLen/4)%testsLen])
	}
}

func BenchmarkLessInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessInt64(testsInt64[i%testsLen], testsInt64[(i+testsLen/4)%testsLen])
	}
}

func BenchmarkNotLessInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotLessInt64(testsInt64[i%testsLen], testsInt64[(i+testsLen/4)%testsLen])
	}
}

package mathh

import (
	"math"
	"math/big"
	"testing"
)

var testsDivideInt64 = []struct {
	a int64
	b int64
	r int64
}{
	{a: -3, b: -7, r: 0},
	{a: -3, b: -6, r: 1},
	{a: -3, b: -5, r: 1},
	{a: -3, b: -4, r: 1},
	{a: -3, b: -3, r: 1},
	{a: -3, b: -2, r: 2},
	{a: -3, b: -1, r: 3},
	//{a: -3, b: 0, r: 1},
	{a: -3, b: 1, r: -3},
	{a: -3, b: 2, r: -2},
	{a: -3, b: 3, r: -1},
	{a: -3, b: 4, r: -1},
	{a: -3, b: 5, r: -1},
	{a: -3, b: 6, r: -1},
	{a: -3, b: 7, r: 0},

	{a: math.MinInt64, b: math.MinInt64, r: 1},
	{a: math.MinInt64, b: math.MinInt64 + 1, r: 1},
	{a: math.MinInt64, b: math.MinInt64 / 2, r: 2},
	{a: math.MinInt64, b: -3, r: 3074457345618258603},
	{a: math.MinInt64, b: -2, r: 4611686018427387904},
	{a: math.MinInt64, b: -1, r: math.MaxInt64}, // Overflow int
	{a: math.MinInt64, b: 1, r: math.MinInt64},
	{a: math.MinInt64, b: 2, r: -4611686018427387904},
	{a: math.MinInt64, b: 3, r: -3074457345618258603},
	{a: math.MinInt64, b: math.MaxInt64 / 2, r: -2},
	{a: math.MinInt64, b: math.MaxInt64 - 1, r: -1},
	{a: math.MinInt64, b: math.MaxInt64, r: -1},

	{a: math.MinInt64 + 1, b: math.MinInt64, r: 1},
	{a: math.MinInt64 + 1, b: math.MinInt64 + 1, r: 1},
	{a: math.MinInt64 + 1, b: math.MinInt64 / 2, r: 2},
	{a: math.MinInt64 + 1, b: -3, r: 3074457345618258602},
	{a: math.MinInt64 + 1, b: -2, r: 4611686018427387904},
	{a: math.MinInt64 + 1, b: -1, r: math.MaxInt64},
	{a: math.MinInt64 + 1, b: 1, r: math.MinInt64 + 1},
	{a: math.MinInt64 + 1, b: 2, r: -4611686018427387904},
	{a: math.MinInt64 + 1, b: 3, r: -3074457345618258602},
	{a: math.MinInt64 + 1, b: math.MaxInt64 / 2, r: -2},
	{a: math.MinInt64 + 1, b: math.MaxInt64 - 1, r: -1},
	{a: math.MinInt64 + 1, b: math.MaxInt64, r: -1},
}

func TestAbsInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if test == math.MinInt64 {
			continue
		}
		if r := AbsInt64(test); r < 0 || ((test < 0 && r != -test) || (test >= 0 && r != test)) {
			t.Errorf("Error Abs(%v) - got %v", test, r)
		}
	}
}

func TestAbsFixInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if r := AbsFixInt64(test); r < 0 || ((test < 0 && r != -test && (test != math.MinInt64 || r != math.MaxInt64)) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsFix(%v) - got %v", test, r)
		}
	}
}

func TestAntiAbsInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if r := AntiAbsInt64(test); r > 0 || ((test > 0 && r != -test) || (test <= 0 && r != test)) {
			t.Errorf("Error AntiAbs(%v) - got %v", test, r)
		}
	}
}

func TestDivideInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if test.a == math.MinInt64 && test.b == -1 {
			continue
		}
		if r := DivideRoundInt64(test.a, test.b); r != test.r {
			t.Errorf("Error Divide(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideFixInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if r := DivideRoundFixInt64(test.a, test.b); r != test.r {
			t.Errorf("Error DivideFix(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func divideInt64Big(a, b int64) (r int64) {
	ba := big.NewInt(a)
	bb := big.NewInt(b)

	ba.Abs(ba)
	bb.Abs(bb)

	bc := big.NewInt(0)
	bd := big.NewInt(0)

	bc.DivMod(ba, bb, bd)

	one := big.NewInt(1)
	two := big.NewInt(2)

	tmp := big.NewInt(0).Div(big.NewInt(0).Add(bb, one), two)
	if bd.Cmp(tmp) != -1 {
		//t.Logf("Fix: %v / %v : %v (%v) ", a, b, bc, bd)
		bc.Add(bc, big.NewInt(1))
	}

	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		bc.Neg(bc)
	}

	// Bound
	if bc.Cmp(big.NewInt(math.MinInt64)) == -1 {
		r = math.MinInt64
	} else if bc.Cmp(big.NewInt(math.MaxInt64)) == 1 {
		r = math.MaxInt64
	} else {
		r = bc.Int64()
	}

	return
}

func TestDivideInt64Overflow(t *testing.T) {
	for _, a := range testsInt64 {
		for _, b := range testsInt64 {
			if b == 0 || (a == math.MinInt64 && b == -1) {
				continue
			}
			validR := divideInt64Big(a, b)
			r := DivideRoundInt64(a, b)
			if r != validR {
				t.Errorf("Error Divide(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func TestDivideFixInt64Overflow(t *testing.T) {
	for _, a := range testsInt64 {
		for _, b := range testsInt64 {
			if b == 0 {
				continue
			}
			validR := divideInt64Big(a, b)
			r := DivideRoundFixInt64(a, b)
			if r != validR {
				t.Errorf("Error DivideFix(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func BenchmarkAbsInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsInt64(testsInt64[i% testsLenInt64])
	}
}

func BenchmarkAbsFixInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsFixInt64(testsInt64[i% testsLenInt64])
	}
}

func BenchmarkAntiAbsInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AntiAbsInt64(testsInt64[i% testsLenInt64])
	}
}

func BenchmarkDivideInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundInt64(testsInt64[i% testsLenInt64], testsInt64[(i+ testsLenInt64 /4)%(testsLenInt64 -1)+1]) // -+1 is to avoid division by zero
	}
}

func BenchmarkDivideFixInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundFixInt64(testsInt64[i% testsLenInt64], testsInt64[(i+ testsLenInt64 /4)%(testsLenInt64 -1)+1]) // -+1 is to avoid division by zero
	}
}

package mathh

import (
	"math/big"
	"testing"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE

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
		bc.Add(bc, big.NewInt(1))
	}

	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		bc.Neg(bc)
	}

	// Bound
	if bc.Cmp(big.NewInt(MinInt64)) == -1 {
		r = MinInt64
	} else if bc.Cmp(big.NewInt(MaxInt64)) == 1 {
		r = MaxInt64
	} else {
		r = bc.Int64()
	}

	return
}

//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32

type testDivideInt64 struct {
	a int64
	b int64
	r int64
}

var testsDivideInt64 = []testDivideInt64{
	{a: 3, b: 7, r: 0},
	{a: 3, b: 6, r: 1},
	{a: 3, b: 5, r: 1},
	{a: 3, b: 4, r: 1},
	{a: 3, b: 3, r: 1},
	{a: 3, b: 2, r: 2},
	{a: 3, b: 1, r: 3},
}

func init() {
	// Extend tests based on the following rules:
	// a/b=c =>
	// 1) -a/b=-c
	// 2) a/-b=-c
	// 3) -a/-b=c
	var ts []testDivideInt64
	for _, v := range testsDivideInt64 {
		ts = append(ts, v)
		if v.a != MinInt64 && v.r != MinInt64 {
			ts = append(ts, testDivideInt64{a: -v.a, b: v.b, r: -v.r})
		}
		if v.b != MinInt64 && v.r != MinInt64 {
			ts = append(ts, testDivideInt64{a: v.a, b: -v.b, r: -v.r})
		}
		if v.a != MinInt64 && v.b != MinInt64 {
			ts = append(ts, testDivideInt64{a: -v.a, b: -v.b, r: v.r})
		}
	}
	testsDivideInt64 = ts
}

func TestAbsInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if test == MinInt64 {
			continue
		}
		if r := AbsInt64(test); r < 0 || ((test < 0 && r != -test) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsInt64(%v) - got %v", test, r)
		}
	}
}

func TestAbsFixInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if r := AbsFixInt64(test); r < 0 || ((test < 0 && r != -test && (test != MinInt64 || r != MaxInt64)) || (test >= 0 && r != test)) {
			t.Errorf("Error AbsFixInt64(%v) - got %v", test, r)
		}
	}
}

func TestAntiAbsInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if r := AntiAbsInt64(test); r > 0 || ((test > 0 && r != -test) || (test <= 0 && r != test)) {
			t.Errorf("Error AntiAbsInt64(%v) - got %v", test, r)
		}
	}
}

func TestDivideInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if test.a == MinInt64 && test.b == -1 {
			continue
		}
		if r := DivideRoundInt64(test.a, test.b); r != test.r {
			t.Errorf("Error DivideInt64(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideFixInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if r := DivideRoundFixInt64(test.a, test.b); r != test.r {
			t.Errorf("Error DivideFixInt64(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideInt64Overflow(t *testing.T) {
	for _, a := range testsInt64 {
		for _, b := range testsInt64 {
			if b == 0 || (a == MinInt64 && b == -1) {
				continue
			}
			validR := divideInt64Big(int64(a), int64(b))
			r := DivideRoundInt64(a, b)
			if r != validR {
				t.Errorf("Error DivideInt64(%v, %v) - got %v, expected %v", a, b, r, validR)
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
			validR := divideInt64Big(int64(a), int64(b))
			r := DivideRoundFixInt64(a, b)
			if r != validR {
				t.Errorf("Error DivideFixInt64(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func BenchmarkAbsInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkAbsFixInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsFixInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkAntiAbsInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AntiAbsInt64(testsInt64[i%testsLenInt64])
	}
}

func BenchmarkDivideInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%(testsLenInt64-1)+1]) // -+1 is to avoid division by zero
	}
}

func BenchmarkDivideFixInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundFixInt64(testsInt64[i%testsLenInt64], testsInt64[(i+testsLenInt64/4)%(testsLenInt64-1)+1]) // -+1 is to avoid division by zero
	}
}

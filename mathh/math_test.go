package mathh

import "testing"

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
import "math/big"

// Hack to avoid unnecessary replace
type customI int64

func divideRoundAsBig(a, b customI) customI { return customI(divideInt64AsBig(int64(a), int64(b))) }

func divideInt64AsBig(a, b int64) (r int64) {
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

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32

type testDivideInt64 struct {
	a     int64
	b     int64
	round int64
	ceil  int64
	floor int64
	rafz  int64
}

var testsDivideInt64 = []testDivideInt64{
	{a: 3, b: 1, round: 3, ceil: 3, floor: 3, rafz: 3},                                     // 3
	{a: 3, b: 2, round: 2, ceil: 2, floor: 1, rafz: 2},                                     // 1.5
	{a: 3, b: 3, round: 1, ceil: 1, floor: 1, rafz: 1},                                     // 1
	{a: 3, b: 4, round: 1, ceil: 1, floor: 0, rafz: 1},                                     // 0.75
	{a: 3, b: 5, round: 1, ceil: 1, floor: 0, rafz: 1},                                     // 0.6
	{a: 3, b: 6, round: 1, ceil: 1, floor: 0, rafz: 1},                                     // 0.5
	{a: 3, b: 7, round: 0, ceil: 1, floor: 0, rafz: 1},                                     // 0.43...
	{a: 0, b: 7, round: 0, ceil: 0, floor: 0, rafz: 0},                                     // 0
	{a: -3, b: 7, round: 0, ceil: 0, floor: -1, rafz: -1},                                  // -0.43...
	{a: -3, b: 6, round: -1, ceil: 0, floor: -1, rafz: -1},                                 // -0.5
	{a: -3, b: 5, round: -1, ceil: 0, floor: -1, rafz: -1},                                 // -0.6
	{a: -3, b: 4, round: -1, ceil: 0, floor: -1, rafz: -1},                                 // -0.75
	{a: -3, b: 3, round: -1, ceil: -1, floor: -1, rafz: -1},                                // -1
	{a: -3, b: 2, round: -2, ceil: -1, floor: -2, rafz: -2},                                // -1.5
	{a: -3, b: 1, round: -3, ceil: -3, floor: -3, rafz: -3},                                // -3
	{a: MinInt64, b: -1, round: MinInt64, ceil: MinInt64, floor: MinInt64, rafz: MinInt64}, // MinInt64 / -1 = MinInt64
}

func init() {
	// Extend tests based on the following rules:
	// a/b=c => -a/-b=c
	var ts []testDivideInt64
	for _, v := range testsDivideInt64 {
		if v.a != MinInt64 && v.b != MinInt64 {
			v.a, v.b = -v.a, -v.b
			ts = append(ts, v)
		}
	}
	testsDivideInt64 = append(testsDivideInt64, ts...)
}

func TestAbsInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if test == MinInt64 {
			continue
		}
		if r := AbsInt64(test); r < 0 || ((test < 0 && r != -test) || (test >= 0 && r != test)) {
			t.Errorf("%v: got %v", test, r)
		}
	}
}

func TestAbsFixInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if r := AbsFixInt64(test); r < 0 || ((test < 0 && r != -test && (test != MinInt64 || r != MaxInt64)) || (test >= 0 && r != test)) {
			t.Errorf("%v: got %v", test, r)
		}
	}
}

func TestAntiAbsInt64(t *testing.T) {
	for _, test := range testsInt64 {
		if r := AntiAbsInt64(test); r > 0 || ((test > 0 && r != -test) || (test <= 0 && r != test)) {
			t.Errorf("%v: got %v", test, r)
		}
	}
}

func TestDivideRoundInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if r := DivideRoundInt64(test.a, test.b); r != test.round {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideRoundFixInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		rightR := test.round
		if test.a == MinInt64 && test.b == -1 {
			rightR = MaxInt64
		}
		if r := DivideRoundFixInt64(test.a, test.b); r != rightR {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, rightR, r)
		}
	}
}

func TestDivideRoundInt64Overflow(t *testing.T) {
	for _, a := range testsInt64 {
		for _, b := range testsInt64 {
			if b == 0 || (a == MinInt64 && b == -1) {
				continue
			}
			validR := int64(divideRoundAsBig(customI(a), customI(b)))
			r := DivideRoundInt64(a, b)
			if r != validR {
				t.Errorf("%v,%v: expect %v, got %v", a, b, validR, r)
			}
		}
	}
}

func TestDivideRoundFixInt64Overflow(t *testing.T) {
	for _, a := range testsInt64 {
		for _, b := range testsInt64 {
			if b == 0 {
				continue
			}
			var validR int64
			if a == MinInt64 && b == -1 {
				validR = MaxInt64
			} else {
				validR = int64(divideRoundAsBig(customI(a), customI(b)))
			}
			r := DivideRoundFixInt64(a, b)
			if r != validR {
				t.Errorf("%v,%v: expect %v,got %v", a, b, validR, r)
			}
		}
	}
}

func TestDivideCeilInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if r := DivideCeilInt64(test.a, test.b); r != test.ceil {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.ceil, r)
		}
	}
}

func TestDivideCeilFixInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		rightR := test.ceil
		if test.a == MinInt64 && test.b == -1 {
			rightR = MaxInt64
		}
		if r := DivideCeilFixInt64(test.a, test.b); r != rightR {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, rightR, r)
		}
	}
}

func TestDivideFloorInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if r := DivideFloorInt64(test.a, test.b); r != test.floor {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.floor, r)
		}
	}
}

func TestDivideFloorFixInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		rightR := test.floor
		if test.a == MinInt64 && test.b == -1 {
			rightR = MaxInt64
		}
		if r := DivideFloorFixInt64(test.a, test.b); r != rightR {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, rightR, r)
		}
	}
}

func TestDivideRafzInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if r := DivideRafzInt64(test.a, test.b); r != test.rafz {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.rafz, r)
		}
	}
}

func TestDivideRafzFixInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		rightR := test.rafz
		if test.a == MinInt64 && test.b == -1 {
			rightR = MaxInt64
		}
		if r := DivideRafzFixInt64(test.a, test.b); r != rightR {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, rightR, r)
		}
	}
}

func TestDivideTruncInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		if r := DivideTruncInt64(test.a, test.b); r != test.a/test.b {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func TestDivideTruncFixInt64(t *testing.T) {
	for _, test := range testsDivideInt64 {
		rightR := test.a / test.b
		if test.a == MinInt64 && test.b == -1 {
			rightR = MaxInt64
		}
		if r := DivideTruncFixInt64(test.a, test.b); r != rightR {
			t.Errorf("%v,%v: expect %v, got %v", test.a, test.b, rightR, r)
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

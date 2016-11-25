package mathh

import (
	"math/big"
	"testing"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE

func divideUint64Big(a, b uint64) (r uint64) {
	ba := big.NewInt(0).SetUint64(a)
	bb := big.NewInt(0).SetUint64(b)

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
	if bc.Cmp(big.NewInt(MinUint64)) == -1 {
		r = MinUint64
	} else if bc.Cmp(big.NewInt(0).SetUint64(MaxUint64)) == 1 {
		r = MaxUint64
	} else {
		r = bc.Uint64()
	}

	return
}

//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32

type testDivideUint64 struct {
	a uint64
	b uint64
	r uint64
}

var testsDivideUint64 = []testDivideUint64{
	{a: 3, b: 7, r: 0},
	{a: 3, b: 6, r: 1},
	{a: 3, b: 5, r: 1},
	{a: 3, b: 4, r: 1},
	{a: 3, b: 3, r: 1},
	{a: 3, b: 2, r: 2},
	{a: 3, b: 1, r: 3},
}

func TestDivideUint64(t *testing.T) {
	for _, test := range testsDivideUint64 {
		if r := DivideRoundUint64(test.a, test.b); r != test.r {
			t.Errorf("Error DivideUint64(%v, %v) - expected %v, got %v", test.a, test.b, test.r, r)
		}
	}
}

func TestDivideUint64Overflow(t *testing.T) {
	for _, a := range testsUint64 {
		for _, b := range testsUint64 {
			if b == 0 || (a == MinUint64 && b == -1) {
				continue
			}
			validR := divideUint64Big(uint64(a), uint64(b))
			r := DivideRoundUint64(a, b)
			if r != validR {
				t.Errorf("Error Divide(%v, %v) - got %v, expected %v", a, b, r, validR)
			}
		}
	}
}

func BenchmarkDivideUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%(testsLenUint64-1)+1]) // -+1 is to avoid division by zero
	}
}

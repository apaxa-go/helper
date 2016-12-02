package mathh

import "testing"

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
import "math/big"

// Hack to avoid unnecessary replace
type customU uint64

func divideAsUBig(a, b customU) (r customU) { return customU(divideUint64AsBig(uint64(a), uint64(b))) }

func divideUint64AsBig(a, b uint64) (r uint64) {
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

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32

type testDivideUint64 struct {
	a     uint64
	b     uint64
	round uint64
	up    uint64
	down  uint64
}

var testsDivideUint64 = []testDivideUint64{
	{a: 3, b: 1, round: 3, up: 3, down: 3}, // 3
	{a: 3, b: 2, round: 2, up: 2, down: 1}, // 1.5
	{a: 3, b: 3, round: 1, up: 1, down: 1}, // 1
	{a: 3, b: 4, round: 1, up: 1, down: 0}, // 0.75
	{a: 3, b: 5, round: 1, up: 1, down: 0}, // 0.6
	{a: 3, b: 6, round: 1, up: 1, down: 0}, // 0.5
	{a: 3, b: 7, round: 0, up: 1, down: 0}, // 0.43...
	{a: 0, b: 7, round: 0, up: 0, down: 0}, // 0
}

func TestDivideRoundUint64(t *testing.T) {
	for _, test := range testsDivideUint64 {
		if r := DivideRoundUint64(test.a, test.b); r != test.round {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.round, r)
		}
	}
}

func TestDivideUint64Overflow(t *testing.T) {
	for _, a := range testsUint64 {
		for _, b := range testsUint64 {
			if b == 0 {
				continue
			}
			validR := uint64(divideAsUBig(customU(a), customU(b)))
			r := DivideRoundUint64(a, b)
			if r != validR {
				t.Errorf("%v, %v: expect %v, got %v, ", a, b, validR, r)
			}
		}
	}
}

func TestDivideCeilUint64(t *testing.T) {
	for _, test := range testsDivideUint64 {
		if r := DivideCeilUint64(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideFloorUint64(t *testing.T) {
	for _, test := range testsDivideUint64 {
		if r := DivideFloorUint64(test.a, test.b); r != test.down {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.down, r)
		}
	}
}

func TestDivideRafzUint64(t *testing.T) {
	for _, test := range testsDivideUint64 {
		if r := DivideRafzUint64(test.a, test.b); r != test.up {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.up, r)
		}
	}
}

func TestDivideTruncUint64(t *testing.T) {
	for _, test := range testsDivideUint64 {
		if r := DivideTruncUint64(test.a, test.b); r != test.a/test.b {
			t.Errorf("%v, %v: expect %v, got %v", test.a, test.b, test.a/test.b, r)
		}
	}
}

func BenchmarkDivideUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideRoundUint64(testsUint64[i%testsLenUint64], testsUint64[(i+testsLenUint64/4)%(testsLenUint64-1)+1]) // -+1 is to avoid division by zero
	}
}

package tokenh

import (
	"go/token"
	"testing"
)

func TestIs(t *testing.T) {
	type testElement struct {
		t       token.Token
		e, o, s bool // Equality, Order & Shift
	}

	const maxToken = 200

	tests := []testElement{
		{token.EQL, true, false, false},
		{token.NEQ, true, false, false},
		{token.LSS, false, true, false},
		{token.LEQ, false, true, false},
		{token.GTR, false, true, false},
		{token.GEQ, false, true, false},
		{token.SHL, false, false, true},
		{token.SHR, false, false, true},
	}

	checked := make(map[token.Token]struct{})

	for _, test := range tests {
		checked[test.t] = struct{}{}

		c := IsComparison(test.t)
		e := IsEqualityCheck(test.t)
		o := IsOrderCheck(test.t)
		s := IsShift(test.t)

		if c != (test.e || test.o) {
			t.Errorf("%v: expect %v, got %v", test.t, test.e || test.o, c)
		}
		if e != test.e {
			t.Errorf("%v: expect %v, got %v", test.t, test.e, e)
		}
		if o != test.o {
			t.Errorf("%v: expect %v, got %v", test.t, test.o, o)
		}
		if s != test.s {
			t.Errorf("%v: expect %v, got %v", test.t, test.s, s)
		}
	}

	// Check all other token [0; maxToken) except already checked
	for tok := token.Token(0); tok < maxToken; tok++ {
		if _, skip := checked[tok]; skip {
			continue
		}
		if IsComparison(tok) {
			t.Errorf("%v: expect false, got true", tok)
		}
		if IsEqualityCheck(tok) {
			t.Errorf("%v: expect false, got true", tok)
		}
		if IsOrderCheck(tok) {
			t.Errorf("%v: expect false, got true", tok)
		}
		if IsShift(tok) {
			t.Errorf("%v: expect false, got true", tok)
		}
	}
}

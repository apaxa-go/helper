package asth

import (
	"testing"
)

func TestIsValidIdent(t *testing.T) {
	type testElement struct {
		i                 string
		blank             bool
		private, exported bool
	}

	tests := []testElement{
		{"", false, false, false},
		{" ", false, false, false},
		{" a", false, false, false},
		{" A", false, false, false},
		{"_", true, false, false},
		{"1", false, false, false},
		{"a", false, true, false},
		{"_a", false, true, false},
		{"A", false, false, true},
		{"aAa1", false, true, false},
		{"AaA1", false, false, true},
		{"A.B", false, false, false},
	}

	for _, test := range tests {
		b := IsBlankIdent(test.i)
		v := IsValidIdent(test.i)
		p := IsValidNotExportedIdent(test.i)
		e := IsValidExportedIdent(test.i)

		if b != test.blank {
			t.Errorf("%v: expect blank %v, got %v", test.i, test.blank, b)
		}
		if v != (test.private || test.exported) {
			t.Errorf("%v: expect valid %v, got %v", test.i, test.private || test.exported, v)
		}
		if p != test.private {
			t.Errorf("%v: expect valid not exported %v, got %v", test.i, test.private, p)
		}
		if e != test.exported {
			t.Errorf("%v: expect valid exported %v, got %v", test.i, test.exported, e)
		}
	}
}

func TestValidateIdent(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = validateIdent("a", 10)
	t.Error("expect panic")
}

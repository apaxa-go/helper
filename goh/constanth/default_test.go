package constanth

import (
	"go/constant"
	"go/token"
	"reflect"
	"testing"
)

func TestDefaultValue(t *testing.T) {
	type testElement struct {
		c  constant.Value
		r  interface{}
		ok bool
	}

	tests := []testElement{
		{MakeBool(true), true, true},
		{MakeString("str"), "str", true},
		{MakeInt(123), 123, true},
		{constant.MakeFromLiteral("123456789012345678901234567890", token.INT, 0), nil, false},
		{MakeFloat64(123.456), 123.456, true},
		{constant.MakeFromLiteral("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890.123", token.FLOAT, 0), nil, false},
		{MakeComplex128(1.2 - 3.4i), 1.2 - 3.4i, true},
		{constant.MakeFromLiteral("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890.123i", token.IMAG, 0), nil, false},
		{constant.MakeUnknown(), nil, false},
	}

	for _, test := range tests {
		if r, ok := DefaultValueInterface(test.c); ok != test.ok || (test.ok && r != test.r) {
			t.Errorf("%v: expect %v %v, got %v %v", test.c.ExactString(), test.r, test.ok, r, ok)
		}
		if r, ok := DefaultValue(test.c); ok != test.ok || (test.ok && r.Interface() != test.r) {
			t.Errorf("%v: expect %v %v, got %v %v", test.c.ExactString(), test.r, test.ok, r, ok)
		}
		if !test.ok {
			continue
		}
		if r := DefaultType(test.c); r != reflect.TypeOf(test.r) {
			t.Errorf("%v: expect %v, got %v", test.c.ExactString(), reflect.TypeOf(test.r), r)
		}
	}
}

func TestDefaultTypeForKind(t *testing.T) {
	if DefaultTypeForKind(constant.Unknown) != nil {
		t.Error("expect nil")
	}
}

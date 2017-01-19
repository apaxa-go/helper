package constanth

import (
	"go/constant"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	for _, k := range []constant.Kind{constant.Unknown, constant.Bool, constant.String, constant.Complex + 1} {
		if IsNumeric(k) {
			t.Errorf("%v: expect false", KindString(k))
		}
	}
	for _, k := range []constant.Kind{constant.Int, constant.Float, constant.Complex} {
		if !IsNumeric(k) {
			t.Errorf("%v: expect false", KindString(k))
		}
	}
}

func TestKindString(t *testing.T) {
	if KindString(constant.Unknown) != "Unknown" || KindString(constant.Bool) != "Bool" || KindString(constant.String) != "String" || KindString(constant.Int) != "Int" || KindString(constant.Float) != "Float" || KindString(constant.Complex) != "Complex" || KindString(constant.Complex+1) != "Invalid" {
		t.Error("invalid kind string")
	}
}

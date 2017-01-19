package reflecth

import (
	"go/token"
	"reflect"
	"testing"
)

func TestCompareOp(t *testing.T) {
	type testElement struct {
		x   interface{}
		op  token.Token
		y   interface{}
		r   bool
		err bool
	}

	tests := []testElement{
		{true, token.EQL, true, true, false},
		{1, token.LSS, 2, true, false},
		{2, token.LEQ, 1, false, false},
		{1, token.GTR, 2, false, false},
		{2, token.GEQ, 1, true, false},
		{uint(1), token.LSS, uint(2), true, false},
		{uint(2), token.LEQ, uint(1), false, false},
		{uint(1), token.GTR, uint(2), false, false},
		{uint(2), token.GEQ, uint(1), true, false},
		{float32(1.2), token.LSS, float32(3.4), true, false},
		{float32(3.4), token.LEQ, float32(1.2), false, false},
		{float32(1.2), token.GTR, float32(3.4), false, false},
		{float32(3.4), token.GEQ, float32(1.2), true, false},
		{"1", token.LSS, "2", true, false},
		{"2", token.LEQ, "1", false, false},
		{"1", token.GTR, "2", false, false},
		{"2", token.GEQ, "1", true, false},
		// Negative
		{"1", token.EQL, 1, false, true},
		{[]int{1}, token.EQL, []int{2}, false, true},
		{complex64(1), token.LSS, complex64(2), false, true},
		{1, token.ADD, 2, false, true},
	}

	for _, test := range tests {
		r, err := CompareOp(reflect.ValueOf(test.x), test.op, reflect.ValueOf(test.y))
		if err != nil != test.err || r != test.r {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", test.x, test.op, test.y, test.r, test.err, r, err)
		}
	}
}

func TestCompareOpInt(t *testing.T) {
	r, err := compareOpInt(1, token.ADD, 2)
	if r || err == nil {
		t.Error("expect error")
	}
}

func TestCompareOpUint(t *testing.T) {
	r, err := compareOpUint(1, token.ADD, 2)
	if r || err == nil {
		t.Error("expect error")
	}
}

func TestCompareOpFloat(t *testing.T) {
	r, err := compareOpFloat(1, token.ADD, 2)
	if r || err == nil {
		t.Error("expect error")
	}
}

func TestCompareOpString(t *testing.T) {
	r, err := compareOpString("1", token.ADD, "2")
	if r || err == nil {
		t.Error("expect error")
	}
}

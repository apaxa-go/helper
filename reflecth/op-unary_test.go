package reflecth

import (
	"github.com/apaxa-go/helper/mathh"
	"go/token"
	"reflect"
	"testing"
)

func TestUnaryOp(t *testing.T) {
	type testElement struct {
		op  token.Token
		y   interface{}
		r   interface{}
		err bool
	}

	tests := []testElement{
		{token.SUB, 1, -1, false},
		{token.SUB, mathh.MinInt, mathh.MinInt, false}, // specific behaviour for MinInt
		{token.SUB, float32(1), float32(-1), false},
		{token.SUB, complex64(1 - 2i), complex64(-1 + 2i), false},
		{token.SUB, "1", nil, true},
		{token.XOR, int(-85), ^int(-85), false},
		{token.XOR, int8(-85), ^int8(-85), false},
		{token.XOR, int16(-85), ^int16(-85), false},
		{token.XOR, int32(-85), ^int32(-85), false},
		{token.XOR, int64(85), ^int64(85), false},
		{token.XOR, uint(85), ^uint(85), false},
		{token.XOR, uint8(85), ^uint8(85), false},
		{token.XOR, uint16(85), ^uint16(85), false},
		{token.XOR, uint32(85), ^uint32(85), false},
		{token.XOR, uint64(85), ^uint64(85), false},
		{token.XOR, float32(-85), nil, true},
		{token.NOT, true, false, false},
		{token.NOT, "1", nil, true},
		{token.ADD, int8(-85), int8(-85), false},
		{token.ADD, float32(85), float32(85), false},
		{token.ADD, complex128(1.2 - 3.4i), complex128(1.2 - 3.4i), false},
		{token.ADD, "-85", nil, true},
		// Negative
		{token.EQL, 1, nil, true},
	}

	for _, test := range tests {
		r, err := UnaryOp(test.op, reflect.ValueOf(test.y))
		if err != nil != test.err || (!test.err && (r.Interface() != test.r)) {
			t.Errorf("%v %v: expect %v %v, got %v %v", test.op, test.y, test.r, test.err, r, err)
		}
	}
}

func TestUnaryOpAnd(t *testing.T) {
	var a = 43

	// addressable argument
	r, err := UnaryOp(token.AND, reflect.ValueOf(&a).Elem())
	if err != nil || r.Interface() != &a {
		t.Error("error")
	}

	// not addressable argument
	r, err = UnaryOp(token.AND, reflect.ValueOf(a))
	if err != nil || *(r.Interface().(*int)) != a {
		t.Error("error")
	}
}

func TestUnaryOpArrow(t *testing.T) {
	a := make(chan int, 10)
	a <- 4
	a <- 5
	a <- 6
	a <- 7

	r, err := UnaryOp(token.ARROW, reflect.ValueOf(a))
	if err != nil || r.Interface() != 4 {
		t.Error("error")
	}

	r, err = UnaryOp(token.ARROW, reflect.ValueOf((<-chan int)(a)))
	if err != nil || r.Interface() != 5 {
		t.Error("error")
	}

	r, err = UnaryOp(token.ARROW, reflect.ValueOf((chan<- int)(a)))
	if err == nil {
		t.Error("error expected")
	}

	r, err = UnaryOp(token.ARROW, reflect.ValueOf(1))
	if err == nil {
		t.Error("error expected")
	}
}

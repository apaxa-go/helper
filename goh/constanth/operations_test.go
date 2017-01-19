package constanth

import (
	"github.com/apaxa-go/helper/reflecth"
	"go/constant"
	"go/token"
	"reflect"
	"testing"
)

func TestCompareOp(t *testing.T) {
	type testElement struct {
		x   constant.Value
		op  token.Token
		y   constant.Value
		r   bool
		err bool
	}

	tests := []testElement{
		{MakeInt(10), token.EQL, MakeInt(11), false, false},
		{MakeString("10"), token.EQL, MakeString("10"), true, false},
		{MakeInt(10), token.EQL, MakeString("10"), false, true},
	}

	for _, test := range tests {
		if r, err := CompareOp(test.x, test.op, test.y); err != nil != test.err || r != test.r {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", test.x, test.op, test.y, test.r, test.err, r, err)
		}
	}
}

func TestBinaryOp(t *testing.T) {
	type testElement struct {
		x   constant.Value
		op  token.Token
		y   constant.Value
		r   constant.Value
		err bool
	}

	tests := []testElement{
		{MakeInt(10), token.SUB, MakeInt(11), MakeInt(-1), false},
		{MakeString("10"), token.ADD, MakeString("11"), MakeString("1011"), false},
		{MakeInt(10), token.ADD, MakeString("10"), nil, true},
		{constant.MakeUnknown(), token.SUB, MakeInt(11), nil, true},
	}

	for _, test := range tests {
		if r, err := BinaryOp(test.x, test.op, test.y); err != nil != test.err || r != test.r {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", test.x, test.op, test.y, test.r, test.err, r, err)
		}
	}
}

func TestUnaryOp(t *testing.T) {
	type testElement struct {
		op  token.Token
		y   constant.Value
		r   constant.Value
		err bool
	}

	tests := []testElement{
		{token.SUB, MakeInt(11), MakeInt(-11), false},
		{token.ADD, MakeInt(11), MakeInt(11), false},
		{token.SUB, MakeString("10"), nil, true},
		{token.SUB, constant.MakeUnknown(), nil, true},
	}

	for _, test := range tests {
		if r, err := UnaryOp(test.op, test.y, 0); err != nil != test.err || r != test.r {
			t.Errorf("%v %v: expect %v %v, got %v %v", test.op, test.y, test.r, test.err, r, err)
		}
	}
}

func TestShiftOp(t *testing.T) {
	type testElement struct {
		x   constant.Value
		op  token.Token
		s   uint
		r   constant.Value
		err bool
	}

	tests := []testElement{
		{MakeInt(10), token.SHL, 1, MakeInt(20), false},
		{MakeUint(8), token.SHR, 2, MakeUint(2), false},
		{MakeInt(10), token.ADD, 1, nil, true},
		{constant.MakeUnknown(), token.SHR, 1, nil, true},
	}

	for _, test := range tests {
		if r, err := ShiftOp(test.x, test.op, test.s); err != nil != test.err || r != test.r {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", test.x, test.op, test.s, test.r, test.err, r, err)
		}
	}
}

func TestCompareOpTyped(t *testing.T) {
	type testElement struct {
		x   constant.Value
		xT  reflect.Type
		op  token.Token
		y   constant.Value
		yT  reflect.Type
		r   bool
		err bool
	}

	tests := []testElement{
		{MakeInt(10), reflecth.TypeInt(), token.EQL, MakeInt(11), reflecth.TypeInt(), false, false},
		{MakeString("10"), reflecth.TypeString(), token.EQL, MakeString("10"), reflecth.TypeString(), true, false},
		{MakeInt(10), reflecth.TypeInt16(), token.EQL, MakeString("10"), reflecth.TypeString(), false, true},
		{MakeInt(10), reflecth.TypeInt(), token.ADD, MakeInt(11), reflecth.TypeInt(), false, true},
	}

	for _, test := range tests {
		x := MustMakeTypedValue(test.x, test.xT)
		y := MustMakeTypedValue(test.y, test.yT)

		if r, err := CompareOpTyped(x, test.op, y); err != nil != test.err || r != test.r {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", x, test.op, y, test.r, test.err, r, err)
		}
	}
}

func TestBinaryOpTyped(t *testing.T) {
	type testElement struct {
		x   constant.Value
		xT  reflect.Type
		op  token.Token
		y   constant.Value
		yT  reflect.Type
		r   constant.Value
		rT  reflect.Type
		err bool
	}

	tests := []testElement{
		{MakeInt(10), reflecth.TypeInt16(), token.SUB, MakeInt(11), reflecth.TypeInt16(), MakeInt(-1), reflecth.TypeInt16(), false},
		{MakeString("10"), reflecth.TypeString(), token.ADD, MakeString("11"), reflecth.TypeString(), MakeString("1011"), reflecth.TypeString(), false},
		{MakeInt(10), reflecth.TypeInt16(), token.EQL, MakeInt(11), reflecth.TypeInt16(), nil, nil, true},
		{MakeInt(10), reflecth.TypeInt16(), token.ADD, MakeString("10"), reflecth.TypeString(), nil, nil, true},
		{MakeInt(127), reflecth.TypeInt8(), token.ADD, MakeInt(127), reflecth.TypeInt8(), nil, nil, true},
	}

	for _, test := range tests {
		x := TypedValue{test.x, test.xT}
		y := TypedValue{test.y, test.yT}
		testR := TypedValue{test.r, test.rT}

		if r, err := BinaryOpTyped(x, test.op, y); err != nil != test.err || r != testR {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", x, test.op, y, testR, test.err, r, err)
		}
	}
}

func TestUnaryOpTyped(t *testing.T) {
	type testElement struct {
		op  token.Token
		y   constant.Value
		yT  reflect.Type
		r   constant.Value
		rT  reflect.Type
		err bool
	}

	tests := []testElement{
		{token.SUB, MakeInt(11), reflecth.TypeInt32(), MakeInt(-11), reflecth.TypeInt32(), false},
		{token.ADD, MakeInt(11), reflecth.TypeUint64(), MakeInt(11), reflecth.TypeUint64(), false},
		{token.SUB, MakeString("10"), reflecth.TypeString(), nil, nil, true},
		{token.SUB, constant.MakeUnknown(), reflecth.TypeInt(), nil, nil, true},
		{token.SUB, MakeUint32(11), reflecth.TypeUint32(), nil, nil, true},
	}

	for _, test := range tests {
		y := TypedValue{test.y, test.yT}
		testR := TypedValue{test.r, test.rT}

		if r, err := UnaryOpTyped(test.op, y, 0); err != nil != test.err || r != testR {
			t.Errorf("%v %v: expect %v %v, got %v %v", test.op, y, testR, test.err, r, err)
		}
	}
}

func TestShiftOpTyped(t *testing.T) {
	type testElement struct {
		x   constant.Value
		xT  reflect.Type
		op  token.Token
		s   uint
		r   constant.Value
		rT  reflect.Type
		err bool
	}

	tests := []testElement{
		{MakeInt(10), reflecth.TypeInt(), token.SHL, 1, MakeInt(20), reflecth.TypeInt(), false},
		{MakeUint(8), reflecth.TypeUint8(), token.SHR, 2, MakeUint(2), reflecth.TypeUint8(), false},
		{MakeInt(10), reflecth.TypeInt8(), token.ADD, 1, nil, nil, true},
		{constant.MakeUnknown(), reflecth.TypeUint16(), token.SHR, 1, nil, nil, true},
		{MakeUint(8), reflecth.TypeUint8(), token.SHL, 8, nil, nil, true},
	}

	for _, test := range tests {
		x := TypedValue{test.x, test.xT}
		testR := TypedValue{test.r, test.rT}

		if r, err := ShiftOpTyped(x, test.op, test.s); err != nil != test.err || r != testR {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", x, test.op, test.s, testR, test.err, r, err)
		}
	}
}

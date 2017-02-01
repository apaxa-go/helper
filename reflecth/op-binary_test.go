package reflecth

import "go/token"

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE

import (
	"reflect"
	"testing"
)

type testBinaryOpElement struct {
	x   interface{}
	op  token.Token
	y   interface{}
	r   interface{}
	err bool
}

//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32
//replacer:new uint64	Uint64

var testsBinaryOpInt64 = []testBinaryOpElement{
	{int64(2), token.ADD, int64(1), int64(3), false},
	{int64(2), token.SUB, int64(1), int64(1), false},
	{int64(2), token.MUL, int64(3), int64(6), false},
	{int64(11), token.QUO, int64(3), int64(3), false},
	{int64(8), token.REM, int64(3), int64(2), false},
	{int64(6), token.AND, int64(3), int64(2), false},
	{int64(6), token.OR, int64(3), int64(7), false},
	{int64(6), token.XOR, int64(3), int64(5), false},
	{int64(6), token.AND_NOT, int64(3), int64(6) &^ int64(3), false},
	{int64(6), token.EQL, int64(3), nil, true},
}

//replacer:replace
//replacer:old float32		Float32
//replacer:new float64		Float64

var testsBinaryOpFloat32 = []testBinaryOpElement{
	{float32(3.4), token.ADD, float32(1.3), float32(4.7), false},
	{float32(3.4), token.SUB, float32(1.2), float32(2.2), false},
	{float32(1.2), token.MUL, float32(0.3), float32(0.36), false},
	{float32(3), token.QUO, float32(1.5), float32(2), false},
	{float32(3.4), token.EQL, float32(17), nil, true},
}

//replacer:replace
//replacer:old complex64	Complex64
//replacer:new complex128	Complex128

var testsBinaryOpComplex64 = []testBinaryOpElement{
	{complex64(3.4 + 2i), token.ADD, complex64(1.3 - 1i), complex64(4.7 + 1i), false},
	{complex64(3.4 + 2i), token.SUB, complex64(1.2 + 1i), complex64(2.2 + 1i), false},
	{complex64(1 + 3i), token.MUL, complex64(5 + 6i), complex64(-13 + 21i), false},
	{complex64(3i), token.QUO, complex64(1.5), complex64(2i), false},
	{complex64(3.4 + 2.1i), token.EQL, complex64(17 - 0.1i), nil, true},
}

//replacer:ignore

func TestBinaryOp(t *testing.T) {
	tests := []testBinaryOpElement{
		// Bool
		{true, token.LAND, false, false, false},
		{true, token.LOR, false, true, false},
		{true, token.ADD, false, nil, true},
		// String
		{"s1", token.ADD, "s2", "s1s2", false},
		{"s1", token.SUB, "s2", nil, true},
		// Negative
		{1, token.ADD, "s1", nil, true},
		{struct{ a int }{1}, token.SUB, struct{ a int }{2}, nil, true},
	}

	tests = append(tests, testsBinaryOpInt...)
	tests = append(tests, testsBinaryOpInt8...)
	tests = append(tests, testsBinaryOpInt16...)
	tests = append(tests, testsBinaryOpInt32...)
	tests = append(tests, testsBinaryOpInt64...)
	tests = append(tests, testsBinaryOpUint...)
	tests = append(tests, testsBinaryOpUint8...)
	tests = append(tests, testsBinaryOpUint16...)
	tests = append(tests, testsBinaryOpUint32...)
	tests = append(tests, testsBinaryOpUint64...)
	tests = append(tests, testsBinaryOpFloat32...)
	tests = append(tests, testsBinaryOpFloat64...)
	tests = append(tests, testsBinaryOpComplex64...)
	tests = append(tests, testsBinaryOpComplex128...)

	for _, test := range tests {
		r, err := BinaryOp(reflect.ValueOf(test.x), test.op, reflect.ValueOf(test.y))
		if err != nil != test.err || (!test.err && (r.Interface() != test.r)) {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", test.x, test.op, test.y, test.r, test.err, r, err)
		}
	}
}

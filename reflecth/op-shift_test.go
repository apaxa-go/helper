package reflecth

import (
	"go/token"
	"reflect"
	"testing"
)

func TestShiftOp(t *testing.T) {
	type testElement struct {
		x   interface{}
		op  token.Token
		s   uint
		r   interface{}
		err bool
	}

	tests := []testElement{
		{int(8), token.SHL, 1, int(16), false},
		{int8(8), token.SHL, 1, int8(16), false},
		{int16(8), token.SHL, 2, int16(32), false},
		{int32(8), token.SHL, 2, int32(32), false},
		{int64(8), token.SHL, 3, int64(64), false},
		{uint(8), token.SHL, 1, uint(16), false},
		{uint8(8), token.SHL, 1, uint8(16), false},
		{uint16(8), token.SHL, 2, uint16(32), false},
		{uint32(8), token.SHL, 2, uint32(32), false},
		{uint64(8), token.SHL, 3, uint64(64), false},
		{int(8), token.SHR, 1, int(4), false},
		{int8(8), token.SHR, 1, int8(4), false},
		{int16(8), token.SHR, 2, int16(2), false},
		{int32(8), token.SHR, 2, int32(2), false},
		{int64(8), token.SHR, 3, int64(1), false},
		{uint(8), token.SHR, 1, uint(4), false},
		{uint8(8), token.SHR, 1, uint8(4), false},
		{uint16(8), token.SHR, 2, uint16(2), false},
		{uint32(8), token.SHR, 2, uint32(2), false},
		{uint64(8), token.SHR, 3, uint64(1), false},
		// negative
		{float32(8), token.SHL, 1, nil, true},
		{float32(8), token.SHR, 1, nil, true},
		{int(8), token.EQL, 1, nil, true},
	}

	for _, test := range tests {
		r, err := ShiftOp(reflect.ValueOf(test.x), test.op, test.s)
		if err != nil != test.err || (!test.err && (r.Interface() != test.r)) {
			t.Errorf("%v %v %v: expect %v %v, got %v %v", test.x, test.op, test.s, test.r, test.err, r, err)
		}
	}
}

package constanth

import (
	"github.com/apaxa-go/helper/reflecth"
	"go/constant"
	"go/token"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type (
		myBool       bool
		myInt        int
		myInt8       int8
		myInt16      int16
		myInt32      int32
		myInt64      int64
		myUint       uint
		myUint8      uint8
		myUint16     uint16
		myUint32     uint32
		myUint64     uint64
		myFloat32    float32
		myFloat64    float64
		myComplex64  complex64
		myComplex128 complex128
		myString     string
	)

	type testElement struct {
		c  constant.Value
		t  reflect.Type
		r  constant.Value // result typed const must be of type t, so it is not required to store it for check; set this field to nil if in result it must be the same as c
		ok bool
	}

	tests := []testElement{
		// fully direct, no conversion
		{MakeBool(true), reflect.TypeOf(myBool(false)), nil, true},

		{MakeInt(123), reflect.TypeOf(myInt(0)), nil, true},
		{MakeInt8(-123), reflect.TypeOf(myInt8(0)), nil, true},
		{MakeInt16(123), reflect.TypeOf(myInt16(0)), nil, true},
		{MakeInt32(123), reflect.TypeOf(myInt32(0)), nil, true},
		{MakeInt64(123), reflect.TypeOf(myInt64(0)), nil, true},
		{MakeUint(123), reflect.TypeOf(myUint(0)), nil, true},
		{MakeUint8(123), reflect.TypeOf(myUint8(0)), nil, true},
		{MakeUint16(123), reflect.TypeOf(myUint16(0)), nil, true},
		{MakeUint32(123), reflect.TypeOf(myUint32(0)), nil, true},
		{MakeUint64(123), reflect.TypeOf(myUint64(0)), nil, true},

		{MakeFloat32(123.456), reflect.TypeOf(myFloat32(0)), nil, true},
		{MakeFloat64(123.456), reflect.TypeOf(myFloat64(0)), nil, true},

		{MakeComplex64(1.2 + 3.4i), reflect.TypeOf(myComplex64(0)), nil, true},
		{MakeComplex128(1.2 - 3.4i), reflect.TypeOf(myComplex128(0)), nil, true},

		{MakeString("str"), reflect.TypeOf(myString("")), nil, true},

		// no conversion
		{MakeInt64(123), reflect.TypeOf(myFloat32(0)), MakeFloat32(123), true},
		{MakeInt64(123), reflect.TypeOf(myComplex128(0)), MakeComplex128(123), true},

		{MakeFloat64(123.0), reflect.TypeOf(myInt16(0)), MakeInt16(123), true},
		{MakeFloat32(123.456), reflect.TypeOf(myComplex64(0)), MakeComplex64(123.456), true},

		{MakeComplex128(1.0 - 0i), reflect.TypeOf(myUint8(0)), MakeUint8(1), true},
		{MakeComplex128(1.2 + 0i), reflect.TypeOf(myFloat64(0)), MakeFloat64(1.2), true},

		// conversion
		{MakeInt8(65), reflect.TypeOf(myString("")), MakeString("A"), true},
		{MakeUint64(1234567890123), reflect.TypeOf(myString("")), MakeString("\uFFFD"), true},

		// negative
		{MakeBool(true), reflecth.TypeInt(), nil, false},
		{MakeBool(true), reflecth.TypeUint8(), nil, false},
		{MakeBool(true), reflecth.TypeFloat64(), nil, false},
		{MakeBool(true), reflecth.TypeComplex64(), nil, false},
		{MakeBool(true), reflecth.TypeString(), nil, false},

		{MakeInt(123), reflecth.TypeBool(), nil, false},

		{MakeFloat32(123.456), reflecth.TypeBool(), nil, false},
		{MakeFloat32(123.456), reflecth.TypeInt(), nil, false},
		{MakeFloat32(123.456), reflecth.TypeUint(), nil, false},
		{MakeFloat32(123.456), reflecth.TypeString(), nil, false},

		{MakeComplex128(1.2 - 3.4i), reflecth.TypeBool(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeInt(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeUint(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeFloat64(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeString(), nil, false},

		{MakeString("str"), reflecth.TypeBool(), nil, false},
		{MakeString("str"), reflecth.TypeInt(), nil, false},
		{MakeString("str"), reflecth.TypeUint(), nil, false},
		{MakeString("str"), reflecth.TypeFloat64(), nil, false},
		{MakeString("str"), reflecth.TypeComplex64(), nil, false},

		{MakeInt(123), reflect.TypeOf(struct{ x int }{0}), nil, false},

		{constant.MakeUnknown(), reflect.TypeOf(myBool(false)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myInt(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myInt8(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myInt16(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myInt32(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myInt64(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myUint(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myUint8(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myUint16(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myUint32(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myUint64(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myFloat32(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myFloat64(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myComplex64(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myComplex128(0)), nil, false},
		{constant.MakeUnknown(), reflect.TypeOf(myString("")), nil, false},

		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myInt(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myInt8(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myInt16(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myInt32(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myInt64(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myUint(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myUint8(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myUint16(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myUint32(0)), nil, false},
		{constant.MakeFromLiteral(veryLongNumber, token.INT, 0), reflect.TypeOf(myUint64(0)), nil, false},
	}

	for _, test := range tests {
		if test.r == nil {
			test.r = test.c
		}

		r, ok := Convert(test.c, test.t)
		testR := TypedValue{t: test.t, v: test.r}

		if ok != test.ok || (ok && !r.Equal(testR)) {
			t.Errorf("%v %v: expect %v %v, got %v %v", test.c, test.t, testR, test.ok, r, ok)
		}
		if r := IsConvertible(test.c, test.t); r != test.ok {
			t.Errorf("%v %v: expect %v, got %v", test.c, test.t, test.ok, r)
		}
		if !test.ok {
			continue
		}
		if r := MustConvert(test.c, test.t); !r.Equal(testR) {
			t.Errorf("%v %v: expect %v, got %v", test.c, test.t, testR, r)
		}
	}
}

func TestMustConvert(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = MustConvert(MakeBool(true), reflecth.TypeInt())
	t.Error("panic expected")
}

func TestAssign(t *testing.T) {
	type (
		myBool       bool
		myInt        int
		myInt8       int8
		myInt16      int16
		myInt32      int32
		myInt64      int64
		myUint       uint
		myUint8      uint8
		myUint16     uint16
		myUint32     uint32
		myUint64     uint64
		myFloat32    float32
		myFloat64    float64
		myComplex64  complex64
		myComplex128 complex128
		myString     string
	)

	type testElement struct {
		c  constant.Value
		t  reflect.Type
		r  interface{}
		ok bool
	}

	tests := []testElement{
		// fully direct, no conversion
		{MakeBool(true), reflect.TypeOf(myBool(false)), myBool(true), true},

		{MakeInt(123), reflect.TypeOf(myInt(0)), myInt(123), true},
		{MakeInt8(-123), reflect.TypeOf(myInt8(0)), myInt8(-123), true},
		{MakeInt16(123), reflect.TypeOf(myInt16(0)), myInt16(123), true},
		{MakeInt32(123), reflect.TypeOf(myInt32(0)), myInt32(123), true},
		{MakeInt64(123), reflect.TypeOf(myInt64(0)), myInt64(123), true},
		{MakeUint(123), reflect.TypeOf(myUint(0)), myUint(123), true},
		{MakeUint8(123), reflect.TypeOf(myUint8(0)), myUint8(123), true},
		{MakeUint16(123), reflect.TypeOf(myUint16(0)), myUint16(123), true},
		{MakeUint32(123), reflect.TypeOf(myUint32(0)), myUint32(123), true},
		{MakeUint64(123), reflect.TypeOf(myUint64(0)), myUint64(123), true},

		{MakeFloat32(123.456), reflect.TypeOf(myFloat32(0)), myFloat32(123.456), true},
		{MakeFloat64(123.456), reflect.TypeOf(myFloat64(0)), myFloat64(123.456), true},

		{MakeComplex64(1.2 + 3.4i), reflect.TypeOf(myComplex64(0)), myComplex64(1.2 + 3.4i), true},
		{MakeComplex128(1.2 - 3.4i), reflect.TypeOf(myComplex128(0)), myComplex128(1.2 - 3.4i), true},

		{MakeString("str"), reflect.TypeOf(myString("")), myString("str"), true},

		// no conversion
		{MakeInt64(123), reflect.TypeOf(myFloat32(0)), myFloat32(123), true},
		{MakeInt64(123), reflect.TypeOf(myComplex128(0)), myComplex128(123), true},

		{MakeFloat64(123.0), reflect.TypeOf(myInt16(0)), myInt16(123), true},
		{MakeFloat32(123.456), reflect.TypeOf(myComplex64(0)), myComplex64(123.456), true},

		{MakeComplex128(1.0 - 0i), reflect.TypeOf(myUint8(0)), myUint8(1), true},
		{MakeComplex128(1.2 + 0i), reflect.TypeOf(myFloat64(0)), myFloat64(1.2), true},

		// negative
		{MakeBool(true), reflecth.TypeInt(), nil, false},
		{MakeBool(true), reflecth.TypeUint8(), nil, false},
		{MakeBool(true), reflecth.TypeFloat64(), nil, false},
		{MakeBool(true), reflecth.TypeComplex64(), nil, false},
		{MakeBool(true), reflecth.TypeString(), nil, false},

		{MakeInt(123), reflecth.TypeBool(), nil, false},
		{MakeInt(123), reflecth.TypeString(), nil, false},

		{MakeFloat32(123.456), reflecth.TypeBool(), nil, false},
		{MakeFloat32(123.456), reflecth.TypeInt(), nil, false},
		{MakeFloat32(123.456), reflecth.TypeUint(), nil, false},
		{MakeFloat32(123.456), reflecth.TypeString(), nil, false},

		{MakeComplex128(1.2 - 3.4i), reflecth.TypeBool(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeInt(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeUint(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeFloat64(), nil, false},
		{MakeComplex128(1.2 - 3.4i), reflecth.TypeString(), nil, false},

		{MakeString("str"), reflecth.TypeBool(), nil, false},
		{MakeString("str"), reflecth.TypeInt(), nil, false},
		{MakeString("str"), reflecth.TypeUint(), nil, false},
		{MakeString("str"), reflecth.TypeFloat64(), nil, false},
		{MakeString("str"), reflecth.TypeComplex64(), nil, false},

		{MakeInt(123), reflect.TypeOf(struct{ x int }{0}), nil, false},
	}

	for _, test := range tests {
		r, ok := Assign(test.c, test.t)

		if ok != test.ok || (ok && r.Interface() != test.r) {
			t.Errorf("%v %v: expect %v (type %v) %v, got %v (type %v) %v", test.c, test.t, test.r, reflect.TypeOf(test.r).String(), test.ok, r, r.Type().String(), ok)
		}
		if r := IsAssignable(test.c, test.t); r != test.ok {
			t.Errorf("%v %v: expect %v, got %v", test.c, test.t, test.ok, r)
		}
		if !test.ok {
			continue
		}
		if r := MustAssign(test.c, test.t); r.Interface() != test.r {
			t.Errorf("%v %v: expect %v, got %v", test.c, test.t, test.r, r)
		}
	}
}

func TestAssignToInterface(t *testing.T) {
	c := MakeInt64(123)

	var tmp0 interface{}
	targetT := reflecth.TypeOfPtr(&tmp0)
	if r, ok := Assign(c, targetT); !ok || r.Type() != targetT || r.Interface() != int(123) {
		t.Errorf("expect %v %v %v, got %v %v %v", int(123), targetT, true, r.Interface(), r.Type(), ok)
	}

	var tmp1 interface {
		f1() string
	}
	targetT = reflecth.TypeOfPtr(&tmp1)
	if _, ok := Assign(c, targetT); ok {
		t.Error("expect false")
	}
}

func TestMustAssign(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = MustAssign(MakeBool(true), reflecth.TypeInt())
	t.Error("panic expected")
}

package constanth

import (
	"github.com/apaxa-go/helper/reflecth"
	"go/constant"
	"go/token"
	"reflect"
	"testing"
)

func TestMustMakeTypedValue(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = MustMakeTypedValue(MakeString("str"), reflecth.TypeInt())
	t.Error("expect panic")
}

func TestTypedValue_Type(t *testing.T) {
	c := MakeString("str")
	cT := MustMakeTypedValue(c, reflecth.TypeString())

	if cT.Type() != reflecth.TypeString() {
		t.Errorf("expect %v, got %v", reflecth.TypeString(), cT.Type())
	}
}

func TestTypedValue_Untyped(t *testing.T) {
	c := MakeString("str")
	cT := MustMakeTypedValue(c, reflecth.TypeString())

	if cT.Untyped() != c {
		t.Errorf("expect %v, got %v", c, cT.Untyped())
	}
}

func TestTypedValue_Equal(t *testing.T) {
	c0 := MakeString("str0")
	c1 := MakeString("str1")
	cT0 := MustMakeTypedValue(c0, reflecth.TypeString())
	cT1 := MustMakeTypedValue(c1, reflecth.TypeString())

	if !cT0.Equal(cT0) {
		t.Error("expect true")
	}
	if cT0.Equal(cT1) {
		t.Error("expect false")
	}
}

func TestTypedValue_Value(t *testing.T) {
	const s = "str"
	c := MakeString(s)
	cT := MustMakeTypedValue(c, reflecth.TypeString())
	cV := cT.Value()
	if cV.Interface() != s {
		t.Error("expect equal")
	}
}

func TestTypedValue_Assign(t *testing.T) {
	type testElement struct {
		xC constant.Value
		xT reflect.Type
		t  reflect.Type
		r  interface{}
		ok bool
	}

	type myInt int

	tests := []testElement{
		{MakeString("str"), reflecth.TypeString(), reflecth.TypeString(), "str", true},
		{MakeString("str"), reflecth.TypeString(), reflecth.TypeInt(), nil, false},
		{MakeInt(65), reflecth.TypeInt(), reflecth.TypeString(), nil, false},
		{MakeInt(65), reflecth.TypeInt(), reflect.TypeOf(myInt(0)), myInt(65), false},
	}

	for _, test := range tests {
		x := MustMakeTypedValue(test.xC, test.xT)

		if r, ok := x.Assign(test.t); ok != test.ok || (ok && (r.Interface() != test.r)) {
			t.Errorf("%v %v: expect %v %v, got %v %v", x, test.t, test.r, test.ok, r, ok)
		}

		if ok := x.AssignableTo(test.t); ok != test.ok {
			t.Errorf("%v %v: expect %v, got %v", x, test.t, test.ok, ok)
		}

		if !test.ok {
			continue
		}

		if r := x.MustAssign(test.t); r.Interface() != test.r {
			t.Errorf("%v %v: expect %v, got %v", x, test.t, test.r, r)
		}
	}
}

func TestTypedValue_AssignToInterface(t *testing.T) {
	type myInt int
	targetT := reflect.TypeOf(myInt(0))
	c := MustMakeTypedValue(MakeInt64(123), targetT)

	var tmp0 interface{}
	targetT = reflecth.TypeOfPtr(&tmp0)
	if r, ok := c.Assign(targetT); !ok || r.Type() != targetT || r.Interface() != myInt(123) {
		t.Errorf("expect %v %v %v, got %v %v %v", myInt(123), targetT, true, r.Interface(), r.Type(), ok)
	}

	var tmp1 interface {
		f1() string
	}
	targetT = reflecth.TypeOfPtr(&tmp1)
	if _, ok := c.Assign(targetT); ok {
		t.Error("expect false")
	}
}

func TestTypedValue_MustAssign(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = MustMakeTypedValue(MakeString("str"), reflecth.TypeString()).MustAssign(reflecth.TypeInt())
	t.Error("expect panic")
}

func TestTypedValue_Convert(t *testing.T) {
	type testElement struct {
		xC constant.Value
		xT reflect.Type
		t  reflect.Type
		rC constant.Value
		ok bool
	}

	type myInt int

	tests := []testElement{
		{MakeString("str"), reflecth.TypeString(), reflecth.TypeString(), nil, true},
		{MakeString("str"), reflecth.TypeString(), reflecth.TypeInt(), nil, false},
		{MakeInt(65), reflecth.TypeInt(), reflecth.TypeString(), MakeString("A"), true},
		{MakeInt(65), reflecth.TypeInt(), reflect.TypeOf(myInt(0)), nil, true},
	}

	for _, test := range tests {
		if test.ok && test.rC == nil {
			test.rC = test.xC
		}

		x := MustMakeTypedValue(test.xC, test.xT)
		testR := TypedValue{test.rC, test.t}

		if r, ok := x.Convert(test.t); ok != test.ok || (ok && (!r.Equal(testR))) {
			t.Errorf("%v %v: expect %v %v, got %v %v", x, test.t, testR, test.ok, r, ok)
		}

		if ok := x.ConvertibleTo(test.t); ok != test.ok {
			t.Errorf("%v %v: expect %v, got %v", x, test.t, test.ok, ok)
		}

		if !test.ok {
			continue
		}

		if r := x.MustConvert(test.t); !r.Equal(testR) {
			t.Errorf("%v %v: expect %v, got %v", x, test.t, testR, r)
		}
	}
}

func TestTypedValue_String(t *testing.T) {
	c := constant.MakeFromLiteral("123.456", token.FLOAT, 0)
	ct := MustMakeTypedValue(c, reflecth.TypeFloat32())
	cs := "123.456"
	ces := "15432/125"
	cst := "123.456 (type float32)"
	cest := "15432/125 (type float32)"
	if r := ct.String(); r != cs {
		t.Errorf(`expect "%v", got "%v"`, cs, r)
	}
	if r := ct.ExactString(); r != ces {
		t.Errorf(`expect "%v", got "%v"`, ces, r)
	}
	if r := ct.StringType(); r != cst {
		t.Errorf(`expect "%v", got "%v"`, cst, r)
	}
	if r := ct.ExactStringType(); r != cest {
		t.Errorf(`expect "%v", got "%v"`, cest, r)
	}
}

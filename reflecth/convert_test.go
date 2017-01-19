package reflecth

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type testElement struct {
		x  interface{}
		t  reflect.Type
		ok bool
		r  interface{}
	}

	tests := []testElement{
		{int(0), TypeFloat32(), true, float32(0)},
		{float32(0), TypeString(), false, nil},
	}

	for i, test := range tests {
		if ok := ConvertibleTo(reflect.ValueOf(test.x), test.t); ok != test.ok {
			t.Errorf("#%v expect %v, got %v", i, test.ok, ok)
			continue
		}
		if r, ok := Convert(reflect.ValueOf(test.x), test.t); ok != test.ok || (ok && r.Interface() != test.r) {
			t.Errorf("#%v expect %v %v, got %v %v", i, test.r, test.ok, r, ok)
			continue
		}
		if !test.ok { // Check panic in separately
			continue
		}
		if r := MustConvert(reflect.ValueOf(test.x), test.t); r.Interface() != test.r {
			t.Errorf("#%v expect %v, got %v", i, test.r, r)
			continue
		}
	}
}

func TestMustConvert(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = MustConvert(reflect.ValueOf("string"), TypeInt())
	t.Error("panic expected")
}

func TestAssign(t *testing.T) {
	type testElement struct {
		x  interface{}
		t  reflect.Type
		ok bool
		r  interface{}
	}

	type myType struct {
		x int
	}

	tests := []testElement{
		{struct{ x int }{1}, reflect.TypeOf(myType{}), true, myType{1}},
		{float32(0), TypeString(), false, nil},
	}

	for i, test := range tests {
		if ok := AssignableTo(reflect.ValueOf(test.x), test.t); ok != test.ok {
			t.Errorf("#%v expect %v, got %v", i, test.ok, ok)
			continue
		}
		if r, ok := Assign(reflect.ValueOf(test.x), test.t); ok != test.ok || (ok && r.Interface() != test.r) {
			t.Errorf("#%v expect %v %v, got %v %v", i, test.r, test.ok, r, ok)
			continue
		}
		if !test.ok { // Check panic in separately
			continue
		}
		if r := MustAssign(reflect.ValueOf(test.x), test.t); r.Interface() != test.r {
			t.Errorf("#%v expect %v, got %v", i, test.r, r)
			continue
		}
	}
}

func TestMustAssign(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = MustAssign(reflect.ValueOf("string"), TypeInt())
	t.Error("panic expected")
}

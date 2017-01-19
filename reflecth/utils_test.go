package reflecth

import (
	"github.com/apaxa-go/helper/goh/asth"
	"github.com/apaxa-go/helper/strconvh"
	"go/ast"
	"reflect"
	"testing"
)

func TestTypeOfPtr(t *testing.T) {
	type testElement struct {
		i interface{}
		r reflect.Type
	}

	var b bool

	tests := []testElement{
		{&b, TypeBool()},
		{b, nil},
	}

	for i, test := range tests {
		if r := TypeOfPtr(test.i); r != test.r {
			t.Errorf("#%v expect %v, got %v", i, test.r, r)
		}
	}
}

func TestValueOfPtr(t *testing.T) {
	var b = true

	if bV := ValueOfPtr(&b); bV.Type() != TypeBool() || bV.Bool() != b {
		t.Errorf("expect " + strconvh.FormatBool(b) + " (type bool), got " + bV.String())
	}

	if bV := ValueOfPtr(b); bV != (reflect.Value{}) {
		t.Errorf("expect zero value, got " + bV.String())
	}
}

func TestChanDir(t *testing.T) {
	type testElement struct {
		a ast.ChanDir
		r reflect.ChanDir
	}

	tests := []testElement{
		{asth.BothDir, reflect.BothDir},
		{asth.RecvDir, reflect.RecvDir},
		{asth.SendDir, reflect.SendDir},
		{0, 0},
	}

	for _, test := range tests {
		if r := ChanDirFromAst(test.a); r != test.r {
			t.Errorf("%v: expect %v, got %v", test.a, test.r, r)
		}
		if a := ChanDirToAst(test.r); a != test.a {
			t.Errorf("%v: expect %v, got %v", test.r, test.a, a)
		}
	}
}

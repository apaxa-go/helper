package evalh

import (
	"go/ast"
	"go/constant"
	"go/parser"
	"reflect"
	"testing"
)

/*
func TestExpr(t *testing.T) {
	type testElement struct {
		expr string
		vars []Var
		r    interface{}
		err  bool
	}

	tests := []testElement{
		{"1", nil, 1, false},
		{"1+2", nil, 3, false},
		{"\"abc\"==\"abc\"", nil, true, false},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}

		r, err := Expr(exprAst, v.vars)

		if err != nil != v.err || (!v.err && r != v.r) {
			t.Errorf("%v: expect %v %v, got %v %v", v.expr, v.r, v.err, r, err)
		}
	}
}
*/

func TestBinary(t *testing.T) {
	// TODO add more tests on complex
	type testElement struct {
		expr string
		vars []Var
		r    interface{}
		err  bool
	}

	tests := []testElement{
		{"1+2", nil, constant.MakeInt64(3), false},
		{"1+\"2\"", nil, nil, true},
		{"1<<2", nil, constant.MakeInt64(4), false},
		{"\"1\"<<2", nil, nil, true},
		{"1<<\"2\"", nil, nil, true},
		{"1==2", nil, constant.MakeBool(false), false},
		{"1==\"2\"", nil, nil, true},
		{"a+2", []Var{{"a", 1}}, 3, false},
		{"a+\"2\"", []Var{{"a", 1}}, nil, true},
		{"2+a", []Var{{"a", 1}}, 3, false},
		{"\"2\"+a", []Var{{"a", 1}}, nil, true},
		{"a+b", []Var{{"a", 1}, {"b", 2}}, 3, false},
		{"a+b", []Var{{"a", 1}, {"b", "2"}}, nil, true},
		{"true&&a", []Var{{"a", false}}, false, false},
		{"a+b", []Var{{"a", "1"}, {"b", "2"}}, "12", false},
		{"a+b", []Var{{"a", 1}, {"b", int8(2)}}, nil, true},
		{"a+2", []Var{{"a", int8(1)}}, int8(3), false},
		{"a+2", []Var{{"a", int16(1)}}, int16(3), false},
		{"a+2", []Var{{"a", int32(1)}}, int32(3), false},
		{"a+2", []Var{{"a", int64(1)}}, int64(3), false},
		{"a+2", []Var{{"a", uint(1)}}, uint(3), false},
		{"a+2", []Var{{"a", uint8(1)}}, uint8(3), false},
		{"a+2", []Var{{"a", uint16(1)}}, uint16(3), false},
		{"a+2", []Var{{"a", uint32(1)}}, uint32(3), false},
		{"a+2", []Var{{"a", uint64(1)}}, uint64(3), false},
		{"a+2.0", []Var{{"a", uint64(1)}}, uint64(3), false},
		{"a+2", []Var{{"a", float32(1)}}, float32(3), false},
		{"a+2", []Var{{"a", float64(1)}}, float64(3), false},
		{"a+0-11i", []Var{{"a", complex64(2 + 3i)}}, complex64(2 - 8i), false},
		{"a+0-11i", []Var{{"a", complex128(2 + 3i)}}, complex128(2 - 8i), false},
		// shift
		{"a<<b", []Var{{"a", 4}, {"b", 2}}, nil, true},
		{"a>>b", []Var{{"a", 4}, {"b", 2}}, nil, true},
		{"a<<b", []Var{{"a", 4}, {"b", uint8(2)}}, 16, false},
		{"a>>b", []Var{{"a", 4}, {"b", uint8(2)}}, 1, false},
		{"a<<b", []Var{{"a", int8(4)}, {"b", uint16(2)}}, int8(16), false},
		{"a>>b", []Var{{"a", int8(4)}, {"b", uint16(2)}}, int8(1), false},
		{"a<<b", []Var{{"a", int16(4)}, {"b", uint32(2)}}, int16(16), false},
		{"a>>b", []Var{{"a", int16(4)}, {"b", uint32(2)}}, int16(1), false},
		{"a<<b", []Var{{"a", int32(4)}, {"b", uint64(2)}}, int32(16), false},
		{"a>>b", []Var{{"a", int32(4)}, {"b", uint64(2)}}, int32(1), false},
		{"a<<b", []Var{{"a", int64(4)}, {"b", uint(2)}}, int64(16), false},
		{"a>>b", []Var{{"a", int64(4)}, {"b", uint(2)}}, int64(1), false},
		{"a<<b", []Var{{"a", uint(4)}, {"b", uint8(2)}}, uint(16), false},
		{"a>>b", []Var{{"a", uint(4)}, {"b", uint8(2)}}, uint(1), false},
		{"a<<b", []Var{{"a", uint8(4)}, {"b", uint16(2)}}, uint8(16), false},
		{"a>>b", []Var{{"a", uint8(4)}, {"b", uint16(2)}}, uint8(1), false},
		{"a<<b", []Var{{"a", uint16(4)}, {"b", uint32(2)}}, uint16(16), false},
		{"a>>b", []Var{{"a", uint16(4)}, {"b", uint32(2)}}, uint16(1), false},
		{"a<<b", []Var{{"a", uint32(4)}, {"b", uint64(2)}}, uint32(16), false},
		{"a>>b", []Var{{"a", uint32(4)}, {"b", uint64(2)}}, uint32(1), false},
		{"a<<b", []Var{{"a", uint64(4)}, {"b", uint(2)}}, uint64(16), false},
		{"a>>b", []Var{{"a", uint64(4)}, {"b", uint(2)}}, uint64(1), false},
		{"4<<2", nil, constant.MakeInt64(16), false},
		{"4>>2", nil, constant.MakeInt64(1), false},
		{"4<<a", []Var{{"a", uint(2)}}, constant.MakeInt64(16), false},
		{"a>>2", []Var{{"a", int(4)}}, 1, false},
		{`"4"<<a`, []Var{{"a", uint(2)}}, nil, true},
		{`a>>"2"`, []Var{{"a", int(4)}}, nil, true},
		{`"4">>2`, nil, nil, true},
		{`4>>"2"`, nil, nil, true},
		// binary compare
		{"a==b", []Var{{"a", 1}, {"b", 2}}, constant.MakeBool(false), false},
		{"a>=b", []Var{{"a", int8(1)}, {"b", int8(2)}}, constant.MakeBool(false), false},
		{"a<=b", []Var{{"a", int16(1)}, {"b", int16(2)}}, constant.MakeBool(true), false},
		{"a!=b", []Var{{"a", int32(1)}, {"b", int32(2)}}, constant.MakeBool(true), false},
		{"a>b", []Var{{"a", int64(1)}, {"b", int64(2)}}, constant.MakeBool(false), false},
		{"a<b", []Var{{"a", int64(1)}, {"b", int64(2)}}, constant.MakeBool(true), false},
		{"a==b", []Var{{"a", uint(1)}, {"b", uint(2)}}, constant.MakeBool(false), false},
		{"a>=b", []Var{{"a", uint8(1)}, {"b", uint8(2)}}, constant.MakeBool(false), false},
		{"a<=b", []Var{{"a", uint16(1)}, {"b", uint16(2)}}, constant.MakeBool(true), false},
		{"a!=b", []Var{{"a", uint32(1)}, {"b", uint32(2)}}, constant.MakeBool(true), false},
		{"a>b", []Var{{"a", uint64(1)}, {"b", uint64(2)}}, constant.MakeBool(false), false},
		{"a<b", []Var{{"a", uint64(1)}, {"b", uint64(2)}}, constant.MakeBool(true), false},
		{"a==b", []Var{{"a", float32(1)}, {"b", float32(2)}}, constant.MakeBool(false), false},
		{"a>=b", []Var{{"a", float64(1)}, {"b", float64(2)}}, constant.MakeBool(false), false},
		{"a<=b", []Var{{"a", float32(1)}, {"b", float32(2)}}, constant.MakeBool(true), false},
		{"a!=b", []Var{{"a", float64(1)}, {"b", float64(2)}}, constant.MakeBool(true), false},
		{"a>b", []Var{{"a", float32(1)}, {"b", float32(2)}}, constant.MakeBool(false), false},
		{"a<b", []Var{{"a", float64(1)}, {"b", float64(2)}}, constant.MakeBool(true), false},
		{"a==b", []Var{{"a", "1"}, {"b", "2"}}, constant.MakeBool(false), false},
		{"a>=b", []Var{{"a", "1"}, {"b", "2"}}, constant.MakeBool(false), false},
		{"a<=b", []Var{{"a", "1"}, {"b", "2"}}, constant.MakeBool(true), false},
		{"a!=b", []Var{{"a", "1"}, {"b", "2"}}, constant.MakeBool(true), false},
		{"a>b", []Var{{"a", "1"}, {"b", "2"}}, constant.MakeBool(false), false},
		{"a<b", []Var{{"a", "1"}, {"b", "2"}}, constant.MakeBool(true), false},
		{"a==b", []Var{{"a", true}, {"b", false}}, constant.MakeBool(false), false},
		{"a!=b", []Var{{"a", true}, {"b", false}}, constant.MakeBool(true), false},
		{"a==b", []Var{{"a", complex64(1 - 2i)}, {"b", complex64(1 - 2i)}}, constant.MakeBool(true), false},
		{"a!=b", []Var{{"a", complex128(1 - 2i)}, {"b", complex128(2 + 3i)}}, constant.MakeBool(true), false},
		{"a==b", []Var{{"a", uintptr(1)}, {"b", uintptr(2)}}, constant.MakeBool(false), false},
		{"a!=b", []Var{{"a", uintptr(1)}, {"b", uintptr(2)}}, constant.MakeBool(true), false},
		{"a==1", []Var{{"a", uint8(1)}}, constant.MakeBool(true), false},
		{"2==a", []Var{{"a", int32(1)}}, constant.MakeBool(false), false},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		binaryAst, ok := exprAst.(*ast.BinaryExpr)
		if !ok {
			t.Errorf("%v: not a BinaryExpr", v.expr)
			continue
		}

		r, err := Binary(binaryAst, v.vars)
		if err != nil != v.err || r != v.r {
			var vrTS, rTS string
			if vrT := reflect.TypeOf(v.r); vrT != nil {
				vrTS = vrT.String()
			}
			if rT := reflect.TypeOf(r); rT != nil {
				rTS = rT.String()
			}

			t.Errorf("'%v' (%#v): expect %v(%v) %v, got %v(%v) %v", v.expr, v.vars, v.r, vrTS, v.err, r, rTS, err)
		}
	}
}

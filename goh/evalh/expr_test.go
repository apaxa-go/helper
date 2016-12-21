package evalh

import (
	"fmt"
	"github.com/apaxa-go/helper/mathh"
	"go/ast"
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

type testExprElement struct {
	expr string
	vars Identifiers
	r    Value
	err  bool
}

func (t testExprElement) Validate(r Value, err error) bool {
	// validate error
	if err != nil != t.err {
		return false
	}

	if t.r == nil && r == nil {
		return true
	}
	if t.r == nil || r == nil {
		return false
	}
	return t.r.Equal(r)
}

func (t testExprElement) ErrorMsg(r Value, err error) string {
	return fmt.Sprintf("'%v' (%+v): expect %v %v, got %v %v", t.expr, t.vars, t.r, t.err, r, err)
}

// SampleStruct is a sample structure with field F and method M
type SampleStruct struct {
	F uint16
}

func (s SampleStruct) M(x uint32) uint64 { return uint64(s.F) * uint64(x) }
func (s SampleStruct) M2(x uint32) (uint64, int64) {
	return uint64(s.F) * uint64(x), int64(s.F) - int64(x)
}

type SampleInt int8

func (s SampleInt) Mv(x int16) int32  { return -int32(s) * int32(x) }
func (s *SampleInt) Mp(x int16) int32 { return -int32(*s)*int32(x) + 1 }

func TestIdent(t *testing.T) {
	tests := []testExprElement{
		{"true", nil, MakeUntypedBool(true), false},
		{"false", nil, MakeUntypedBool(false), false},
		{"nil", nil, MakeNil(), false},
		{"v1", IdentifiersInterface{"v0": 2, "v1": 3, "v2": 4}.Identifiers(), MakeRegularInterface(int(3)), false},
		{"v10", IdentifiersInterface{"v0": 2, "v1": 3, "v2": 4}.Identifiers(), nil, true},
		{"pow", IdentifiersInterface{"v0": 2, "pow": mathh.PowInt16, "v2": 4}.Identifiers(), MakeRegularInterface(mathh.PowInt16), false},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		identAst, ok := exprAst.(*ast.Ident)
		if !ok {
			t.Errorf("%v: not an Ident", v.expr)
			continue
		}

		r, err := Ident(identAst, v.vars)
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

func TestSelector(t *testing.T) {
	tmp := &(SampleStruct{2})
	tests := []testExprElement{
		{"a.F", IdentifiersInterface{"a": SampleStruct{2}}.Identifiers(), MakeRegularInterface(uint16(2)), false},
		{"a.F", IdentifiersInterface{"a": &SampleStruct{2}}.Identifiers(), MakeRegularInterface(uint16(2)), false},
		{"a.F", IdentifiersInterface{"a": &tmp}.Identifiers(), nil, true}, // unable to double dereference on-the-fly, only single is possible
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		selectorAst, ok := exprAst.(*ast.SelectorExpr)
		if !ok {
			t.Errorf("%v: not a SelectorExpr", v.expr)
			continue
		}

		r, err := Selector(selectorAst, v.vars)
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

// Check for method extracting.
// It is not trivial to fully check returned value, so do it separately from other tests
func TestSelector2(t *testing.T) {
	type testSelectorElement struct {
		expr string
		vars Identifiers
		arg  interface{}
		r    interface{}
	}

	tests := []testSelectorElement{
		{"a.M", IdentifiersInterface{"a": SampleStruct{2}}.Identifiers(), uint32(3), uint64(6)},
		{"b.Mv", IdentifiersInterface{"b": SampleInt(4)}.Identifiers(), int16(5), int32(-20)},
		{"b.Mv", IdentifiersInterface{"b": new(SampleInt)}.Identifiers(), int16(6), int32(0)},
		{"b.Mp", IdentifiersInterface{"b": new(SampleInt)}.Identifiers(), int16(7), int32(1)},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		selectorAst, ok := exprAst.(*ast.SelectorExpr)
		if !ok {
			t.Errorf("%v: not a SelectorExpr", v.expr)
			continue
		}

		r, err := Selector(selectorAst, v.vars)
		if err != nil {
			t.Errorf("expect not error, got %v", err.Error())
			continue
		}

		rV := reflect.ValueOf(r.Interface())
		if rK := rV.Kind(); rK != reflect.Func {
			t.Errorf("expect function, got %v", rK.String())
			continue
		}

		rs := rV.Call([]reflect.Value{reflect.ValueOf(v.arg)})
		if l := len(rs); l != 1 {
			t.Errorf("expect 1 result, got %v", l)
			continue
		}

		if rI := rs[0].Interface(); rI != v.r {
			t.Errorf("expect %v, got %v", v.r, rI)
		}
	}
}

func TestBinary(t *testing.T) {
	// TODO add more tests on complex
	tests := []testExprElement{
		{"1+2", nil, MakeUntypedInt64(3), false},
		{"1+\"2\"", nil, nil, true},
		{"1<<2", nil, MakeUntypedInt64(4), false},
		{"\"1\"<<2", nil, nil, true},
		{"1<<\"2\"", nil, nil, true},
		{"1==2", nil, MakeUntypedBool(false), false},
		{"1==\"2\"", nil, nil, true},
		{"a+2", IdentifiersInterface{"a": 1}.Identifiers(), MakeRegularInterface(3), false},
		{"a+\"2\"", IdentifiersInterface{"a": 1}.Identifiers(), nil, true},
		{"2+a", IdentifiersInterface{"a": 1}.Identifiers(), MakeRegularInterface(3), false},
		{"\"2\"+a", IdentifiersInterface{"a": 1}.Identifiers(), nil, true},
		{"a+b", IdentifiersInterface{"a": 1, "b": 2}.Identifiers(), MakeRegularInterface(3), false},
		{"a+b", IdentifiersInterface{"a": 1, "b": "2"}.Identifiers(), nil, true},
		{"true&&a", IdentifiersInterface{"a": false}.Identifiers(), MakeRegularInterface(false), false},
		{"a+b", IdentifiersInterface{"a": "1", "b": "2"}.Identifiers(), MakeRegularInterface("12"), false},
		{"a+b", IdentifiersInterface{"a": 1, "b": int8(2)}.Identifiers(), nil, true},
		{"a+2", IdentifiersInterface{"a": int8(1)}.Identifiers(), MakeRegularInterface(int8(3)), false},
		{"a+2", IdentifiersInterface{"a": int16(1)}.Identifiers(), MakeRegularInterface(int16(3)), false},
		{"a+2", IdentifiersInterface{"a": int32(1)}.Identifiers(), MakeRegularInterface(int32(3)), false},
		{"a+2", IdentifiersInterface{"a": int64(1)}.Identifiers(), MakeRegularInterface(int64(3)), false},
		{"a+2", IdentifiersInterface{"a": uint(1)}.Identifiers(), MakeRegularInterface(uint(3)), false},
		{"a+2", IdentifiersInterface{"a": uint8(1)}.Identifiers(), MakeRegularInterface(uint8(3)), false},
		{"a+2", IdentifiersInterface{"a": uint16(1)}.Identifiers(), MakeRegularInterface(uint16(3)), false},
		{"a+2", IdentifiersInterface{"a": uint32(1)}.Identifiers(), MakeRegularInterface(uint32(3)), false},
		{"a+2", IdentifiersInterface{"a": uint64(1)}.Identifiers(), MakeRegularInterface(uint64(3)), false},
		{"a+2.0", IdentifiersInterface{"a": uint64(1)}.Identifiers(), MakeRegularInterface(uint64(3)), false},
		{"a+2", IdentifiersInterface{"a": float32(1)}.Identifiers(), MakeRegularInterface(float32(3)), false},
		{"a+2", IdentifiersInterface{"a": float64(1)}.Identifiers(), MakeRegularInterface(float64(3)), false},
		{"a+0-11i", IdentifiersInterface{"a": complex64(2 + 3i)}.Identifiers(), MakeRegularInterface(complex64(2 - 8i)), false},
		{"a+0-11i", IdentifiersInterface{"a": complex128(2 + 3i)}.Identifiers(), MakeRegularInterface(complex128(2 - 8i)), false},
		// shift
		{"a<<b", IdentifiersInterface{"a": 4, "b": 2}.Identifiers(), nil, true},
		{"a>>b", IdentifiersInterface{"a": 4, "b": 2}.Identifiers(), nil, true},
		{"a<<b", IdentifiersInterface{"a": 4, "b": uint8(2)}.Identifiers(), MakeRegularInterface(16), false},
		{"a>>b", IdentifiersInterface{"a": 4, "b": uint8(2)}.Identifiers(), MakeRegularInterface(1), false},
		{"a<<b", IdentifiersInterface{"a": int8(4), "b": uint16(2)}.Identifiers(), MakeRegularInterface(int8(16)), false},
		{"a>>b", IdentifiersInterface{"a": int8(4), "b": uint16(2)}.Identifiers(), MakeRegularInterface(int8(1)), false},
		{"a<<b", IdentifiersInterface{"a": int16(4), "b": uint32(2)}.Identifiers(), MakeRegularInterface(int16(16)), false},
		{"a>>b", IdentifiersInterface{"a": int16(4), "b": uint32(2)}.Identifiers(), MakeRegularInterface(int16(1)), false},
		{"a<<b", IdentifiersInterface{"a": int32(4), "b": uint64(2)}.Identifiers(), MakeRegularInterface(int32(16)), false},
		{"a>>b", IdentifiersInterface{"a": int32(4), "b": uint64(2)}.Identifiers(), MakeRegularInterface(int32(1)), false},
		{"a<<b", IdentifiersInterface{"a": int64(4), "b": uint(2)}.Identifiers(), MakeRegularInterface(int64(16)), false},
		{"a>>b", IdentifiersInterface{"a": int64(4), "b": uint(2)}.Identifiers(), MakeRegularInterface(int64(1)), false},
		{"a<<b", IdentifiersInterface{"a": uint(4), "b": uint8(2)}.Identifiers(), MakeRegularInterface(uint(16)), false},
		{"a>>b", IdentifiersInterface{"a": uint(4), "b": uint8(2)}.Identifiers(), MakeRegularInterface(uint(1)), false},
		{"a<<b", IdentifiersInterface{"a": uint8(4), "b": uint16(2)}.Identifiers(), MakeRegularInterface(uint8(16)), false},
		{"a>>b", IdentifiersInterface{"a": uint8(4), "b": uint16(2)}.Identifiers(), MakeRegularInterface(uint8(1)), false},
		{"a<<b", IdentifiersInterface{"a": uint16(4), "b": uint32(2)}.Identifiers(), MakeRegularInterface(uint16(16)), false},
		{"a>>b", IdentifiersInterface{"a": uint16(4), "b": uint32(2)}.Identifiers(), MakeRegularInterface(uint16(1)), false},
		{"a<<b", IdentifiersInterface{"a": uint32(4), "b": uint64(2)}.Identifiers(), MakeRegularInterface(uint32(16)), false},
		{"a>>b", IdentifiersInterface{"a": uint32(4), "b": uint64(2)}.Identifiers(), MakeRegularInterface(uint32(1)), false},
		{"a<<b", IdentifiersInterface{"a": uint64(4), "b": uint(2)}.Identifiers(), MakeRegularInterface(uint64(16)), false},
		{"a>>b", IdentifiersInterface{"a": uint64(4), "b": uint(2)}.Identifiers(), MakeRegularInterface(uint64(1)), false},
		{"4<<2", nil, MakeUntypedInt64(16), false},
		{"4>>2", nil, MakeUntypedInt64(1), false},
		{"4<<a", IdentifiersInterface{"a": uint(2)}.Identifiers(), MakeUntypedInt64(16), false},
		{"a>>2", IdentifiersInterface{"a": int(4)}.Identifiers(), MakeRegularInterface(1), false},
		{`"4"<<a`, IdentifiersInterface{"a": uint(2)}.Identifiers(), nil, true},
		{`a>>"2"`, IdentifiersInterface{"a": int(4)}.Identifiers(), nil, true},
		{`"4">>2`, nil, nil, true},
		{`4>>"2"`, nil, nil, true},
		// binary compare
		{"a==b", IdentifiersInterface{"a": 1, "b": 2}.Identifiers(), MakeUntypedBool(false), false},
		{"a>=b", IdentifiersInterface{"a": int8(1), "b": int8(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a<=b", IdentifiersInterface{"a": int16(1), "b": int16(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a!=b", IdentifiersInterface{"a": int32(1), "b": int32(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a>b", IdentifiersInterface{"a": int64(1), "b": int64(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a<b", IdentifiersInterface{"a": int64(1), "b": int64(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a==b", IdentifiersInterface{"a": uint(1), "b": uint(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a>=b", IdentifiersInterface{"a": uint8(1), "b": uint8(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a<=b", IdentifiersInterface{"a": uint16(1), "b": uint16(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a!=b", IdentifiersInterface{"a": uint32(1), "b": uint32(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a>b", IdentifiersInterface{"a": uint64(1), "b": uint64(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a<b", IdentifiersInterface{"a": uint64(1), "b": uint64(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a==b", IdentifiersInterface{"a": float32(1), "b": float32(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a>=b", IdentifiersInterface{"a": float64(1), "b": float64(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a<=b", IdentifiersInterface{"a": float32(1), "b": float32(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a!=b", IdentifiersInterface{"a": float64(1), "b": float64(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a>b", IdentifiersInterface{"a": float32(1), "b": float32(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a<b", IdentifiersInterface{"a": float64(1), "b": float64(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a==b", IdentifiersInterface{"a": "1", "b": "2"}.Identifiers(), MakeUntypedBool(false), false},
		{"a>=b", IdentifiersInterface{"a": "1", "b": "2"}.Identifiers(), MakeUntypedBool(false), false},
		{"a<=b", IdentifiersInterface{"a": "1", "b": "2"}.Identifiers(), MakeUntypedBool(true), false},
		{"a!=b", IdentifiersInterface{"a": "1", "b": "2"}.Identifiers(), MakeUntypedBool(true), false},
		{"a>b", IdentifiersInterface{"a": "1", "b": "2"}.Identifiers(), MakeUntypedBool(false), false},
		{"a<b", IdentifiersInterface{"a": "1", "b": "2"}.Identifiers(), MakeUntypedBool(true), false},
		{"a==b", IdentifiersInterface{"a": true, "b": false}.Identifiers(), MakeUntypedBool(false), false},
		{"a!=b", IdentifiersInterface{"a": true, "b": false}.Identifiers(), MakeUntypedBool(true), false},
		{"a==b", IdentifiersInterface{"a": complex64(1 - 2i), "b": complex64(1 - 2i)}.Identifiers(), MakeUntypedBool(true), false},
		{"a!=b", IdentifiersInterface{"a": complex128(1 - 2i), "b": complex128(2 + 3i)}.Identifiers(), MakeUntypedBool(true), false},
		{"a==b", IdentifiersInterface{"a": uintptr(1), "b": uintptr(2)}.Identifiers(), MakeUntypedBool(false), false},
		{"a!=b", IdentifiersInterface{"a": uintptr(1), "b": uintptr(2)}.Identifiers(), MakeUntypedBool(true), false},
		{"a==1", IdentifiersInterface{"a": uint8(1)}.Identifiers(), MakeUntypedBool(true), false},
		{"2==a", IdentifiersInterface{"a": int32(1)}.Identifiers(), MakeUntypedBool(false), false},
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
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

func TestCall(t *testing.T) {
	tests := []testExprElement{
		{"f(3)", IdentifiersInterface{"f": func(x uint8) uint64 { return 2 * uint64(x) }}.Identifiers(), MakeRegularInterface(uint64(6)), false},
		{"f(2)", nil, nil, true},
		{"a.M(3)", IdentifiersInterface{"a": SampleStruct{2}}.Identifiers(), MakeRegularInterface(uint64(6)), false},
		{"a.M(5)", IdentifiersInterface{"a": &SampleStruct{4}}.Identifiers(), MakeRegularInterface(uint64(20)), false},
		{"a.M(b)", IdentifiersInterface{"a": &SampleStruct{4}, "b": uint32(5)}.Identifiers(), MakeRegularInterface(uint64(20)), false},
		{"a.M9(7)", IdentifiersInterface{"a": &SampleStruct{6}}.Identifiers(), nil, true},
		{"a.F(7)", IdentifiersInterface{"a": &SampleStruct{6}}.Identifiers(), nil, true},
		{"a.M(7,8)", IdentifiersInterface{"a": &SampleStruct{6}}.Identifiers(), nil, true},
		{"a.M2(7)", IdentifiersInterface{"a": &SampleStruct{6}}.Identifiers(), nil, true},
		{"a.M(b)", IdentifiersInterface{"a": &SampleStruct{6}}.Identifiers(), nil, true},
		{`a.M("7")`, IdentifiersInterface{"a": &SampleStruct{6}}.Identifiers(), nil, true},
		{"a.M(b)", IdentifiersInterface{"a": &SampleStruct{6}, "b": "bad"}.Identifiers(), nil, true},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		callAst, ok := exprAst.(*ast.CallExpr)
		if !ok {
			t.Errorf("%v: not a CallExpr", v.expr)
			continue
		}

		r, err := Call(callAst, v.vars)
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

func TestStar(t *testing.T) {
	tests := []testExprElement{
		{"*v", IdentifiersInterface{"v": new(int8)}.Identifiers(), MakeRegularInterface(int8(0)), false},
		{"*v", IdentifiersInterface{"v": int8(3)}.Identifiers(), nil, true},
		{"*v", nil, nil, true},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		starAst, ok := exprAst.(*ast.StarExpr)
		if !ok {
			t.Errorf("%v: not a StarExpr", v.expr)
			continue
		}

		r, err := Star(starAst, v.vars)
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

func TestParen(t *testing.T) {
	tests := []testExprElement{
		{"(v)", IdentifiersInterface{"v": int8(3)}.Identifiers(), MakeRegularInterface(int8(3)), false},
		{"(v)", nil, nil, true},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		parenAst, ok := exprAst.(*ast.ParenExpr)
		if !ok {
			t.Errorf("%v: not a ParenExpr", v.expr)
			continue
		}

		r, err := Paren(parenAst, v.vars)
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

func TestUnary(t *testing.T) {
	tests := []testExprElement{
		{"-1", nil, MakeUntypedInt64(-1), false},
		{"+2", nil, MakeUntypedInt64(+2), false},
		{"-a", IdentifiersInterface{"a": int8(3)}.Identifiers(), MakeRegularInterface(int8(-3)), false},
		{"+a", IdentifiersInterface{"a": int8(4)}.Identifiers(), MakeRegularInterface(int8(4)), false},
		{"^a", IdentifiersInterface{"a": int8(5)}.Identifiers(), MakeRegularInterface(int8(-6)), false},
		{"!a", IdentifiersInterface{"a": true}.Identifiers(), MakeRegularInterface(false), false},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		unaryAst, ok := exprAst.(*ast.UnaryExpr)
		if !ok {
			t.Errorf("%v: not a UnaryExpr", v.expr)
			continue
		}

		r, err := Unary(unaryAst, v.vars)
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

// Check for getting address (&).
// It is not trivial to fully check returned value, so do it separately from other tests
func TestUnary2(t *testing.T) {
	tmp := SampleStruct{5}
	tmp2 := []int8{6}
	tests := []testExprElement{
		{"&a.F", IdentifiersInterface{"a": &tmp}.Identifiers(), MakeRegularInterface(&tmp.F), false},
		{"&a[0]", IdentifiersInterface{"a": tmp2}.Identifiers(), MakeRegularInterface(&tmp2[0]), false},
	}

	for _, v := range tests {
		exprAst, err := parser.ParseExpr(v.expr)
		if err != nil {
			t.Errorf("%v: %v", v.expr, err)
			continue
		}
		unaryAst, ok := exprAst.(*ast.UnaryExpr)
		if !ok {
			t.Errorf("%v: not a UnaryExpr", v.expr)
			continue
		}

		r, err := Unary(unaryAst, v.vars)
		if !v.Validate(r, err) {
			t.Errorf(v.ErrorMsg(r, err))
		}
	}
}

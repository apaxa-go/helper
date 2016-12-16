package evalh

import (
	"errors"
	"fmt"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/goh/tokenh"
	"github.com/apaxa-go/helper/strconvh"
	"go/ast"
	"go/constant"
	"reflect"
)

type Var struct {
	Name  string
	Value interface{}
}

// r is a typed variable, untyped bool constant or nil
func Ident(e *ast.Ident, vars []Var) (r interface{}, err error) {
	// true and false literals does not covered by BasicLit (because it is look like normal identifier?).
	if e.Name == "true" {
		return constant.MakeBool(true), nil
	}
	if e.Name == "false" {
		return constant.MakeBool(false), nil
	}

	// nil literal
	if e.Name == "nil" {
		return nil, nil
	}

	// find required identifier in vars.
	for i := range vars {
		if vars[i].Name == e.Name {
			return vars[i].Value, nil
		}
	}
	return nil, errors.New("variable '" + e.Name + "' does not defined")
}

// r is a typed variable
func Selector(e *ast.SelectorExpr, vars []Var) (r interface{}, err error) {
	// Calc object (left of '.')
	x, err := ExprUntyped(e.X, vars)
	if err != nil {
		return nil, err
	}

	// Dereference a if it is a pointer
	xV := reflect.ValueOf(x)
	xK := xV.Kind()
	if xK == reflect.Ptr {
		xV = xV.Elem()
	}

	// Extract field
	fieldName := e.Sel.Name
	if xV.Kind() != reflect.Struct {
		return nil, errors.New("unable to get field '" + fieldName + "' from " + xK.String())
	}
	field := xV.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, errors.New("no such field: '" + fieldName + "' on " + xK.String())
	}
	return field.Interface(), nil
}

func Binary(e *ast.BinaryExpr, vars []Var) (r interface{}, err error) {
	x, err := ExprUntyped(e.X, vars)
	if err != nil {
		return nil, err
	}
	y, err := ExprUntyped(e.Y, vars)
	if err != nil {
		return nil, err
	}

	// Perform calc depending on operation type
	if tokenh.IsComparison(e.Op) {
		return binaryCompare(x, e.Op, y)
	} else if tokenh.IsShift(e.Op) {
		return binaryShift(x, e.Op, y)
	} else {
		return binaryOther(x, e.Op, y)
	}
}

func BasicLit(e *ast.BasicLit, vars []Var) (r interface{}, err error) {
	return constant.MakeFromLiteral(e.Value, e.Kind, 0), nil
}

func Paren(e *ast.ParenExpr, vars []Var) (r interface{}, err error) {
	return ExprUntyped(e.X, vars)
}

func Call(e *ast.CallExpr, vars []Var) (r interface{}, err error) {
	// Resolve func
	f, err := ExprUntyped(e.Fun, vars)
	if err != nil {
		return nil, err
	}
	fV := reflect.ValueOf(f)
	if fV.Kind() != reflect.Func {
		return nil, errors.New("no such function " + fV.String())
	}

	// Check in/out arguments count
	fT := fV.Type()
	if fT.NumIn() != len(e.Args) {
		return nil, errors.New("required " + strconvh.FormatInt(fT.NumIn()) + " but got " + strconvh.FormatInt(len(e.Args)))
	}
	if fT.NumOut() != 1 {
		return nil, errors.New("function must return exactly 1 parameter but returns " + strconvh.FormatInt(fT.NumOut()))
	}

	// Prepare arguments
	args := make([]reflect.Value, fT.NumIn())
	for i := range e.Args {
		var arg interface{}
		arg, err = ExprUntyped(e.Args[i], vars)
		if err != nil {
			return nil, err
		}

		// Resolve constant
		if c, ok := arg.(constant.Value); ok {
			arg, ok = constanth.SameTypeInterface(c, fT.In(i))
			if !ok {
				return nil, errors.New("cannot convert argument " + strconvh.FormatInt(i) + " " + c.String() + " (untyped constant) to required type " + fT.In(i).String())
			}
		}

		args[i] = reflect.ValueOf(arg)
	}

	defer func() {
		if rec := recover(); rec != nil {
			if str, ok := rec.(string); ok {
				err = errors.New(str)
			} else {
				err = errors.New(fmt.Sprint(rec))
			}
		}
	}()
	rs := fV.Call(args)
	return rs[0], nil
}

func Star(e *ast.StarExpr, vars []Var) (r interface{}, err error) {
	// TODO what if * means not dereference, but part of type???
	v, err := ExprUntyped(e.X, vars)
	if err != nil {
		return nil, err
	}
	vV := reflect.ValueOf(v)
	if kind := vV.Kind(); kind != reflect.Ptr {
		return reflect.Value{}, errors.New("unable to dereference " + kind.String())
	}
	return vV.Elem().Interface(), nil
}

func Unary(e *ast.UnaryExpr, vars []Var) (r interface{}, err error) {
	x, err := ExprUntyped(e.X, vars)
	if err != nil {
		return nil, err
	}
	return unary(x, e.Op)
}

func Index(e *ast.IndexExpr, vars []Var) (r interface{}, err error) {
	x, err := ExprUntyped(e.X, vars)
	if err != nil {
		return nil, err
	}

	i, err := ExprUntyped(e.Index, vars)
	if err != nil {
		return nil, err
	}

	if k := reflect.ValueOf(x).Kind(); k == reflect.Map {
		return indexMap(x, i)
	}
	return indexOther(x, i)
}

func ExprUntyped(e ast.Expr, vars []Var) (r interface{}, err error) {
	switch v := e.(type) {
	case *ast.Ident:
		return Ident(v, vars)
	case *ast.SelectorExpr:
		return Selector(v, vars)
	case *ast.BinaryExpr:
		return Binary(v, vars)
	case *ast.BasicLit:
		return BasicLit(v, vars)
	case *ast.ParenExpr:
		return Paren(v, vars)
	case *ast.CallExpr:
		return Call(v, vars)
	case *ast.StarExpr:
		return Star(v, vars)
	case *ast.UnaryExpr:
		return Unary(v, vars)
	case *ast.IndexExpr:
		return Index(v, vars)
	default:
		return nil, errors.New("expression evaluation does not support " + reflect.TypeOf(e).String())
	}
}

func Expr(e ast.Expr, vars []Var) (r interface{}, err error) {
	r, err = ExprUntyped(e, vars)
	if err != nil {
		return
	}

	// Resolve constant
	if c, ok := r.(constant.Value); ok {
		r, ok = constanth.DefaultType(c)
		if !ok {
			return nil, errors.New("unable to use default type for untyped constant " + c.String())
		}
	}

	return
}

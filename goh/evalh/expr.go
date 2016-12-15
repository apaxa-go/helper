package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/tokenh"
	"go/ast"
	"go/constant"
	"reflect"
)

type Var struct {
	Name  string
	Value interface{}
}

// r is a typed variable, untyped bool constant or nil
func Ident(e ast.Ident, vars []Var) (r interface{}, err error) {
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
func Selector(e ast.SelectorExpr, vars []Var) (r interface{}, err error) {
	// Calc object (left of '.')
	x, err := Expr(e.X, vars)
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
		return nil, reflect.Value{}, errors.New("unable to get field '" + fieldName + "' from " + xK.String())
	}
	field := xV.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, reflect.Value{}, errors.New("no such field: '" + fieldName + "' on " + xK.String())
	}
	return field.Interface(), nil
}

func Binary(e ast.BinaryExpr, vars []Var) (r interface{}, err error) {
	x, err := Expr(e.X, vars)
	if err != nil {
		return nil, err
	}
	y, err := Expr(e.Y, vars)
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

func BasicLit(e ast.BasicLit, vars []Var) (r interface{}, err error) {
	return constant.MakeFromLiteral(e.Value, e.Kind, 0), nil
}

func Paren(e ast.ParenExpr, vars []Var) (r interface{}, err error) {
	return Expr(e.X, vars)
}

func Call(e ast.CallExpr, vars []Var)(r interface{},err error){
	switch e.Fun. {
	
	}
}

func Expr(e ast.Expr, vars []Var) (r interface{}, err error) {
	switch v := e.(type) {
	case ast.Ident:
		return Ident(v, vars)
	case ast.SelectorExpr:
		return Selector(v, vars)
	case ast.BinaryExpr:
		return Binary(v, vars)
	case ast.BasicLit:
		return BasicLit(v, vars)
	case ast.ParenExpr:
		return Paren(v, vars)
	default:
		return nil, errors.New("expression evaluation does not support " + reflect.TypeOf(e).String())
	}
}

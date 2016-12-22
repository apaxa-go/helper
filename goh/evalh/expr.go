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

type (
	Identifiers          map[string]Value
	IdentifiersRegular   map[string]reflect.Value
	IdentifiersInterface map[string]interface{}
)

func (idents IdentifiersRegular) Identifiers() (r Identifiers) {
	r = make(Identifiers)
	for i := range idents {
		r[i] = MakeRegular(idents[i])
	}
	return
}
func (idents IdentifiersInterface) IdentifiersRegular() (r IdentifiersRegular) {
	r = make(IdentifiersRegular)
	for i := range idents {
		r[i] = reflect.ValueOf(idents[i])
	}
	return
}
func (idents IdentifiersInterface) Identifiers() (r Identifiers) {
	r = make(Identifiers)
	for i := range idents {
		r[i] = MakeRegular(reflect.ValueOf(idents[i]))
	}
	return
}

// r is a typed variable, untyped bool constant or nil
func Ident(e *ast.Ident, idents Identifiers) (r Value, err error) {
	// true and false literals does not covered by BasicLit (because it is look like normal identifier?).
	if e.Name == "true" {
		return MakeUntyped(constant.MakeBool(true)), nil
	}
	if e.Name == "false" {
		return MakeUntyped(constant.MakeBool(false)), nil
	}

	// nil literal
	if e.Name == "nil" {
		return MakeNil(), nil
	}

	// find required identifier in vars.
	r, ok := idents[e.Name]
	if !ok {
		err = errors.New("variable '" + e.Name + "' does not defined")
	}
	return
}

// Selector can:
// 	* get field from struct or pointer to struct
//	* get method (defined with receiver V) from variable of type V or pointer variable to type V
//	* get method (defined with pointer receiver V) from pointer variable to type V
func Selector(e *ast.SelectorExpr, idents Identifiers) (r Value, err error) {
	// TODO implement "pkg.Method(...)"
	// Calc object (left of '.')
	x, err := Expr(e.X, idents)
	if err != nil {
		return nil, err
	}
	if x.Kind() != Regular {
		return nil, errors.New("unable to select field or method on untyped constant")
	}

	// Extract field/method name
	name := e.Sel.Name

	////////////////
	xV := x.Regular()
	xK := xV.Kind()

	// If kind is pointer than try to get method.
	// If no method can be get than dereference pointer.
	if xK == reflect.Ptr {
		if method := xV.MethodByName(name); method.IsValid() {
			return MakeRegular(method), nil
		}
		xV = xV.Elem()
	}

	// If kind is struct than try to get field
	if xV.Kind() == reflect.Struct {
		if field := xV.FieldByName(name); field.IsValid() {
			return MakeRegular(field), nil
		}
	}

	// Last case - try to get method on dereferenced variable
	if method := xV.MethodByName(name); method.IsValid() {
		return MakeRegular(method), nil
	}

	return nil, errors.New("no such field or method: '" + name + "' on " + xK.String())
}

func Binary(e *ast.BinaryExpr, idents Identifiers) (r Value, err error) {
	x, err := Expr(e.X, idents)
	if err != nil {
		return nil, err
	}
	y, err := Expr(e.Y, idents)
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

func BasicLit(e *ast.BasicLit, idents Identifiers) (r Value, err error) {
	return MakeUntyped(constant.MakeFromLiteral(e.Value, e.Kind, 0)), nil
}

func Paren(e *ast.ParenExpr, idents Identifiers) (r Value, err error) {
	return Expr(e.X, idents)
}

// ok means that CallExpr is built-in function, not that calling done without error.
func tryCallBuiltIn(e *ast.CallExpr, idents Identifiers) (r Value, ok bool, err error) {
	fIdent, ok := e.Fun.(*ast.Ident)
	if !ok {
		return
	}

	ok = true
	switch fIdent.Name {
	case "len":
		if len(e.Args) != 1 {
			err = errors.New("tcbi0") // TODO
			return
		}
		var arg0 Value
		if arg0, err = Expr(e.Args[0], idents); err != nil {
			return
		}
		r, err = BuiltInLen(arg0)
	case "cap":
		if len(e.Args) != 1 {
			err = errors.New("tcbi1") // TODO
			return
		}
		var arg0 Value
		if arg0, err = Expr(e.Args[0], idents); err != nil {
			return
		}
		r, err = BuiltInCap(arg0)
	case "complex":
		if len(e.Args) != 2 {
			err = errors.New("tcbi2") // TODO
			return
		}
		var arg0, arg1 Value
		if arg0, err = Expr(e.Args[0], idents); err != nil {
			return
		}
		if arg1, err = Expr(e.Args[1], idents); err != nil {
			return
		}
		r, err = BuiltInComplex(arg0, arg1)
	case "real":
		if len(e.Args) != 1 {
			err = errors.New("tcbi3") // TODO
			return
		}
		var arg0 Value
		if arg0, err = Expr(e.Args[0], idents); err != nil {
			return
		}
		r, err = BuiltInReal(arg0)
	case "imag":
		if len(e.Args) != 1 {
			err = errors.New("tcbi4") // TODO
			return
		}
		var arg0 Value
		if arg0, err = Expr(e.Args[0], idents); err != nil {
			return
		}
		r, err = BuiltInImag(arg0)
	default:
		ok = false
	}

	return
}

func Call(e *ast.CallExpr, idents Identifiers) (r Value, err error) {
	// Try built-in
	var ok bool
	r, ok, err = tryCallBuiltIn(e, idents)
	if ok {
		return
	}

	// Resolve func
	f, err := Expr(e.Fun, idents)
	if err != nil {
		return nil, err
	}
	if f.Kind() != Regular || f.Regular().Kind() != reflect.Func {
		return nil, errors.New("no such function " + f.String())
	}

	// Check in/out arguments count
	fT := f.Regular().Type()
	if fT.NumIn() != len(e.Args) {
		return nil, errors.New("required " + strconvh.FormatInt(fT.NumIn()) + " but got " + strconvh.FormatInt(len(e.Args)))
	}
	if fT.NumOut() != 1 {
		return nil, errors.New("function must return exactly 1 parameter but returns " + strconvh.FormatInt(fT.NumOut()))
	}

	// Prepare arguments
	args := make([]reflect.Value, fT.NumIn())
	for i := range e.Args {
		var arg Value
		arg, err = Expr(e.Args[i], idents)
		if err != nil {
			return nil, err
		}

		switch arg.Kind() {
		case Untyped: // TODO what if argument may be untyped because function is built-in?
			var ok bool
			args[i], ok = constanth.SameType(arg.Untyped(), fT.In(i))
			if !ok {
				return nil, errors.New("cannot convert argument " + strconvh.FormatInt(i) + " " + arg.String() + " (untyped constant) to required type " + fT.In(i).String())
			}
		case Nil:
			args[i] = reflect.ValueOf(nil) // TODO check this, may be reflect.ValueOf(nil)
		case Regular:
			if aT := arg.Regular().Type(); !aT.AssignableTo(fT.In(i)) {
				return nil, errors.New("cannot convert argument " + strconvh.FormatInt(i) + " " + aT.String() + " to required type " + fT.In(i).String())
			}
			args[i] = arg.Regular()
		default:
			panic("unknown argument kind")
		}
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
	rs := f.Regular().Call(args)
	return MakeRegular(rs[0]), nil // TODO what if result is untyped because function is built-in?
}

func Star(e *ast.StarExpr, idents Identifiers) (r Value, err error) {
	// TODO what if * means not dereference, but part of type???
	v, err := Expr(e.X, idents)
	if err != nil {
		return nil, err
	}
	if v.Kind() != Regular || v.Regular().Kind() != reflect.Ptr {
		return nil, errors.New("unable to dereference " + v.String())
	}
	return MakeRegular(v.Regular().Elem()), nil
}

func Unary(e *ast.UnaryExpr, idents Identifiers) (r Value, err error) {
	x, err := Expr(e.X, idents)
	if err != nil {
		return nil, err
	}
	return unary(x, e.Op)
}

func Index(e *ast.IndexExpr, idents Identifiers) (r Value, err error) {
	x, err := Expr(e.X, idents)
	if err != nil {
		return nil, err
	}

	i, err := Expr(e.Index, idents)
	if err != nil {
		return nil, err
	}

	if x.Kind() == Untyped {
		return indexConstant(x.Untyped(), i)
	}
	if x.Kind() != Regular {
		return nil, errors.New("unable to index " + x.String())
	}

	if x.Regular().Kind() == reflect.Map {
		return indexMap(x.Regular(), i)
	}
	return indexOther(x.Regular(), i)
}

func Slice(e *ast.SliceExpr, idents Identifiers) (r Value, err error) {
	x, err := Expr(e.X, idents)
	if err != nil {
		return nil, err
	}

	// Calc indexes
	low, err := getSliceIndex(e.Low, idents)
	if err != nil {
		return nil, err
	}
	high, err := getSliceIndex(e.High, idents)
	if err != nil {
		return nil, err
	}
	var max int
	if e.Slice3 {
		max, err = getSliceIndex(e.Max, idents)
		if err != nil {
			return nil, err
		}
	}

	var v reflect.Value
	switch x.Kind() {
	case Untyped:
		if x.Untyped().Kind() != constant.String {
			return nil, errors.New("unable to indexing " + x.String())
		}
		v = reflect.ValueOf(x.Untyped().String())
	case Regular:
		v = x.Regular()
	case Nil:
		return nil, errors.New("unable to slicing Nil")
	default:
		panic("unknown kind")
	}

	if e.Slice3 {
		slice3(v, low, high, max)
	}
	return slice2(v, low, high)
}

func Expr(e ast.Expr, idents Identifiers) (r Value, err error) {
	switch v := e.(type) {
	case *ast.Ident:
		return Ident(v, idents)
	case *ast.SelectorExpr:
		return Selector(v, idents)
	case *ast.BinaryExpr:
		return Binary(v, idents)
	case *ast.BasicLit:
		return BasicLit(v, idents)
	case *ast.ParenExpr:
		return Paren(v, idents)
	case *ast.CallExpr:
		return Call(v, idents)
	case *ast.StarExpr:
		return Star(v, idents)
	case *ast.UnaryExpr:
		return Unary(v, idents)
	case *ast.IndexExpr:
		return Index(v, idents)
	case *ast.SliceExpr:
		return Slice(v, idents)
	default:
		return nil, errors.New("expression evaluation does not support " + reflect.TypeOf(e).String())
	}
}

func ExprRegular(e ast.Expr, idents IdentifiersRegular) (r reflect.Value, err error) {
	rV, err := Expr(e, idents.Identifiers())
	if err != nil {
		return
	}

	switch rV.Kind() {
	case Regular:
		r = rV.Regular()
	case Untyped:
		var ok bool
		r, ok = constanth.DefaultType(rV.Untyped())
		if !ok {
			return r, errors.New("unable to represent untyped value in default type")
		}
	case Nil:
		r = reflect.ValueOf(nil) // TODO is it normal?
	default:
		panic("unknown kind")
	}
	return
}

func ExprInterface(e ast.Expr, idents IdentifiersInterface) (r interface{}, err error) {
	rV, err := ExprRegular(e, idents.IdentifiersRegular())
	if err != nil {
		return
	}

	if !rV.CanInterface() {
		return r, errors.New("result value can not be converted into interface")
	}

	return rV.Interface(), nil
}

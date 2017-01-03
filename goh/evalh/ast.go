package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/goh/tokenh"
	"github.com/apaxa-go/helper/reflecth"
	"go/ast"
	"go/constant"
	"reflect"
)

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

	if isBuiltInFunc(e.Name) {
		return MakeBuiltInFunc(e.Name), nil
	}

	if isBuiltInType(e.Name) {
		return MakeType(builtInTypes[e.Name]), nil
	}

	// find required identifier in vars.
	r, ok := idents[e.Name]
	if !ok {
		err = errors.New("variable '" + e.Name + "' does not defined")
	}
	return
}

// SelectorExpr can:
// 	* get field from struct or pointer to struct
//	* get method (defined with receiver V) from variable of type V or pointer variable to type V
//	* get method (defined with pointer receiver V) from pointer variable to type V
func SelectorExpr(e *ast.SelectorExpr, idents Identifiers) (r Value, err error) {
	// Calc object (left of '.')
	x, err := expr(e.X, idents)
	if err != nil {
		return nil, err
	}

	// Extract field/method name
	name := e.Sel.Name

	switch x.Kind() {
	case Package:
		var ok bool
		r, ok = x.Package()[name]
		if !ok {
			return nil, errors.New("package has no method/variable/type " + name)
		}

		return
	case Regular:
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

		// Last case - try to get method (on already dereferenced variable)
		if method := xV.MethodByName(name); method.IsValid() {
			return MakeRegular(method), nil
		}

		return nil, errors.New("no such field or method: '" + name + "' on " + xK.String())
	default:
		return nil, errors.New("required Package or Regular")
	}
}

func BinaryExpr(e *ast.BinaryExpr, idents Identifiers) (r Value, err error) {
	x, err := expr(e.X, idents)
	if err != nil {
		return nil, err
	}
	y, err := expr(e.Y, idents)
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

func ParenExpr(e *ast.ParenExpr, idents Identifiers) (r Value, err error) {
	if e.X == nil { // TODO add this check to other expressions
		return nil, errors.New("empty ParenExpr")
	}
	return expr(e.X, idents)
}

func CallExpr(e *ast.CallExpr, idents Identifiers) (r Value, err error) {
	// Resolve func
	f, err := expr(e.Fun, idents)
	if err != nil {
		return nil, err
	}

	// Resolve args
	args := make([]Value, len(e.Args))
	for i := range e.Args {
		args[i], err = expr(e.Args[i], idents)
		if err != nil {
			return
		}
	}

	switch f.Kind() {
	case Regular:
		return callRegular(f.Regular(), args)
	case BuiltInFunc:
		return callBuiltInFunc(f.BuiltInFunc(), args)
	case Type:
		return callType(f.Type(), args)
	default:
		return nil, errors.New("unable to call on " + f.String())
	}
}

func StarExpr(e *ast.StarExpr, idents Identifiers) (r Value, err error) {
	v, err := expr(e.X, idents)
	if err != nil {
		return nil, err
	}
	switch {
	case v.Kind() == Type:
		return MakeType(reflect.PtrTo(v.Type())), nil
	case v.Kind() == Regular && v.Regular().Kind() == reflect.Ptr:
		return MakeRegular(v.Regular().Elem()), nil
	default:
		return nil, errors.New("unable to dereference " + v.String())
	}
}

func UnaryExpr(e *ast.UnaryExpr, idents Identifiers) (r Value, err error) {
	x, err := expr(e.X, idents)
	if err != nil {
		return nil, err
	}
	return unary(x, e.Op)
}

func ChanType(e *ast.ChanType, idents Identifiers) (r Value, err error) {
	v, err := expr(e.Value, idents)
	if err != nil {
		return nil, err
	}
	switch v.Kind() {
	case Type:
		return MakeType(reflect.ChanOf(reflecth.ChanDirFromAst(e.Dir), v.Type())), nil
	default:
		return nil, errors.New("") // TODO
	}
}

// TODO works only for list of arguments types (not value, not array size)
func Ellipsis(e *ast.Ellipsis, idents Identifiers) (r Value, err error) {
	v, err := expr(e.Elt, idents)
	if err != nil {
		return
	}

	if v.Kind() != Type {
		return nil, errors.New("ellipsis aplies only to Type")
	}
	return MakeType(reflect.SliceOf(v.Type())), nil
}

func FuncType(e *ast.FuncType, idents Identifiers) (r Value, err error) {
	in, variadic, err := funcTranslateArgs(e.Params, true, idents)
	if err != nil {
		return
	}
	out, _, err := funcTranslateArgs(e.Results, false, idents)
	if err != nil {
		return
	}
	return MakeType(reflect.FuncOf(in, out, variadic)), nil
}

func IndexExpr(e *ast.IndexExpr, idents Identifiers) (r Value, err error) {
	x, err := expr(e.X, idents)
	if err != nil {
		return nil, err
	}

	i, err := expr(e.Index, idents)
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

func SliceExpr(e *ast.SliceExpr, idents Identifiers) (r Value, err error) {
	x, err := expr(e.X, idents)
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

func CompositeLit(e *ast.CompositeLit, idents Identifiers) (r Value, err error) {
	v, err := expr(e.Type, idents) // TODO docs says: e.Type may be nil
	if err != nil {
		return
	}

	if v.Kind() != Type {
		return nil, errors.New("composite literal requires type, but got " + v.String())
	}

	vT := v.Type()
	switch vT.Kind() {
	case reflect.Struct:
		var withKeys bool
		if len(e.Elts) == 0 {
			withKeys = true // Treat empty initialization list as with keys
		} else {
			_, withKeys = e.Elts[0].(*ast.KeyValueExpr)
		}

		switch withKeys {
		case true:
			elts := make(map[string]Value)
			for i := range e.Elts {
				kve, ok := e.Elts[i].(*ast.KeyValueExpr)
				if !ok {
					return nil, errors.New("mixed keys & no keys initialization in composite literal")
				}

				key, ok := kve.Key.(*ast.Ident)
				if !ok {
					return nil, errors.New("keys in struct composite literal must be identifier")
				}

				elts[key.Name], err = expr(kve.Value, idents)
				if err != nil {
					return
				}
			}
			return compositeLitStructKeys(vT, elts)
		case false:
			elts := make([]Value, len(e.Elts))
			for i := range e.Elts {
				if _, ok := e.Elts[i].(*ast.KeyValueExpr); ok {
					return nil, errors.New("mixed keys & no keys initialization in composite literal")
				}

				elts[i], err = expr(e.Elts[i], idents)
				if err != nil {
					return
				}
			}
			return compositeLitStructOrdered(vT, elts)
		default:
			return nil, nil // Unreachable
		}
	case reflect.Array, reflect.Slice:
		elts := make(map[int]Value)
		nextIndex := 0
		for i := range e.Elts {
			var valueExpr ast.Expr
			if kve, ok := e.Elts[i].(*ast.KeyValueExpr); ok {
				var v Value
				v, err = expr(kve.Key, idents)
				if err != nil {
					return
				}
				if v.Kind() != Untyped {
					return nil, errors.New("key in array initialization must be constant")
				}

				if nextIndex, ok = constanth.IntVal(v.Untyped()); !ok || nextIndex < 0 {
					return nil, errors.New("key in array initialization must be constant not negative integer")
				}

				valueExpr = kve.Value
			} else {
				valueExpr = e.Elts[i]
			}

			if _, ok := elts[nextIndex]; ok {
				return nil, errors.New("duplicate key in array initialization")
			}

			elts[nextIndex], err = expr(valueExpr, idents)
			if err != nil {
				return
			}
		}

		return compositeLitArrayLike(vT, elts)
	case reflect.Map:
		elts := make(map[Value]Value)
		for i := range e.Elts {
			kve, ok := e.Elts[i].(*ast.KeyValueExpr)
			if !ok {
				return nil, errors.New("map initialization must have keys")
			}

			var key Value
			key, err = expr(kve.Key, idents)
			if err != nil {
				return
			}
			elts[key], err = expr(kve.Value, idents) // TODO looks like it is impossible to overwrite value here because key!=prev_key
			if err != nil {
				return
			}
		}

		return compositeLitMap(vT, elts)
	default:
		return nil, errors.New("composite literal can be used only for structs, arrays, slices & maps, not for " + vT.Kind().String())
	}
}

func expr(e ast.Expr, idents Identifiers) (r Value, err error) {
	switch v := e.(type) {
	case *ast.Ident:
		return Ident(v, idents)
	case *ast.SelectorExpr:
		return SelectorExpr(v, idents)
	case *ast.BinaryExpr:
		return BinaryExpr(v, idents)
	case *ast.BasicLit:
		return BasicLit(v, idents)
	case *ast.ParenExpr:
		return ParenExpr(v, idents)
	case *ast.CallExpr:
		return CallExpr(v, idents)
	case *ast.StarExpr:
		return StarExpr(v, idents)
	case *ast.UnaryExpr:
		return UnaryExpr(v, idents)
	case *ast.Ellipsis:
		return Ellipsis(v, idents)
	case *ast.ChanType:
		return ChanType(v, idents)
	case *ast.FuncType:
		return FuncType(v, idents)
	case *ast.IndexExpr:
		return IndexExpr(v, idents)
	case *ast.SliceExpr:
		return SliceExpr(v, idents)
	case *ast.CompositeLit:
		return CompositeLit(v, idents)
	default:
		return nil, errors.New("expression evaluation does not support " + reflect.TypeOf(e).String())
	}
}

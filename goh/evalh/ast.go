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

func astIdent(e *ast.Ident, idents Identifiers) (r Value, err error) {
	switch e.Name {
	case "true":
		return MakeUntyped(constant.MakeBool(true)), nil
	case "false":
		return MakeUntyped(constant.MakeBool(false)), nil
	case "nil":
		return MakeNil(), nil
	}

	switch {
	case isBuiltInFunc(e.Name):
		return MakeBuiltInFunc(e.Name), nil
	case isBuiltInType(e.Name):
		return MakeType(builtInTypes[e.Name]), nil
	default:
		var ok bool
		r, ok = idents[e.Name]
		if !ok {
			err = errors.New("variable '" + e.Name + "' does not defined") // TODO error
		}
		return
	}
}

// astSelectorExpr can:
// 	* get field from struct or pointer to struct
//	* get method (defined with receiver V) from variable of type V or pointer variable to type V
//	* get method (defined with pointer receiver V) from pointer variable to type V
func astSelectorExpr(e *ast.SelectorExpr, idents Identifiers) (r Value, err error) {
	// Calc object (left of '.')
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	// Extract field/method name
	if e.Sel == nil {
		return nil, errors.New("try to select nil method/field/type/variable")
	}
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

		// If kind is pointer than try to get method.
		// If no method can be get than dereference pointer.
		if xV.Kind() == reflect.Ptr {
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

		return nil, errors.New("no such field or method: '" + name + "' on " + xV.Kind().String())
	default:
		return nil, errors.New("required Package or Regular")
	}
}

func astBinaryExpr(e *ast.BinaryExpr, idents Identifiers) (r Value, err error) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}
	y, err := astExpr(e.Y, idents)
	if err != nil {
		return
	}

	// Perform calc depending on operation type
	switch {
	case tokenh.IsComparison(e.Op):
		return binaryCompare(x, e.Op, y)
	case tokenh.IsShift(e.Op):
		return binaryShift(x, e.Op, y)
	default:
		return binaryOther(x, e.Op, y)
	}
}

func astBasicLit(e *ast.BasicLit, idents Identifiers) (r Value, err error) {
	rC := constant.MakeFromLiteral(e.Value, e.Kind, 0)
	if rC.Kind() == constant.Unknown {
		return nil, errors.New("invalid basic literal syntax")
	}
	return MakeUntyped(rC), nil
}

func astParenExpr(e *ast.ParenExpr, idents Identifiers) (r Value, err error) {
	return astExpr(e.X, idents)
}

func astCallExpr(e *ast.CallExpr, idents Identifiers) (r Value, err error) {
	// Resolve func
	f, err := astExpr(e.Fun, idents)
	if err != nil {
		return
	}

	// Resolve args
	args := make([]Value, len(e.Args))
	for i := range e.Args {
		args[i], err = astExpr(e.Args[i], idents)
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
		return convertCall(f.Type(), args)
	default:
		return nil, errors.New("unable to call on " + f.String())
	}
}

func astStarExpr(e *ast.StarExpr, idents Identifiers) (r Value, err error) {
	v, err := astExpr(e.X, idents)
	if err != nil {
		return
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

func astUnaryExpr(e *ast.UnaryExpr, idents Identifiers) (r Value, err error) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	switch x.Kind() {
	case Regular:
		return unary(x.Regular(), e.Op)
	case Untyped:
		return unaryConstant(x.Untyped(), e.Op)
	default:
		return nil, errors.New("Regular or Untyped required")
	}
}

func astChanType(e *ast.ChanType, idents Identifiers) (r Value, err error) {
	v, err := astExpr(e.Value, idents)
	if err != nil {
		return
	}

	switch v.Kind() {
	case Type:
		return MakeType(reflect.ChanOf(reflecth.ChanDirFromAst(e.Dir), v.Type())), nil
	default:
		return nil, errors.New("") // TODO error
	}
}

// TODO works only for list of arguments types (not value, not array size)
func astEllipsis(e *ast.Ellipsis, idents Identifiers) (r Value, err error) {
	v, err := astExpr(e.Elt, idents)
	if err != nil {
		return
	}

	if v.Kind() != Type {
		return nil, errors.New("ellipsis aplies only to Type")
	}

	return MakeType(reflect.SliceOf(v.Type())), nil
}

func astFuncType(e *ast.FuncType, idents Identifiers) (r Value, err error) {
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

func astIndexExpr(e *ast.IndexExpr, idents Identifiers) (r Value, err error) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	i, err := astExpr(e.Index, idents)
	if err != nil {
		return nil, err
	}

	switch x.Kind() {
	case Untyped:
		return indexConstant(x.Untyped(), i)
	case Regular:
		switch x.Regular().Kind() {
		case reflect.Map:
			return indexMap(x.Regular(), i)
		default:
			return indexOther(x.Regular(), i)
		}
	default:
		return nil, errors.New("unable to index " + x.String())
	}
}

func astSliceExpr(e *ast.SliceExpr, idents Identifiers) (r Value, err error) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	// Calc indexes
	low, err := getSliceIndex(e.Low, idents)
	if err != nil {
		return
	}
	high, err := getSliceIndex(e.High, idents)
	if err != nil {
		return
	}
	var max int
	if e.Slice3 {
		max, err = getSliceIndex(e.Max, idents)
		if err != nil {
			return
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
	default:
		return nil, errors.New("Untyped or Regular required")
	}

	if e.Slice3 {
		slice3(v, low, high, max)
	}
	return slice2(v, low, high)
}

func astCompositeLit(e *ast.CompositeLit, idents Identifiers) (r Value, err error) {
	v, err := astExpr(e.Type, idents)
	if err != nil {
		return
	}

	if v.Kind() != Type {
		return nil, errors.New("composite literal requires type, but got " + v.String())
	}

	switch vT := v.Type(); vT.Kind() {
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

				elts[key.Name], err = astExpr(kve.Value, idents)
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

				elts[i], err = astExpr(e.Elts[i], idents)
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
				v, err = astExpr(kve.Key, idents)
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

			elts[nextIndex], err = astExpr(valueExpr, idents)
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
			key, err = astExpr(kve.Key, idents)
			if err != nil {
				return
			}
			elts[key], err = astExpr(kve.Value, idents) // looks like it is impossible to overwrite value here because key!=prev_key (it is interface)
			if err != nil {
				return
			}
		}

		return compositeLitMap(vT, elts)
	default:
		return nil, errors.New("composite literal can be used only for structs, arrays, slices & maps, not for " + vT.Kind().String())
	}
}

func astExpr(e ast.Expr, idents Identifiers) (r Value, err error) {
	if e == nil {
		return nil, errors.New("try to eval nil Expr")
	}

	switch v := e.(type) {
	case *ast.Ident:
		return astIdent(v, idents)
	case *ast.SelectorExpr:
		return astSelectorExpr(v, idents)
	case *ast.BinaryExpr:
		return astBinaryExpr(v, idents)
	case *ast.BasicLit:
		return astBasicLit(v, idents)
	case *ast.ParenExpr:
		return astParenExpr(v, idents)
	case *ast.CallExpr:
		return astCallExpr(v, idents)
	case *ast.StarExpr:
		return astStarExpr(v, idents)
	case *ast.UnaryExpr:
		return astUnaryExpr(v, idents)
	case *ast.Ellipsis:
		return astEllipsis(v, idents)
	case *ast.ChanType:
		return astChanType(v, idents)
	case *ast.FuncType:
		return astFuncType(v, idents)
	case *ast.IndexExpr:
		return astIndexExpr(v, idents)
	case *ast.SliceExpr:
		return astSliceExpr(v, idents)
	case *ast.CompositeLit:
		return astCompositeLit(v, idents)
	default:
		return nil, errors.New("expression evaluation does not support " + reflect.TypeOf(e).String())
	}
}

package evalh

import (
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/goh/tokenh"
	"github.com/apaxa-go/helper/reflecth"
	"go/ast"
	"go/constant"
	"go/token"
	"reflect"
)

func astIdent(e *ast.Ident, idents Identifiers) (r Value, err *posError) {
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
			err = identUndefinedError(e.Name).pos(e)
		}
		return
	}
}

// astSelectorExpr can:
// 	* get field from struct or pointer to struct
//	* get method (defined with receiver V) from variable of type V or pointer variable to type V
//	* get method (defined with pointer receiver V) from pointer variable to type V
func astSelectorExpr(e *ast.SelectorExpr, idents Identifiers) (r Value, err *posError) {
	// Calc object (left of '.')
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	// Extract field/method name
	if e.Sel == nil {
		return nil, invAstSelectorError().pos(e)
	}
	name := e.Sel.Name

	switch x.Kind() {
	case Package:
		var ok bool
		r, ok = x.Package()[name]
		if !ok {
			return nil, identUndefinedError("." + name).pos(e)
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

		return nil, identUndefinedError("." + name).pos(e)
	case Type:
		xT := x.Type()

		if xT.Kind() == reflect.Interface {
			return nil, newIntError("Method expressions for interface types currently does not supported").pos(e) // BUG
		}

		f, ok := xT.MethodByName(name)
		if !ok || !f.Func.IsValid() {
			return nil, selectorUndefIdentError(xT, name).pos(e)
		}
		return MakeRegular(f.Func), nil
	default:
		return nil, invSelectorXError(x).pos(e)
	}
}

func astBinaryExpr(e *ast.BinaryExpr, idents Identifiers) (r Value, err *posError) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}
	y, err := astExpr(e.Y, idents)
	if err != nil {
		return
	}

	// Perform calc depending on operation type
	var intErr *intError
	switch {
	case tokenh.IsComparison(e.Op):
		r, intErr = binaryCompare(x, e.Op, y)
	case tokenh.IsShift(e.Op):
		r, intErr = binaryShift(x, e.Op, y)
	default:
		r, intErr = binaryOther(x, e.Op, y)
	}

	err = intErr.pos(e)
	return
}

func astBasicLit(e *ast.BasicLit, idents Identifiers) (r Value, err *posError) {
	rC := constant.MakeFromLiteral(e.Value, e.Kind, 0)
	if rC.Kind() == constant.Unknown {
		return nil, syntaxInvBasLitError(e.Value).pos(e)
	}
	return MakeUntyped(rC), nil
}

func astParenExpr(e *ast.ParenExpr, idents Identifiers) (r Value, err *posError) {
	return astExpr(e.X, idents)
}

func astCallExpr(e *ast.CallExpr, idents Identifiers) (r Value, err *posError) {
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

	var intErr *intError
	switch f.Kind() {
	case Regular:
		r, intErr = callRegular(f.Regular(), args, e.Ellipsis != token.NoPos)
	case BuiltInFunc:
		r, intErr = callBuiltInFunc(f.BuiltInFunc(), args, e.Ellipsis != token.NoPos)
	case Type:
		if e.Ellipsis != token.NoPos {
			return nil, convertWithEllipsisError(f.Type()).pos(e)
		}
		r, intErr = convertCall(f.Type(), args)
	default:
		intErr = callNonFuncError(f)
	}

	err = intErr.pos(e)
	return
}

func astStarExpr(e *ast.StarExpr, idents Identifiers) (r Value, err *posError) {
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
		return nil, indirectInvalError(v).pos(e)
	}
}

func astUnaryExpr(e *ast.UnaryExpr, idents Identifiers) (r Value, err *posError) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	var intErr *intError
	switch x.Kind() {
	case Regular:
		r, intErr = unary(x.Regular(), e.Op)
	case Untyped:
		r, intErr = unaryConstant(x.Untyped(), e.Op)
	default:
		intErr = notExprError(x)
	}

	err = intErr.pos(e)
	return
}

func astChanType(e *ast.ChanType, idents Identifiers) (r Value, err *posError) {
	v, err := astExpr(e.Value, idents)
	if err != nil {
		return
	}

	switch v.Kind() {
	case Type:
		return MakeType(reflect.ChanOf(reflecth.ChanDirFromAst(e.Dir), v.Type())), nil
	default:
		return nil, syntaxMisChanTypeError().pos(e)
	}
}

// Here implements only for list of arguments types ("func(a ...string)").
// For ellipsis array literal ("[...]int{1,2}") see astCompositeLit.
// For ellipsis argument for call ("f(1,a...)") see astCallExpr.
func astEllipsis(e *ast.Ellipsis, idents Identifiers) (r Value, err *posError) {
	v, err := astExpr(e.Elt, idents)
	if err != nil {
		return
	}

	if v.Kind() != Type {
		return nil, syntaxMisVariadicTypeError().pos(e)
	}

	return MakeType(reflect.SliceOf(v.Type())), nil
}

func astFuncType(e *ast.FuncType, idents Identifiers) (r Value, err *posError) {
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

func astArrayType(e *ast.ArrayType, idents Identifiers) (r Value, err *posError) {
	v, err := astExpr(e.Elt, idents)
	if err != nil {
		return
	}

	if v.Kind() != Type {
		return nil, syntaxMisArrayTypeError().pos(e)
	}

	switch e.Len {
	case nil: // Slice
		rT := reflect.SliceOf(v.Type())
		return MakeType(rT), nil
	default: // Array
		var lV Value
		lV, err = astExpr(e.Len, idents) // Case with ellipsis in length must be caught by caller
		if err != nil {
			return
		}
		lInt, _, ok := lV.ConvertToInt() // BUG Not by spec: spec require constant (typed or untyped), we accept any Val wich can be convert to int exactly.
		if !ok || lInt < 0 {
			return nil, arrayBoundNegError().pos(e)
		}
		rT := reflect.ArrayOf(lInt, v.Type())
		return MakeType(rT), nil
	}
}

func astIndexExpr(e *ast.IndexExpr, idents Identifiers) (r Value, err *posError) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	i, err := astExpr(e.Index, idents)
	if err != nil {
		return nil, err
	}

	var intErr *intError
	switch x.Kind() {
	case Untyped:
		r, intErr = indexConstant(x.Untyped(), i)
	case Regular:
		switch x.Regular().Kind() {
		case reflect.Map:
			r, intErr = indexMap(x.Regular(), i)
		default:
			r, intErr = indexOther(x.Regular(), i)
		}
	default:
		intErr = notExprError(x)
	}

	err = intErr.pos(e)
	return
}

func astSliceExpr(e *ast.SliceExpr, idents Identifiers) (r Value, err *posError) {
	x, err := astExpr(e.X, idents)
	if err != nil {
		return
	}

	indexResolve := func(e ast.Expr) (iInt int, err1 *posError) {
		var i Value
		if e != nil {
			i, err1 = astExpr(e, idents)
			if err1 != nil {
				return
			}
		}

		var intErr *intError
		iInt, intErr = getSliceIndex(i)
		err1 = intErr.pos(e)
		return
	}

	// Calc indexes
	low, err := indexResolve(e.Low)
	if err != nil {
		return
	}
	high, err := indexResolve(e.High)
	if err != nil {
		return
	}
	var max int
	if e.Slice3 {
		max, err = indexResolve(e.Max)
		if err != nil {
			return
		}
	}

	var v reflect.Value
	switch x.Kind() {
	case Untyped:
		if x.Untyped().Kind() != constant.String {
			return nil, sliceInvTypeError(x).pos(e.X)
		}
		v = reflect.ValueOf(x.Untyped().String())
	case Regular:
		v = x.Regular()
	default:
		return nil, sliceInvTypeError(x).pos(e.X)
	}

	var intErr *intError
	if e.Slice3 {
		r, intErr = slice3(v, low, high, max)
	} else {
		r, intErr = slice2(v, low, high)
	}

	err = intErr.pos(e)
	return
}

func astCompositeLit(e *ast.CompositeLit, idents Identifiers) (r Value, err *posError) {
	// type
	var vT reflect.Type
	// case where type is an ellipsis array
	if aType, ok := e.Type.(*ast.ArrayType); ok {
		if _, ok := aType.Len.(*ast.Ellipsis); ok {
			// Resolve array elements type
			var v Value
			v, err = astExpr(aType.Elt, idents)
			if v.Kind() != Type {
				return nil, notTypeError(v).pos(aType.Elt)
			}
			vT = reflect.ArrayOf(len(e.Elts), v.Type())
		}
	}
	// other cases
	if vT == nil {
		var v Value
		v, err = astExpr(e.Type, idents)
		if err != nil {
			return
		}

		if v.Kind() != Type {
			return nil, notTypeError(v).pos(e.Type)
		}
		vT = v.Type()
	}

	// Construct
	var intErr *intError
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
					return nil, initMixError().pos(e)
				}

				key, ok := kve.Key.(*ast.Ident)
				if !ok {
					return nil, initStructInvFieldNameError().pos(kve)
				}

				elts[key.Name], err = astExpr(kve.Value, idents)
				if err != nil {
					return
				}
			}
			r, intErr = compositeLitStructKeys(vT, elts)
		case false:
			elts := make([]Value, len(e.Elts))
			for i := range e.Elts {
				if _, ok := e.Elts[i].(*ast.KeyValueExpr); ok {
					return nil, initMixError().pos(e)
				}

				elts[i], err = astExpr(e.Elts[i], idents)
				if err != nil {
					return
				}
			}
			r, intErr = compositeLitStructOrdered(vT, elts)
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
					return nil, initArrayInvIndexError().pos(kve)
				}

				if nextIndex, ok = constanth.IntVal(v.Untyped()); !ok || nextIndex < 0 {
					return nil, initArrayInvIndexError().pos(kve)
				}

				valueExpr = kve.Value
			} else {
				valueExpr = e.Elts[i]
			}

			if _, ok := elts[nextIndex]; ok {
				return nil, initArrayDupIndexError(nextIndex).pos(e.Elts[i])
			}

			elts[nextIndex], err = astExpr(valueExpr, idents)
			if err != nil {
				return
			}
			nextIndex++
		}

		r, intErr = compositeLitArrayLike(vT, elts)
	case reflect.Map:
		elts := make(map[Value]Value)
		for i := range e.Elts {
			kve, ok := e.Elts[i].(*ast.KeyValueExpr)
			if !ok {
				return nil, initMapMisKeyError().pos(e.Elts[i])
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

		r, intErr = compositeLitMap(vT, elts)
	default:
		return nil, initInvTypeError(vT).pos(e.Type)
	}

	err = intErr.pos(e)
	return
}

func astExpr(e ast.Expr, idents Identifiers) (r Value, err *posError) {
	if e == nil {
		return nil, invAstNilError().noPos()
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
	case *ast.ArrayType:
		return astArrayType(v, idents)
	case *ast.IndexExpr:
		return astIndexExpr(v, idents)
	case *ast.SliceExpr:
		return astSliceExpr(v, idents)
	case *ast.CompositeLit:
		return astCompositeLit(v, idents)
	default:
		return nil, invAstUnsupportedError(e).pos(e)
	}
}

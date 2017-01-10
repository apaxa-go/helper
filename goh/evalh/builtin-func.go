package evalh

import (
	"github.com/apaxa-go/helper/goh/constanth"
	"go/constant"
	"go/token"
	"reflect"
)

func isBuiltInFunc(ident string) bool {
	switch ident {
	case "len", "cap", "complex", "real", "imag", "new", "make", "append":
		return true
	default:
		return false
	}
}

// ok means that astCallExpr is built-in function, not that calling done without error.
func callBuiltInFunc(f string, args []Value, ellipsis bool) (r Value, err *intError) {
	if f != "append" && ellipsis {
		return nil, callBuiltInWithEllipsisError(f)
	}
	switch f {
	case "len":
		if len(args) != 1 {
			err = callBuiltInArgsCountMismError(f, 1, len(args))
			return
		}
		return BuiltInLen(args[0])
	case "cap":
		if len(args) != 1 {
			err = callBuiltInArgsCountMismError(f, 1, len(args))
			return
		}
		return BuiltInCap(args[0])
	case "complex":
		if len(args) != 2 {
			err = callBuiltInArgsCountMismError(f, 2, len(args))
			return
		}
		return BuiltInComplex(args[0], args[1])
	case "real":
		if len(args) != 1 {
			err = callBuiltInArgsCountMismError(f, 1, len(args))
			return
		}
		return BuiltInReal(args[0])
	case "imag":
		if len(args) != 1 {
			err = callBuiltInArgsCountMismError(f, 1, len(args))
			return
		}
		return BuiltInImag(args[0])
	case "new":
		if len(args) != 1 {
			err = callBuiltInArgsCountMismError(f, 1, len(args))
			return
		}
		return BuiltInNew(args[0])
	case "make":
		return BuiltInMake(args)
	case "append":
		if len(args) < 1 {
			err = callBuiltInArgsCountMismError(f, 1, len(args))
			return
		}
		return BuiltInAppend(args[0], args[1:], ellipsis)
	default:
		return nil, undefIdentError(f)
	}
}

func builtInLenConstant(v constant.Value) (r Value, err *intError) {
	const fn = "len"
	if v.Kind() != constant.String {
		return nil, invBuiltInArgError(fn, MakeUntyped(v))
	}
	l := len(constant.StringVal(v))
	return MakeUntypedInt64(int64(l)), nil
}

func builtInLen(v reflect.Value) (r Value, err *intError) {
	const fn = "len"
	// Resolve pointer to array
	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Array {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return MakeRegularInterface(v.Len()), nil
	default:
		return nil, invBuiltInArgError(fn, MakeRegular(v))
	}
}

// Not fully following GoLang spec (always returns int instead of untyped for array & pointer to array).
func BuiltInLen(v Value) (r Value, err *intError) {
	const fn = "len"
	switch v.Kind() {
	case Untyped:
		return builtInLenConstant(v.Untyped())
	case Regular:
		return builtInLen(v.Regular())
	default:
		return nil, invBuiltInArgError(fn, v)
	}
}

func builtInCap(v reflect.Value) (r Value, err *intError) {
	const fn = "cap"
	// Resolve pointer to array
	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Array {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice:
		return MakeRegularInterface(v.Cap()), nil
	default:
		return nil, invBuiltInArgError(fn, MakeRegular(v))
	}
}

//BUG: Not fully following GoLang spec (always returns int instead of untyped for array & pointer to array).
func BuiltInCap(v Value) (r Value, err *intError) {
	const fn = "cap"
	switch v.Kind() {
	case Regular:
		return builtInCap(v.Regular())
	default:
		return nil, invBuiltInArgError(fn, v)
	}
}

func builtInComplexConstant(realPart, imaginaryPart constant.Value) (r Value, err *intError) {
	const fn = "complex"
	switch realPart.Kind() {
	case constant.Int, constant.Float:
	default:
		return nil, invBuiltInArgAtError(fn, 0, MakeUntyped(realPart))
	}

	switch imaginaryPart.Kind() {
	case constant.Int, constant.Float:
	default:
		return nil, invBuiltInArgAtError(fn, 1, MakeUntyped(imaginaryPart))
	}

	rC := constant.BinaryOp(realPart, token.ADD, constant.MakeImag(imaginaryPart))
	if rC.Kind() != constant.Complex {
		return nil, invBuiltInArgsError(fn, []Value{MakeUntyped(realPart), MakeUntyped(imaginaryPart)})
	}
	return MakeUntyped(rC), nil
}

func BuiltInComplexArgParse(a Value) (r float64, can32, can64 bool) {
	switch a.Kind() {
	case Untyped:
		aConst := a.Untyped()
		aConst = constant.ToFloat(aConst)
		if aConst.Kind() != constant.Float {
			return 0, false, false
		}
		r, can64 = constanth.Float64Val(aConst)
		if !can64 {
			return 0, false, false
		}
		_, can32 = constanth.Float32Val(aConst)
		return
	case Regular:
		if !a.Regular().CanInterface() {
			return 0, false, false
		}
		switch aFloat := a.Regular().Interface().(type) {
		case float32:
			return float64(aFloat), true, false
		case float64:
			return aFloat, false, true
		}
	}
	return 0, false, false
}

func BuiltInComplex(realPart, imaginaryPart Value) (r Value, err *intError) {
	const fn = "complex"
	if realPart.Kind() == Untyped && imaginaryPart.Kind() == Untyped {
		return builtInComplexConstant(realPart.Untyped(), imaginaryPart.Untyped())
	}

	// Prepare arguments
	rF, r32, r64 := BuiltInComplexArgParse(realPart)
	if !r32 && !r64 {
		return nil, invBuiltInArgAtError(fn, 0, realPart)
	}
	iF, i32, i64 := BuiltInComplexArgParse(imaginaryPart)
	if !i32 && !i64 {
		return nil, invBuiltInArgAtError(fn, 1, imaginaryPart)
	}

	// Calc
	if r32 && i32 {
		return MakeRegularInterface(complex(float32(rF), float32(iF))), nil
	}
	if r64 && i64 {
		return MakeRegularInterface(complex(rF, iF)), nil
	}
	return nil, invBuiltInArgsError(fn, []Value{realPart, imaginaryPart})
}

func builtInRealConstant(v constant.Value) (r Value, err *intError) {
	const fn = "real"
	if !constanth.IsNumeric(v) {
		return nil, invBuiltInArgError(fn, MakeUntyped(v))
	}
	rV := constant.Real(v)
	if rV.Kind() == constant.Unknown {
		return nil, invBuiltInArgError(fn, MakeUntyped(v))
	}
	return MakeUntyped(rV), nil
}

func builtInReal(v reflect.Value) (r Value, err *intError) {
	const fn = "real"
	if !v.CanInterface() {
		return nil, invBuiltInArgError(fn, MakeRegular(v))
	}
	switch vI := v.Interface().(type) {
	case complex64:
		return MakeRegularInterface(real(vI)), nil
	case complex128:
		return MakeRegularInterface(real(vI)), nil
	default:
		return nil, invBuiltInArgError(fn, MakeRegular(v))
	}
}

func BuiltInReal(v Value) (r Value, err *intError) {
	const fn = "real"
	switch v.Kind() {
	case Untyped:
		return builtInRealConstant(v.Untyped())
	case Regular:
		return builtInReal(v.Regular())
	default:
		return nil, invBuiltInArgError(fn, v)
	}
}

func builtInImagConstant(v constant.Value) (r Value, err *intError) {
	const fn = "imag"
	if !constanth.IsNumeric(v) {
		return nil, invBuiltInArgError(fn, MakeUntyped(v))
	}
	rV := constant.Imag(v)
	if rV.Kind() == constant.Unknown {
		return nil, invBuiltInArgError(fn, MakeUntyped(v))
	}
	return MakeUntyped(rV), nil
}

func builtInImag(v reflect.Value) (r Value, err *intError) {
	const fn = "imag"
	if !v.CanInterface() {
		return nil, invBuiltInArgError(fn, MakeRegular(v))
	}
	switch vI := v.Interface().(type) {
	case complex64:
		return MakeRegularInterface(imag(vI)), nil
	case complex128:
		return MakeRegularInterface(imag(vI)), nil
	default:
		return nil, invBuiltInArgError(fn, MakeRegular(v))
	}
}

func BuiltInImag(v Value) (r Value, err *intError) {
	const fn = "imag"
	switch v.Kind() {
	case Untyped:
		return builtInImagConstant(v.Untyped())
	case Regular:
		return builtInImag(v.Regular())
	default:
		return nil, invBuiltInArgError(fn, v)
	}
}

func BuiltInNew(v Value) (r Value, err *intError) {
	const fn = "new"
	switch v.Kind() {
	case Type:
		return MakeRegular(reflect.New(v.Type())), nil
	default:
		return nil, notTypeError(v)
	}
}

func BuiltInMake(v []Value) (r Value, err *intError) {
	const fn = "make"
	if len(v) < 1 {
		return nil, callBuiltInArgsCountMismError(fn, 1, len(v))
	}
	if v[0].Kind() != Type {
		return nil, notTypeError(v[0])
	}
	switch t := v[0].Type(); t.Kind() {
	case reflect.Slice:
		var n, m int
		switch len(v) {
		case 2:
			var ok bool
			n, _, ok = v[1].AsInt()
			if !ok {
				return nil, makeNotIntArgError(t, "len", v[1])
			}
			m = n
		case 3:
			var ok bool
			n, _, ok = v[1].AsInt()
			if !ok {
				return nil, makeNotIntArgError(t, "len", v[1])
			}
			m, _, ok = v[2].AsInt()
			if !ok {
				return nil, makeNotIntArgError(t, "cap", v[2])
			}
		default:
			return nil, callBuiltInArgsCountMismError(fn, 2, len(v))
		}
		if n < 0 {
			return nil, makeNegArgError(t, "len")
		}
		if n > m {
			return nil, makeSliceMismArgsError(t)
		}
		return MakeRegular(reflect.MakeSlice(t, n, m)), nil
	case reflect.Map:
		// BUG make(<map>,n) ignore n (but check it type).
		switch len(v) {
		case 1:
			// nothing to do
		case 2:
			n, _, ok := v[1].AsInt()
			if !ok {
				return nil, makeNotIntArgError(t, "size", v[1])
			}
			if n < 0 {
				return nil, makeNegArgError(t, "size")
			}
		default:
			return nil, callBuiltInArgsCountMismError(fn, 1, len(v))
		}
		return MakeRegular(reflect.MakeMap(t)), nil
	case reflect.Chan:
		var n int
		switch len(v) {
		case 1:
			// nothing to do
		case 2:
			var ok bool
			n, _, ok = v[1].AsInt()
			if !ok {
				return nil, makeNotIntArgError(t, "buffer", v[1])
			}
			if n < 0 {
				return nil, makeNegArgError(t, "buffer")
			}
		}
		return MakeRegular(reflect.MakeChan(t, n)), nil
	default:
		return nil, makeInvalidTypeError(t)
	}
}

func BuiltInAppend(v Value, a []Value, ellipsis bool) (r Value, err *intError) {
	const fn = "append"
	if v.Kind() != Regular {
		return nil, appendFirstNotSliceError(v)
	}
	vV := v.Regular()
	if vV.Kind() != reflect.Slice {
		return nil, appendFirstNotSliceError(v)
	}

	elemT := v.Regular().Type().Elem()
	//var aV []reflect.Value
	switch ellipsis {
	case true:
		if len(a) != 1 {
			return nil, callBuiltInArgsCountMismError(fn, 2, 1+len(a))
		}
		if a[0].Kind() != Regular {
			return nil, appendMismTypeError(reflect.SliceOf(elemT), a[0])
		}
		aV := a[0].Regular()
		if aV.Kind() != reflect.Slice {
			return nil, appendMismTypeError(reflect.SliceOf(elemT), a[0])
		}
		if !aV.Type().Elem().AssignableTo(elemT) {
			return nil, appendMismTypeError(reflect.SliceOf(elemT), a[0])
		}
		return MakeRegular(reflect.AppendSlice(vV, aV)), nil
	case false:
		aV := make([]reflect.Value, len(a))
		for i := range a {
			var ok bool
			aV[i], _, ok = a[i].AsType(elemT)
			if !ok {
				return nil, appendMismTypeError(elemT, a[i])
			}
		}
		return MakeRegular(reflect.Append(vV, aV...)), nil
	default:
		return nil, nil // unreachable
	}
}

package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/constanth"
	"go/constant"
	"go/token"
	"log"
	"reflect"
)

func isBuiltInFunc(ident string) bool {
	switch ident {
	case "len", "cap", "complex", "real", "imag":
		return true
	default:
		return false
	}
}

// ok means that CallExpr is built-in function, not that calling done without error.
func callBuiltInFunc(f string, args []Value) (r Value, err error) {	// TODO new, <-, ->
	switch f {
	case "len":
		if len(args) != 1 {
			err = errors.New("tcbi0") // TODO
			return
		}
		return BuiltInLen(args[0])
	case "cap":
		if len(args) != 1 {
			err = errors.New("tcbi1") // TODO
			return
		}
		return BuiltInCap(args[0])
	case "complex":
		if len(args) != 2 {
			err = errors.New("tcbi2") // TODO
			return
		}
		return BuiltInComplex(args[0], args[1])
	case "real":
		if len(args) != 1 {
			err = errors.New("tcbi3") // TODO
			return
		}
		return BuiltInReal(args[0])
	case "imag":
		if len(args) != 1 {
			err = errors.New("tcbi4") // TODO
			return
		}
		return BuiltInImag(args[0])
	default:
		return nil, errors.New("not a built-in function: " + f)
	}
}

func builtInLenConstant(v constant.Value) (r Value, err error) {
	if v.Kind() != constant.String {
		return nil, errors.New("unable to call len with " + v.String())
	}
	l := len(constant.StringVal(v))
	return MakeUntypedInt64(int64(l)), nil
}

func builtInLen(v reflect.Value) (r Value, err error) {
	// Resolve pointer to array
	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Array {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return MakeRegularInterface(v.Len()), nil
	default:
		return nil, errors.New("unable to call len with " + v.String())
	}
}

// Not fully following GoLang spec (always returns int instead of untyped for array & pointer to array).
func BuiltInLen(v Value) (r Value, err error) {
	switch v.Kind() {
	case Untyped:
		return builtInLenConstant(v.Untyped())
	case Regular:
		return builtInLen(v.Regular())
	default:
		return nil, errors.New("unable to call len with " + v.String())
	}
}

func builtInCap(v reflect.Value) (r Value, err error) {
	// Resolve pointer to array
	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Array {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice:
		return MakeRegularInterface(v.Cap()), nil
	default:
		return nil, errors.New("unable to call cap with " + v.String())
	}
}

// Not fully following GoLang spec (always returns int instead of untyped for array & pointer to array).
func BuiltInCap(v Value) (r Value, err error) {
	switch v.Kind() {
	case Regular:
		return builtInCap(v.Regular())
	default:
		return nil, errors.New("unable to call cap with " + v.String())
	}
}

func builtInComplexConstant(realPart, imaginaryPart constant.Value) (r Value, err error) {
	switch realPart.Kind() {
	case constant.Int, constant.Float:
	default:
		return nil, errors.New("") // TODO
	}

	switch imaginaryPart.Kind() {
	case constant.Int, constant.Float:
	default:
		return nil, errors.New("") // TODO
	}

	rC := constant.BinaryOp(realPart, token.ADD, constant.MakeImag(imaginaryPart))
	if rC.Kind() != constant.Complex {
		return nil, errors.New("") // TODO
	}
	return MakeUntyped(rC), nil
}

func BuiltInComplexArgParse(a Value) (r float64, can32, can64 bool) {
	switch a.Kind() {
	case Untyped:
		aConst := a.Untyped()
		aConst = constant.ToFloat(aConst)
		if aConst.Kind() != constant.Float {
			log.Println("1# ", constanth.KindString(aConst.Kind()))
			return 0, false, false
		}
		r, can64 = constanth.Float64Val(aConst)
		if !can64 {
			log.Println("2# ", a)
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

func BuiltInComplex(realPart, imaginaryPart Value) (r Value, err error) {
	if realPart.Kind() == Untyped && imaginaryPart.Kind() == Untyped {
		return builtInComplexConstant(realPart.Untyped(), imaginaryPart.Untyped())
	}

	rF, r32, r64 := BuiltInComplexArgParse(realPart)
	if err != nil {
		return nil, err
	}
	iF, i32, i64 := BuiltInComplexArgParse(imaginaryPart)
	if err != nil {
		return nil, err
	}

	if r32 && i32 {
		return MakeRegularInterface(complex(float32(rF), float32(iF))), nil
	}
	if r64 && i64 {
		return MakeRegularInterface(complex(rF, iF)), nil
	}
	log.Println(realPart)
	log.Println(BuiltInComplexArgParse(realPart))
	log.Println(imaginaryPart)
	log.Println(BuiltInComplexArgParse(imaginaryPart))
	return nil, errors.New("unable to call built-in complex on " + realPart.String() + " and " + imaginaryPart.String())
}

func builtInRealConstant(v constant.Value) (r Value, err error) {
	if !constanth.IsNumeric(v) {
		return nil, errors.New("") // TODO
	}
	rV := constant.Real(v)
	if rV.Kind() == constant.Unknown {
		return nil, errors.New("") // TODO
	}
	return MakeUntyped(rV), nil
}

func builtInReal(v reflect.Value) (r Value, err error) {
	if !v.CanInterface() {
		return nil, errors.New("") // TODO
	}
	switch vI := v.Interface().(type) {
	case complex64:
		return MakeRegularInterface(real(vI)), nil
	case complex128:
		return MakeRegularInterface(real(vI)), nil
	default:
		return nil, errors.New("") // TODO
	}
}

func BuiltInReal(v Value) (r Value, err error) {
	switch v.Kind() {
	case Untyped:
		return builtInRealConstant(v.Untyped())
	case Regular:
		return builtInReal(v.Regular())
	default:
		return nil, errors.New("unable to call real with " + v.String())
	}
}

func builtInImagConstant(v constant.Value) (r Value, err error) {
	if !constanth.IsNumeric(v) {
		return nil, errors.New("") // TODO
	}
	rV := constant.Imag(v)
	if rV.Kind() == constant.Unknown {
		return nil, errors.New("") // TODO
	}
	return MakeUntyped(rV), nil
}

func builtInImag(v reflect.Value) (r Value, err error) {
	if !v.CanInterface() {
		return nil, errors.New("") // TODO
	}
	switch vI := v.Interface().(type) {
	case complex64:
		return MakeRegularInterface(imag(vI)), nil
	case complex128:
		return MakeRegularInterface(imag(vI)), nil
	default:
		return nil, errors.New("") // TODO
	}
}

func BuiltInImag(v Value) (r Value, err error) {
	switch v.Kind() {
	case Untyped:
		return builtInImagConstant(v.Untyped())
	case Regular:
		return builtInImag(v.Regular())
	default:
		return nil, errors.New("unable to call imag with " + v.String())
	}
}

package evalh

import (
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/goh/tokenh"
	"github.com/apaxa-go/helper/reflecth"
	"go/constant"
	"go/token"
	"reflect"
)

func binaryCompare(x Value, op token.Token, y Value) (r Value, err *intError) {
	var xV, yV reflect.Value

	// Prepare arguments
	switch xK, yK := x.Kind(), y.Kind(); {
	case xK == Nil && yK == Regular:
		return binaryCompareWithNil(y.Regular(), op)
	case xK == Regular && yK == Nil:
		return binaryCompareWithNil(x.Regular(), op)
	case xK == Untyped && yK == Untyped:
		return binaryCompareConstant(x.Untyped(), op, y.Untyped())
	case xK == Untyped && yK == Regular:
		yV = y.Regular()
		var ok bool
		xV, ok = constanth.AsType(x.Untyped(), yV.Type())
		if !ok {
			return nil, invBinOpTypesMismError(x, op, y)
		}
	case xK == Regular && yK == Untyped:
		xV = x.Regular()
		var ok bool
		yV, ok = constanth.AsType(y.Untyped(), xV.Type())
		if !ok {
			return nil, invBinOpTypesMismError(x, op, y)
		}
	case xK == Regular && yK == Regular:
		xV = x.Regular()
		yV = y.Regular()
	default:
		return nil, invBinOpTypesInvalError(x, op, y)
	}

	// Basic check
	if xT, yT := xV.Type(), yV.Type(); !xT.AssignableTo(yT) && !yT.AssignableTo(xT) {
		return nil, invBinOpTypesIncompError(x, op, y)
	}

	// Choose compare
	switch {
	case tokenh.IsComparisonCompare(op):
		if xT, yT := xV.Type(), yV.Type(); !xT.Comparable() || !yT.Comparable() {
			return nil, invBinOpTypesIncompError(x, op, y)
		}
		equality := op == token.EQL
		return MakeUntypedBool(xV.Interface() == yV.Interface() == equality), nil
	case tokenh.IsComparisonOrder(op):
		switch k := xV.Kind(); {
		case reflecth.IsInt(k):
			return binaryCompareInt(xV.Int(), op, yV.Int())
		case reflecth.IsUint(k):
			return binaryCompareUint(xV.Uint(), op, yV.Uint())
		case reflecth.IsAnyFloat(k):
			return binaryCompareFloat(xV.Float(), op, yV.Float())
		case k == reflect.String:
			return binaryCompareString(xV.String(), op, yV.String())
		default:
			return nil, invBinOpTypesUnorderError(x, op, y)
		}
	default:
		return nil, invBinOpInvalError(x, op, y)
	}

	/*
		// Preform kind-depending compare
		switch k := xV.Kind(); {
		case k == reflect.Bool:
			return binaryCompareBool(xV.Bool(), op, yV.Bool())
		case reflecth.IsInt(k):
			return binaryCompareInt(xV.Int(), op, yV.Int())
		case reflecth.IsUint(k):
			return binaryCompareUint(xV.Uint(), op, yV.Uint())
		case reflecth.IsAnyFloat(k):
			return binaryCompareFloat(xV.Float(), op, yV.Float())
		case reflecth.IsAnyComplex(k):
			return binaryCompareComplex(xV.Complex(), op, yV.Complex())
		case k == reflect.String:
			return binaryCompareString(xV.String(), op, yV.String())
		//case k == reflect.Ptr:	TO DO what does this mean?
		//	fallthrough
		case k == reflect.Uintptr:
			return binaryComparePointer(uintptr(xV.Uint()), op, uintptr(yV.Uint()))
		case k == reflect.Chan:
			return binaryComparePointer(xV.Pointer(), op, yV.Pointer())
		TO DO compare interfaces
		// case k==reflect.Interface:
		TO DO compare structs
		// case k==reflect.Struct:
		TO DO compare arrays
		// case k==reflect.Array:
		default:
			return nil, errors.New("comparison of " + k.String() + " and " + yV.Kind().String() + " does not allowed or currently does not implemented")
		}
	*/

}

func binaryShift(x Value, op token.Token, y Value) (r Value, err *intError) {
	// Calc right operand
	var yUint uint
	switch y.Kind() {
	case Untyped:
		var ok bool
		yUint, ok = constanth.UintVal(y.Untyped())
		if !ok {
			return nil, invBinOpShiftCountError(x, op, y)
		}
	case Regular:
		if !reflecth.IsUint(y.Regular().Kind()) {
			return nil, invBinOpShiftCountError(x, op, y)
		}
		yUint = uint(y.Regular().Uint())
	default:
		return nil, invBinOpShiftCountError(x, op, y)
	}

	switch x.Kind() {
	case Untyped:
		return binaryShiftConstant(x.Untyped(), op, yUint)
	case Regular:
		break
	default:
		return nil, invBinOpShiftArgError(x, op, y)
	}

	// Check left operand kind and set result type
	xV := x.Regular()
	if !reflecth.IsAnyInt(xV.Kind()) {
		return nil, invBinOpShiftArgError(x, op, y)
	}
	rV := reflect.New(xV.Type()).Elem()

	switch op {
	case token.SHL:
		switch xV.Kind() {
		case reflect.Int:
			rV.SetInt(int64(int(xV.Int()) << yUint))
		case reflect.Int8:
			rV.SetInt(int64(int8(xV.Int()) << yUint))
		case reflect.Int16:
			rV.SetInt(int64(int16(xV.Int()) << yUint))
		case reflect.Int32:
			rV.SetInt(int64(int32(xV.Int()) << yUint))
		case reflect.Int64:
			rV.SetInt(xV.Int() << yUint)
		case reflect.Uint:
			rV.SetUint(uint64(uint(xV.Uint()) << yUint))
		case reflect.Uint8:
			rV.SetUint(uint64(uint8(xV.Uint()) << yUint))
		case reflect.Uint16:
			rV.SetUint(uint64(uint16(xV.Uint()) << yUint))
		case reflect.Uint32:
			rV.SetUint(uint64(uint32(xV.Uint()) << yUint))
		case reflect.Uint64:
			rV.SetUint(xV.Uint() << yUint)
		default:
			return nil, invBinOpShiftArgError(x, op, y)
		}
	case token.SHR:
		switch xV.Kind() {
		case reflect.Int:
			rV.SetInt(int64(int(xV.Int()) >> yUint))
		case reflect.Int8:
			rV.SetInt(int64(int8(xV.Int()) >> yUint))
		case reflect.Int16:
			rV.SetInt(int64(int16(xV.Int()) >> yUint))
		case reflect.Int32:
			rV.SetInt(int64(int32(xV.Int()) >> yUint))
		case reflect.Int64:
			rV.SetInt(xV.Int() >> yUint)
		case reflect.Uint:
			rV.SetUint(uint64(uint(xV.Uint()) >> yUint))
		case reflect.Uint8:
			rV.SetUint(uint64(uint8(xV.Uint()) >> yUint))
		case reflect.Uint16:
			rV.SetUint(uint64(uint16(xV.Uint()) >> yUint))
		case reflect.Uint32:
			rV.SetUint(uint64(uint32(xV.Uint()) >> yUint))
		case reflect.Uint64:
			rV.SetUint(xV.Uint() >> yUint)
		default:
			return nil, invBinOpShiftArgError(x, op, y)
		}
	default:
		return nil, invBinOpInvalError(x, op, y)
	}

	return MakeRegular(rV), nil
}

func binaryOther(x Value, op token.Token, y Value) (r Value, err *intError) {
	var xV, yV reflect.Value

	// Prepare arguments
	switch xK, yK := x.Kind(), y.Kind(); {
	case xK == Untyped && yK == Untyped:
		return binaryOtherConstant(x.Untyped(), op, y.Untyped())
	case xK == Untyped && yK == Regular:
		yV = y.Regular()
		var ok bool
		xV, ok = constanth.AsType(x.Untyped(), yV.Type())
		if !ok {
			return nil, invBinOpTypesMismError(x, op, y)
		}
	case xK == Regular && yK == Untyped:
		xV = x.Regular()
		var ok bool
		yV, ok = constanth.AsType(y.Untyped(), xV.Type())
		if !ok {
			return nil, invBinOpTypesMismError(x, op, y)
		}
	case xK == Regular && yK == Regular:
		xV = x.Regular()
		yV = y.Regular()

		if xT, yT := xV.Type(), yV.Type(); xT != yT {
			return nil, invBinOpTypesMismError(x, op, y)
		}
	default:
		return nil, invBinOpTypesInvalError(x, op, y)
	}

	rV := reflect.New(xV.Type()).Elem()
	switch xV.Kind() {
	case reflect.Bool:
		v, err := binaryOtherBool(xV.Bool(), op, yV.Bool())
		if err != nil {
			return nil, err
		}
		rV.SetBool(v)
	case reflect.Int:
		v, err := binaryOtherInt(int(xV.Int()), op, int(yV.Int()))
		if err != nil {
			return nil, err
		}
		rV.SetInt(int64(v))
	case reflect.Int8:
		v, err := binaryOtherInt8(int8(xV.Int()), op, int8(yV.Int()))
		if err != nil {
			return nil, err
		}
		rV.SetInt(int64(v))
	case reflect.Int16:
		v, err := binaryOtherInt16(int16(xV.Int()), op, int16(yV.Int()))
		if err != nil {
			return nil, err
		}
		rV.SetInt(int64(v))
	case reflect.Int32:
		v, err := binaryOtherInt32(int32(xV.Int()), op, int32(yV.Int()))
		if err != nil {
			return nil, err
		}
		rV.SetInt(int64(v))
	case reflect.Int64:
		v, err := binaryOtherInt64(xV.Int(), op, yV.Int())
		if err != nil {
			return nil, err
		}
		rV.SetInt(v)
	case reflect.Uint:
		v, err := binaryOtherUint(uint(xV.Uint()), op, uint(yV.Uint()))
		if err != nil {
			return nil, err
		}
		rV.SetUint(uint64(v))
	case reflect.Uint8:
		v, err := binaryOtherUint8(uint8(xV.Uint()), op, uint8(yV.Uint()))
		if err != nil {
			return nil, err
		}
		rV.SetUint(uint64(v))
	case reflect.Uint16:
		v, err := binaryOtherUint16(uint16(xV.Uint()), op, uint16(yV.Uint()))
		if err != nil {
			return nil, err
		}
		rV.SetUint(uint64(v))
	case reflect.Uint32:
		v, err := binaryOtherUint32(uint32(xV.Uint()), op, uint32(yV.Uint()))
		if err != nil {
			return nil, err
		}
		rV.SetUint(uint64(v))
	case reflect.Uint64:
		v, err := binaryOtherUint64(xV.Uint(), op, yV.Uint())
		if err != nil {
			return nil, err
		}
		rV.SetUint(v)
	case reflect.Float32:
		v, err := binaryOtherFloat32(float32(xV.Float()), op, float32(yV.Float()))
		if err != nil {
			return nil, err
		}
		rV.SetFloat(float64(v))
	case reflect.Float64:
		v, err := binaryOtherFloat64(xV.Float(), op, yV.Float())
		if err != nil {
			return nil, err
		}
		rV.SetFloat(v)
	case reflect.Complex64:
		v, err := binaryOtherComplex64(complex64(xV.Complex()), op, complex64(yV.Complex()))
		if err != nil {
			return nil, err
		}
		rV.SetComplex(complex128(v))
	case reflect.Complex128:
		v, err := binaryOtherComplex128(xV.Complex(), op, yV.Complex())
		if err != nil {
			return nil, err
		}
		rV.SetComplex(v)
	case reflect.String:
		v, err := binaryOtherString(xV.String(), op, yV.String())
		if err != nil {
			return nil, err
		}
		rV.SetString(v)
	default:
		return nil, invBinOpTypesInvalError(x, op, y)
	}
	return MakeRegular(rV), nil
}

func binaryCompareConstant(x constant.Value, op token.Token, y constant.Value) (r Value, err *intError) {
	if x.Kind() == constant.Unknown || y.Kind() == constant.Unknown {
		return nil, invBinOpTypesInvalError(MakeUntyped(x), op, MakeUntyped(y))
	}
	defer func() {
		rec := recover()
		if rec != nil {
			r = nil
			err = newIntErrorf(invBinOp, x, op, y, rec)
		}
	}()
	return MakeUntyped(constant.MakeBool(constant.Compare(x, op, y))), nil
}

func binaryShiftConstant(x constant.Value, op token.Token, y uint) (r Value, err *intError) {
	if x.Kind() == constant.Unknown {
		return nil, invBinOpTypesInvalError(MakeUntyped(x), op, MakeRegularInterface(y))
	}
	defer func() {
		rec := recover()
		if rec != nil {
			r = nil
			err = newIntErrorf(invBinOp, x, op, y, rec)
		}
	}()
	return MakeUntyped(constant.Shift(x, op, y)), nil // should not return unknown because x already checked for unknown
}

func binaryOtherConstant(x constant.Value, op token.Token, y constant.Value) (r Value, err *intError) {
	if x.Kind() == constant.Unknown || y.Kind() == constant.Unknown {
		return nil, invBinOpTypesInvalError(MakeUntyped(x), op, MakeUntyped(y))
	}
	defer func() {
		rec := recover()
		if rec != nil {
			r = nil
			err = newIntErrorf(invBinOp, x, op, y, rec)
		}
	}()
	return MakeUntyped(constant.BinaryOp(x, op, y)), nil
}

package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/reflecth"
	"github.com/apaxa-go/helper/strconvh"
	"go/constant"
	"go/token"
	"reflect"
)

func binaryCompare(x Value, op token.Token, y Value) (r Value, err error) {
	if x.Kind() == Untyped && y.Kind() == Untyped {
		return binaryCompareConstant(x.Untyped(), op, y.Untyped())
	}

	var xV, yV reflect.Value

	// Convert untyped const to appropriate variable or check type assignability
	if x.Kind() == Untyped {
		yV = y.Regular() // TODO what if y is of type Nil
		var ok bool
		xV, ok = constanth.SameType(x.Untyped(), yV.Type())
		if !ok {
			return nil, errors.New("unable to convert " + x.String() + " to " + yV.Type().String())
		}
	} else if y.Kind() == Untyped {
		xV = x.Regular() // TODO what if x is of type Nil
		var ok bool
		yV, ok = constanth.SameType(y.Untyped(), xV.Type())
		if !ok {
			return nil, errors.New("unable to convert untyped const " + y.String() + " to " + xV.Type().String())
		}
	} else {
		xV = x.Regular() // TODO what if x is of type Nil
		yV = y.Regular() // TODO what if y is of type Nil

		if xT, yT := xV.Type(), yV.Type(); !xT.AssignableTo(yT) && !yT.AssignableTo(xT) {
			return nil, errors.New("unable to compare " + xT.String() + " and " + yT.String())
		}
	}

	// Preform kind-depending compare
	switch k := xV.Kind(); {
	case k == reflect.Bool:
		return binaryCompareBool(xV.Bool(), op, yV.Bool())
	case reflecth.IsInt(k):
		return binaryCompareInt(xV.Int(), op, yV.Int())
	case reflecth.IsUint(k):
		return binaryCompareUint(xV.Uint(), op, yV.Uint())
	case reflecth.IsFloat(k):
		return binaryCompareFloat(xV.Float(), op, yV.Float())
	case reflecth.IsComplex(k):
		return binaryCompareComplex(xV.Complex(), op, yV.Complex())
	case k == reflect.String:
		return binaryCompareString(xV.String(), op, yV.String())
	//case k == reflect.Ptr:	// TODO what does this mean?
	//	fallthrough
	case k == reflect.Uintptr:
		return binaryComparePointer(uintptr(xV.Uint()), op, uintptr(yV.Uint()))
	// TODO compare channels
	//case k == reflect.Chan: // Not sure about channels comparison via pointer
	//	return binaryComparePointer(xV.Pointer(), op, yV.Pointer())
	// TODO compare interfaces
	// case k==reflect.Interface:
	// TODO compare structs
	// case k==reflect.Struct:
	// TODO compare arrays
	// case k==reflect.Array:
	// TODO compare slice, map & function with nil
	default:
		return nil, errors.New("comparison of " + k.String() + " and " + k.String() + " does not allowed or currently does not implemented")
	}
}

func binaryShift(x Value, op token.Token, y Value) (r Value, err error) {
	// Calc right operand
	var yUint uint
	switch y.Kind() {
	case Untyped:
		var ok bool
		yUint, ok = constanth.UintVal(y.Untyped())
		if !ok {
			return nil, errors.New("unable to perform binary shift on smth and " + y.String())
		}
	case Regular:
		if !reflecth.IsUint(y.Regular().Kind()) {
			return nil, errors.New("unable to perform shift on smth and " + y.String())
		}
		yUint = uint(y.Regular().Uint())
	case Nil:
		return nil, errors.New("unable to perform shift on smth and " + y.String())
	default:
		panic("unknown y kind")
	}

	switch x.Kind() {
	case Untyped:
		// Perform calc separately if left operand is untyped constant
		return binaryShiftConstant(x.Untyped(), op, yUint)
	case Nil:
		return nil, errors.New("unable to perform shift on " + y.String())
	case Regular:
		// No panic
	default:
		panic("unknown x kind")
	}

	// Check left operand kind and set result type
	xV := x.Regular()
	if !reflecth.IsAnyInt(xV.Kind()) {
		return nil, errors.New("unable to perform shift on " + x.String() + " and smth")
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
			return nil, errors.New("unable to perform left shift on " + x.String())
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
			return nil, errors.New("unable to perform right shift on " + x.String())
		}
	default:
		return nil, errors.New("unable to perform shift opearation " + op.String())
	}

	return MakeRegular(rV), nil
}

func binaryOther(x Value, op token.Token, y Value) (r Value, err error) {
	// Perform calc separately if both args is untyped constants
	if x.Kind() == Untyped && y.Kind() == Untyped {
		return binaryOtherConstant(x.Untyped(), op, y.Untyped())
	}

	var xV, yV reflect.Value

	// Convert untyped const to appropriate variable or check type assignability
	if x.Kind() == Untyped {
		yV = y.Regular() // TODO what if y is of type Nil
		var ok bool
		xV, ok = constanth.SameType(x.Untyped(), yV.Type())
		if !ok {
			return nil, errors.New("unable to convert " + x.String() + " to " + yV.Type().String())
		}
	} else if y.Kind() == Untyped {
		xV = x.Regular() // TODO what if x is of type Nil
		var ok bool
		yV, ok = constanth.SameType(y.Untyped(), xV.Type())
		if !ok {
			return nil, errors.New("unable to convert " + y.String() + " to " + xV.Type().String())
		}
	} else {
		xV = x.Regular() // TODO what if x is of type Nil
		yV = y.Regular() // TODO what if y is of type Nil

		if xT, yT := xV.Type(), yV.Type(); xT != yT {
			return nil, errors.New("unable to perform binary operation on " + xT.String() + " and " + yT.String())
		}
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
		return nil, errors.New("binary operation " + op.String() + " currently not implemented")
	}
	return MakeRegular(rV), nil
}

func binaryCompareConstant(x constant.Value, op token.Token, y constant.Value) (r Value, err error) {
	if x.Kind() == constant.Unknown || y.Kind() == constant.Unknown {
		return nil, errors.New("unable to perform compare on unknown constant")
	}
	defer func() {
		rec := recover()
		if rec != nil {
			r = nil
			err = errors.New("unable to perform binary operation '" + op.String() + "' on " + x.String() + " and " + y.String())
		}
	}()
	return MakeUntyped(constant.MakeBool(constant.Compare(x, op, y))), nil
}

func binaryShiftConstant(x constant.Value, op token.Token, y uint) (r Value, err error) {
	if x.Kind() == constant.Unknown {
		return nil, errors.New("unable to perform shift on unknown constant")
	}
	defer func() {
		rec := recover()
		if rec != nil {
			r = nil
			err = errors.New("unable to perform binary operation '" + op.String() + "' on " + x.String() + " and " + strconvh.FormatUint(y))
		}
	}()
	return MakeUntyped(constant.Shift(x, op, y)), nil // should not return unknown because x already checked for unknown
}

func binaryOtherConstant(x constant.Value, op token.Token, y constant.Value) (r Value, err error) {
	if x.Kind() == constant.Unknown || y.Kind() == constant.Unknown {
		return nil, errors.New("unable to perform compare on unknown constant")
	}
	defer func() {
		rec := recover()
		if rec != nil {
			r = nil
			err = errors.New("unable to perform binary operation '" + op.String() + "' on " + x.String() + " and " + y.String())
		}
	}()
	return MakeUntyped(constant.BinaryOp(x, op, y)), nil
}

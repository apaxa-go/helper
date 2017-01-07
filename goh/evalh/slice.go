package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/strconvh"
	"go/ast"
	"reflect"
)

func invSliceOpError(x Value) *intError {
	return newIntError("cannot slice " + x.String() + " (type " + x.DeepType() + ")")
}
func invSliceIndexError(low, high int) *intError {
	return newIntError("invalid slice index: " + strconvh.FormatInt(low) + " > " + strconvh.FormatInt(high))
}
func invSlice3IndexOmitted() *intError {
	return newIntError("only first index in 3-index slice can be omitted")
}

const indexSkipped int = -1

// getSliceIndex returns index value (int).
// Returned index may be negative (-1) only if e is nil (this means that index is skipped in source).
// In all other cases negative index causes error.
func getSliceIndex(e ast.Expr, idents Identifiers) (r int, err error) {
	if e == nil {
		return indexSkipped, nil
	}

	v, err := astExpr(e, idents)
	if err != nil {
		return 0, err
	}

	r,_,ok:=v.Int()
	if !ok{
		return nil, convertUnableError(reflect.TypeOf(int(0)), v)
	}

	//switch v.Kind() {
	//case Regular:
	//	// Check kind
	//	if v.Regular().Kind() != reflect.Int {
	//		return 0, errors.New("unable to slicing using " + v.String())
	//	}
	//	// Check exact type
	//	var ok bool
	//	r, ok = v.Regular().Interface().(int)
	//	if !ok {
	//		return 0, errors.New("unable to slicing using " + v.String())
	//	}
	//case Untyped:
	//	var ok bool
	//	r, ok = constanth.IntVal(v.Untyped())
	//	if !ok {
	//		return 0, errors.New("unable to slicing using " + v.String())
	//	}
	//case Nil:
	//	return 0, errors.New("unable to slicing using Nil")
	//default:
	//	panic("unknown kind")	// TODO
	//}

	if r < 0 {
		return 0, errors.New("negative index value")
	}
	return
}

func slice2(x reflect.Value, low, high int) (r Value, err *intError) {
	// resolve pointer to array
	if x.Kind() == reflect.Ptr && x.Elem().Kind() == reflect.Array {
		x = x.Elem()
	}

	// check slicing possibility
	if k := x.Kind(); k != reflect.Array && k != reflect.Slice && k != reflect.String {
		return nil, invSliceOpError(MakeRegular(x))
	} else if k == reflect.Array && !x.CanAddr() {
		return nil, invSliceOpError(MakeRegular(x))
	}

	// resolve default value
	if low == indexSkipped {
		low = 0
	}
	if high == indexSkipped {
		high = x.Len()
	}

	// validate indexes
	switch {
	case low < 0:
		return nil, indexOutOfRange(low)
	case high > x.Len():
		return nil, indexOutOfRange(high)
	case low > high:
		return nil, invSliceIndexError(low, high)
	}

	return MakeRegular(x.Slice(low, high)), nil
}

func slice3(x reflect.Value, low, high, max int) (r Value, err *intError) {
	// resolve pointer to array
	if x.Kind() == reflect.Ptr && x.Elem().Kind() == reflect.Array {
		x = x.Elem()
	}

	// check slicing possibility
	if k := x.Kind(); k != reflect.Array && k != reflect.Slice {
		return nil, invSliceOpError(MakeRegular(x))
	} else if k == reflect.Array && !x.CanAddr() {
		return nil, invSliceOpError(MakeRegular(x))
	}

	// resolve default value
	if low == indexSkipped {
		low = 0
	}

	// validate indexes
	if high == indexSkipped || max == indexSkipped {
		return nil, invSlice3IndexOmitted()
	}
	switch {
	case low < 0:
		return nil, indexOutOfRange(low)
	case max > x.Cap():
		return nil, indexOutOfRange(high)
	case low > high:
		return nil, invSliceIndexError(low, high)
	case high > max:
		return nil, invSliceIndexError(high, max)
	}

	return MakeRegular(x.Slice3(low, high, max)), nil
}

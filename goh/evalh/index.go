package evalh

import (
	"github.com/apaxa-go/helper/strconvh"
	"go/constant"
	"reflect"
)

func invOpError(op, reason string) *intError {
	return newIntError("invalid operation: " + op + " (" + reason + ")")
}

func invIndexOpError(x Value, i Value) *intError {
	return invOpError(x.String()+"["+i.String()+"]", "type "+x.DeepType()+" does not support indexing")
}
func indexOutOfRangeError(i int) *intError {
	return newIntError("index " + strconvh.FormatInt(i) + " out of range")
}

func indexMap(x reflect.Value, i Value) (r Value, err *intError) {
	if k := x.Kind(); k != reflect.Map {
		return nil, invIndexOpError(MakeRegular(x), i)
	}

	iReqT := x.Type().Key()

	iV, _, ok := i.AsType(iReqT)
	if !ok {
		return nil, convertUnableError(iReqT, i)
	}
	// calc index
	//var iV reflect.Value
	//switch i.Kind() { TO DO extract to convertCall function
	//case Untyped:
	//	var ok bool
	//	iV, ok = constanth.AsType(i.Untyped(), iReqT)
	//	if !ok {
	//		return nil, convertUnableError(iReqT, i)
	//	}
	//case Regular:
	//	iV = i.Regular()
	//	if !iV.Type().AssignableTo(iReqT) {
	//		return nil, convertUnableError(iReqT, i)
	//	}
	//case Nil:
	//	return nil, convertUnableError(iReqT, i)
	//default:
	//	panic("unknown kind")
	//}

	//
	rV := x.MapIndex(iV)
	if !rV.IsValid() { // Return zero value if no such key in map
		return MakeRegular(reflect.New(x.Type().Elem()).Elem()), nil
	}

	return MakeRegular(rV), nil
}

func indexOther(x reflect.Value, i Value) (r Value, err *intError) {
	if k := x.Kind(); k != reflect.String && k != reflect.Array && k != reflect.Slice {
		return nil, invIndexOpError(MakeRegular(x), i)
	}

	iInt, _, ok := i.AsInt()
	if !ok {
		return nil, convertUnableError(reflect.TypeOf(int(0)), i)
	}

	// check out-of-range
	if iInt < 0 || iInt >= x.Len() {
		return nil, indexOutOfRangeError(iInt)
	}

	return MakeRegular(x.Index(iInt)), nil
}

func indexConstant(x constant.Value, i Value) (r Value, err *intError) {
	if x.Kind() != constant.String {
		return nil, invIndexOpError(MakeUntyped(x), i)
	}

	iInt, isUntyped, ok := i.AsInt()
	if !ok {
		return nil, convertUnableError(reflect.TypeOf(int(0)), i)
	}

	xStr := constant.StringVal(x)
	// check out-of-range
	if iInt < 0 || iInt >= len(xStr) {
		return nil, indexOutOfRangeError(iInt)
	}

	if isUntyped {
		return MakeUntypedInt64(int64(xStr[iInt])), nil
	}
	return MakeRegularInterface(byte(xStr[iInt])), nil
}

package evalh

import (
	"github.com/apaxa-go/helper/strconvh"
	"reflect"
)

func assign(to reflect.Value, value Value) (err *intError) {
	if !to.CanSet() {
		return assignDstUnsettableError(MakeRegular(to))
	}

	newValue, _, ok := value.AsType(to.Type())
	if !ok {
		return assignTypesMismError(to.Type(), value)
	}

	//var newValue reflect.Value
	//switch value.Kind() {
	//case Untyped:
	//	var ok bool
	//	newValue, ok = constanth.AsType(value.Untyped(), to.Type())
	//	if !ok {
	//		return assignTypesMismError(MakeRegular(to), value)
	//	}
	//case Regular:
	//	newValue = value.Regular()
	//case Nil:
	//	var tmp reflect.Value
	//	tmp, err = convertNil2(to.Type())
	//	if err != nil {
	//		return
	//	}
	//	if !tmp.Type().AssignableTo(to.Type()) {
	//		return assignTypesMismError(MakeRegular(to), value)
	//	}
	//	to.Set(tmp)
	//default:
	//	return assignTypesMismError(MakeRegular(to), value)
	//}
	//
	if !newValue.Type().AssignableTo(to.Type()) {
		return assignTypesMismError(to.Type(), value)
	}
	to.Set(newValue)
	return nil
}

func compositeLitStructKeys(t reflect.Type, elts map[string]Value) (r Value, err *intError) {
	if t.Kind() != reflect.Struct {
		return nil, compLitInvTypeError(t)
	}
	rV := reflect.New(t).Elem()

	for field, value := range elts {
		fV := rV.FieldByName(field)
		if !fV.IsValid() {
			return nil, compLitUnknFieldError(rV, field)
		}
		err = assign(fV, value)
		if err != nil {
			return
		}
	}

	return MakeRegular(rV), nil
}

func compositeLitStructOrdered(t reflect.Type, elts []Value) (r Value, err *intError) {
	if t.Kind() != reflect.Struct {
		return nil, compLitInvTypeError(t)
	}
	if t.NumField() != len(elts) {
		return nil, compLitArgsCountMismError(t.NumField(), len(elts))
	}
	rV := reflect.New(t).Elem()

	for i := range elts {
		fV := rV.Field(i)
		if !fV.IsValid() {
			return nil, compLitUnknFieldError(rV, strconvh.FormatInt(i))
		}
		err = assign(fV, elts[i])
		if err != nil {
			return
		}
	}

	return MakeRegular(rV), nil
}

func compositeLitArrayLike(t reflect.Type, elts map[int]Value) (r Value, err *intError) {
	// Calc max index (len of slice)
	var maxIndex = -1
	for i := range elts {
		if i < 0 {
			return nil, compLitNegIndexError()
		}
		if i > maxIndex {
			maxIndex = i
		}
	}

	// Init result
	var rV reflect.Value
	switch t.Kind() {
	case reflect.Array:
		rV = reflect.New(t).Elem()
		if maxIndex > rV.Len()-1 {
			return nil, compLitIndexOutOfBoundsError(rV.Len()-1, maxIndex)
		}
	case reflect.Slice:
		rV = reflect.MakeSlice(t, maxIndex+1, maxIndex+1)
	default:
		return nil, compLitInvTypeError(t)
	}

	// Fill result
	for i := range elts {
		iV := rV.Index(i)
		if !iV.IsValid() {
			return nil, compLitUnknFieldError(rV, strconvh.FormatInt(i))
		}
		err = assign(iV, elts[i])
		if err != nil {
			return
		}
	}

	return MakeRegular(rV), nil
}

func compositeLitMap(t reflect.Type, elts map[Value]Value) (r Value, err *intError) {
	if t.Kind() != reflect.Map {
		return nil, compLitInvTypeError(t)
	}

	rV := reflect.New(t).Elem()
	kV := reflect.New(t.Key()).Elem()
	vV := reflect.New(t.Elem()).Elem()
	for k, v := range elts {
		err = assign(kV, k)
		if err != nil {
			return
		}
		err = assign(vV, v)
		if err != nil {
			return
		}
		rV.SetMapIndex(kV, vV)
	}

	return MakeRegular(rV), nil
}

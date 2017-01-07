package evalh

import (
	"reflect"
)

func convertNilUnableError(t reflect.Type) *intError {
	return newIntError("cannot convertCall nil to type " + t.String())
}

func convertCall(t reflect.Type, args []Value) (r Value, err *intError) {
	if len(args) != 1 {
		return nil, convertMultError(t, args)	// TODO what if not multi, but no args
	}
	return convert(t,args[0])
}

func convert(t reflect.Type, x Value) (r Value, err *intError) {
	rV,_,ok:=x.ToType(t)
	if ok{
		r=MakeRegular(rV)
	}else{
		err=convertUnableError(t, x)
	}
	return
}
//
//func convertNil2(t reflect.Type) (r reflect.Value, err *intError) {
//	switch t.Kind() {
//	case reflect.Slice, reflect.Ptr, reflect.Func, reflect.Interface, reflect.Map, reflect.Chan:
//		return reflect.New(t), nil // TODO check if result is adequate
//	default:
//		return r, convertNilUnableError(t)
//	}
//}

package constanth

import (
	"go/constant"
	"reflect"
)

func SameKind(x constant.Value, kind reflect.Kind) (r interface{}, ok bool) {
	switch kind {
	case reflect.Bool:
		return BoolVal(x)
	case reflect.Int:
		return IntVal(x)
	case reflect.Int8:
		return Int8Val(x)
	case reflect.Int16:
		return Int16Val(x)
	case reflect.Int32:
		return Int32Val(x)
	case reflect.Int64:
		return Int64Val(x)
	case reflect.Uint:
		return UintVal(x)
	case reflect.Uint8:
		return Uint8Val(x)
	case reflect.Uint16:
		return Uint16Val(x)
	case reflect.Uint32:
		return Uint32Val(x)
	case reflect.Uint64:
		return Uint64Val(x)
	case reflect.Float32:
		return Float32Val(x)
	case reflect.Float64:
		return Float64Val(x)
	case reflect.Complex64:
		return Complex64Val(x)
	case reflect.Complex128:
		return Complex128Val(x)
	case reflect.String:
		return StringVal(x)
	default:
		return nil, false
	}
}

func SameType(x constant.Value, t reflect.Type) (r reflect.Value, ok bool) {
	r = reflect.New(t).Elem()
	switch t.Kind() {
	case reflect.Bool:
		var v bool
		v, ok = BoolVal(x)
		if ok {
			r.SetBool(v)
		}
	case reflect.Int:
		var v int
		v, ok = IntVal(x)
		if ok {
			r.SetInt(int64(v))
		}
	case reflect.Int8:
		var v int8
		v, ok = Int8Val(x)
		if ok {
			r.SetInt(int64(v))
		}
	case reflect.Int16:
		var v int16
		v, ok = Int16Val(x)
		if ok {
			r.SetInt(int64(v))
		}
	case reflect.Int32:
		var v int32
		v, ok = Int32Val(x)
		if ok {
			r.SetInt(int64(v))
		}
	case reflect.Int64:
		var v int64
		v, ok = Int64Val(x)
		if ok {
			r.SetInt(v)
		}
	case reflect.Uint:
		var v uint
		v, ok = UintVal(x)
		if ok {
			r.SetUint(uint64(v))
		}
	case reflect.Uint8:
		var v uint8
		v, ok = Uint8Val(x)
		if ok {
			r.SetUint(uint64(v))
		}
	case reflect.Uint16:
		var v uint16
		v, ok = Uint16Val(x)
		if ok {
			r.SetUint(uint64(v))
		}
	case reflect.Uint32:
		var v uint32
		v, ok = Uint32Val(x)
		if ok {
			r.SetUint(uint64(v))
		}
	case reflect.Uint64:
		var v uint64
		v, ok = Uint64Val(x)
		if ok {
			r.SetUint(v)
		}
	case reflect.Float32:
		var v float32
		v, ok = Float32Val(x)
		if ok {
			r.SetFloat(float64(v))
		}
	case reflect.Float64:
		var v float64
		v, ok = Float64Val(x)
		if ok {
			r.SetFloat(v)
		}
	case reflect.Complex64:
		var v complex64
		v, ok = Complex64Val(x)
		if ok {
			r.SetComplex(complex128(v))
		}
	case reflect.Complex128:
		var v complex128
		v, ok = Complex128Val(x)
		if ok {
			r.SetComplex(v)
		}
	case reflect.String:
		var v string
		v, ok = StringVal(x)
		if ok {
			r.SetString(v)
		}
	default:
		ok = false
	}
	return
}

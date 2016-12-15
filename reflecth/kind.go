package reflecth

import "reflect"

func IsInt(k reflect.Kind) bool {
	return k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64
}

func IsUint(k reflect.Kind) bool {
	return k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64
}

func IsAnyInt(k reflect.Kind) bool {
	return IsInt(k) || IsUint(k)
}

func IsFloat(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

func IsComplex(k reflect.Kind) bool {
	return k == reflect.Complex64 || k == reflect.Complex128
}

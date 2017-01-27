package reflecth

import (
	"reflect"
	"testing"
)

func TestTypes(t *testing.T) {
	ok := TypeBool() == reflect.TypeOf(true) && TypeString() == reflect.TypeOf("") && TypeInt() == reflect.TypeOf(int(0)) && TypeInt8() == reflect.TypeOf(int8(0)) && TypeInt16() == reflect.TypeOf(int16(0)) && TypeInt32() == reflect.TypeOf(int32(0)) && TypeInt64() == reflect.TypeOf(int64(0)) && TypeUint() == reflect.TypeOf(uint(0)) && TypeUint8() == reflect.TypeOf(uint8(0)) && TypeUint16() == reflect.TypeOf(uint16(0)) && TypeUint32() == reflect.TypeOf(uint32(0)) && TypeUint64() == reflect.TypeOf(uint64(0)) && TypeFloat32() == reflect.TypeOf(float32(0)) && TypeFloat64() == reflect.TypeOf(float64(0)) && TypeComplex64() == reflect.TypeOf(complex64(0)) && TypeComplex128() == reflect.TypeOf(complex128(0)) && TypeByte() == reflect.TypeOf(byte(0)) && TypeRune() == reflect.TypeOf(rune(0)) && TypeUintptr() == reflect.TypeOf(uintptr(0)) && TypeEmptyInterface() == reflect.TypeOf(interface{}(nil))
	if !ok {
		t.Error("error in Type*")
	}
}

package evalh

import "reflect"

func isBuiltInType(ident string) bool {
	_, ok := builtInTypes[ident]
	return ok
}

var builtInTypes = map[string]reflect.Type{
	"bool": reflect.TypeOf(false),

	//////////////// Numeric ////////////////
	"uint":   reflect.TypeOf(uint(0)),
	"uint8":  reflect.TypeOf(uint8(0)),
	"uint16": reflect.TypeOf(uint16(0)),
	"uint32": reflect.TypeOf(uint32(0)),
	"uint64": reflect.TypeOf(uint64(0)),

	"int":   reflect.TypeOf(int(0)),
	"int8":  reflect.TypeOf(int8(0)),
	"int16": reflect.TypeOf(int16(0)),
	"int32": reflect.TypeOf(int32(0)),
	"int64": reflect.TypeOf(int64(0)),

	"float32": reflect.TypeOf(float32(0)),
	"float64": reflect.TypeOf(float64(0)),

	"complex64":  reflect.TypeOf(complex64(0)),
	"complex128": reflect.TypeOf(complex128(0)),

	"byte": reflect.TypeOf(byte(0)),
	"rune": reflect.TypeOf(rune(0)),

	"uintptr": reflect.TypeOf(uintptr(0)),
	/////////////////////////////////////////

	"string": reflect.TypeOf(""),
}

// In some places required specific types. This variables allow to avoid using types map.
var (
	stringT     = reflect.TypeOf(string(""))
	bytesSliceT = reflect.SliceOf(reflect.TypeOf(byte(0)))
)

package reflecth

import (
	"reflect"
	"unsafe"
)

// Set assigns src to the value dst.
// It is similar to dst.Set(src) but this function also allow to set private fields.
// Primary reason for this is to avoid restriction with your own struct variable.
// It panics if CanAddr returns false.
// As in Go, x's value must be assignable to v's type.
//
// Deprecated: In any case it is bad practice to change private fields in 3rd party variables/classes.
func Set(dst, src reflect.Value) {
	switch dst.CanSet() {
	case false:
		addr := dst.UnsafeAddr()
		dst = reflect.NewAt(dst.Type(), unsafe.Pointer(addr)).Elem()
		fallthrough
	case true:
		dst.Set(src)
	}
	return
}

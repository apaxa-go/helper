package reflecth

import (
	"reflect"
	"testing"
)

func TestIs(t *testing.T) {
	type testElement struct {
		k                                           reflect.Kind
		isInt, isUint, isAnyInt, isFloat, isComplex bool
	}

	tests := []testElement{
		{reflect.Int, true, false, true, false, false},
		{reflect.Int8, true, false, true, false, false},
		{reflect.Int16, true, false, true, false, false},
		{reflect.Int32, true, false, true, false, false},
		{reflect.Int64, true, false, true, false, false},
		{reflect.Uint, false, true, true, false, false},
		{reflect.Uint8, false, true, true, false, false},
		{reflect.Uint16, false, true, true, false, false},
		{reflect.Uint32, false, true, true, false, false},
		{reflect.Uint64, false, true, true, false, false},
		{reflect.Float32, false, false, false, true, false},
		{reflect.Float64, false, false, false, true, false},
		{reflect.Complex64, false, false, false, false, true},
		{reflect.Complex128, false, false, false, false, true},
	}

	for _, test := range tests {
		ok := IsInt(test.k) == test.isInt && IsUint(test.k) == test.isUint && IsAnyInt(test.k) == test.isAnyInt && IsFloat(test.k) == test.isFloat && IsComplex(test.k) == test.isComplex
		if !ok {
			t.Errorf("%v: invalid Is*", test.k)
		}
	}
}

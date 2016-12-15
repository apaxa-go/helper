package constanth

import (
	"github.com/apaxa-go/helper/mathh"
	"go/constant"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE

func BoolVal(x constant.Value) (r bool, ok bool) {
	if x.Kind() != constant.Bool {
		return false, false
	}
	return constant.BoolVal(x), true
}

func Int64Val(x constant.Value) (int64, bool) {
	if x.Kind() != constant.Int {
		return 0, false
	}
	return constant.Int64Val(x)
}

func Uint64Val(x constant.Value) (uint64, bool) {
	if x.Kind() != constant.Int {
		return 0, false
	}
	return constant.Uint64Val(x)
}

func Float32Val(x constant.Value) (float32, bool) {
	if x.Kind() != constant.Float {
		return 0, false
	}
	return constant.Float32Val(x)
}

func Float64Val(x constant.Value) (float64, bool) {
	if x.Kind() != constant.Float {
		return 0, false
	}
	return constant.Float64Val(x)
}

func Complex64Val(x constant.Value) (complex64, bool) {
	if x.Kind() != constant.Complex {
		return 0, false
	}
	realC := constant.Real(x)
	imagC := constant.Imag(x)
	r, ok := Float32Val(realC)
	if !ok {
		return 0, false
	}
	i, ok := Float32Val(imagC)
	if !ok {
		return 0, false
	}
	return complex(r, i), true
}

func Complex128Val(x constant.Value) (complex128, bool) {
	if x.Kind() != constant.Complex {
		return 0, false
	}
	realC := constant.Real(x)
	imagC := constant.Imag(x)
	r, ok := Float64Val(realC)
	if !ok {
		return 0, false
	}
	i, ok := Float64Val(imagC)
	if !ok {
		return 0, false
	}
	return complex(r, i), true
}

func StringVal(x constant.Value) (string, bool) {
	if x.Kind() != constant.String {
		return "", false
	}
	return constant.StringVal(x)
}

//replacer:replace
//replacer:old int32	Int32
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16

func Int32Val(x constant.Value) (int32, bool) {
	i64, ok := Int64Val(x)
	if !ok {
		return 0, false
	}
	if i64 < mathh.MinInt32 || i64 > mathh.MaxInt32 {
		return 0, false
	}
	return int32(i64), true
}

func Uint32Val(x constant.Value) (uint32, bool) {
	u64, ok := Uint64Val(x)
	if !ok {
		return 0, false
	}
	if u64 > mathh.MaxUint32 {
		return 0, false
	}
	return uint32(u64), true
}

package constanth

import (
	"go/constant"
	"go/token"
)

func MakeComplex128(x complex128) constant.Value {
	return constant.BinaryOp(constant.MakeFloat64(real(x)), token.ADD, constant.MakeImag(constant.MakeFloat64(imag(x))))
}

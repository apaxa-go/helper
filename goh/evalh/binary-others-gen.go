//replacer:generated-file

package evalh

import (
	"errors"
	"go/token"
)

func binaryOtherInt(x int, op token.Token, y int) (r int, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherInt8(x int8, op token.Token, y int8) (r int8, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherInt16(x int16, op token.Token, y int16) (r int16, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherInt32(x int32, op token.Token, y int32) (r int32, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherUint(x uint, op token.Token, y uint) (r uint, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherUint8(x uint8, op token.Token, y uint8) (r uint8, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherUint16(x uint16, op token.Token, y uint16) (r uint16, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherUint32(x uint32, op token.Token, y uint32) (r uint32, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherUint64(x uint64, op token.Token, y uint64) (r uint64, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	case token.REM:
		return x % y, nil
	case token.AND:
		return x & y, nil
	case token.OR:
		return x | y, nil
	case token.XOR:
		return x ^ y, nil
	case token.AND_NOT:
		return x &^ y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherFloat64(x float64, op token.Token, y float64) (r float64, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherComplex64(x complex64, op token.Token, y complex64) (r complex64, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryOtherComplex128(x complex128, op token.Token, y complex128) (r complex128, err *intError) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	default:
		return 0, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

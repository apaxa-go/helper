//replacer:generated-file

package evalh

import (
	"errors"
	"go/token"
)

func binaryOtherInt(x int, op token.Token, y int) (r int, err error) {
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
		return 0, errors.New("int operands does not support operation " + op.String())
	}
}

func binaryOtherInt8(x int8, op token.Token, y int8) (r int8, err error) {
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
		return 0, errors.New("int8 operands does not support operation " + op.String())
	}
}

func binaryOtherInt16(x int16, op token.Token, y int16) (r int16, err error) {
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
		return 0, errors.New("int16 operands does not support operation " + op.String())
	}
}

func binaryOtherInt32(x int32, op token.Token, y int32) (r int32, err error) {
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
		return 0, errors.New("int32 operands does not support operation " + op.String())
	}
}

func binaryOtherUint(x uint, op token.Token, y uint) (r uint, err error) {
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
		return 0, errors.New("uint operands does not support operation " + op.String())
	}
}

func binaryOtherUint8(x uint8, op token.Token, y uint8) (r uint8, err error) {
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
		return 0, errors.New("uint8 operands does not support operation " + op.String())
	}
}

func binaryOtherUint16(x uint16, op token.Token, y uint16) (r uint16, err error) {
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
		return 0, errors.New("uint16 operands does not support operation " + op.String())
	}
}

func binaryOtherUint32(x uint32, op token.Token, y uint32) (r uint32, err error) {
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
		return 0, errors.New("uint32 operands does not support operation " + op.String())
	}
}

func binaryOtherUint64(x uint64, op token.Token, y uint64) (r uint64, err error) {
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
		return 0, errors.New("uint64 operands does not support operation " + op.String())
	}
}

func binaryOtherFloat64(x float64, op token.Token, y float64) (r float64, err error) {
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
		return 0, errors.New("float64 operands does not support operation " + op.String())
	}
}

func binaryOtherComplex64(x complex64, op token.Token, y complex64) (r complex64, err error) {
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
		return 0, errors.New("complex64 operands does not support operation " + op.String())
	}
}

func binaryOtherComplex128(x complex128, op token.Token, y complex128) (r complex128, err error) {
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
		return 0, errors.New("complex128 operands does not support operation " + op.String())
	}
}

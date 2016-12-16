package evalh

import (
	"errors"
	"go/token"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE

func binaryOtherBool(x bool, op token.Token, y bool) (r bool, err error) {
	switch op {
	case token.LAND:
		return x && y, nil
	case token.LOR:
		return x || y, nil
	default:
		return false, errors.New("boolean operands does not support operation " + op.String())
	}
}

func binaryOtherString(x string, op token.Token, y string) (r string, err error) {
	switch op {
	case token.ADD:
		return x + y, nil
	default:
		return "", errors.New("strings operands does not support operation " + op.String())
	}
}

//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32
//replacer:new uint64	Uint64

func binaryOtherInt64(x int64, op token.Token, y int64) (r int64, err error) {
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
		return 0, errors.New("int64 operands does not support operation " + op.String())
	}
}

//replacer:replace
//replacer:old float32		Float32
//replacer:new float64		Float64
//replacer:new complex64	Complex64
//replacer:new complex128	Complex128

func binaryOtherFloat32(x float32, op token.Token, y float32) (r float32, err error) {
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
		return 0, errors.New("float32 operands does not support operation " + op.String())
	}
}

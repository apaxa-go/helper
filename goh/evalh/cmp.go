package evalh

import (
	"errors"
	"go/constant"
	"go/token"
)

func binaryCompareBool(x bool, op token.Token, y bool) (r interface{}, err error) {
	switch op {
	case token.EQL:
		return constant.MakeBool(x == y), nil
	case token.NEQ:
		return constant.MakeBool(x != y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on booleans")
	}
}

func binaryCompareInt(x int64, op token.Token, y int64) (r interface{}, err error) {
	switch op {
	case token.EQL:
		return constant.MakeBool(x == y), nil
	case token.NEQ:
		return constant.MakeBool(x != y), nil
	case token.LSS:
		return constant.MakeBool(x < y), nil
	case token.LEQ:
		return constant.MakeBool(x <= y), nil
	case token.GTR:
		return constant.MakeBool(x > y), nil
	case token.GEQ:
		return constant.MakeBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on signed integers")
	}
}

func binaryCompareUint(x uint64, op token.Token, y uint64) (r interface{}, err error) {
	switch op {
	case token.EQL:
		return constant.MakeBool(x == y), nil
	case token.NEQ:
		return constant.MakeBool(x != y), nil
	case token.LSS:
		return constant.MakeBool(x < y), nil
	case token.LEQ:
		return constant.MakeBool(x <= y), nil
	case token.GTR:
		return constant.MakeBool(x > y), nil
	case token.GEQ:
		return constant.MakeBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on unsigned integers")
	}
}

func binaryCompareFloat(x float64, op token.Token, y float64) (r interface{}, err error) {
	switch op {
	case token.EQL:
		return constant.MakeBool(x == y), nil
	case token.NEQ:
		return constant.MakeBool(x != y), nil
	case token.LSS:
		return constant.MakeBool(x < y), nil
	case token.LEQ:
		return constant.MakeBool(x <= y), nil
	case token.GTR:
		return constant.MakeBool(x > y), nil
	case token.GEQ:
		return constant.MakeBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on floats")
	}
}

func binaryCompareComplex(x complex128, op token.Token, y complex128) (r interface{}, err error) {
	switch op {
	case token.EQL:
		return constant.MakeBool(x == y), nil
	case token.NEQ:
		return constant.MakeBool(x != y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on complexes")
	}
}

func binaryCompareString(x string, op token.Token, y string) (r interface{}, err error) {
	switch op {
	case token.EQL:
		return constant.MakeBool(x == y), nil
	case token.NEQ:
		return constant.MakeBool(x != y), nil
	case token.LSS:
		return constant.MakeBool(x < y), nil
	case token.LEQ:
		return constant.MakeBool(x <= y), nil
	case token.GTR:
		return constant.MakeBool(x > y), nil
	case token.GEQ:
		return constant.MakeBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on strings")
	}
}

func binaryComparePointer(x uintptr, op token.Token, y uintptr) (r interface{}, err error) {
	switch op {
	case token.EQL:
		return constant.MakeBool(x == y), nil
	case token.NEQ:
		return constant.MakeBool(x != y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on pointers")
	}
}

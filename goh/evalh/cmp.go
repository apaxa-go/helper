package evalh

import (
	"errors"
	"go/token"
)

func binaryCompareBool(x bool, op token.Token, y bool) (r Value, err error) {
	switch op {
	case token.EQL:
		return MakeUntypedBool(x == y), nil
	case token.NEQ:
		return MakeUntypedBool(x != y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on booleans")
	}
}

func binaryCompareInt(x int64, op token.Token, y int64) (r Value, err error) {
	switch op {
	case token.EQL:
		return MakeUntypedBool(x == y), nil
	case token.NEQ:
		return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on signed integers")
	}
}

func binaryCompareUint(x uint64, op token.Token, y uint64) (r Value, err error) {
	switch op {
	case token.EQL:
		return MakeUntypedBool(x == y), nil
	case token.NEQ:
		return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on unsigned integers")
	}
}

func binaryCompareFloat(x float64, op token.Token, y float64) (r Value, err error) {
	switch op {
	case token.EQL:
		return MakeUntypedBool(x == y), nil
	case token.NEQ:
		return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on floats")
	}
}

func binaryCompareComplex(x complex128, op token.Token, y complex128) (r Value, err error) {
	switch op {
	case token.EQL:
		return MakeUntypedBool(x == y), nil
	case token.NEQ:
		return MakeUntypedBool(x != y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on complexes")
	}
}

func binaryCompareString(x string, op token.Token, y string) (r Value, err error) {
	switch op {
	case token.EQL:
		return MakeUntypedBool(x == y), nil
	case token.NEQ:
		return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on strings")
	}
}

func binaryComparePointer(x uintptr, op token.Token, y uintptr) (r Value, err error) {
	switch op {
	case token.EQL:
		return MakeUntypedBool(x == y), nil
	case token.NEQ:
		return MakeUntypedBool(x != y), nil
	default:
		return nil, errors.New("unable to perform comparison '" + op.String() + "' on pointers")
	}
}

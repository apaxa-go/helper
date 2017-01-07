package evalh

import (
	"go/token"
	"reflect"
)

func binaryCompareWithNil(x reflect.Value, op token.Token) (r Value, err *intError) {
	var equality bool // true if op == "=="
	switch op {
	case token.EQL:
		equality = false
	case token.NEQ:
		equality = true
	default:
		return nil, newIntErrorf(invBinOp, x.String(), op.String(), "nil", "operator "+op.String()+" not defined on nil")
	}

	switch x.Kind() {
	case reflect.Slice, reflect.Map, reflect.Func, reflect.Ptr, reflect.Chan, reflect.Interface:
		return MakeUntypedBool(x.IsNil() == equality), nil
	default:
		return nil, newIntErrorf(invBinOp, x.String(), op.String(), "nil", "compare with nil is not defined on "+x.Type().String())
	}
}

//func binaryCompareBool(x bool, op token.Token, y bool) (r Value, err error) {
//	switch op {
//	case token.EQL:
//		return MakeUntypedBool(x == y), nil
//	case token.NEQ:
//		return MakeUntypedBool(x != y), nil
//	default:
//		return nil, errors.New("unable to perform comparison '" + op.String() + "' on booleans")
//	}
//}

func binaryCompareInt(x int64, op token.Token, y int64) (r Value, err *intError) {
	switch op {
	//case token.EQL:
	//	return MakeUntypedBool(x == y), nil
	//case token.NEQ:
	//	return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryCompareUint(x uint64, op token.Token, y uint64) (r Value, err *intError) {
	switch op {
	//case token.EQL:
	//	return MakeUntypedBool(x == y), nil
	//case token.NEQ:
	//	return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

func binaryCompareFloat(x float64, op token.Token, y float64) (r Value, err *intError) {
	switch op {
	//case token.EQL:
	//	return MakeUntypedBool(x == y), nil
	//case token.NEQ:
	//	return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

//func binaryCompareComplex(x complex128, op token.Token, y complex128) (r Value, err error) {
//	switch op {
//	case token.EQL:
//		return MakeUntypedBool(x == y), nil
//	case token.NEQ:
//		return MakeUntypedBool(x != y), nil
//	default:
//		return nil, errors.New("unable to perform comparison '" + op.String() + "' on complexes")
//	}
//}

func binaryCompareString(x string, op token.Token, y string) (r Value, err *intError) {
	switch op {
	//case token.EQL:
	//	return MakeUntypedBool(x == y), nil
	//case token.NEQ:
	//	return MakeUntypedBool(x != y), nil
	case token.LSS:
		return MakeUntypedBool(x < y), nil
	case token.LEQ:
		return MakeUntypedBool(x <= y), nil
	case token.GTR:
		return MakeUntypedBool(x > y), nil
	case token.GEQ:
		return MakeUntypedBool(x >= y), nil
	default:
		return nil, invBinOpInvalError(MakeRegularInterface(x), op, MakeRegularInterface(y))
	}
}

//func binaryComparePointer(x uintptr, op token.Token, y uintptr) (r Value, err error) {
//	switch op {
//	case token.EQL:
//		return MakeUntypedBool(x == y), nil
//	case token.NEQ:
//		return MakeUntypedBool(x != y), nil
//	default:
//		return nil, errors.New("unable to perform comparison '" + op.String() + "' on pointers")
//	}
//}

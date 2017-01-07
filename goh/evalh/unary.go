package evalh

import (
	"github.com/apaxa-go/helper/reflecth"
	"go/constant"
	"go/token"
	"reflect"
)

func invUnaryOp(x Value, op token.Token) *intError {
	return newIntError("invalid operation: " + op.String() + " " + x.DeepType())
}
func invUnaryOpReason(x Value, op token.Token, reason interface{}) *intError {
	return newIntErrorf("invalid operation: %v %v: %v", op.String(), x.DeepType(), reason)
}

func unary(x reflect.Value, op token.Token) (r Value, err *intError) {
	switch op {
	case token.SUB:
		return unarySub(x)
	case token.XOR:
		return unaryXor(x)
	case token.NOT:
		return unaryNot(x)
	case token.AND:
		return unaryAnd(x)
	case token.ADD:
		if k := x.Kind(); reflecth.IsAnyInt(k) || reflecth.IsAnyFloat(k) || reflecth.IsAnyComplex(k) {
			return MakeRegular(x), nil
		}
		fallthrough
	default:
		return nil, invUnaryOp(MakeRegular(x), op)
	}
}

func unaryAnd(x reflect.Value) (r Value, err *intError) {
	if x.CanAddr() {
		return MakeRegular(x.Addr()), nil
	}

	rV := reflect.New(x.Type())
	rV.Elem().Set(x)
	return MakeRegular(rV), nil
}

func unarySub(x reflect.Value) (r Value, err *intError) {
	rV := reflect.New(x.Type()).Elem()

	switch k := x.Kind(); {
	case reflecth.IsInt(k):
		rV.SetInt(-x.Int()) // looks like overflow correct (see tests)
	case reflecth.IsAnyFloat(k):
		rV.SetFloat(-x.Float())
	case reflecth.IsAnyComplex(k):
		rV.SetComplex(-x.Complex())
	default:
		return nil, invUnaryOp(MakeRegular(x), token.SUB)
	}
	return MakeRegular(rV), nil
}

func unaryNot(x reflect.Value) (r Value, err *intError) {
	rV := reflect.New(x.Type()).Elem()

	if k := x.Kind(); k != reflect.Bool {
		return nil, invUnaryOp(MakeRegular(x), token.NOT)
	}

	rV.SetBool(!x.Bool())
	return MakeRegular(rV), nil
}

func unaryXor(x reflect.Value) (r Value, err *intError) {
	rV := reflect.New(x.Type()).Elem()

	switch k := x.Kind(); k {
	case reflect.Int:
		rV.SetInt(int64(^int(x.Int())))
	case reflect.Int8:
		rV.SetInt(int64(^int8(x.Int())))
	case reflect.Int16:
		rV.SetInt(int64(^int16(x.Int())))
	case reflect.Int32:
		rV.SetInt(int64(^int32(x.Int())))
	case reflect.Int64:
		rV.SetInt(^x.Int())
	case reflect.Uint:
		rV.SetUint(uint64(^uint(x.Uint())))
	case reflect.Uint8:
		rV.SetUint(uint64(^uint8(x.Uint())))
	case reflect.Uint16:
		rV.SetUint(uint64(^uint16(x.Uint())))
	case reflect.Uint32:
		rV.SetUint(uint64(^uint32(x.Uint())))
	case reflect.Uint64:
		rV.SetUint(^x.Uint())
	default:
		return nil, invUnaryOp(MakeRegular(x), token.XOR)
	}

	return MakeRegular(rV), nil
}

func unaryConstant(x constant.Value, op token.Token) (r Value, err *intError) {
	defer func() {
		if rec := recover(); rec != nil {
			r = nil
			err = invUnaryOpReason(MakeUntyped(x), op, rec)
		}
	}()
	v := constant.UnaryOp(op, x, 0) // TODO prec should be set?
	if v.Kind() == constant.Unknown {
		return nil, invUnaryOp(MakeUntyped(x), op)
	}
	return MakeUntyped(v), nil
}

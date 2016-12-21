package evalh

import (
	"errors"
	"fmt"
	"github.com/apaxa-go/helper/reflecth"
	"go/constant"
	"go/token"
	"reflect"
)

func unary(x Value, op token.Token) (r Value, err error) {
	// Perform calc separately if arg is untyped constant
	if x.Kind() == Untyped {
		return unaryConstant(x.Untyped(), op)
	}

	if x.Kind() != Regular {
		return nil, errors.New("unable to perform unary operation on " + x.String())
	}

	switch op {
	case token.ADD:
		if k := x.Regular().Kind(); reflecth.IsAnyInt(k) || reflecth.IsFloat(k) || reflecth.IsComplex(k) {
			return x, nil
		}
	case token.SUB:
		return unarySub(x.Regular())
	case token.XOR:
		return unaryXor(x.Regular())
	case token.NOT:
		return unaryNot(x.Regular())
	case token.AND:
		return unaryAnd(x.Regular())
	}

	return nil, errors.New("unable to perform unary operation " + op.String() + " on " + x.String())
}

func unaryAnd(x reflect.Value) (r Value, err error) {
	if x.CanAddr() {
		return MakeRegular(x.Addr()), nil
	}

	rV := reflect.New(x.Type())
	rV.Elem().Set(x)
	return MakeRegular(rV), nil
}

func unarySub(x reflect.Value) (r Value, err error) {
	rV := reflect.New(x.Type()).Elem()

	switch k := x.Kind(); {
	case reflecth.IsInt(k):
		rV.SetInt(-x.Int()) // TODO possible wrong overflow
	case reflecth.IsFloat(k):
		rV.SetFloat(-x.Float())
	case reflecth.IsComplex(k):
		rV.SetComplex(-x.Complex())
	default:
		return nil, errors.New("unable to negate " + x.String())
	}
	return MakeRegular(rV), nil
}

func unaryNot(x reflect.Value) (r Value, err error) {
	rV := reflect.New(x.Type()).Elem()

	if k := x.Kind(); k != reflect.Bool {
		return nil, errors.New("unable to not " + k.String())
	}

	rV.SetBool(!x.Bool())
	return MakeRegular(rV), nil
}

func unaryXor(x reflect.Value) (r Value, err error) {
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
		return nil, errors.New("unable to unary xor " + k.String())
	}

	return MakeRegular(rV), nil
}

func unaryConstant(x constant.Value, op token.Token) (r Value, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if str, ok := rec.(string); ok {
				err = errors.New(str)
			} else {
				err = errors.New(fmt.Sprint(rec))
			}
		}
	}()
	v := constant.UnaryOp(op, x, 0) // TODO what 0 does mean?
	if v.Kind() == constant.Unknown {
		return nil, errors.New("unable to perform unary operation on untyped constant: unknown result")
	}
	return MakeUntyped(v), nil
}

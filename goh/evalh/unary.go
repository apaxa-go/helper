package evalh

import (
	"errors"
	"fmt"
	"github.com/apaxa-go/helper/reflecth"
	"go/constant"
	"go/token"
	"reflect"
)

func unary(x interface{}, op token.Token) (r interface{}, err error) {
	xConst, xConstOk := x.(constant.Value)

	// Perform calc separately if both args is untyped constants
	if xConstOk {
		return unaryConstant(xConst, op)
	}

	switch op {
	case token.ADD:
		xV := reflect.ValueOf(x)
		if k := xV.Kind(); reflecth.IsAnyInt(k) || reflecth.IsFloat(k) || reflecth.IsComplex(k) {
			return x, nil
		}
	case token.SUB:
		return unarySub(x)
	case token.XOR:
		return unaryXor(x)
	case token.NOT:
		return unaryNot(x)
	case token.AND:
		return unaryAnd(x)
	}

	return nil, errors.New("unable to perform unary operation " + op.String() + " on " + reflect.TypeOf(x).String())
}

func unaryAnd(x interface{}) (r interface{}, err error) {
	xV := reflect.ValueOf(x)
	if xV.CanAddr() {
		return xV.Addr(), nil
	}

	rV := reflect.New(xV.Type())
	rV.Elem().Set(xV)
	return rV, nil
}

func unarySub(x interface{}) (r interface{}, err error) {
	xV := reflect.ValueOf(x)
	rV := reflect.New(xV.Type()).Elem()

	switch k := xV.Kind(); {
	case reflecth.IsInt(k):
		rV.SetInt(-xV.Int())
	case reflecth.IsFloat(k):
		rV.SetFloat(-xV.Float())
	case reflecth.IsComplex(k):
		rV.SetComplex(-xV.Complex())
	default:
		return nil, errors.New("unable to negate " + xV.Kind().String())
	}
	return rV.Interface(), nil
}

func unaryNot(x interface{}) (r interface{}, err error) {
	xV := reflect.ValueOf(x)
	rV := reflect.New(xV.Type()).Elem()

	if k := xV.Kind(); k != reflect.Bool {
		return nil, errors.New("unable to not " + k.String())
	}

	rV.Set(!xV.Bool())
	return rV.Interface(), nil
}

func unaryXor(x interface{}) (r interface{}, err error) {
	xV := reflect.ValueOf(x)
	rV := reflect.New(xV.Type()).Elem()

	switch k := xV.Kind(); k {
	case reflect.Int:
		rV.SetInt(int64(^int(xV.Int())))
	case reflect.Int8:
		rV.SetInt(int64(^int8(xV.Int())))
	case reflect.Int16:
		rV.SetInt(int64(^int16(xV.Int())))
	case reflect.Int32:
		rV.SetInt(int64(^int32(xV.Int())))
	case reflect.Int64:
		rV.SetInt(^xV.Int())
	case reflect.Uint:
		rV.SetUint(uint64(^uint(xV.Uint())))
	case reflect.Uint8:
		rV.SetUint(uint64(^uint8(xV.Uint())))
	case reflect.Uint16:
		rV.SetUint(uint64(^uint16(xV.Uint())))
	case reflect.Uint32:
		rV.SetUint(uint64(^uint32(xV.Uint())))
	case reflect.Uint64:
		rV.SetUint(^xV.Uint())
	default:
		return nil, errors.New("unable to unary xor " + k.String())
	}

	return rV.Interface(), nil
}

func unaryConstant(x constant.Value, op token.Token) (r interface{}, err error) {
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
	return v, nil
}

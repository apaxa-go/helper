package evalh

import (
	"errors"
	"fmt"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/strconvh"
	"reflect"
)

func callRegular(f reflect.Value, args []Value) (r Value, err error) {
	if f.Kind() != reflect.Func {
		return nil, errors.New("no such function " + f.String())
	}

	// Check in/out arguments count
	fT := f.Type()
	if fT.NumIn() != len(args) {
		return nil, errors.New("required " + strconvh.FormatInt(fT.NumIn()) + " but got " + strconvh.FormatInt(len(args)))
	}
	if fT.NumOut() != 1 {
		return nil, errors.New("function must return exactly 1 parameter but returns " + strconvh.FormatInt(fT.NumOut()))
	}

	// Prepare arguments
	typedArgs := make([]reflect.Value, fT.NumIn())
	for i := range args {
		switch args[i].Kind() {
		case Untyped:
			var ok bool
			typedArgs[i], ok = constanth.SameType(args[i].Untyped(), fT.In(i))
			if !ok {
				return nil, errors.New("cannot convert argument " + strconvh.FormatInt(i) + " " + args[i].String() + " (untyped constant) to required type " + fT.In(i).String())
			}
		case Nil:
			typedArgs[i] = reflect.ValueOf(nil) // TODO check this, may be reflect.ValueOf(nil)
		case Regular:
			if aT := args[i].Regular().Type(); !aT.AssignableTo(fT.In(i)) {
				return nil, errors.New("cannot convert argument " + strconvh.FormatInt(i) + " " + aT.String() + " to required type " + fT.In(i).String())
			}
			typedArgs[i] = args[i].Regular()
		default:
			return nil, errors.New("cannot convert argument " + strconvh.FormatInt(i) + " " + args[i].String() + " to required type " + fT.In(i).String())
		}
	}

	defer func() {
		if rec := recover(); rec != nil {
			if str, ok := rec.(string); ok {
				err = errors.New(str)
			} else {
				err = errors.New(fmt.Sprint(rec))
			}
		}
	}()
	rs := f.Call(typedArgs)
	return MakeRegular(rs[0]), nil
}

func callType(t reflect.Type, args []Value) (r Value, err error) {
	if len(args) != 1 {
		return nil, errors.New("type conversion requires exactly one argument")
	}

	switch args[0].Kind() { // TODO Nil?
	case Untyped:
		return convertUntyped(t, args[0].Untyped())
	case Regular:
		return convertRegular(t, args[0].Regular())
	default:
		return nil, errors.New("unable to convert " + args[0].String())
	}
}

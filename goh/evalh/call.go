package evalh

import (
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/strconvh"
	"reflect"
)

func callArgsCountLessError() *intError {
	return newIntError("not enough arguments in call")
}
func callArgsCountMoreError() *intError {
	return newIntError("too many arguments in call")
}
func callArgsCountMismError(req, got int) *intError {
	if req > got {
		return callArgsCountLessError()
	}
	return callArgsCountMoreError()
}
func callNonFuncError(f Value) *intError {
	return newIntError("cannot call non-function (type " + f.DeepType() + ")")
}
func multValueError() *intError {
	return newIntError("multiple-value in single-value context")
}
func callWithNoResultError() *intError {
	return newIntError("function call with no result used as value")
}
func callResultCountMismError(got int) *intError {
	if got > 1 {
		return multValueError()
	}
	return callWithNoResultError()
}
func callInvArgAtError(pos int, x Value, reqT reflect.Type) *intError {
	return newIntError("cannot use " + x.String() + " (type " + x.DeepType() + ") as type " + reqT.String() + " in argument #" + strconvh.FormatInt(pos))
}
func callPanicError(p interface{}) *intError {
	return newIntErrorf("runtime panic in function call (%v)", p)
}
func convertMultError(t reflect.Type, x []Value) *intError {
	msg := "too many arguments to conversion to " + t.String() + ": "
	for i := range x {
		if i != 0 {
			msg += ", "
		}
		msg += x[i].String()
	}
	return newIntError(msg)
}
func convertUnableError(t reflect.Type, x Value) *intError {
	return newIntError("cannot convertCall " + x.String() + " (type " + x.DeepType() + ") to type " + t.String())
}

func callRegular(f reflect.Value, args []Value) (r Value, err *intError) {
	if f.Kind() != reflect.Func {
		return nil, callNonFuncError(MakeRegular(f))
	}

	// Check in/out arguments count
	fT := f.Type()
	if fT.NumIn() != len(args) {
		return nil, callArgsCountMismError(fT.NumIn(), len(args))
	}
	if fT.NumOut() != 1 {
		return nil, callResultCountMismError(fT.NumOut())
	}

	// Prepare arguments
	typedArgs := make([]reflect.Value, fT.NumIn())
	for i := range args {
		switch args[i].Kind() {
		case Untyped:
			var ok bool
			typedArgs[i], ok = constanth.SameType(args[i].Untyped(), fT.In(i))
			if !ok {
				return nil, callInvArgAtError(i, args[i], fT.In(i))
			}
		case Nil:
			typedArgs[i] = reflect.ValueOf(nil) // TODO check this, may be "convertNil()"
		case Regular:
			if aT := args[i].Regular().Type(); !aT.AssignableTo(fT.In(i)) {
				return nil, callInvArgAtError(i, args[i], fT.In(i))
			}
			typedArgs[i] = args[i].Regular()
		default:
			return nil, callInvArgAtError(i, args[i], fT.In(i))
		}
	}

	defer func() {
		if rec := recover(); rec != nil {
			r = nil
			err = callPanicError(rec)
		}
	}()
	rs := f.Call(typedArgs)
	return MakeRegular(rs[0]), nil
}

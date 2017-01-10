package evalh

import (
	"reflect"
)

// ellipsis true if last argument has ellipsis notation ("f(a,b,c...)").
func callRegular(f reflect.Value, args []Value, ellipsis bool) (r Value, err *intError) {
	if f.Kind() != reflect.Func {
		return nil, callNonFuncError(MakeRegular(f))
	}
	fT := f.Type()
	if fT.NumOut() != 1 {
		return nil, callResultCountMismError(fT.NumOut())
	}

	switch fT.IsVariadic() {
	case true:
		switch ellipsis {
		case true:
			return callRegularVariadicEllipsis(f, args)
		case false:
			return callRegularVariadic(f, args)
		default:
			return // unreachable
		}
	case false:
		if ellipsis {
			return nil, callRegularWithEllipsisError()
		}
		return callRegularNonVariadic(f, args)
	default:
		return // unreachable
	}
}

// f must be variadic func with single result (check must perform caller).
func callRegularVariadicEllipsis(f reflect.Value, args []Value) (r Value, err *intError) {
	fT := f.Type()
	if len(args) != fT.NumIn() {
		return nil, callArgsCountMismError(fT.NumIn(), len(args))
	}

	// Prepare arguments
	typedArgs := make([]reflect.Value, len(args))
	for i := range args {
		var ok bool
		typedArgs[i], _, ok = args[i].AsType(fT.In(i))
		if !ok {
			return nil, callInvArgAtError(i, args[i], fT.In(i))
		}
	}

	defer func() {
		if rec := recover(); rec != nil {
			r = nil
			err = callPanicError(rec)
		}
	}()
	rs := f.CallSlice(typedArgs)
	return MakeRegular(rs[0]), nil
}

// f must be variadic func with single result (check must perform caller).
func callRegularVariadic(f reflect.Value, args []Value) (r Value, err *intError) {
	fT := f.Type()
	if len(args) < fT.NumIn()-1 {
		return nil, callArgsCountMismError(fT.NumIn(), len(args)-1)
	}

	// Prepare arguments
	typedArgs := make([]reflect.Value, len(args))
	// non-variadic arguments
	for i := 0; i < fT.NumIn()-1; i++ {
		var ok bool
		typedArgs[i], _, ok = args[i].AsType(fT.In(i))
		if !ok {
			return nil, callInvArgAtError(i, args[i], fT.In(i))
		}
	}
	// variadic arguments
	variadicT := fT.In(fT.NumIn() - 1).Elem()
	for i := fT.NumIn() - 1; i < len(args); i++ {
		var ok bool
		typedArgs[i], _, ok = args[i].AsType(variadicT)
		if !ok {
			return nil, callInvArgAtError(i, args[i], variadicT)
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

// f must be non-variadic func with single result (check must perform caller).
func callRegularNonVariadic(f reflect.Value, args []Value) (r Value, err *intError) {
	// Check in/out arguments count
	fT := f.Type()
	if len(args) != fT.NumIn() {
		return nil, callArgsCountMismError(fT.NumIn(), len(args))
	}

	// Prepare arguments
	typedArgs := make([]reflect.Value, len(args))
	for i := range args {
		//switch args[i].Kind() {
		//case Untyped:
		//	var ok bool
		//	typedArgs[i], ok = constanth.AsType(args[i].Untyped(), fT.In(i))
		//	if !ok {
		//		return nil, callInvArgAtError(i, args[i], fT.In(i))
		//	}
		//case Nil:
		//	typedArgs[i] = reflect.ValueOf(nil)  TO DO check this, may be "convertNil()"
		//case Regular:
		//	if aT := args[i].Regular().Type(); !aT.AssignableTo(fT.In(i)) {
		//		return nil, callInvArgAtError(i, args[i], fT.In(i))
		//	}
		//	typedArgs[i] = args[i].Regular()
		//default:
		//	return nil, callInvArgAtError(i, args[i], fT.In(i))
		//}
		var ok bool
		typedArgs[i], _, ok = args[i].AsType(fT.In(i))
		if !ok {
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

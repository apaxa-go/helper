package evalh

import (
	"github.com/apaxa-go/helper/strconvh"
	"go/ast"
	"go/token"
	"reflect"
)

func identUndefinedError(ident string) *intError {
	return newIntError("undefined: " + ident)
}

func invAstError(msg string) *intError {
	return newIntError("invalid AST: " + msg)
}

func invAstNilError() *intError {
	return invAstError("evaluate nil expr")
}

func invAstUnsupportedError(e ast.Expr) *intError {
	return invAstError("expression evaluation does not support " + reflect.TypeOf(e).String())
}

func invAstSelectorError() *intError {
	return invAstError("no field specified (Sel is nil)")
}

func invSelectorXError(x Value) *intError {
	return newIntError("unable to select from " + x.DeepType())
}

func syntaxError(msg string) *intError {
	return newIntError("syntax error: " + msg)
}

func syntaxInvBasLitError(literal string) *intError {
	return syntaxError("invalid basic literal \"" + literal + "\"")
}

func indirectInvalError(x Value) *intError {
	return newIntError("invalid indirect of " + x.String() + " (type " + x.DeepType() + ")")
}

func notExprError(x Value) *intError {
	return newIntError(x.DeepType() + " is not an expression")
}

func syntaxMisChanTypeError() *intError {
	return syntaxError("syntax error: missing channel element type")
}

func syntaxMisArrayTypeError() *intError {
	return syntaxError("syntax error: missing array element type")
}

func syntaxMisVariadicTypeError() *intError {
	return syntaxError("final argument in variadic function missing type")
}

func sliceInvTypeError(x Value) *intError {
	return newIntError("cannot slice " + x.String() + " (type " + x.DeepType() + ")")
}

func notTypeError(x Value) *intError {
	return newIntError(x.String() + " is not a type")
}

func initMixError() *intError {
	return newIntError("mixture of field:value and value initializers")
}

func initStructInvFieldNameError() *intError {
	return newIntError("invalid field name in struct initializer")
}

func initArrayInvIndexError() *intError {
	return newIntError("index must be non-negative integer constant")
}

func initArrayDupIndexError(i int) *intError {
	return newIntError("duplicate index in array literal: " + strconvh.FormatInt(i))
}

func initMapMisKeyError() *intError {
	return newIntError("missing key in map literal")
}

func initInvTypeError(t reflect.Type) *intError {
	return newIntError("invalid type for composite literal: " + t.String())
}

func funcInvEllipsisPos() *intError {
	return newIntError("can only use ... with final parameter in list")
}

func cannotUseAsError(dst reflect.Type, src Value, in string) *intError {
	return newIntError("cannot use " + src.String() + " (type " + src.DeepType() + ") as type " + dst.String() + " in " + in)
}
func assignTypesMismError(dst reflect.Type, src Value) *intError {
	return cannotUseAsError(dst, src, "assigment")
}
func appendMismTypeError(dst reflect.Type, src Value) *intError {
	return cannotUseAsError(dst, src, "append")
}
func assignDstUnsettableError(dst Value) *intError {
	return newIntError("cannot change " + dst.String() + " (type " + dst.DeepType() + ") in assignment")
}
func compLitInvTypeError(t reflect.Type) *intError {
	return newIntError("invalid type for composite literal: " + t.String())
}
func compLitUnknFieldError(s reflect.Value, f string) *intError {
	return newIntError("unknown " + s.Type().String() + " field '" + f + "' in struct literal")
}
func compLitArgsCountMismError(req, got int) *intError { // TODO in other *ArgsCountMismError move errors into signle function?
	if req > got {
		return newIntError("too few values in struct initializer")
	}
	return newIntError("too many values in struct initializer")
}
func compLitNegIndexError() *intError {
	return newIntError("index must be non-negative integer constant")
}
func compLitIndexOutOfBoundsError(max, i int) *intError {
	return newIntError("array index " + strconvh.FormatInt(i) + " out of bounds [0:" + strconvh.FormatInt(max) + "]")
}

const invBinOp = "invalid operation: %v %v %v (%v)" // Widely used error
func invBinOpTypesMismError(x Value, op token.Token, y Value) *intError {
	return newIntErrorf(invBinOp, x.String(), op.String(), y.String(), "mismatched types "+x.DeepType()+" and "+y.DeepType())
}

func invBinOpTypesInvalError(x Value, op token.Token, y Value) *intError {
	return newIntErrorf(invBinOp, x.String(), op.String(), y.String(), "invalid types "+x.DeepType()+" and/or "+y.DeepType())
}

func invBinOpTypesIncompError(x Value, op token.Token, y Value) *intError {
	return newIntErrorf(invBinOp, x.String(), op.String(), y.String(), "incomparable types "+x.DeepType()+" and "+y.DeepType())
}

func invBinOpTypesUnorderError(x Value, op token.Token, y Value) *intError {
	return newIntErrorf(invBinOp, x.String(), op.String(), y.String(), "unordered types "+x.DeepType()+" and "+y.DeepType())
}

func invBinOpInvalError(x Value, op token.Token, y Value) *intError {
	return newIntErrorf(invBinOp, x.String(), op.String(), y.String(), "invalid operator")
}

func invBinOpShiftCountError(x Value, op token.Token, y Value) *intError {
	return newIntErrorf(invBinOp, x.String(), op.String(), y.String(), "shift count type "+y.DeepType()+", must be unsigned integer")
}

func invBinOpShiftArgError(x Value, op token.Token, y Value) *intError {
	return newIntErrorf(invBinOp, x.String(), op.String(), y.String(), "shift of type "+y.DeepType())
}

func callBuiltInArgsCountLessError(fn string) *intError {
	return newIntError("not enough arguments in call to " + fn)
}
func callBuiltInArgsCountMoreError(fn string) *intError {
	return newIntError("too many arguments in call to " + fn)
}
func callBuiltInArgsCountMismError(fn string, req, got int) *intError {
	if req > got {
		return callBuiltInArgsCountLessError(fn)
	}
	return callBuiltInArgsCountMoreError(fn)
}
func invBuiltInArgError(fn string, x Value) *intError {
	return newIntError("invalid argument " + x.String() + " (type " + x.DeepType() + ") for " + fn)
}
func invBuiltInArgAtError(fn string, pos int, x Value) *intError {
	return newIntError("invalid argument #" + strconvh.FormatInt(pos) + " " + x.String() + " (type " + x.DeepType() + ") for " + fn)
}

func invBuiltInArgsError(fn string, x []Value) *intError {
	var msg string
	for i := range x {
		if i != 0 {
			msg += ", "
		}
		msg += x[i].String() + " (type " + x[i].DeepType() + ")"
	}
	return newIntError("invalid arguments " + msg + " for " + fn)
}

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
func convertArgsCountMismError(t reflect.Type, req int, x []Value) *intError {
	var msg string
	switch {
	case len(x) > req:
		msg = "too many arguments to conversion to " + t.String() + ": "
		for i := range x {
			if i != 0 {
				msg += ", "
			}
			msg += x[i].String()
		}
	default:
		msg = "no arguments to conversion to " + t.String()
	}
	return newIntError(msg)
}
func convertUnableError(t reflect.Type, x Value) *intError {
	return newIntError("cannot convertCall " + x.String() + " (type " + x.DeepType() + ") to type " + t.String())
}

func convertNilUnableError(t reflect.Type) *intError {
	return newIntError("cannot convertCall nil to type " + t.String())
}

func undefIdentError(ident string) *intError {
	return newIntError("undefined: " + ident)
}

func invSliceOpError(x Value) *intError {
	return newIntError("cannot slice " + x.String() + " (type " + x.DeepType() + ")")
}
func invSliceIndexError(low, high int) *intError {
	return newIntError("invalid slice index: " + strconvh.FormatInt(low) + " > " + strconvh.FormatInt(high))
}
func invSlice3IndexOmitted() *intError {
	return newIntError("only first index in 3-index slice can be omitted")
}

func invUnaryOp(x Value, op token.Token) *intError {
	return newIntError("invalid operation: " + op.String() + " " + x.DeepType())
}
func invUnaryOpReason(x Value, op token.Token, reason interface{}) *intError {
	return newIntErrorf("invalid operation: %v %v: %v", op.String(), x.DeepType(), reason)
}
func invUnaryReceiveError(x Value, op token.Token) *intError {
	return invUnaryOpReason(x, op, "receive from non-chan type "+x.Type().String())
}
func selectorUndefIdentError(t reflect.Type, name string) *intError {
	return undefIdentError(t.String() + "." + name)
}

func arrayBoundNegError() *intError {
	return newIntError("array bound must be non-negative")
}

func convertWithEllipsisError(t reflect.Type) *intError {
	return newIntError("invalid use of ... in type conversion to " + t.String())
}

func callBuiltInWithEllipsisError(f string) *intError {
	return newIntError("invalid use of ... with builtin " + f)
}

func callRegularWithEllipsisError() *intError {
	return newIntError("invalid use of ... in call")
}
func makeInvalidTypeError(t reflect.Type) *intError {
	return newIntError("cannot make type " + t.String())
}

func makeNotIntArgError(t reflect.Type, argName string, arg Value) *intError {
	return newIntError("non-integer " + argName + " argument in make(" + t.String() + ") - " + arg.DeepType())
}
func makeNegArgError(t reflect.Type, argName string) *intError {
	return newIntError("negative " + argName + " argument in make(" + t.String() + ")")
}
func makeSliceMismArgsError(t reflect.Type) *intError {
	return newIntError("len larger than cap in make(" + t.String() + ")")
}
func appendFirstNotSliceError(x Value) *intError {
	return newIntError("first argument to append must be slice; have " + x.DeepType())
}

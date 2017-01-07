package evalh

import (
	"fmt"
	"github.com/apaxa-go/helper/goh/constanth"
	"go/constant"
	"go/token"
	"reflect"
	"unicode"
)

type Kind int

const (
	Nil         Kind = iota
	Regular          = iota
	Untyped          = iota
	Type             = iota
	BuiltInFunc      = iota
	Package          = iota
)

type Value interface {
	Kind() Kind
	DeepType() string
	String() string
	Regular() reflect.Value
	Untyped() constant.Value
	Type() reflect.Type
	BuiltInFunc() string
	Package() map[string]Value
	Equal(Value) bool // TODO move to test?
	Interface() interface{}
	Int() (r int, isConst, ok bool)                            // No any conversion. Return int for regular variable with exactly int type and for untyped constant which can be represent as int.
	AsType(t reflect.Type) (r reflect.Value, isConst, ok bool) // No any conversion. Return r of type t if it is possible. Only regular, untyped and nil can return non false result.
	ToType(t reflect.Type) (r reflect.Value, isConst, ok bool) // Convert Value to type t if it is possible. Only regular, untyped and nil can return non false result.
	implementsValue()
}

type (
	nilVal         struct{}
	regVal         reflect.Value
	untypedVal     struct{ v constant.Value }
	typeVal        struct{ v reflect.Type }
	builtInFuncVal string
	packageVal     map[string]Value
)

func (nilVal) Kind() Kind         { return Nil }
func (regVal) Kind() Kind         { return Regular }
func (untypedVal) Kind() Kind     { return Untyped }
func (typeVal) Kind() Kind        { return Type }
func (builtInFuncVal) Kind() Kind { return BuiltInFunc }
func (packageVal) Kind() Kind     { return Package }

func (nilVal) DeepKind() string   { return "nil" }
func (x regVal) DeepKind() string { return x.Regular().Type().String() }
func (x untypedVal) DeepKind() string {
	return constanth.KindString(x.Untyped().Kind()) + " (" + x.Untyped().ExactString() + ")"
}
func (x typeVal) DeepKind() string      { return "type" }
func (builtInFuncVal) DeepKind() string { return "built-in function" }
func (packageVal) DeepKind() string     { return "package" }

func (nilVal) String() string { return "nil" }
func (x regVal) String() string {
	return fmt.Sprintf("regular value %v (%v)", reflect.Value(x), reflect.Value(x).Type().String())
}
func (x untypedVal) String() string {
	return fmt.Sprintf("untyped value %v (%v)", x.v.ExactString(), constanth.KindString(x.v.Kind()))
}
func (x typeVal) String() string        { return fmt.Sprintf("type value %v", x.v.String()) }
func (x builtInFuncVal) String() string { return fmt.Sprintf("built-in function value %v", string(x)) }
func (x packageVal) String() string     { return fmt.Sprintf("package (%v)", x) }

func (nilVal) Regular() reflect.Value         { panic("") }
func (x regVal) Regular() reflect.Value       { return reflect.Value(x) }
func (untypedVal) Regular() reflect.Value     { panic("") }
func (typeVal) Regular() reflect.Value        { panic("") }
func (builtInFuncVal) Regular() reflect.Value { panic("") }
func (packageVal) Regular() reflect.Value     { panic("") }

func (nilVal) Untyped() constant.Value         { panic("") }
func (regVal) Untyped() constant.Value         { panic("") }
func (x untypedVal) Untyped() constant.Value   { return x.v }
func (typeVal) Untyped() constant.Value        { panic("") }
func (builtInFuncVal) Untyped() constant.Value { panic("") }
func (packageVal) Untyped() constant.Value     { panic("") }

func (nilVal) Type() reflect.Type         { panic("") }
func (regVal) Type() reflect.Type         { panic("") }
func (untypedVal) Type() reflect.Type     { panic("") }
func (x typeVal) Type() reflect.Type      { return x.v }
func (builtInFuncVal) Type() reflect.Type { panic("") }
func (packageVal) Type() reflect.Type     { panic("") }

func (nilVal) BuiltInFunc() string           { panic("") }
func (regVal) BuiltInFunc() string           { panic("") }
func (untypedVal) BuiltInFunc() string       { panic("") }
func (typeVal) BuiltInFunc() string          { panic("") }
func (x builtInFuncVal) BuiltInFunc() string { return string(x) }
func (packageVal) BuiltInFunc() string       { panic("") }

func (nilVal) Package() map[string]Value         { panic("") }
func (regVal) Package() map[string]Value         { panic("") }
func (untypedVal) Package() map[string]Value     { panic("") }
func (typeVal) Package() map[string]Value        { panic("") }
func (builtInFuncVal) Package() map[string]Value { panic("") }
func (x packageVal) Package() map[string]Value   { return map[string]Value(x) }

func (nilVal) Equal(v Value) bool { return v.Kind() == Nil }
func (x regVal) Equal(v Value) (r bool) {
	if v.Kind() != Regular {
		return false
	}

	xV := x.Regular()
	vV := v.Regular()

	if xV.Kind() != vV.Kind() {
		return false
	}

	// Compare functions
	if xV.Kind() == reflect.Func {
		return xV.Pointer() == vV.Pointer() // may return wrong result: http://stackoverflow.com/questions/9643205/how-do-i-compare-two-functions-for-pointer-equality-in-the-latest-go-weekly
	}
	// Compare slices
	if xV.Kind() == reflect.Slice {
		return reflect.DeepEqual(xV.Interface(), vV.Interface()) // not a good check
	}

	defer func() {
		if rec := recover(); rec != nil {
			r = false
		}
	}()
	r = xV.Interface() == vV.Interface()
	return
}
func (x untypedVal) Equal(v Value) bool {
	if v.Kind() != Untyped {
		return false
	}
	return constant.Compare(x.v, token.EQL, v.Untyped())
}
func (x typeVal) Equal(v Value) bool {
	if v.Kind() != Type {
		return false
	}
	return x.v == v.Type()
}
func (x builtInFuncVal) Equal(v Value) bool {
	if v.Kind() != BuiltInFunc {
		return false
	}
	return x.BuiltInFunc() == v.BuiltInFunc()
}

func (x packageVal) Equal(v Value) bool {
	if v.Kind() != Package {
		return false
	}
	return reflect.DeepEqual(x.Package(), v.Package())
}

func (nilVal) Interface() interface{}         { return nil }
func (x regVal) Interface() interface{}       { return reflect.Value(x).Interface() }
func (x untypedVal) Interface() interface{}   { r, _ := constanth.DefaultTypeInterface(x.v); return r }
func (typeVal) Interface() interface{}        { panic("") }
func (builtInFuncVal) Interface() interface{} { panic("") }
func (packageVal) Interface() interface{}     { panic("") }

func (nilVal) Int() (r int, isConst, ok bool) { return 0, true, false }
func (x regVal) Int() (r int, isConst, ok bool) {
	if !x.Regular().CanInterface() {
		return 0, false, false
	}
	r, ok = x.Regular().Interface().(int)
	return
}
func (x untypedVal) Int() (r int, isConst, ok bool) {
	isConst = true
	r, ok = constanth.IntVal(x.Untyped())
	return
}
func (typeVal) Int() (r int, isConst, ok bool)        { return 0, false, false }
func (builtInFuncVal) Int() (r int, isConst, ok bool) { return 0, false, false }
func (packageVal) Int() (r int, isConst, ok bool)     { return 0, false, false }

func (nilVal) AsType(t reflect.Type) (r reflect.Value, isConst, ok bool) {
	isConst = true
	switch t.Kind() {
	case reflect.Slice, reflect.Ptr, reflect.Func, reflect.Interface, reflect.Map, reflect.Chan:
		r = reflect.New(t) // TODO check if result is adequate
		ok = true
	}
	return
}
func (x regVal) AsType(t reflect.Type) (r reflect.Value, isConst, ok bool) {
	if x.Regular().Type() == t {
		r = x.Regular()
		ok = true
	}
	return
}
func (x untypedVal) AsType(t reflect.Type) (r reflect.Value, isConts, ok bool) {
	isConts = true
	if r, ok = constanth.SameType(x.Untyped(), t); ok {
		return
	}

	if x.Untyped().Kind() == constant.Int && t.Kind() == reflect.String { // TODO may be move this case to constanth.SameType()
		var i rune
		i, ok = constanth.RuneVal(x.Untyped())
		if !ok {
			i = unicode.ReplacementChar
			ok = true
		}
		r = reflect.New(t).Elem()
		r.SetString(string(i))
	}
	return
}
func (typeVal) AsType(t reflect.Type) (r reflect.Value, isConst, ok bool) { return reflect.Value{}, false, false }
func (builtInFuncVal) AsType(t reflect.Type) (r reflect.Value, isConst, ok bool) {return reflect.Value{}, false, false}
func (packageVal) AsType(t reflect.Type) (r reflect.Value, isConst, ok bool) { return reflect.Value{}, false, false }

func (x nilVal) ToType(t reflect.Type) (r reflect.Value, isConst, ok bool) { return x.AsType(t) }
func (x regVal) ToType(t reflect.Type) (r reflect.Value, isConst, ok bool) {
	xT := x.Regular().Type()
	if xT.ConvertibleTo(t) {
		r=x.Regular().Convert(t)
		ok=true
	}
	return
}
func (x untypedVal)ToType(t reflect.Type)(r reflect.Value, isConst, ok bool) { return x.AsType(t) }
func (typeVal) ToType(t reflect.Type) (r reflect.Value, isConst, ok bool) { return reflect.Value{}, false, false }
func (builtInFuncVal) ToType(t reflect.Type) (r reflect.Value, isConst, ok bool) {return reflect.Value{}, false, false}
func (packageVal) ToType(t reflect.Type) (r reflect.Value, isConst, ok bool) { return reflect.Value{}, false, false }

func (nilVal) implementsValue()         {}
func (regVal) implementsValue()         {}
func (untypedVal) implementsValue()     {}
func (typeVal) implementsValue()        {}
func (builtInFuncVal) implementsValue() {}
func (packageVal) implementsValue()     {}

func MakeNil() Value                           { return nilVal{} }
func MakeRegular(x reflect.Value) Value        { return regVal(x) }
func MakeRegularInterface(x interface{}) Value { return regVal(reflect.ValueOf(x)) }
func MakeUntyped(x constant.Value) Value       { return untypedVal{x} }
func MakeUntypedBool(x bool) Value             { return MakeUntyped(constant.MakeBool(x)) }
func MakeUntypedFloat64(x float64) Value       { return MakeUntyped(constant.MakeFloat64(x)) }
func MakeUntypedInt64(x int64) Value           { return MakeUntyped(constant.MakeInt64(x)) }
func MakeUntypedString(x string) Value         { return MakeUntyped(constant.MakeString(x)) }
func MakeUntypedUint64(x uint64) Value         { return MakeUntyped(constant.MakeUint64(x)) }
func MakeUntypedComplex128(x complex128) Value { return MakeUntyped(constanth.MakeComplex128(x)) }
func MakeType(x reflect.Type) Value            { return typeVal{x} }
func MakeBuiltInFunc(x string) Value           { return builtInFuncVal(x) }
func MakePackage(idents Identifiers) Value     { return packageVal(idents) } // keys in idents must not have dots in names

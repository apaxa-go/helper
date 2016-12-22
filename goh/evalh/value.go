package evalh

import (
	"fmt"
	"github.com/apaxa-go/helper/goh/constanth"
	"go/constant"
	"go/token"
	"reflect"
)

// TODO implement panic-safe constant API and use only it

type Kind int

const (
	Nil     Kind = iota
	Regular      = iota
	Untyped      = iota
	BuiltIn      = iota
)

type Value interface {
	Kind() Kind
	String() string
	Regular() reflect.Value
	Untyped() constant.Value
	Equal(Value) bool // TODO move to test?
	Interface() interface{}
	implementsValue()
}

type (
	nilVal     struct{}
	regVal     reflect.Value
	untypedVal struct{ v constant.Value }
)

func (nilVal) Kind() Kind     { return Nil }
func (regVal) Kind() Kind     { return Regular }
func (untypedVal) Kind() Kind { return Untyped }

func (nilVal) String() string { return "nil" }
func (x regVal) String() string {
	return fmt.Sprintf("regular value %v (%v)", reflect.Value(x), reflect.Value(x).Type().String())
}
func (x untypedVal) String() string {
	return fmt.Sprintf("untyped value %v (%v)", x.v.ExactString(), constanth.KindString(x.v.Kind()))
}

func (nilVal) Regular() reflect.Value     { panic("") }
func (x regVal) Regular() reflect.Value   { return reflect.Value(x) }
func (untypedVal) Regular() reflect.Value { panic("") }

func (nilVal) Untyped() constant.Value       { panic("") }
func (regVal) Untyped() constant.Value       { panic("") }
func (x untypedVal) Untyped() constant.Value { return x.v }

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

func (nilVal) Interface() interface{}       { return nil }
func (x regVal) Interface() interface{}     { return reflect.Value(x).Interface() }
func (x untypedVal) Interface() interface{} { r, _ := constanth.DefaultTypeInterface(x.v); return r }

func (nilVal) implementsValue()     {}
func (regVal) implementsValue()     {}
func (untypedVal) implementsValue() {}

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

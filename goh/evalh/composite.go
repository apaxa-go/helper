package evalh

import (
	"errors"
	"fmt"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/strconvh"
	"reflect"
)

func assign(to reflect.Value, value Value) (err error) {
	if !to.CanSet() {
		return errors.New("unable to set")
	}

	var newValue reflect.Value
	switch value.Kind() {
	case Untyped:
		var ok bool
		newValue, ok = constanth.SameType(value.Untyped(), to.Type())
		if !ok {
			return errors.New("unable to set " + to.String() + " to value " + value.String())
		}
	case Regular:
		newValue = value.Regular()
	// TODO case Nil:
	default:
		return errors.New("unable to set " + to.String() + " to value " + value.String())
	}

	if !newValue.Type().AssignableTo(to.Type()) {
		return errors.New("unable to set " + to.String() + " to value " + value.String())
	}
	to.Set(newValue)
	return nil
}

func compositeLitStructKeys(t reflect.Type, elts map[string]Value) (r Value, err error) {
	if t.Kind() != reflect.Struct {
		return nil, errors.New("struct type required") // Unreachable?
	}
	rV := reflect.New(t).Elem()

	for field, value := range elts {
		fV := rV.FieldByName(field)
		if !fV.IsValid() {
			return nil, errors.New("no such field: " + field)
		}
		err = assign(fV, value)
		if err != nil {
			return
		}
	}

	return MakeRegular(rV), nil
}

func compositeLitStructOrdered(t reflect.Type, elts []Value) (r Value, err error) {
	if t.Kind() != reflect.Struct {
		return nil, errors.New("struct type required") // Unreachable?
	}
	if t.NumField() != len(elts) {
		return nil, errors.New(fmt.Sprintf("struct literal requires %v elements, but got %v", t.NumField(), len(elts)))
	}
	rV := reflect.New(t).Elem()

	for i := range elts {
		fV := rV.Field(i)
		if !fV.IsValid() {
			return nil, errors.New("no field with index " + strconvh.FormatInt(i))
		}
		err = assign(fV, elts[i])
		if err != nil {
			return
		}
	}

	return MakeRegular(rV), nil
}

func compositeLitArrayLike(t reflect.Type, elts map[int]Value) (r Value, err error) {
	// Calc max index (len of slice)
	var maxIndex = -1
	for i := range elts {
		if i < 0 {
			return nil, errors.New("index cannot be negative")
		}
		if i > maxIndex {
			maxIndex = i
		}
	}

	// Init result
	rV := reflect.New(t).Elem()
	switch t.Kind() {
	case reflect.Array:
		if maxIndex > rV.Len()-1 {
			return nil, errors.New("index out of range")
		}
	case reflect.Slice:
		rV.SetLen(maxIndex + 1)
	default:
		return nil, errors.New("slice or array required") // Unreachable?
	}

	// Fill result
	for i := range elts {
		iV := rV.Index(i)
		if !iV.IsValid() {
			return nil, errors.New("unreachable?") // TODO
		}
		err = assign(iV, elts[i])
		if err != nil {
			return
		}
	}

	return MakeRegular(rV), nil
}

func compositeLitMap(t reflect.Type, elts map[Value]Value) (r Value, err error) {
	if t.Kind() != reflect.Map {
		return nil, errors.New("map required")
	}

	rV := reflect.New(t).Elem()
	kV := reflect.New(t.Key()).Elem()
	vV := reflect.New(t.Elem()).Elem()
	for k, v := range elts {
		err = assign(kV, k)
		if err != nil {
			return
		}
		err = assign(vV, v)
		if err != nil {
			return
		}
		rV.SetMapIndex(kV, vV)
	}

	return MakeRegular(rV), nil
}

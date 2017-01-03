package evalh

import (
	"errors"
	"fmt"
	"github.com/apaxa-go/helper/goh/constanth"
	"go/constant"
	"reflect"
	"unicode"
)

func convertUntyped(t reflect.Type, x constant.Value) (r Value, err error) {
	if v, ok := constanth.SameType(x, t); ok {
		return MakeRegular(v), nil
	}

	if x.Kind() == constant.Int && t.Kind() == reflect.String {
		i, ok := constanth.RuneVal(x)
		if !ok {
			i = unicode.ReplacementChar
		}
		v := reflect.New(t).Elem()
		v.SetString(string(i))
		return MakeRegular(v), nil
	}

	return nil, errors.New(fmt.Sprintf("unable to convert %v to %v", x.String(), t.String()))
}

func convertRegular(t reflect.Type, x reflect.Value) (r Value, err error) {
	xT := x.Type()
	if !xT.ConvertibleTo(t) {
		return nil, errors.New(fmt.Sprintf("unable to convert %v to %v", x.String(), t.String()))
	}
	rV := x.Convert(t) //TODO Is it required to catch panic?
	return MakeRegular(rV), nil
}

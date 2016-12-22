package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/constanth"
	"go/ast"
	"reflect"
)

const indexSkipped int = -1

// getSliceIndex returns index value (int).
// Returned index may be negative (-1) only if e is nil (this means that index is skipped in source).
// In all other cases negative index causes error.
func getSliceIndex(e ast.Expr, idents Identifiers) (r int, err error) {
	if e == nil {
		return indexSkipped, nil
	}

	v, err := Expr(e, idents)
	if err != nil {
		return 0, err
	}

	switch v.Kind() {
	case Regular:
		// Check kind
		if v.Regular().Kind() != reflect.Int {
			return 0, errors.New("unable to slicing using " + v.String())
		}
		// Check exact type
		var ok bool
		r, ok = v.Regular().Interface().(int)
		if !ok {
			return 0, errors.New("unable to slicing using " + v.String())
		}
	case Untyped:
		var ok bool
		r, ok = constanth.IntVal(v.Untyped())
		if !ok {
			return 0, errors.New("unable to slicing using " + v.String())
		}
	case Nil:
		return 0, errors.New("unable to slicing using Nil")
	default:
		panic("unknown kind")
	}

	if r < 0 {
		return 0, errors.New("negative index value")
	}
	return
}

func slice2(x reflect.Value, low, high int) (r Value, err error) {
	// resolve pointer to array
	if x.Kind() == reflect.Ptr && x.Elem().Kind() == reflect.Array {
		x = x.Elem()
	}

	// check slicing possibility
	if k := x.Kind(); k != reflect.Array && k != reflect.Slice && k != reflect.String {
		return nil, errors.New("unable to slicing " + k.String())
	} else if k == reflect.Array && !x.CanAddr() {
		return nil, errors.New("unable to slicing unaddressable array")
	}

	// resolve default value
	if low == indexSkipped {
		low = 0
	}
	if high == indexSkipped {
		high = x.Len()
	}

	// validate indexes
	if 0 > low || low > high || high > x.Len() {
		return nil, errors.New("invalid index value")
	}

	return MakeRegular(x.Slice(low, high)), nil
}

func slice3(x reflect.Value, low, high, max int) (r Value, err error) {
	// resolve pointer to array
	if x.Kind() == reflect.Ptr && x.Elem().Kind() == reflect.Array {
		x = x.Elem()
	}

	// check slicing possibility
	if k := x.Kind(); k != reflect.Array && k != reflect.Slice {
		return nil, errors.New("unable to slicing " + k.String())
	} else if k == reflect.Array && !x.CanAddr() {
		return nil, errors.New("unable to slicing unaddressable array")
	}

	// resolve default value
	if low == indexSkipped {
		low = 0
	}

	// validate indexes
	if high == indexSkipped || max == indexSkipped {
		return nil, errors.New("only first index in 3-index slicing can be omitted")
	}
	if 0 > low || low > high || high > max || max > x.Cap() {
		return nil, errors.New("invalid index value")
	}

	return MakeRegular(x.Slice3(low, high, max)), nil
}

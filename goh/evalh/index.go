package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/strconvh"
	"reflect"
)

func indexMap(x reflect.Value, i Value) (r Value, err error) {
	if k := x.Kind(); k != reflect.Map {
		return nil, errors.New("unable to index (map) " + k.String())
	}

	iReqT := x.Type().Key()

	// calc index
	var iV reflect.Value
	switch i.Kind() {
	case Untyped:
		var ok bool
		iV, ok = constanth.SameType(i.Untyped(), iReqT)
		if !ok {
			return nil, errors.New("unable convert index (untyped constant " + i.String() + ") to " + iReqT.String())
		}
	case Regular:
		iV = i.Regular()
		if !iV.Type().AssignableTo(iReqT) {
			return nil, errors.New("unable to use " + iV.String() + " as " + iReqT.String() + " in map index")
		}
	case Nil:
		return nil, errors.New("unable to index Nil")
	default:
		panic("unknown kind")
	}

	//
	rV := x.MapIndex(iV)
	if !rV.IsValid() {
		return nil, errors.New("no such map index: " + iV.String())
	}

	return MakeRegular(rV), nil
}

func indexOther(x reflect.Value, i Value) (r Value, err error) {
	if k := x.Kind(); k != reflect.String && k != reflect.Array && k != reflect.Slice {
		return nil, errors.New("unable to index " + k.String())
	}

	// calc index
	var iInt int

	switch i.Kind() {
	case Untyped:
		var ok bool
		iInt, ok = constanth.IntVal(i.Untyped())
		if !ok {
			return nil, errors.New("unable to use untyped constant " + i.String() + " as index")
		}
	case Regular:
		if i.Regular().Kind() != reflect.Int {
			return nil, errors.New("unable to use " + i.String() + " as index")
		}
		iInt = int(i.Regular().Int()) // TODO check that type of i is int, not just underlying type is int
	case Nil:
		return nil, errors.New("unable to index Nil")
	default:
		panic("unknown kind")
	}

	// check out-of-range
	if iInt < 0 || iInt >= x.Len() {
		return nil, errors.New("index " + strconvh.FormatInt(iInt) + " out of range")
	}

	return MakeRegular(x.Index(iInt)), nil
}

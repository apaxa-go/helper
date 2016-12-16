package evalh

import (
	"errors"
	"github.com/apaxa-go/helper/goh/constanth"
	"github.com/apaxa-go/helper/strconvh"
	"go/constant"
	"reflect"
)

func indexMap(x interface{}, i interface{}) (r interface{}, err error) {
	xV := reflect.ValueOf(x)
	if k := xV.Kind(); k != reflect.Map {
		return nil, errors.New("unable to index (map) " + k.String())
	}

	iReqT := xV.Type().Key()

	// calc index
	var iV reflect.Value
	if iConst, ok := i.(constant.Value); ok {
		iV, ok = constanth.SameType(iConst, iReqT)
		if !ok {
			return nil, errors.New("unable convert index (untyped constant " + iConst.String() + ") to " + iReqT.String())
		}
	} else {
		iV = reflect.ValueOf(i)
		if !iV.Type().AssignableTo(iReqT) {
			return nil, errors.New("unable to use " + iV.String() + " as " + iReqT.String() + " in map index")
		}
	}

	rV := xV.MapIndex(iV)
	if !rV.IsValid() {
		return nil, errors.New("no such map index: " + iV.String())
	}

	return rV.Interface(), nil
}

func indexOther(x interface{}, i interface{}) (r interface{}, err error) {
	xV := reflect.ValueOf(x)
	if k := xV.Kind(); k != reflect.String && k != reflect.Array && k != reflect.Slice {
		return nil, errors.New("unable to index " + k.String())
	}

	// calc index
	//var iInt int
	iInt, ok := i.(int)
	if !ok {
		iConst, ok := i.(constant.Value)
		if !ok {
			return nil, errors.New("unable to use " + reflect.TypeOf(i).String() + " as index")
		}
		iInt, ok = constanth.IntVal(iConst)
		if !ok {
			return nil, errors.New("unable to use untyped constant " + iConst.String() + " as index")
		}
	}

	// check out-of-range
	if iInt < 0 || iInt >= xV.Len() {
		return nil, errors.New("index " + strconvh.FormatInt(iInt) + " out of range")
	}

	return xV.Index(iInt).Interface(), nil
}

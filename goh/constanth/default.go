package constanth

import "go/constant"

func DefaultType(x constant.Value) (r interface{}, ok bool) {
	switch k := x.Kind(); k {
	case constant.Bool:
		return constant.BoolVal(x), true
	case constant.String:
		return constant.StringVal(x), true
	case constant.Int:
		return IntVal(x)
	case constant.Float:
		return constant.Float64Val(x)
	case constant.Complex:
		return Complex128Val(x)
	default:
		return nil, false
	}
}

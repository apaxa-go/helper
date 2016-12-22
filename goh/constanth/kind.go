package constanth

import "go/constant"

func KindString(k constant.Kind) (r string) {
	switch k {
	case constant.Unknown:
		return "Unknown"
	case constant.Bool:
		return "Bool"
	case constant.String:
		return "String"
	case constant.Int:
		return "Int"
	case constant.Float:
		return "Float"
	case constant.Complex:
		return "Complex"
	default:
		return "Invalid"
	}
}

func IsNumeric(k constant.Value) bool {
	switch k.Kind() {
	case constant.Int, constant.Float, constant.Complex:
		return true
	default:
		return false
	}
}
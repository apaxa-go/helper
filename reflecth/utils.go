package reflecth

import (
	"github.com/apaxa-go/helper/goh/asth"
	"go/ast"
	"reflect"
)

func TypeOfPtr(i interface{}) reflect.Type {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Ptr {
		return nil
	}
	return t.Elem()
}

// for interfaces
func ValueOfPtr(i interface{}) reflect.Value{
	r:=reflect.ValueOf(i)
	if r.Type().Kind()!=reflect.Ptr{
		return reflect.Value{}
	}
	return r.Elem()
}

func ChanDirFromAst(dir ast.ChanDir) (r reflect.ChanDir) {
	switch dir {
	case asth.SendDir:
		return reflect.SendDir
	case asth.RecvDir:
		return reflect.RecvDir
	case asth.BothDir:
		return reflect.BothDir
	default:
		return 0
	}
}

func ChanDirToAst(dir reflect.ChanDir) (r ast.ChanDir) {
	switch dir {
	case reflect.SendDir:
		return asth.SendDir
	case reflect.RecvDir:
		return asth.RecvDir
	case reflect.BothDir:
		return asth.BothDir
	default:
		return 0
	}
}

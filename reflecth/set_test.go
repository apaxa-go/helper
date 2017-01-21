package reflecth

import (
	"testing"
	"reflect"
)

func TestSet(t *testing.T) {
	type myStruct0 struct {
		i int
		s string
	}

	my0 := myStruct0{}
	myV:=ValueOfPtr(&my0)

	myIV:=myV.FieldByName("i")
	Set(myIV,reflect.ValueOf(2))
	if testR:=(myStruct0{2, ""}); myV.Interface()!=testR{
		t.Errorf("expect %#v, got %#v", testR,myV.Interface())
	}

	//
	//
	//

	type myStruct1 struct {
		I int
		S string
	}

	my1 := myStruct1{}
	myV=ValueOfPtr(&my1)

	myIV=myV.FieldByName("I")
	Set(myIV,reflect.ValueOf(3))
	if testR:=(myStruct1{3, ""}); myV.Interface()!=testR{
		t.Errorf("expect %#v, got %#v", testR,myV.Interface())
	}
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type myStruct0 struct {
			i int
			s string
		}

		my0 := myStruct0{}
		myV:=ValueOfPtr(&my0)

		myIV:=myV.FieldByName("i")
		Set(myIV,reflect.ValueOf(2))
		if myV.Interface()!=(myStruct0{2, ""}){
			b.Errorf("%#v", myV.Interface())
		}
	}
}

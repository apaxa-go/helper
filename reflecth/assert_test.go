package reflecth

import (
	"errors"
	"testing"
)

func TestTypeAssert(t *testing.T) {
	b := true
	var i interface{} = b
	// interface with bool to bool
	if r, ok, valid := TypeAssert(ValueOfPtr(&i), TypeBool()); !valid || !ok || r.Interface() != b {
		t.Errorf("expect %v %v %v, got %v %v %v", b, true, true, r.Interface(), ok, valid)
	}
	// interface with bool to int
	if _, ok, valid := TypeAssert(ValueOfPtr(&i), TypeInt()); !valid || ok {
		t.Errorf("expect %v %v, got %v %v", false, true, ok, valid)
	}
	// interface woth bool to interface
	if r, ok, valid := TypeAssert(ValueOfPtr(&i), TypeOfPtr(&i)); !valid || !ok || r.Interface() != b {
		t.Errorf("expect %v %v %v, got %v %v %v", b, true, true, r.Interface(), ok, valid)
	}

	e := errors.New("str")
	i = e
	// error interface with error to interface
	if r, ok, valid := TypeAssert(ValueOfPtr(&e), TypeOfPtr(&i)); !valid || !ok || r.Interface() != e {
		t.Errorf("expect %v %v %v, got %v %v %v", e, true, true, r.Interface(), ok, valid)
	}
	// interface with error to error interface
	if r, ok, valid := TypeAssert(ValueOfPtr(&i), TypeOfPtr(&e)); !valid || !ok || r.Interface() != e {
		t.Errorf("expect %v %v %v, got %v %v %v", e, true, true, r.Interface(), ok, valid)
	}
	// error interface with error to bool
	if _, ok, valid := TypeAssert(ValueOfPtr(&e), TypeBool()); valid || ok {
		t.Errorf("expect %v %v, got %v %v", false, false, ok, valid)
	}
	i = b
	// interface with error to error interface
	if _, ok, valid := TypeAssert(ValueOfPtr(&i), TypeOfPtr(&e)); !valid || ok {
		t.Errorf("expect %v %v, got %v %v", false, true, ok, valid)
	}

	i = nil
	// interface with nil to bool
	if _, ok, valid := TypeAssert(ValueOfPtr(&i), TypeBool()); !valid || ok {
		t.Errorf("expect %v %v, got %v %v", false, true, ok, valid)
	}
}

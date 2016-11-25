package neth

import (
	"errors"
	"net"
	"testing"
)

type timeoutError struct {
	TimeoutValue bool
}

func (te timeoutError) Timeout() bool { return te.TimeoutValue }
func (te timeoutError) Error() string { return "TimeOut error" }

func TestIsTimeout(t *testing.T) {
	if IsTimeout(nil) {
		t.Error("TestIsTimeout: nil should cause false, not true")
	}
	if IsTimeout(errors.New("Some error")) {
		t.Error("TestIsTimeout: non timeout-aware error should cause false, not true")
	}
	var err net.OpError
	err.Err = timeoutError{true}
	if !IsTimeout(&err) {
		t.Error("TestIsTimeout: timeout error should cause true, not false")
	}
	err.Err = timeoutError{false}
	if IsTimeout(&err) {
		t.Error("TestIsTimeout: non timeout error should cause false, not true")
	}
}

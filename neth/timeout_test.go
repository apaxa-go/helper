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
		t.Error("nil should cause false, not true")
	}
	if IsTimeout(errors.New("some error")) {
		t.Error("non timeout-aware error should cause false, not true")
	}
	var err net.OpError
	err.Err = timeoutError{true}
	if !IsTimeout(&err) {
		t.Error("timeout error should cause true, not false")
	}
	err.Err = timeoutError{false}
	if IsTimeout(&err) {
		t.Error("non timeout error should cause false, not true")
	}
}

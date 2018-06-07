//replacer:generated-file

package mathh

import "testing"

func TestIsPositiveInfFloat64(t *testing.T) {
	if !IsPositiveInfFloat64(PositiveInfFloat64()) {
		t.Error("error in PositiveInfFloat64 or IsPositiveInfFloat64")
	}
}

func TestIsNegativeInfFloat64(t *testing.T) {
	if !IsNegativeInfFloat64(NegativeInfFloat64()) {
		t.Error("error in NegativeInfFloat64 or IsNegativeInfFloat64")
	}
}

func TestIsInfFloat64(t *testing.T) {
	if !IsInfFloat64(PositiveInfFloat64()) || !IsInfFloat64(NegativeInfFloat64()) {
		t.Error("error in PositiveInfFloat64, NegativeInfFloat64 or IsInfFloat64")
	}
}

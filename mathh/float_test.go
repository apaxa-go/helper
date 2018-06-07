package mathh

import "testing"

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old Float32
//replacer:new Float64

func TestIsPositiveInfFloat32(t *testing.T) {
	if !IsPositiveInfFloat32(PositiveInfFloat32()) {
		t.Error("error in PositiveInfFloat32 or IsPositiveInfFloat32")
	}
}

func TestIsNegativeInfFloat32(t *testing.T) {
	if !IsNegativeInfFloat32(NegativeInfFloat32()) {
		t.Error("error in NegativeInfFloat32 or IsNegativeInfFloat32")
	}
}

func TestIsInfFloat32(t *testing.T) {
	if !IsInfFloat32(PositiveInfFloat32()) || !IsInfFloat32(NegativeInfFloat32()) {
		t.Error("error in PositiveInfFloat32, NegativeInfFloat32 or IsInfFloat32")
	}
}

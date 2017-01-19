package constanth

import (
	"github.com/apaxa-go/helper/mathh"
	"go/constant"
	"go/token"
	"testing"
)

const veryLongNumber = "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

func TestFloat32Val(t *testing.T) {
	if _, ok := Float32Val(constant.MakeFromLiteral(veryLongNumber, token.FLOAT, 0)); ok {
		t.Error("expect not ok")
	}
}

func TestComplexVal(t *testing.T) {
	c1, _ := BinaryOp(constant.MakeFromLiteral(veryLongNumber, token.FLOAT, 0), token.ADD, constant.MakeFromLiteral("1i", token.IMAG, 0))
	c2, _ := BinaryOp(MakeInt(1), token.ADD, constant.MakeFromLiteral(veryLongNumber+"i", token.IMAG, 0))

	if c1.Kind() != constant.Complex || c2.Kind() != constant.Complex {
		t.Fatal("c1 or c2 not complex")
	}

	if _, ok := Complex64Val(c1); ok {
		t.Error("expect not ok")
	}
	if _, ok := Complex64Val(c2); ok {
		t.Error("expect not ok")
	}
	if _, ok := Complex128Val(c1); ok {
		t.Error("expect not ok")
	}
	if _, ok := Complex128Val(c2); ok {
		t.Error("expect not ok")
	}
}

func TestIntVal(t *testing.T) {
	c := MakeInt64(mathh.MaxInt64)

	if _, ok := IntVal(c); ok == (mathh.IntBits < 64) {
		t.Error("expect not ok")
	}
	if _, ok := Int8Val(c); ok {
		t.Error("expect not ok")
	}
	if _, ok := Int16Val(c); ok {
		t.Error("expect not ok")
	}
	if _, ok := Int32Val(c); ok {
		t.Error("expect not ok")
	}
	if _, ok := UintVal(c); ok == (mathh.IntBits < 64) {
		t.Error("expect not ok")
	}
	if _, ok := Uint8Val(c); ok {
		t.Error("expect not ok")
	}
	if _, ok := Uint16Val(c); ok {
		t.Error("expect not ok")
	}
	if _, ok := Uint32Val(c); ok {
		t.Error("expect not ok")
	}
}

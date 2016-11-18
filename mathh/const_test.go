package mathh

import (
	"math"
	"runtime"
	"testing"
)

func TestConsts(t *testing.T) {
	if Uint8Bytes != 1 || Int8Bytes != 1 || Uint16Bytes != 2 || Int16Bytes != 2 || Uint32Bytes != 4 || Int32Bytes != 4 || Uint64Bytes != 8 || Int64Bytes != 8 {
		t.Error("TestConsts: bad Bytes")
	}
	if Uint8Bits != 8 || Int8Bits != 8 || Uint16Bits != 16 || Int16Bits != 16 || Uint32Bits != 32 || Int32Bits != 32 || Uint64Bits != 64 || Int64Bits != 64 {
		t.Error("TestConsts: bad Bits")
	}
	if MinUint8 != 0 || MinInt8 != math.MinInt8 || MinUint16 != 0 || MinInt16 != math.MinInt16 || MinUint32 != 0 || MinInt32 != math.MinInt32 || MinUint64 != 0 || MinInt64 != math.MinInt64 {
		t.Error("TestConsts: bad Min")
	}
	if MaxUint8 != math.MaxUint8 || MaxInt8 != math.MaxInt8 || MaxUint16 != math.MaxUint16 || MaxInt16 != math.MaxInt16 || MaxUint32 != math.MaxUint32 || MaxInt32 != math.MaxInt32 || MaxUint64 != math.MaxUint64 || MaxInt64 != math.MaxInt64 {
		t.Error("TestConsts: bad Max")
	}

	////////////////

	switch runtime.GOARCH {
	case "386":
		if UintBytes != Uint32Bytes || IntBytes != Int32Bytes || UintBits != Uint32Bits || IntBits != Int32Bits || MinUint != MinUint32 || MinInt != MinInt32 || MaxUint != MaxUint32 || MaxInt != MaxInt32 {
			t.Error("TestConsts: bad Int")
		}
	case "amd64":
		if UintBytes != Uint64Bytes || IntBytes != Int64Bytes || UintBits != Uint64Bits || IntBits != Int64Bits || MinUint != MinUint64 || MinInt != MinInt64 || MaxUint != MaxUint64 || MaxInt != MaxInt64 {
			t.Error("TestConsts: bad Int")
		}
	default:

	}
}

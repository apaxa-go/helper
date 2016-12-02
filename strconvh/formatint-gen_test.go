//replacer:generated-file
package strconvh

import (
	"github.com/apaxa-go/helper/mathh"
	"testing"
)

func TestFormatUint(t *testing.T) {
	type testElement struct {
		i uint
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxUint, maxUintStr},
		{mathh.MinUint, minUintStr},
	}
	for _, v := range test {
		s := FormatUint(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatUint8(t *testing.T) {
	type testElement struct {
		i uint8
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxUint8, maxUint8Str},
		{mathh.MinUint8, minUint8Str},
	}
	for _, v := range test {
		s := FormatUint8(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatUint16(t *testing.T) {
	type testElement struct {
		i uint16
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxUint16, maxUint16Str},
		{mathh.MinUint16, minUint16Str},
	}
	for _, v := range test {
		s := FormatUint16(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatUint32(t *testing.T) {
	type testElement struct {
		i uint32
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxUint32, maxUint32Str},
		{mathh.MinUint32, minUint32Str},
	}
	for _, v := range test {
		s := FormatUint32(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatInt(t *testing.T) {
	type testElement struct {
		i int
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxInt, maxIntStr},
		{mathh.MinInt, minIntStr},
	}
	for _, v := range test {
		s := FormatInt(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatInt8(t *testing.T) {
	type testElement struct {
		i int8
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxInt8, maxInt8Str},
		{mathh.MinInt8, minInt8Str},
	}
	for _, v := range test {
		s := FormatInt8(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatInt16(t *testing.T) {
	type testElement struct {
		i int16
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxInt16, maxInt16Str},
		{mathh.MinInt16, minInt16Str},
	}
	for _, v := range test {
		s := FormatInt16(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatInt32(t *testing.T) {
	type testElement struct {
		i int32
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxInt32, maxInt32Str},
		{mathh.MinInt32, minInt32Str},
	}
	for _, v := range test {
		s := FormatInt32(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

func TestFormatInt64(t *testing.T) {
	type testElement struct {
		i int64
		s string
	}
	test := []testElement{
		{0, "0"},
		{10, "10"},
		{mathh.MaxInt64, maxInt64Str},
		{mathh.MinInt64, minInt64Str},
	}
	for _, v := range test {
		s := FormatInt64(v.i)
		if s != v.s {
			t.Errorf("expect %s, got %s", v.s, s)
		}
	}
}

package strconvh

import "testing"

func TestFormatFloat32(t *testing.T) {
	type testElement struct {
		i float32
		s string
	}
	tests := []testElement{
		{0, "0"},
		{.12345, "0.12345"},
		{-.12345, "-0.12345"},
		{1, "1"},
		{-1, "-1"},
		{1.12345, "1.12345"},
		{-1.12345, "-1.12345"},
		{123456, "123456"},
		{-123456, "-123456"},
		{12345.678, "12345.678"},
		{-12345.678, "-12345.678"},
	}
	for _, v := range tests {
		s := FormatFloat32(v.i)
		if s != v.s {
			t.Errorf("Expected: %s, got: %s", v.s, s)
		}
	}
}

func TestFormatFloat32Prec(t *testing.T) {
	type testElement struct {
		i float32
		p int
		s string
	}
	tests := []testElement{
		{0, 10, "0.0000000000"},
		{.12345, 6, "0.123450"},
		{-.12345, 5, "-0.12345"},
		{.12345, 4, "0.1235"},
		{1.12345, 0, "1"},
		{-1.6, 0, "-2"},
	}
	for _, v := range tests {
		s := FormatFloat32Prec(v.i, v.p)
		if s != v.s {
			t.Errorf("Expected: %s, got: %s", v.s, s)
		}
	}
}

func TestFormatFloat64(t *testing.T) {
	type testElement struct {
		i float64
		s string
	}
	test := []testElement{
		{0, "0"},
		{.12345, "0.12345"},
		{-.12345, "-0.12345"},
		{1, "1"},
		{-1, "-1"},
		{1.12345, "1.12345"},
		{-1.12345, "-1.12345"},
		{123456, "123456"},
		{-123456, "-123456"},
		{12345678.901, "12345678.901"},
		{-12345678.901, "-12345678.901"},
	}
	for _, v := range test {
		s := FormatFloat64(v.i)
		if s != v.s {
			t.Errorf("Expected: %s, got: %s", v.s, s)
		}
	}
}

func TestFormatFloat64Prec(t *testing.T) {
	type testElement struct {
		i float64
		p int
		s string
	}
	test := []testElement{
		{0, 10, "0.0000000000"},
		{.12345, 6, "0.123450"},
		{-.12345, 5, "-0.12345"},
		{.12345, 4, "0.1235"},
		{1.12345, 0, "1"},
		{-1.6, 0, "-2"},
	}
	for _, v := range test {
		s := FormatFloat64Prec(v.i, v.p)
		if s != v.s {
			t.Errorf("Expected: %s, got: %s", v.s, s)
		}
	}
}
//replacer:generated-file

package strconvh

import (
	"github.com/apaxa-go/helper/mathh"
	"testing"
)

func TestParseInt(t *testing.T) {

	type testElement struct {
		s   string
		i   int
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minIntStr, mathh.MinInt, false},
		{maxIntStr, mathh.MaxInt, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseInt(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseInt8(t *testing.T) {

	type testElement struct {
		s   string
		i   int8
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minInt8Str, mathh.MinInt8, false},
		{maxInt8Str, mathh.MaxInt8, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseInt8(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseInt16(t *testing.T) {

	type testElement struct {
		s   string
		i   int16
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minInt16Str, mathh.MinInt16, false},
		{maxInt16Str, mathh.MaxInt16, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseInt16(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseInt32(t *testing.T) {

	type testElement struct {
		s   string
		i   int32
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minInt32Str, mathh.MinInt32, false},
		{maxInt32Str, mathh.MaxInt32, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseInt32(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseUint(t *testing.T) {

	type testElement struct {
		s   string
		i   uint
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minUintStr, mathh.MinUint, false},
		{maxUintStr, mathh.MaxUint, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseUint(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseUint8(t *testing.T) {

	type testElement struct {
		s   string
		i   uint8
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minUint8Str, mathh.MinUint8, false},
		{maxUint8Str, mathh.MaxUint8, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseUint8(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseUint16(t *testing.T) {

	type testElement struct {
		s   string
		i   uint16
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minUint16Str, mathh.MinUint16, false},
		{maxUint16Str, mathh.MaxUint16, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseUint16(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseUint32(t *testing.T) {

	type testElement struct {
		s   string
		i   uint32
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minUint32Str, mathh.MinUint32, false},
		{maxUint32Str, mathh.MaxUint32, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseUint32(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

func TestParseUint64(t *testing.T) {

	type testElement struct {
		s   string
		i   uint64
		err bool
	}

	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{"10", 10, false},
		{"10.05", 0, true},
		{minUint64Str, mathh.MinUint64, false},
		{maxUint64Str, mathh.MaxUint64, false},
		{"18446744073709551616", 0, true}, // Maximum unsigned integer plus 1
		{"-9223372036854775809", 0, true}, // Minimum signed integer minus 1
	}

	for _, v := range test {
		r, err := ParseUint64(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got %v", v.i, r)
			}
		}
	}
}

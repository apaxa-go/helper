package strconvh

import "testing"

func TestFormatInt(t *testing.T) {
	type testElement struct {
		i int
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{9223372036854775807, "9223372036854775807"},
		testElement{-9223372036854775808, "-9223372036854775808"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatInt(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatInt8(t *testing.T) {
	type testElement struct {
		i int8
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{127, "127"},
		testElement{-128, "-128"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatInt8(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatInt16(t *testing.T) {
	type testElement struct {
		i int16
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{32767, "32767"},
		testElement{-32768, "-32768"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatInt16(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatInt32(t *testing.T) {
	type testElement struct {
		i int32
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{2147483647, "2147483647"},
		testElement{-2147483648, "-2147483648"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatInt32(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatInt64(t *testing.T) {
	type testElement struct {
		i int64
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{9223372036854775807, "9223372036854775807"},
		testElement{-9223372036854775808, "-9223372036854775808"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatInt64(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatUint(t *testing.T) {
	type testElement struct {
		i uint
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{18446744073709551615, "18446744073709551615"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatUint(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatUint8(t *testing.T) {
	type testElement struct {
		i uint8
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{255, "255"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatUint8(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatUint16(t *testing.T) {
	type testElement struct {
		i uint16
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{65535, "65535"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatUint16(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatUint32(t *testing.T) {
	type testElement struct {
		i uint32
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{4294967295, "4294967295"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatUint32(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

func TestFormatUint64(t *testing.T) {
	type testElement struct {
		i uint64
		s string
	}
	test := []testElement{
		testElement{10, "10"},
		testElement{18446744073709551615, "18446744073709551615"},
		testElement{0, "0"},
	}
	for i, v := range test {
		s := FormatUint64(v.i)
		if s != v.s {
			t.Errorf("Test-%v\nExpected string: %s, got: %s", i, v.s, s)
		}
	}
}

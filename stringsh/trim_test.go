package stringsh

import "testing"

func TestTrimLeftBytes(t *testing.T) {
	type testElement struct {
		s string
		b byte
		r string
	}
	tests := []testElement{
		{"", '0', ""},
		{"0", '0', ""},
		{"0000000000", '0', ""},
		{"12345", '0', "12345"},
		{"012345", '0', "12345"},
		{"00000123450", '0', "123450"},
	}
	for _, v := range tests {
		if r := TrimLeftBytes(v.s, v.b); r != v.r {
			t.Errorf("Expected %v, but got %v", v.r, r)
		}
	}
}

func TestTrimRightBytes(t *testing.T) {
	type testElement struct {
		s string
		b byte
		r string
	}
	tests := []testElement{
		{"", '0', ""},
		{"0", '0', ""},
		{"0000000000", '0', ""},
		{"12345", '0', "12345"},
		{"123450", '0', "12345"},
		{"01234500000", '0', "012345"},
	}
	for _, v := range tests {
		if r := TrimRightBytes(v.s, v.b); r != v.r {
			t.Errorf("Expected %v, but got %v", v.r, r)
		}
	}
}

func TestTrimBytes(t *testing.T) {
	type testElement struct {
		s string
		b byte
		r string
	}
	tests := []testElement{
		{"", '0', ""},
		{"0", '0', ""},
		{"0000000000", '0', ""},
		{"12345", '0', "12345"},
		{"123450", '0', "12345"},
		{"012345", '0', "12345"},
		{"01234500000", '0', "12345"},
		{"00000123450", '0', "12345"},
	}
	for _, v := range tests {
		if r := TrimBytes(v.s, v.b); r != v.r {
			t.Errorf("Expected %v, but got %v", v.r, r)
		}
	}
}

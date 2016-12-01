package stringsh

import "testing"

func TestPadLeft(t *testing.T) {
	type testElement struct {
		s string
		g string
		l int
		r string
	}
	tests := []testElement{
		{"", "ё", 0, ""},
		{"", "ё", 1, "ё"},
		{"", "ё", 10, "ёёёёёёёёёё"},
		{"ййййй", "ё", 0, "ййййй"},
		{"ййййй", "ё", 5, "ййййй"},
		{"ййййй", "ё", 6, "ёййййй"},
		{"ййййй", "ё", 10, "ёёёёёййййй"},
	}
	for _, v := range tests {
		if r := PadLeft(v.s, v.g, v.l); r != v.r {
			t.Errorf("expect %v, got %v", v.r, r)
		}
	}
}

func TestPadRight(t *testing.T) {
	type testElement struct {
		s string
		g string
		l int
		r string
	}
	tests := []testElement{
		{"", "ё", 0, ""},
		{"", "ё", 1, "ё"},
		{"", "ё", 10, "ёёёёёёёёёё"},
		{"ййййй", "ё", 0, "ййййй"},
		{"ййййй", "ё", 5, "ййййй"},
		{"ййййй", "ё", 6, "йййййё"},
		{"ййййй", "ё", 10, "йййййёёёёё"},
	}
	for _, v := range tests {
		if r := PadRight(v.s, v.g, v.l); r != v.r {
			t.Errorf("expect %v, got %v", v.r, r)
		}
	}
}

func TestPadLeftWithByte(t *testing.T) {
	type testElement struct {
		s string
		b byte
		l int
		r string
	}
	tests := []testElement{
		{"", '0', 0, ""},
		{"", '0', 1, "0"},
		{"", '0', 10, "0000000000"},
		{"12345", '0', 0, "12345"},
		{"12345", '0', 5, "12345"},
		{"12345", '0', 6, "012345"},
		{"12345", '0', 10, "0000012345"},
	}
	for _, v := range tests {
		if r := PadLeftWithByte(v.s, v.b, v.l); r != v.r {
			t.Errorf("expect %v, got %v", v.r, r)
		}
	}
}

func TestPadRightWithByte(t *testing.T) {
	type testElement struct {
		s string
		b byte
		l int
		r string
	}
	tests := []testElement{
		{"", '0', 0, ""},
		{"", '0', 1, "0"},
		{"", '0', 10, "0000000000"},
		{"12345", '0', 0, "12345"},
		{"12345", '0', 5, "12345"},
		{"12345", '0', 6, "123450"},
		{"12345", '0', 10, "1234500000"},
	}
	for _, v := range tests {
		if r := PadRightWithByte(v.s, v.b, v.l); r != v.r {
			t.Errorf("expect %v, got %v", v.r, r)
		}
	}
}

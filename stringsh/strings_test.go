package stringsh

import (
	"strings"
	"testing"
)

func TestLen(t *testing.T) {
	type testElement struct {
		s string
		l int
	}
	tests := []testElement{
		{"", 0},
		{"a", 1},
		{"abc", 3},
		{"Hello, 世界", 9},
		{"ẛ̣", 1},
		{"が", 1},
		{"ç", 1},
		{"й", 1},
		{"ё", 1},
	}
	for _, v := range tests {
		if r := Len(v.s); r != v.l {
			t.Errorf("string '%v', expected: %v, got: %v", v.s, v.l, r)
		}
	}
}

type firstLineTestElement struct {
	s   string
	l   string
	p   int
	rem string
}

var firstLineTests = []firstLineTestElement{
	{"", "", 0, ""},
	{"\r", "\r", 1, ""},
	{"\n", "", 1, ""},
	{"\r\n", "", 2, ""},
	{"1\n", "1", 2, ""},
	{"1\r", "1\r", 2, ""},
	{"1\r\n", "1", 3, ""},
	{"1\n2", "1", 2, "2"},
	{"1\r2", "1\r2", 3, ""},
	{"1\r\n2", "1", 3, "2"},
	{"\n2", "", 1, "2"},
	{"\r\n2", "", 2, "2"},
	{"\r2", "\r2", 2, ""},
	{"12\n34\n56\r\n78\r90", "12", 3, "34\n56\r\n78\r90"},
	{"12\r\n34\n56\r\n78\r90", "12", 4, "34\n56\r\n78\r90"},
	{"12\r34\n56\r\n78\r90", "12\r34", 6, "56\r\n78\r90"},
}

func TestGetLine(t *testing.T) {
	for _, v := range firstLineTests {
		if l, p := GetLine(v.s); l != v.l || p != v.p {
			s := strings.Replace(v.s, "\n", "\\n", -1)
			s = strings.Replace(s, "\r", "\\r", -1)
			t.Errorf("%v: expect '%v' %v, got '%v' %v", s, v.l, v.p, l, p)
		}
	}
}

func TestGetFirstLine(t *testing.T) {
	for _, v := range firstLineTests {
		if l := GetFirstLine(v.s); l != v.l {
			s := strings.Replace(v.s, "\n", "\\n", -1)
			s = strings.Replace(s, "\r", "\\r", -1)
			t.Errorf("%v: expect '%v', got '%v'", s, v.l, l)
		}
	}
}

func TestExtractLine(t *testing.T) {
	for _, v := range firstLineTests {
		if l, e := ExtractLine(v.s); l != v.l || e != v.rem {
			s := strings.Replace(v.s, "\n", "\\n", -1)
			s = strings.Replace(s, "\r", "\\r", -1)
			t.Errorf("%v: expect '%v' '%v', got '%v' '%v'", s, v.l, v.rem, l, e)
		}
	}
}

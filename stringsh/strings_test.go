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

func TestIndexMulti(t *testing.T) {
	type testElement struct {
		s        string
		seps     []string
		rI, rSep int
	}
	tests := []testElement{
		{"", []string{}, -1, -1},
		{"", []string{"d"}, -1, -1},
		{"", []string{"d", "ef"}, -1, -1},
		{"abc", []string{}, -1, -1},
		{"abc", []string{"d"}, -1, -1},
		{"abc", []string{"d", "ef"}, -1, -1},
		{"", []string{""}, 0, 0},
		{"", []string{"d", ""}, 0, 1},
		{"abcde", []string{"cd", "de"}, 2, 0},
		{"abcde", []string{"de", "cd"}, 2, 1},
		{"abcde", []string{""}, 0, 0},
		{"abcde", []string{"b", "bc"}, 1, 0},
		{"abcde", []string{"bc", "b"}, 1, 0},
	}
	for i, v := range tests {
		if rI, rSep := IndexMulti(v.s, v.seps...); rI != v.rI || rSep != v.rSep {
			t.Errorf("#%v expect %v %v, got %v %v", i, v.rI, v.rSep, rI, rSep)
		}
	}
}

func TestReplaceMulti(t *testing.T) {
	type testElement struct {
		s        string
		old, new []string
		r        string
	}
	tests := []testElement{
		{"", []string{}, []string{}, ""},
		{"", []string{"a"}, []string{"A"}, ""},
		{"", []string{"a", "b"}, []string{"A", "B"}, ""},
		{"abcdefghijklmnopqrstuwxyz", []string{}, []string{}, "abcdefghijklmnopqrstuwxyz"},
		{"abcdefghijklmnopqrstuwxyz", []string{"ac"}, []string{"AC"}, "abcdefghijklmnopqrstuwxyz"},
		{"abcdefghijklmnopqrstuwxyz", []string{"ac", "np"}, []string{"AC", "NP"}, "abcdefghijklmnopqrstuwxyz"},
		{"abcdefghijklmnopqrstuwxyz", []string{"a"}, []string{"A"}, "Abcdefghijklmnopqrstuwxyz"},
		{"abcdefghijklmnopqrstuwxyz", []string{"ab"}, []string{"ABC"}, "ABCcdefghijklmnopqrstuwxyz"},
		{"abcdefghijklmnopqrstuwxyz", []string{"ab", "yz"}, []string{"AB", "YZ"}, "ABcdefghijklmnopqrstuwxYZ"},
		{"abcdefghijklmnopqrstuwxyz", []string{"ab", "xz"}, []string{"AB", "XZ"}, "ABcdefghijklmnopqrstuwxyz"},
		{"abcdefghijklmnopqrstuwxyz", []string{"a", "ab"}, []string{"A", "AB"}, "Abcdefghijklmnopqrstuwxyz"},
		{"abcdefghijklmnopqrstuwxyz", []string{"ab", "a"}, []string{"AB", "A"}, "ABcdefghijklmnopqrstuwxyz"},
	}
	for i, v := range tests {
		if r := ReplaceMulti(v.s, v.old, v.new); r != v.r {
			t.Errorf("#%v expect '%v', got '%v'", i, v.r, r)
		}
	}
}

func TestReplaceMulti2(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expect panic")
		}
	}()
	ReplaceMulti("", []string{}, []string{""})
	t.Error("expect panic")
}

func TestReplaceMulti3(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expect panic")
		}
	}()
	ReplaceMulti("", []string{""}, []string{"A"})
	t.Error("expect panic")
}

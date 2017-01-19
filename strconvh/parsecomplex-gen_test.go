//replacer:generated-file

package strconvh

import "testing"

func TestParseComplex128(t *testing.T) {
	type testElement struct {
		s   string
		c   complex128
		err bool
	}

	tests := []testElement{
		{"", 0, true},
		{"i", 0, true},
		{"+", 0, true},
		{"-", 0, true},
		{"a", 0, true},
		{"(", 0, true},
		{")", 0, true},
		{"(0", 0, true},
		{"0)", 0, true},
		{"1+1", 0, true},
		{"123a45+1i", 0, true},
		{"1+123a45i", 0, true},
		{"0", 0, false},
		{"1.2", 1.2, false},
		{"+1.2", 1.2, false},
		{"-1.2", -1.2, false},
		{"3.4i", 3.4i, false},
		{"+3.4i", 3.4i, false},
		{"-3.4i", -3.4i, false},
		{"1.2-3.4i", 1.2 - 3.4i, false},
		{"1.2+3.4i", 1.2 + 3.4i, false},
		{"+1.2-3.4i", 1.2 - 3.4i, false},
		{"+1.2+3.4i", 1.2 + 3.4i, false},
		{"-1.2-3.4i", -1.2 - 3.4i, false},
		{"-1.2+3.4i", -1.2 + 3.4i, false},
		{"(1.2-3.4i)", 1.2 - 3.4i, false},
	}

	for _, test := range tests {
		c, err := ParseComplex128(test.s)
		if c != test.c || err != nil != test.err {
			t.Errorf("%v: expect %v %v, got %v %v", test.s, test.c, test.err, c, err)
		}
	}
}

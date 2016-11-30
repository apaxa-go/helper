package stringsh

import "testing"

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
			t.Errorf("String '%v', expected: %v, got: %v", v.s, v.l, r)
		}
	}
}

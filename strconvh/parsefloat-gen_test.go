//replacer:generated-file
package strconvh

import "testing"


func TestParseFloat64(t *testing.T) {
	type testElement struct {
		s   string
		i   float64
		err bool
	}
	test := []testElement{
		{"", 0, true},
		{"0", 0, false},
		{".12345", .12345, false},
		{"-.12345", -.12345, false},
		{"0.12345", .12345, false},
		{"-0.12345", -.12345, false},
		{"1", 1, false},
		{"-1", -1, false},
		{"1.12345", 1.12345, false},
		{"-1.12345", -1.12345, false},
		{"123456", 123456, false},
		{"-123456", -123456, false},
		{"12345.678", 12345.678, false},
		{"-12345.678", -12345.678, false},
	}
	for _, v := range test {
		r, err := ParseFloat64(v.s)
		if (err != nil) != v.err {
			t.Errorf("error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("expect %v, got: %v", v.i, r)
			}
		}
	}
}

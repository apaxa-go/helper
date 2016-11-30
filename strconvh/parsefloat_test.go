package strconvh

import "testing"

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old float32	Float32
//replacer:new float64	Float64

func TestParseFloat32(t *testing.T) {
	type testElement struct {
		s   string
		i   float32
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
		r, err := ParseFloat32(v.s)
		if (err != nil) != v.err {
			t.Errorf("Error expected: %v, got: %v", v.err, err)
		}
		if !v.err && (err == nil) {
			if r != v.i {
				t.Errorf("Wrong parse. Expected float32: %v, got: %v", v.i, r)
			}
		}
	}
}

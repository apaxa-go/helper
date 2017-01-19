package strconvh

import "testing"

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old 32	64
//replacer:new 64	128

func TestFormatComplex64(t *testing.T) {
	type testElement struct {
		c complex64
		r string
	}

	tests := []testElement{
		{0, "(0+0i)"},
		{1, "(1+0i)"},
		{-1, "(-1+0i)"},
		{0i, "(0+0i)"},
		{1i, "(0+1i)"},
		{-1i, "(0-1i)"},
		{1 + 2i, "(1+2i)"},
		{1 - 2i, "(1-2i)"},
		{-1 + 2i, "(-1+2i)"},
		{-1 - 2i, "(-1-2i)"},
		{0.123456 + 7890.123i, "(0.123456+7890.123i)"},
	}

	for _, test := range tests {
		r := FormatComplex64(test.c)
		if r != test.r {
			t.Errorf("%v: expect %v, got %v", test.c, test.r, r)
		}
	}
}

func TestFormatComplex64Prec(t *testing.T) {
	type testElement struct {
		c    complex64
		prec int
		r    string
	}

	tests := []testElement{
		{0, 2, "(0.00+0.00i)"},
		{1, 2, "(1.00+0.00i)"},
		{-1, 2, "(-1.00+0.00i)"},
		{0i, 2, "(0.00+0.00i)"},
		{1i, 2, "(0.00+1.00i)"},
		{-1i, 2, "(0.00-1.00i)"},
		{1 + 2i, 2, "(1.00+2.00i)"},
		{1 - 2i, 2, "(1.00-2.00i)"},
		{-1 + 2i, 2, "(-1.00+2.00i)"},
		{-1 - 2i, 2, "(-1.00-2.00i)"},
		{0.123456 + 7890.1235i, -1, "(0.123456+7890.1235i)"},
		{0.123456 + 7890.1235i, 0, "(0+7890i)"},
		{0.123456 + 7890.1235i, 1, "(0.1+7890.1i)"},
		{0.123456 + 7890.1235i, 2, "(0.12+7890.12i)"},
		{0.123456 + 7890.123i, 3, "(0.123+7890.123i)"},
	}

	for _, test := range tests {
		r := FormatComplex64Prec(test.c, test.prec)
		if r != test.r {
			t.Errorf("%v %v: expect %v, got %v", test.c, test.prec, test.r, r)
		}
	}
}

package asth

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestTokenPositionToPoint(t *testing.T) {
	pos := token.Position{Filename: "fn", Offset: 11, Line: 12, Column: 13}
	point := PointInSource{Offset: 11, Line: 12, Column: 13}
	r := TokenPositionToPoint(pos)
	if r != point {
		t.Errorf("expect %v, got %v", point, r)
	}
}

func TestPointInSource_IsValid(t *testing.T) {
	pointValid := PointInSource{Offset: 11, Line: 12, Column: 13}
	pointInvalid0 := PointInSource{Offset: 11, Line: 0, Column: 13}
	pointInvalid1 := PointInSource{}

	if !pointValid.IsValid() {
		t.Error("expect valid")
	}
	if pointInvalid0.IsValid() {
		t.Error("expect invalid")
	}
	if pointInvalid1.IsValid() {
		t.Error("expect invalid")
	}
}

func TestPointInSource_String(t *testing.T) {
	type testElement struct {
		p PointInSource
		r string
	}

	tests := []testElement{
		{PointInSource{}, "-"},
		{PointInSource{Offset: 100, Line: 11, Column: 12}, "11:12"},
	}

	for _, test := range tests {
		r := test.p.String()
		if r != test.r {
			t.Errorf("%#v: expect %v, got %v", test.p, test.r, r)
		}
	}
}

func TestMakePosition(t *testing.T) {
	type testElement struct {
		expr string
		r    Position
	}

	const fn = "expr-file-name"

	tests := []testElement{
		{"a", Position{fn, PointInSource{0, 1, 1}, PointInSource{}}},
		{"    a    ", Position{fn, PointInSource{4, 1, 5}, PointInSource{}}},
		{"\n    a    \n", Position{fn, PointInSource{5, 2, 5}, PointInSource{}}},
		{"a(2)", Position{fn, PointInSource{0, 1, 1}, PointInSource{3, 1, 4}}},
		{"    a(2)    ", Position{fn, PointInSource{4, 1, 5}, PointInSource{7, 1, 8}}},
		{"\n    a(2)    \n", Position{fn, PointInSource{5, 2, 5}, PointInSource{8, 2, 8}}},
		{"[]int{\n	1,\n	2,\n}", Position{fn, PointInSource{0, 1, 1}, PointInSource{15, 4, 1}}},
		{"    []int{\n	1,\n	2,\n}    ", Position{fn, PointInSource{4, 1, 5}, PointInSource{19, 4, 1}}},
		{"\n    []int{\n	1,\n	2,\n}    \n", Position{fn, PointInSource{5, 2, 5}, PointInSource{20, 5, 1}}},
	}

	for _, test := range tests {
		fs := token.NewFileSet()
		fs.AddFile(fn, -1, len(test.expr))

		expr, err := parser.ParseExprFrom(fs, fn, test.expr, 0)
		if err != nil {
			t.Error("invalid expr")
			continue
		}

		r := MakePosition(expr.Pos(), expr.End(), fs)
		if r != test.r {
			t.Errorf("src:\n\"\"\"\n%v\n\"\"\"\nexpect:\n%#v\ngot:\n%#v", test.expr, test.r, r)
		}

		r = NodePosition(expr, fs)
		if r != test.r {
			t.Errorf("src:\n\"\"\"\n%v\n\"\"\"\nexpect:\n%#v\ngot:\n%#v", test.expr, test.r, r)
		}
	}
}

func TestPosition_String(t *testing.T) {
	type testElement struct {
		p Position
		r string
	}

	const fn = "file-name"

	tests := []testElement{
		{Position{}, "-"},
		{Position{fn, PointInSource{}, PointInSource{}}, fn},
		{Position{"", PointInSource{Offset: 100, Line: 11, Column: 12}, PointInSource{}}, "11:12"},
		{Position{"", PointInSource{}, PointInSource{Offset: 100, Line: 11, Column: 12}}, "-"}, // Malformed Position (invalid Pos & valid End)
		{Position{"", PointInSource{Offset: 100, Line: 0, Column: 12}, PointInSource{}}, "-"},  // Invalid Pos
		{Position{"", PointInSource{Offset: 100, Line: 11, Column: 12}, PointInSource{Offset: 200, Line: 21, Column: 22}}, "11:12-21:22"},
		{Position{fn, PointInSource{Offset: 100, Line: 11, Column: 12}, PointInSource{}}, fn + ":11:12"},
		{Position{fn, PointInSource{Offset: 100, Line: 11, Column: 12}, PointInSource{Offset: 200, Line: 21, Column: 22}}, fn + ":11:12-21:22"},
	}

	for _, test := range tests {
		r := test.p.String()
		if r != test.r {
			t.Errorf("%#v: expect %v, got %v", test.p, test.r, r)
		}
	}
}

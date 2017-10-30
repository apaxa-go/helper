package boundaryh

import "testing"

func TestBoundary_Len(t *testing.T) {
	b := Boundary{1, 4}
	if b.Len() != 3 {
		t.Error("Wrong boundary length.")
	}
	b = Invalid()
	if b.Len() != 0 {
		t.Error(`Wrong length for "Invalid" boundary`)
	}
}

func TestBoundary_IsValid(t *testing.T) {
	type testElement struct {
		b Boundary
		v bool // valid
	}
	tests := []testElement{
		{Boundary{0, 0}, true},
		{Boundary{0, 1}, true},
		{Boundary{1, 0}, false},
		{Boundary{1, 1}, true},
		{Boundary{0, 10}, true},
		{Boundary{10, 0}, false},
		{Boundary{10, 10}, true},
		{Boundary{1, 10}, true},
		{Boundary{10, 1}, false},

		{Boundary{-1, 0}, false},
		{Boundary{0, -1}, false},
		{Boundary{-1, -1}, false},
		{Boundary{-10, 0}, false},
		{Boundary{0, -10}, false},
		{Boundary{-10, -10}, false},
		{Boundary{-1, -10}, false},
		{Boundary{-10, -1}, false},

		{Boundary{-1, 1}, false},
		{Boundary{1, -1}, false},
		{Boundary{-10, 10}, false},
		{Boundary{10, -10}, false},
	}
	for _, test := range tests {
		if v := test.b.IsValid(); v != test.v {
			t.Errorf("%#v: expect %v, got %v", test.b, test.v, v)
		}
	}
}

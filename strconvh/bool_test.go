package strconvh

import "testing"

func TestFormatBool(t *testing.T) {
	if FormatBool(false) != "false" || FormatBool(true) != "true" {
		t.Error("invalid string from FormatBool")
	}
}

func TestParseBool(t *testing.T) {
	type testElement struct {
		bStr string
		b    bool
		err  bool
	}

	//1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
	tests := []testElement{
		// Positive tests
		{"1", true, false},
		{"t", true, false},
		{"T", true, false},
		{"TRUE", true, false},
		{"true", true, false},
		{"True", true, false},
		{"0", false, false},
		{"f", false, false},
		{"F", false, false},
		{"FALSE", false, false},
		{"false", false, false},
		{"False", false, false},
		// Negative tests
		{"-1", false, true},
		{"2", false, true},
		{"A", false, true},
		{"trUe", false, true},
		{"falsE", false, true},
	}

	for _, test := range tests {
		b, err := ParseBool(test.bStr)
		if b != test.b || err != nil != test.err {
			t.Errorf("%v: expect %v %v, got %v %v", test.bStr, test.b, test.err, b, err)
		}
	}
}

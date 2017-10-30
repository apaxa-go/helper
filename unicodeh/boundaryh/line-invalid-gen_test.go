//replacer:generated-file

package boundaryh

import "testing"

func TestLineInvalidArgsInString(t *testing.T) {
	//
	// Invalid pos arguments
	//
	type testElement struct {
		s   string
		pos int
	}
	tests := []testElement{
		{"ab", -2},
		{"ab", -1},
		{"ab", 2},
		{"ab", 3},
	}
	for testI, test := range tests {
		r := LineBreakAfterInString(test.s, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, InvalidPos, r)
		}
		r = LineBreakBeforeInString(test.s, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, InvalidPos, r)
		}
	}

	//
	// Empty string
	//
	for _, s := range []string{""} {
		r := FirstLineBreakInString(s)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", s, InvalidPos, r)
		}
		r = LastLineBreakInString(s)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", s, InvalidPos, r)
		}
		ris := LineBreaksInString(s)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", s, ris, cap(ris))
		}
	}

}

func TestLineInvalidArgs(t *testing.T) {
	//
	// Invalid pos arguments
	//
	type testElement struct {
		bytes []byte
		pos   int
	}
	tests := []testElement{
		{[]byte{'a', 'b'}, -2},
		{[]byte{'a', 'b'}, -1},
		{[]byte{'a', 'b'}, 2},
		{[]byte{'a', 'b'}, 3},
	}
	for testI, test := range tests {
		r := LineBreakAfter(test.bytes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, InvalidPos, r)
		}
		r = LineBreakBefore(test.bytes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, InvalidPos, r)
		}
	}

	//
	// Empty string
	//
	for _, bytes := range [][]byte{nil, {}} {
		r := FirstLineBreak(bytes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", bytes, InvalidPos, r)
		}
		r = LastLineBreak(bytes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", bytes, InvalidPos, r)
		}
		ris := LineBreaks(bytes)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", bytes, ris, cap(ris))
		}
	}

}

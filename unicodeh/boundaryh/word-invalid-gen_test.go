//replacer:generated-file

package boundaryh

import "testing"

func TestWordInvalidArgsInString(t *testing.T) {
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
		r := WordBeginInString(test.s, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, InvalidPos, r)
		}
		r = WordEndInString(test.s, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, InvalidPos, r)
		}
		rb := WordAtInString(test.s, test.pos)
		if !rb.IsInvalid() {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, Invalid(), rb)
		}
	}

	//
	// Empty string
	//
	for _, s := range []string{""} {
		r := FirstWordInString(s)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", s, InvalidPos, r)
		}
		r = LastWordInString(s)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", s, InvalidPos, r)
		}
		rbs := WordsInString(s)
		if len(rbs) != 0 || cap(rbs) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", s, rbs, cap(rbs))
		}
		ris := WordBreaksInString(s)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", s, ris, cap(ris))
		}
	}

}

func TestWordInvalidArgs(t *testing.T) {
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
		r := WordBegin(test.bytes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, InvalidPos, r)
		}
		r = WordEnd(test.bytes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, InvalidPos, r)
		}
		rb := WordAt(test.bytes, test.pos)
		if !rb.IsInvalid() {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, Invalid(), rb)
		}
	}

	//
	// Empty string
	//
	for _, bytes := range [][]byte{nil, {}} {
		r := FirstWord(bytes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", bytes, InvalidPos, r)
		}
		r = LastWord(bytes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", bytes, InvalidPos, r)
		}
		rbs := Words(bytes)
		if len(rbs) != 0 || cap(rbs) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", bytes, rbs, cap(rbs))
		}
		ris := WordBreaks(bytes)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", bytes, ris, cap(ris))
		}
	}

}

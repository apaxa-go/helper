//replacer:generated-file

package boundaryh

import "testing"

func TestSentenceInvalidArgsInString(t *testing.T) {
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
		r := SentenceBeginInString(test.s, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, InvalidPos, r)
		}
		r = SentenceEndInString(test.s, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, InvalidPos, r)
		}
		rb := SentenceAtInString(test.s, test.pos)
		if !rb.IsInvalid() {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.s, test.pos, Invalid(), rb)
		}
	}

	//
	// Empty string
	//
	for _, s := range []string{""} {
		r := FirstSentenceInString(s)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", s, InvalidPos, r)
		}
		r = LastSentenceInString(s)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", s, InvalidPos, r)
		}
		rbs := SentencesInString(s)
		if len(rbs) != 0 || cap(rbs) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", s, rbs, cap(rbs))
		}
		ris := SentenceBreaksInString(s)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", s, ris, cap(ris))
		}
	}

}

func TestSentenceInvalidArgs(t *testing.T) {
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
		r := SentenceBegin(test.bytes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, InvalidPos, r)
		}
		r = SentenceEnd(test.bytes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, InvalidPos, r)
		}
		rb := SentenceAt(test.bytes, test.pos)
		if !rb.IsInvalid() {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.bytes, test.pos, Invalid(), rb)
		}
	}

	//
	// Empty string
	//
	for _, bytes := range [][]byte{nil, {}} {
		r := FirstSentence(bytes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", bytes, InvalidPos, r)
		}
		r = LastSentence(bytes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", bytes, InvalidPos, r)
		}
		rbs := Sentences(bytes)
		if len(rbs) != 0 || cap(rbs) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", bytes, rbs, cap(rbs))
		}
		ris := SentenceBreaks(bytes)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", bytes, ris, cap(ris))
		}
	}

}

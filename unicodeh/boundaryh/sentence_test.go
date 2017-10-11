package boundaryh

import "testing"

func TestSentenceInvalidArgs(t *testing.T) {
	//
	// Invalid pos arguments
	//
	type testElement struct {
		runes []rune
		pos   int
	}
	tests := []testElement{
		{[]rune{0, 1}, -2},
		{[]rune{0, 1}, -1},
		{[]rune{0, 1}, 2},
		{[]rune{0, 1}, 3},
	}
	for testI, test := range tests {
		r := SentenceBegin(test.runes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.runes, test.pos, InvalidPos, r)
		}
		r = SentenceEnd(test.runes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.runes, test.pos, InvalidPos, r)
		}
		rb := SentenceAt(test.runes, test.pos)
		if !rb.IsInvalid() {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.runes, test.pos, Invalid(), rb)
		}
	}

	//
	// Empty string
	//
	for _, runes := range [][]rune{nil, {}} {
		r := FirstSentence(runes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", runes, InvalidPos, r)
		}
		r = LastSentence(runes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", runes, InvalidPos, r)
		}
		rbs := Sentences(runes)
		if len(rbs) != 0 || cap(rbs) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", runes, rbs, cap(rbs))
		}
		ris := SentenceBreaks(runes)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", runes, ris, cap(ris))
		}
	}

}

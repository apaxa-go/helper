package boundaryh

import "testing"

//replacer:ignore
// TODO replace windows path separator
//go:generate go run $GOPATH\src\github.com\apaxa-go\generator\replacer\main.go -- $GOFILE
//replacer:replace
//replacer:old InRunes	"[]rune{'a', 'b'}"	"{nil, {}}"	[]rune	Runes	runes
//replacer:new InString	'"ab"'				'{""}'		string	String	s
//replacer:new ""		"[]byte{'a', 'b'}"	"{nil, {}}"	[]byte	Bytes	bytes

func TestGraphemeClusterInvalidArgsInRunes(t *testing.T) {
	//
	// Invalid pos arguments
	//
	type testElement struct {
		runes []rune
		pos   int
	}
	tests := []testElement{
		{[]rune{'a', 'b'}, -2},
		{[]rune{'a', 'b'}, -1},
		{[]rune{'a', 'b'}, 2},
		{[]rune{'a', 'b'}, 3},
	}
	for testI, test := range tests {
		r := GraphemeClusterBeginInRunes(test.runes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.runes, test.pos, InvalidPos, r)
		}
		r = GraphemeClusterEndInRunes(test.runes, test.pos)
		if r != InvalidPos {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.runes, test.pos, InvalidPos, r)
		}
		rb := GraphemeClusterAtInRunes(test.runes, test.pos)
		if !rb.IsInvalid() {
			t.Errorf("%v \"%v\" [%v]: expect %v, got %v", testI, test.runes, test.pos, Invalid(), rb)
		}
	}

	//
	// Empty string
	//
	for _, runes := range [][]rune{nil, {}} {
		r := FirstGraphemeClusterInRunes(runes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", runes, InvalidPos, r)
		}
		r = LastGraphemeClusterInRunes(runes)
		if r != InvalidPos {
			t.Errorf("\"%v\": expect %v, got %v", runes, InvalidPos, r)
		}
		rbs := GraphemeClustersInRunes(runes)
		if len(rbs) != 0 || cap(rbs) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", runes, rbs, cap(rbs))
		}
		ris := GraphemeClusterBreaksInRunes(runes)
		if len(ris) != 0 || cap(ris) != 0 {
			t.Errorf("\"%v\": expect empty slice with cap=0, got %v with cap=%v", runes, ris, cap(ris))
		}
	}
}

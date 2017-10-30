package boundaryh

import (
	"reflect"
	"testing"
)

//replacer:ignore
// TODO replace windows path separator
//go:generate go run $GOPATH\src\github.com\apaxa-go\generator\replacer\main.go -- $GOFILE
//replacer:replace
//replacer:old InRunes	[]rune	test.runes	runes
//replacer:new InString	string	test.runes	s
//replacer:new ""		[]byte	test.runes	bytes

func TestGraphemeClustersInRunes(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInRunes()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := GraphemeClustersInRunes(test.SampleInRunes())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastGraphemeClusterInRunes(t *testing.T) {
	var stat Stat

	// Same as function "GraphemeClustersInRunes", but going from end to begin.
	revGCs := func(runes []rune) (boundaries []Boundary) {
		for len(runes) > 0 {
			pos := LastGraphemeClusterInRunes(runes)
			boundaries = append([]Boundary{{pos, len(runes)}}, boundaries...)
			runes = runes[:pos]
		}
		return
	}

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInRunes()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := revGCs(test.SampleInRunes())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestGraphemeClusterAtInRunes(t *testing.T) {
	var stat Stat

	in := func(b Boundary, bs []Boundary) bool {
		for _, b1 := range bs {
			if b == b1 {
				return true
			}
		}
		return false
	}

	for testI, test := range ucdGraphemeClusterTests {
		expBoundaries, skip := test.breaksToBoundariesInRunes()
		if skip {
			l := len(test.SampleInRunes())
			stat.Add(l)
			stat.Skip(l)
			continue
		}
		for pos := 0; pos < len(test.SampleInRunes()); pos++ {
			stat.Add()
			b := GraphemeClusterAtInRunes(test.SampleInRunes(), pos)
			if b.From > pos || b.To <= pos {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: invalid boundary %v", testI, test.runes, pos, b)
				continue
			}
			if !in(b, expBoundaries) {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: wrong boundary %v, possible %v", testI, test.runes, pos, b, expBoundaries)
			}
		}
	}

	stat.Log(t)
}

func TestGraphemeClusterBreaksInRunes(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBreaks, skip := test.BreaksInRunes()
		if skip {
			stat.Skip()
			continue
		}
		breaks := GraphemeClusterBreaksInRunes(test.SampleInRunes())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

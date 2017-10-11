package boundaryh

import (
	"reflect"
	"testing"
)

type ucdWordTest ucdTest

func TestWords(t *testing.T) {
	var stat Stat

	for testI, test := range ucdWordTests {
		stat.Add()
		boundaries := Words(test.runes)
		expBoundaries := breaksToBoundaries(test.breaks)
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastWord(t *testing.T) {
	var stat Stat

	// Same as function "Words", but going from end to begin.
	revWs := func(runes []rune) (boundaries []Boundary) {
		boundaries = make([]Boundary, 0, len(runes))
		for len(runes) > 0 {
			pos := LastWord(runes)
			boundaries = append([]Boundary{{pos, len(runes)}}, boundaries...)
			runes = runes[:pos]
		}
		return
	}

	for testI, test := range ucdWordTests {
		stat.Add()
		boundaries := revWs(test.runes)
		expBoundaries := breaksToBoundaries(test.breaks)
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestWordAt(t *testing.T) {
	var stat Stat

	in := func(b Boundary, bs []Boundary) bool {
		for _, b1 := range bs {
			if b == b1 {
				return true
			}
		}
		return false
	}

	for testI, test := range ucdWordTests {
		expBoundaries := breaksToBoundaries(test.breaks)
		for runeI := range test.runes {
			stat.Add()
			b := WordAt(test.runes, runeI)
			if b.From > runeI || b.To <= runeI {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: invalid boundary %v", testI, test.runes, runeI, b)
				continue
			}
			if !in(b, expBoundaries) {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: wrong boundary %v, possible %v", testI, test.runes, runeI, b, expBoundaries)
			}
		}
	}

	stat.Log(t)
}

func TestWordBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdWordTests {
		stat.Add()
		breaks := WordBreaks(test.runes)
		if !reflect.DeepEqual(breaks, test.breaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

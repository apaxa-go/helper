package boundaryh

import (
	"reflect"
	"testing"
)

type ucdSentenceTest ucdTest

func TestSentences(t *testing.T) {
	var stat Stat

	for testI, test := range ucdSentenceTests {
		stat.Add()
		boundaries := Sentences(test.runes)
		expBoundaries := breaksToBoundaries(test.breaks)
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, (test.runes), expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastSentence(t *testing.T) {
	var stat Stat

	// Same as function "Sentences", but going from end to begin.
	revSs := func(runes []rune) (boundaries []Boundary) {
		boundaries = make([]Boundary, 0, len(runes))
		for len(runes) > 0 {
			pos := LastSentence(runes)
			boundaries = append([]Boundary{{pos, len(runes)}}, boundaries...)
			runes = runes[:pos]
		}
		return
	}

	for testI, test := range ucdSentenceTests {
		stat.Add()
		boundaries := revSs(test.runes)
		expBoundaries := breaksToBoundaries(test.breaks)
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestSentenceAt(t *testing.T) {
	var stat Stat

	in := func(b Boundary, bs []Boundary) bool {
		for _, b1 := range bs {
			if b == b1 {
				return true
			}
		}
		return false
	}

	for testI, test := range ucdSentenceTests {
		expBoundaries := breaksToBoundaries(test.breaks)
		for runeI := range test.runes {
			stat.Add()
			b := SentenceAt(test.runes, runeI)
			if b.From > runeI || b.To <= runeI {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: invalid boundary %v", testI, test.runes, runeI, b)
				continue
			}
			if !in(b, expBoundaries) {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: wrong boundary %v, possible are %v", testI, test.runes, runeI, b, expBoundaries)
			}
		}
	}

	stat.Log(t)
}

func TestSentenceBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdSentenceTests {
		stat.Add()
		breaks := SentenceBreaks(test.runes)
		if !reflect.DeepEqual(breaks, test.breaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}
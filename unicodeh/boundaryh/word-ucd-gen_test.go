//replacer:generated-file

package boundaryh

import (
	"reflect"
	"testing"
)

func TestWordsInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdWordTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := WordsInString(test.SampleInString())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastWordInString(t *testing.T) {
	var stat Stat

	// Same as function "WordsInString", but going from end to begin.
	revWs := func(s string) (boundaries []Boundary) {
		boundaries = make([]Boundary, 0, len(s))
		for len(s) > 0 {
			pos := LastWordInString(s)
			boundaries = append([]Boundary{{pos, len(s)}}, boundaries...)
			s = s[:pos]
		}
		return
	}

	for testI, test := range ucdWordTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := revWs(test.SampleInString())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestWordAtInString(t *testing.T) {
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
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			l := len(test.SampleInString()) // TODO
			stat.Add(l)
			stat.Skip(l)
			continue
		}
		for pos := 0; pos < len(test.SampleInString()); pos++ {
			stat.Add()
			b := WordAtInString(test.SampleInString(), pos)
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

func TestWordBreaksInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdWordTests {
		stat.Add()
		expBreaks, skip := test.BreaksInString()
		if skip {
			stat.Skip()
			continue
		}
		breaks := WordBreaksInString(test.SampleInString())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestWords(t *testing.T) {
	var stat Stat

	for testI, test := range ucdWordTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := Words(test.Sample())
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
	revWs := func(bytes []byte) (boundaries []Boundary) {
		boundaries = make([]Boundary, 0, len(bytes))
		for len(bytes) > 0 {
			pos := LastWord(bytes)
			boundaries = append([]Boundary{{pos, len(bytes)}}, boundaries...)
			bytes = bytes[:pos]
		}
		return
	}

	for testI, test := range ucdWordTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := revWs(test.Sample())
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
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			l := len(test.Sample()) // TODO
			stat.Add(l)
			stat.Skip(l)
			continue
		}
		for pos := 0; pos < len(test.Sample()); pos++ {
			stat.Add()
			b := WordAt(test.Sample(), pos)
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

func TestWordBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdWordTests {
		stat.Add()
		expBreaks, skip := test.Breaks()
		if skip {
			stat.Skip()
			continue
		}
		breaks := WordBreaks(test.Sample())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

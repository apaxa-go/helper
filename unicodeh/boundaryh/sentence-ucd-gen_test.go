//replacer:generated-file

package boundaryh

import (
	"reflect"
	"testing"
)

func TestSentencesInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdSentenceTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := SentencesInString(test.SampleInString())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastSentenceInString(t *testing.T) {
	var stat Stat

	// Same as function "SentencesInString", but going from end to begin.
	revSs := func(s string) (boundaries []Boundary) {
		boundaries = make([]Boundary, 0, len(s))
		for len(s) > 0 {
			pos := LastSentenceInString(s)
			boundaries = append([]Boundary{{pos, len(s)}}, boundaries...)
			s = s[:pos]
		}
		return
	}

	for testI, test := range ucdSentenceTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := revSs(test.SampleInString())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestSentenceAtInString(t *testing.T) {
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
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			l := len(test.SampleInString())
			stat.Add(l)
			stat.Skip(l)
			continue
		}
		for pos := 0; pos < len(test.SampleInString()); pos++ {
			stat.Add()
			b := SentenceAtInString(test.SampleInString(), pos)
			if b.From > pos || b.To <= pos {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: invalid boundary %v", testI, test.runes, pos, b)
				continue
			}
			if !in(b, expBoundaries) {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: wrong boundary %v, possible are %v", testI, test.runes, pos, b, expBoundaries)
			}
		}
	}

	stat.Log(t)
}

func TestSentenceBreaksInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdSentenceTests {
		stat.Add()
		expBreaks, skip := test.BreaksInString()
		if skip {
			stat.Skip()
			continue
		}
		breaks := SentenceBreaksInString(test.SampleInString())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBreaks, breaks)
		}
	}

	stat.Log(t)
}

func TestSentences(t *testing.T) {
	var stat Stat

	for testI, test := range ucdSentenceTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := Sentences(test.Sample())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastSentence(t *testing.T) {
	var stat Stat

	// Same as function "Sentences", but going from end to begin.
	revSs := func(bytes []byte) (boundaries []Boundary) {
		boundaries = make([]Boundary, 0, len(bytes))
		for len(bytes) > 0 {
			pos := LastSentence(bytes)
			boundaries = append([]Boundary{{pos, len(bytes)}}, boundaries...)
			bytes = bytes[:pos]
		}
		return
	}

	for testI, test := range ucdSentenceTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := revSs(test.Sample())
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
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			l := len(test.Sample())
			stat.Add(l)
			stat.Skip(l)
			continue
		}
		for pos := 0; pos < len(test.Sample()); pos++ {
			stat.Add()
			b := SentenceAt(test.Sample(), pos)
			if b.From > pos || b.To <= pos {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: invalid boundary %v", testI, test.runes, pos, b)
				continue
			}
			if !in(b, expBoundaries) {
				stat.Fail()
				t.Errorf("%v \"%v\" [%v]: wrong boundary %v, possible are %v", testI, test.runes, pos, b, expBoundaries)
			}
		}
	}

	stat.Log(t)
}

func TestSentenceBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdSentenceTests {
		stat.Add()
		expBreaks, skip := test.Breaks()
		if skip {
			stat.Skip()
			continue
		}
		breaks := SentenceBreaks(test.Sample())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBreaks, breaks)
		}
	}

	stat.Log(t)
}

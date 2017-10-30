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

func TestLineBreaksInRunes(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		stat.Add()
		expBreaks, skip := test.BreaksInRunes()
		if skip {
			stat.Skip()
			continue
		}
		breaks := LineBreaksInRunes(test.SampleInRunes())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLastLineBreakInRunes(t *testing.T) {
	var stat Stat

	// Same as function "LineBreaksInRunes", but going from end to begin.
	revLBs := func(runes []rune) (breaks []int) {
		breaks = []int{len(runes)}
		for len(runes) > 0 {
			pos := LastLineBreakInRunes(runes)
			if pos == NoLineBreak {
				break
			}
			breaks = append([]int{pos}, breaks...)
			runes = runes[:pos]
		}
		return
	}

	for testI, test := range ucdLineTests {
		stat.Add()
		expBreaks, skip := test.BreaksInRunes()
		if skip {
			stat.Skip()
			continue
		}
		breaks := revLBs(test.SampleInRunes())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLineBreakBeforeAfterInRunes(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		expBreaks, skip := test.BreaksInRunes()
		if skip {
			lastBreak := expBreaks[len(expBreaks)-1]
			stat.Add(lastBreak)
			stat.Skip(lastBreak)
			continue
		}
		breakPrev := 0
		for _, breakNext := range expBreaks {
			for pos := breakPrev; pos < breakNext; pos++ {
				stat.Add()
				breakBefore := LineBreakBeforeInRunes(test.SampleInRunes(), pos)
				if breakBefore != breakPrev {
					stat.Fail()
					t.Errorf("%v \"%v\" before %v: expect %v, got %v", testI, test.runes, pos, breakPrev, breakBefore)
				}
				stat.Add()
				breakAfter := LineBreakAfterInRunes(test.SampleInRunes(), pos)
				if breakAfter != breakNext {
					stat.Fail()
					t.Errorf("%v \"%v\" after %v: expect %v, got %v", testI, test.runes, pos, breakNext, breakAfter)
				}
			}
			breakPrev = breakNext
		}
	}

	stat.Log(t)
}

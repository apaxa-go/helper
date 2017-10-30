//replacer:generated-file

package boundaryh

import (
	"reflect"
	"testing"
)

func TestLineBreaksInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		stat.Add()
		expBreaks, skip := test.BreaksInString()
		if skip {
			stat.Skip()
			continue
		}
		breaks := LineBreaksInString(test.SampleInString())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLastLineBreakInString(t *testing.T) {
	var stat Stat

	// Same as function "LineBreaksInString", but going from end to begin.
	revLBs := func(s string) (breaks []int) {
		breaks = []int{len(s)}
		for len(s) > 0 {
			pos := LastLineBreakInString(s)
			if pos == NoLineBreak {
				break
			}
			breaks = append([]int{pos}, breaks...)
			s = s[:pos]
		}
		return
	}

	for testI, test := range ucdLineTests {
		stat.Add()
		expBreaks, skip := test.BreaksInString()
		if skip {
			stat.Skip()
			continue
		}
		breaks := revLBs(test.SampleInString())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLineBreakBeforeAfterInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		expBreaks, skip := test.BreaksInString()
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
				breakBefore := LineBreakBeforeInString(test.SampleInString(), pos)
				if breakBefore != breakPrev {
					stat.Fail()
					t.Errorf("%v \"%v\" before %v: expect %v, got %v", testI, test.runes, pos, breakPrev, breakBefore)
				}
				stat.Add()
				breakAfter := LineBreakAfterInString(test.SampleInString(), pos)
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

func TestLineBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		stat.Add()
		expBreaks, skip := test.Breaks()
		if skip {
			stat.Skip()
			continue
		}
		breaks := LineBreaks(test.Sample())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLastLineBreak(t *testing.T) {
	var stat Stat

	// Same as function "LineBreaks", but going from end to begin.
	revLBs := func(bytes []byte) (breaks []int) {
		breaks = []int{len(bytes)}
		for len(bytes) > 0 {
			pos := LastLineBreak(bytes)
			if pos == NoLineBreak {
				break
			}
			breaks = append([]int{pos}, breaks...)
			bytes = bytes[:pos]
		}
		return
	}

	for testI, test := range ucdLineTests {
		stat.Add()
		expBreaks, skip := test.Breaks()
		if skip {
			stat.Skip()
			continue
		}
		breaks := revLBs(test.Sample())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLineBreakBeforeAfter(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		expBreaks, skip := test.Breaks()
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
				breakBefore := LineBreakBefore(test.Sample(), pos)
				if breakBefore != breakPrev {
					stat.Fail()
					t.Errorf("%v \"%v\" before %v: expect %v, got %v", testI, test.runes, pos, breakPrev, breakBefore)
				}
				stat.Add()
				breakAfter := LineBreakAfter(test.Sample(), pos)
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

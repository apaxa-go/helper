package boundaryh

import (
	"reflect"
	"testing"
)

type ucdLineTest ucdTest

func TestLineBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		stat.Add()
		breaks := LineBreaks(test.runes)
		if !reflect.DeepEqual(test.breaks, breaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLastLineBreak(t *testing.T) {
	var stat Stat

	// Same as function "LineBreaks", but going from end to begin.
	revLBs := func(runes []rune) (breaks []int) {
		breaks = []int{len(runes)}
		for len(runes) > 0 {
			pos := LastLineBreak(runes)
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
		breaks := revLBs(test.runes)
		if !reflect.DeepEqual(test.breaks, breaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestLineBreakBeforeAfter(t *testing.T) {
	var stat Stat

	for testI, test := range ucdLineTests {
		start := 0
		for _, breakV := range test.breaks {
			for runeI := start; runeI < breakV; runeI++ {
				stat.Add()
				breakBefore := LineBreakBefore(test.runes, runeI)
				if breakBefore != start {
					stat.Fail()
					t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, start, breakBefore)
				}
				stat.Add()
				breakAfter := LineBreakAfter(test.runes, runeI)
				if breakAfter != breakV {
					stat.Fail()
					t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, breakV, breakAfter)
				}
			}
			start = breakV
		}
	}

	stat.Log(t)
}

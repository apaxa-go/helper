package boundaryh

import (
	"reflect"
	"testing"
)

type ucdGraphemeClusterTest ucdTest

func TestGraphemeClusters(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		boundaries := GraphemeClusters(test.runes)
		expBoundaries := breaksToBoundaries(test.breaks)
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastGraphemeCluster(t *testing.T) {
	var stat Stat

	// Same as function "GraphemeClusters", but going from end to begin.
	revGCs := func(runes []rune) (boundaries []Boundary) {
		for len(runes) > 0 {
			pos := LastGraphemeCluster(runes)
			boundaries = append([]Boundary{{pos, len(runes)}}, boundaries...)
			runes = runes[:pos]
		}
		return
	}

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBoundaries := breaksToBoundaries(test.breaks)
		boundaries := revGCs(test.runes)
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestGraphemeClusterAt(t *testing.T) {
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
		expBoundaries := breaksToBoundaries(test.breaks)
		for runeI := range test.runes {
			stat.Add()
			b := GraphemeClusterAt(test.runes, runeI)
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

func TestGraphemeClusterBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		breaks := GraphemeClusterBreaks(test.runes)
		if !reflect.DeepEqual(breaks, test.breaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

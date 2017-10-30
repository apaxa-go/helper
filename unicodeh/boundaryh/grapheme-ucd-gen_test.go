//replacer:generated-file

package boundaryh

import (
	"reflect"
	"testing"
)

func TestGraphemeClustersInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := GraphemeClustersInString(test.SampleInString())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestLastGraphemeClusterInString(t *testing.T) {
	var stat Stat

	// Same as function "GraphemeClustersInString", but going from end to begin.
	revGCs := func(s string) (boundaries []Boundary) {
		for len(s) > 0 {
			pos := LastGraphemeClusterInString(s)
			boundaries = append([]Boundary{{pos, len(s)}}, boundaries...)
			s = s[:pos]
		}
		return
	}

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := revGCs(test.SampleInString())
		if !reflect.DeepEqual(expBoundaries, boundaries) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, expBoundaries, boundaries)
		}
	}

	stat.Log(t)
}

func TestGraphemeClusterAtInString(t *testing.T) {
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
		expBoundaries, skip := test.breaksToBoundariesInString()
		if skip {
			l := len(test.SampleInString())
			stat.Add(l)
			stat.Skip(l)
			continue
		}
		for pos := 0; pos < len(test.SampleInString()); pos++ {
			stat.Add()
			b := GraphemeClusterAtInString(test.SampleInString(), pos)
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

func TestGraphemeClusterBreaksInString(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBreaks, skip := test.BreaksInString()
		if skip {
			stat.Skip()
			continue
		}
		breaks := GraphemeClusterBreaksInString(test.SampleInString())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

func TestGraphemeClusters(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := GraphemeClusters(test.Sample())
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
	revGCs := func(bytes []byte) (boundaries []Boundary) {
		for len(bytes) > 0 {
			pos := LastGraphemeCluster(bytes)
			boundaries = append([]Boundary{{pos, len(bytes)}}, boundaries...)
			bytes = bytes[:pos]
		}
		return
	}

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			stat.Skip()
			continue
		}
		boundaries := revGCs(test.Sample())
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
		expBoundaries, skip := test.breaksToBoundaries()
		if skip {
			l := len(test.Sample())
			stat.Add(l)
			stat.Skip(l)
			continue
		}
		for pos := 0; pos < len(test.Sample()); pos++ {
			stat.Add()
			b := GraphemeClusterAt(test.Sample(), pos)
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

func TestGraphemeClusterBreaks(t *testing.T) {
	var stat Stat

	for testI, test := range ucdGraphemeClusterTests {
		stat.Add()
		expBreaks, skip := test.Breaks()
		if skip {
			stat.Skip()
			continue
		}
		breaks := GraphemeClusterBreaks(test.Sample())
		if !reflect.DeepEqual(breaks, expBreaks) {
			stat.Fail()
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.breaks, breaks)
		}
	}

	stat.Log(t)
}

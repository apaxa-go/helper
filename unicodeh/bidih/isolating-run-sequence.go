package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

type runSequence struct {
	sos, eos      bidi.Class
	ranges        []struct{ from, to int } // [from; to)
	rangeI, runeI int                      // for iteration over sequence
}

func makeRunSequence(from, to int) runSequence {
	return runSequence{ranges: []struct{ from, to int }{{from, to}}}
}

func (s *runSequence) start() {
	s.rangeI = 0
	s.runeI = s.ranges[0].from
}
func (s *runSequence) end() {
	s.rangeI = len(s.ranges) - 1
	s.runeI = s.ranges[s.rangeI].to - 1
}
func (s *runSequence) next() {
	if s.isEnd() {
		return
	}
	s.runeI++
	if s.runeI >= s.ranges[s.rangeI].to && s.rangeI < len(s.ranges)-1 {
		s.rangeI++
		s.runeI = s.ranges[s.rangeI].from
	}
}
func (s *runSequence) prev() {
	if s.isStart() {
		return
	}
	s.runeI--
	if s.runeI < s.ranges[s.rangeI].from && s.rangeI > 0 {
		s.rangeI--
		s.runeI = s.ranges[s.rangeI].to - 1
	}
}
func (s *runSequence) isEnd() bool {
	return s.runeI >= s.ranges[len(s.ranges)-1].to
}
func (s *runSequence) isStart() bool {
	return s.runeI < s.ranges[0].from
}

type sequenceState struct {
	rangeI, runeI int
}

func (s *runSequence) saveState() sequenceState {
	return sequenceState{s.rangeI, s.runeI}
}
func (s *runSequence) loadState(state sequenceState) {
	s.rangeI = state.rangeI
	s.runeI = state.runeI
}

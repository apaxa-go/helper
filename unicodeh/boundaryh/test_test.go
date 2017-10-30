package boundaryh

// Common data, types & so on for tests.

import (
	"testing"
	"unicode/utf8"
)

//go:generate go run ./internal/test-generator/main.go ../internal/ucd-data

func breaksToBoundaries(breaks []int) (boundaries []Boundary) {
	if len(breaks) == 0 {
		return
	}
	boundaries = make([]Boundary, len(breaks)-1)
	for boundaryI := range boundaries {
		boundaries[boundaryI] = Boundary{breaks[boundaryI], breaks[boundaryI+1]}
	}
	return
}

type ucdTest struct {
	runes  []rune
	breaks []int
}

func (t ucdTest) SampleInRunes() []rune  { return t.runes }
func (t ucdTest) SampleInString() string { return string(t.runes) }

// ...InBytes
func (t ucdTest) Sample() []byte                           { return []byte(string(t.runes)) }
func (t ucdTest) BreaksInRunes() (breaks []int, skip bool) { return t.breaks, false }
func (t ucdTest) BreaksInString() (breaks []int, skip bool) {
	breaks = make([]int, len(t.breaks))
	pos := 0
	runePos := 0
	for breakI, breakV := range t.breaks {
		for ; runePos < breakV; runePos++ {
			l := utf8.RuneLen(t.runes[runePos])
			if l == -1 {
				return nil, true
			}
			pos += l
		}
		breaks[breakI] = pos
	}
	return
}

// ...InBytes
func (t ucdTest) Breaks() (breaks []int, skip bool) { return t.BreaksInString() }
func (t *ucdTest) breaksToBoundariesInRunes() (boundaries []Boundary, skip bool) {
	return breaksToBoundaries(t.breaks), false
}
func (t *ucdTest) breaksToBoundariesInString() (boundaries []Boundary, skip bool) {
	breaks, skip := t.BreaksInString()
	boundaries = breaksToBoundaries(breaks)
	return
}

// ...InBytes
func (t *ucdTest) breaksToBoundaries() (boundaries []Boundary, skip bool) {
	return t.breaksToBoundariesInString()
}

type Stat struct {
	total, skipped, failed int
}

func (s *Stat) Add(counts ...int) {
	if len(counts) == 0 {
		s.total++
		return
	}
	for _, count := range counts {
		s.total += count
	}
}
func (s *Stat) Fail() { s.failed++ }
func (s *Stat) Skip(counts ...int) {
	if len(counts) == 0 {
		s.skipped++
		return
	}
	for _, count := range counts {
		s.skipped += count
	}
}
func (s *Stat) Passed() int { return s.total - s.skipped - s.failed }
func (s *Stat) Log(t *testing.T) {
	if s.total == 0 {
		t.Error("Internal test error - 0 total tests")
	}
	if s.skipped == s.total {
		t.Error("Internal test error - all tests are skipped")
	}
	t.Logf("\n Total: %v\nSkipped: %v\nPassed: %v\nFailed: %v\n", s.total, s.skipped, s.Passed(), s.failed)
}

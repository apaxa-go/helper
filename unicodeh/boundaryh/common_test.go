package boundaryh

import "testing"

//go:generate go run ./internal/test-generator/main.go ../internal/ucd-data

type ucdTest struct {
	runes  []rune
	breaks []int
}

func breaksToBoundaries(breaks []int) (boundaries []Boundary) {
	boundaries = make([]Boundary, len(breaks)-1)
	for boundaryI := range boundaries {
		boundaries[boundaryI] = Boundary{breaks[boundaryI], breaks[boundaryI+1]}
	}
	return
}

type Stat struct {
	total, failed int
}

func (s *Stat) Add()  { s.total++ }
func (s *Stat) Fail() { s.failed++ }
func (s *Stat) Log(t *testing.T) {
	if s.total == 0 {
		t.Error("Internal test error - 0 total tests")
	}
	t.Logf("\n Total: %v\nPassed: %v\nFailed: %v\n", s.total, s.total-s.failed, s.failed)
}

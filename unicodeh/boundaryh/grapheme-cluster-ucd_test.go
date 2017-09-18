package boundaryh

//go:generate go run ./internal/test-generator/main.go ../internal/ucd-data

import (
	"testing"
	"reflect"
)

type ucdTest struct {
	runes []rune
	boundaries []Boundary
}

func TestBoundariesRunes(t *testing.T) {
	for testI, test := range ucdTests {
		boundaries := GraphemeClusters(test.runes)
		if !reflect.DeepEqual(test.boundaries, boundaries) {
			t.Errorf("%v \"%v\": expect %v, got %v", testI, (test.runes), test.boundaries, boundaries)
		}
	}

	t.Log(1536, getGCClass(1536))
	t.Log(1424, getGCClass(1424))
}
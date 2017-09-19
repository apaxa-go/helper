package boundaryh

import (
	"reflect"
	"testing"
)

type ucdLineTest ucdTest

func TestLines(t *testing.T) {
	for testI, test := range ucdLineTests {
		boundaries := Lines(test.runes)
		if !reflect.DeepEqual(test.boundaries, boundaries) {
			t.Errorf("%v \"%v\": expect %v, got %v", testI, test.runes, test.boundaries, boundaries)
		}
	}
}

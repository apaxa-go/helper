package boundaryh

import (
	"reflect"
	"testing"
)

type ucdWordTest ucdTest

func TestWords(t *testing.T) {
	for testI, test := range ucdWordTests{
		boundaries := Words(test.runes)
		if !reflect.DeepEqual(test.boundaries, boundaries) {
			t.Errorf("%v \"%v\": expect %v, got %v", testI, (test.runes), test.boundaries, boundaries)
		}
	}
}

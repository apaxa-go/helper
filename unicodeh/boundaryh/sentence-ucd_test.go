package boundaryh

import (
	"reflect"
	"testing"
)

type ucdSentenceTest ucdTest

func TestSentences(t *testing.T) {
	for testI, test := range ucdSentenceTests{
		boundaries := Sentences(test.runes)
		if !reflect.DeepEqual(test.boundaries, boundaries) {
			t.Errorf("%v \"%v\": expect %v, got %v", testI, (test.runes), test.boundaries, boundaries)
		}
	}
}

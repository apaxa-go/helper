//replacer:generated-file
package mathh

import "testing"


func TestRound32(t *testing.T) {
	test := []struct {
		f float32
		i int64
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, -1},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, -1},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},
	}
	for j, v := range test {
		i := Round32(v.f)
		if i != v.i {
			t.Errorf("\nTestRound32 - %v.\nExpected int64: %v\ngot: %v", j, v.i, i)
		}
	}
}

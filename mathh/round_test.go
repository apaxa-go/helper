package mathh

import "testing"

func TestRound(t *testing.T) {
	type testElement struct {
		f float64
		i int64
	}

	test := []testElement{
		// 0
		{
			0.5,
			1,
		},

		// 1
		{
			0,
			0,
		},

		// 2
		{
			-0.5,
			-1,
		},

		// 3
		{
			0.48999999,
			0,
		},

		// 4
		{
			-0.4899999,
			0,
		},

		// 5
		{
			0.599999,
			1,
		},

		// 6
		{
			-0.599999,
			-1,
		},

		// 7
		{
			99.999999999,
			100,
		},

		// 8
		{
			99.48999999999999999999999,
			99,
		},

		// 9
		{
			99.0000000001,
			99,
		},
	}
	for j, v := range test {
		i := Round(v.f)
		if i != v.i {
			t.Errorf("\nTestRound - %v.\nExpected int64: %v\ngot: %v", j, v.i, i)
		}
	}
}

package mathh

import "testing"

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old Round64	float64
//replacer:new Round32	float32

func TestRound64(t *testing.T) {
	test := []struct {
		f float64
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
		i := Round64(v.f)
		if i != v.i {
			t.Errorf("#%v expect %v, got: %v", j, v.i, i)
		}
	}
}

package templateh

import "errors"

func NewEmptySlice(n int) []struct{} { return make([]struct{}, n) }

// for (i=from; i<to; i+=step)
func NewForSlice(from, to, step int) (r []int) {
	r = make([]int, (to-from)/step)
	for i := range r {
		r[i] = from + step*i
	}
	return
}

func NewFromSlice(from, num int) []int {
	return NewForSlice(from, from+num, 1)
}

func Dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func Add(a, b int) int {
	return a + b
}

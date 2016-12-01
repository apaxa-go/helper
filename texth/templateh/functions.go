package templateh

import (
	"errors"
	"github.com/apaxa-go/helper/mathh"
	"github.com/apaxa-go/helper/strconvh"
)

// NewEmptySlice returns new empty slice of integers with length and capacity = n.
// n should be not less than 0.
func NewEmptySlice(n int) []struct{} { return make([]struct{}, n) }

// NewRange returns slice of integers initialized with specified arithmetic progression.
// Progression defined by from and step: r[0]=from; r[i+1]=r[i]+step.
// to defines length (and capacity) of slice: (to-from)/step.
// If length (as described above) is negative or step is zero then panic will be raised.
// for (i=from; i<to; i+=step) for positive arguments.
func NewRange(from, to, step int) (r []int) {
	if step == 0 ||
		(to > from && step < 0) ||
		(to < from && step > 0) {
		panic("NewRange: from, to and step are inconsistent: " + strconvh.FormatInt(from) + ", " + strconvh.FormatInt(to) + ", " + strconvh.FormatInt(step))
	}
	r = make([]int, mathh.DivideCeilInt(to-from, step))
	for i := range r {
		r[i] = from + step*i
	}
	return
}

// Dict groups passed values by pair in map[string]interface{}.
// It requires that length of values is even.
// Each odd value used as key and should be string. Following after it element (each even) is value for this key and may of any type.
// Dict("one", v1, "two", v2) => {"one": v1, "two": v2}
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

// Add just returns a+b
func Add(a, b int) int { return a + b }

// Sub just returns a-b
func Sub(a, b int) int { return a - b }

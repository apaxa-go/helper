package templateh

import (
	"reflect"
	"testing"
)

func TestNewEmptySlice(t *testing.T) {
	tests := []int{0, 1, 2, 10, 100}
	for _, v := range tests {
		if r := NewEmptySlice(v); len(r) != v || cap(r) != v {
			t.Errorf("expected length and capacity are %v, but got %v and %v", v, len(r), cap(r))
		}
	}
}

func TestNewRange(t *testing.T) {
	type testElement struct {
		f, t, s int
		r       []int
	}
	tests := []testElement{
		{0, 0, 10, []int{}},
		{0, 4, 1, []int{0, 1, 2, 3}},
		{0, 4, 2, []int{0, 2}},
		{1, 4, 1, []int{1, 2, 3}},
		{1, 4, 2, []int{1, 3}},
		{1, 4, 3, []int{1}},
		{10, -10, -3, []int{10, 7, 4, 1, -2, -5, -8}},
		{-1, -4, -1, []int{-1, -2, -3}},
		{-1, -4, -2, []int{-1, -3}},
		{-1, -4, -3, []int{-1}},
	}
	for _, v := range tests {
		if r := NewRange(v.f, v.t, v.s); !reflect.DeepEqual(r, v.r) {
			t.Errorf("expected slice %v, but got %v", v.r, r)
		}
	}
}

func TestNewRange2(t *testing.T) {
	defer func() { recover() }()
	NewRange(1, 10, -1)
	t.Error("panic expected, but no panic")
}

func TestDict(t *testing.T) {
	type testElement struct {
		v   []interface{}
		r   map[string]interface{}
		err bool
	}
	tests := []testElement{
		{[]interface{}{}, map[string]interface{}{}, false},
		{[]interface{}{1}, map[string]interface{}{}, true},
		{[]interface{}{1, 2}, map[string]interface{}{}, true},
		{[]interface{}{"1", 2}, map[string]interface{}{"1": 2}, false},
		{[]interface{}{"1", 2, 3}, map[string]interface{}{}, true},
		{[]interface{}{"1", 2, "3", "four"}, map[string]interface{}{"1": 2, "3": "four"}, false},
	}
	for _, v := range tests {
		r, err := Dict(v.v...)
		if (err != nil) != v.err {
			t.Errorf("expect error for %v - %v, but got %v", v.v, v.err, err)
		}
		if !v.err && err == nil && !reflect.DeepEqual(v.r, r) {
			t.Errorf("expect map %v, got %v", v.r, r)
		}
	}
}

var arithmeticTests = []int{-10, -6, -4, -3, -1, 0, 1, 2, 3, 4, 5, 6}

func TestAdd(t *testing.T) {
	for _, v1 := range arithmeticTests {
		for _, v2 := range arithmeticTests {
			if r := Add(v1, v2); r != v1+v2 {
				t.Errorf("expect %v, got %v", v1+v2, r)
			}
		}
	}
}

func TestSub(t *testing.T) {
	for _, v1 := range arithmeticTests {
		for _, v2 := range arithmeticTests {
			if r := Sub(v1, v2); r != v1-v2 {
				t.Errorf("expect %v, got %v", v1-v2, r)
			}
		}
	}
}

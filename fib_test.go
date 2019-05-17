package main

import (
	"testing"
	"reflect"
)

var fibTests = []struct {
	n        int
	expected []int
}{
	{1, []int{0}},
	{2, []int{0,1}},
	{3, []int{0,1,1}},
	{4, []int{0,1,1,2}},
	{5, []int{0,1,1,2,3}},
	{6, []int{0,1,1,2,3,5}},
	{7, []int{0,1,1,2,3,5,8}},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		result := Fib(tt.n, 0, 1, []int{})
		if reflect.DeepEqual(result, tt.expected) == false {
			t.Errorf("Fib(%d): expected %d, result %d", tt.n, tt.expected, result)
		}
	}
}

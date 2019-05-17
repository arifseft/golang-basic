package main

import (
	"testing"
	"reflect"
)

var primeTests = []struct {
	n        int
	expected []int
}{
	{1, []int{2}},
	{2, []int{2,3}},
	{3, []int{2,3,5}},
	{4, []int{2,3,5,7}},
	{5, []int{2,3,5,7,11}},
	{6, []int{2,3,5,7,11,13}},
	{7, []int{2,3,5,7,11,13,17}},
}

func TestPrime(t *testing.T) {
	for _, tt := range primeTests {
		result := Prime(tt.n)
		if reflect.DeepEqual(result, tt.expected) == false {
			t.Errorf("Prime(%d): expected %d, result %d", tt.n, tt.expected, result)
		}
	}
}

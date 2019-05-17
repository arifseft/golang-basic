package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		expected int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		result := Sum(table.x, table.y)
		if result != table.expected {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, result, table.expected)
		}
	}
}

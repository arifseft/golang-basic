package main

import "testing"

func TestMultiply(t *testing.T) {
	tables := []struct {
		x int
		y int
		expected int
	}{
		{1, 3, 3},
		{2, 4, 8},
		{3, 6, 18},
		{5, 2, 10},
	}

	for _, table := range tables {
		result := Multiply(table.x, table.y)
		if result != table.expected {
			t.Errorf("Multiply of (%d*%d) was incorrect, got: %d, want: %d.", table.x, table.y, result, table.expected)
		}
	}
}

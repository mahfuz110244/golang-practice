package main

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"Positive Numbers", 5, 3, 8},
		{"Negative Numbers", -2, -4, -6},
		{"Mixed Numbers", 7, -3, 4},
		{"Zero", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.x, tc.y)
			if result != tc.expected {
				t.Errorf("TestAdd %s failed: expected: %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}

package puzzles

import "testing"

type smallestMissingPositiveNumberTestCase struct {
	input    []int
	expected int
}

func TestSmallestMissingPositiveNumber(t *testing.T) {
	testCases := []smallestMissingPositiveNumberTestCase{
		{
			input:    []int{3, 4, -1, -1},
			expected: 2,
		},
		{
			input:    []int{1, 2, 0},
			expected: 3,
		},
		{
			input:    []int{-2, -1, 1, 2, 4, 5, 8, 10},
			expected: 3,
		},
		{
			input:    []int{-5, -3, 6, 7, 9},
			expected: 5,
		},
		{
			input:    []int{4, 5, 10, 12, 43},
			expected: 3,
		},
	}
	for _, c := range testCases {
		if got := smallestMissingPositiveNumber(c.input); got != c.expected {
			t.Errorf("expected : %v, got : %v, for case : %v", c.expected, got, c)
		}
	}
}

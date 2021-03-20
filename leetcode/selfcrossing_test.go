package puzzles

import "testing"

type isSelfCorssingTestCase struct {
	input    []int
	expected bool
}

func TestIsSelfCrossing(t *testing.T) {
	testCases := []isSelfCorssingTestCase{
		{
			input:    []int{2, 1, 1, 2},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 1},
			expected: true,
		},
	}
	for _, c := range testCases {
		got := IsSelfCrossing(c.input)
		if got != c.expected {
			t.Fatalf("expected : %v, got: %v\n", c.expected, got)
		}
	}
}

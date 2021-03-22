package puzzles

import "testing"

type canSumToKTestCase struct {
	inputNumbers []int
	k            int
	expected     bool
}

func TestCanSumToK(t *testing.T) {
	testcases := []canSumToKTestCase{
		{
			inputNumbers: []int{10, 15, 3, 7},
			k:            17,
			expected:     true,
		},
		{
			inputNumbers: []int{5, 6, 11},
			k:            100,
			expected:     false,
		},
		{
			inputNumbers: []int{5, 3, 14, 433, 143, 12},
			k:            155,
			expected:     true,
		},
	}
	for _, c := range testcases {
		got := CanSumToK(c.inputNumbers, c.k)
		if c.expected != got {
			t.Fatalf("expected : %v got : %v\n", c.expected, got)
		}
	}
}

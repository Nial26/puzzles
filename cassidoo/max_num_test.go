package puzzles

import "testing"

type maxNumTestCase struct {
	listA    []int
	listB    []int
	k        int
	expected int
}

func TestMaxNum(t *testing.T) {
	testCases := []maxNumTestCase{
		{
			listA:    []int{3, 4, 6, 5},
			listB:    []int{9, 0, 2, 5, 8, 3},
			k:        5,
			expected: 98655,
		},
		{
			listA:    []int{5, 5, 3, 4, 6},
			listB:    []int{1, 1, 7, 3},
			k:        3,
			expected: 765,
		},
	}
	for _, c := range testCases {
		got := maxNum(c.listA, c.listB, c.k)
		if got != c.expected {
			t.Fatalf("expected: %v, got : %v, for input : %v", c.expected, got, c)
		}
	}
}

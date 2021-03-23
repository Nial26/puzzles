package puzzles

import (
	"reflect"
	"testing"
)

type multipliedValueTestCase struct {
	input    []int
	expected []int
}

func TestMultipledValueWithoutDivision(t *testing.T) {
	testCases := []multipliedValueTestCase{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{120, 60, 40, 30, 24},
		},
		{
			input:    []int{3, 2, 1},
			expected: []int{2, 3, 6},
		},
		{
			input:    []int{1, 2, 4, 50, 0},
			expected: []int{0, 0, 0, 0, 400},
		},
	}
	for _, c := range testCases {
		got := multipledValueWithoutDivision(c.input)
		if !reflect.DeepEqual(c.expected, got) {
			t.Fatalf("expected : %v, got : %v, for input : %v", c.expected, got, c)
		}
	}
}

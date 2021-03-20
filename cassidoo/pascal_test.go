package puzzles

import (
	"reflect"
	"testing"
)

type pascalTestCase struct {
	input    int
	expected []int
}

func TestPascal(t *testing.T) {
	cases := []pascalTestCase{
		{
			input:    0,
			expected: []int{1},
		},
		{
			input:    1,
			expected: []int{1, 1},
		},
		{
			input:    5,
			expected: []int{1, 5, 10, 10, 5, 1},
		},
	}
	for _, c := range cases {
		got := Pascal(c.input)
		if !reflect.DeepEqual(c.expected, got) {
			t.Fatalf("expected %v, got %v\n", c.expected, got)
		}
	}
}

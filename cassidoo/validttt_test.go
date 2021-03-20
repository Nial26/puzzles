package puzzles

import "testing"

type validTTTPositionTestCase struct {
	input    []string
	expected bool
}

func TestValidTTTPosition(t *testing.T) {
	testCases := []validTTTPositionTestCase{
		{
			input:    []string{"XXX", "OOO", "XXX"},
			expected: false,
		},
		{
			input:    []string{"XOX", " X ", "   "},
			expected: false,
		},
		{
			input:    []string{"XOX", "O O", "XOX"},
			expected: true,
		},
		{
			input:    []string{"OOO", "   ", "XXX"},
			expected: false,
		},
	}
	for _, c := range testCases {
		got := ValidTTTPosition(c.input)
		if got != c.expected {
			t.Fatalf("expected %v, got %v\n", c.expected, got)
		}
	}
}

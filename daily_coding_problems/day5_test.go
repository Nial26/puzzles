package puzzles

import "testing"

type day5TestCase struct {
	a         int
	b         int
	carOutput int
	cdrOutput int
}

func TestDay5(t *testing.T) {
	testCases := []day5TestCase{
		{
			a:         3,
			b:         4,
			carOutput: 3,
			cdrOutput: 4,
		},
		{
			a:         0,
			b:         -1,
			carOutput: 0,
			cdrOutput: -1,
		},
	}
	for _, c := range testCases {
		pair := cons(c.a, c.b)
		if carGot, cdrGot := car(pair), cdr(pair); carGot != c.carOutput && cdrGot != c.cdrOutput {
			t.Fatalf("expected : %v, got : %v, for case : %v", []int{c.carOutput, c.cdrOutput}, []int{carGot, cdrGot}, c)
		}
	}
}

package puzzles

import (
	"math/rand"
	"testing"
	"time"
)

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
		got := canSumToK(c.inputNumbers, c.k)
		if c.expected != got {
			t.Fatalf("expected : %v got : %v for input : %v\n", c.expected, got, c.inputNumbers)
		}
	}
	for _, c := range testcases {
		got := optimisedCanSumToK(c.inputNumbers, c.k)
		if c.expected != got {
			t.Fatalf("expected : %v got : %v for input : %v\n", c.expected, got, c.inputNumbers)
		}
	}
}

var canBenchToK bool

func BenchmarkCanSumToK(b *testing.B) {
	var result bool
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		numbers := rand.Perm(10)
		result = canSumToK(numbers, numbers[5]+numbers[7])
	}
	canBenchToK = result
}

var optimisedCanBenchToK bool

func BenchmarkOptimisedCanSumToK(b *testing.B) {
	var result bool
	rand.Seed(time.Now().Unix())
	for n := 0; n < b.N; n++ {
		numbers := rand.Perm(10)
		result = optimisedCanSumToK(numbers, numbers[5]+numbers[7])
	}
	optimisedCanBenchToK = result
}

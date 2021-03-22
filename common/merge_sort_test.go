package puzzles

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type mergeSortTestCase struct {
	input    []int
	expected []int
}

func TestMergeSort(t *testing.T) {
	testCases := []mergeSortTestCase{
		{
			input:    []int{4, 123, 43, 16, 76},
			expected: []int{4, 16, 43, 76, 123},
		},
		{
			input:    []int{7},
			expected: []int{7},
		},
		{
			input:    []int{},
			expected: []int{},
		},
	}

	for _, c := range testCases {
		got := mergeSort(c.input)
		if !reflect.DeepEqual(c.expected, got) {
			fmt.Println(c.expected)
			fmt.Println(got)
			t.Fatalf("expected : %v, got : %v, for input : %v", c.expected, got, c.input)
		}
	}
}

func TestMergeSortOnRandomInput(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var expected []int
	for i := range [100]int{} {
		expected = append(expected, i)
	}
	input := rand.Perm(100)
	got := mergeSort(input)
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected : %v, got : %v, for input : %v", expected, got, input)
	}
}

var benchMergeResult []int

func BenchmarkMergeSort(b *testing.B) {
	var out []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		rand.Perm(100)
		out = mergeSort(rand.Perm(100))
	}
	benchMergeResult = out
}

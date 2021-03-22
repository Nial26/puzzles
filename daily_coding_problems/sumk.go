package puzzles

// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?

func canSumToK(numbers []int, k int) bool {
	for i, num1 := range numbers {
		for _, num2 := range numbers[i+1:] {
			if num1+num2 == k {
				return true
			}
		}
	}
	return false
}

// 10, 15, 3, 7
// 7, 2, 14, 10
// For one pass I was thinking we could the diff to a set and take an intersection, but the add into a set is usually O(N) (Naive Implementation)
// But looks like that is done here https://www.geeksforgeeks.org/given-an-array-a-and-a-number-x-check-for-pair-in-a-with-sum-as-x/, but instead
// of taking an intersection they directly check a diff

func optimisedCanSumToK(numbers []int, k int) bool {
	set := make(map[int]bool, 2)
	for _, num := range numbers {
		if set[num] {
			return true
		}
		set[k-num] = true
	}
	return false
}

// Funnily enough the optimised one pass calls takes more time in benchmark
// So I must be doing something wrong 🤔

// goos: darwin
// goarch: amd64
// pkg: github.com/Nial26/puzzles/daily_coding_problems
// cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
// BenchmarkCanSumToK-8            	 5166394	       224.0 ns/op	      80 B/op	       1 allocs/op
// BenchmarkOptimisedCanSumToK-8   	 3799670	       307.6 ns/op	      80 B/op	       1 allocs/op

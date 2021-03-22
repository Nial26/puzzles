package puzzles

import (
	"math"
	"sort"
)

// You’re given two integer arrays (n and m), and an integer k. Using the digits from n and m, return the largest number you can of length k.
// Example:
// n = [3,4,6,5]
// m = [9,0,2,5,8,3]
// k = 5
// $ maxNum(n, m, k)
// $ 98653

func maxNum(l1, l2 []int, k int) int {
	mergedList := append(l1, l2...)
	sort.Sort(sort.Reverse(sort.IntSlice(mergedList)))
	return buildNumber(mergedList[:k])

}

func buildNumber(n []int) int {
	number := 0
	power := len(n) - 1
	for _, v := range n {
		number += v * int(math.Pow(10, float64(power)))
		power--
	}
	return number
}

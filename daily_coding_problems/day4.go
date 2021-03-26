package puzzles

import "fmt"

/*
This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

[-2, -1, 1, 2, 4, 5, 8, 10] -> 3
[1, 5, 9, 10, -1, 2, 4, 8]

[4, 5, 10, 54, 12]
You can modify the input array in-place.
*/

const INF = 10000

func smallestMissingPositiveNumber(numbers []int) int {
	smallestNumber, secondSmallestNumber := INF, INF
	hasZero := false
	for _, v := range numbers {
		if v > 0 && v < smallestNumber {
			secondSmallestNumber = smallestNumber
			smallestNumber = v
		} else if v > 0 && v < secondSmallestNumber {
			secondSmallestNumber = v
		}
		if v == 0 {
			hasZero = true
		}
	}
	fmt.Printf("%d %d %v\n", smallestNumber, secondSmallestNumber, hasZero)
	if secondSmallestNumber == smallestNumber+1 {
		// Either We go towards zero if we didn't have a zero
		if !hasZero && smallestNumber-1 > 0 {
			return smallestNumber - 1
		}
		// Otherwise we go towards +INF
		return secondSmallestNumber + 1
	}
	return smallestNumber + 1
}

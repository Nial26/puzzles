package puzzles

import "fmt"

// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?
func CanSumToK(numbers []int, k int) bool {
	for i, num1 := range numbers {
		for _, num2 := range numbers[i+1:] {
			if num1+num2 == k {
				fmt.Printf("%d + %d\n", num1, num2)
				return true
			}
		}
	}
	return false
}

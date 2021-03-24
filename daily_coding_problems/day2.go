package puzzles

/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

func multipledValueWithoutDivision(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}
	result := make([]int, len(numbers))
	for i := range numbers {
		result[i] = multiply(numbers[:i]) * multiply(numbers[i+1:])
	}
	return result
}

func multiply(numbers []int) int {
	total := 1
	for _, v := range numbers {
		total *= v
	}
	return total
}

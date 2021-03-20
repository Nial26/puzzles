package puzzles

func Pascal(i int) []int {
	if i == 0 {
		return []int{1}
	} else {
		previousRow := Pascal(i - 1)
		// We'll get something like [1 3 3 1], we want to convert it to [0 1 3 3 1 0]
		previousRow = append([]int{0}, previousRow...)
		previousRow = append(previousRow, 0)
		return pairwiseSum(previousRow)
	}
}

func pairwiseSum(row []int) []int {
	result := make([]int, len(row)-1)
	for i := range row {
		if i <= len(row)-2 {
			result[i] = row[i] + row[i+1]
		}
	}
	return result
}

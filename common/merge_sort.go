package common

func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	return merge(MergeSort(numbers[:len(numbers)/2]), MergeSort(numbers[len(numbers)/2:]))

}

func merge(sortedListA, sortedListB []int) []int {
	var i, j int
	var mergedList []int
	for i < len(sortedListA) && j < len(sortedListB) {
		if sortedListA[i] < sortedListB[j] {
			mergedList = append(mergedList, sortedListA[i])
			i++
		} else {
			mergedList = append(mergedList, sortedListB[j])
			j++
		}
	}
	for i < len(sortedListA) {
		mergedList = append(mergedList, sortedListA[i])
		i++
	}
	for j < len(sortedListB) {
		mergedList = append(mergedList, sortedListB[j])
		j++
	}
	return mergedList
}

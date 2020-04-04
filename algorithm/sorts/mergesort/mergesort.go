package mergesort

func MergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	left, right := divide(slice)
	leftSlice := MergeSort(left)
	rightSlice := MergeSort(right)

	sortedSlice := merge(leftSlice, rightSlice)

	return sortedSlice
}

func merge(leftSlice, rightSlice []int) []int {
	sortedSliceLength := len(leftSlice) + len(rightSlice)
	sortedSlice := make([]int, sortedSliceLength)

	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(leftSlice) || rightIndex < len(rightSlice) {
		if leftIndex >= len(leftSlice) ||
			rightIndex < len(rightSlice) && (leftSlice[leftIndex] > rightSlice[rightIndex]) {

			sortedSlice[leftIndex+rightIndex] = rightSlice[rightIndex]
			rightIndex++
		} else {
			sortedSlice[leftIndex+rightIndex] = leftSlice[leftIndex]
			leftIndex++
		}
	}

	return sortedSlice
}

func divide(slice []int) ([]int, []int) {
	sliceLength := len(slice)
	dividePointIndex := sliceLength / 2
	return slice[:dividePointIndex], slice[dividePointIndex:]
}

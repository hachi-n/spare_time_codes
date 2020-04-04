package quicksort

func QuickSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	pivotIndex := len(slice) / 2

	pivot, popedSlice := popPivot(pivotIndex, slice)
	leftSlice, rightSlice := divide(popedSlice, pivot)

	sortedLeftSlice := QuickSort(leftSlice)
	sortedRightSlice := QuickSort(rightSlice)

	var sortedSlice []int
	sortedSlice = append(sortedLeftSlice, pivot)
	sortedSlice = append(sortedSlice, sortedRightSlice...)

	return sortedSlice
}

func divide(slice []int, pivot int) ([]int, []int) {
	var leftSlice, rightSlice []int

	for _, e := range slice {
		if e < pivot {
			leftSlice = append(leftSlice, e)
		} else {
			rightSlice = append(rightSlice, e)
		}
	}

	return leftSlice, rightSlice
}

func popPivot(pivotIndex int, slice []int) (int, []int) {
	var popedSlice []int
	popedSlice = append(popedSlice, slice[:pivotIndex]...)
	popedSlice = append(popedSlice, slice[pivotIndex+1:]...)

	return slice[pivotIndex], popedSlice
}

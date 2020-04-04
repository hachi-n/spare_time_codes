package selectionsort

import "math"

func SelectionSort(slice []int) []int {

	for index := range slice {
		selectionSort(slice, index)
	}

	return slice
}

func selectionSort(slice []int, index int) {
	min, minIndex := math.MaxInt32, index
	for j := index; j < len(slice); j++ {
		if min > slice[j] {
			min, minIndex = slice[j], j
		}
	}
	slice[index], slice[minIndex] = min, slice[index]
}

package insertionsort

func InsertionSort(slice []int) []int {
	for index := range slice {
		// not insertion
		if index == 0 {
			continue
		}

		insertionSort(slice, index)

	}
	return slice
}

func insertionSort(slice []int, index int) {
	for j := 0; j < index; j++ {

		if slice[index] < slice[j] {
			tmp := slice[index]

			for index > j {
				slice[index] = slice[index-1]
				index--
			}
			slice[j] = tmp
		}
	}
}

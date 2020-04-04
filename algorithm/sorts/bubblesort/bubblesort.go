package bubblesort

func BubbleSort(slice []int) []int {
	for range slice {
		slice = bubble(slice)
	}
	return slice
}

func bubble(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			tmp := slice[i+1]
			slice[i+1] = slice[i]
			slice[i] = tmp
		}
	}
	return slice
}

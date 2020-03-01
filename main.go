package main

import (
	"fmt"
	"github.com/hachi-n/awesomeProject/sorts/mergesort"
	"github.com/hachi-n/awesomeProject/sorts/quicksort"
	"time"
)

func main() {
	slice := []int{6, 3, 0, 5, 5, 1, 8, 2, 9, 4, 7}

	t1 := time.Now()
	mergeSortedSlice := mergesort.MergeSort(slice)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1), mergeSortedSlice)

	t1 = time.Now()
	quickSortedSlice := quicksort.QuickSort(slice)
	t2 = time.Now()

	fmt.Println(t2.Sub(t1), quickSortedSlice)
}

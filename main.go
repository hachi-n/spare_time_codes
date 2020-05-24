package main

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/spare_time_codes/interfaces/json_marshaler"
	"os"
)

func main() {
	//nodesMap := map[string]map[string]int{
	//	"A": {"B": 8, "C": 6},
	//	"B": {"A": 8, "C": 5, "D": 8, "E": 9},
	//	"C": {"A": 6, "B": 5, "D": 2, "E": 7},
	//	"D": {"B": 8, "C": 2, "E": 1, "F": 7},
	//	"E": {"B": 9, "C": 7, "D": 1, "F": 9},
	//	"F": {"D": 7, "E": 9},
	//}
	//
	//pathInfo := shortestpath.NewPathInfo(nodesMap)
	//shortestPath := pathInfo.FindShortestPath("A", "F")
	//fmt.Println(shorktestPath)

	//originalSlice := []int{54, 34, 54, 41, 54, 5, 45, 46, 547, 68, 578, 78, 9, 76, 54, 7, 8, 6, 56, 4564, 5654, 6, 525, 435, 43, 543, 5, 435, 543, 543, 3256, 6, 6, 236}
	//sortedSlice := insertionsort.InsertionSort(originalSlice)
	//// sortedSlice := selectionsort.SelectionSort(originalSlice)
	//fmt.Println(sortedSlice)

	jsonByte := []byte(`{"name": "mike", "age": 16, "nicknames": ["a","b","c"]}`)
	obj := new(json_marshaler.Person)
	if err := json.Unmarshal(jsonByte, obj); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(obj)

	data, _ := json.Marshal(obj)
	fmt.Println(string(data))

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var m = []string{
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
	}

	fmt.Println(memoryAllocationJoin(m))
	fmt.Println(memoryAllocationJoin2(m))
	fmt.Println(memoryAllocationJoin3(m))

}

func memoryAllocationJoin(str []string) string {
	return strings.Join(str, "")
}

func memoryAllocationJoin2(str []string) string {
	generatedString := ""
	for i := 0; i < len(str); i++ {
		generatedString += str[i]
	}
	return generatedString
}

func memoryAllocationJoin3(str []string) string {
	b := make([]byte, 63)
	index := 0
	for i := 0; i < len(str); i++ {
		copy(b[index:], str[i])
		index += len(str[i])
	}
	return string(b)
}

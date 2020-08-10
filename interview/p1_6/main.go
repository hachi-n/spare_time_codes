package p1_6

import (
	"strconv"
)

func OriginalCompress(str string) string {
	compressedStr := ""
	var testRune rune
	var currentRuneCount int

	for _, v := range str {
		if testRune == 0 {
			testRune = v
			currentRuneCount++
			continue
		}

		if testRune == v {
			currentRuneCount++
			continue
		}

		compressedStr += string(testRune) + strconv.Itoa(currentRuneCount)
		testRune = v
		currentRuneCount = 1
	}
	compressedStr += string(testRune) + strconv.Itoa(currentRuneCount)

	if len(compressedStr) >= len(str) {
		return str
	}
	return compressedStr
}

func OriginalCompress2(str string) string {
	compressedStr := ""
	var currentRuneCount int

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		currentRuneCount++
		if i == len(runes)-1 || runes[i] != runes[i+1] {
			compressedStr += string(runes[i]) + strconv.Itoa(currentRuneCount)
			currentRuneCount = 0
		}
	}

	if len(compressedStr) >= len(str) {
		return str
	}
	return compressedStr
}

func OriginalCompress3(str string) string {
	compressedL := compressedLength(str)

	if len([]rune(str)) <= compressedL {
		return str
	}
	compressedRunes := make([]rune, compressedL)
	compressedRunesIndex := 0

	var currentRuneCount int
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		currentRuneCount++
		if i == len(runes)-1 || runes[i] != runes[i+1] {
			compressedRunes[compressedRunesIndex], compressedRunes[compressedRunesIndex+1] =
				runes[i], []rune(strconv.Itoa(currentRuneCount))[0]

			compressedRunesIndex += 2
			currentRuneCount = 0
		}
	}

	if len(compressedRunes) >= len([]rune(str)) {
		return str
	}
	return string(compressedRunes)
}

func compressedLength(str string) int {
	var counter int

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if i == len(runes)-1 || runes[i] != runes[i+1] {
			counter++
		}
	}
	return counter * 2
}

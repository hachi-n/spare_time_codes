package p1_4

import "strings"

func IsPermutation(str string) bool {
	lowerRune := []rune(strings.ToLower(str))

	excludeRunes := []rune{' ', '.', '!'}
	checkerMap := make(map[rune]int)
	for _, v := range lowerRune {
		if containsRune(excludeRunes, v) {
			continue
		}
		if n, ok := checkerMap[v]; ok && n != 0 {
			checkerMap[v]--
			continue
		}
		checkerMap[v]++
	}

	var checkOddCount int
	for _, v := range checkerMap {
		if v&1 == 1 {
			checkOddCount++
		}
	}
	return checkOddCount <= 1
}

func containsRune(runes []rune, r rune) bool {
	for _, v := range runes {
		if v == r {
			return true
		}
	}
	return false
}

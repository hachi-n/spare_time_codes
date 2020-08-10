package p1_2

import "sort"

func IsSortedCharacters(str1, str2 string) bool {
	if str1 == str2 || len(str1) != len(str2) {
		return false
	}

	castedStr1, castedStr2 := []uint8(str1), []uint8(str2)

	sortFunc := func(s []uint8) []uint8 {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		return s
	}

	if string(sortFunc(castedStr1)) == string(sortFunc(castedStr2)) {
		return true
	}
	return false
}

func IsSortedCharacters2(str1, str2 string) bool {
	if str1 == str2 || len(str1) != len(str2) {
		return false
	}

	charactorsMap := make(map[uint8]int)
	for _, v := range []uint8(str1) {
		charactorsMap[v]++
	}

	for _, v := range []uint8(str2) {
		charactorsMap[v]--
		if charactorsMap[v] < 0 {
			return false
		}
	}

	for _, v := range charactorsMap {
		if v != 0 {
			return false
		}
	}
	return true
}

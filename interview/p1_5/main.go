package p1_5

import "math"

func IsEditedString(baseStr, editedStr string) bool {
	if math.Abs(float64(len([]rune(baseStr))-len([]rune(editedStr)))) > 1 {
		return false
	}

	checkerMap := make(map[rune]int)
	for _, v := range baseStr {
		checkerMap[v]++
	}

	for _, v := range editedStr {
		checkerMap[v]--
	}

	changeUpCount, changeDownCount := 0, 0
	for _, v := range checkerMap {
		if v > 0 {
			changeUpCount += v
		} else if v < 0 {
			changeDownCount += v
		}
	}

	if changeUpCount < 2 && changeDownCount < 2 {
		return true
	}
	return false
}

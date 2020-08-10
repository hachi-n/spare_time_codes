package p1_3

import (
	"strings"
)

func ConvertSpace(str string, realLength int) string {
	sentence := strings.TrimRight(str, " ")
	return strings.ReplaceAll(sentence, " ", "%20")
}

func ConvertSpace2(str string, realLength int) string {
	var blankCount int
	var realRealLength int

	castedStr := []rune(str)
	trimRightFlag := true

	for i := len(castedStr) - 1; i >= 0; i-- {
		if castedStr[i] == ' ' && trimRightFlag {
			continue
		}
		trimRightFlag = false
		if castedStr[i] == ' ' {
			blankCount++
		}
		realRealLength++
	}

	if blankCount == 0 {
		return str
	}

	rLength := realRealLength + blankCount*2
	r := make([]rune, rLength)
	rIndex := 0
	for _, v := range castedStr {
		if rIndex == rLength {
			break
		}
		if v == ' ' {
			r[rIndex] = '%'
			r[rIndex+1] = '2'
			r[rIndex+2] = '0'
			rIndex += 3
			continue
		}
		r[rIndex] = v
		rIndex++
	}

	return string(r)
}

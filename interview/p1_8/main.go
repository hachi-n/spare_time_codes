package p1_8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertMatrix(matrix [][]int) [][]int {
	var zeroPoints []string

	for i, row := range matrix {
		for j, c := range row {
			if c == 0 {
				zeroPoints =  append(zeroPoints, matrixKeyToKey(j, i))
			}
		}
	}
	for _, v := range zeroPoints {
		j, i := keyToMatrixKey(v)
		l := len(matrix[i])
		matrix[i] = make([]int, l)

		for mi, mv := range matrix {
			if mi == i {
				continue
			}
			mv[j] = 0
		}
	}
	fmt.Println(matrix)
	return matrix
}

const (
	separator = "::"
)

func matrixKeyToKey(m, n int) string {
	return fmt.Sprintf("%d%s%d", m, separator, n)
}

func keyToMatrixKey(key string) (int, int) {
	keys := strings.Split(key, separator)
	i, err := strconv.Atoi(keys[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	j, err := strconv.Atoi(keys[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i, j
}

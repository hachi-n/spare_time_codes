package p1_7

import "fmt"

func Rotate(top, right, bottom, left []string) ([]string, []string, []string, []string) {
	// 0 1 2 0
	// 2     1
	// 1     2
	// 0 2 1 0
	temp := make([]string, len(top))
	fmt.Println(top)
	copy(temp, top)
	copy(top, left)
	copy(left, bottom)
	copy(bottom, right)
	copy(right, temp)
	return top, right, bottom, left
}

func Rotate1(top, right, bottom, left []string) ([]string, []string, []string, []string) {
	temp := make([]string, len(top))
	copySlice(temp, top)
	copySlice(top, left)
	copySlice(left, bottom)
	copySlice(bottom, right)
	copySlice(right, temp)
	return top, right, bottom, left

}

func copySlice(dst, src []string) {
	for i, v := range src {
		dst[i] = v
	}
}

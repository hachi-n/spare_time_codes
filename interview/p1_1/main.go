package p1_1

import (
	"math/bits"
	"sort"
)

func IsDuplicationString(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}

func IsDupulicationString2(s string) bool {
	if len(s) > 128 {
		return true
	}

	strMap := make(map[uint8]bool)

	for i := 0; i < len(s); i++ {
		if _, ok := strMap[s[i]]; ok {
			return true
		}
		strMap[s[i]] = true
	}
	return false
}

func IsDupulicationString3(s string) bool {
	if len(s) > 128 {
		return true
	}

	n := 2
	for i := 0; i < (1 << uint64(len(s))); i++ {
		if bits.OnesCount64(uint64(i)) != n {
			continue
		}
		var subset []uint8
		for j := uint(0); j < uint(len(s)); j++ {
			if (i>>j)&1 == 1 {
				subset = append(subset, s[j])
			}
		}
		if subset[0] == subset[1] {
			return true
		}
	}
	return false
}

func IsDupulicationString4(s string) bool {
	hoge := []uint8(s)
	sort.Slice(hoge, func(i, j int) bool {
		return hoge[i] < hoge[j]
	})

	for i := 1; i < len(hoge); i++ {
		if hoge[i-1] == hoge[i] {
			return true
		}
	}
	return false
}

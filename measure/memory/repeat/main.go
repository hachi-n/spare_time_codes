package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	r := memoryAllocationSample6("hogefuga")
	fmt.Println(r)
}

func memoryAllocationSample(str string) string {
	mergedStr := ""
	for i := 0; i < 1000; i++ {
		mergedStr += str
	}
	return mergedStr
}

func memoryAllocationSample2(str string) string {
	stringList := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		stringList[i] = str
	}
	return strings.Join(stringList, "")
}

func memoryAllocationSample3(str string) string {
	return strings.Repeat(str, 1000)
}

func memoryAllocationSample4(str string) string {
	n := len(str) * 1000
	buf := bytes.NewBufferString(str)

	for buf.Len() < n {
		if buf.Len() <= n/2 {
			buf.WriteString(buf.String())
		} else {
			buf.WriteString(buf.String()[:n-buf.Len()])
			break
		}
	}
	return buf.String()
}

// my original repeat func
func memoryAllocationSample5(str string) string {
	l := uint(len(str))
	n := l * 1000
	b := make([]byte, n)

	copy(b, str)
	counter := uint(0)
	for c := uint(1) * l; c < n; c = (1 << counter) * l {
		if counter <= n/2 {
			copy(b[c:], b[:c])
		} else {
			copy(b[c:], b[:n-c])
			break
		}
		counter++
	}
	return string(b)
}

// copy => append.
func memoryAllocationSample6(str string) string {
	l := uint(len(str))
	n := l * 1000
	b := make([]byte, 0, n)

	b = append(b, str...)
	counter := uint(0)
	for c := uint(1) * l; c < n; c = (1 << counter) * l {
		if counter <= n/2 {
			b = append(b, b[:c]...)
		} else {
			b = append(b, b[:n-c]...)
			break
		}
		counter++
	}
	return string(b)
}

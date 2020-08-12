package main

import "testing"

var m = []string{
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
}


func Benchmark_memoryAllocationJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationJoin(m)
	}
}

func Benchmark_memoryAllocationJoin2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationJoin2(m)
	}
}

func Benchmark_memoryAllocationJoin3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationJoin3(m)
	}
}


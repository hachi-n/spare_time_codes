package main

import (
	"testing"
)

func Benchmark_memoryAllocationSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationSample("fuga")
	}
}

func Benchmark_memoryAllocationSample2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationSample2("fuga")
	}
}

func Benchmark_memoryAllocationSample3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationSample3("fuga")
	}
}

func Benchmark_memoryAllocationSample4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationSample4("fuga")
	}
}

func Benchmark_memoryAllocationSample5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationSample5("fuga")
	}
}

func Benchmark_memoryAllocationSample6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocationSample6("fuga")
	}
}

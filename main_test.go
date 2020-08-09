package main

import (
	"spare_time_codes/stream"
	"testing"
)

//func BenchmarkFileCreate(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		stream.FileCreate()
//	}
//}

func BenchmarkFileCopy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stream.FileCopy("./tmp/input.txt")
	}
}

func BenchmarkFileCopy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stream.FileCopy("./tmp/input2.txt")
	}
}

func BenchmarkFileMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stream.FileMerge()
	}
}

func BenchmarkFileMerge2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stream.FileMerge2()
	}
}

package p1_6

import (
	"testing"
)

type args struct {
	str string
}

var tests = []struct {
	name string
	args args
	want string
}{
	// TODO: Add test cases.
	{name: "case1", args: args{str: "aaaabbbcccccdddd"}, want: "a4b3c5d4"},
	{name: "case2", args: args{str: "aabbbcccccdddd"}, want: "a2b3c5d4"},
	{name: "case3", args: args{str: "abccd"}, want: "abccd"},
	{name: "case4", args: args{str: "aabbccdd"}, want: "aabbccdd"},
	{name: "case5", args: args{str: "aaabbccdd"}, want: "a3b2c2d2"},
	{name: "case6", args: args{str: "abbbcccccdddd"}, want: "a1b3c5d4"},
	{name: "case7", args: args{str: "abbbcccccd"}, want: "a1b3c5d1"},
}

func TestOriginalCompress(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OriginalCompress(tt.args.str); got != tt.want {
				t.Errorf("OriginalCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOriginalCompress(b *testing.B) {
	t := tests[0]
	for i := 0; i < b.N; i++ {
		OriginalCompress(t.args.str)
	}
}

func TestOriginalCompress2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OriginalCompress2(tt.args.str); got != tt.want {
				t.Errorf("OriginalCompress2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOriginalCompress2(b *testing.B) {
	t := tests[0]
	for i := 0; i < b.N; i++ {
		OriginalCompress2(t.args.str)
	}
}

func TestOriginalCompress3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OriginalCompress3(tt.args.str); got != tt.want {
				t.Errorf("OriginalCompress3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOriginalCompress3(b *testing.B) {
	t := tests[0]
	for i := 0; i < b.N; i++ {
		OriginalCompress3(t.args.str)
	}
}

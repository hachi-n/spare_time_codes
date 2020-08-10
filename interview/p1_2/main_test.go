package p1_2

import (
	"testing"
)

type args struct {
	str1 string
	str2 string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	// TODO: Add test cases.
	{name: "case1", args: args{"strings", "sstring"}, want: true},
	{name: "case2", args: args{"strings", "strings"}, want: false},
	{name: "case3", args: args{"strings", "srintgs"}, want: true},
	{name: "case4", args: args{"strings", "false"}, want: false},
	{name: "case5", args: args{"strings", "true"}, want: false},
	{name: "case6", args: args{"", "strings"}, want: false},
	{name: "case7", args: args{"strings", ""}, want: false},
	{name: "case8", args: args{"strings", "stringS"}, want: false},
}

func TestIsSortedCharacters(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSortedCharacters(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("IsSortedCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsSortedCharacters(b *testing.B) {
	args := tests[0].args

	for i := 0; i < b.N; i++ {
		IsSortedCharacters(args.str1, args.str2)
	}
}

func TestIsSortedCharacters2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSortedCharacters2(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("IsSortedCharacters2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsSortedCharacters2(b *testing.B) {
	args := tests[0].args

	for i := 0; i < b.N; i++ {
		IsSortedCharacters2(args.str1, args.str2)
	}
}


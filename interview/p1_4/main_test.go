package p1_4

import "testing"

type args struct {
	str string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	// TODO: Add test cases.
	{name: "case1", args: args{str: "Tact Coa"}, want: true},
	{name: "case2", args: args{str: "mom"}, want: true},
	{name: "case3", args: args{str: "moma"}, want: false},
	{name: "case4", args: args{str: "Now I won"}, want: true},
	{name: "case5", args: args{str: "Brwo orr oobr"}, want: true},
	{name: "case5", args: args{str: "Brw orr oobr"}, want: false},
	{name: "case6", args: args{str: "Pul ul ip If ull upp"}, want: true},
	{name: "case7", args: args{str: "tactcoapapa"}, want: true},
}

func TestIsPermutation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermutation(tt.args.str); got != tt.want {
				t.Errorf("IsPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

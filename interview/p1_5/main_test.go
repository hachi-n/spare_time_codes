package p1_5

import "testing"

type args struct {
	baseStr   string
	editedStr string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	// TODO: Add test cases.
	{name: "case1", args: args{baseStr: "pale", editedStr: "ple"}, want: true},
	{name: "case2", args: args{baseStr: "pales", editedStr: "pale"}, want: true},
	{name: "case3", args: args{baseStr: "pale", editedStr: "pales"}, want: true},
	{name: "case4", args: args{baseStr: "pale", editedStr: "bale"}, want: true},
	{name: "case5", args: args{baseStr: "pale", editedStr: "bake"}, want: false},
	{name: "case6", args: args{baseStr: "paale", editedStr: "pale"}, want: true},
	{name: "case6", args: args{baseStr: "ppaale", editedStr: "pale"}, want: false},
}

func TestIsEditedString(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEditedString(tt.args.baseStr, tt.args.editedStr); got != tt.want {
				t.Errorf("IsEditedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

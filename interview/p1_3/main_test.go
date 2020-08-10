package p1_3

import (
	"testing"
)

type args struct {
	str        string
	realLength int
}

var tests = []struct {
	name string
	args args
	want string
}{
	{name: "case1", args: args{str: "Mr John Smith ", realLength: 13}, want: "Mr%20John%20Smith"},
	{name: "case1", args: args{str: "Mr John  Smith ", realLength: 14}, want: "Mr%20John%20%20Smith"},
	{name: "case1", args: args{str: "Mr  John Smith ", realLength: 14}, want: "Mr%20%20John%20Smith"},
	{name: "case1", args: args{str: "Mr John Smith   ", realLength: 13}, want: "Mr%20John%20Smith"},
	{name: "case1", args: args{str: " Mr John Smith ", realLength: 13}, want: "%20Mr%20John%20Smith"},
}

func TestConvertSpace(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertSpace(tt.args.str, tt.args.realLength); got != tt.want {
				t.Errorf("ConvertSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertSpace2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertSpace2(tt.args.str, tt.args.realLength); got != tt.want {
				t.Errorf("ConvertSpace2() = %v, want %v", got, tt.want)
			}
		})
	}
}

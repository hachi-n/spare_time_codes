package p1_1

import (
	"testing"
)

func BenchmarkIsDuplicationString(b *testing.B) {
	testString := "strings"

	for i := 0; i < b.N; i++ {
		IsDuplicationString(testString)
	}
}

func BenchmarkIsDuplicationString2(b *testing.B) {
	testString := "strings"

	for i := 0; i < b.N; i++ {
		IsDupulicationString2(testString)
	}
}

func BenchmarkIsDuplicationString3(b *testing.B) {
	testString := "strings"

	for i := 0; i < b.N; i++ {
		IsDupulicationString3(testString)
	}
}

func BenchmarkIsDuplicationString4(b *testing.B) {
	testString := "strings"

	for i := 0; i < b.N; i++ {
		IsDupulicationString4(testString)
	}
}

func Test_isDuplicationString(t *testing.T) {
	type args struct{ s string }
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{s: "strings"}, want: true},
		{name: "case2", args: args{s: "string"}, want: false},
		{name: "case3", args: args{s: "hoge"}, want: false},
		{name: "case4", args: args{s: "hogehoge"}, want: true},
		{name: "case5", args: args{s: "fjksaldfjdak;sdfl"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDuplicationString(tt.args.s); got != tt.want {
				t.Errorf("isDuplicationString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDupulicationString2(t *testing.T) {
	type args struct{ s string }
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{s: "strings"}, want: true},
		{name: "case2", args: args{s: "string"}, want: false},
		{name: "case3", args: args{s: "hoge"}, want: false},
		{name: "case4", args: args{s: "hogehoge"}, want: true},
		{name: "case5", args: args{s: "fjksaldfjdak;sdfl"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDupulicationString2(tt.args.s); got != tt.want {
				t.Errorf("isDupulicationString2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDupulicationString3(t *testing.T) {
	type args struct{ s string }
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{s: "strings"}, want: true},
		{name: "case2", args: args{s: "string"}, want: false},
		{name: "case3", args: args{s: "hoge"}, want: false},
		{name: "case4", args: args{s: "hogehoge"}, want: true},
		{name: "case5", args: args{s: "fjksaldfjdak;sdfl"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDupulicationString3(tt.args.s); got != tt.want {
				t.Errorf("isDupulicationString3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDupulicationString4(t *testing.T) {
	type args struct{ s string }
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{s: "strings"}, want: true},
		{name: "case2", args: args{s: "string"}, want: false},
		{name: "case3", args: args{s: "hoge"}, want: false},
		{name: "case4", args: args{s: "hogehoge"}, want: true},
		{name: "case5", args: args{s: "fjksaldfjdak;sdfl"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDupulicationString4(tt.args.s); got != tt.want {
				t.Errorf("isDupulicationString4() = %v, want %v", got, tt.want)
			}
		})
	}
}

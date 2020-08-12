package p1_8

import (
	"reflect"
	"testing"
)

type args struct {
	matrix [][]int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	// TODO: Add test cases.
	{name: "case1",
		args: args{
			matrix: [][]int{
				{0, 1, 2, 3, 4, 5, 6},
				{1, 2, 3, 4, 5, 6, 7},
				{1, 2, 3, 4, 5, 6, 7},
				{1, 2, 3, 4, 5, 6, 7},
				{1, 2, 3, 4, 5, 6, 7},
			},
		},
		want: [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 2, 3, 4, 5, 6, 7},
			{0, 2, 3, 4, 5, 6, 7},
			{0, 2, 3, 4, 5, 6, 7},
			{0, 2, 3, 4, 5, 6, 7},
		},
	},
}

func TestConvertMatrix(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

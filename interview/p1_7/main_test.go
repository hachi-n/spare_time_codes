package p1_7

import (
	"reflect"
	"testing"
)

type args struct {
	top    []string
	right  []string
	bottom []string
	left   []string
}

var tests = []struct {
	name  string
	args  args
	want  []string
	want1 []string
	want2 []string
	want3 []string
}{
	// TODO: Add test cases.
	{name: "case1",
		args: args{
			top:    []string{"a", "b", "c"},
			right:  []string{"d", "e", "f"},
			bottom: []string{"g", "h", "i"},
			left:   []string{"j", "k", "l"},
		},
		want:  []string{"j", "k", "l"},
		want1: []string{"a", "b", "c"},
		want2: []string{"d", "e", "f"},
		want3: []string{"g", "h", "i"},
	},
}

func TestRotate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := Rotate(tt.args.top, tt.args.right, tt.args.bottom, tt.args.left)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Rotate() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Rotate() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("Rotate() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

// Because Mem Copy.
//func TestRotate1(t *testing.T) {
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1, got2, got3 := Rotate1(tt.args.top, tt.args.right, tt.args.bottom, tt.args.left)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Rotate1() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("Rotate1() got1 = %v, want %v", got1, tt.want1)
//			}
//			if !reflect.DeepEqual(got2, tt.want2) {
//				t.Errorf("Rotate1() got2 = %v, want %v", got2, tt.want2)
//			}
//			if !reflect.DeepEqual(got3, tt.want3) {
//				t.Errorf("Rotate1() got3 = %v, want %v", got3, tt.want3)
//			}
//		})
//	}
//}

package leetcode

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "all zero",
			args: args{
				matrix: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "0 in mid",
			args: args{
				matrix: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "test 3",
			args: args{
				matrix: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{-3, -2, -1, 0, 0, 1, 2, 3},
				target: 0,
			},
			want: [][]int{
				{-3, -2, 2, 3},
				{-3, -1, 1, 3},
				{-3, 0, 0, 3},
				{-3, 0, 1, 2},
				{-2, -1, 0, 3},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name: "3",
			args: args{
				nums:   []int{-1, -5, -5, -3, 2, 5, 0, 4},
				target: -7,
			},
			want: [][]int{
				{-5, -5, -1, 4},
				{-5, -3, -1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

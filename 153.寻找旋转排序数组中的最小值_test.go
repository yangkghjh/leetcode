package leetcode

import (
	"math"
	"testing"
)

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,4,5,1,2]",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "[4,5,6,7,0,1,2]",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "[0,1,2,4,5,6,7]",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 6, 7},
			},
			want: 0,
		},
		{
			name: "[]",
			args: args{
				nums: []int{},
			},
			want: math.MinInt64,
		},
		{
			name: "[1]",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

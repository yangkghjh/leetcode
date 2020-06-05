package leetcode

import (
	"reflect"
	"testing"
)

func Test_getNextStartPoint(t *testing.T) {
	type args struct {
		m int
		n int
		i int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "3*3,2",
			args: args{
				m: 3,
				n: 3,
				i: 2,
			},
			want:  2,
			want1: 0,
		},
		{
			name: "3*3,3",
			args: args{
				m: 3,
				n: 3,
				i: 3,
			},
			want:  1,
			want1: 2,
		},
		{
			name: "3*3,1",
			args: args{
				m: 3,
				n: 3,
				i: 1,
			},
			want:  0,
			want1: 1,
		},
		{
			name: "1*1,0",
			args: args{
				m: 1,
				n: 1,
				i: 0,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "3*2,3",
			args: args{
				m: 3,
				n: 2,
				i: 3,
			},
			want:  2,
			want1: 1,
		},
		{
			name: "2*3,3",
			args: args{
				m: 2,
				n: 3,
				i: 3,
			},
			want:  1,
			want1: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getNextStartPoint(tt.args.m, tt.args.n, tt.args.i)
			if got != tt.want {
				t.Errorf("getNextStartPoint() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getNextStartPoint() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[]",
			args: args{
				matrix: [][]int{},
			},
			want: []int{},
		},
		{
			name: "3*3",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "2,3",
			args: args{
				matrix: [][]int{
					{2, 3},
				},
			},
			want: []int{2, 3},
		},
		{
			name: "[2],[3]",
			args: args{
				matrix: [][]int{
					{2},
					{3},
				},
			},
			want: []int{2, 3},
		},
		{
			name: "[2],[3],[4]",
			args: args{
				matrix: [][]int{
					{2},
					{3},
					{4},
				},
			},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

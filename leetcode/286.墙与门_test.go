package leetcode

import (
	"reflect"
	"testing"
)

func Test_wallsAndGates(t *testing.T) {
	const INF = 2147483647

	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				rooms: [][]int{
					{INF, -1, 0, INF},
					{INF, INF, INF, -1},
					{INF, -1, INF, -1},
					{0, -1, INF, INF},
				},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				rooms: [][]int{
					{INF, -1, INF, INF},
					{INF, INF, INF, -1},
					{INF, -1, INF, -1},
					{INF, -1, INF, INF},
				},
			},
			want: [][]int{
				{INF, -1, INF, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{INF, -1, INF, INF},
			},
		},
		{
			name: "3",
			args: args{
				rooms: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "3",
			args: args{
				rooms: [][]int{{}},
			},
			want: [][]int{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallsAndGates(tt.args.rooms)
			if !reflect.DeepEqual(tt.args.rooms, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", tt.args.rooms, tt.want)
			}
		})
	}
}

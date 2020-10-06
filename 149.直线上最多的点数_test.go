package leetcode

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				points: [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				points: [][]int{{0, 0}},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				points: [][]int{},
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				points: [][]int{{0, 0}, {0, 0}},
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				points: [][]int{{0, 0}, {1, 1}, {0, 0}},
			},
			want: 3,
		},
		{
			name: "7",
			args: args{
				points: [][]int{{1, 1}, {1, 1}, {2, 3}},
			},
			want: 3,
		},
		{
			name: "8",
			args: args{
				points: [][]int{{1, 1}, {1, 1}, {0, 0}, {3, 4}, {4, 5}, {5, 6}, {7, 8}, {8, 9}},
			},
			want: 5,
		},
		{
			name: "9",
			args: args{
				points: [][]int{{1, 1}, {2, 1}, {2, 2}, {1, 4}, {3, 3}},
			},
			want: 3,
		},
		{
			name: "10",
			args: args{
				points: [][]int{{0, 0}, {94911151, 94911150}, {94911152, 94911151}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

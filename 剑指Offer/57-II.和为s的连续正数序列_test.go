package offer

import (
	"reflect"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
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
				target: 1,
			},
			want: [][]int{[]int{1}},
		},
		{
			name: "2",
			args: args{
				target: 9,
			},
			want: [][]int{[]int{2, 3, 4}, []int{4, 5}},
		},
		{
			name: "3",
			args: args{
				target: 15,
			},
			want: [][]int{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6}, []int{7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

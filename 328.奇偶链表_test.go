package leetcode

import (
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1, 2, 3, 4, 5]",
			args: args{
				head: NewList(1, 2, 3, 4, 5),
			},
			want: []int{1, 3, 5, 2, 4},
		},
		{
			name: "[]",
			args: args{
				head: NewList(),
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oddEvenList(tt.args.head)

			for _, v := range tt.want {
				if v != got.Val {
					t.Errorf("oddEvenList() = %v, want %v", got.Val, v)
					break
				}
				got = got.Next
			}
		})
	}
}

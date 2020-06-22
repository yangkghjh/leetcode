package leetcode

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil",
			args: args{
				head: nil,
				val:  0,
			},
			want: []int{},
		},
		{
			name: "[1, 2, 6, 3, 4, 5, 6], 6",
			args: args{
				head: NewList(1, 2, 6, 3, 4, 5, 6),
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[1, 2, 6, 3, 4, 5, 6], 1",
			args: args{
				head: NewList(1, 2, 6, 3, 4, 5, 6),
				val:  1,
			},
			want: []int{2, 6, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.args.head, tt.args.val)
			for _, v := range tt.want {
				if v != got.Val {
					t.Errorf("removeElements() = %v, want %v", got.Val, v)
					break
				}
				got = got.Next
			}
		})
	}
}

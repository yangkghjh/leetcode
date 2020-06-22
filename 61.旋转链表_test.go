package leetcode

import (
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1->2->3->4->5->NULL, k = 2",
			args: args{
				head: NewList(1, 2, 3, 4, 5),
				k:    2,
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "0->1->2->NULL, k = 4",
			args: args{
				head: NewList(0, 1, 2),
				k:    4,
			},
			want: []int{2, 0, 1},
		},
		{
			name: "NULL, k = 4",
			args: args{
				head: NewList(),
				k:    4,
			},
			want: []int{},
		},
		{
			name: "1, k = 4",
			args: args{
				head: NewList(1),
				k:    4,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(tt.args.head, tt.args.k)
			for _, v := range tt.want {
				if v != got.Val {
					t.Errorf("rotateRight() = %v, want %v", got.Val, v)
				}

				got = got.Next
			}

		})
	}
}

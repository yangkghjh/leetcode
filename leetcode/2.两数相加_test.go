package leetcode

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "(2 -> 4 -> 3) + (5 -> 6 -> 4)",
			args: args{
				l1: NewList(2, 4, 3),
				l2: NewList(5, 6, 4),
			},
			want: NewList(7, 0, 8),
		},
		{
			name: "(2 -> 4 -> 3 -> 1) + (5 -> 6 -> 4)",
			args: args{
				l1: NewList(2, 4, 3, 1),
				l2: NewList(5, 6, 4),
			},
			want: NewList(7, 0, 8, 1),
		},
		{
			name: "(2 -> 4 -> 3 -> 1) + ()",
			args: args{
				l1: NewList(2, 4, 3, 1),
				l2: nil,
			},
			want: NewList(2, 4, 3, 1),
		},
		{
			name: "(5) + (5)",
			args: args{
				l1: NewList(5),
				l2: NewList(5),
			},
			want: NewList(0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			w := tt.want
			for w != nil {
				if got.Val != w.Val {
					t.Errorf("addTwoNumbers() = %v, want %v", got.Val, w.Val)
					break
				}
				got = got.Next
				w = w.Next
			}
		})
	}
}

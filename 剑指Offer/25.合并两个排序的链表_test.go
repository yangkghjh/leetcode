package offer

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			name: "1",
			args: args{
				l1: NewList(1, 2, 4),
				l2: NewList(1, 3, 4),
			},
			want: NewList(1, 1, 2, 3, 4, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.l1, tt.args.l2)
			w := tt.want
			for got != nil {
				if got.Val != w.Val {
					t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
				w = w.Next
			}
		})
	}
}

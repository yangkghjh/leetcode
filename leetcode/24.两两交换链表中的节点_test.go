package leetcode

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: NewList(1, 2, 3, 4),
			},
			want: NewList(2, 1, 4, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.args.head)
			want := tt.want
			for want != nil {
				if want.Val != got.Val {
					t.Errorf("swapPairs() = %v, want %v", got, tt.want)
					break
				}
				want = want.Next
				got = got.Next
			}
		})
	}
}

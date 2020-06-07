package leetcode

import (
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,4,5],2,4",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				m: 2,
				n: 4,
			},
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "[3,5],1,2",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
				m: 1,
				n: 2,
			},
			want: []int{5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); got != nil {
				for i := 0; i < len(tt.want); i++ {
					if tt.want[i] != got.Val {
						t.Errorf("step = %v, reverseBetween() = %v, want %v", i, got.Val, tt.want)
					}
					got = got.Next
				}
			}
		})
	}
}

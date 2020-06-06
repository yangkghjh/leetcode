package leetcode

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	l1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l1.Next.Next.Next.Next = l1

	l2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l2.Next.Next.Next.Next = l2.Next

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "nil",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "0",
			args: args{
				head: &ListNode{Val: 0},
			},
			want: nil,
		},
		{
			name: "0,1,2,3",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "0,1,2,3,0",
			args: args{
				head: l1,
			},
			want: l1,
		},
		{
			name: "0,1,2,3,0",
			args: args{
				head: l2,
			},
			want: l2.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

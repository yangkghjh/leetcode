package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
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

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			name: "0",
			args: args{
				head: &ListNode{Val: 0},
			},
			want: false,
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
			want: false,
		},
		{
			name: "0,1,2,3,0",
			args: args{
				head: l1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

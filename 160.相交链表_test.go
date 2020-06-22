package leetcode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	test1A := NewList(4, 1)
	test1B := NewList(5, 0, 1)
	test1C := NewList(8, 4, 5)
	test1A.Next.Next = test1C
	test1B.Next.Next.Next = test1C

	test2A := NewList(4, 1)
	test2B := NewList(5, 0, 1)

	test4A := NewList(5, 0, 1)
	test4B := NewList(4, 1)
	test4C := NewList(8, 4, 5)
	test4A.Next.Next.Next = test4C
	test4B.Next.Next = test4C

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				headA: test1A,
				headB: test1B,
			},
			want: test1C,
		},
		{
			name: "test2",
			args: args{
				headA: nil,
				headB: nil,
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				headA: test2A,
				headB: test2B,
			},
			want: nil,
		},
		{
			name: "test4",
			args: args{
				headA: test4A,
				headB: test4B,
			},
			want: test4C,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

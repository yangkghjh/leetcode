package offer

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	a := NewList(1, 2, 3)
	b := NewList(4, 5)
	c := NewList(6, 7)
	a.Next.Next.Next = c
	b.Next.Next = c

	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				headA: a,
				headB: b,
			},
			want: c,
		},
		{
			name: "2",
			args: args{
				headA: NewList(1, 2),
				headB: NewList(3, 4),
			},
			want: nil,
		},
		{
			name: "3",
			args: args{
				headA: b,
				headB: a,
			},
			want: c,
		},
		{
			name: "2",
			args: args{
				headA: NewList(1, 2),
			},
			want: nil,
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

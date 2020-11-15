package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1->2->3->3->4->4->5",
			args: args{
				head: NewList(1, 2, 3, 3, 4, 4, 5),
			},
			want: NewList(1, 2, 5),
		},
		{
			name: "1->2->3->3->3->4->4->5",
			args: args{
				head: NewList(1, 2, 3, 3, 3, 4, 4, 5),
			},
			want: NewList(1, 2, 5),
		},
		{
			name: "1->1->1->2->3",
			args: args{
				head: NewList(1, 1, 1, 2, 3),
			},
			want: NewList(2, 3),
		},
		{
			name: "1->1->1->2->3->3",
			args: args{
				head: NewList(1, 1, 1, 2, 3, 3),
			},
			want: NewList(2),
		},
		{
			name: "1->1->1->2->2->3->3",
			args: args{
				head: NewList(1, 1, 1, 2, 2, 3, 3),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.args.head)
			want := tt.want

			gotnums := []int{}
			wantnums := []int{}
			w := want
			g := got
			for w != nil {
				wantnums = append(wantnums, w.Val)
				w = w.Next
			}
			for g != nil {
				gotnums = append(gotnums, g.Val)
				g = g.Next
			}
			fmt.Println(gotnums, wantnums)

			if !reflect.DeepEqual(gotnums, wantnums) {
				t.Errorf("deleteDuplicates() = %v, want %v", gotnums, wantnums)
			}
		})
	}
}

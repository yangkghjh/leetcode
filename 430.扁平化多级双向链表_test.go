package leetcode

import (
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *FlattenNode
	}

	node1 := NewFlattenList(1, 2, 3, 4, 5, 6)
	node2 := NewFlattenList(7, 8, 9, 10)
	node3 := NewFlattenList(11, 12)
	node1.Next.Next.Child = node2
	node2.Next.Child = node3

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				root: node1,
			},
			want: []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flatten(tt.args.root)
			for _, v := range tt.want {
				if v != got.Val {
					t.Errorf("flattenChildOrTail() = %v, want %v", got.Val, v)
					break
				}
				got = got.Next
			}
		})
	}
}

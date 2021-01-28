package leetcode

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree1 := NewBinaryTree(6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5)
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: tree1,
				p:    tree1.Left,
				q:    tree1.Right,
			},
			want: tree1,
		},
		{
			name: "2",
			args: args{
				root: tree1,
				p:    tree1.Right,
				q:    tree1.Left,
			},
			want: tree1,
		},
		{
			name: "3",
			args: args{
				root: tree1,
				p:    tree1.Left,
				q:    tree1.Left.Right,
			},
			want: tree1.Left,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

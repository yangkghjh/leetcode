package leetcode

import (
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: NewBinaryTree(4, 2, 7, 1, 3),
				val:  5,
			},
			want: NewBinaryTree(4, 2, 7, 1, 3, 5),
		},
		{
			name: "2",
			args: args{
				root: NewBinaryTree(),
				val:  5,
			},
			want: NewBinaryTree(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertIntoBST(tt.args.root, tt.args.val)
			if !TreeNodeDeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

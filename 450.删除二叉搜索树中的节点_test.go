package leetcode

import (
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: NewBinaryTree(5, 3, 6, 2, 4, -1, 7),
				key:  3,
			},
			want: NewBinaryTree(5, 4, 6, 2, -1, -1, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode(tt.args.root, tt.args.key)
			if !TreeNodeDeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

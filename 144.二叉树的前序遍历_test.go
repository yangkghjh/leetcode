package leetcode

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	// NewTreeNode()
	node0 := NewTreeNode(1)
	node1 := NewTreeNode(2)
	node2 := NewTreeNode(3)
	node0.Right = node1
	node1.Left = node2

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,null,2,3]",
			args: args{
				root: NewBinaryTree(1, -1, 2, 3),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

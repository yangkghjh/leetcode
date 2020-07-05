package leetcode

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	tree1 := NewBinaryTree(3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4)
	tree2 := NewBinaryTree(1, 2)

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
				q:    tree1.Left.Right.Right,
			},
			want: tree1.Left,
		},
		{
			name: "2",
			args: args{
				root: tree2,
				p:    tree2,
				q:    tree2.Left,
			},
			want: tree2,
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

func Test_findBinaryTreeTrace(t *testing.T) {
	type args struct {
		node   *TreeNode
		target *TreeNode
	}

	tree1 := NewBinaryTree(1, 2)
	tree2 := NewBinaryTree(3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4)

	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "1",
			args: args{
				node:   tree1,
				target: tree1.Left,
			},
			want: []*TreeNode{tree1.Left, tree1},
		},
		{
			name: "2",
			args: args{
				node:   tree1,
				target: tree1,
			},
			want: []*TreeNode{tree1},
		},
		{
			name: "3",
			args: args{
				node:   tree2,
				target: tree2.Left,
			},
			want: []*TreeNode{tree2.Left, tree2},
		},
		{
			name: "4",
			args: args{
				node:   tree2,
				target: tree2.Left.Right,
			},
			want: []*TreeNode{tree2.Left.Right, tree2.Left, tree2},
		},
		{
			name: "5",
			args: args{
				node:   tree2,
				target: tree2.Left.Right.Right,
			},
			want: []*TreeNode{tree2.Left.Right.Right, tree2.Left.Right, tree2.Left, tree2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBinaryTreeTrace(tt.args.node, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBinaryTreeTrace() = %v, want %v", got, tt.want)
			}
		})
	}
}

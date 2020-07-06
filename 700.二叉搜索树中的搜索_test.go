package leetcode

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}

	tree1 := NewBinaryTree(4, 2, 7, 1, 3)
	tree2 := NewBinaryTree(2, 1, 3)

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: tree1,
				val:  2,
			},
			want: tree1.Left,
		},
		{
			name: "2",
			args: args{
				root: tree2,
				val:  5,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
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
			want: []int{3, 2, 1},
		},
		{
			name: "[3,1,2]",
			args: args{
				root: NewBinaryTree(3, 1, 2),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

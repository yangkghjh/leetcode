package leetcode

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				root: NewBinaryTree(3, 9, 20, -1, -1, 15, 7),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				root: NewBinaryTree(1, 2, 2, 3, 3, -1, -1, 4, 4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
